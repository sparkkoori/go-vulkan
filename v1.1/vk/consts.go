package vk

func MAKE_VERSION(major, minor, patch uint32) uint32 {
	return (((major) << 22) | ((minor) << 12) | (patch))
}

func VERSION_MAJOR(version uint32) uint32 {
	return version >> 22
}

func VERSION_MINOR(version uint32) uint32 {
	return (version >> 12) & 0x3ff // 10 bits
}

func VERSION_PATCH(version uint32) uint32 {
	return version & 0xfff // 12 bits
}

const HEADER_VERSION = 82
const LOD_CLAMP_NONE = 1000.0
const REMAINING_MIP_LEVELS = ^uint32(0)
const REMAINING_ARRAY_LAYERS = ^uint32(0)
const WHOLE_SIZE = ^uint64(0)
const ATTACHMENT_UNUSED = ^uint32(0)

// const TRUE                        =    1
// const FALSE                       =    0
const QUEUE_FAMILY_IGNORED = ^uint32(0)
const SUBPASS_EXTERNAL = ^uint32(0)
const MAX_PHYSICAL_DEVICE_NAME_SIZE = 256
const UUID_SIZE = 16
const MAX_MEMORY_TYPES = 32
const MAX_MEMORY_HEAPS = 16
const MAX_EXTENSION_NAME_SIZE = 256
const MAX_DESCRIPTION_SIZE = 256
