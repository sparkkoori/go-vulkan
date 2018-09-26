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

void callPFN_vkVoidFunction(PFN_vkVoidFunction f)
{
  return f();
};

VkResult callPFN_vkCreateInstance(PFN_vkCreateInstance f, const VkInstanceCreateInfo * arg0,  const VkAllocationCallbacks * arg1,  VkInstance * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkDestroyInstance(PFN_vkDestroyInstance f, VkInstance arg0,  const VkAllocationCallbacks * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices f, VkInstance arg0,  uint32_t * arg1,  VkPhysicalDevice * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures f, VkPhysicalDevice arg0,  VkPhysicalDeviceFeatures * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceFormatProperties(PFN_vkGetPhysicalDeviceFormatProperties f, VkPhysicalDevice arg0,  VkFormat arg1,  VkFormatProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceImageFormatProperties(PFN_vkGetPhysicalDeviceImageFormatProperties f, VkPhysicalDevice arg0,  VkFormat arg1,  VkImageType arg2,  VkImageTiling arg3,  VkImageUsageFlags arg4,  VkImageCreateFlags arg5,  VkImageFormatProperties * arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties f, VkPhysicalDevice arg0,  VkPhysicalDeviceProperties * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkQueueFamilyProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceMemoryProperties(PFN_vkGetPhysicalDeviceMemoryProperties f, VkPhysicalDevice arg0,  VkPhysicalDeviceMemoryProperties * arg1)
{
  return f(arg0, arg1);
};

PFN_vkVoidFunction callPFN_vkGetInstanceProcAddr(PFN_vkGetInstanceProcAddr f, VkInstance arg0,  const char * arg1)
{
  return f(arg0, arg1);
};

PFN_vkVoidFunction callPFN_vkGetDeviceProcAddr(PFN_vkGetDeviceProcAddr f, VkDevice arg0,  const char * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkCreateDevice(PFN_vkCreateDevice f, VkPhysicalDevice arg0,  const VkDeviceCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDevice * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDevice(PFN_vkDestroyDevice f, VkDevice arg0,  const VkAllocationCallbacks * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkEnumerateInstanceExtensionProperties(PFN_vkEnumerateInstanceExtensionProperties f, const char * arg0,  uint32_t * arg1,  VkExtensionProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkEnumerateDeviceExtensionProperties(PFN_vkEnumerateDeviceExtensionProperties f, VkPhysicalDevice arg0,  const char * arg1,  uint32_t * arg2,  VkExtensionProperties * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkEnumerateInstanceLayerProperties(PFN_vkEnumerateInstanceLayerProperties f, uint32_t * arg0,  VkLayerProperties * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkEnumerateDeviceLayerProperties(PFN_vkEnumerateDeviceLayerProperties f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkLayerProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDeviceQueue(PFN_vkGetDeviceQueue f, VkDevice arg0,  uint32_t arg1,  uint32_t arg2,  VkQueue * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkQueueSubmit(PFN_vkQueueSubmit f, VkQueue arg0,  uint32_t arg1,  const VkSubmitInfo * arg2,  VkFence arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkQueueWaitIdle(PFN_vkQueueWaitIdle f, VkQueue arg0)
{
  return f(arg0);
};

VkResult callPFN_vkDeviceWaitIdle(PFN_vkDeviceWaitIdle f, VkDevice arg0)
{
  return f(arg0);
};

VkResult callPFN_vkAllocateMemory(PFN_vkAllocateMemory f, VkDevice arg0,  const VkMemoryAllocateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDeviceMemory * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkFreeMemory(PFN_vkFreeMemory f, VkDevice arg0,  VkDeviceMemory arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkMapMemory(PFN_vkMapMemory f, VkDevice arg0,  VkDeviceMemory arg1,  VkDeviceSize arg2,  VkDeviceSize arg3,  VkMemoryMapFlags arg4,  void ** arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkUnmapMemory(PFN_vkUnmapMemory f, VkDevice arg0,  VkDeviceMemory arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkFlushMappedMemoryRanges(PFN_vkFlushMappedMemoryRanges f, VkDevice arg0,  uint32_t arg1,  const VkMappedMemoryRange * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkInvalidateMappedMemoryRanges(PFN_vkInvalidateMappedMemoryRanges f, VkDevice arg0,  uint32_t arg1,  const VkMappedMemoryRange * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDeviceMemoryCommitment(PFN_vkGetDeviceMemoryCommitment f, VkDevice arg0,  VkDeviceMemory arg1,  VkDeviceSize * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkBindBufferMemory(PFN_vkBindBufferMemory f, VkDevice arg0,  VkBuffer arg1,  VkDeviceMemory arg2,  VkDeviceSize arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkBindImageMemory(PFN_vkBindImageMemory f, VkDevice arg0,  VkImage arg1,  VkDeviceMemory arg2,  VkDeviceSize arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetBufferMemoryRequirements(PFN_vkGetBufferMemoryRequirements f, VkDevice arg0,  VkBuffer arg1,  VkMemoryRequirements * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageMemoryRequirements(PFN_vkGetImageMemoryRequirements f, VkDevice arg0,  VkImage arg1,  VkMemoryRequirements * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageSparseMemoryRequirements(PFN_vkGetImageSparseMemoryRequirements f, VkDevice arg0,  VkImage arg1,  uint32_t * arg2,  VkSparseImageMemoryRequirements * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetPhysicalDeviceSparseImageFormatProperties(PFN_vkGetPhysicalDeviceSparseImageFormatProperties f, VkPhysicalDevice arg0,  VkFormat arg1,  VkImageType arg2,  VkSampleCountFlagBits arg3,  VkImageUsageFlags arg4,  VkImageTiling arg5,  uint32_t * arg6,  VkSparseImageFormatProperties * arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

VkResult callPFN_vkQueueBindSparse(PFN_vkQueueBindSparse f, VkQueue arg0,  uint32_t arg1,  const VkBindSparseInfo * arg2,  VkFence arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateFence(PFN_vkCreateFence f, VkDevice arg0,  const VkFenceCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkFence * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyFence(PFN_vkDestroyFence f, VkDevice arg0,  VkFence arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkResetFences(PFN_vkResetFences f, VkDevice arg0,  uint32_t arg1,  const VkFence * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetFenceStatus(PFN_vkGetFenceStatus f, VkDevice arg0,  VkFence arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkWaitForFences(PFN_vkWaitForFences f, VkDevice arg0,  uint32_t arg1,  const VkFence * arg2,  VkBool32 arg3,  uint64_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkCreateSemaphore(PFN_vkCreateSemaphore f, VkDevice arg0,  const VkSemaphoreCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkSemaphore * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroySemaphore(PFN_vkDestroySemaphore f, VkDevice arg0,  VkSemaphore arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateEvent(PFN_vkCreateEvent f, VkDevice arg0,  const VkEventCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkEvent * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyEvent(PFN_vkDestroyEvent f, VkDevice arg0,  VkEvent arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetEventStatus(PFN_vkGetEventStatus f, VkDevice arg0,  VkEvent arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkSetEvent(PFN_vkSetEvent f, VkDevice arg0,  VkEvent arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkResetEvent(PFN_vkResetEvent f, VkDevice arg0,  VkEvent arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkCreateQueryPool(PFN_vkCreateQueryPool f, VkDevice arg0,  const VkQueryPoolCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkQueryPool * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyQueryPool(PFN_vkDestroyQueryPool f, VkDevice arg0,  VkQueryPool arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetQueryPoolResults(PFN_vkGetQueryPoolResults f, VkDevice arg0,  VkQueryPool arg1,  uint32_t arg2,  uint32_t arg3,  size_t arg4,  void * arg5,  VkDeviceSize arg6,  VkQueryResultFlags arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

VkResult callPFN_vkCreateBuffer(PFN_vkCreateBuffer f, VkDevice arg0,  const VkBufferCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkBuffer * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyBuffer(PFN_vkDestroyBuffer f, VkDevice arg0,  VkBuffer arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateBufferView(PFN_vkCreateBufferView f, VkDevice arg0,  const VkBufferViewCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkBufferView * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyBufferView(PFN_vkDestroyBufferView f, VkDevice arg0,  VkBufferView arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateImage(PFN_vkCreateImage f, VkDevice arg0,  const VkImageCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkImage * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyImage(PFN_vkDestroyImage f, VkDevice arg0,  VkImage arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageSubresourceLayout(PFN_vkGetImageSubresourceLayout f, VkDevice arg0,  VkImage arg1,  const VkImageSubresource * arg2,  VkSubresourceLayout * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateImageView(PFN_vkCreateImageView f, VkDevice arg0,  const VkImageViewCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkImageView * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyImageView(PFN_vkDestroyImageView f, VkDevice arg0,  VkImageView arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateShaderModule(PFN_vkCreateShaderModule f, VkDevice arg0,  const VkShaderModuleCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkShaderModule * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyShaderModule(PFN_vkDestroyShaderModule f, VkDevice arg0,  VkShaderModule arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreatePipelineCache(PFN_vkCreatePipelineCache f, VkDevice arg0,  const VkPipelineCacheCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkPipelineCache * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyPipelineCache(PFN_vkDestroyPipelineCache f, VkDevice arg0,  VkPipelineCache arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPipelineCacheData(PFN_vkGetPipelineCacheData f, VkDevice arg0,  VkPipelineCache arg1,  size_t * arg2,  void * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkMergePipelineCaches(PFN_vkMergePipelineCaches f, VkDevice arg0,  VkPipelineCache arg1,  uint32_t arg2,  const VkPipelineCache * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines f, VkDevice arg0,  VkPipelineCache arg1,  uint32_t arg2,  const VkGraphicsPipelineCreateInfo * arg3,  const VkAllocationCallbacks * arg4,  VkPipeline * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

VkResult callPFN_vkCreateComputePipelines(PFN_vkCreateComputePipelines f, VkDevice arg0,  VkPipelineCache arg1,  uint32_t arg2,  const VkComputePipelineCreateInfo * arg3,  const VkAllocationCallbacks * arg4,  VkPipeline * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkDestroyPipeline(PFN_vkDestroyPipeline f, VkDevice arg0,  VkPipeline arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreatePipelineLayout(PFN_vkCreatePipelineLayout f, VkDevice arg0,  const VkPipelineLayoutCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkPipelineLayout * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyPipelineLayout(PFN_vkDestroyPipelineLayout f, VkDevice arg0,  VkPipelineLayout arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateSampler(PFN_vkCreateSampler f, VkDevice arg0,  const VkSamplerCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkSampler * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroySampler(PFN_vkDestroySampler f, VkDevice arg0,  VkSampler arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateDescriptorSetLayout(PFN_vkCreateDescriptorSetLayout f, VkDevice arg0,  const VkDescriptorSetLayoutCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDescriptorSetLayout * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDescriptorSetLayout(PFN_vkDestroyDescriptorSetLayout f, VkDevice arg0,  VkDescriptorSetLayout arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateDescriptorPool(PFN_vkCreateDescriptorPool f, VkDevice arg0,  const VkDescriptorPoolCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDescriptorPool * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDescriptorPool(PFN_vkDestroyDescriptorPool f, VkDevice arg0,  VkDescriptorPool arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkResetDescriptorPool(PFN_vkResetDescriptorPool f, VkDevice arg0,  VkDescriptorPool arg1,  VkDescriptorPoolResetFlags arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkAllocateDescriptorSets(PFN_vkAllocateDescriptorSets f, VkDevice arg0,  const VkDescriptorSetAllocateInfo * arg1,  VkDescriptorSet * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkFreeDescriptorSets(PFN_vkFreeDescriptorSets f, VkDevice arg0,  VkDescriptorPool arg1,  uint32_t arg2,  const VkDescriptorSet * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkUpdateDescriptorSets(PFN_vkUpdateDescriptorSets f, VkDevice arg0,  uint32_t arg1,  const VkWriteDescriptorSet * arg2,  uint32_t arg3,  const VkCopyDescriptorSet * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkCreateFramebuffer(PFN_vkCreateFramebuffer f, VkDevice arg0,  const VkFramebufferCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkFramebuffer * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyFramebuffer(PFN_vkDestroyFramebuffer f, VkDevice arg0,  VkFramebuffer arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateRenderPass(PFN_vkCreateRenderPass f, VkDevice arg0,  const VkRenderPassCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkRenderPass * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyRenderPass(PFN_vkDestroyRenderPass f, VkDevice arg0,  VkRenderPass arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetRenderAreaGranularity(PFN_vkGetRenderAreaGranularity f, VkDevice arg0,  VkRenderPass arg1,  VkExtent2D * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateCommandPool(PFN_vkCreateCommandPool f, VkDevice arg0,  const VkCommandPoolCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkCommandPool * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyCommandPool(PFN_vkDestroyCommandPool f, VkDevice arg0,  VkCommandPool arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkResetCommandPool(PFN_vkResetCommandPool f, VkDevice arg0,  VkCommandPool arg1,  VkCommandPoolResetFlags arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkAllocateCommandBuffers(PFN_vkAllocateCommandBuffers f, VkDevice arg0,  const VkCommandBufferAllocateInfo * arg1,  VkCommandBuffer * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkFreeCommandBuffers(PFN_vkFreeCommandBuffers f, VkDevice arg0,  VkCommandPool arg1,  uint32_t arg2,  const VkCommandBuffer * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkBeginCommandBuffer(PFN_vkBeginCommandBuffer f, VkCommandBuffer arg0,  const VkCommandBufferBeginInfo * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkEndCommandBuffer(PFN_vkEndCommandBuffer f, VkCommandBuffer arg0)
{
  return f(arg0);
};

VkResult callPFN_vkResetCommandBuffer(PFN_vkResetCommandBuffer f, VkCommandBuffer arg0,  VkCommandBufferResetFlags arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdBindPipeline(PFN_vkCmdBindPipeline f, VkCommandBuffer arg0,  VkPipelineBindPoint arg1,  VkPipeline arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdSetViewport(PFN_vkCmdSetViewport f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  const VkViewport * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdSetScissor(PFN_vkCmdSetScissor f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  const VkRect2D * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdSetLineWidth(PFN_vkCmdSetLineWidth f, VkCommandBuffer arg0,  float arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdSetDepthBias(PFN_vkCmdSetDepthBias f, VkCommandBuffer arg0,  float arg1,  float arg2,  float arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdSetBlendConstants(PFN_vkCmdSetBlendConstants f, VkCommandBuffer arg0,  const float * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdSetDepthBounds(PFN_vkCmdSetDepthBounds f, VkCommandBuffer arg0,  float arg1,  float arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdSetStencilCompareMask(PFN_vkCmdSetStencilCompareMask f, VkCommandBuffer arg0,  VkStencilFaceFlags arg1,  uint32_t arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdSetStencilWriteMask(PFN_vkCmdSetStencilWriteMask f, VkCommandBuffer arg0,  VkStencilFaceFlags arg1,  uint32_t arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdSetStencilReference(PFN_vkCmdSetStencilReference f, VkCommandBuffer arg0,  VkStencilFaceFlags arg1,  uint32_t arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdBindDescriptorSets(PFN_vkCmdBindDescriptorSets f, VkCommandBuffer arg0,  VkPipelineBindPoint arg1,  VkPipelineLayout arg2,  uint32_t arg3,  uint32_t arg4,  const VkDescriptorSet * arg5,  uint32_t arg6,  const uint32_t * arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

void callPFN_vkCmdBindIndexBuffer(PFN_vkCmdBindIndexBuffer f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkIndexType arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdBindVertexBuffers(PFN_vkCmdBindVertexBuffers f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  const VkBuffer * arg3,  const VkDeviceSize * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdDraw(PFN_vkCmdDraw f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  uint32_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdDrawIndexed(PFN_vkCmdDrawIndexed f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  int32_t arg4,  uint32_t arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdDrawIndirect(PFN_vkCmdDrawIndirect f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  uint32_t arg3,  uint32_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdDrawIndexedIndirect(PFN_vkCmdDrawIndexedIndirect f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  uint32_t arg3,  uint32_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdDispatch(PFN_vkCmdDispatch f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdDispatchIndirect(PFN_vkCmdDispatchIndirect f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdCopyBuffer(PFN_vkCmdCopyBuffer f, VkCommandBuffer arg0,  VkBuffer arg1,  VkBuffer arg2,  uint32_t arg3,  const VkBufferCopy * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdCopyImage(PFN_vkCmdCopyImage f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  VkImage arg3,  VkImageLayout arg4,  uint32_t arg5,  const VkImageCopy * arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkCmdBlitImage(PFN_vkCmdBlitImage f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  VkImage arg3,  VkImageLayout arg4,  uint32_t arg5,  const VkImageBlit * arg6,  VkFilter arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

void callPFN_vkCmdCopyBufferToImage(PFN_vkCmdCopyBufferToImage f, VkCommandBuffer arg0,  VkBuffer arg1,  VkImage arg2,  VkImageLayout arg3,  uint32_t arg4,  const VkBufferImageCopy * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdCopyImageToBuffer(PFN_vkCmdCopyImageToBuffer f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  VkBuffer arg3,  uint32_t arg4,  const VkBufferImageCopy * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdUpdateBuffer(PFN_vkCmdUpdateBuffer f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkDeviceSize arg3,  const void * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdFillBuffer(PFN_vkCmdFillBuffer f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkDeviceSize arg3,  uint32_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdClearColorImage(PFN_vkCmdClearColorImage f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  const VkClearColorValue * arg3,  uint32_t arg4,  const VkImageSubresourceRange * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdClearDepthStencilImage(PFN_vkCmdClearDepthStencilImage f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  const VkClearDepthStencilValue * arg3,  uint32_t arg4,  const VkImageSubresourceRange * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdClearAttachments(PFN_vkCmdClearAttachments f, VkCommandBuffer arg0,  uint32_t arg1,  const VkClearAttachment * arg2,  uint32_t arg3,  const VkClearRect * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdResolveImage(PFN_vkCmdResolveImage f, VkCommandBuffer arg0,  VkImage arg1,  VkImageLayout arg2,  VkImage arg3,  VkImageLayout arg4,  uint32_t arg5,  const VkImageResolve * arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkCmdSetEvent(PFN_vkCmdSetEvent f, VkCommandBuffer arg0,  VkEvent arg1,  VkPipelineStageFlags arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdResetEvent(PFN_vkCmdResetEvent f, VkCommandBuffer arg0,  VkEvent arg1,  VkPipelineStageFlags arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdWaitEvents(PFN_vkCmdWaitEvents f, VkCommandBuffer arg0,  uint32_t arg1,  const VkEvent * arg2,  VkPipelineStageFlags arg3,  VkPipelineStageFlags arg4,  uint32_t arg5,  const VkMemoryBarrier * arg6,  uint32_t arg7,  const VkBufferMemoryBarrier * arg8,  uint32_t arg9,  const VkImageMemoryBarrier * arg10)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10);
};

void callPFN_vkCmdPipelineBarrier(PFN_vkCmdPipelineBarrier f, VkCommandBuffer arg0,  VkPipelineStageFlags arg1,  VkPipelineStageFlags arg2,  VkDependencyFlags arg3,  uint32_t arg4,  const VkMemoryBarrier * arg5,  uint32_t arg6,  const VkBufferMemoryBarrier * arg7,  uint32_t arg8,  const VkImageMemoryBarrier * arg9)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9);
};

void callPFN_vkCmdBeginQuery(PFN_vkCmdBeginQuery f, VkCommandBuffer arg0,  VkQueryPool arg1,  uint32_t arg2,  VkQueryControlFlags arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdEndQuery(PFN_vkCmdEndQuery f, VkCommandBuffer arg0,  VkQueryPool arg1,  uint32_t arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdResetQueryPool(PFN_vkCmdResetQueryPool f, VkCommandBuffer arg0,  VkQueryPool arg1,  uint32_t arg2,  uint32_t arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdWriteTimestamp(PFN_vkCmdWriteTimestamp f, VkCommandBuffer arg0,  VkPipelineStageFlagBits arg1,  VkQueryPool arg2,  uint32_t arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdCopyQueryPoolResults(PFN_vkCmdCopyQueryPoolResults f, VkCommandBuffer arg0,  VkQueryPool arg1,  uint32_t arg2,  uint32_t arg3,  VkBuffer arg4,  VkDeviceSize arg5,  VkDeviceSize arg6,  VkQueryResultFlags arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

void callPFN_vkCmdPushConstants(PFN_vkCmdPushConstants f, VkCommandBuffer arg0,  VkPipelineLayout arg1,  VkShaderStageFlags arg2,  uint32_t arg3,  uint32_t arg4,  const void * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdBeginRenderPass(PFN_vkCmdBeginRenderPass f, VkCommandBuffer arg0,  const VkRenderPassBeginInfo * arg1,  VkSubpassContents arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdNextSubpass(PFN_vkCmdNextSubpass f, VkCommandBuffer arg0,  VkSubpassContents arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdEndRenderPass(PFN_vkCmdEndRenderPass f, VkCommandBuffer arg0)
{
  return f(arg0);
};

void callPFN_vkCmdExecuteCommands(PFN_vkCmdExecuteCommands f, VkCommandBuffer arg0,  uint32_t arg1,  const VkCommandBuffer * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkEnumerateInstanceVersion(PFN_vkEnumerateInstanceVersion f, uint32_t * arg0)
{
  return f(arg0);
};

VkResult callPFN_vkBindBufferMemory2(PFN_vkBindBufferMemory2 f, VkDevice arg0,  uint32_t arg1,  const VkBindBufferMemoryInfo * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkBindImageMemory2(PFN_vkBindImageMemory2 f, VkDevice arg0,  uint32_t arg1,  const VkBindImageMemoryInfo * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDeviceGroupPeerMemoryFeatures(PFN_vkGetDeviceGroupPeerMemoryFeatures f, VkDevice arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  VkPeerMemoryFeatureFlags * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdSetDeviceMask(PFN_vkCmdSetDeviceMask f, VkCommandBuffer arg0,  uint32_t arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdDispatchBase(PFN_vkCmdDispatchBase f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  uint32_t arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

VkResult callPFN_vkEnumeratePhysicalDeviceGroups(PFN_vkEnumeratePhysicalDeviceGroups f, VkInstance arg0,  uint32_t * arg1,  VkPhysicalDeviceGroupProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageMemoryRequirements2(PFN_vkGetImageMemoryRequirements2 f, VkDevice arg0,  const VkImageMemoryRequirementsInfo2 * arg1,  VkMemoryRequirements2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetBufferMemoryRequirements2(PFN_vkGetBufferMemoryRequirements2 f, VkDevice arg0,  const VkBufferMemoryRequirementsInfo2 * arg1,  VkMemoryRequirements2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageSparseMemoryRequirements2(PFN_vkGetImageSparseMemoryRequirements2 f, VkDevice arg0,  const VkImageSparseMemoryRequirementsInfo2 * arg1,  uint32_t * arg2,  VkSparseImageMemoryRequirements2 * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetPhysicalDeviceFeatures2(PFN_vkGetPhysicalDeviceFeatures2 f, VkPhysicalDevice arg0,  VkPhysicalDeviceFeatures2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceProperties2(PFN_vkGetPhysicalDeviceProperties2 f, VkPhysicalDevice arg0,  VkPhysicalDeviceProperties2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceFormatProperties2(PFN_vkGetPhysicalDeviceFormatProperties2 f, VkPhysicalDevice arg0,  VkFormat arg1,  VkFormatProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceImageFormatProperties2(PFN_vkGetPhysicalDeviceImageFormatProperties2 f, VkPhysicalDevice arg0,  const VkPhysicalDeviceImageFormatInfo2 * arg1,  VkImageFormatProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceQueueFamilyProperties2(PFN_vkGetPhysicalDeviceQueueFamilyProperties2 f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkQueueFamilyProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceMemoryProperties2(PFN_vkGetPhysicalDeviceMemoryProperties2 f, VkPhysicalDevice arg0,  VkPhysicalDeviceMemoryProperties2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceSparseImageFormatProperties2(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 f, VkPhysicalDevice arg0,  const VkPhysicalDeviceSparseImageFormatInfo2 * arg1,  uint32_t * arg2,  VkSparseImageFormatProperties2 * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkTrimCommandPool(PFN_vkTrimCommandPool f, VkDevice arg0,  VkCommandPool arg1,  VkCommandPoolTrimFlags arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDeviceQueue2(PFN_vkGetDeviceQueue2 f, VkDevice arg0,  const VkDeviceQueueInfo2 * arg1,  VkQueue * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateSamplerYcbcrConversion(PFN_vkCreateSamplerYcbcrConversion f, VkDevice arg0,  const VkSamplerYcbcrConversionCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkSamplerYcbcrConversion * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroySamplerYcbcrConversion(PFN_vkDestroySamplerYcbcrConversion f, VkDevice arg0,  VkSamplerYcbcrConversion arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateDescriptorUpdateTemplate(PFN_vkCreateDescriptorUpdateTemplate f, VkDevice arg0,  const VkDescriptorUpdateTemplateCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDescriptorUpdateTemplate * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDescriptorUpdateTemplate(PFN_vkDestroyDescriptorUpdateTemplate f, VkDevice arg0,  VkDescriptorUpdateTemplate arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkUpdateDescriptorSetWithTemplate(PFN_vkUpdateDescriptorSetWithTemplate f, VkDevice arg0,  VkDescriptorSet arg1,  VkDescriptorUpdateTemplate arg2,  const void * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetPhysicalDeviceExternalBufferProperties(PFN_vkGetPhysicalDeviceExternalBufferProperties f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalBufferInfo * arg1,  VkExternalBufferProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceExternalFenceProperties(PFN_vkGetPhysicalDeviceExternalFenceProperties f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalFenceInfo * arg1,  VkExternalFenceProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceExternalSemaphoreProperties(PFN_vkGetPhysicalDeviceExternalSemaphoreProperties f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalSemaphoreInfo * arg1,  VkExternalSemaphoreProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDescriptorSetLayoutSupport(PFN_vkGetDescriptorSetLayoutSupport f, VkDevice arg0,  const VkDescriptorSetLayoutCreateInfo * arg1,  VkDescriptorSetLayoutSupport * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkDestroySurfaceKHR(PFN_vkDestroySurfaceKHR f, VkInstance arg0,  VkSurfaceKHR arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR f, VkPhysicalDevice arg0,  uint32_t arg1,  VkSurfaceKHR arg2,  VkBool32 * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR f, VkPhysicalDevice arg0,  VkSurfaceKHR arg1,  VkSurfaceCapabilitiesKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR f, VkPhysicalDevice arg0,  VkSurfaceKHR arg1,  uint32_t * arg2,  VkSurfaceFormatKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR f, VkPhysicalDevice arg0,  VkSurfaceKHR arg1,  uint32_t * arg2,  VkPresentModeKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateSwapchainKHR(PFN_vkCreateSwapchainKHR f, VkDevice arg0,  const VkSwapchainCreateInfoKHR * arg1,  const VkAllocationCallbacks * arg2,  VkSwapchainKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroySwapchainKHR(PFN_vkDestroySwapchainKHR f, VkDevice arg0,  VkSwapchainKHR arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR f, VkDevice arg0,  VkSwapchainKHR arg1,  uint32_t * arg2,  VkImage * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkAcquireNextImageKHR(PFN_vkAcquireNextImageKHR f, VkDevice arg0,  VkSwapchainKHR arg1,  uint64_t arg2,  VkSemaphore arg3,  VkFence arg4,  uint32_t * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

VkResult callPFN_vkQueuePresentKHR(PFN_vkQueuePresentKHR f, VkQueue arg0,  const VkPresentInfoKHR * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetDeviceGroupPresentCapabilitiesKHR(PFN_vkGetDeviceGroupPresentCapabilitiesKHR f, VkDevice arg0,  VkDeviceGroupPresentCapabilitiesKHR * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetDeviceGroupSurfacePresentModesKHR(PFN_vkGetDeviceGroupSurfacePresentModesKHR f, VkDevice arg0,  VkSurfaceKHR arg1,  VkDeviceGroupPresentModeFlagsKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDevicePresentRectanglesKHR(PFN_vkGetPhysicalDevicePresentRectanglesKHR f, VkPhysicalDevice arg0,  VkSurfaceKHR arg1,  uint32_t * arg2,  VkRect2D * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkAcquireNextImage2KHR(PFN_vkAcquireNextImage2KHR f, VkDevice arg0,  const VkAcquireNextImageInfoKHR * arg1,  uint32_t * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceDisplayPropertiesKHR(PFN_vkGetPhysicalDeviceDisplayPropertiesKHR f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkDisplayPropertiesKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR(PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkDisplayPlanePropertiesKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetDisplayPlaneSupportedDisplaysKHR(PFN_vkGetDisplayPlaneSupportedDisplaysKHR f, VkPhysicalDevice arg0,  uint32_t arg1,  uint32_t * arg2,  VkDisplayKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetDisplayModePropertiesKHR(PFN_vkGetDisplayModePropertiesKHR f, VkPhysicalDevice arg0,  VkDisplayKHR arg1,  uint32_t * arg2,  VkDisplayModePropertiesKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateDisplayModeKHR(PFN_vkCreateDisplayModeKHR f, VkPhysicalDevice arg0,  VkDisplayKHR arg1,  const VkDisplayModeCreateInfoKHR * arg2,  const VkAllocationCallbacks * arg3,  VkDisplayModeKHR * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkGetDisplayPlaneCapabilitiesKHR(PFN_vkGetDisplayPlaneCapabilitiesKHR f, VkPhysicalDevice arg0,  VkDisplayModeKHR arg1,  uint32_t arg2,  VkDisplayPlaneCapabilitiesKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateDisplayPlaneSurfaceKHR(PFN_vkCreateDisplayPlaneSurfaceKHR f, VkInstance arg0,  const VkDisplaySurfaceCreateInfoKHR * arg1,  const VkAllocationCallbacks * arg2,  VkSurfaceKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateSharedSwapchainsKHR(PFN_vkCreateSharedSwapchainsKHR f, VkDevice arg0,  uint32_t arg1,  const VkSwapchainCreateInfoKHR * arg2,  const VkAllocationCallbacks * arg3,  VkSwapchainKHR * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkGetPhysicalDeviceFeatures2KHR(PFN_vkGetPhysicalDeviceFeatures2KHR f, VkPhysicalDevice arg0,  VkPhysicalDeviceFeatures2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceProperties2KHR(PFN_vkGetPhysicalDeviceProperties2KHR f, VkPhysicalDevice arg0,  VkPhysicalDeviceProperties2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceFormatProperties2KHR(PFN_vkGetPhysicalDeviceFormatProperties2KHR f, VkPhysicalDevice arg0,  VkFormat arg1,  VkFormatProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceImageFormatProperties2KHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceImageFormatInfo2 * arg1,  VkImageFormatProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR(PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkQueueFamilyProperties2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceMemoryProperties2KHR(PFN_vkGetPhysicalDeviceMemoryProperties2KHR f, VkPhysicalDevice arg0,  VkPhysicalDeviceMemoryProperties2 * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceSparseImageFormatInfo2 * arg1,  uint32_t * arg2,  VkSparseImageFormatProperties2 * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetDeviceGroupPeerMemoryFeaturesKHR(PFN_vkGetDeviceGroupPeerMemoryFeaturesKHR f, VkDevice arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  VkPeerMemoryFeatureFlags * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdSetDeviceMaskKHR(PFN_vkCmdSetDeviceMaskKHR f, VkCommandBuffer arg0,  uint32_t arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdDispatchBaseKHR(PFN_vkCmdDispatchBaseKHR f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  uint32_t arg3,  uint32_t arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkTrimCommandPoolKHR(PFN_vkTrimCommandPoolKHR f, VkDevice arg0,  VkCommandPool arg1,  VkCommandPoolTrimFlags arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkEnumeratePhysicalDeviceGroupsKHR(PFN_vkEnumeratePhysicalDeviceGroupsKHR f, VkInstance arg0,  uint32_t * arg1,  VkPhysicalDeviceGroupProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR(PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalBufferInfo * arg1,  VkExternalBufferProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetMemoryFdKHR(PFN_vkGetMemoryFdKHR f, VkDevice arg0,  const VkMemoryGetFdInfoKHR * arg1,  int * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetMemoryFdPropertiesKHR(PFN_vkGetMemoryFdPropertiesKHR f, VkDevice arg0,  VkExternalMemoryHandleTypeFlagBits arg1,  int arg2,  VkMemoryFdPropertiesKHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR(PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalSemaphoreInfo * arg1,  VkExternalSemaphoreProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkImportSemaphoreFdKHR(PFN_vkImportSemaphoreFdKHR f, VkDevice arg0,  const VkImportSemaphoreFdInfoKHR * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetSemaphoreFdKHR(PFN_vkGetSemaphoreFdKHR f, VkDevice arg0,  const VkSemaphoreGetFdInfoKHR * arg1,  int * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdPushDescriptorSetKHR(PFN_vkCmdPushDescriptorSetKHR f, VkCommandBuffer arg0,  VkPipelineBindPoint arg1,  VkPipelineLayout arg2,  uint32_t arg3,  uint32_t arg4,  const VkWriteDescriptorSet * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

void callPFN_vkCmdPushDescriptorSetWithTemplateKHR(PFN_vkCmdPushDescriptorSetWithTemplateKHR f, VkCommandBuffer arg0,  VkDescriptorUpdateTemplate arg1,  VkPipelineLayout arg2,  uint32_t arg3,  const void * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkCreateDescriptorUpdateTemplateKHR(PFN_vkCreateDescriptorUpdateTemplateKHR f, VkDevice arg0,  const VkDescriptorUpdateTemplateCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkDescriptorUpdateTemplate * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDescriptorUpdateTemplateKHR(PFN_vkDestroyDescriptorUpdateTemplateKHR f, VkDevice arg0,  VkDescriptorUpdateTemplate arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkUpdateDescriptorSetWithTemplateKHR(PFN_vkUpdateDescriptorSetWithTemplateKHR f, VkDevice arg0,  VkDescriptorSet arg1,  VkDescriptorUpdateTemplate arg2,  const void * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateRenderPass2KHR(PFN_vkCreateRenderPass2KHR f, VkDevice arg0,  const VkRenderPassCreateInfo2KHR * arg1,  const VkAllocationCallbacks * arg2,  VkRenderPass * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdBeginRenderPass2KHR(PFN_vkCmdBeginRenderPass2KHR f, VkCommandBuffer arg0,  const VkRenderPassBeginInfo * arg1,  const VkSubpassBeginInfoKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdNextSubpass2KHR(PFN_vkCmdNextSubpass2KHR f, VkCommandBuffer arg0,  const VkSubpassBeginInfoKHR * arg1,  const VkSubpassEndInfoKHR * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdEndRenderPass2KHR(PFN_vkCmdEndRenderPass2KHR f, VkCommandBuffer arg0,  const VkSubpassEndInfoKHR * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetSwapchainStatusKHR(PFN_vkGetSwapchainStatusKHR f, VkDevice arg0,  VkSwapchainKHR arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceExternalFencePropertiesKHR(PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceExternalFenceInfo * arg1,  VkExternalFenceProperties * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkImportFenceFdKHR(PFN_vkImportFenceFdKHR f, VkDevice arg0,  const VkImportFenceFdInfoKHR * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetFenceFdKHR(PFN_vkGetFenceFdKHR f, VkDevice arg0,  const VkFenceGetFdInfoKHR * arg1,  int * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR(PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceSurfaceInfo2KHR * arg1,  VkSurfaceCapabilities2KHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceFormats2KHR(PFN_vkGetPhysicalDeviceSurfaceFormats2KHR f, VkPhysicalDevice arg0,  const VkPhysicalDeviceSurfaceInfo2KHR * arg1,  uint32_t * arg2,  VkSurfaceFormat2KHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetPhysicalDeviceDisplayProperties2KHR(PFN_vkGetPhysicalDeviceDisplayProperties2KHR f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkDisplayProperties2KHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR(PFN_vkGetPhysicalDeviceDisplayPlaneProperties2KHR f, VkPhysicalDevice arg0,  uint32_t * arg1,  VkDisplayPlaneProperties2KHR * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetDisplayModeProperties2KHR(PFN_vkGetDisplayModeProperties2KHR f, VkPhysicalDevice arg0,  VkDisplayKHR arg1,  uint32_t * arg2,  VkDisplayModeProperties2KHR * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetDisplayPlaneCapabilities2KHR(PFN_vkGetDisplayPlaneCapabilities2KHR f, VkPhysicalDevice arg0,  const VkDisplayPlaneInfo2KHR * arg1,  VkDisplayPlaneCapabilities2KHR * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageMemoryRequirements2KHR(PFN_vkGetImageMemoryRequirements2KHR f, VkDevice arg0,  const VkImageMemoryRequirementsInfo2 * arg1,  VkMemoryRequirements2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetBufferMemoryRequirements2KHR(PFN_vkGetBufferMemoryRequirements2KHR f, VkDevice arg0,  const VkBufferMemoryRequirementsInfo2 * arg1,  VkMemoryRequirements2 * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetImageSparseMemoryRequirements2KHR(PFN_vkGetImageSparseMemoryRequirements2KHR f, VkDevice arg0,  const VkImageSparseMemoryRequirementsInfo2 * arg1,  uint32_t * arg2,  VkSparseImageMemoryRequirements2 * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkCreateSamplerYcbcrConversionKHR(PFN_vkCreateSamplerYcbcrConversionKHR f, VkDevice arg0,  const VkSamplerYcbcrConversionCreateInfo * arg1,  const VkAllocationCallbacks * arg2,  VkSamplerYcbcrConversion * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroySamplerYcbcrConversionKHR(PFN_vkDestroySamplerYcbcrConversionKHR f, VkDevice arg0,  VkSamplerYcbcrConversion arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkBindBufferMemory2KHR(PFN_vkBindBufferMemory2KHR f, VkDevice arg0,  uint32_t arg1,  const VkBindBufferMemoryInfo * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkBindImageMemory2KHR(PFN_vkBindImageMemory2KHR f, VkDevice arg0,  uint32_t arg1,  const VkBindImageMemoryInfo * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkGetDescriptorSetLayoutSupportKHR(PFN_vkGetDescriptorSetLayoutSupportKHR f, VkDevice arg0,  const VkDescriptorSetLayoutCreateInfo * arg1,  VkDescriptorSetLayoutSupport * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdDrawIndirectCountKHR(PFN_vkCmdDrawIndirectCountKHR f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkBuffer arg3,  VkDeviceSize arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkCmdDrawIndexedIndirectCountKHR(PFN_vkCmdDrawIndexedIndirectCountKHR f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkBuffer arg3,  VkDeviceSize arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

VkBool32 callPFN_vkDebugReportCallbackEXT(PFN_vkDebugReportCallbackEXT f, VkDebugReportFlagsEXT arg0,  VkDebugReportObjectTypeEXT arg1,  uint64_t arg2,  size_t arg3,  int32_t arg4,  const char * arg5,  const char * arg6,  void * arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

VkResult callPFN_vkCreateDebugReportCallbackEXT(PFN_vkCreateDebugReportCallbackEXT f, VkInstance arg0,  const VkDebugReportCallbackCreateInfoEXT * arg1,  const VkAllocationCallbacks * arg2,  VkDebugReportCallbackEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDebugReportCallbackEXT(PFN_vkDestroyDebugReportCallbackEXT f, VkInstance arg0,  VkDebugReportCallbackEXT arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkDebugReportMessageEXT(PFN_vkDebugReportMessageEXT f, VkInstance arg0,  VkDebugReportFlagsEXT arg1,  VkDebugReportObjectTypeEXT arg2,  uint64_t arg3,  size_t arg4,  int32_t arg5,  const char * arg6,  const char * arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

VkResult callPFN_vkDebugMarkerSetObjectTagEXT(PFN_vkDebugMarkerSetObjectTagEXT f, VkDevice arg0,  const VkDebugMarkerObjectTagInfoEXT * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkDebugMarkerSetObjectNameEXT(PFN_vkDebugMarkerSetObjectNameEXT f, VkDevice arg0,  const VkDebugMarkerObjectNameInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdDebugMarkerBeginEXT(PFN_vkCmdDebugMarkerBeginEXT f, VkCommandBuffer arg0,  const VkDebugMarkerMarkerInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdDebugMarkerEndEXT(PFN_vkCmdDebugMarkerEndEXT f, VkCommandBuffer arg0)
{
  return f(arg0);
};

void callPFN_vkCmdDebugMarkerInsertEXT(PFN_vkCmdDebugMarkerInsertEXT f, VkCommandBuffer arg0,  const VkDebugMarkerMarkerInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdDrawIndirectCountAMD(PFN_vkCmdDrawIndirectCountAMD f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkBuffer arg3,  VkDeviceSize arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

void callPFN_vkCmdDrawIndexedIndirectCountAMD(PFN_vkCmdDrawIndexedIndirectCountAMD f, VkCommandBuffer arg0,  VkBuffer arg1,  VkDeviceSize arg2,  VkBuffer arg3,  VkDeviceSize arg4,  uint32_t arg5,  uint32_t arg6)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
};

VkResult callPFN_vkGetShaderInfoAMD(PFN_vkGetShaderInfoAMD f, VkDevice arg0,  VkPipeline arg1,  VkShaderStageFlagBits arg2,  VkShaderInfoTypeAMD arg3,  size_t * arg4,  void * arg5)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5);
};

VkResult callPFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV(PFN_vkGetPhysicalDeviceExternalImageFormatPropertiesNV f, VkPhysicalDevice arg0,  VkFormat arg1,  VkImageType arg2,  VkImageTiling arg3,  VkImageUsageFlags arg4,  VkImageCreateFlags arg5,  VkExternalMemoryHandleTypeFlagsNV arg6,  VkExternalImageFormatPropertiesNV * arg7)
{
  return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
};

void callPFN_vkCmdBeginConditionalRenderingEXT(PFN_vkCmdBeginConditionalRenderingEXT f, VkCommandBuffer arg0,  const VkConditionalRenderingBeginInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdEndConditionalRenderingEXT(PFN_vkCmdEndConditionalRenderingEXT f, VkCommandBuffer arg0)
{
  return f(arg0);
};

void callPFN_vkCmdProcessCommandsNVX(PFN_vkCmdProcessCommandsNVX f, VkCommandBuffer arg0,  const VkCmdProcessCommandsInfoNVX * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdReserveSpaceForCommandsNVX(PFN_vkCmdReserveSpaceForCommandsNVX f, VkCommandBuffer arg0,  const VkCmdReserveSpaceForCommandsInfoNVX * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkCreateIndirectCommandsLayoutNVX(PFN_vkCreateIndirectCommandsLayoutNVX f, VkDevice arg0,  const VkIndirectCommandsLayoutCreateInfoNVX * arg1,  const VkAllocationCallbacks * arg2,  VkIndirectCommandsLayoutNVX * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyIndirectCommandsLayoutNVX(PFN_vkDestroyIndirectCommandsLayoutNVX f, VkDevice arg0,  VkIndirectCommandsLayoutNVX arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateObjectTableNVX(PFN_vkCreateObjectTableNVX f, VkDevice arg0,  const VkObjectTableCreateInfoNVX * arg1,  const VkAllocationCallbacks * arg2,  VkObjectTableNVX * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyObjectTableNVX(PFN_vkDestroyObjectTableNVX f, VkDevice arg0,  VkObjectTableNVX arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkRegisterObjectsNVX(PFN_vkRegisterObjectsNVX f, VkDevice arg0,  VkObjectTableNVX arg1,  uint32_t arg2,  const VkObjectTableEntryNVX *const * arg3,  const uint32_t * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkUnregisterObjectsNVX(PFN_vkUnregisterObjectsNVX f, VkDevice arg0,  VkObjectTableNVX arg1,  uint32_t arg2,  const VkObjectEntryTypeNVX * arg3,  const uint32_t * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX(PFN_vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX f, VkPhysicalDevice arg0,  VkDeviceGeneratedCommandsFeaturesNVX * arg1,  VkDeviceGeneratedCommandsLimitsNVX * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkCmdSetViewportWScalingNV(PFN_vkCmdSetViewportWScalingNV f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  const VkViewportWScalingNV * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkReleaseDisplayEXT(PFN_vkReleaseDisplayEXT f, VkPhysicalDevice arg0,  VkDisplayKHR arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT(PFN_vkGetPhysicalDeviceSurfaceCapabilities2EXT f, VkPhysicalDevice arg0,  VkSurfaceKHR arg1,  VkSurfaceCapabilities2EXT * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkDisplayPowerControlEXT(PFN_vkDisplayPowerControlEXT f, VkDevice arg0,  VkDisplayKHR arg1,  const VkDisplayPowerInfoEXT * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkRegisterDeviceEventEXT(PFN_vkRegisterDeviceEventEXT f, VkDevice arg0,  const VkDeviceEventInfoEXT * arg1,  const VkAllocationCallbacks * arg2,  VkFence * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkRegisterDisplayEventEXT(PFN_vkRegisterDisplayEventEXT f, VkDevice arg0,  VkDisplayKHR arg1,  const VkDisplayEventInfoEXT * arg2,  const VkAllocationCallbacks * arg3,  VkFence * arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

VkResult callPFN_vkGetSwapchainCounterEXT(PFN_vkGetSwapchainCounterEXT f, VkDevice arg0,  VkSwapchainKHR arg1,  VkSurfaceCounterFlagBitsEXT arg2,  uint64_t * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetRefreshCycleDurationGOOGLE(PFN_vkGetRefreshCycleDurationGOOGLE f, VkDevice arg0,  VkSwapchainKHR arg1,  VkRefreshCycleDurationGOOGLE * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkGetPastPresentationTimingGOOGLE(PFN_vkGetPastPresentationTimingGOOGLE f, VkDevice arg0,  VkSwapchainKHR arg1,  uint32_t * arg2,  VkPastPresentationTimingGOOGLE * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdSetDiscardRectangleEXT(PFN_vkCmdSetDiscardRectangleEXT f, VkCommandBuffer arg0,  uint32_t arg1,  uint32_t arg2,  const VkRect2D * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkSetHdrMetadataEXT(PFN_vkSetHdrMetadataEXT f, VkDevice arg0,  uint32_t arg1,  const VkSwapchainKHR * arg2,  const VkHdrMetadataEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkBool32 callPFN_vkDebugUtilsMessengerCallbackEXT(PFN_vkDebugUtilsMessengerCallbackEXT f, VkDebugUtilsMessageSeverityFlagBitsEXT arg0,  VkDebugUtilsMessageTypeFlagsEXT arg1,  const VkDebugUtilsMessengerCallbackDataEXT * arg2,  void * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkSetDebugUtilsObjectNameEXT(PFN_vkSetDebugUtilsObjectNameEXT f, VkDevice arg0,  const VkDebugUtilsObjectNameInfoEXT * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkSetDebugUtilsObjectTagEXT(PFN_vkSetDebugUtilsObjectTagEXT f, VkDevice arg0,  const VkDebugUtilsObjectTagInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkQueueBeginDebugUtilsLabelEXT(PFN_vkQueueBeginDebugUtilsLabelEXT f, VkQueue arg0,  const VkDebugUtilsLabelEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkQueueEndDebugUtilsLabelEXT(PFN_vkQueueEndDebugUtilsLabelEXT f, VkQueue arg0)
{
  return f(arg0);
};

void callPFN_vkQueueInsertDebugUtilsLabelEXT(PFN_vkQueueInsertDebugUtilsLabelEXT f, VkQueue arg0,  const VkDebugUtilsLabelEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdBeginDebugUtilsLabelEXT(PFN_vkCmdBeginDebugUtilsLabelEXT f, VkCommandBuffer arg0,  const VkDebugUtilsLabelEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkCmdEndDebugUtilsLabelEXT(PFN_vkCmdEndDebugUtilsLabelEXT f, VkCommandBuffer arg0)
{
  return f(arg0);
};

void callPFN_vkCmdInsertDebugUtilsLabelEXT(PFN_vkCmdInsertDebugUtilsLabelEXT f, VkCommandBuffer arg0,  const VkDebugUtilsLabelEXT * arg1)
{
  return f(arg0, arg1);
};

VkResult callPFN_vkCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT f, VkInstance arg0,  const VkDebugUtilsMessengerCreateInfoEXT * arg1,  const VkAllocationCallbacks * arg2,  VkDebugUtilsMessengerEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT f, VkInstance arg0,  VkDebugUtilsMessengerEXT arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

void callPFN_vkSubmitDebugUtilsMessageEXT(PFN_vkSubmitDebugUtilsMessageEXT f, VkInstance arg0,  VkDebugUtilsMessageSeverityFlagBitsEXT arg1,  VkDebugUtilsMessageTypeFlagsEXT arg2,  const VkDebugUtilsMessengerCallbackDataEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdSetSampleLocationsEXT(PFN_vkCmdSetSampleLocationsEXT f, VkCommandBuffer arg0,  const VkSampleLocationsInfoEXT * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetPhysicalDeviceMultisamplePropertiesEXT(PFN_vkGetPhysicalDeviceMultisamplePropertiesEXT f, VkPhysicalDevice arg0,  VkSampleCountFlagBits arg1,  VkMultisamplePropertiesEXT * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkCreateValidationCacheEXT(PFN_vkCreateValidationCacheEXT f, VkDevice arg0,  const VkValidationCacheCreateInfoEXT * arg1,  const VkAllocationCallbacks * arg2,  VkValidationCacheEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkDestroyValidationCacheEXT(PFN_vkDestroyValidationCacheEXT f, VkDevice arg0,  VkValidationCacheEXT arg1,  const VkAllocationCallbacks * arg2)
{
  return f(arg0, arg1, arg2);
};

VkResult callPFN_vkMergeValidationCachesEXT(PFN_vkMergeValidationCachesEXT f, VkDevice arg0,  VkValidationCacheEXT arg1,  uint32_t arg2,  const VkValidationCacheEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetValidationCacheDataEXT(PFN_vkGetValidationCacheDataEXT f, VkDevice arg0,  VkValidationCacheEXT arg1,  size_t * arg2,  void * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

VkResult callPFN_vkGetMemoryHostPointerPropertiesEXT(PFN_vkGetMemoryHostPointerPropertiesEXT f, VkDevice arg0,  VkExternalMemoryHandleTypeFlagBits arg1,  const void * arg2,  VkMemoryHostPointerPropertiesEXT * arg3)
{
  return f(arg0, arg1, arg2, arg3);
};

void callPFN_vkCmdWriteBufferMarkerAMD(PFN_vkCmdWriteBufferMarkerAMD f, VkCommandBuffer arg0,  VkPipelineStageFlagBits arg1,  VkBuffer arg2,  VkDeviceSize arg3,  uint32_t arg4)
{
  return f(arg0, arg1, arg2, arg3, arg4);
};

void callPFN_vkCmdSetCheckpointNV(PFN_vkCmdSetCheckpointNV f, VkCommandBuffer arg0,  const void * arg1)
{
  return f(arg0, arg1);
};

void callPFN_vkGetQueueCheckpointDataNV(PFN_vkGetQueueCheckpointDataNV f, VkQueue arg0,  uint32_t * arg1,  VkCheckpointDataNV * arg2)
{
  return f(arg0, arg1, arg2);
};

