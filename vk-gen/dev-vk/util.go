package vk

/*
#include "stdlib.h"
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"math/rand"
	"sync"
	"unsafe"
)

var registry objectRegistry
var pool stackAllocatorPool

func init() {
	registry = newObjectRegistry()
	pool.init(16)
}

type allocator interface {
	alloc(s uint) unsafe.Pointer
	free(p unsafe.Pointer)
}

type sysAllocator struct{}

func (sa *sysAllocator) alloc(s uint) unsafe.Pointer {
	return C.malloc(C.size_t(s))
}

func (sa *sysAllocator) free(ptr unsafe.Pointer) {
	C.free(ptr)
}

/*
stackAllocator is a stack memory allocator using C memory.

Why?
C code can't access Go's *dynamic memory*, which means that it's discontinuity,
due to Cgo rules about pointer passing.
When a C function's accepts a *dynamic memory* as a parameter, we can pass a
stack memory or heap memory to it in C, but in Cgo, it's the only choice to pass
a heap memory, it takes costs.
So, a stack memory allocator on heap is needed.
*/
type stackAllocator struct {
	start  unsafe.Pointer
	offset uintptr
	total  uintptr

	ptrs []unsafe.Pointer
}

func (sa *stackAllocator) init(totalSize uint) {
	if sa.start != nil {
		sa.deinit()
	}
	sa.start = C.malloc(C.size_t(totalSize))
	sa.offset = 0
	sa.total = uintptr(totalSize)
}

//Deinit releases all allocated memories, which include the buffer and the others
//allocated by malloc().
func (sa *stackAllocator) deinit() {
	C.free(sa.start)
	sa.start = nil
	sa.offset = 0
	sa.total = 0

	for _, p := range sa.ptrs {
		C.free(p)
	}
	sa.ptrs = sa.ptrs[:0]
}

//Alloc allocate c memory(saignmented for any type) on the buffer,
//if the buffer is out of memory, using malloc() instead,
//but the allocated memories are still mamanged by this allocator.
func (sa *stackAllocator) alloc(s uint) unsafe.Pointer {
	size := uintptr(s)
	current := uintptr(sa.start) + sa.offset
	padding := calculatePaddingWithOneByteHeader(current, C.sizeof_intmax_t)

	if len(sa.ptrs) == 0 && sa.offset+padding+size < sa.total {
		next := current + padding
		header := next - 1

		headerPtr := (*byte)(unsafe.Pointer(header))
		*headerPtr = byte(padding)

		sa.offset += padding + size

		ptr := unsafe.Pointer(next)
		return ptr
	}

	//malloc
	ptr := C.malloc(C.size_t(s))
	sa.ptrs = append(sa.ptrs, ptr)
	return ptr
}

//Free all the memories allocated after the allocation of the passed pointer,
//includes those memories using malloc().
func (sa *stackAllocator) free(ptr unsafe.Pointer) {
	//malloc
	for i := len(sa.ptrs) - 1; i > 0; i-- {
		_ptr := sa.ptrs[i]
		C.free(_ptr)
		sa.ptrs = sa.ptrs[:i]
		if ptr == _ptr {
			return
		}
	}

	current := uintptr(ptr)
	header := current - 1

	headerPtr := (*byte)(unsafe.Pointer(header))
	padding := uintptr(*headerPtr)

	sa.offset = current - padding - uintptr(sa.start)
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

func newCString(s string, al allocator) *C.char {
	if s == "" {
		return nil
	}
	n := len(s) + 1
	p := al.alloc(uint(n))
	slice := (*[1 << 31]C.char)(p)[0:n]

	for i := 0; i < n-1; i++ {
		slice[i] = C.char(s[i])
	}
	slice[n-1] = 0

	_s := (*C.char)(p)
	return _s
}

func freeCString(s *C.char, al allocator) {
	al.free(unsafe.Pointer(s))
}

func newCStringArray(ss []string, al allocator) (**C.char, C.uint32_t) {
	n := len(ss)
	if n == 0 {
		return nil, 0
	}

	size := uintptr(n) * unsafe.Sizeof(uintptr(0))
	_ss := (**C.char)(al.alloc(uint(size)))
	slice := (*[1 << 31]*C.char)(unsafe.Pointer(_ss))[0:n]

	for i := 0; i < n; i++ {
		slice[i] = newCString(ss[i], al)
	}

	return _ss, C.uint32_t(n)
}

func freeCStringArray(ss **C.char, n C.uint32_t, al allocator) {
	_n := int(n)
	slice := (*[1 << 31]*C.char)(unsafe.Pointer(ss))[0:_n]

	for i := 0; i < int(_n); i++ {
		al.free(unsafe.Pointer(slice[i]))
	}
	al.free(unsafe.Pointer(ss))
}

type disposer []func()

func (d *disposer) add(disposeFn func()) {
	*d = append(*d, disposeFn)
}

func (d *disposer) dispose() {
	for _, fn := range *d {
		fn()
	}
}

type stackAllocatorPool struct {
	allocators []*stackAllocator
	max        int
	sync.Mutex
}

func (ap *stackAllocatorPool) init(max int) {
	ap.allocators = make([]*stackAllocator, 0, max)
	ap.max = max
}

func (ap *stackAllocatorPool) take() *stackAllocator {
	ap.Lock()
	defer ap.Unlock()
	n := len(ap.allocators)
	if n > 0 {
		sa := ap.allocators[n-1]
		ap.allocators = ap.allocators[:n-1]
		return sa
	} else {
		sa := &stackAllocator{}
		sa.init(1024 * 2)
		return sa
	}
}

func (ap *stackAllocatorPool) give(sa *stackAllocator) {
	ap.Lock()
	defer ap.Unlock()
	n := len(ap.allocators)
	if n < ap.max {
		ap.allocators = append(ap.allocators, sa)
	} else {
		sa.deinit()
	}
}

//ObjectRegistry keeps the mapping from id to go object.
type objectRegistry []*sharedObjectRegistry

type sharedObjectRegistry struct {
	objs map[uintptr]interface{}
	id   uintptr
	sync.RWMutex
}

func newObjectRegistry() objectRegistry {
	reg := make(objectRegistry, 256)
	for i := uintptr(0); i < 256; i++ {
		reg[i] = &sharedObjectRegistry{objs: make(map[uintptr]interface{})}
	}
	return reg
}

func (reg objectRegistry) getShard(id uintptr) *sharedObjectRegistry {
	return reg[id&0xff]
}

func (reg objectRegistry) register(value interface{}) uintptr {
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

func (reg objectRegistry) access(id uintptr) (interface{}, bool) {
	shard := reg.getShard(id)
	shard.RLock()
	val, ok := shard.objs[id]
	shard.RUnlock()
	return val, ok
}

func (reg objectRegistry) unregister(id uintptr) {
	shard := reg.getShard(id)
	shard.Lock()
	delete(shard.objs, id)
	shard.Unlock()
}
