package vk

//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"
import "unsafe"

func print() {
	C.print()
}
func num() (ret int32) {
	c_ret := C.num()
	ret = int32(c_ret)
	return
}
func movePointer(a *int32, b unsafe.Pointer) (ret **int32) {
	_sa := pool.take()
	defer pool.give(_sa)
	var c_a *C.int
	var c_b unsafe.Pointer
	{
		c_a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c_a = C.int(*a)
	}
	c_b = b
	c_ret := C.movePointer(c_a, c_b)
	{
		ret = new(*int32)
		{
			*ret = new(int32)
			**ret = int32(**c_ret)
		}
	}
	*a = int32(*c_a)
	return
}
func setArray(a *int32) {
	_sa := pool.take()
	defer pool.give(_sa)
	var c_a *C.int
	{
		c_a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c_a = C.int(*a)
	}
	C.setArray(c_a)
	*a = int32(*c_a)
}

type bigN uint32

func setbigN(n bigN) {
	var c_n C.bigN
	{
		var _temp C.uint
		_temp = C.uint((uint32)(n))
		c_n = C.bigN(_temp)
	}
	C.setbigN(c_n)
}
func setFn(fn *[0]byte) {
	var c_fn *[0]byte
	c_fn = fn
	C.setFn(c_fn)
}

type FUNC *[0]byte

func changeFunc(fun FUNC) (ret FUNC) {
	var c_fun C.FUNC
	{
		var _temp *[0]byte
		_temp = (*[0]byte)(fun)
		c_fun = C.FUNC(_temp)
	}
	c_ret := C.changeFunc(c_fun)
	{
		var _temp *[0]byte
		_temp = (*[0]byte)(c_ret)
		ret = FUNC(_temp)
	}
	return
}
