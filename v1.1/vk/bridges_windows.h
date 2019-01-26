#ifndef VK_USE_PLATFORM_WIN32_KHR
#define VK_USE_PLATFORM_WIN32_KHR
#endif
#include "vulkan/vulkan.h"

VkResult callPFN_vkCreateWin32SurfaceKHR(PFN_vkCreateWin32SurfaceKHR f, VkInstance arg0,  const VkWin32SurfaceCreateInfoKHR * arg1,  const VkAllocationCallbacks * arg2,  VkSurfaceKHR * arg3);

VkBool32 callPFN_vkGetPhysicalDeviceWin32PresentationSupportKHR(PFN_vkGetPhysicalDeviceWin32PresentationSupportKHR f, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex);
