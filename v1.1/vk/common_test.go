package vk

import (
	"sync"
	"testing"
)

func TestStackAllocator(t *testing.T) {
	al := stackAllocator{}
	al.init(1024)
	defer al.deinit()

	p0 := al.alloc(100)
	p1 := al.alloc(300)
	p2 := al.alloc(500)
	if len(al.ptrs) != 0 {
		t.Error()
	}
	p3 := al.alloc(800)
	if len(al.ptrs) != 1 {
		t.Error()
	}

	al.free(p3)
	al.free(p2)
	al.free(p1)
	al.free(p0)

	if al.offset != 0 {
		t.Error()
	}
}

func BenchmarkStackAllocator1MB(b *testing.B) {
	al := stackAllocator{}
	al.init(1024 * 1024) // 1 MB
	defer al.deinit()
	for n := 0; n < b.N; n++ {
		p := al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.free(p)
	}
}

func BenchmarkStackAllocator1Byte(b *testing.B) {
	al := stackAllocator{}
	al.init(1)
	defer al.deinit()
	for n := 0; n < b.N; n++ {
		p := al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.alloc(100)
		al.free(p)
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
	id := reg.register(obj0)
	obj1, _ := reg.access(id)
	if obj0 != obj1 {
		t.Error("Unxpected obj0 != obj1")
	}

	reg.unregister(id)
	_, ok := reg.access(id)
	if ok {
		t.Error("Unexpected ok")
	}
}

func TestObjectRegistryShard(t *testing.T) {
	reg := newObjectRegistry()

	for i := 0; i < 10000; i++ {
		reg.register("any")
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
		reg.register("value")
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
				reg.register("value")
				// id := reg.register("value")
				// reg.access(id)
				// reg.unregister(id)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
