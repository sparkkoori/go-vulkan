package vk

//#include "vulkan/vulkan.h"
import "C"
import "unsafe"

type Flags uint32
type DeviceSize uint64
type SampleMask uint32
type Instance C.VkInstance
type PhysicalDevice C.VkPhysicalDevice
type Device C.VkDevice
type Queue C.VkQueue
type Semaphore C.VkSemaphore
type CommandBuffer C.VkCommandBuffer
type Fence C.VkFence
type DeviceMemory C.VkDeviceMemory
type Buffer C.VkBuffer
type Image C.VkImage
type Event C.VkEvent
type QueryPool C.VkQueryPool
type BufferView C.VkBufferView
type ImageView C.VkImageView
type ShaderModule C.VkShaderModule
type PipelineCache C.VkPipelineCache
type PipelineLayout C.VkPipelineLayout
type RenderPass C.VkRenderPass
type Pipeline C.VkPipeline
type DescriptorSetLayout C.VkDescriptorSetLayout
type Sampler C.VkSampler
type DescriptorPool C.VkDescriptorPool
type DescriptorSet C.VkDescriptorSet
type Framebuffer C.VkFramebuffer
type CommandPool C.VkCommandPool
type PipelineCacheHeaderVersion int

const (
	PIPELINE_CACHE_HEADER_VERSION_ONE         PipelineCacheHeaderVersion = 1
	PIPELINE_CACHE_HEADER_VERSION_BEGIN_RANGE PipelineCacheHeaderVersion = PIPELINE_CACHE_HEADER_VERSION_ONE
	PIPELINE_CACHE_HEADER_VERSION_END_RANGE   PipelineCacheHeaderVersion = PIPELINE_CACHE_HEADER_VERSION_ONE
	PIPELINE_CACHE_HEADER_VERSION_RANGE_SIZE  PipelineCacheHeaderVersion = (PIPELINE_CACHE_HEADER_VERSION_ONE - PIPELINE_CACHE_HEADER_VERSION_ONE + 1)
	PIPELINE_CACHE_HEADER_VERSION_MAX_ENUM    PipelineCacheHeaderVersion = 2147483647
)

type Result int

const (
	SUCCESS                           Result = 0
	NOT_READY                         Result = 1
	TIMEOUT                           Result = 2
	EVENT_SET                         Result = 3
	EVENT_RESET                       Result = 4
	INCOMPLETE                        Result = 5
	ERROR_OUT_OF_HOST_MEMORY          Result = -1
	ERROR_OUT_OF_DEVICE_MEMORY        Result = -2
	ERROR_INITIALIZATION_FAILED       Result = -3
	ERROR_DEVICE_LOST                 Result = -4
	ERROR_MEMORY_MAP_FAILED           Result = -5
	ERROR_LAYER_NOT_PRESENT           Result = -6
	ERROR_EXTENSION_NOT_PRESENT       Result = -7
	ERROR_FEATURE_NOT_PRESENT         Result = -8
	ERROR_INCOMPATIBLE_DRIVER         Result = -9
	ERROR_TOO_MANY_OBJECTS            Result = -10
	ERROR_FORMAT_NOT_SUPPORTED        Result = -11
	ERROR_FRAGMENTED_POOL             Result = -12
	ERROR_OUT_OF_POOL_MEMORY          Result = -1000069000
	ERROR_INVALID_EXTERNAL_HANDLE     Result = -1000072003
	ERROR_SURFACE_LOST_KHR            Result = -1000000000
	ERROR_NATIVE_WINDOW_IN_USE_KHR    Result = -1000000001
	SUBOPTIMAL_KHR                    Result = 1000001003
	ERROR_OUT_OF_DATE_KHR             Result = -1000001004
	ERROR_INCOMPATIBLE_DISPLAY_KHR    Result = -1000003001
	ERROR_VALIDATION_FAILED_EXT       Result = -1000011001
	ERROR_INVALID_SHADER_NV           Result = -1000012000
	ERROR_FRAGMENTATION_EXT           Result = -1000161000
	ERROR_NOT_PERMITTED_EXT           Result = -1000174001
	ERROR_OUT_OF_POOL_MEMORY_KHR      Result = ERROR_OUT_OF_POOL_MEMORY
	ERROR_INVALID_EXTERNAL_HANDLE_KHR Result = ERROR_INVALID_EXTERNAL_HANDLE
	RESULT_BEGIN_RANGE                Result = ERROR_FRAGMENTED_POOL
	RESULT_END_RANGE                  Result = INCOMPLETE
	RESULT_RANGE_SIZE                 Result = (INCOMPLETE - ERROR_FRAGMENTED_POOL + 1)
	RESULT_MAX_ENUM                   Result = 2147483647
)

type SystemAllocationScope int

const (
	SYSTEM_ALLOCATION_SCOPE_COMMAND     SystemAllocationScope = 0
	SYSTEM_ALLOCATION_SCOPE_OBJECT      SystemAllocationScope = 1
	SYSTEM_ALLOCATION_SCOPE_CACHE       SystemAllocationScope = 2
	SYSTEM_ALLOCATION_SCOPE_DEVICE      SystemAllocationScope = 3
	SYSTEM_ALLOCATION_SCOPE_INSTANCE    SystemAllocationScope = 4
	SYSTEM_ALLOCATION_SCOPE_BEGIN_RANGE SystemAllocationScope = SYSTEM_ALLOCATION_SCOPE_COMMAND
	SYSTEM_ALLOCATION_SCOPE_END_RANGE   SystemAllocationScope = SYSTEM_ALLOCATION_SCOPE_INSTANCE
	SYSTEM_ALLOCATION_SCOPE_RANGE_SIZE  SystemAllocationScope = (SYSTEM_ALLOCATION_SCOPE_INSTANCE - SYSTEM_ALLOCATION_SCOPE_COMMAND + 1)
	SYSTEM_ALLOCATION_SCOPE_MAX_ENUM    SystemAllocationScope = 2147483647
)

type InternalAllocationType int

const (
	INTERNAL_ALLOCATION_TYPE_EXECUTABLE  InternalAllocationType = 0
	INTERNAL_ALLOCATION_TYPE_BEGIN_RANGE InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_END_RANGE   InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_RANGE_SIZE  InternalAllocationType = (INTERNAL_ALLOCATION_TYPE_EXECUTABLE - INTERNAL_ALLOCATION_TYPE_EXECUTABLE + 1)
	INTERNAL_ALLOCATION_TYPE_MAX_ENUM    InternalAllocationType = 2147483647
)

type Format int

