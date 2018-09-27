# go-vulkan
Go binding for Vulkan

## Usage

### How to use

Fetch the package:

```
$ go get -u github.com/sparkkoori/go-vulkan
```
Install the [Vulkan SDK](https://www.lunarg.com/vulkan-sdk/), make sure loader vulkan-1 or vulkan.1 exists in search path.

Use `vk` in your package:

```
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

```

### Rules

The warping rules used in this package.

#### Type Rules

- No translation for handle types.

  - eg. `type Instance C.VkInstance`


- Map `Bool32` to `bool`.

- Map `char *` to `string`.

- Merge size and pointer fields to slice.

  - eg. `EnabledLayerNames     []string` for
  `uint32_t enabledLayerCount` and `const char* const* ppEnabledLayerNames`


- Remove `sType` field, it can be inferred from struct name.

- Map `void *` to `Structure` in `pNext` field.

- Map `PFN*` to `uintptr`.

	- This allow conversion between `PFN*`.

- Map `union` to `struct`, keep it's memory layout.


#### Memory Rules

- Copy data between go and c, instead of pass memory.

  - But no copy for `unsafe.Pointer`.


- No manual memory management is need.


#### Call Rules

- Because cgo do not support calling function pointer, so for every `PFN*`, a bridge method `Call()` is created.

#### Name Rules

- Prefix has been trimed.

  - eg. Vk, vk, VK_, p, pp

## License
MIT
