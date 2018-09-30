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

const HEADER_VERSION uint32 = 82