const (
	FORMAT_UNDEFINED                                      Format = 0
	FORMAT_R4G4_UNORM_PACK8                               Format = 1
	FORMAT_R4G4B4A4_UNORM_PACK16                          Format = 2
	FORMAT_B4G4R4A4_UNORM_PACK16                          Format = 3
	FORMAT_R5G6B5_UNORM_PACK16                            Format = 4
	FORMAT_B5G6R5_UNORM_PACK16                            Format = 5
	FORMAT_R5G5B5A1_UNORM_PACK16                          Format = 6
	FORMAT_B5G5R5A1_UNORM_PACK16                          Format = 7
	FORMAT_A1R5G5B5_UNORM_PACK16                          Format = 8
	FORMAT_R8_UNORM                                       Format = 9
	FORMAT_R8_SNORM                                       Format = 10
	FORMAT_R8_USCALED                                     Format = 11
	FORMAT_R8_SSCALED                                     Format = 12
	FORMAT_R8_UINT                                        Format = 13
	FORMAT_R8_SINT                                        Format = 14
	FORMAT_R8_SRGB                                        Format = 15
	FORMAT_R8G8_UNORM                                     Format = 16
	FORMAT_R8G8_SNORM                                     Format = 17
	FORMAT_R8G8_USCALED                                   Format = 18
	FORMAT_R8G8_SSCALED                                   Format = 19
	FORMAT_R8G8_UINT                                      Format = 20
	FORMAT_R8G8_SINT                                      Format = 21
	FORMAT_R8G8_SRGB                                      Format = 22
	FORMAT_R8G8B8_UNORM                                   Format = 23
	FORMAT_R8G8B8_SNORM                                   Format = 24
	FORMAT_R8G8B8_USCALED                                 Format = 25
	FORMAT_R8G8B8_SSCALED                                 Format = 26
	FORMAT_R8G8B8_UINT                                    Format = 27
	FORMAT_R8G8B8_SINT                                    Format = 28
	FORMAT_R8G8B8_SRGB                                    Format = 29
	FORMAT_B8G8R8_UNORM                                   Format = 30
	FORMAT_B8G8R8_SNORM                                   Format = 31
	FORMAT_B8G8R8_USCALED                                 Format = 32
	FORMAT_B8G8R8_SSCALED                                 Format = 33
	FORMAT_B8G8R8_UINT                                    Format = 34
	FORMAT_B8G8R8_SINT                                    Format = 35
	FORMAT_B8G8R8_SRGB                                    Format = 36
	FORMAT_R8G8B8A8_UNORM                                 Format = 37
	FORMAT_R8G8B8A8_SNORM                                 Format = 38
	FORMAT_R8G8B8A8_USCALED                               Format = 39
	FORMAT_R8G8B8A8_SSCALED                               Format = 40
	FORMAT_R8G8B8A8_UINT                                  Format = 41
	FORMAT_R8G8B8A8_SINT                                  Format = 42
	FORMAT_R8G8B8A8_SRGB                                  Format = 43
	FORMAT_B8G8R8A8_UNORM                                 Format = 44
	FORMAT_B8G8R8A8_SNORM                                 Format = 45
	FORMAT_B8G8R8A8_USCALED                               Format = 46
	FORMAT_B8G8R8A8_SSCALED                               Format = 47
	FORMAT_B8G8R8A8_UINT                                  Format = 48
	FORMAT_B8G8R8A8_SINT                                  Format = 49
	FORMAT_B8G8R8A8_SRGB                                  Format = 50
	FORMAT_A8B8G8R8_UNORM_PACK32                          Format = 51
	FORMAT_A8B8G8R8_SNORM_PACK32                          Format = 52
	FORMAT_A8B8G8R8_USCALED_PACK32                        Format = 53
	FORMAT_A8B8G8R8_SSCALED_PACK32                        Format = 54
	FORMAT_A8B8G8R8_UINT_PACK32                           Format = 55
	FORMAT_A8B8G8R8_SINT_PACK32                           Format = 56
	FORMAT_A8B8G8R8_SRGB_PACK32                           Format = 57
	FORMAT_A2R10G10B10_UNORM_PACK32                       Format = 58
	FORMAT_A2R10G10B10_SNORM_PACK32                       Format = 59
	FORMAT_A2R10G10B10_USCALED_PACK32                     Format = 60
	FORMAT_A2R10G10B10_SSCALED_PACK32                     Format = 61
	FORMAT_A2R10G10B10_UINT_PACK32                        Format = 62
	FORMAT_A2R10G10B10_SINT_PACK32                        Format = 63
	FORMAT_A2B10G10R10_UNORM_PACK32                       Format = 64
	FORMAT_A2B10G10R10_SNORM_PACK32                       Format = 65
	FORMAT_A2B10G10R10_USCALED_PACK32                     Format = 66
	FORMAT_A2B10G10R10_SSCALED_PACK32                     Format = 67
	FORMAT_A2B10G10R10_UINT_PACK32                        Format = 68
	FORMAT_A2B10G10R10_SINT_PACK32                        Format = 69
	FORMAT_R16_UNORM                                      Format = 70
	FORMAT_R16_SNORM                                      Format = 71
	FORMAT_R16_USCALED                                    Format = 72
	FORMAT_R16_SSCALED                                    Format = 73
	FORMAT_R16_UINT                                       Format = 74
	FORMAT_R16_SINT                                       Format = 75
	FORMAT_R16_SFLOAT                                     Format = 76
	FORMAT_R16G16_UNORM                                   Format = 77
	FORMAT_R16G16_SNORM                                   Format = 78
	FORMAT_R16G16_USCALED                                 Format = 79
	FORMAT_R16G16_SSCALED                                 Format = 80
	FORMAT_R16G16_UINT                                    Format = 81
	FORMAT_R16G16_SINT                                    Format = 82
	FORMAT_R16G16_SFLOAT                                  Format = 83
	FORMAT_R16G16B16_UNORM                                Format = 84
	FORMAT_R16G16B16_SNORM                                Format = 85
	FORMAT_R16G16B16_USCALED                              Format = 86
	FORMAT_R16G16B16_SSCALED                              Format = 87
	FORMAT_R16G16B16_UINT                                 Format = 88
	FORMAT_R16G16B16_SINT                                 Format = 89
	FORMAT_R16G16B16_SFLOAT                               Format = 90
	FORMAT_R16G16B16A16_UNORM                             Format = 91
	FORMAT_R16G16B16A16_SNORM                             Format = 92
	FORMAT_R16G16B16A16_USCALED                           Format = 93
	FORMAT_R16G16B16A16_SSCALED                           Format = 94
	FORMAT_R16G16B16A16_UINT                              Format = 95
	FORMAT_R16G16B16A16_SINT                              Format = 96
	FORMAT_R16G16B16A16_SFLOAT                            Format = 97
	FORMAT_R32_UINT                                       Format = 98
	FORMAT_R32_SINT                                       Format = 99
	FORMAT_R32_SFLOAT                                     Format = 100
	FORMAT_R32G32_UINT                                    Format = 101
	FORMAT_R32G32_SINT                                    Format = 102
	FORMAT_R32G32_SFLOAT                                  Format = 103
	FORMAT_R32G32B32_UINT                                 Format = 104
	FORMAT_R32G32B32_SINT                                 Format = 105
	FORMAT_R32G32B32_SFLOAT                               Format = 106
	FORMAT_R32G32B32A32_UINT                              Format = 107
	FORMAT_R32G32B32A32_SINT                              Format = 108
	FORMAT_R32G32B32A32_SFLOAT                            Format = 109
	FORMAT_R64_UINT                                       Format = 110
	FORMAT_R64_SINT                                       Format = 111
	FORMAT_R64_SFLOAT                                     Format = 112
	FORMAT_R64G64_UINT                                    Format = 113
	FORMAT_R64G64_SINT                                    Format = 114
	FORMAT_R64G64_SFLOAT                                  Format = 115
	FORMAT_R64G64B64_UINT                                 Format = 116
	FORMAT_R64G64B64_SINT                                 Format = 117
	FORMAT_R64G64B64_SFLOAT                               Format = 118
	FORMAT_R64G64B64A64_UINT                              Format = 119
	FORMAT_R64G64B64A64_SINT                              Format = 120
	FORMAT_R64G64B64A64_SFLOAT                            Format = 121
	FORMAT_B10G11R11_UFLOAT_PACK32                        Format = 122
	FORMAT_E5B9G9R9_UFLOAT_PACK32                         Format = 123
	FORMAT_D16_UNORM                                      Format = 124
	FORMAT_X8_D24_UNORM_PACK32                            Format = 125
	FORMAT_D32_SFLOAT                                     Format = 126
	FORMAT_S8_UINT                                        Format = 127
	FORMAT_D16_UNORM_S8_UINT                              Format = 128
	FORMAT_D24_UNORM_S8_UINT                              Format = 129
	FORMAT_D32_SFLOAT_S8_UINT                             Format = 130
	FORMAT_BC1_RGB_UNORM_BLOCK                            Format = 131
	FORMAT_BC1_RGB_SRGB_BLOCK                             Format = 132
	FORMAT_BC1_RGBA_UNORM_BLOCK                           Format = 133
	FORMAT_BC1_RGBA_SRGB_BLOCK                            Format = 134
	FORMAT_BC2_UNORM_BLOCK                                Format = 135
	FORMAT_BC2_SRGB_BLOCK                                 Format = 136
	FORMAT_BC3_UNORM_BLOCK                                Format = 137
	FORMAT_BC3_SRGB_BLOCK                                 Format = 138
	FORMAT_BC4_UNORM_BLOCK                                Format = 139
	FORMAT_BC4_SNORM_BLOCK                                Format = 140
	FORMAT_BC5_UNORM_BLOCK                                Format = 141
	FORMAT_BC5_SNORM_BLOCK                                Format = 142
	FORMAT_BC6H_UFLOAT_BLOCK                              Format = 143
	FORMAT_BC6H_SFLOAT_BLOCK                              Format = 144
	FORMAT_BC7_UNORM_BLOCK                                Format = 145
	FORMAT_BC7_SRGB_BLOCK                                 Format = 146
	FORMAT_ETC2_R8G8B8_UNORM_BLOCK                        Format = 147
	FORMAT_ETC2_R8G8B8_SRGB_BLOCK                         Format = 148
	FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK                      Format = 149
	FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK                       Format = 150
	FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK                      Format = 151
	FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK                       Format = 152
	FORMAT_EAC_R11_UNORM_BLOCK                            Format = 153
	FORMAT_EAC_R11_SNORM_BLOCK                            Format = 154
	FORMAT_EAC_R11G11_UNORM_BLOCK                         Format = 155
	FORMAT_EAC_R11G11_SNORM_BLOCK                         Format = 156
	FORMAT_ASTC_4x4_UNORM_BLOCK                           Format = 157
	FORMAT_ASTC_4x4_SRGB_BLOCK                            Format = 158
	FORMAT_ASTC_5x4_UNORM_BLOCK                           Format = 159
	FORMAT_ASTC_5x4_SRGB_BLOCK                            Format = 160
	FORMAT_ASTC_5x5_UNORM_BLOCK                           Format = 161
	FORMAT_ASTC_5x5_SRGB_BLOCK                            Format = 162
	FORMAT_ASTC_6x5_UNORM_BLOCK                           Format = 163
	FORMAT_ASTC_6x5_SRGB_BLOCK                            Format = 164
	FORMAT_ASTC_6x6_UNORM_BLOCK                           Format = 165
	FORMAT_ASTC_6x6_SRGB_BLOCK                            Format = 166
	FORMAT_ASTC_8x5_UNORM_BLOCK                           Format = 167
	FORMAT_ASTC_8x5_SRGB_BLOCK                            Format = 168
	FORMAT_ASTC_8x6_UNORM_BLOCK                           Format = 169
	FORMAT_ASTC_8x6_SRGB_BLOCK                            Format = 170
	FORMAT_ASTC_8x8_UNORM_BLOCK                           Format = 171
	FORMAT_ASTC_8x8_SRGB_BLOCK                            Format = 172
	FORMAT_ASTC_10x5_UNORM_BLOCK                          Format = 173
	FORMAT_ASTC_10x5_SRGB_BLOCK                           Format = 174
	FORMAT_ASTC_10x6_UNORM_BLOCK                          Format = 175
	FORMAT_ASTC_10x6_SRGB_BLOCK                           Format = 176
	FORMAT_ASTC_10x8_UNORM_BLOCK                          Format = 177
	FORMAT_ASTC_10x8_SRGB_BLOCK                           Format = 178
	FORMAT_ASTC_10x10_UNORM_BLOCK                         Format = 179
	FORMAT_ASTC_10x10_SRGB_BLOCK                          Format = 180
	FORMAT_ASTC_12x10_UNORM_BLOCK                         Format = 181
	FORMAT_ASTC_12x10_SRGB_BLOCK                          Format = 182
	FORMAT_ASTC_12x12_UNORM_BLOCK                         Format = 183
	FORMAT_ASTC_12x12_SRGB_BLOCK                          Format = 184
	FORMAT_G8B8G8R8_422_UNORM                             Format = 1000156000
	FORMAT_B8G8R8G8_422_UNORM                             Format = 1000156001
	FORMAT_G8_B8_R8_3PLANE_420_UNORM                      Format = 1000156002
	FORMAT_G8_B8R8_2PLANE_420_UNORM                       Format = 1000156003
	FORMAT_G8_B8_R8_3PLANE_422_UNORM                      Format = 1000156004
	FORMAT_G8_B8R8_2PLANE_422_UNORM                       Format = 1000156005
	FORMAT_G8_B8_R8_3PLANE_444_UNORM                      Format = 1000156006
	FORMAT_R10X6_UNORM_PACK16                             Format = 1000156007
	FORMAT_R10X6G10X6_UNORM_2PACK16                       Format = 1000156008
	FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16             Format = 1000156009
	FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16         Format = 1000156010
	FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16         Format = 1000156011
	FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16     Format = 1000156012
	FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16      Format = 1000156013
	FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16     Format = 1000156014
	FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16      Format = 1000156015
	FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16     Format = 1000156016
	FORMAT_R12X4_UNORM_PACK16                             Format = 1000156017
	FORMAT_R12X4G12X4_UNORM_2PACK16                       Format = 1000156018
	FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16             Format = 1000156019
	FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16         Format = 1000156020
	FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16         Format = 1000156021
	FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16     Format = 1000156022
	FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16      Format = 1000156023
	FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16     Format = 1000156024
	FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16      Format = 1000156025
	FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16     Format = 1000156026
	FORMAT_G16B16G16R16_422_UNORM                         Format = 1000156027
	FORMAT_B16G16R16G16_422_UNORM                         Format = 1000156028
	FORMAT_G16_B16_R16_3PLANE_420_UNORM                   Format = 1000156029
	FORMAT_G16_B16R16_2PLANE_420_UNORM                    Format = 1000156030
	FORMAT_G16_B16_R16_3PLANE_422_UNORM                   Format = 1000156031
	FORMAT_G16_B16R16_2PLANE_422_UNORM                    Format = 1000156032
	FORMAT_G16_B16_R16_3PLANE_444_UNORM                   Format = 1000156033
	FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG                    Format = 1000054000
	FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG                    Format = 1000054001
	FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG                    Format = 1000054002
	FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG                    Format = 1000054003
	FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG                     Format = 1000054004
	FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG                     Format = 1000054005
	FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG                     Format = 1000054006
	FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG                     Format = 1000054007
	FORMAT_G8B8G8R8_422_UNORM_KHR                         Format = FORMAT_G8B8G8R8_422_UNORM
	FORMAT_B8G8R8G8_422_UNORM_KHR                         Format = FORMAT_B8G8R8G8_422_UNORM
	FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_420_UNORM
	FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR                   Format = FORMAT_G8_B8R8_2PLANE_420_UNORM
	FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_422_UNORM
	FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR                   Format = FORMAT_G8_B8R8_2PLANE_422_UNORM
	FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_444_UNORM
	FORMAT_R10X6_UNORM_PACK16_KHR                         Format = FORMAT_R10X6_UNORM_PACK16
	FORMAT_R10X6G10X6_UNORM_2PACK16_KHR                   Format = FORMAT_R10X6G10X6_UNORM_2PACK16
	FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR         Format = FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR     Format = FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR     Format = FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR  Format = FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR  Format = FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	FORMAT_R12X4_UNORM_PACK16_KHR                         Format = FORMAT_R12X4_UNORM_PACK16
	FORMAT_R12X4G12X4_UNORM_2PACK16_KHR                   Format = FORMAT_R12X4G12X4_UNORM_2PACK16
	FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR         Format = FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR     Format = FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR     Format = FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR  Format = FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR  Format = FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	FORMAT_G16B16G16R16_422_UNORM_KHR                     Format = FORMAT_G16B16G16R16_422_UNORM
	FORMAT_B16G16R16G16_422_UNORM_KHR                     Format = FORMAT_B16G16R16G16_422_UNORM
	FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_420_UNORM
	FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR                Format = FORMAT_G16_B16R16_2PLANE_420_UNORM
	FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_422_UNORM
	FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR                Format = FORMAT_G16_B16R16_2PLANE_422_UNORM
	FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_444_UNORM
	FORMAT_BEGIN_RANGE                                    Format = FORMAT_UNDEFINED
	FORMAT_END_RANGE                                      Format = FORMAT_ASTC_12x12_SRGB_BLOCK
	FORMAT_RANGE_SIZE                                     Format = (FORMAT_ASTC_12x12_SRGB_BLOCK - FORMAT_UNDEFINED + 1)
	FORMAT_MAX_ENUM                                       Format = 2147483647
)

type ImageType int

const (
	IMAGE_TYPE_1D          ImageType = 0
	IMAGE_TYPE_2D          ImageType = 1
	IMAGE_TYPE_3D          ImageType = 2
	IMAGE_TYPE_BEGIN_RANGE ImageType = IMAGE_TYPE_1D
	IMAGE_TYPE_END_RANGE   ImageType = IMAGE_TYPE_3D
	IMAGE_TYPE_RANGE_SIZE  ImageType = (IMAGE_TYPE_3D - IMAGE_TYPE_1D + 1)
	IMAGE_TYPE_MAX_ENUM    ImageType = 2147483647
)

type ImageTiling int

const (
	IMAGE_TILING_OPTIMAL     ImageTiling = 0
	IMAGE_TILING_LINEAR      ImageTiling = 1
	IMAGE_TILING_BEGIN_RANGE ImageTiling = IMAGE_TILING_OPTIMAL
	IMAGE_TILING_END_RANGE   ImageTiling = IMAGE_TILING_LINEAR
	IMAGE_TILING_RANGE_SIZE  ImageTiling = (IMAGE_TILING_LINEAR - IMAGE_TILING_OPTIMAL + 1)
	IMAGE_TILING_MAX_ENUM    ImageTiling = 2147483647
)

type PhysicalDeviceType int

