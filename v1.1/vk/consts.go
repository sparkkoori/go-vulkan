package vk

func MakeVersion(major, minor, patch uint32) uint32 {
	return (((major) << 22) | ((minor) << 12) | (patch))
}

func VersionMajor(version uint32) uint32 {
	return version >> 22
}

func VersionMinor(version uint32) uint32 {
	return (version >> 12) & 0x3ff // 10 bits
}

func VersionPatch(version uint32) uint32 {
	return version & 0xfff // 12 bits
}

const HeaderVersion uint32 = 82
