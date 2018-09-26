package main

import (
	"log"

	"github.com/sparkkoori/go-vulkan/v1.1/vk"
)

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
	rs := vk.CreateInstance(createInfo, nil, &ins)
	if rs == vk.ERROR_INCOMPATIBLE_DRIVER {
		log.Fatalln("cannot find a compatible Vulkan ICD")
	} else if rs != vk.SUCCESS {
		log.Fatalln("unknown error")
	}

	vk.DestroyInstance(ins, nil)
}