const (
	PHYSICAL_DEVICE_TYPE_OTHER          PhysicalDeviceType = 0
	PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU PhysicalDeviceType = 1
	PHYSICAL_DEVICE_TYPE_DISCRETE_GPU   PhysicalDeviceType = 2
	PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU    PhysicalDeviceType = 3
	PHYSICAL_DEVICE_TYPE_CPU            PhysicalDeviceType = 4
	PHYSICAL_DEVICE_TYPE_BEGIN_RANGE    PhysicalDeviceType = PHYSICAL_DEVICE_TYPE_OTHER
	PHYSICAL_DEVICE_TYPE_END_RANGE      PhysicalDeviceType = PHYSICAL_DEVICE_TYPE_CPU
	PHYSICAL_DEVICE_TYPE_RANGE_SIZE     PhysicalDeviceType = (PHYSICAL_DEVICE_TYPE_CPU - PHYSICAL_DEVICE_TYPE_OTHER + 1)
	PHYSICAL_DEVICE_TYPE_MAX_ENUM       PhysicalDeviceType = 2147483647
)

type QueryType int

const (
	QUERY_TYPE_OCCLUSION           QueryType = 0
	QUERY_TYPE_PIPELINE_STATISTICS QueryType = 1
	QUERY_TYPE_TIMESTAMP           QueryType = 2
	QUERY_TYPE_BEGIN_RANGE         QueryType = QUERY_TYPE_OCCLUSION
	QUERY_TYPE_END_RANGE           QueryType = QUERY_TYPE_TIMESTAMP
	QUERY_TYPE_RANGE_SIZE          QueryType = (QUERY_TYPE_TIMESTAMP - QUERY_TYPE_OCCLUSION + 1)
	QUERY_TYPE_MAX_ENUM            QueryType = 2147483647
)

type SharingMode int

const (
	SHARING_MODE_EXCLUSIVE   SharingMode = 0
	SHARING_MODE_CONCURRENT  SharingMode = 1
	SHARING_MODE_BEGIN_RANGE SharingMode = SHARING_MODE_EXCLUSIVE
	SHARING_MODE_END_RANGE   SharingMode = SHARING_MODE_CONCURRENT
	SHARING_MODE_RANGE_SIZE  SharingMode = (SHARING_MODE_CONCURRENT - SHARING_MODE_EXCLUSIVE + 1)
	SHARING_MODE_MAX_ENUM    SharingMode = 2147483647
)

type ImageLayout int

const (
	IMAGE_LAYOUT_UNDEFINED                                      ImageLayout = 0
	IMAGE_LAYOUT_GENERAL                                        ImageLayout = 1
	IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL                       ImageLayout = 2
	IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL               ImageLayout = 3
	IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL                ImageLayout = 4
	IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL                       ImageLayout = 5
	IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL                           ImageLayout = 6
	IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL                           ImageLayout = 7
	IMAGE_LAYOUT_PREINITIALIZED                                 ImageLayout = 8
	IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL     ImageLayout = 1000117000
	IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL     ImageLayout = 1000117001
	IMAGE_LAYOUT_PRESENT_SRC_KHR                                ImageLayout = 1000001002
	IMAGE_LAYOUT_SHARED_PRESENT_KHR                             ImageLayout = 1000111000
	IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR ImageLayout = IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
	IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR ImageLayout = IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	IMAGE_LAYOUT_BEGIN_RANGE                                    ImageLayout = IMAGE_LAYOUT_UNDEFINED
	IMAGE_LAYOUT_END_RANGE                                      ImageLayout = IMAGE_LAYOUT_PREINITIALIZED
	IMAGE_LAYOUT_RANGE_SIZE                                     ImageLayout = (IMAGE_LAYOUT_PREINITIALIZED - IMAGE_LAYOUT_UNDEFINED + 1)
	IMAGE_LAYOUT_MAX_ENUM                                       ImageLayout = 2147483647
)

type ImageViewType int

const (
	IMAGE_VIEW_TYPE_1D          ImageViewType = 0
	IMAGE_VIEW_TYPE_2D          ImageViewType = 1
	IMAGE_VIEW_TYPE_3D          ImageViewType = 2
	IMAGE_VIEW_TYPE_CUBE        ImageViewType = 3
	IMAGE_VIEW_TYPE_1D_ARRAY    ImageViewType = 4
	IMAGE_VIEW_TYPE_2D_ARRAY    ImageViewType = 5
	IMAGE_VIEW_TYPE_CUBE_ARRAY  ImageViewType = 6
	IMAGE_VIEW_TYPE_BEGIN_RANGE ImageViewType = IMAGE_VIEW_TYPE_1D
	IMAGE_VIEW_TYPE_END_RANGE   ImageViewType = IMAGE_VIEW_TYPE_CUBE_ARRAY
	IMAGE_VIEW_TYPE_RANGE_SIZE  ImageViewType = (IMAGE_VIEW_TYPE_CUBE_ARRAY - IMAGE_VIEW_TYPE_1D + 1)
	IMAGE_VIEW_TYPE_MAX_ENUM    ImageViewType = 2147483647
)

type ComponentSwizzle int

const (
	COMPONENT_SWIZZLE_IDENTITY    ComponentSwizzle = 0
	COMPONENT_SWIZZLE_ZERO        ComponentSwizzle = 1
	COMPONENT_SWIZZLE_ONE         ComponentSwizzle = 2
	COMPONENT_SWIZZLE_R           ComponentSwizzle = 3
	COMPONENT_SWIZZLE_G           ComponentSwizzle = 4
	COMPONENT_SWIZZLE_B           ComponentSwizzle = 5
	COMPONENT_SWIZZLE_A           ComponentSwizzle = 6
	COMPONENT_SWIZZLE_BEGIN_RANGE ComponentSwizzle = COMPONENT_SWIZZLE_IDENTITY
	COMPONENT_SWIZZLE_END_RANGE   ComponentSwizzle = COMPONENT_SWIZZLE_A
	COMPONENT_SWIZZLE_RANGE_SIZE  ComponentSwizzle = (COMPONENT_SWIZZLE_A - COMPONENT_SWIZZLE_IDENTITY + 1)
	COMPONENT_SWIZZLE_MAX_ENUM    ComponentSwizzle = 2147483647
)

type VertexInputRate int

const (
	VERTEX_INPUT_RATE_VERTEX      VertexInputRate = 0
	VERTEX_INPUT_RATE_INSTANCE    VertexInputRate = 1
	VERTEX_INPUT_RATE_BEGIN_RANGE VertexInputRate = VERTEX_INPUT_RATE_VERTEX
	VERTEX_INPUT_RATE_END_RANGE   VertexInputRate = VERTEX_INPUT_RATE_INSTANCE
	VERTEX_INPUT_RATE_RANGE_SIZE  VertexInputRate = (VERTEX_INPUT_RATE_INSTANCE - VERTEX_INPUT_RATE_VERTEX + 1)
	VERTEX_INPUT_RATE_MAX_ENUM    VertexInputRate = 2147483647
)

type PrimitiveTopology int

const (
	PRIMITIVE_TOPOLOGY_POINT_LIST                    PrimitiveTopology = 0
	PRIMITIVE_TOPOLOGY_LINE_LIST                     PrimitiveTopology = 1
	PRIMITIVE_TOPOLOGY_LINE_STRIP                    PrimitiveTopology = 2
	PRIMITIVE_TOPOLOGY_TRIANGLE_LIST                 PrimitiveTopology = 3
	PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP                PrimitiveTopology = 4
	PRIMITIVE_TOPOLOGY_TRIANGLE_FAN                  PrimitiveTopology = 5
	PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY      PrimitiveTopology = 6
	PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY     PrimitiveTopology = 7
	PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY  PrimitiveTopology = 8
	PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY PrimitiveTopology = 9
	PRIMITIVE_TOPOLOGY_PATCH_LIST                    PrimitiveTopology = 10
	PRIMITIVE_TOPOLOGY_BEGIN_RANGE                   PrimitiveTopology = PRIMITIVE_TOPOLOGY_POINT_LIST
	PRIMITIVE_TOPOLOGY_END_RANGE                     PrimitiveTopology = PRIMITIVE_TOPOLOGY_PATCH_LIST
	PRIMITIVE_TOPOLOGY_RANGE_SIZE                    PrimitiveTopology = (PRIMITIVE_TOPOLOGY_PATCH_LIST - PRIMITIVE_TOPOLOGY_POINT_LIST + 1)
	PRIMITIVE_TOPOLOGY_MAX_ENUM                      PrimitiveTopology = 2147483647
)

type PolygonMode int

const (
	POLYGON_MODE_FILL              PolygonMode = 0
	POLYGON_MODE_LINE              PolygonMode = 1
	POLYGON_MODE_POINT             PolygonMode = 2
	POLYGON_MODE_FILL_RECTANGLE_NV PolygonMode = 1000153000
	POLYGON_MODE_BEGIN_RANGE       PolygonMode = POLYGON_MODE_FILL
	POLYGON_MODE_END_RANGE         PolygonMode = POLYGON_MODE_POINT
	POLYGON_MODE_RANGE_SIZE        PolygonMode = (POLYGON_MODE_POINT - POLYGON_MODE_FILL + 1)
	POLYGON_MODE_MAX_ENUM          PolygonMode = 2147483647
)

type FrontFace int

const (
	FRONT_FACE_COUNTER_CLOCKWISE FrontFace = 0
	FRONT_FACE_CLOCKWISE         FrontFace = 1
	FRONT_FACE_BEGIN_RANGE       FrontFace = FRONT_FACE_COUNTER_CLOCKWISE
	FRONT_FACE_END_RANGE         FrontFace = FRONT_FACE_CLOCKWISE
	FRONT_FACE_RANGE_SIZE        FrontFace = (FRONT_FACE_CLOCKWISE - FRONT_FACE_COUNTER_CLOCKWISE + 1)
	FRONT_FACE_MAX_ENUM          FrontFace = 2147483647
)

type CompareOp int

const (
	COMPARE_OP_NEVER            CompareOp = 0
	COMPARE_OP_LESS             CompareOp = 1
	COMPARE_OP_EQUAL            CompareOp = 2
	COMPARE_OP_LESS_OR_EQUAL    CompareOp = 3
	COMPARE_OP_GREATER          CompareOp = 4
	COMPARE_OP_NOT_EQUAL        CompareOp = 5
	COMPARE_OP_GREATER_OR_EQUAL CompareOp = 6
	COMPARE_OP_ALWAYS           CompareOp = 7
	COMPARE_OP_BEGIN_RANGE      CompareOp = COMPARE_OP_NEVER
	COMPARE_OP_END_RANGE        CompareOp = COMPARE_OP_ALWAYS
	COMPARE_OP_RANGE_SIZE       CompareOp = (COMPARE_OP_ALWAYS - COMPARE_OP_NEVER + 1)
	COMPARE_OP_MAX_ENUM         CompareOp = 2147483647
)

type StencilOp int

const (
	STENCIL_OP_KEEP                StencilOp = 0
	STENCIL_OP_ZERO                StencilOp = 1
	STENCIL_OP_REPLACE             StencilOp = 2
	STENCIL_OP_INCREMENT_AND_CLAMP StencilOp = 3
	STENCIL_OP_DECREMENT_AND_CLAMP StencilOp = 4
	STENCIL_OP_INVERT              StencilOp = 5
	STENCIL_OP_INCREMENT_AND_WRAP  StencilOp = 6
	STENCIL_OP_DECREMENT_AND_WRAP  StencilOp = 7
	STENCIL_OP_BEGIN_RANGE         StencilOp = STENCIL_OP_KEEP
	STENCIL_OP_END_RANGE           StencilOp = STENCIL_OP_DECREMENT_AND_WRAP
	STENCIL_OP_RANGE_SIZE          StencilOp = (STENCIL_OP_DECREMENT_AND_WRAP - STENCIL_OP_KEEP + 1)
	STENCIL_OP_MAX_ENUM            StencilOp = 2147483647
)

type LogicOp int

const (
	LOGIC_OP_CLEAR         LogicOp = 0
	LOGIC_OP_AND           LogicOp = 1
	LOGIC_OP_AND_REVERSE   LogicOp = 2
	LOGIC_OP_COPY          LogicOp = 3
	LOGIC_OP_AND_INVERTED  LogicOp = 4
	LOGIC_OP_NO_OP         LogicOp = 5
	LOGIC_OP_XOR           LogicOp = 6
	LOGIC_OP_OR            LogicOp = 7
	LOGIC_OP_NOR           LogicOp = 8
	LOGIC_OP_EQUIVALENT    LogicOp = 9
	LOGIC_OP_INVERT        LogicOp = 10
	LOGIC_OP_OR_REVERSE    LogicOp = 11
	LOGIC_OP_COPY_INVERTED LogicOp = 12
	LOGIC_OP_OR_INVERTED   LogicOp = 13
	LOGIC_OP_NAND          LogicOp = 14
	LOGIC_OP_SET           LogicOp = 15
	LOGIC_OP_BEGIN_RANGE   LogicOp = LOGIC_OP_CLEAR
	LOGIC_OP_END_RANGE     LogicOp = LOGIC_OP_SET
	LOGIC_OP_RANGE_SIZE    LogicOp = (LOGIC_OP_SET - LOGIC_OP_CLEAR + 1)
	LOGIC_OP_MAX_ENUM      LogicOp = 2147483647
)

