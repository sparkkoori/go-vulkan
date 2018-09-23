package vk
//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"

func print() {
	C.print()
}
func num() (_ret int32) {
	var c struct{ _ret C.int }
	c._ret = C.num()
	_ret = int32(c._ret)
	return
}
