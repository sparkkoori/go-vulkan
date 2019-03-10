package vkerr

import (
	"errors"

	"github.com/sparkkoori/go-vulkan/v1.1/vk"
)

//Success Codes:
var SUCCESS = error(nil) //Command successfully completed
var NOT_READY = errors.New("VK_NOT_READY: A fence or query has not yet completed")
var TIMEOUT = errors.New("VK_TIMEOUT: A wait operation has not completed in the specified time")
var EVENT_SET = errors.New("VK_EVENT_SET: An event is signaled")
var EVENT_RESET = errors.New("VK_EVENT_RESET: An event is unsignaled")
var INCOMPLETE = errors.New("VK_INCOMPLETE: A return array was too small for the result")
var SUBOPTIMAL_KHR = errors.New("VK_SUBOPTIMAL_KHR: A swapchain no longer matches the surface properties exactly, but can still be used to present to the surface successfully")

//Error codes:
var ERROR_OUT_OF_HOST_MEMORY = errors.New("VK_ERROR_OUT_OF_HOST_MEMORY: A host memory allocation has failed")
var ERROR_OUT_OF_DEVICE_MEMORY = errors.New("VK_ERROR_OUT_OF_DEVICE_MEMORY: A device memory allocation has failed")
var ERROR_INITIALIZATION_FAILED = errors.New("VK_ERROR_INITIALIZATION_FAILED: Initialization of an object could not be completed for implementation-specific reasons")
var ERROR_DEVICE_LOST = errors.New("VK_ERROR_DEVICE_LOST: The logical or physical device has been lost")
var ERROR_MEMORY_MAP_FAILED = errors.New("VK_ERROR_MEMORY_MAP_FAILED: Mapping of a memory object has failed")
var ERROR_LAYER_NOT_PRESENT = errors.New("VK_ERROR_LAYER_NOT_PRESENT: A requested layer is not present or could not be loaded")
var ERROR_EXTENSION_NOT_PRESENT = errors.New("VK_ERROR_EXTENSION_NOT_PRESENT: A requested extension is not supported")
var ERROR_FEATURE_NOT_PRESENT = errors.New("VK_ERROR_FEATURE_NOT_PRESENT: A requested feature is not supported")
var ERROR_INCOMPATIBLE_DRIVER = errors.New("VK_ERROR_INCOMPATIBLE_DRIVER: The requested version of Vulkan is not supported by the driver or is otherwise incompatible for implementation-specific reasons")
var ERROR_TOO_MANY_OBJECTS = errors.New("VK_ERROR_TOO_MANY_OBJECTS: Too many objects of the type have already been created")
var ERROR_FORMAT_NOT_SUPPORTED = errors.New("VK_ERROR_FORMAT_NOT_SUPPORTED: A requested format is not supported on this device")
var ERROR_FRAGMENTED_POOL = errors.New("VK_ERROR_FRAGMENTED_POOL: A pool allocation has failed due to fragmentation of the pool’s memory")
var ERROR_OUT_OF_POOL_MEMORY = errors.New("VK_ERROR_OUT_OF_POOL_MEMORY: A pool memory allocation has failed")
var ERROR_INVALID_EXTERNAL_HANDLE = errors.New("VK_ERROR_INVALID_EXTERNAL_HANDLE: An external handle is not a valid handle of the specified type")

var ERROR_SURFACE_LOST_KHR = errors.New("VK_ERROR_SURFACE_LOST_KHR: A surface is no longer available")
var ERROR_NATIVE_WINDOW_IN_USE_KHR = errors.New("VK_ERROR_NATIVE_WINDOW_IN_USE_KHR: The requested window is already in use by Vulkan or another API in a manner which prevents it from being used again")
var ERROR_OUT_OF_DATE_KHR = errors.New("VK_ERROR_OUT_OF_DATE_KHR: A surface has changed in such a way that it is no longer compatible with the swapchain, and further presentation requests using the swapchain will fail")
var ERROR_INCOMPATIBLE_DISPLAY_KHR = errors.New("VK_ERROR_INCOMPATIBLE_DISPLAY_KHR: The display used by a swapchain does not use the same presentable image layout, or is incompatible in a way that prevents sharing an image")
var ERROR_INVALID_SHADER_NV = errors.New("VK_ERROR_INVALID_SHADER_NV: One or more shaders failed to compile or link")
var ERROR_FRAGMENTATION_EXT = errors.New("VK_ERROR_FRAGMENTATION_EXT: A descriptor pool creation has failed due to fragmentation")
var ERROR_VALIDATION_FAILED_EXT = errors.New("ERROR_VALIDATION_FAILED_EXT")
var ERROR_NOT_PERMITTED_EXT = errors.New("ERROR_NOT_PERMITTED_EXT")
var ERROR_OUT_OF_POOL_MEMORY_KHR = errors.New("ERROR_OUT_OF_POOL_MEMORY_KHR")
var ERROR_INVALID_EXTERNAL_HANDLE_KHR = errors.New("ERROR_INVALID_EXTERNAL_HANDLE_KHR")

