package vk

import (
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
