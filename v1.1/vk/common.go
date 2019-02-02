package vk

/*
#include "stdlib.h"
#include "string.h"
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"bytes"
	"math/rand"
	"sync"
	"unsafe"
)

var Registry ObjectRegistry
var pool cmemoryPool

type Structure interface {
	GetNext() Structure
	SetNext(s Structure)
	sType() C.VkStructureType
	toCStructure(*cmemory) unsafe.Pointer
	fromCStructure(unsafe.Pointer)
}

func init() {
	Registry = newObjectRegistry()
	pool.init(16)
}

/*
cmemory is a stack memory allocator using C memory.

Why?
C memory is need.
C code can't access Go's *discontinuity memory*, due to the pointer passing rule of Cgo.
But, malloc() is expensive.
*/
type cmemory struct {
	start  unsafe.Pointer
	offset uintptr
	total  uintptr

	ptrs []unsafe.Pointer
}

func (m *cmemory) init(totalSize uint) {
	if m.start != nil {
		panic("Cmemory is already initialized")
	}

	m.start = C.malloc(C.size_t(totalSize))
	C.memset(m.start, 0, C.size_t(totalSize))
	m.offset = 0
	m.total = uintptr(totalSize)
	m.ptrs = nil
}

func (m *cmemory) dispose() {
	if m.start == nil {
		panic("Cmemory is uninitialized")
	}

	m.free()

	C.free(m.start)
	m.start = nil
	m.total = 0
}

//Alloc allocate c memory on the buffer,
//if the buffer is out of memory, using malloc() instead,
//but the allocated memories are still mamanged by this allocator.
func (m *cmemory) alloc(s uint) unsafe.Pointer {
	if s == 0 {
		panic("Allocate zero size memory")
	}
	size := uintptr(s)
	current := uintptr(m.start) + m.offset
	padding := calculatePaddingWithOneByteHeader(current, C.sizeof_intmax_t)

	if len(m.ptrs) == 0 && m.offset+padding+size < m.total {
		next := current + padding
		header := next - 1

		headerPtr := (*byte)(unsafe.Pointer(header))
		*headerPtr = byte(padding)

		m.offset += padding + size

		ptr := unsafe.Pointer(next)
		return ptr
	}

	//malloc
	ptr := C.malloc(C.size_t(s))
	m.ptrs = append(m.ptrs, ptr)
	return ptr
}

func (m *cmemory) free() {
	for i := range m.ptrs {
		C.free(m.ptrs[i])
	}
	m.ptrs = m.ptrs[:0]
	m.offset = 0
	C.memset(m.start, 0, C.size_t(m.total))
}

func calculatePadding(base, alignment uintptr) uintptr {
	multiplier := (base / alignment) + 1
	aligned := multiplier * alignment
	return aligned - base
}

func calculatePaddingWithOneByteHeader(base, alignment uintptr) uintptr {
	padding := calculatePadding(base, alignment)
	space := uintptr(1)

	if padding < space {
		space -= padding
		if space%alignment > 0 {
			padding += alignment * (1 + (space / alignment))
		} else {
			padding += alignment * (space / alignment)
		}
	}
	return padding
}

func toCString(s string, m *cmemory) *C.char {
	if s == "" {
		return nil
	}
	n := len(s)
	p := m.alloc(uint(n + 1))
	slice := (*[1 << 31]C.char)(p)[0 : n+1]

	for i := 0; i < n; i++ {
		slice[i] = C.char(s[i])
	}
	slice[n] = 0

	_s := (*C.char)(p)
	return _s
}

func toGoString(p *C.char) string {
	if p == nil {
		return ""
	}
	slice := (*[1 << 31]C.char)(unsafe.Pointer(p))
	var buffer bytes.Buffer

	for i := 0; ; i++ {
		if slice[i] == 0 {
			break
		}
		buffer.WriteByte(byte(slice[i]))
	}
	return buffer.String()
}

type cmemoryPool struct {
	cmemories []*cmemory
	max       int
	sync.Mutex
}

func (p *cmemoryPool) init(max int) {
	p.cmemories = make([]*cmemory, 0, max)
	p.max = max
}

func (p *cmemoryPool) take() *cmemory {
	p.Lock()
	defer p.Unlock()
	n := len(p.cmemories)
	if n > 0 {
		m := p.cmemories[n-1]
		p.cmemories = p.cmemories[:n-1]
		return m
	}
	m := &cmemory{}
	m.init(1024 * 2)
	return m
}

func (p *cmemoryPool) give(m *cmemory) {
	p.Lock()
	defer p.Unlock()
	n := len(p.cmemories)
	if n < p.max {
		m.free()
		p.cmemories = append(p.cmemories, m)
		return
	}
	m.dispose()
}

//ObjectRegistry keeps the mapping from id to go object.
type ObjectRegistry []*sharedObjectRegistry

type sharedObjectRegistry struct {
	objs map[uintptr]interface{}
	id   uintptr
	sync.RWMutex
}

func newObjectRegistry() ObjectRegistry {
	reg := make(ObjectRegistry, 256)
	for i := uintptr(0); i < 256; i++ {
		reg[i] = &sharedObjectRegistry{objs: make(map[uintptr]interface{})}
	}
	return reg
}

func (reg ObjectRegistry) getShard(id uintptr) *sharedObjectRegistry {
	return reg[id&0xff]
}

func (reg ObjectRegistry) Register(value interface{}) uintptr {
	shardN := uintptr(rand.Intn(256))
	shard := reg.getShard(shardN)
	var id uintptr
	shard.Lock()
	shard.id++
	id = shard.id<<8 + shardN
	shard.objs[id] = value
	shard.Unlock()
	return id
}

func (reg ObjectRegistry) Access(id uintptr) (interface{}, bool) {
	shard := reg.getShard(id)
	shard.RLock()
	val, ok := shard.objs[id]
	shard.RUnlock()
	return val, ok
}

func (reg ObjectRegistry) Unregister(id uintptr) {
	shard := reg.getShard(id)
	shard.Lock()
	delete(shard.objs, id)
	shard.Unlock()
}
