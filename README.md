# go-vulkan
Go binding for Vulkan

## Rules

- Prefix was trimed.

  - eg. Vk, vk, VK_, p, pp, s


- No translation for handle types.

  - eg. `type Instance C.VkInstance`


- Merge Vk\*FlagBit to Vk\*Flags.

- Use bool instead of Bool32.


## License
MIT
