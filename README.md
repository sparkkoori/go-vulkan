# go-vulkan
Go binding for Vulkan

## Rules

- Prefix was trimed.

  - eg. Vk, vk, VK_, p, pp, s


- No translation for handle types.

  - eg. `type Instance C.VkInstance`


- Merge `Vk*FlagBit` to `Vk*Flags`.

- Map `Bool32` to `bool`.

- Map `char *` to `string`.

- Merge size and pointer fields to slice.
  - eg. `EnabledLayerNames     []string` for
  `uint32_t enabledLayerCount` and `const char* const* ppEnabledLayerNames`


- Remove `sType` field, it can be inferred from struct name.

- Map `void *` to `Structure` in `pNext` field.

- Map function pointer types to structs with Call() method which calls bridge function.

## License
MIT
