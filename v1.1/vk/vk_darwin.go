package vk

/*
#include "bridges_darwin.h"
typedef void * void_pointer;
*/
import "C"
import "unsafe"

const MVK_MACOS_SURFACE_SPEC_VERSION uint32 = 2
const MVK_MACOS_SURFACE_EXTENSION_NAME string = "VK_MVK_macos_surface"

type PFNCreateMacOSSurfaceMVK C.PFN_vkCreateMacOSSurfaceMVK

type MacOSSurfaceCreateFlagsMVK Flags
type MacOSSurfaceCreateInfoMVK struct {
	Next  Structure
	Flags MacOSSurfaceCreateFlagsMVK
	View  unsafe.Pointer
}

func (g *MacOSSurfaceCreateInfoMVK) toC(c *C.VkMacOSSurfaceCreateInfoMVK, m *cmemory) {
	c.sType = g.sType()
	if g.Next != nil {
		c.pNext = g.Next.toCStructure(m)
	}
	c.flags = C.VkMacOSSurfaceCreateFlagsMVK(g.Flags)
	c.pView = g.View
}

func (g *MacOSSurfaceCreateInfoMVK) fromC(c *C.VkMacOSSurfaceCreateInfoMVK) {
	if g.Next != nil {
		g.Next.fromCStructure(c.pNext)
	}
	g.Flags = MacOSSurfaceCreateFlagsMVK(c.flags)
	g.View = c.pView
}
func (s *MacOSSurfaceCreateInfoMVK) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
}
func (s *MacOSSurfaceCreateInfoMVK) toCStructure(m *cmemory) unsafe.Pointer {
	c := (*C.VkMacOSSurfaceCreateInfoMVK)(m.alloc(C.sizeof_VkMacOSSurfaceCreateInfoMVK))
	s.toC(c, m)
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

type FuncCreateMacOSSurfaceMVK func(instance Instance, createInfo *MacOSSurfaceCreateInfoMVK, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result)

func ToCreateMacOSSurfaceMVK(p PFNVoidFunction) (fn FuncCreateMacOSSurfaceMVK) {
	return func(instance Instance, createInfo *MacOSSurfaceCreateInfoMVK, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result) {
		var c struct {
			instance    C.VkInstance
			pCreateInfo *C.VkMacOSSurfaceCreateInfoMVK
			pAllocator  *C.VkAllocationCallbacks
			pSurface    *C.VkSurfaceKHR
			_ret        C.VkResult
		}
		m := pool.take()
		defer pool.give(m)
		c.instance = C.VkInstance(instance)
		if createInfo != nil {
			c.pCreateInfo = (*C.VkMacOSSurfaceCreateInfoMVK)(m.alloc(C.sizeof_VkMacOSSurfaceCreateInfoMVK))
			createInfo.toC(c.pCreateInfo, m)
		} else {
			c.pCreateInfo = nil
		}
		if allocator != nil {
			c.pAllocator = (*C.VkAllocationCallbacks)(m.alloc(C.sizeof_VkAllocationCallbacks))
			allocator.toC(c.pAllocator)
		} else {
			c.pAllocator = nil
		}
		if surface != nil {
			c.pSurface = (*C.VkSurfaceKHR)(m.alloc(C.sizeof_VkSurfaceKHR))
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
}
