package vk

//#include "vulkan/vulkan.h"
import "C"

type Structure interface {
	GetNext() Structure
	SetNext(s Structure)
	sType() C.VkStructureType
}
