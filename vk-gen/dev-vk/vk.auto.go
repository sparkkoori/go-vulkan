package vk
//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"

func print() {
	C.print()
	return
}
func num() (ret int32) {
	c_ret := C.num()
	ret = int32(c_ret)
	return
}
func movePointer(a *int32) (ret **int32) {
	_sa := pool.take()
	defer pool.give(_sa)
	c_a := (*C.int)(_sa.alloc(C.sizeof_int))
	*c_a = C.int(*a)
	c_ret := C.movePointer(c_a)
	*a = int32(*c_a)
	ret = new(*int32)
	*ret = new(int32)
	**ret = int32(**c_ret)
	return
}
func setArray(a *int32) {
	_sa := pool.take()
	defer pool.give(_sa)
	c_a := (*C.int)(_sa.alloc(C.sizeof_int))
	*c_a = C.int(*a)
	C.setArray(c_a)
	*a = int32(*c_a)
	return
}
func setbigN() {
	C.setbigN()
	return
}