var UNKOWN = errors.New("Unknown VkResult")

func Convert(rs vk.Result) error {
	switch rs {
	case vk.SUCCESS:
		return SUCCESS
	case vk.NOT_READY:
		return NOT_READY
	case vk.TIMEOUT:
		return TIMEOUT
	case vk.EVENT_SET:
		return EVENT_SET
	case vk.EVENT_RESET:
		return EVENT_RESET
	case vk.INCOMPLETE:
		return INCOMPLETE
	case vk.SUBOPTIMAL_KHR:
		return SUBOPTIMAL_KHR

	case vk.ERROR_OUT_OF_HOST_MEMORY:
		return ERROR_OUT_OF_HOST_MEMORY
	case vk.ERROR_OUT_OF_DEVICE_MEMORY:
		return ERROR_OUT_OF_DEVICE_MEMORY
	case vk.ERROR_INITIALIZATION_FAILED:
		return ERROR_INITIALIZATION_FAILED
	case vk.ERROR_DEVICE_LOST:
		return ERROR_DEVICE_LOST
	case vk.ERROR_MEMORY_MAP_FAILED:
		return ERROR_MEMORY_MAP_FAILED
	case vk.ERROR_LAYER_NOT_PRESENT:
		return ERROR_LAYER_NOT_PRESENT
	case vk.ERROR_EXTENSION_NOT_PRESENT:
		return ERROR_EXTENSION_NOT_PRESENT
	case vk.ERROR_FEATURE_NOT_PRESENT:
		return ERROR_FEATURE_NOT_PRESENT
	case vk.ERROR_INCOMPATIBLE_DRIVER:
		return ERROR_INCOMPATIBLE_DRIVER
	case vk.ERROR_TOO_MANY_OBJECTS:
		return ERROR_TOO_MANY_OBJECTS
	case vk.ERROR_FORMAT_NOT_SUPPORTED:
		return ERROR_FORMAT_NOT_SUPPORTED
	case vk.ERROR_FRAGMENTED_POOL:
		return ERROR_FRAGMENTED_POOL
	case vk.ERROR_OUT_OF_POOL_MEMORY:
		return ERROR_OUT_OF_POOL_MEMORY
	case vk.ERROR_INVALID_EXTERNAL_HANDLE:
		return ERROR_INVALID_EXTERNAL_HANDLE

	case vk.ERROR_SURFACE_LOST_KHR:
		return ERROR_SURFACE_LOST_KHR
	case vk.ERROR_NATIVE_WINDOW_IN_USE_KHR:
		return ERROR_NATIVE_WINDOW_IN_USE_KHR
	case vk.ERROR_OUT_OF_DATE_KHR:
		return ERROR_OUT_OF_DATE_KHR
	case vk.ERROR_INCOMPATIBLE_DISPLAY_KHR:
		return ERROR_INCOMPATIBLE_DISPLAY_KHR
	case vk.ERROR_INVALID_SHADER_NV:
		return ERROR_INVALID_SHADER_NV
	case vk.ERROR_FRAGMENTATION_EXT:
		return ERROR_FRAGMENTATION_EXT
	case vk.ERROR_VALIDATION_FAILED_EXT:
		return ERROR_VALIDATION_FAILED_EXT
	case vk.ERROR_NOT_PERMITTED_EXT:
		return ERROR_NOT_PERMITTED_EXT
	// case vk.ERROR_OUT_OF_POOL_MEMORY_KHR:
	// 	return ERROR_OUT_OF_POOL_MEMORY_KHR
	// case vk.ERROR_INVALID_EXTERNAL_HANDLE_KHR:
	// 	return ERROR_INVALID_EXTERNAL_HANDLE_KHR

	default:
		return UNKOWN
	}
}
