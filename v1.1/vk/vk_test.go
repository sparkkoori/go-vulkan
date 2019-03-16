package vk

import (
	"testing"
)

func TestCreateInstance(t *testing.T) {
	appInfo := &ApplicationInfo{
		ApplicationName:    "AppName",
		ApplicationVersion: 1,
		EngineName:         "AppName",
		EngineVersion:      1,
		ApiVersion:         MAKE_VERSION(1, 1, 0),
	}

	createInfo := &InstanceCreateInfo{
		Flags:                 0,
		ApplicationInfo:       appInfo,
		EnabledLayerNames:     nil,
		EnabledExtensionNames: nil,
	}

	var ins Instance

	//Directly Call
	rs := CreateInstance(createInfo, nil, &ins)
	if rs == ERROR_INCOMPATIBLE_DRIVER {
		t.Fatal("cannot find a compatible Vulkan ICD")
	} else if rs != SUCCESS {
		t.Fatal("unknown error")
	}
	DestroyInstance(ins, nil)

	//Indirectly Call
	{
		fn := ToCreateInstance(GetInstanceProcAddr(nil, "vkCreateInstance"))
		rs = fn(createInfo, nil, &ins)
		if rs == ERROR_INCOMPATIBLE_DRIVER {
			t.Fatal("cannot find a compatible Vulkan ICD")
		} else if rs != SUCCESS {
			t.Fatal("unknown error")
		}
	}
	{
		fn := ToDestroyInstance(GetInstanceProcAddr(ins, "vkDestroyInstance"))
		fn(ins, nil)
	}
}
