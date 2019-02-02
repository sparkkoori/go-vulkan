package vk

import (
	"sync"
	"testing"
	"unsafe"
)

func TestCMemory(t *testing.T) {
	m := cmemory{}
	m.init(1024)
	defer m.dispose()

	for i := 0; i < 20; i++ {
		p := (*[10]int)(m.alloc(uint(unsafe.Sizeof([10]int{}))))
		for j := range p {
			(*p)[j] = 123
		}
	}
	if len(m.ptrs) == 0 {
		t.Error()
	}

	m.free()
	if m.offset != 0 || len(m.ptrs) != 0 {
		t.Error()
	}
}

func BenchmarkCMemory1MB(b *testing.B) {
	m := cmemory{}
	m.init(1024 * 1024) // 1 MB
	defer m.dispose()
	for n := 0; n < b.N; n++ {
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.free()
	}
}

func BenchmarkCMemory1Byte(b *testing.B) {
	m := cmemory{}
	m.init(1)
	defer m.dispose()
	for n := 0; n < b.N; n++ {
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.alloc(100)
		m.free()
	}
}

func TestCalculatePadding(t *testing.T) {
	pd := calculatePadding(100, 32)
	if pd != 28 { // 128 -100
		t.Error(pd)
	}
}

func TestCalculatePaddingWithOneByteHeader(t *testing.T) {
	pd := calculatePaddingWithOneByteHeader(127, 32)
	if pd != 1 {
		t.Errorf("Expected 1, got %d.\n", pd)
	}

	pd = calculatePaddingWithOneByteHeader(128, 32)
	if pd != 32 {
		t.Errorf("Expected 32, got %d.\n", pd)
	}
}

func TestObjectRegistry(t *testing.T) {
	reg := newObjectRegistry()
	obj0 := "any object"
	id := reg.Register(obj0)
	obj1, _ := reg.Access(id)
	if obj0 != obj1 {
		t.Error("Unxpected obj0 != obj1")
	}

	reg.Unregister(id)
	_, ok := reg.Access(id)
	if ok {
		t.Error("Unexpected ok")
	}
}

func TestObjectRegistryShard(t *testing.T) {
	reg := newObjectRegistry()

	for i := 0; i < 10000; i++ {
		reg.Register("any")
	}

	for _, s := range reg {
		if len(s.objs) == 0 {
			t.Fatal("Shard probaly dosn't work")
		}
	}
}

func BenchmarkObjectRegistryGo1(b *testing.B) {
	reg := newObjectRegistry()
	for i := 0; i < b.N; i++ {
		reg.Register("value")
		// id := reg.register("value")
		// reg.access(id)
		// reg.unregister(id)
	}
}

func BenchmarkSyncedMapGo1(b *testing.B) {
	m := make(map[int]interface{})
	l := sync.RWMutex{}
	for i := 0; i < b.N; i++ {
		l.Lock()
		m[i] = "value"
		l.Unlock()
	}
}

func BenchmarkSyncedMapGo100(b *testing.B) {
	m := make(map[int]interface{})
	l := sync.RWMutex{}
	aid := 0
	wg := sync.WaitGroup{}
	wg.Add(100)
	for j := 0; j < 100; j++ {
		go func() {
			for i := 0; i < b.N/100; i++ {
				l.Lock()
				aid++
				id := aid<<8 + i
				m[id] = "value"
				l.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkObjectRegistryGo100(b *testing.B) {
	reg := newObjectRegistry()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for j := 0; j < 100; j++ {
		go func() {
			for i := 0; i < b.N/100; i++ {
				reg.Register("value")
				// id := reg.register("value")
				// reg.access(id)
				// reg.unregister(id)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
