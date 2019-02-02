package main

import (
	"github.com/sparkkoori/go-vulkan/v1.1/vk"
)

func main() {
	appInfo := &vk.ApplicationInfo{
		ApplicationName:    "AppName",
		ApplicationVersion: 1,
		EngineName:         "AppName",
		EngineVersion:      1,
		ApiVersion:         vk.MAKE_VERSION(1, 1, 0),
	}

	createInfo := &vk.InstanceCreateInfo{
		Flags:                 0,
		ApplicationInfo:       appInfo,
		EnabledLayerNames:     nil,
		EnabledExtensionNames: nil,
	}

	var ins vk.Instance

	//Directly Call
	rs := vk.CreateInstance(createInfo, nil, &ins)
	if rs == vk.ERROR_INCOMPATIBLE_DRIVER {
		panic("cannot find a compatible Vulkan ICD")
	} else if rs != vk.SUCCESS {
		panic("unknown error")
	}
	vk.DestroyInstance(ins, nil)

	//Indirectly Call
	{
		fn := vk.ToCreateInstance(vk.GetInstanceProcAddr(nil, "vkCreateInstance"))
		rs = fn(createInfo, nil, &ins)
		if rs == vk.ERROR_INCOMPATIBLE_DRIVER {
			panic("cannot find a compatible Vulkan ICD")
		} else if rs != vk.SUCCESS {
			panic("unknown error")
		}
	}
	{
		fn := vk.ToDestroyInstance(vk.GetInstanceProcAddr(ins, "vkDestroyInstance"))
		fn(ins, nil)
	}
}