type BlendFactor int

const (
	BLEND_FACTOR_ZERO                     BlendFactor = 0
	BLEND_FACTOR_ONE                      BlendFactor = 1
	BLEND_FACTOR_SRC_COLOR                BlendFactor = 2
	BLEND_FACTOR_ONE_MINUS_SRC_COLOR      BlendFactor = 3
	BLEND_FACTOR_DST_COLOR                BlendFactor = 4
	BLEND_FACTOR_ONE_MINUS_DST_COLOR      BlendFactor = 5
	BLEND_FACTOR_SRC_ALPHA                BlendFactor = 6
	BLEND_FACTOR_ONE_MINUS_SRC_ALPHA      BlendFactor = 7
	BLEND_FACTOR_DST_ALPHA                BlendFactor = 8
	BLEND_FACTOR_ONE_MINUS_DST_ALPHA      BlendFactor = 9
	BLEND_FACTOR_CONSTANT_COLOR           BlendFactor = 10
	BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR BlendFactor = 11
	BLEND_FACTOR_CONSTANT_ALPHA           BlendFactor = 12
	BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA BlendFactor = 13
	BLEND_FACTOR_SRC_ALPHA_SATURATE       BlendFactor = 14
	BLEND_FACTOR_SRC1_COLOR               BlendFactor = 15
	BLEND_FACTOR_ONE_MINUS_SRC1_COLOR     BlendFactor = 16
	BLEND_FACTOR_SRC1_ALPHA               BlendFactor = 17
	BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA     BlendFactor = 18
	BLEND_FACTOR_BEGIN_RANGE              BlendFactor = BLEND_FACTOR_ZERO
	BLEND_FACTOR_END_RANGE                BlendFactor = BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
	BLEND_FACTOR_RANGE_SIZE               BlendFactor = (BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA - BLEND_FACTOR_ZERO + 1)
	BLEND_FACTOR_MAX_ENUM                 BlendFactor = 2147483647
)

type BlendOp int

const (
	BLEND_OP_ADD                    BlendOp = 0
	BLEND_OP_SUBTRACT               BlendOp = 1
	BLEND_OP_REVERSE_SUBTRACT       BlendOp = 2
	BLEND_OP_MIN                    BlendOp = 3
	BLEND_OP_MAX                    BlendOp = 4
	BLEND_OP_ZERO_EXT               BlendOp = 1000148000
	BLEND_OP_SRC_EXT                BlendOp = 1000148001
	BLEND_OP_DST_EXT                BlendOp = 1000148002
	BLEND_OP_SRC_OVER_EXT           BlendOp = 1000148003
	BLEND_OP_DST_OVER_EXT           BlendOp = 1000148004
	BLEND_OP_SRC_IN_EXT             BlendOp = 1000148005
	BLEND_OP_DST_IN_EXT             BlendOp = 1000148006
	BLEND_OP_SRC_OUT_EXT            BlendOp = 1000148007
	BLEND_OP_DST_OUT_EXT            BlendOp = 1000148008
	BLEND_OP_SRC_ATOP_EXT           BlendOp = 1000148009
	BLEND_OP_DST_ATOP_EXT           BlendOp = 1000148010
	BLEND_OP_XOR_EXT                BlendOp = 1000148011
	BLEND_OP_MULTIPLY_EXT           BlendOp = 1000148012
	BLEND_OP_SCREEN_EXT             BlendOp = 1000148013
	BLEND_OP_OVERLAY_EXT            BlendOp = 1000148014
	BLEND_OP_DARKEN_EXT             BlendOp = 1000148015
	BLEND_OP_LIGHTEN_EXT            BlendOp = 1000148016
	BLEND_OP_COLORDODGE_EXT         BlendOp = 1000148017
	BLEND_OP_COLORBURN_EXT          BlendOp = 1000148018
	BLEND_OP_HARDLIGHT_EXT          BlendOp = 1000148019
	BLEND_OP_SOFTLIGHT_EXT          BlendOp = 1000148020
	BLEND_OP_DIFFERENCE_EXT         BlendOp = 1000148021
	BLEND_OP_EXCLUSION_EXT          BlendOp = 1000148022
	BLEND_OP_INVERT_EXT             BlendOp = 1000148023
	BLEND_OP_INVERT_RGB_EXT         BlendOp = 1000148024
	BLEND_OP_LINEARDODGE_EXT        BlendOp = 1000148025
	BLEND_OP_LINEARBURN_EXT         BlendOp = 1000148026
	BLEND_OP_VIVIDLIGHT_EXT         BlendOp = 1000148027
	BLEND_OP_LINEARLIGHT_EXT        BlendOp = 1000148028
	BLEND_OP_PINLIGHT_EXT           BlendOp = 1000148029
	BLEND_OP_HARDMIX_EXT            BlendOp = 1000148030
	BLEND_OP_HSL_HUE_EXT            BlendOp = 1000148031
	BLEND_OP_HSL_SATURATION_EXT     BlendOp = 1000148032
	BLEND_OP_HSL_COLOR_EXT          BlendOp = 1000148033
	BLEND_OP_HSL_LUMINOSITY_EXT     BlendOp = 1000148034
	BLEND_OP_PLUS_EXT               BlendOp = 1000148035
	BLEND_OP_PLUS_CLAMPED_EXT       BlendOp = 1000148036
	BLEND_OP_PLUS_CLAMPED_ALPHA_EXT BlendOp = 1000148037
	BLEND_OP_PLUS_DARKER_EXT        BlendOp = 1000148038
	BLEND_OP_MINUS_EXT              BlendOp = 1000148039
	BLEND_OP_MINUS_CLAMPED_EXT      BlendOp = 1000148040
	BLEND_OP_CONTRAST_EXT           BlendOp = 1000148041
	BLEND_OP_INVERT_OVG_EXT         BlendOp = 1000148042
	BLEND_OP_RED_EXT                BlendOp = 1000148043
	BLEND_OP_GREEN_EXT              BlendOp = 1000148044
	BLEND_OP_BLUE_EXT               BlendOp = 1000148045
	BLEND_OP_BEGIN_RANGE            BlendOp = BLEND_OP_ADD
	BLEND_OP_END_RANGE              BlendOp = BLEND_OP_MAX
	BLEND_OP_RANGE_SIZE             BlendOp = (BLEND_OP_MAX - BLEND_OP_ADD + 1)
	BLEND_OP_MAX_ENUM               BlendOp = 2147483647
)

type DynamicState int

const (
	DYNAMIC_STATE_VIEWPORT              DynamicState = 0
	DYNAMIC_STATE_SCISSOR               DynamicState = 1
	DYNAMIC_STATE_LINE_WIDTH            DynamicState = 2
	DYNAMIC_STATE_DEPTH_BIAS            DynamicState = 3
	DYNAMIC_STATE_BLEND_CONSTANTS       DynamicState = 4
	DYNAMIC_STATE_DEPTH_BOUNDS          DynamicState = 5
	DYNAMIC_STATE_STENCIL_COMPARE_MASK  DynamicState = 6
	DYNAMIC_STATE_STENCIL_WRITE_MASK    DynamicState = 7
	DYNAMIC_STATE_STENCIL_REFERENCE     DynamicState = 8
	DYNAMIC_STATE_VIEWPORT_W_SCALING_NV DynamicState = 1000087000
	DYNAMIC_STATE_DISCARD_RECTANGLE_EXT DynamicState = 1000099000
	DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT  DynamicState = 1000143000
	DYNAMIC_STATE_BEGIN_RANGE           DynamicState = DYNAMIC_STATE_VIEWPORT
	DYNAMIC_STATE_END_RANGE             DynamicState = DYNAMIC_STATE_STENCIL_REFERENCE
	DYNAMIC_STATE_RANGE_SIZE            DynamicState = (DYNAMIC_STATE_STENCIL_REFERENCE - DYNAMIC_STATE_VIEWPORT + 1)
	DYNAMIC_STATE_MAX_ENUM              DynamicState = 2147483647
)

type Filter int

const (
	FILTER_NEAREST     Filter = 0
	FILTER_LINEAR      Filter = 1
	FILTER_CUBIC_IMG   Filter = 1000015000
	FILTER_BEGIN_RANGE Filter = FILTER_NEAREST
	FILTER_END_RANGE   Filter = FILTER_LINEAR
	FILTER_RANGE_SIZE  Filter = (FILTER_LINEAR - FILTER_NEAREST + 1)
	FILTER_MAX_ENUM    Filter = 2147483647
)

type SamplerMipmapMode int

const (
	SAMPLER_MIPMAP_MODE_NEAREST     SamplerMipmapMode = 0
	SAMPLER_MIPMAP_MODE_LINEAR      SamplerMipmapMode = 1
	SAMPLER_MIPMAP_MODE_BEGIN_RANGE SamplerMipmapMode = SAMPLER_MIPMAP_MODE_NEAREST
	SAMPLER_MIPMAP_MODE_END_RANGE   SamplerMipmapMode = SAMPLER_MIPMAP_MODE_LINEAR
	SAMPLER_MIPMAP_MODE_RANGE_SIZE  SamplerMipmapMode = (SAMPLER_MIPMAP_MODE_LINEAR - SAMPLER_MIPMAP_MODE_NEAREST + 1)
	SAMPLER_MIPMAP_MODE_MAX_ENUM    SamplerMipmapMode = 2147483647
)

type SamplerAddressMode int

const (
	SAMPLER_ADDRESS_MODE_REPEAT               SamplerAddressMode = 0
	SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT      SamplerAddressMode = 1
	SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE        SamplerAddressMode = 2
	SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER      SamplerAddressMode = 3
	SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE SamplerAddressMode = 4
	SAMPLER_ADDRESS_MODE_BEGIN_RANGE          SamplerAddressMode = SAMPLER_ADDRESS_MODE_REPEAT
	SAMPLER_ADDRESS_MODE_END_RANGE            SamplerAddressMode = SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
	SAMPLER_ADDRESS_MODE_RANGE_SIZE           SamplerAddressMode = (SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER - SAMPLER_ADDRESS_MODE_REPEAT + 1)
	SAMPLER_ADDRESS_MODE_MAX_ENUM             SamplerAddressMode = 2147483647
)

type BorderColor int

const (
	BORDER_COLOR_FLOAT_TRANSPARENT_BLACK BorderColor = 0
	BORDER_COLOR_INT_TRANSPARENT_BLACK   BorderColor = 1
	BORDER_COLOR_FLOAT_OPAQUE_BLACK      BorderColor = 2
	BORDER_COLOR_INT_OPAQUE_BLACK        BorderColor = 3
	BORDER_COLOR_FLOAT_OPAQUE_WHITE      BorderColor = 4
	BORDER_COLOR_INT_OPAQUE_WHITE        BorderColor = 5
	BORDER_COLOR_BEGIN_RANGE             BorderColor = BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	BORDER_COLOR_END_RANGE               BorderColor = BORDER_COLOR_INT_OPAQUE_WHITE
	BORDER_COLOR_RANGE_SIZE              BorderColor = (BORDER_COLOR_INT_OPAQUE_WHITE - BORDER_COLOR_FLOAT_TRANSPARENT_BLACK + 1)
	BORDER_COLOR_MAX_ENUM                BorderColor = 2147483647
)

type DescriptorType int

const (
	DESCRIPTOR_TYPE_SAMPLER                DescriptorType = 0
	DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER DescriptorType = 1
	DESCRIPTOR_TYPE_SAMPLED_IMAGE          DescriptorType = 2
	DESCRIPTOR_TYPE_STORAGE_IMAGE          DescriptorType = 3
	DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER   DescriptorType = 4
	DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER   DescriptorType = 5
	DESCRIPTOR_TYPE_UNIFORM_BUFFER         DescriptorType = 6
	DESCRIPTOR_TYPE_STORAGE_BUFFER         DescriptorType = 7
	DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC DescriptorType = 8
	DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC DescriptorType = 9
	DESCRIPTOR_TYPE_INPUT_ATTACHMENT       DescriptorType = 10
	DESCRIPTOR_TYPE_BEGIN_RANGE            DescriptorType = DESCRIPTOR_TYPE_SAMPLER
	DESCRIPTOR_TYPE_END_RANGE              DescriptorType = DESCRIPTOR_TYPE_INPUT_ATTACHMENT
	DESCRIPTOR_TYPE_RANGE_SIZE             DescriptorType = (DESCRIPTOR_TYPE_INPUT_ATTACHMENT - DESCRIPTOR_TYPE_SAMPLER + 1)
	DESCRIPTOR_TYPE_MAX_ENUM               DescriptorType = 2147483647
)

type AttachmentLoadOp int

