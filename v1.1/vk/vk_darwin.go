package vk

/*
#include "bridges_darwin.h"
typedef void * void_pointer;
*/
import "C"
import "unsafe"

const MVK_MACOS_SURFACE_SPEC_VERSION uint32 = 2
const MVK_MACOS_SURFACE_EXTENSION_NAME string = "VK_MVK_macos_surface"

type PFNCreateMacOSSurfaceMVK uintptr
type MacOSSurfaceCreateFlagsMVK Flags
type MacOSSurfaceCreateInfoMVK struct {
	Next  Structure
	Flags MacOSSurfaceCreateFlagsMVK
	View  unsafe.Pointer
}

func (g *MacOSSurfaceCreateInfoMVK) toC(c *C.VkMacOSSurfaceCreateInfoMVK, _sa *stackAllocator) {
	c.sType = g.sType()
	if g.Next != nil {
		c.pNext = g.Next.toCStructure(_sa)
	}
	{
		var temp_in_VkMacOSSurfaceCreateFlagsMVK C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkMacOSSurfaceCreateFlagsMVK = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkMacOSSurfaceCreateFlagsMVK(temp_in_VkMacOSSurfaceCreateFlagsMVK)
	}
	c.pView = g.View
}
func (g *MacOSSurfaceCreateInfoMVK) fromC(c *C.VkMacOSSurfaceCreateInfoMVK) {
	if g.Next != nil {
		g.Next.fromCStructure(c.pNext)
	}
	{
		var temp_in_VkMacOSSurfaceCreateFlagsMVK Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkMacOSSurfaceCreateFlagsMVK = Flags(temp_in_VkFlags)
		}
		g.Flags = MacOSSurfaceCreateFlagsMVK(temp_in_VkMacOSSurfaceCreateFlagsMVK)
	}
	g.View = c.pView
}
func (s *MacOSSurfaceCreateInfoMVK) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
}
func (s *MacOSSurfaceCreateInfoMVK) toCStructure(_sa *stackAllocator) unsafe.Pointer {
	c := (*C.VkMacOSSurfaceCreateInfoMVK)(_sa.alloc(C.sizeof_VkMacOSSurfaceCreateInfoMVK))
	s.toC(c, _sa)
	return unsafe.Pointer(c)
}
func (s *MacOSSurfaceCreateInfoMVK) fromCStructure(p unsafe.Pointer) {
	c := (*C.VkMacOSSurfaceCreateInfoMVK)(p)
	s.fromC(c)
}
func (s *MacOSSurfaceCreateInfoMVK) GetNext() Structure {
	return s.Next
}
func (s *MacOSSurfaceCreateInfoMVK) SetNext(n Structure) {
	s.Next = n
}
func (p PFNCreateMacOSSurfaceMVK) Call(instance Instance, createInfo *MacOSSurfaceCreateInfoMVK, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result) {
	var c struct {
		instance    C.VkInstance
		pCreateInfo *C.VkMacOSSurfaceCreateInfoMVK
		pAllocator  *C.VkAllocationCallbacks
		pSurface    *C.VkSurfaceKHR
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	if createInfo != nil {
		c.pCreateInfo = (*C.VkMacOSSurfaceCreateInfoMVK)(_sa.alloc(C.sizeof_VkMacOSSurfaceCreateInfoMVK))
		createInfo.toC(c.pCreateInfo, _sa)
	} else {
		c.pCreateInfo = nil
	}
	if allocator != nil {
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	} else {
		c.pAllocator = nil
	}
	if surface != nil {
		c.pSurface = (*C.VkSurfaceKHR)(_sa.alloc(C.sizeof_VkSurfaceKHR))
		*c.pSurface = C.VkSurfaceKHR(*surface)
	} else {
		c.pSurface = nil
	}
	c._ret = C.callPFN_vkCreateMacOSSurfaceMVK(C.PFN_vkCreateMacOSSurfaceMVK(unsafe.Pointer(p)), c.instance, c.pCreateInfo, c.pAllocator, c.pSurface)
	_ret = Result(c._ret)
	if surface != nil {
		*surface = SurfaceKHR(*c.pSurface)
	}
	return
}

func CreateMacOSSurfaceMVK(instance Instance, createInfo *MacOSSurfaceCreateInfoMVK, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result) {
	var c struct {
		instance    C.VkInstance
		pCreateInfo *C.VkMacOSSurfaceCreateInfoMVK
		pAllocator  *C.VkAllocationCallbacks
		pSurface    *C.VkSurfaceKHR
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	if createInfo != nil {
		c.pCreateInfo = (*C.VkMacOSSurfaceCreateInfoMVK)(_sa.alloc(C.sizeof_VkMacOSSurfaceCreateInfoMVK))
		createInfo.toC(c.pCreateInfo, _sa)
	} else {
		c.pCreateInfo = nil
	}
	if allocator != nil {
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	} else {
		c.pAllocator = nil
	}
	if surface != nil {
		c.pSurface = (*C.VkSurfaceKHR)(_sa.alloc(C.sizeof_VkSurfaceKHR))
		*c.pSurface = C.VkSurfaceKHR(*surface)
	} else {
		c.pSurface = nil
	}
	c._ret = C.vkCreateMacOSSurfaceMVK(c.instance, c.pCreateInfo, c.pAllocator, c.pSurface)
	_ret = Result(c._ret)
	if surface != nil {
		*surface = SurfaceKHR(*c.pSurface)
	}
	return
}
