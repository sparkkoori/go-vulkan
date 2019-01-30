package vk

/*
#include "bridges_windows.h"
typedef void * void_pointer;
*/
import "C"
import "unsafe"

const KHR_WIN32_SURFACE_SPEC_VERSION uint32 = 6
const KHR_WIN32_SURFACE_EXTENSION_NAME string = "VK_KHR_win32_surface"

type Win32SurfaceCreateFlagsKHR Flags
type Win32SurfaceCreateInfoKHR struct {
	Next      Structure
	Flags     Win32SurfaceCreateFlagsKHR
	Hinstance unsafe.Pointer
	Hwnd      unsafe.Pointer
}

func (g *Win32SurfaceCreateInfoKHR) toC(c *C.VkWin32SurfaceCreateInfoKHR, m *cmemory) {
	c.sType = g.sType()
	if g.Next != nil {
		c.pNext = g.Next.toCStructure(m)
	}
	c.flags = C.VkWin32SurfaceCreateFlagsKHR(g.Flags)
	c.hinstance = C.HINSTANCE(g.Hinstance)
	c.hwnd = C.HWND(g.Hwnd)
}
func (g *Win32SurfaceCreateInfoKHR) fromC(c *C.VkWin32SurfaceCreateInfoKHR) {
	if g.Next != nil {
		g.Next.fromCStructure(c.pNext)
	}
	g.Flags = Win32SurfaceCreateFlagsKHR(c.flags)
	g.Hinstance = unsafe.Pointer(c.hinstance)
	g.Hwnd = unsafe.Pointer(c.hwnd)
}
func (s *Win32SurfaceCreateInfoKHR) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
}
func (s *Win32SurfaceCreateInfoKHR) toCStructure(m *cmemory) unsafe.Pointer {
	c := (*C.VkWin32SurfaceCreateInfoKHR)(m.alloc(C.sizeof_VkWin32SurfaceCreateInfoKHR))
	s.toC(c, m)
	return unsafe.Pointer(c)
}
func (s *Win32SurfaceCreateInfoKHR) fromCStructure(p unsafe.Pointer) {
	c := (*C.VkWin32SurfaceCreateInfoKHR)(p)
	s.fromC(c)
}
func (s *Win32SurfaceCreateInfoKHR) GetNext() Structure {
	return s.Next
}
func (s *Win32SurfaceCreateInfoKHR) SetNext(n Structure) {
	s.Next = n
}

type PFNCreateWin32SurfaceKHR uintptr

func (p PFNCreateWin32SurfaceKHR) Call(instance Instance, createInfo *Win32SurfaceCreateInfoKHR, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result) {
	var c struct {
		instance    C.VkInstance
		pCreateInfo *C.VkWin32SurfaceCreateInfoKHR
		pAllocator  *C.VkAllocationCallbacks
		pSurface    *C.VkSurfaceKHR
		_ret        C.VkResult
	}
	m := pool.take()
	defer pool.give(m)
	c.instance = C.VkInstance(instance)
	if createInfo != nil {
		c.pCreateInfo = (*C.VkWin32SurfaceCreateInfoKHR)(m.alloc(C.sizeof_VkWin32SurfaceCreateInfoKHR))
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
	c._ret = C.callPFN_vkCreateWin32SurfaceKHR(C.PFN_vkCreateWin32SurfaceKHR(unsafe.Pointer(p)), c.instance, c.pCreateInfo, c.pAllocator, c.pSurface)
	_ret = Result(c._ret)
	if surface != nil {
		*surface = SurfaceKHR(*c.pSurface)
	}
	return
}

type PFNGetPhysicalDeviceWin32PresentationSupportKHR uintptr

func (p PFNGetPhysicalDeviceWin32PresentationSupportKHR) Call(physicalDevice PhysicalDevice, queueFamilyIndex uint32) (_ret bool) {
	var c struct {
		physicalDevice   C.VkPhysicalDevice
		queueFamilyIndex C.uint32_t
		_ret             C.VkBool32
	}
	m := pool.take()
	defer pool.give(m)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.queueFamilyIndex = C.uint32_t(queueFamilyIndex)
	c._ret = C.callPFN_vkGetPhysicalDeviceWin32PresentationSupportKHR(C.PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR(unsafe.Pointer(p)), c.physicalDevice, c.queueFamilyIndex)
	_ret = c._ret != 0
	return
}

//TODO: other win32 extensions