const (
	ATTACHMENT_LOAD_OP_LOAD        AttachmentLoadOp = 0
	ATTACHMENT_LOAD_OP_CLEAR       AttachmentLoadOp = 1
	ATTACHMENT_LOAD_OP_DONT_CARE   AttachmentLoadOp = 2
	ATTACHMENT_LOAD_OP_BEGIN_RANGE AttachmentLoadOp = ATTACHMENT_LOAD_OP_LOAD
	ATTACHMENT_LOAD_OP_END_RANGE   AttachmentLoadOp = ATTACHMENT_LOAD_OP_DONT_CARE
	ATTACHMENT_LOAD_OP_RANGE_SIZE  AttachmentLoadOp = (ATTACHMENT_LOAD_OP_DONT_CARE - ATTACHMENT_LOAD_OP_LOAD + 1)
	ATTACHMENT_LOAD_OP_MAX_ENUM    AttachmentLoadOp = 2147483647
)

type AttachmentStoreOp int

const (
	ATTACHMENT_STORE_OP_STORE       AttachmentStoreOp = 0
	ATTACHMENT_STORE_OP_DONT_CARE   AttachmentStoreOp = 1
	ATTACHMENT_STORE_OP_BEGIN_RANGE AttachmentStoreOp = ATTACHMENT_STORE_OP_STORE
	ATTACHMENT_STORE_OP_END_RANGE   AttachmentStoreOp = ATTACHMENT_STORE_OP_DONT_CARE
	ATTACHMENT_STORE_OP_RANGE_SIZE  AttachmentStoreOp = (ATTACHMENT_STORE_OP_DONT_CARE - ATTACHMENT_STORE_OP_STORE + 1)
	ATTACHMENT_STORE_OP_MAX_ENUM    AttachmentStoreOp = 2147483647
)

type PipelineBindPoint int

const (
	PIPELINE_BIND_POINT_GRAPHICS    PipelineBindPoint = 0
	PIPELINE_BIND_POINT_COMPUTE     PipelineBindPoint = 1
	PIPELINE_BIND_POINT_BEGIN_RANGE PipelineBindPoint = PIPELINE_BIND_POINT_GRAPHICS
	PIPELINE_BIND_POINT_END_RANGE   PipelineBindPoint = PIPELINE_BIND_POINT_COMPUTE
	PIPELINE_BIND_POINT_RANGE_SIZE  PipelineBindPoint = (PIPELINE_BIND_POINT_COMPUTE - PIPELINE_BIND_POINT_GRAPHICS + 1)
	PIPELINE_BIND_POINT_MAX_ENUM    PipelineBindPoint = 2147483647
)

type CommandBufferLevel int

const (
	COMMAND_BUFFER_LEVEL_PRIMARY     CommandBufferLevel = 0
	COMMAND_BUFFER_LEVEL_SECONDARY   CommandBufferLevel = 1
	COMMAND_BUFFER_LEVEL_BEGIN_RANGE CommandBufferLevel = COMMAND_BUFFER_LEVEL_PRIMARY
	COMMAND_BUFFER_LEVEL_END_RANGE   CommandBufferLevel = COMMAND_BUFFER_LEVEL_SECONDARY
	COMMAND_BUFFER_LEVEL_RANGE_SIZE  CommandBufferLevel = (COMMAND_BUFFER_LEVEL_SECONDARY - COMMAND_BUFFER_LEVEL_PRIMARY + 1)
	COMMAND_BUFFER_LEVEL_MAX_ENUM    CommandBufferLevel = 2147483647
)

type IndexType int

const (
	INDEX_TYPE_UINT16      IndexType = 0
	INDEX_TYPE_UINT32      IndexType = 1
	INDEX_TYPE_BEGIN_RANGE IndexType = INDEX_TYPE_UINT16
	INDEX_TYPE_END_RANGE   IndexType = INDEX_TYPE_UINT32
	INDEX_TYPE_RANGE_SIZE  IndexType = (INDEX_TYPE_UINT32 - INDEX_TYPE_UINT16 + 1)
	INDEX_TYPE_MAX_ENUM    IndexType = 2147483647
)

type SubpassContents int

const (
	SUBPASS_CONTENTS_INLINE                    SubpassContents = 0
	SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS SubpassContents = 1
	SUBPASS_CONTENTS_BEGIN_RANGE               SubpassContents = SUBPASS_CONTENTS_INLINE
	SUBPASS_CONTENTS_END_RANGE                 SubpassContents = SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	SUBPASS_CONTENTS_RANGE_SIZE                SubpassContents = (SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS - SUBPASS_CONTENTS_INLINE + 1)
	SUBPASS_CONTENTS_MAX_ENUM                  SubpassContents = 2147483647
)

type ObjectType int

const (
	OBJECT_TYPE_UNKNOWN                        ObjectType = 0
	OBJECT_TYPE_INSTANCE                       ObjectType = 1
	OBJECT_TYPE_PHYSICAL_DEVICE                ObjectType = 2
	OBJECT_TYPE_DEVICE                         ObjectType = 3
	OBJECT_TYPE_QUEUE                          ObjectType = 4
	OBJECT_TYPE_SEMAPHORE                      ObjectType = 5
	OBJECT_TYPE_COMMAND_BUFFER                 ObjectType = 6
	OBJECT_TYPE_FENCE                          ObjectType = 7
	OBJECT_TYPE_DEVICE_MEMORY                  ObjectType = 8
	OBJECT_TYPE_BUFFER                         ObjectType = 9
	OBJECT_TYPE_IMAGE                          ObjectType = 10
	OBJECT_TYPE_EVENT                          ObjectType = 11
	OBJECT_TYPE_QUERY_POOL                     ObjectType = 12
	OBJECT_TYPE_BUFFER_VIEW                    ObjectType = 13
	OBJECT_TYPE_IMAGE_VIEW                     ObjectType = 14
	OBJECT_TYPE_SHADER_MODULE                  ObjectType = 15
	OBJECT_TYPE_PIPELINE_CACHE                 ObjectType = 16
	OBJECT_TYPE_PIPELINE_LAYOUT                ObjectType = 17
	OBJECT_TYPE_RENDER_PASS                    ObjectType = 18
	OBJECT_TYPE_PIPELINE                       ObjectType = 19
	OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT          ObjectType = 20
	OBJECT_TYPE_SAMPLER                        ObjectType = 21
	OBJECT_TYPE_DESCRIPTOR_POOL                ObjectType = 22
	OBJECT_TYPE_DESCRIPTOR_SET                 ObjectType = 23
	OBJECT_TYPE_FRAMEBUFFER                    ObjectType = 24
	OBJECT_TYPE_COMMAND_POOL                   ObjectType = 25
	OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION       ObjectType = 1000156000
	OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE     ObjectType = 1000085000
	OBJECT_TYPE_SURFACE_KHR                    ObjectType = 1000000000
	OBJECT_TYPE_SWAPCHAIN_KHR                  ObjectType = 1000001000
	OBJECT_TYPE_DISPLAY_KHR                    ObjectType = 1000002000
	OBJECT_TYPE_DISPLAY_MODE_KHR               ObjectType = 1000002001
	OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT      ObjectType = 1000011000
	OBJECT_TYPE_OBJECT_TABLE_NVX               ObjectType = 1000086000
	OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX   ObjectType = 1000086001
	OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT      ObjectType = 1000128000
	OBJECT_TYPE_VALIDATION_CACHE_EXT           ObjectType = 1000160000
	OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR ObjectType = OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR   ObjectType = OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
	OBJECT_TYPE_BEGIN_RANGE                    ObjectType = OBJECT_TYPE_UNKNOWN
	OBJECT_TYPE_END_RANGE                      ObjectType = OBJECT_TYPE_COMMAND_POOL
	OBJECT_TYPE_RANGE_SIZE                     ObjectType = (OBJECT_TYPE_COMMAND_POOL - OBJECT_TYPE_UNKNOWN + 1)
	OBJECT_TYPE_MAX_ENUM                       ObjectType = 2147483647
)

type VendorId int

const (
	VENDOR_ID_VIV         VendorId = 65537
	VENDOR_ID_VSI         VendorId = 65538
	VENDOR_ID_KAZAN       VendorId = 65539
	VENDOR_ID_BEGIN_RANGE VendorId = VENDOR_ID_VIV
	VENDOR_ID_END_RANGE   VendorId = VENDOR_ID_KAZAN
	VENDOR_ID_RANGE_SIZE  VendorId = (VENDOR_ID_KAZAN - VENDOR_ID_VIV + 1)
	VENDOR_ID_MAX_ENUM    VendorId = 2147483647
)

type InstanceCreateFlags Flags
type FormatFeatureFlags Flags

const (
	FORMAT_FEATURE_SAMPLED_IMAGE_BIT                                                               FormatFeatureFlags = 1
	FORMAT_FEATURE_STORAGE_IMAGE_BIT                                                               FormatFeatureFlags = 2
	FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT                                                        FormatFeatureFlags = 4
	FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT                                                        FormatFeatureFlags = 8
	FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT                                                        FormatFeatureFlags = 16
	FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT                                                 FormatFeatureFlags = 32
	FORMAT_FEATURE_VERTEX_BUFFER_BIT                                                               FormatFeatureFlags = 64
	FORMAT_FEATURE_COLOR_ATTACHMENT_BIT                                                            FormatFeatureFlags = 128
	FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT                                                      FormatFeatureFlags = 256
	FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT                                                    FormatFeatureFlags = 512
	FORMAT_FEATURE_BLIT_SRC_BIT                                                                    FormatFeatureFlags = 1024
	FORMAT_FEATURE_BLIT_DST_BIT                                                                    FormatFeatureFlags = 2048
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT                                                 FormatFeatureFlags = 4096
	FORMAT_FEATURE_TRANSFER_SRC_BIT                                                                FormatFeatureFlags = 16384
	FORMAT_FEATURE_TRANSFER_DST_BIT                                                                FormatFeatureFlags = 32768
	FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT                                                     FormatFeatureFlags = 131072
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT                                FormatFeatureFlags = 262144
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT               FormatFeatureFlags = 524288
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT               FormatFeatureFlags = 1048576
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT     FormatFeatureFlags = 2097152
	FORMAT_FEATURE_DISJOINT_BIT                                                                    FormatFeatureFlags = 4194304
	FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT                                                      FormatFeatureFlags = 8388608
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG                                              FormatFeatureFlags = 8192
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT                                             FormatFeatureFlags = 65536
	FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR                                                            FormatFeatureFlags = FORMAT_FEATURE_TRANSFER_SRC_BIT
	FORMAT_FEATURE_TRANSFER_DST_BIT_KHR                                                            FormatFeatureFlags = FORMAT_FEATURE_TRANSFER_DST_BIT
	FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR                                                 FormatFeatureFlags = FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR                            FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR           FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR           FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	FORMAT_FEATURE_DISJOINT_BIT_KHR                                                                FormatFeatureFlags = FORMAT_FEATURE_DISJOINT_BIT
	FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR                                                  FormatFeatureFlags = FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	FORMAT_FEATURE_FLAG_BITS_MAX_ENUM                                                              FormatFeatureFlags = 2147483647
)

type ImageUsageFlags Flags

const (
	IMAGE_USAGE_TRANSFER_SRC_BIT             ImageUsageFlags = 1
	IMAGE_USAGE_TRANSFER_DST_BIT             ImageUsageFlags = 2
	IMAGE_USAGE_SAMPLED_BIT                  ImageUsageFlags = 4
	IMAGE_USAGE_STORAGE_BIT                  ImageUsageFlags = 8
	IMAGE_USAGE_COLOR_ATTACHMENT_BIT         ImageUsageFlags = 16
	IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT ImageUsageFlags = 32
	IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT     ImageUsageFlags = 64
	IMAGE_USAGE_INPUT_ATTACHMENT_BIT         ImageUsageFlags = 128
	IMAGE_USAGE_FLAG_BITS_MAX_ENUM           ImageUsageFlags = 2147483647
)

type ImageCreateFlags Flags

