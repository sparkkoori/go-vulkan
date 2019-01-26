#include "bridges_darwin.h"

VkResult callPFN_vkCreateMacOSSurfaceMVK(PFN_vkCreateMacOSSurfaceMVK f, VkInstance arg0,  const VkMacOSSurfaceCreateInfoMVK * arg1,  const VkAllocationCallbacks * arg2,  VkSurfaceKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};
