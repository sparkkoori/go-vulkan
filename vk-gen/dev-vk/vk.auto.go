package vk

//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"
import "unsafe"

func print() {
	C.print()
}
func num() (_ret int32) {
	var c struct{ _ret C.int }
	c._ret = C.num()
	_ret = int32(c._ret)
	return
}
func movePointer(a *int32, b unsafe.Pointer) (_ret **int32) {
	var c struct {
		a    *C.int
		b    unsafe.Pointer
		_ret **C.int
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.a = C.int(*a)
	}
	c.b = b
	c._ret = C.movePointer(c.a, c.b)
	{
		_ret = new(*int32)
		{
			*_ret = new(int32)
			**_ret = int32(**c._ret)
		}
	}
	*a = int32(*c.a)
	return
}
func setArray(a *int32) {
	var c struct{ a *C.int }
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.a = C.int(*a)
	}
	C.setArray(c.a)
	*a = int32(*c.a)
}

type bigN uint32

func setbigN(n bigN) {
	var c struct{ n C.bigN }
	{
		var _temp C.uint
		_temp = C.uint((uint32)(n))
		c.n = C.bigN(_temp)
	}
	C.setbigN(c.n)
}
func setFn(fn *[0]byte) {
	var c struct{ fn *[0]byte }
	c.fn = fn
	C.setFn(c.fn)
}

type FUNC *[0]byte

func changeFunc(fun FUNC) (_ret FUNC) {
	var c struct {
		fun  C.FUNC
		_ret C.FUNC
	}
	{
		var _temp *[0]byte
		_temp = (*[0]byte)(fun)
		c.fun = C.FUNC(_temp)
	}
	c._ret = C.changeFunc(c.fun)
	{
		var _temp *[0]byte
		_temp = (*[0]byte)(c._ret)
		_ret = FUNC(_temp)
	}
	return
}