const (
	IMAGE_CREATE_SPARSE_BINDING_BIT                        ImageCreateFlags = 1
	IMAGE_CREATE_SPARSE_RESIDENCY_BIT                      ImageCreateFlags = 2
	IMAGE_CREATE_SPARSE_ALIASED_BIT                        ImageCreateFlags = 4
	IMAGE_CREATE_MUTABLE_FORMAT_BIT                        ImageCreateFlags = 8
	IMAGE_CREATE_CUBE_COMPATIBLE_BIT                       ImageCreateFlags = 16
	IMAGE_CREATE_ALIAS_BIT                                 ImageCreateFlags = 1024
	IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT           ImageCreateFlags = 64
	IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT                   ImageCreateFlags = 32
	IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT           ImageCreateFlags = 128
	IMAGE_CREATE_EXTENDED_USAGE_BIT                        ImageCreateFlags = 256
	IMAGE_CREATE_PROTECTED_BIT                             ImageCreateFlags = 2048
	IMAGE_CREATE_DISJOINT_BIT                              ImageCreateFlags = 512
	IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT ImageCreateFlags = 4096
	IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR       ImageCreateFlags = IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT
	IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR               ImageCreateFlags = IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR       ImageCreateFlags = IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR                    ImageCreateFlags = IMAGE_CREATE_EXTENDED_USAGE_BIT
	IMAGE_CREATE_DISJOINT_BIT_KHR                          ImageCreateFlags = IMAGE_CREATE_DISJOINT_BIT
	IMAGE_CREATE_ALIAS_BIT_KHR                             ImageCreateFlags = IMAGE_CREATE_ALIAS_BIT
	IMAGE_CREATE_FLAG_BITS_MAX_ENUM                        ImageCreateFlags = 2147483647
)

type SampleCountFlags Flags

const (
	SAMPLE_COUNT_1_BIT              SampleCountFlags = 1
	SAMPLE_COUNT_2_BIT              SampleCountFlags = 2
	SAMPLE_COUNT_4_BIT              SampleCountFlags = 4
	SAMPLE_COUNT_8_BIT              SampleCountFlags = 8
	SAMPLE_COUNT_16_BIT             SampleCountFlags = 16
	SAMPLE_COUNT_32_BIT             SampleCountFlags = 32
	SAMPLE_COUNT_64_BIT             SampleCountFlags = 64
	SAMPLE_COUNT_FLAG_BITS_MAX_ENUM SampleCountFlags = 2147483647
)

type QueueFlags Flags

const (
	QUEUE_GRAPHICS_BIT       QueueFlags = 1
	QUEUE_COMPUTE_BIT        QueueFlags = 2
	QUEUE_TRANSFER_BIT       QueueFlags = 4
	QUEUE_SPARSE_BINDING_BIT QueueFlags = 8
	QUEUE_PROTECTED_BIT      QueueFlags = 16
	QUEUE_FLAG_BITS_MAX_ENUM QueueFlags = 2147483647
)

type MemoryPropertyFlags Flags

const (
	MEMORY_PROPERTY_DEVICE_LOCAL_BIT     MemoryPropertyFlags = 1
	MEMORY_PROPERTY_HOST_VISIBLE_BIT     MemoryPropertyFlags = 2
	MEMORY_PROPERTY_HOST_COHERENT_BIT    MemoryPropertyFlags = 4
	MEMORY_PROPERTY_HOST_CACHED_BIT      MemoryPropertyFlags = 8
	MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT MemoryPropertyFlags = 16
	MEMORY_PROPERTY_PROTECTED_BIT        MemoryPropertyFlags = 32
	MEMORY_PROPERTY_FLAG_BITS_MAX_ENUM   MemoryPropertyFlags = 2147483647
)

type MemoryHeapFlags Flags

const (
	MEMORY_HEAP_DEVICE_LOCAL_BIT       MemoryHeapFlags = 1
	MEMORY_HEAP_MULTI_INSTANCE_BIT     MemoryHeapFlags = 2
	MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR MemoryHeapFlags = MEMORY_HEAP_MULTI_INSTANCE_BIT
	MEMORY_HEAP_FLAG_BITS_MAX_ENUM     MemoryHeapFlags = 2147483647
)

type DeviceCreateFlags Flags
type DeviceQueueCreateFlags Flags

const (
	DEVICE_QUEUE_CREATE_PROTECTED_BIT      DeviceQueueCreateFlags = 1
	DEVICE_QUEUE_CREATE_FLAG_BITS_MAX_ENUM DeviceQueueCreateFlags = 2147483647
)

type PipelineStageFlags Flags

const (
	PIPELINE_STAGE_TOP_OF_PIPE_BIT                    PipelineStageFlags = 1
	PIPELINE_STAGE_DRAW_INDIRECT_BIT                  PipelineStageFlags = 2
	PIPELINE_STAGE_VERTEX_INPUT_BIT                   PipelineStageFlags = 4
	PIPELINE_STAGE_VERTEX_SHADER_BIT                  PipelineStageFlags = 8
	PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT    PipelineStageFlags = 16
	PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT PipelineStageFlags = 32
	PIPELINE_STAGE_GEOMETRY_SHADER_BIT                PipelineStageFlags = 64
	PIPELINE_STAGE_FRAGMENT_SHADER_BIT                PipelineStageFlags = 128
	PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT           PipelineStageFlags = 256
	PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT            PipelineStageFlags = 512
	PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT        PipelineStageFlags = 1024
	PIPELINE_STAGE_COMPUTE_SHADER_BIT                 PipelineStageFlags = 2048
	PIPELINE_STAGE_TRANSFER_BIT                       PipelineStageFlags = 4096
	PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT                 PipelineStageFlags = 8192
	PIPELINE_STAGE_HOST_BIT                           PipelineStageFlags = 16384
	PIPELINE_STAGE_ALL_GRAPHICS_BIT                   PipelineStageFlags = 32768
	PIPELINE_STAGE_ALL_COMMANDS_BIT                   PipelineStageFlags = 65536
	PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT      PipelineStageFlags = 262144
	PIPELINE_STAGE_COMMAND_PROCESS_BIT_NVX            PipelineStageFlags = 131072
	PIPELINE_STAGE_FLAG_BITS_MAX_ENUM                 PipelineStageFlags = 2147483647
)

type MemoryMapFlags Flags
type ImageAspectFlags Flags

const (
	IMAGE_ASPECT_COLOR_BIT          ImageAspectFlags = 1
	IMAGE_ASPECT_DEPTH_BIT          ImageAspectFlags = 2
	IMAGE_ASPECT_STENCIL_BIT        ImageAspectFlags = 4
	IMAGE_ASPECT_METADATA_BIT       ImageAspectFlags = 8
	IMAGE_ASPECT_PLANE_0_BIT        ImageAspectFlags = 16
	IMAGE_ASPECT_PLANE_1_BIT        ImageAspectFlags = 32
	IMAGE_ASPECT_PLANE_2_BIT        ImageAspectFlags = 64
	IMAGE_ASPECT_PLANE_0_BIT_KHR    ImageAspectFlags = IMAGE_ASPECT_PLANE_0_BIT
	IMAGE_ASPECT_PLANE_1_BIT_KHR    ImageAspectFlags = IMAGE_ASPECT_PLANE_1_BIT
	IMAGE_ASPECT_PLANE_2_BIT_KHR    ImageAspectFlags = IMAGE_ASPECT_PLANE_2_BIT
	IMAGE_ASPECT_FLAG_BITS_MAX_ENUM ImageAspectFlags = 2147483647
)

type SparseImageFormatFlags Flags

const (
	SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT         SparseImageFormatFlags = 1
	SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT       SparseImageFormatFlags = 2
	SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT SparseImageFormatFlags = 4
	SPARSE_IMAGE_FORMAT_FLAG_BITS_MAX_ENUM         SparseImageFormatFlags = 2147483647
)

type SparseMemoryBindFlags Flags

const (
	SPARSE_MEMORY_BIND_METADATA_BIT       SparseMemoryBindFlags = 1
	SPARSE_MEMORY_BIND_FLAG_BITS_MAX_ENUM SparseMemoryBindFlags = 2147483647
)

type FenceCreateFlags Flags

const (
	FENCE_CREATE_SIGNALED_BIT       FenceCreateFlags = 1
	FENCE_CREATE_FLAG_BITS_MAX_ENUM FenceCreateFlags = 2147483647
)

type SemaphoreCreateFlags Flags
type EventCreateFlags Flags
type QueryPoolCreateFlags Flags
type QueryPipelineStatisticFlags Flags

const (
	QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT                    QueryPipelineStatisticFlags = 1
	QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT                  QueryPipelineStatisticFlags = 2
	QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT                  QueryPipelineStatisticFlags = 4
	QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT                QueryPipelineStatisticFlags = 8
	QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT                 QueryPipelineStatisticFlags = 16
	QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT                       QueryPipelineStatisticFlags = 32
	QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT                        QueryPipelineStatisticFlags = 64
	QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT                QueryPipelineStatisticFlags = 128
	QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT        QueryPipelineStatisticFlags = 256
	QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT QueryPipelineStatisticFlags = 512
	QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT                 QueryPipelineStatisticFlags = 1024
	QUERY_PIPELINE_STATISTIC_FLAG_BITS_MAX_ENUM                             QueryPipelineStatisticFlags = 2147483647
)

type QueryResultFlags Flags

const (
	QUERY_RESULT_64_BIT                QueryResultFlags = 1
	QUERY_RESULT_WAIT_BIT              QueryResultFlags = 2
	QUERY_RESULT_WITH_AVAILABILITY_BIT QueryResultFlags = 4
	QUERY_RESULT_PARTIAL_BIT           QueryResultFlags = 8
	QUERY_RESULT_FLAG_BITS_MAX_ENUM    QueryResultFlags = 2147483647
)

type BufferCreateFlags Flags

const (
	BUFFER_CREATE_SPARSE_BINDING_BIT   BufferCreateFlags = 1
	BUFFER_CREATE_SPARSE_RESIDENCY_BIT BufferCreateFlags = 2
	BUFFER_CREATE_SPARSE_ALIASED_BIT   BufferCreateFlags = 4
	BUFFER_CREATE_PROTECTED_BIT        BufferCreateFlags = 8
	BUFFER_CREATE_FLAG_BITS_MAX_ENUM   BufferCreateFlags = 2147483647
)

type BufferUsageFlags Flags

const (
	BUFFER_USAGE_TRANSFER_SRC_BIT              BufferUsageFlags = 1
	BUFFER_USAGE_TRANSFER_DST_BIT              BufferUsageFlags = 2
	BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT      BufferUsageFlags = 4
	BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT      BufferUsageFlags = 8
	BUFFER_USAGE_UNIFORM_BUFFER_BIT            BufferUsageFlags = 16
	BUFFER_USAGE_STORAGE_BUFFER_BIT            BufferUsageFlags = 32
	BUFFER_USAGE_INDEX_BUFFER_BIT              BufferUsageFlags = 64
	BUFFER_USAGE_VERTEX_BUFFER_BIT             BufferUsageFlags = 128
	BUFFER_USAGE_INDIRECT_BUFFER_BIT           BufferUsageFlags = 256
	BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT BufferUsageFlags = 512
	BUFFER_USAGE_FLAG_BITS_MAX_ENUM            BufferUsageFlags = 2147483647
)

type BufferViewCreateFlags Flags
type ImageViewCreateFlags Flags
type ShaderModuleCreateFlags Flags
type PipelineCacheCreateFlags Flags
type PipelineCreateFlags Flags

const (
	PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT             PipelineCreateFlags = 1
	PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT                PipelineCreateFlags = 2
	PIPELINE_CREATE_DERIVATIVE_BIT                       PipelineCreateFlags = 4
	PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT     PipelineCreateFlags = 8
	PIPELINE_CREATE_DISPATCH_BASE                        PipelineCreateFlags = 16
	PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR PipelineCreateFlags = PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
	PIPELINE_CREATE_DISPATCH_BASE_KHR                    PipelineCreateFlags = PIPELINE_CREATE_DISPATCH_BASE
	PIPELINE_CREATE_FLAG_BITS_MAX_ENUM                   PipelineCreateFlags = 2147483647
)

type PipelineShaderStageCreateFlags Flags
type ShaderStageFlags Flags

const (
	SHADER_STAGE_VERTEX_BIT                  ShaderStageFlags = 1
	SHADER_STAGE_TESSELLATION_CONTROL_BIT    ShaderStageFlags = 2
	SHADER_STAGE_TESSELLATION_EVALUATION_BIT ShaderStageFlags = 4
	SHADER_STAGE_GEOMETRY_BIT                ShaderStageFlags = 8
	SHADER_STAGE_FRAGMENT_BIT                ShaderStageFlags = 16
	SHADER_STAGE_COMPUTE_BIT                 ShaderStageFlags = 32
	SHADER_STAGE_ALL_GRAPHICS                ShaderStageFlags = 31
	SHADER_STAGE_ALL                         ShaderStageFlags = 2147483647
	SHADER_STAGE_FLAG_BITS_MAX_ENUM          ShaderStageFlags = 2147483647
)

type PipelineVertexInputStateCreateFlags Flags
type PipelineInputAssemblyStateCreateFlags Flags
type PipelineTessellationStateCreateFlags Flags
type PipelineViewportStateCreateFlags Flags
type PipelineRasterizationStateCreateFlags Flags
type CullModeFlags Flags

const (
	CULL_MODE_NONE               CullModeFlags = 0
	CULL_MODE_FRONT_BIT          CullModeFlags = 1
	CULL_MODE_BACK_BIT           CullModeFlags = 2
	CULL_MODE_FRONT_AND_BACK     CullModeFlags = 3
	CULL_MODE_FLAG_BITS_MAX_ENUM CullModeFlags = 2147483647
)

