package main

import (
	"github.com/sparkkoori/go-vulkan/v1.1/vk"
)

type ccc uintptr
type bbb uintptr

func main() {
	appInfo := &vk.ApplicationInfo{
		ApplicationName:    "AppName",
		ApplicationVersion: 1,
		EngineName:         "AppName",
		EngineVersion:      1,
		ApiVersion:         vk.MakeVersion(1, 1, 0),
	}

	createInfo := &vk.InstanceCreateInfo{
		Flags:                 0,
		ApplicationInfo:       appInfo,
		EnabledLayerNames:     nil,
		EnabledExtensionNames: nil,
	}

	var ins vk.Instance

	//Directly call
	rs := vk.CreateInstance(createInfo, nil, &ins)
	if rs == vk.ERROR_INCOMPATIBLE_DRIVER {
		panic("cannot find a compatible Vulkan ICD")
	} else if rs != vk.SUCCESS {
		panic("unknown error")
	}
	vk.DestroyInstance(ins, nil)

	//GetInstanceProcAddr
	pfn := vk.GetInstanceProcAddr(nil, "vkCreateInstance")
	rs = vk.PFNCreateInstance(pfn).Call(createInfo, nil, &ins)
	if rs == vk.ERROR_INCOMPATIBLE_DRIVER {
		panic("cannot find a compatible Vulkan ICD")
	} else if rs != vk.SUCCESS {
		panic("unknown error")
	}
	pfn = vk.GetInstanceProcAddr(ins, "vkDestroyInstance")
	vk.PFNDestroyInstance(pfn).Call(ins, nil)
}
