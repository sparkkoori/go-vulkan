package vk

//#include "vulkan/vulkan.h"
import "C"

type Structure interface {
	sType() C.VkStructureType
}