type PipelineMultisampleStateCreateFlags Flags
type PipelineDepthStencilStateCreateFlags Flags
type PipelineColorBlendStateCreateFlags Flags
type ColorComponentFlags Flags

const (
	COLOR_COMPONENT_R_BIT              ColorComponentFlags = 1
	COLOR_COMPONENT_G_BIT              ColorComponentFlags = 2
	COLOR_COMPONENT_B_BIT              ColorComponentFlags = 4
	COLOR_COMPONENT_A_BIT              ColorComponentFlags = 8
	COLOR_COMPONENT_FLAG_BITS_MAX_ENUM ColorComponentFlags = 2147483647
)

type PipelineDynamicStateCreateFlags Flags
type PipelineLayoutCreateFlags Flags
type SamplerCreateFlags Flags
type DescriptorSetLayoutCreateFlags Flags

const (
	DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR        DescriptorSetLayoutCreateFlags = 1
	DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT DescriptorSetLayoutCreateFlags = 2
	DESCRIPTOR_SET_LAYOUT_CREATE_FLAG_BITS_MAX_ENUM             DescriptorSetLayoutCreateFlags = 2147483647
)

type DescriptorPoolCreateFlags Flags

const (
	DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT   DescriptorPoolCreateFlags = 1
	DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT DescriptorPoolCreateFlags = 2
	DESCRIPTOR_POOL_CREATE_FLAG_BITS_MAX_ENUM        DescriptorPoolCreateFlags = 2147483647
)

type DescriptorPoolResetFlags Flags
type FramebufferCreateFlags Flags
type RenderPassCreateFlags Flags
type AttachmentDescriptionFlags Flags

const (
	ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT      AttachmentDescriptionFlags = 1
	ATTACHMENT_DESCRIPTION_FLAG_BITS_MAX_ENUM AttachmentDescriptionFlags = 2147483647
)

type SubpassDescriptionFlags Flags

const (
	SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX      SubpassDescriptionFlags = 1
	SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX SubpassDescriptionFlags = 2
	SUBPASS_DESCRIPTION_FLAG_BITS_MAX_ENUM               SubpassDescriptionFlags = 2147483647
)

type AccessFlags Flags

const (
	ACCESS_INDIRECT_COMMAND_READ_BIT                 AccessFlags = 1
	ACCESS_INDEX_READ_BIT                            AccessFlags = 2
	ACCESS_VERTEX_ATTRIBUTE_READ_BIT                 AccessFlags = 4
	ACCESS_UNIFORM_READ_BIT                          AccessFlags = 8
	ACCESS_INPUT_ATTACHMENT_READ_BIT                 AccessFlags = 16
	ACCESS_SHADER_READ_BIT                           AccessFlags = 32
	ACCESS_SHADER_WRITE_BIT                          AccessFlags = 64
	ACCESS_COLOR_ATTACHMENT_READ_BIT                 AccessFlags = 128
	ACCESS_COLOR_ATTACHMENT_WRITE_BIT                AccessFlags = 256
	ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT         AccessFlags = 512
	ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT        AccessFlags = 1024
	ACCESS_TRANSFER_READ_BIT                         AccessFlags = 2048
	ACCESS_TRANSFER_WRITE_BIT                        AccessFlags = 4096
	ACCESS_HOST_READ_BIT                             AccessFlags = 8192
	ACCESS_HOST_WRITE_BIT                            AccessFlags = 16384
	ACCESS_MEMORY_READ_BIT                           AccessFlags = 32768
	ACCESS_MEMORY_WRITE_BIT                          AccessFlags = 65536
	ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT        AccessFlags = 1048576
	ACCESS_COMMAND_PROCESS_READ_BIT_NVX              AccessFlags = 131072
	ACCESS_COMMAND_PROCESS_WRITE_BIT_NVX             AccessFlags = 262144
	ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT AccessFlags = 524288
	ACCESS_FLAG_BITS_MAX_ENUM                        AccessFlags = 2147483647
)

type DependencyFlags Flags

const (
	DEPENDENCY_BY_REGION_BIT        DependencyFlags = 1
	DEPENDENCY_DEVICE_GROUP_BIT     DependencyFlags = 4
	DEPENDENCY_VIEW_LOCAL_BIT       DependencyFlags = 2
	DEPENDENCY_VIEW_LOCAL_BIT_KHR   DependencyFlags = DEPENDENCY_VIEW_LOCAL_BIT
	DEPENDENCY_DEVICE_GROUP_BIT_KHR DependencyFlags = DEPENDENCY_DEVICE_GROUP_BIT
	DEPENDENCY_FLAG_BITS_MAX_ENUM   DependencyFlags = 2147483647
)

type CommandPoolCreateFlags Flags

const (
	COMMAND_POOL_CREATE_TRANSIENT_BIT            CommandPoolCreateFlags = 1
	COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT CommandPoolCreateFlags = 2
	COMMAND_POOL_CREATE_PROTECTED_BIT            CommandPoolCreateFlags = 4
	COMMAND_POOL_CREATE_FLAG_BITS_MAX_ENUM       CommandPoolCreateFlags = 2147483647
)

type CommandPoolResetFlags Flags

const (
	COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT CommandPoolResetFlags = 1
	COMMAND_POOL_RESET_FLAG_BITS_MAX_ENUM    CommandPoolResetFlags = 2147483647
)

type CommandBufferUsageFlags Flags

const (
	COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT      CommandBufferUsageFlags = 1
	COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT CommandBufferUsageFlags = 2
	COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT     CommandBufferUsageFlags = 4
	COMMAND_BUFFER_USAGE_FLAG_BITS_MAX_ENUM       CommandBufferUsageFlags = 2147483647
)

type QueryControlFlags Flags

const (
	QUERY_CONTROL_PRECISE_BIT        QueryControlFlags = 1
	QUERY_CONTROL_FLAG_BITS_MAX_ENUM QueryControlFlags = 2147483647
)

type CommandBufferResetFlags Flags

const (
	COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT CommandBufferResetFlags = 1
	COMMAND_BUFFER_RESET_FLAG_BITS_MAX_ENUM    CommandBufferResetFlags = 2147483647
)

type StencilFaceFlags Flags

const (
	STENCIL_FACE_FRONT_BIT          StencilFaceFlags = 1
	STENCIL_FACE_BACK_BIT           StencilFaceFlags = 2
	STENCIL_FRONT_AND_BACK          StencilFaceFlags = 3
	STENCIL_FACE_FLAG_BITS_MAX_ENUM StencilFaceFlags = 2147483647
)

type ApplicationInfo struct {
	Next               unsafe.Pointer
	ApplicationName    string
	ApplicationVersion uint32
	EngineName         string
	EngineVersion      uint32
	ApiVersion         uint32
}

func (s *ApplicationInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_APPLICATION_INFO
}

type InstanceCreateInfo struct {
	Next                  unsafe.Pointer
	Flags                 InstanceCreateFlags
	ApplicationInfo       *ApplicationInfo
	EnabledLayerNames     []string
	EnabledExtensionNames []string
}

func (s *InstanceCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
}

type PFN_vkAllocationFunction *func(unsafe.Pointer, uint, uint, SystemAllocationScope) unsafe.Pointer
type PFN_vkReallocationFunction *func(unsafe.Pointer, unsafe.Pointer, uint, uint, SystemAllocationScope) unsafe.Pointer
type PFN_vkFreeFunction *func(unsafe.Pointer, unsafe.Pointer)
type PFN_vkInternalAllocationNotification *func(unsafe.Pointer, uint, InternalAllocationType, SystemAllocationScope)
type PFN_vkInternalFreeNotification *func(unsafe.Pointer, uint, InternalAllocationType, SystemAllocationScope)
type AllocationCallbacks struct {
	UserData              unsafe.Pointer
	PfnAllocation         PFN_vkAllocationFunction
	PfnReallocation       PFN_vkReallocationFunction
	PfnFree               PFN_vkFreeFunction
	PfnInternalAllocation PFN_vkInternalAllocationNotification
	PfnInternalFree       PFN_vkInternalFreeNotification
}
type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      bool
	FullDrawIndexUint32                     bool
	ImageCubeArray                          bool
	IndependentBlend                        bool
	GeometryShader                          bool
	TessellationShader                      bool
	SampleRateShading                       bool
	DualSrcBlend                            bool
	LogicOp                                 bool
	MultiDrawIndirect                       bool
	DrawIndirectFirstInstance               bool
	DepthClamp                              bool
	DepthBiasClamp                          bool
	FillModeNonSolid                        bool
	DepthBounds                             bool
	WideLines                               bool
	LargePoints                             bool
	AlphaToOne                              bool
	MultiViewport                           bool
	SamplerAnisotropy                       bool
	TextureCompressionETC2                  bool
	TextureCompressionASTC_LDR              bool
	TextureCompressionBC                    bool
	OcclusionQueryPrecise                   bool
	PipelineStatisticsQuery                 bool
	VertexPipelineStoresAndAtomics          bool
	FragmentStoresAndAtomics                bool
	ShaderTessellationAndGeometryPointSize  bool
	ShaderImageGatherExtended               bool
	ShaderStorageImageExtendedFormats       bool
	ShaderStorageImageMultisample           bool
	ShaderStorageImageReadWithoutFormat     bool
	ShaderStorageImageWriteWithoutFormat    bool
	ShaderUniformBufferArrayDynamicIndexing bool
	ShaderSampledImageArrayDynamicIndexing  bool
	ShaderStorageBufferArrayDynamicIndexing bool
	ShaderStorageImageArrayDynamicIndexing  bool
	ShaderClipDistance                      bool
	ShaderCullDistance                      bool
	ShaderFloat64                           bool
	ShaderInt64                             bool
	ShaderInt16                             bool
	ShaderResourceResidency                 bool
	ShaderResourceMinLod                    bool
	SparseBinding                           bool
	SparseResidencyBuffer                   bool
	SparseResidencyImage2D                  bool
	SparseResidencyImage3D                  bool
	SparseResidency2Samples                 bool
	SparseResidency4Samples                 bool
	SparseResidency8Samples                 bool
	SparseResidency16Samples                bool
	SparseResidencyAliased                  bool
	VariableMultisampleRate                 bool
	InheritedQueries                        bool
}
type FormatProperties struct {
	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}
type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}
type ImageFormatProperties struct {
	MaxExtent       Extent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    SampleCountFlags
	MaxResourceSize DeviceSize
}
type PhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          DeviceSize
	SparseAddressSpaceSize                          DeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uint
	MinTexelBufferOffsetAlignment                   DeviceSize
	MinUniformBufferOffsetAlignment                 DeviceSize
	MinStorageBufferOffsetAlignment                 DeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    SampleCountFlags
	FramebufferDepthSampleCounts                    SampleCountFlags
	FramebufferStencilSampleCounts                  SampleCountFlags
	FramebufferNoAttachmentsSampleCounts            SampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   SampleCountFlags
	SampledImageIntegerSampleCounts                 SampleCountFlags
	SampledImageDepthSampleCounts                   SampleCountFlags
	SampledImageStencilSampleCounts                 SampleCountFlags
	StorageImageSampleCounts                        SampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     bool
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     bool
	StandardSampleLocations                         bool
	OptimalBufferCopyOffsetAlignment                DeviceSize
	OptimalBufferCopyRowPitchAlignment              DeviceSize
	NonCoherentAtomSize                             DeviceSize
}
type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}
type PhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        PhysicalDeviceType
	DeviceName        [256]byte
	PipelineCacheUUID [16]uint8
	Limits            PhysicalDeviceLimits
	SparseProperties  PhysicalDeviceSparseProperties
}
type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}
type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}
type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags
}
type PhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]MemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]MemoryHeap
}
type DeviceQueueCreateInfo struct {
	Next             unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

func (s *DeviceQueueCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
}

type DeviceCreateInfo struct {
	Next                  unsafe.Pointer
	Flags                 DeviceCreateFlags
	QueueCreateInfos      []DeviceQueueCreateInfo
	EnabledLayerNames     []string
	EnabledExtensionNames []string
	EnabledFeatures       []PhysicalDeviceFeatures
}

func (s *DeviceCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
}

type ExtensionProperties struct {
	ExtensionName [256]byte
	SpecVersion   uint32
}
type LayerProperties struct {
	LayerName             [256]byte
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [256]byte
}
type SubmitInfo struct {
	Next             unsafe.Pointer
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	CommandBuffers   []CommandBuffer
	SignalSemaphores []Semaphore
}

func (s *SubmitInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_SUBMIT_INFO
}

type MemoryAllocateInfo struct {
	Next            unsafe.Pointer
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

func (s *MemoryAllocateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
}

type MappedMemoryRange struct {
	Next   unsafe.Pointer
	Memory DeviceMemory
	Offset DeviceSize
	Size   DeviceSize
}

func (s *MappedMemoryRange) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
}

type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}
type SparseImageFormatProperties struct {
	AspectMask       ImageAspectFlags
	ImageGranularity Extent3D
	Flags            SparseImageFormatFlags
}
type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     DeviceSize
	ImageMipTailOffset   DeviceSize
	ImageMipTailStride   DeviceSize
}
type SparseMemoryBind struct {
	ResourceOffset DeviceSize
	Size           DeviceSize
	Memory         DeviceMemory
	MemoryOffset   DeviceSize
	Flags          SparseMemoryBindFlags
}
type SparseBufferMemoryBindInfo struct {
	Buffer Buffer
	Binds  []SparseMemoryBind
}
type SparseImageOpaqueMemoryBindInfo struct {
	Image Image
	Binds []SparseMemoryBind
}
type ImageSubresource struct {
	AspectMask ImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}
