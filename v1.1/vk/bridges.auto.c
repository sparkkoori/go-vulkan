#include "bridges.auto.h"

void *callPFN_vkAllocationFunction(PFN_vkAllocationFunction f, void * arg0,  size_t arg1,  size_t arg2,  VkSystemAllocationScope arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void *callPFN_vkReallocationFunction(PFN_vkReallocationFunction f, void * arg0,  void * arg1,  size_t arg2,  size_t arg3,  VkSystemAllocationScope arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkFreeFunction(PFN_vkFreeFunction f, void * arg0,  void * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkInternalAllocationNotification(PFN_vkInternalAllocationNotification f, void * arg0,  size_t arg1,  VkInternalAllocationType arg2,  VkSystemAllocationScope arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkInternalFreeNotification(PFN_vkInternalFreeNotification f, void * arg0,  size_t arg1,  VkInternalAllocationType arg2,  VkSystemAllocationScope arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

