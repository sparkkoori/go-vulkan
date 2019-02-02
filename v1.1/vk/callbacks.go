package vk

/*
#include "vulkan/vulkan.h"
extern VkBool32 defaultPFN_vkDebugUtilsMessengerCallbackEXT (
    VkDebugUtilsMessageSeverityFlagBitsEXT           messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT                  messageType,
    VkDebugUtilsMessengerCallbackDataEXT*      pCallbackData,
    void*                                            pUserData);

*/
import "C"
import "unsafe"

//export defaultPFN_vkDebugUtilsMessengerCallbackEXT
func defaultPFN_vkDebugUtilsMessengerCallbackEXT(
	messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT,
	messageType C.VkDebugUtilsMessageTypeFlagsEXT,
	pCallbackData *C.VkDebugUtilsMessengerCallbackDataEXT,
	pUserData unsafe.Pointer) (_ret C.VkBool32) {
	if pUserData == nil {
		panic("Param pUserData is nil")
	}
	var g struct {
		messageSeverity DebugUtilsMessageSeverityFlagBitsEXT
		messageType     DebugUtilsMessageTypeFlagsEXT
		callbackData    *DebugUtilsMessengerCallbackDataEXT
		_ret            bool
	}
	g.messageSeverity = DebugUtilsMessageSeverityFlagBitsEXT(messageSeverity)
	g.messageType = DebugUtilsMessageTypeFlagsEXT(messageType)
	if pCallbackData != nil {
		g.callbackData = &DebugUtilsMessengerCallbackDataEXT{}
		g.callbackData.fromC(pCallbackData)
	}

	ptr, ok := Registry.Access(uintptr(pUserData))
	if !ok {
		panic("Unknown pUserData")
	}
	pfn := ptr.(*func(
		messageSeverity DebugUtilsMessageSeverityFlagBitsEXT,
		messageType DebugUtilsMessageTypeFlagsEXT,
		callbackData *DebugUtilsMessengerCallbackDataEXT,
	) (_ret bool))
	g._ret = (*pfn)(g.messageSeverity, g.messageType, g.callbackData)
	if g._ret {
		_ret = 1
	}
	return
}

func DefaultDebugUtilsMessengerCallbackEXT() PFNDebugUtilsMessengerCallbackEXT {
	return PFNDebugUtilsMessengerCallbackEXT(C.defaultPFN_vkDebugUtilsMessengerCallbackEXT)
}