type Offset3D struct {
	X int32
	Y int32
	Z int32
}
type SparseImageMemoryBind struct {
	Subresource  ImageSubresource
	Offset       Offset3D
	Extent       Extent3D
	Memory       DeviceMemory
	MemoryOffset DeviceSize
	Flags        SparseMemoryBindFlags
}
type SparseImageMemoryBindInfo struct {
	Image Image
	Binds []SparseImageMemoryBind
}
type BindSparseInfo struct {
	Next             unsafe.Pointer
	WaitSemaphores   []Semaphore
	BufferBinds      []SparseBufferMemoryBindInfo
	ImageOpaqueBinds []SparseImageOpaqueMemoryBindInfo
	ImageBinds       []SparseImageMemoryBindInfo
	SignalSemaphores []Semaphore
}

func (s *BindSparseInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
}

type FenceCreateInfo struct {
	Next  unsafe.Pointer
	Flags FenceCreateFlags
}

func (s *FenceCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
}

type SemaphoreCreateInfo struct {
	Next  unsafe.Pointer
	Flags SemaphoreCreateFlags
}

func (s *SemaphoreCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
}

type EventCreateInfo struct {
	Next  unsafe.Pointer
	Flags EventCreateFlags
}

func (s *EventCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
}

type QueryPoolCreateInfo struct {
	Next               unsafe.Pointer
	Flags              QueryPoolCreateFlags
	QueryType          QueryType
	QueryCount         uint32
	PipelineStatistics QueryPipelineStatisticFlags
}

func (s *QueryPoolCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
}

type BufferCreateInfo struct {
	Next               unsafe.Pointer
	Flags              BufferCreateFlags
	Size               DeviceSize
	Usage              BufferUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
}

func (s *BufferCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
}

type BufferViewCreateInfo struct {
	Next   unsafe.Pointer
	Flags  BufferViewCreateFlags
	Buffer Buffer
	Format Format
	Offset DeviceSize
	Range  DeviceSize
}

func (s *BufferViewCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
}

type ImageCreateInfo struct {
	Next               unsafe.Pointer
	Flags              ImageCreateFlags
	ImageType          ImageType
	Format             Format
	Extent             Extent3D
	MipLevels          uint32
	ArrayLayers        uint32
	Samples            SampleCountFlags
	Tiling             ImageTiling
	Usage              ImageUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
	InitialLayout      ImageLayout
}

func (s *ImageCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
}

type SubresourceLayout struct {
	Offset     DeviceSize
	Size       DeviceSize
	RowPitch   DeviceSize
	ArrayPitch DeviceSize
	DepthPitch DeviceSize
}
type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}
type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}
type ImageViewCreateInfo struct {
	Next             unsafe.Pointer
	Flags            ImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

func (s *ImageViewCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
}

type ShaderModuleCreateInfo struct {
	Next     unsafe.Pointer
	Flags    ShaderModuleCreateFlags
	CodeSize uint
	Code     *uint32
}

func (s *ShaderModuleCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
}

type PipelineCacheCreateInfo struct {
	Next            unsafe.Pointer
	Flags           PipelineCacheCreateFlags
	InitialDataSize uint
	InitialData     unsafe.Pointer
}

func (s *PipelineCacheCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
}

type SpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uint
}
type SpecializationInfo struct {
	MapEntries []SpecializationMapEntry
	DataSize   uint
	Data       unsafe.Pointer
}
type PipelineShaderStageCreateInfo struct {
	Next               unsafe.Pointer
	Flags              PipelineShaderStageCreateFlags
	Stage              ShaderStageFlags
	Module             ShaderModule
	Name               string
	SpecializationInfo *SpecializationInfo
}

func (s *PipelineShaderStageCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
}

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}
type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}
type PipelineVertexInputStateCreateInfo struct {
	Next                        unsafe.Pointer
	Flags                       PipelineVertexInputStateCreateFlags
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (s *PipelineVertexInputStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
}

type PipelineInputAssemblyStateCreateInfo struct {
	Next                   unsafe.Pointer
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (s *PipelineInputAssemblyStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
}

type PipelineTessellationStateCreateInfo struct {
	Next               unsafe.Pointer
	Flags              PipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

func (s *PipelineTessellationStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
}

type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}
type Offset2D struct {
	X int32
	Y int32
}
type Extent2D struct {
	Width  uint32
	Height uint32
}
type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}
type PipelineViewportStateCreateInfo struct {
	Next      unsafe.Pointer
	Flags     PipelineViewportStateCreateFlags
	Viewports []Viewport
	Scissors  []Rect2D
}

func (s *PipelineViewportStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
}

type PipelineRasterizationStateCreateInfo struct {
	Next                    unsafe.Pointer
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        bool
	RasterizerDiscardEnable bool
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         bool
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

func (s *PipelineRasterizationStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
}

type PipelineMultisampleStateCreateInfo struct {
	Next                  unsafe.Pointer
	Flags                 PipelineMultisampleStateCreateFlags
	RasterizationSamples  SampleCountFlags
	SampleShadingEnable   bool
	MinSampleShading      float32
	SampleMask            *SampleMask
	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool
}

func (s *PipelineMultisampleStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
}

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}
type PipelineDepthStencilStateCreateInfo struct {
	Next                  unsafe.Pointer
	Flags                 PipelineDepthStencilStateCreateFlags
	DepthTestEnable       bool
	DepthWriteEnable      bool
	DepthCompareOp        CompareOp
	DepthBoundsTestEnable bool
	StencilTestEnable     bool
	Front                 StencilOpState
	Back                  StencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}

func (s *PipelineDepthStencilStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
}

type PipelineColorBlendAttachmentState struct {
	BlendEnable         bool
	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp
	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp
	ColorWriteMask      ColorComponentFlags
}
type PipelineColorBlendStateCreateInfo struct {
	Next           unsafe.Pointer
	Flags          PipelineColorBlendStateCreateFlags
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (s *PipelineColorBlendStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
}

type PipelineDynamicStateCreateInfo struct {
	Next          unsafe.Pointer
	Flags         PipelineDynamicStateCreateFlags
	DynamicStates []DynamicState
}

func (s *PipelineDynamicStateCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
}

type GraphicsPipelineCreateInfo struct {
	Next               unsafe.Pointer
	Flags              PipelineCreateFlags
	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   []PipelineVertexInputStateCreateInfo
	InputAssemblyState []PipelineInputAssemblyStateCreateInfo
	TessellationState  []PipelineTessellationStateCreateInfo
	ViewportState      []PipelineViewportStateCreateInfo
	RasterizationState []PipelineRasterizationStateCreateInfo
	MultisampleState   []PipelineMultisampleStateCreateInfo
	DepthStencilState  []PipelineDepthStencilStateCreateInfo
	ColorBlendState    []PipelineColorBlendStateCreateInfo
	DynamicState       []PipelineDynamicStateCreateInfo
	Layout             PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func (s *GraphicsPipelineCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
}

type ComputePipelineCreateInfo struct {
	Next               unsafe.Pointer
	Flags              PipelineCreateFlags
	Stage              PipelineShaderStageCreateInfo
	Layout             PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func (s *ComputePipelineCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
}

type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}
type PipelineLayoutCreateInfo struct {
	Next               unsafe.Pointer
	Flags              PipelineLayoutCreateFlags
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (s *PipelineLayoutCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
}

type SamplerCreateInfo struct {
	Next                    unsafe.Pointer
	Flags                   SamplerCreateFlags
	MagFilter               Filter
	MinFilter               Filter
	MipmapMode              SamplerMipmapMode
	AddressModeU            SamplerAddressMode
	AddressModeV            SamplerAddressMode
	AddressModeW            SamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        bool
	MaxAnisotropy           float32
	CompareEnable           bool
	CompareOp               CompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             BorderColor
	UnnormalizedCoordinates bool
}

func (s *SamplerCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO
}

type DescriptorSetLayoutBinding struct {
	Binding           uint32
	DescriptorType    DescriptorType
	DescriptorCount   uint32
	StageFlags        ShaderStageFlags
	ImmutableSamplers *Sampler
}
type DescriptorSetLayoutCreateInfo struct {
	Next     unsafe.Pointer
	Flags    DescriptorSetLayoutCreateFlags
	Bindings []DescriptorSetLayoutBinding
}

func (s *DescriptorSetLayoutCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
}

type DescriptorPoolSize struct {
	Type            DescriptorType
	DescriptorCount uint32
}
type DescriptorPoolCreateInfo struct {
	Next      unsafe.Pointer
	Flags     DescriptorPoolCreateFlags
	MaxSets   uint32
	PoolSizes []DescriptorPoolSize
}

func (s *DescriptorPoolCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
}

type DescriptorSetAllocateInfo struct {
	Next           unsafe.Pointer
	DescriptorPool DescriptorPool
	SetLayouts     []DescriptorSetLayout
}

func (s *DescriptorSetAllocateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
}

type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}
type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset DeviceSize
	Range  DeviceSize
}
type WriteDescriptorSet struct {
	Next            unsafe.Pointer
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  DescriptorType
	ImageInfo       *DescriptorImageInfo
	BufferInfo      *DescriptorBufferInfo
	TexelBufferView *BufferView
}

func (s *WriteDescriptorSet) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
}

type CopyDescriptorSet struct {
	Next            unsafe.Pointer
	SrcSet          DescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

func (s *CopyDescriptorSet) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
}

type FramebufferCreateInfo struct {
	Next        unsafe.Pointer
	Flags       FramebufferCreateFlags
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (s *FramebufferCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
}

type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlags
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}
type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
}
type SubpassDescription struct {
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment []AttachmentReference
	PreserveAttachments    []uint32
}
type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}
type RenderPassCreateInfo struct {
	Next         unsafe.Pointer
	Flags        RenderPassCreateFlags
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

func (s *RenderPassCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
}

type CommandPoolCreateInfo struct {
	Next             unsafe.Pointer
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

func (s *CommandPoolCreateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
}

type CommandBufferAllocateInfo struct {
	Next               unsafe.Pointer
	CommandPool        CommandPool
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

func (s *CommandBufferAllocateInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
}

type CommandBufferInheritanceInfo struct {
	Next                 unsafe.Pointer
	RenderPass           RenderPass
	Subpass              uint32
	Framebuffer          Framebuffer
	OcclusionQueryEnable bool
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags
}

func (s *CommandBufferInheritanceInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
}

type CommandBufferBeginInfo struct {
	Next            unsafe.Pointer
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

func (s *CommandBufferBeginInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
}

type BufferCopy struct {
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}
type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}
type ImageCopy struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}
type ImageBlit struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}
type BufferImageCopy struct {
	BufferOffset      DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D
}
type ClearColorValue struct {
	Float32 [4]float32
	Int32   [4]int32
	Uint32  [4]uint32
}
type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}
type ClearValue struct {
	Color        ClearColorValue
	DepthStencil ClearDepthStencilValue
}
type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}
type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}
type ImageResolve struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}
type MemoryBarrier struct {
	Next          unsafe.Pointer
	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

func (s *MemoryBarrier) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_MEMORY_BARRIER
}

type BufferMemoryBarrier struct {
	Next                unsafe.Pointer
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              Buffer
	Offset              DeviceSize
	Size                DeviceSize
}

func (s *BufferMemoryBarrier) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
}

type ImageMemoryBarrier struct {
	Next                unsafe.Pointer
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	OldLayout           ImageLayout
	NewLayout           ImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               Image
	SubresourceRange    ImageSubresourceRange
}

func (s *ImageMemoryBarrier) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
}

type RenderPassBeginInfo struct {
	Next        unsafe.Pointer
	RenderPass  RenderPass
	Framebuffer Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

func (s *RenderPassBeginInfo) sType() C.VkStructureType {
	return C.VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
}

type DispatchIndirectCommand struct {
	X uint32
	Y uint32
	Z uint32
}
type DrawIndexedIndirectCommand struct {
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}
type DrawIndirectCommand struct {
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}
