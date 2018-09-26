package vk
//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
//typedef void * void_pointer;
import "C"

type StructureType int

const (
	STRUCTURE_TYPE_APPLICATION_INFO                                             StructureType = 0
	STRUCTURE_TYPE_INSTANCE_CREATE_INFO                                         StructureType = 1
	STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO                                     StructureType = 2
	STRUCTURE_TYPE_DEVICE_CREATE_INFO                                           StructureType = 3
	STRUCTURE_TYPE_SUBMIT_INFO                                                  StructureType = 4
	STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO                                         StructureType = 5
	STRUCTURE_TYPE_MAPPED_MEMORY_RANGE                                          StructureType = 6
	STRUCTURE_TYPE_BIND_SPARSE_INFO                                             StructureType = 7
	STRUCTURE_TYPE_FENCE_CREATE_INFO                                            StructureType = 8
	STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO                                        StructureType = 9
	STRUCTURE_TYPE_EVENT_CREATE_INFO                                            StructureType = 10
	STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO                                       StructureType = 11
	STRUCTURE_TYPE_BUFFER_CREATE_INFO                                           StructureType = 12
	STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO                                      StructureType = 13
	STRUCTURE_TYPE_IMAGE_CREATE_INFO                                            StructureType = 14
	STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO                                       StructureType = 15
	STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO                                    StructureType = 16
	STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO                                   StructureType = 17
	STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO                            StructureType = 18
	STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO                      StructureType = 19
	STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO                    StructureType = 20
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO                      StructureType = 21
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO                          StructureType = 22
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO                     StructureType = 23
	STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO                       StructureType = 24
	STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO                     StructureType = 25
	STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO                       StructureType = 26
	STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO                           StructureType = 27
	STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO                                StructureType = 28
	STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO                                 StructureType = 29
	STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO                                  StructureType = 30
	STRUCTURE_TYPE_SAMPLER_CREATE_INFO                                          StructureType = 31
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO                            StructureType = 32
	STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO                                  StructureType = 33
	STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO                                 StructureType = 34
	STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET                                         StructureType = 35
	STRUCTURE_TYPE_COPY_DESCRIPTOR_SET                                          StructureType = 36
	STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO                                      StructureType = 37
	STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO                                      StructureType = 38
	STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO                                     StructureType = 39
	STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO                                 StructureType = 40
	STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO                              StructureType = 41
	STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO                                    StructureType = 42
	STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO                                       StructureType = 43
	STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER                                        StructureType = 44
	STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER                                         StructureType = 45
	STRUCTURE_TYPE_MEMORY_BARRIER                                               StructureType = 46
	STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO                                  StructureType = 47
	STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO                                    StructureType = 48
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES                          StructureType = 1000094000
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO                                      StructureType = 1000157000
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO                                       StructureType = 1000157001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES                       StructureType = 1000083000
	STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS                                StructureType = 1000127000
	STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO                               StructureType = 1000127001
	STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO                                   StructureType = 1000060000
	STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO                          StructureType = 1000060003
	STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO                       StructureType = 1000060004
	STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO                                     StructureType = 1000060005
	STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO                                StructureType = 1000060006
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO                         StructureType = 1000060013
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO                          StructureType = 1000060014
	STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES                             StructureType = 1000070000
	STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO                              StructureType = 1000070001
	STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2                            StructureType = 1000146000
	STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2                             StructureType = 1000146001
	STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2                      StructureType = 1000146002
	STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2                                        StructureType = 1000146003
	STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2                           StructureType = 1000146004
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2                                   StructureType = 1000059000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2                                 StructureType = 1000059001
	STRUCTURE_TYPE_FORMAT_PROPERTIES_2                                          StructureType = 1000059002
	STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2                                    StructureType = 1000059003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2                          StructureType = 1000059004
	STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2                                    StructureType = 1000059005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2                          StructureType = 1000059006
	STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2                             StructureType = 1000059007
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2                   StructureType = 1000059008
	STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES                    StructureType = 1000117000
	STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO              StructureType = 1000117001
	STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO                                 StructureType = 1000117002
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO        StructureType = 1000117003
	STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO                            StructureType = 1000053000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES                           StructureType = 1000053001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES                         StructureType = 1000053002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES                    StructureType = 1000120000
	STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO                                        StructureType = 1000145000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES                    StructureType = 1000145001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES                  StructureType = 1000145002
	STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2                                          StructureType = 1000145003
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO                         StructureType = 1000156000
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO                                StructureType = 1000156001
	STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO                                 StructureType = 1000156002
	STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO                         StructureType = 1000156003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES            StructureType = 1000156004
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES             StructureType = 1000156005
	STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO                       StructureType = 1000085000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO                   StructureType = 1000071000
	STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES                             StructureType = 1000071001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO                         StructureType = 1000071002
	STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES                                   StructureType = 1000071003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES                                StructureType = 1000071004
	STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO                           StructureType = 1000072000
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO                            StructureType = 1000072001
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO                                  StructureType = 1000072002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO                          StructureType = 1000112000
	STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES                                    StructureType = 1000112001
	STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO                                     StructureType = 1000113000
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO                                 StructureType = 1000077000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO                      StructureType = 1000076000
	STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES                                StructureType = 1000076001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES                     StructureType = 1000168000
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT                                StructureType = 1000168001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES               StructureType = 1000063000
	STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR                                    StructureType = 1000001000
	STRUCTURE_TYPE_PRESENT_INFO_KHR                                             StructureType = 1000001001
	STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR                        StructureType = 1000060007
	STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR                              StructureType = 1000060008
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR                         StructureType = 1000060009
	STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR                                  StructureType = 1000060010
	STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR                                StructureType = 1000060011
	STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR                       StructureType = 1000060012
	STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR                                 StructureType = 1000002000
	STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR                              StructureType = 1000002001
	STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR                                     StructureType = 1000003000
	STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR                                 StructureType = 1000004000
	STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR                                  StructureType = 1000005000
	STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR                              StructureType = 1000006000
	STRUCTURE_TYPE_MIR_SURFACE_CREATE_INFO_KHR                                  StructureType = 1000007000
	STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR                              StructureType = 1000008000
	STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR                                StructureType = 1000009000
	STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT                        StructureType = 1000011000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD         StructureType = 1000018000
	STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT                            StructureType = 1000022000
	STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT                             StructureType = 1000022001
	STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT                                 StructureType = 1000022002
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV                    StructureType = 1000026000
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV                   StructureType = 1000026001
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV                 StructureType = 1000026002
	STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD                     StructureType = 1000041000
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV                         StructureType = 1000056000
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV                               StructureType = 1000056001
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV                           StructureType = 1000057000
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV                           StructureType = 1000057001
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV                    StructureType = 1000058000
	STRUCTURE_TYPE_VALIDATION_FLAGS_EXT                                         StructureType = 1000061000
	STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN                                    StructureType = 1000062000
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR                          StructureType = 1000073000
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR                          StructureType = 1000073001
	STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR                           StructureType = 1000073002
	STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR                             StructureType = 1000073003
	STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR                                    StructureType = 1000074000
	STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR                                     StructureType = 1000074001
	STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR                                       StructureType = 1000074002
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR                   StructureType = 1000075000
	STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                       StructureType = 1000078000
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                       StructureType = 1000078001
	STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR                                  StructureType = 1000078002
	STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR                          StructureType = 1000078003
	STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR                                 StructureType = 1000079000
	STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR                                    StructureType = 1000079001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR               StructureType = 1000080000
	STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT    StructureType = 1000081000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT           StructureType = 1000081001
	STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT                         StructureType = 1000081002
	STRUCTURE_TYPE_PRESENT_REGIONS_KHR                                          StructureType = 1000084000
	STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX                                 StructureType = 1000086000
	STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX                     StructureType = 1000086001
	STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX                                StructureType = 1000086002
	STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX                      StructureType = 1000086003
	STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX                         StructureType = 1000086004
	STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX                       StructureType = 1000086005
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV             StructureType = 1000087000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT                                   StructureType = 1000090000
	STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT                                       StructureType = 1000091000
	STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT                                        StructureType = 1000091001
	STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT                                       StructureType = 1000091002
	STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT                            StructureType = 1000091003
	STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE                                    StructureType = 1000092000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX StructureType = 1000097000
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV               StructureType = 1000098000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT             StructureType = 1000099000
	STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT             StructureType = 1000099001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT    StructureType = 1000101000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT    StructureType = 1000101001
	STRUCTURE_TYPE_HDR_METADATA_EXT                                             StructureType = 1000105000
	STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR                                 StructureType = 1000109000
	STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR                                   StructureType = 1000109001
	STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR                                    StructureType = 1000109002
	STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR                                     StructureType = 1000109003
	STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR                                StructureType = 1000109004
	STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR                                       StructureType = 1000109005
	STRUCTURE_TYPE_SUBPASS_END_INFO_KHR                                         StructureType = 1000109006
	STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR                      StructureType = 1000111000
	STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR                           StructureType = 1000114000
	STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR                           StructureType = 1000114001
	STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR                              StructureType = 1000114002
	STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR                                     StructureType = 1000115000
	STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR                                        StructureType = 1000115001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR                           StructureType = 1000119000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR                                   StructureType = 1000119001
	STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR                                         StructureType = 1000119002
	STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR                                     StructureType = 1000121000
	STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR                               StructureType = 1000121001
	STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR                                StructureType = 1000121002
	STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR                                     StructureType = 1000121003
	STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR                             StructureType = 1000121004
	STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK                                  StructureType = 1000122000
	STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK                                StructureType = 1000123000
	STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT                             StructureType = 1000128000
	STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT                              StructureType = 1000128001
	STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT                                        StructureType = 1000128002
	STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT                      StructureType = 1000128003
	STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT                        StructureType = 1000128004
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID                        StructureType = 1000129000
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID                   StructureType = 1000129001
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID            StructureType = 1000129002
	STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                  StructureType = 1000129003
	STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID              StructureType = 1000129004
	STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID                                      StructureType = 1000129005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT         StructureType = 1000130000
	STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT                       StructureType = 1000130001
	STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT                                    StructureType = 1000143000
	STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT                  StructureType = 1000143001
	STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT              StructureType = 1000143002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT              StructureType = 1000143003
	STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT                                   StructureType = 1000143004
	STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR                            StructureType = 1000147000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT        StructureType = 1000148000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT      StructureType = 1000148001
	STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT          StructureType = 1000148002
	STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV              StructureType = 1000149000
	STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV            StructureType = 1000152000
	STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT                             StructureType = 1000160000
	STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT               StructureType = 1000160001
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT          StructureType = 1000161000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT             StructureType = 1000161001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT           StructureType = 1000161002
	STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT   StructureType = 1000161003
	STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT  StructureType = 1000161004
	STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT                 StructureType = 1000174000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR                    StructureType = 1000177000
	STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT                          StructureType = 1000178000
	STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT                           StructureType = 1000178001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT          StructureType = 1000178002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD                   StructureType = 1000185000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT      StructureType = 1000190000
	STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT          StructureType = 1000190001
	STRUCTURE_TYPE_CHECKPOINT_DATA_NV                                           StructureType = 1000206000
	STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV                        StructureType = 1000206001
	STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR                        StructureType = STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR                       StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR                     StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR                               StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR                             StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR                                      StructureType = STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR                                StructureType = STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR                      StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR                                StructureType = STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR                      StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR                         StructureType = STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR               StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR                               StructureType = STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR                      StructureType = STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR                   StructureType = STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR                                 StructureType = STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR                            StructureType = STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR                     StructureType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR                      StructureType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR                         StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR                          StructureType = STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR               StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR                         StructureType = STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR                     StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR                               StructureType = STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR                            StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR                       StructureType = STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR                        StructureType = STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR                              StructureType = STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR                            StructureType = STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR                             StructureType = STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR                   StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR                   StructureType = STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR                      StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR                                StructureType = STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR                                 StructureType = STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR                StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR          StructureType = STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR                             StructureType = STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR    StructureType = STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR                StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES
	STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR                            StructureType = STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR                           StructureType = STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR                        StructureType = STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR                         StructureType = STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR                  StructureType = STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR                                    StructureType = STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR                       StructureType = STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR                     StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR                            StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR                             StructureType = STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR                     StructureType = STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR        StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR         StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR                                  StructureType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR                                   StructureType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR                 StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR                            StructureType = STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	STRUCTURE_TYPE_BEGIN_RANGE                                                  StructureType = STRUCTURE_TYPE_APPLICATION_INFO
	STRUCTURE_TYPE_END_RANGE                                                    StructureType = STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO
	STRUCTURE_TYPE_RANGE_SIZE                                                   StructureType = (STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO - STRUCTURE_TYPE_APPLICATION_INFO + 1)
	STRUCTURE_TYPE_MAX_ENUM                                                     StructureType = 2147483647
)

type Flags uint32
type InstanceCreateFlags Flags
type ApplicationInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	ApplicationName    string
	ApplicationVersion uint32
	EngineName         string
	EngineVersion      uint32
	ApiVersion         uint32
}

func (g *ApplicationInfo) toC(c *C.VkApplicationInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.pApplicationName = toCString(g.ApplicationName, _sa)
	c.applicationVersion = C.uint32_t(g.ApplicationVersion)
	c.pEngineName = toCString(g.EngineName, _sa)
	c.engineVersion = C.uint32_t(g.EngineVersion)
	c.apiVersion = C.uint32_t(g.ApiVersion)
}
func (g *ApplicationInfo) fromC(c *C.VkApplicationInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.ApplicationName = toGoString(c.pApplicationName)
	g.ApplicationVersion = uint32(c.applicationVersion)
	g.EngineName = toGoString(c.pEngineName)
	g.EngineVersion = uint32(c.engineVersion)
	g.ApiVersion = uint32(c.apiVersion)
}

type InstanceCreateInfo struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	Flags                 InstanceCreateFlags
	ApplicationInfo       *ApplicationInfo
	EnabledLayerNames     []string
	EnabledExtensionNames []string
}

func (g *InstanceCreateInfo) toC(c *C.VkInstanceCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkInstanceCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkInstanceCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkInstanceCreateFlags(temp_in_VkInstanceCreateFlags)
	}
	{
		c.pApplicationInfo = (*C.VkApplicationInfo)(_sa.alloc(C.sizeof_VkApplicationInfo))
		g.ApplicationInfo.toC(c.pApplicationInfo, _sa)
	}
	c.enabledLayerCount = C.uint32_t(len(g.EnabledLayerNames))
	{
		c.ppEnabledLayerNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.EnabledLayerNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.EnabledLayerNames):len(g.EnabledLayerNames)]
		for i2, _ := range g.EnabledLayerNames {
			slice2[i2] = toCString(g.EnabledLayerNames[i2], _sa)
		}
	}
	c.enabledExtensionCount = C.uint32_t(len(g.EnabledExtensionNames))
	{
		c.ppEnabledExtensionNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.EnabledExtensionNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.EnabledExtensionNames):len(g.EnabledExtensionNames)]
		for i2, _ := range g.EnabledExtensionNames {
			slice2[i2] = toCString(g.EnabledExtensionNames[i2], _sa)
		}
	}
}
func (g *InstanceCreateInfo) fromC(c *C.VkInstanceCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkInstanceCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkInstanceCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = InstanceCreateFlags(temp_in_VkInstanceCreateFlags)
	}
	{
		if g.ApplicationInfo == nil {
			g.ApplicationInfo = new(ApplicationInfo)
		}
		g.ApplicationInfo.fromC(c.pApplicationInfo)
	}
	g.EnabledLayerNames = make([]string, int(c.enabledLayerCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.EnabledLayerNames):len(g.EnabledLayerNames)]
		for i2, _ := range g.EnabledLayerNames {
			g.EnabledLayerNames[i2] = toGoString(slice2[i2])
		}
	}
	g.EnabledExtensionNames = make([]string, int(c.enabledExtensionCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.EnabledExtensionNames):len(g.EnabledExtensionNames)]
		for i2, _ := range g.EnabledExtensionNames {
			g.EnabledExtensionNames[i2] = toGoString(slice2[i2])
		}
	}
}

type PFN_vkAllocationFunction struct {
	Raw C.PFN_vkAllocationFunction
}
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

func (p PFN_vkAllocationFunction) Call(arg0 unsafe.Pointer, arg1 uint, arg2 uint, arg3 SystemAllocationScope) (_ret unsafe.Pointer) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 C.size_t
		arg2 C.size_t
		arg3 C.VkSystemAllocationScope
		_ret unsafe.Pointer
	}
	c.arg0 = arg0
	c.arg1 = C.size_t(arg1)
	c.arg2 = C.size_t(arg2)
	c.arg3 = C.VkSystemAllocationScope(arg3)
	c._ret = C.callPFN_vkAllocationFunction(p.Raw, c.arg0, c.arg1, c.arg2, c.arg3)
	_ret = c._ret
	return
}

type PFN_vkReallocationFunction struct {
	Raw C.PFN_vkReallocationFunction
}

func (p PFN_vkReallocationFunction) Call(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uint, arg3 uint, arg4 SystemAllocationScope) (_ret unsafe.Pointer) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 unsafe.Pointer
		arg2 C.size_t
		arg3 C.size_t
		arg4 C.VkSystemAllocationScope
		_ret unsafe.Pointer
	}
	c.arg0 = arg0
	c.arg1 = arg1
	c.arg2 = C.size_t(arg2)
	c.arg3 = C.size_t(arg3)
	c.arg4 = C.VkSystemAllocationScope(arg4)
	c._ret = C.callPFN_vkReallocationFunction(p.Raw, c.arg0, c.arg1, c.arg2, c.arg3, c.arg4)
	_ret = c._ret
	return
}

type PFN_vkFreeFunction struct {
	Raw C.PFN_vkFreeFunction
}

func (p PFN_vkFreeFunction) Call(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 unsafe.Pointer
	}
	c.arg0 = arg0
	c.arg1 = arg1
	C.callPFN_vkFreeFunction(p.Raw, c.arg0, c.arg1)
}

type PFN_vkInternalAllocationNotification struct {
	Raw C.PFN_vkInternalAllocationNotification
}
type InternalAllocationType int

const (
	INTERNAL_ALLOCATION_TYPE_EXECUTABLE  InternalAllocationType = 0
	INTERNAL_ALLOCATION_TYPE_BEGIN_RANGE InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_END_RANGE   InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_RANGE_SIZE  InternalAllocationType = (INTERNAL_ALLOCATION_TYPE_EXECUTABLE - INTERNAL_ALLOCATION_TYPE_EXECUTABLE + 1)
	INTERNAL_ALLOCATION_TYPE_MAX_ENUM    InternalAllocationType = 2147483647
)

func (p PFN_vkInternalAllocationNotification) Call(arg0 unsafe.Pointer, arg1 uint, arg2 InternalAllocationType, arg3 SystemAllocationScope) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 C.size_t
		arg2 C.VkInternalAllocationType
		arg3 C.VkSystemAllocationScope
	}
	c.arg0 = arg0
	c.arg1 = C.size_t(arg1)
	c.arg2 = C.VkInternalAllocationType(arg2)
	c.arg3 = C.VkSystemAllocationScope(arg3)
	C.callPFN_vkInternalAllocationNotification(p.Raw, c.arg0, c.arg1, c.arg2, c.arg3)
}

type PFN_vkInternalFreeNotification struct {
	Raw C.PFN_vkInternalFreeNotification
}

func (p PFN_vkInternalFreeNotification) Call(arg0 unsafe.Pointer, arg1 uint, arg2 InternalAllocationType, arg3 SystemAllocationScope) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 C.size_t
		arg2 C.VkInternalAllocationType
		arg3 C.VkSystemAllocationScope
	}
	c.arg0 = arg0
	c.arg1 = C.size_t(arg1)
	c.arg2 = C.VkInternalAllocationType(arg2)
	c.arg3 = C.VkSystemAllocationScope(arg3)
	C.callPFN_vkInternalFreeNotification(p.Raw, c.arg0, c.arg1, c.arg2, c.arg3)
}

type AllocationCallbacks struct {
	UserData           []byte
	Allocation         PFN_vkAllocationFunction
	Reallocation       PFN_vkReallocationFunction
	Free               PFN_vkFreeFunction
	InternalAllocation PFN_vkInternalAllocationNotification
	InternalFree       PFN_vkInternalFreeNotification
}

func (g *AllocationCallbacks) toC(c *C.VkAllocationCallbacks, _sa *stackAllocator) {
	{
		c.pUserData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.UserData)))
		slice2 := (*[1 << 31]byte)(c.pUserData)[:len(g.UserData):len(g.UserData)]
		for i2, _ := range g.UserData {
			slice2[i2] = g.UserData[i2]
		}
	}
	c.pfnAllocation = g.Allocation.Raw
	c.pfnReallocation = g.Reallocation.Raw
	c.pfnFree = g.Free.Raw
	c.pfnInternalAllocation = g.InternalAllocation.Raw
	c.pfnInternalFree = g.InternalFree.Raw
}
func (g *AllocationCallbacks) fromC(c *C.VkAllocationCallbacks) {
	{
		slice2 := (*[1 << 31]byte)(c.pUserData)[:len(g.UserData):len(g.UserData)]
		for i2, _ := range g.UserData {
			g.UserData[i2] = slice2[i2]
		}
	}
	g.Allocation.Raw = c.pfnAllocation
	g.Reallocation.Raw = c.pfnReallocation
	g.Free.Raw = c.pfnFree
	g.InternalAllocation.Raw = c.pfnInternalAllocation
	g.InternalFree.Raw = c.pfnInternalFree
}

type Instance C.VkInstance
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

func CreateInstance(createInfo *InstanceCreateInfo, allocator *AllocationCallbacks, instance *Instance) (_ret Result) {
	var c struct {
		pCreateInfo *C.VkInstanceCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pInstance   *C.VkInstance
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pCreateInfo = (*C.VkInstanceCreateInfo)(_sa.alloc(C.sizeof_VkInstanceCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pInstance = (*C.VkInstance)(_sa.alloc(C.sizeof_VkInstance))
		*c.pInstance = C.VkInstance(*instance)
	}
	c._ret = C.vkCreateInstance(c.pCreateInfo, c.pAllocator, c.pInstance)
	_ret = Result(c._ret)
	*instance = Instance(*c.pInstance)
	return
}
func DestroyInstance(instance Instance, allocator *AllocationCallbacks) {
	var c struct {
		instance   C.VkInstance
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyInstance(c.instance, c.pAllocator)
}

type PhysicalDevice C.VkPhysicalDevice

func EnumeratePhysicalDevices(instance Instance, physicalDeviceCount *uint32, physicalDevices []PhysicalDevice) (_ret Result) {
	var c struct {
		instance             C.VkInstance
		pPhysicalDeviceCount *C.uint32_t
		pPhysicalDevices     *C.VkPhysicalDevice
		_ret                 C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pPhysicalDeviceCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPhysicalDeviceCount = C.uint32_t(*physicalDeviceCount)
	}
	{
		c.pPhysicalDevices = (*C.VkPhysicalDevice)(_sa.alloc(C.sizeof_VkPhysicalDevice * uint(len(physicalDevices))))
		slice3 := (*[1 << 31]C.VkPhysicalDevice)(unsafe.Pointer(c.pPhysicalDevices))[:len(physicalDevices):len(physicalDevices)]
		for i3, _ := range physicalDevices {
			slice3[i3] = C.VkPhysicalDevice(physicalDevices[i3])
		}
	}
	c._ret = C.vkEnumeratePhysicalDevices(c.instance, c.pPhysicalDeviceCount, c.pPhysicalDevices)
	_ret = Result(c._ret)
	*physicalDeviceCount = uint32(*c.pPhysicalDeviceCount)
	return
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

func (g *PhysicalDeviceFeatures) toC(c *C.VkPhysicalDeviceFeatures) {
	if g.RobustBufferAccess {
		c.robustBufferAccess = 1
	} else {
		c.robustBufferAccess = 0
	}
	if g.FullDrawIndexUint32 {
		c.fullDrawIndexUint32 = 1
	} else {
		c.fullDrawIndexUint32 = 0
	}
	if g.ImageCubeArray {
		c.imageCubeArray = 1
	} else {
		c.imageCubeArray = 0
	}
	if g.IndependentBlend {
		c.independentBlend = 1
	} else {
		c.independentBlend = 0
	}
	if g.GeometryShader {
		c.geometryShader = 1
	} else {
		c.geometryShader = 0
	}
	if g.TessellationShader {
		c.tessellationShader = 1
	} else {
		c.tessellationShader = 0
	}
	if g.SampleRateShading {
		c.sampleRateShading = 1
	} else {
		c.sampleRateShading = 0
	}
	if g.DualSrcBlend {
		c.dualSrcBlend = 1
	} else {
		c.dualSrcBlend = 0
	}
	if g.LogicOp {
		c.logicOp = 1
	} else {
		c.logicOp = 0
	}
	if g.MultiDrawIndirect {
		c.multiDrawIndirect = 1
	} else {
		c.multiDrawIndirect = 0
	}
	if g.DrawIndirectFirstInstance {
		c.drawIndirectFirstInstance = 1
	} else {
		c.drawIndirectFirstInstance = 0
	}
	if g.DepthClamp {
		c.depthClamp = 1
	} else {
		c.depthClamp = 0
	}
	if g.DepthBiasClamp {
		c.depthBiasClamp = 1
	} else {
		c.depthBiasClamp = 0
	}
	if g.FillModeNonSolid {
		c.fillModeNonSolid = 1
	} else {
		c.fillModeNonSolid = 0
	}
	if g.DepthBounds {
		c.depthBounds = 1
	} else {
		c.depthBounds = 0
	}
	if g.WideLines {
		c.wideLines = 1
	} else {
		c.wideLines = 0
	}
	if g.LargePoints {
		c.largePoints = 1
	} else {
		c.largePoints = 0
	}
	if g.AlphaToOne {
		c.alphaToOne = 1
	} else {
		c.alphaToOne = 0
	}
	if g.MultiViewport {
		c.multiViewport = 1
	} else {
		c.multiViewport = 0
	}
	if g.SamplerAnisotropy {
		c.samplerAnisotropy = 1
	} else {
		c.samplerAnisotropy = 0
	}
	if g.TextureCompressionETC2 {
		c.textureCompressionETC2 = 1
	} else {
		c.textureCompressionETC2 = 0
	}
	if g.TextureCompressionASTC_LDR {
		c.textureCompressionASTC_LDR = 1
	} else {
		c.textureCompressionASTC_LDR = 0
	}
	if g.TextureCompressionBC {
		c.textureCompressionBC = 1
	} else {
		c.textureCompressionBC = 0
	}
	if g.OcclusionQueryPrecise {
		c.occlusionQueryPrecise = 1
	} else {
		c.occlusionQueryPrecise = 0
	}
	if g.PipelineStatisticsQuery {
		c.pipelineStatisticsQuery = 1
	} else {
		c.pipelineStatisticsQuery = 0
	}
	if g.VertexPipelineStoresAndAtomics {
		c.vertexPipelineStoresAndAtomics = 1
	} else {
		c.vertexPipelineStoresAndAtomics = 0
	}
	if g.FragmentStoresAndAtomics {
		c.fragmentStoresAndAtomics = 1
	} else {
		c.fragmentStoresAndAtomics = 0
	}
	if g.ShaderTessellationAndGeometryPointSize {
		c.shaderTessellationAndGeometryPointSize = 1
	} else {
		c.shaderTessellationAndGeometryPointSize = 0
	}
	if g.ShaderImageGatherExtended {
		c.shaderImageGatherExtended = 1
	} else {
		c.shaderImageGatherExtended = 0
	}
	if g.ShaderStorageImageExtendedFormats {
		c.shaderStorageImageExtendedFormats = 1
	} else {
		c.shaderStorageImageExtendedFormats = 0
	}
	if g.ShaderStorageImageMultisample {
		c.shaderStorageImageMultisample = 1
	} else {
		c.shaderStorageImageMultisample = 0
	}
	if g.ShaderStorageImageReadWithoutFormat {
		c.shaderStorageImageReadWithoutFormat = 1
	} else {
		c.shaderStorageImageReadWithoutFormat = 0
	}
	if g.ShaderStorageImageWriteWithoutFormat {
		c.shaderStorageImageWriteWithoutFormat = 1
	} else {
		c.shaderStorageImageWriteWithoutFormat = 0
	}
	if g.ShaderUniformBufferArrayDynamicIndexing {
		c.shaderUniformBufferArrayDynamicIndexing = 1
	} else {
		c.shaderUniformBufferArrayDynamicIndexing = 0
	}
	if g.ShaderSampledImageArrayDynamicIndexing {
		c.shaderSampledImageArrayDynamicIndexing = 1
	} else {
		c.shaderSampledImageArrayDynamicIndexing = 0
	}
	if g.ShaderStorageBufferArrayDynamicIndexing {
		c.shaderStorageBufferArrayDynamicIndexing = 1
	} else {
		c.shaderStorageBufferArrayDynamicIndexing = 0
	}
	if g.ShaderStorageImageArrayDynamicIndexing {
		c.shaderStorageImageArrayDynamicIndexing = 1
	} else {
		c.shaderStorageImageArrayDynamicIndexing = 0
	}
	if g.ShaderClipDistance {
		c.shaderClipDistance = 1
	} else {
		c.shaderClipDistance = 0
	}
	if g.ShaderCullDistance {
		c.shaderCullDistance = 1
	} else {
		c.shaderCullDistance = 0
	}
	if g.ShaderFloat64 {
		c.shaderFloat64 = 1
	} else {
		c.shaderFloat64 = 0
	}
	if g.ShaderInt64 {
		c.shaderInt64 = 1
	} else {
		c.shaderInt64 = 0
	}
	if g.ShaderInt16 {
		c.shaderInt16 = 1
	} else {
		c.shaderInt16 = 0
	}
	if g.ShaderResourceResidency {
		c.shaderResourceResidency = 1
	} else {
		c.shaderResourceResidency = 0
	}
	if g.ShaderResourceMinLod {
		c.shaderResourceMinLod = 1
	} else {
		c.shaderResourceMinLod = 0
	}
	if g.SparseBinding {
		c.sparseBinding = 1
	} else {
		c.sparseBinding = 0
	}
	if g.SparseResidencyBuffer {
		c.sparseResidencyBuffer = 1
	} else {
		c.sparseResidencyBuffer = 0
	}
	if g.SparseResidencyImage2D {
		c.sparseResidencyImage2D = 1
	} else {
		c.sparseResidencyImage2D = 0
	}
	if g.SparseResidencyImage3D {
		c.sparseResidencyImage3D = 1
	} else {
		c.sparseResidencyImage3D = 0
	}
	if g.SparseResidency2Samples {
		c.sparseResidency2Samples = 1
	} else {
		c.sparseResidency2Samples = 0
	}
	if g.SparseResidency4Samples {
		c.sparseResidency4Samples = 1
	} else {
		c.sparseResidency4Samples = 0
	}
	if g.SparseResidency8Samples {
		c.sparseResidency8Samples = 1
	} else {
		c.sparseResidency8Samples = 0
	}
	if g.SparseResidency16Samples {
		c.sparseResidency16Samples = 1
	} else {
		c.sparseResidency16Samples = 0
	}
	if g.SparseResidencyAliased {
		c.sparseResidencyAliased = 1
	} else {
		c.sparseResidencyAliased = 0
	}
	if g.VariableMultisampleRate {
		c.variableMultisampleRate = 1
	} else {
		c.variableMultisampleRate = 0
	}
	if g.InheritedQueries {
		c.inheritedQueries = 1
	} else {
		c.inheritedQueries = 0
	}
}
func (g *PhysicalDeviceFeatures) fromC(c *C.VkPhysicalDeviceFeatures) {
	g.RobustBufferAccess = c.robustBufferAccess != 0
	g.FullDrawIndexUint32 = c.fullDrawIndexUint32 != 0
	g.ImageCubeArray = c.imageCubeArray != 0
	g.IndependentBlend = c.independentBlend != 0
	g.GeometryShader = c.geometryShader != 0
	g.TessellationShader = c.tessellationShader != 0
	g.SampleRateShading = c.sampleRateShading != 0
	g.DualSrcBlend = c.dualSrcBlend != 0
	g.LogicOp = c.logicOp != 0
	g.MultiDrawIndirect = c.multiDrawIndirect != 0
	g.DrawIndirectFirstInstance = c.drawIndirectFirstInstance != 0
	g.DepthClamp = c.depthClamp != 0
	g.DepthBiasClamp = c.depthBiasClamp != 0
	g.FillModeNonSolid = c.fillModeNonSolid != 0
	g.DepthBounds = c.depthBounds != 0
	g.WideLines = c.wideLines != 0
	g.LargePoints = c.largePoints != 0
	g.AlphaToOne = c.alphaToOne != 0
	g.MultiViewport = c.multiViewport != 0
	g.SamplerAnisotropy = c.samplerAnisotropy != 0
	g.TextureCompressionETC2 = c.textureCompressionETC2 != 0
	g.TextureCompressionASTC_LDR = c.textureCompressionASTC_LDR != 0
	g.TextureCompressionBC = c.textureCompressionBC != 0
	g.OcclusionQueryPrecise = c.occlusionQueryPrecise != 0
	g.PipelineStatisticsQuery = c.pipelineStatisticsQuery != 0
	g.VertexPipelineStoresAndAtomics = c.vertexPipelineStoresAndAtomics != 0
	g.FragmentStoresAndAtomics = c.fragmentStoresAndAtomics != 0
	g.ShaderTessellationAndGeometryPointSize = c.shaderTessellationAndGeometryPointSize != 0
	g.ShaderImageGatherExtended = c.shaderImageGatherExtended != 0
	g.ShaderStorageImageExtendedFormats = c.shaderStorageImageExtendedFormats != 0
	g.ShaderStorageImageMultisample = c.shaderStorageImageMultisample != 0
	g.ShaderStorageImageReadWithoutFormat = c.shaderStorageImageReadWithoutFormat != 0
	g.ShaderStorageImageWriteWithoutFormat = c.shaderStorageImageWriteWithoutFormat != 0
	g.ShaderUniformBufferArrayDynamicIndexing = c.shaderUniformBufferArrayDynamicIndexing != 0
	g.ShaderSampledImageArrayDynamicIndexing = c.shaderSampledImageArrayDynamicIndexing != 0
	g.ShaderStorageBufferArrayDynamicIndexing = c.shaderStorageBufferArrayDynamicIndexing != 0
	g.ShaderStorageImageArrayDynamicIndexing = c.shaderStorageImageArrayDynamicIndexing != 0
	g.ShaderClipDistance = c.shaderClipDistance != 0
	g.ShaderCullDistance = c.shaderCullDistance != 0
	g.ShaderFloat64 = c.shaderFloat64 != 0
	g.ShaderInt64 = c.shaderInt64 != 0
	g.ShaderInt16 = c.shaderInt16 != 0
	g.ShaderResourceResidency = c.shaderResourceResidency != 0
	g.ShaderResourceMinLod = c.shaderResourceMinLod != 0
	g.SparseBinding = c.sparseBinding != 0
	g.SparseResidencyBuffer = c.sparseResidencyBuffer != 0
	g.SparseResidencyImage2D = c.sparseResidencyImage2D != 0
	g.SparseResidencyImage3D = c.sparseResidencyImage3D != 0
	g.SparseResidency2Samples = c.sparseResidency2Samples != 0
	g.SparseResidency4Samples = c.sparseResidency4Samples != 0
	g.SparseResidency8Samples = c.sparseResidency8Samples != 0
	g.SparseResidency16Samples = c.sparseResidency16Samples != 0
	g.SparseResidencyAliased = c.sparseResidencyAliased != 0
	g.VariableMultisampleRate = c.variableMultisampleRate != 0
	g.InheritedQueries = c.inheritedQueries != 0
}
func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice, features []PhysicalDeviceFeatures) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFeatures      *C.VkPhysicalDeviceFeatures
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFeatures = (*C.VkPhysicalDeviceFeatures)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures * uint(len(features))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pFeatures))[:len(features):len(features)]
		for i3, _ := range features {
			features[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFeatures(c.physicalDevice, c.pFeatures)
}

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

type FormatFeatureFlags Flags
type FormatProperties struct {
	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}

func (g *FormatProperties) toC(c *C.VkFormatProperties) {
	{
		var temp_in_VkFormatFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.LinearTilingFeatures)))
			temp_in_VkFormatFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.linearTilingFeatures = C.VkFormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
	{
		var temp_in_VkFormatFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.OptimalTilingFeatures)))
			temp_in_VkFormatFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.optimalTilingFeatures = C.VkFormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
	{
		var temp_in_VkFormatFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.BufferFeatures)))
			temp_in_VkFormatFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.bufferFeatures = C.VkFormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
}
func (g *FormatProperties) fromC(c *C.VkFormatProperties) {
	{
		var temp_in_VkFormatFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.linearTilingFeatures)))
			temp_in_VkFormatFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.LinearTilingFeatures = FormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
	{
		var temp_in_VkFormatFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.optimalTilingFeatures)))
			temp_in_VkFormatFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.OptimalTilingFeatures = FormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
	{
		var temp_in_VkFormatFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.bufferFeatures)))
			temp_in_VkFormatFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.BufferFeatures = FormatFeatureFlags(temp_in_VkFormatFeatureFlags)
	}
}
func GetPhysicalDeviceFormatProperties(physicalDevice PhysicalDevice, format Format, formatProperties []FormatProperties) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		format            C.VkFormat
		pFormatProperties *C.VkFormatProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.format = C.VkFormat(format)
	{
		c.pFormatProperties = (*C.VkFormatProperties)(_sa.alloc(C.sizeof_VkFormatProperties * uint(len(formatProperties))))
		slice3 := (*[1 << 31]C.VkFormatProperties)(unsafe.Pointer(c.pFormatProperties))[:len(formatProperties):len(formatProperties)]
		for i3, _ := range formatProperties {
			formatProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFormatProperties(c.physicalDevice, c.format, c.pFormatProperties)
}

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

type ImageUsageFlags Flags
type ImageCreateFlags Flags
type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

func (g *Extent3D) toC(c *C.VkExtent3D) {
	c.width = C.uint32_t(g.Width)
	c.height = C.uint32_t(g.Height)
	c.depth = C.uint32_t(g.Depth)
}
func (g *Extent3D) fromC(c *C.VkExtent3D) {
	g.Width = uint32(c.width)
	g.Height = uint32(c.height)
	g.Depth = uint32(c.depth)
}

type SampleCountFlags Flags
type DeviceSize uint64
type ImageFormatProperties struct {
	MaxExtent       Extent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    SampleCountFlags
	MaxResourceSize DeviceSize
}

func (g *ImageFormatProperties) toC(c *C.VkImageFormatProperties) {
	g.MaxExtent.toC(&c.maxExtent)
	c.maxMipLevels = C.uint32_t(g.MaxMipLevels)
	c.maxArrayLayers = C.uint32_t(g.MaxArrayLayers)
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.sampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MaxResourceSize))
		c.maxResourceSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *ImageFormatProperties) fromC(c *C.VkImageFormatProperties) {
	g.MaxExtent.fromC(&c.maxExtent)
	g.MaxMipLevels = uint32(c.maxMipLevels)
	g.MaxArrayLayers = uint32(c.maxArrayLayers)
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.sampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.SampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.maxResourceSize))
		g.MaxResourceSize = DeviceSize(temp_in_VkDeviceSize)
	}
}
func GetPhysicalDeviceImageFormatProperties(physicalDevice PhysicalDevice, format Format, _type ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, imageFormatProperties []ImageFormatProperties) (_ret Result) {
	var c struct {
		physicalDevice         C.VkPhysicalDevice
		format                 C.VkFormat
		_type                  C.VkImageType
		tiling                 C.VkImageTiling
		usage                  C.VkImageUsageFlags
		flags                  C.VkImageCreateFlags
		pImageFormatProperties *C.VkImageFormatProperties
		_ret                   C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.format = C.VkFormat(format)
	c._type = C.VkImageType(_type)
	c.tiling = C.VkImageTiling(tiling)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(usage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	{
		var temp_in_VkImageCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkImageCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkImageCreateFlags(temp_in_VkImageCreateFlags)
	}
	{
		c.pImageFormatProperties = (*C.VkImageFormatProperties)(_sa.alloc(C.sizeof_VkImageFormatProperties * uint(len(imageFormatProperties))))
		slice3 := (*[1 << 31]C.VkImageFormatProperties)(unsafe.Pointer(c.pImageFormatProperties))[:len(imageFormatProperties):len(imageFormatProperties)]
		for i3, _ := range imageFormatProperties {
			imageFormatProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceImageFormatProperties(c.physicalDevice, c.format, c._type, c.tiling, c.usage, c.flags, c.pImageFormatProperties)
	_ret = Result(c._ret)
	return
}

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
	MaxMemoryAllocationCount                        uint32
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

func (g *PhysicalDeviceLimits) toC(c *C.VkPhysicalDeviceLimits) {
	c.maxImageDimension1D = C.uint32_t(g.MaxImageDimension1D)
	c.maxImageDimension2D = C.uint32_t(g.MaxImageDimension2D)
	c.maxImageDimension3D = C.uint32_t(g.MaxImageDimension3D)
	c.maxImageDimensionCube = C.uint32_t(g.MaxImageDimensionCube)
	c.maxImageArrayLayers = C.uint32_t(g.MaxImageArrayLayers)
	c.maxTexelBufferElements = C.uint32_t(g.MaxTexelBufferElements)
	c.maxUniformBufferRange = C.uint32_t(g.MaxUniformBufferRange)
	c.maxStorageBufferRange = C.uint32_t(g.MaxStorageBufferRange)
	c.maxPushConstantsSize = C.uint32_t(g.MaxPushConstantsSize)
	c.maxMemoryAllocationCount = C.uint32_t(g.MaxMemoryAllocationCount)
	c.maxSamplerAllocationCount = C.uint32_t(g.MaxSamplerAllocationCount)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.BufferImageGranularity))
		c.bufferImageGranularity = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.SparseAddressSpaceSize))
		c.sparseAddressSpaceSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.maxBoundDescriptorSets = C.uint32_t(g.MaxBoundDescriptorSets)
	c.maxPerStageDescriptorSamplers = C.uint32_t(g.MaxPerStageDescriptorSamplers)
	c.maxPerStageDescriptorUniformBuffers = C.uint32_t(g.MaxPerStageDescriptorUniformBuffers)
	c.maxPerStageDescriptorStorageBuffers = C.uint32_t(g.MaxPerStageDescriptorStorageBuffers)
	c.maxPerStageDescriptorSampledImages = C.uint32_t(g.MaxPerStageDescriptorSampledImages)
	c.maxPerStageDescriptorStorageImages = C.uint32_t(g.MaxPerStageDescriptorStorageImages)
	c.maxPerStageDescriptorInputAttachments = C.uint32_t(g.MaxPerStageDescriptorInputAttachments)
	c.maxPerStageResources = C.uint32_t(g.MaxPerStageResources)
	c.maxDescriptorSetSamplers = C.uint32_t(g.MaxDescriptorSetSamplers)
	c.maxDescriptorSetUniformBuffers = C.uint32_t(g.MaxDescriptorSetUniformBuffers)
	c.maxDescriptorSetUniformBuffersDynamic = C.uint32_t(g.MaxDescriptorSetUniformBuffersDynamic)
	c.maxDescriptorSetStorageBuffers = C.uint32_t(g.MaxDescriptorSetStorageBuffers)
	c.maxDescriptorSetStorageBuffersDynamic = C.uint32_t(g.MaxDescriptorSetStorageBuffersDynamic)
	c.maxDescriptorSetSampledImages = C.uint32_t(g.MaxDescriptorSetSampledImages)
	c.maxDescriptorSetStorageImages = C.uint32_t(g.MaxDescriptorSetStorageImages)
	c.maxDescriptorSetInputAttachments = C.uint32_t(g.MaxDescriptorSetInputAttachments)
	c.maxVertexInputAttributes = C.uint32_t(g.MaxVertexInputAttributes)
	c.maxVertexInputBindings = C.uint32_t(g.MaxVertexInputBindings)
	c.maxVertexInputAttributeOffset = C.uint32_t(g.MaxVertexInputAttributeOffset)
	c.maxVertexInputBindingStride = C.uint32_t(g.MaxVertexInputBindingStride)
	c.maxVertexOutputComponents = C.uint32_t(g.MaxVertexOutputComponents)
	c.maxTessellationGenerationLevel = C.uint32_t(g.MaxTessellationGenerationLevel)
	c.maxTessellationPatchSize = C.uint32_t(g.MaxTessellationPatchSize)
	c.maxTessellationControlPerVertexInputComponents = C.uint32_t(g.MaxTessellationControlPerVertexInputComponents)
	c.maxTessellationControlPerVertexOutputComponents = C.uint32_t(g.MaxTessellationControlPerVertexOutputComponents)
	c.maxTessellationControlPerPatchOutputComponents = C.uint32_t(g.MaxTessellationControlPerPatchOutputComponents)
	c.maxTessellationControlTotalOutputComponents = C.uint32_t(g.MaxTessellationControlTotalOutputComponents)
	c.maxTessellationEvaluationInputComponents = C.uint32_t(g.MaxTessellationEvaluationInputComponents)
	c.maxTessellationEvaluationOutputComponents = C.uint32_t(g.MaxTessellationEvaluationOutputComponents)
	c.maxGeometryShaderInvocations = C.uint32_t(g.MaxGeometryShaderInvocations)
	c.maxGeometryInputComponents = C.uint32_t(g.MaxGeometryInputComponents)
	c.maxGeometryOutputComponents = C.uint32_t(g.MaxGeometryOutputComponents)
	c.maxGeometryOutputVertices = C.uint32_t(g.MaxGeometryOutputVertices)
	c.maxGeometryTotalOutputComponents = C.uint32_t(g.MaxGeometryTotalOutputComponents)
	c.maxFragmentInputComponents = C.uint32_t(g.MaxFragmentInputComponents)
	c.maxFragmentOutputAttachments = C.uint32_t(g.MaxFragmentOutputAttachments)
	c.maxFragmentDualSrcAttachments = C.uint32_t(g.MaxFragmentDualSrcAttachments)
	c.maxFragmentCombinedOutputResources = C.uint32_t(g.MaxFragmentCombinedOutputResources)
	c.maxComputeSharedMemorySize = C.uint32_t(g.MaxComputeSharedMemorySize)
	for i, _ := range g.MaxComputeWorkGroupCount {
		c.maxComputeWorkGroupCount[i] = C.uint32_t(g.MaxComputeWorkGroupCount[i])
	}
	c.maxComputeWorkGroupInvocations = C.uint32_t(g.MaxComputeWorkGroupInvocations)
	for i, _ := range g.MaxComputeWorkGroupSize {
		c.maxComputeWorkGroupSize[i] = C.uint32_t(g.MaxComputeWorkGroupSize[i])
	}
	c.subPixelPrecisionBits = C.uint32_t(g.SubPixelPrecisionBits)
	c.subTexelPrecisionBits = C.uint32_t(g.SubTexelPrecisionBits)
	c.mipmapPrecisionBits = C.uint32_t(g.MipmapPrecisionBits)
	c.maxDrawIndexedIndexValue = C.uint32_t(g.MaxDrawIndexedIndexValue)
	c.maxDrawIndirectCount = C.uint32_t(g.MaxDrawIndirectCount)
	c.maxSamplerLodBias = C.float(g.MaxSamplerLodBias)
	c.maxSamplerAnisotropy = C.float(g.MaxSamplerAnisotropy)
	c.maxViewports = C.uint32_t(g.MaxViewports)
	for i, _ := range g.MaxViewportDimensions {
		c.maxViewportDimensions[i] = C.uint32_t(g.MaxViewportDimensions[i])
	}
	for i, _ := range g.ViewportBoundsRange {
		c.viewportBoundsRange[i] = C.float(g.ViewportBoundsRange[i])
	}
	c.viewportSubPixelBits = C.uint32_t(g.ViewportSubPixelBits)
	c.minMemoryMapAlignment = C.size_t(g.MinMemoryMapAlignment)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MinTexelBufferOffsetAlignment))
		c.minTexelBufferOffsetAlignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MinUniformBufferOffsetAlignment))
		c.minUniformBufferOffsetAlignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MinStorageBufferOffsetAlignment))
		c.minStorageBufferOffsetAlignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.minTexelOffset = C.int32_t(g.MinTexelOffset)
	c.maxTexelOffset = C.uint32_t(g.MaxTexelOffset)
	c.minTexelGatherOffset = C.int32_t(g.MinTexelGatherOffset)
	c.maxTexelGatherOffset = C.uint32_t(g.MaxTexelGatherOffset)
	c.minInterpolationOffset = C.float(g.MinInterpolationOffset)
	c.maxInterpolationOffset = C.float(g.MaxInterpolationOffset)
	c.subPixelInterpolationOffsetBits = C.uint32_t(g.SubPixelInterpolationOffsetBits)
	c.maxFramebufferWidth = C.uint32_t(g.MaxFramebufferWidth)
	c.maxFramebufferHeight = C.uint32_t(g.MaxFramebufferHeight)
	c.maxFramebufferLayers = C.uint32_t(g.MaxFramebufferLayers)
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.FramebufferColorSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.framebufferColorSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.FramebufferDepthSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.framebufferDepthSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.FramebufferStencilSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.framebufferStencilSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.FramebufferNoAttachmentsSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.framebufferNoAttachmentsSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	c.maxColorAttachments = C.uint32_t(g.MaxColorAttachments)
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SampledImageColorSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.sampledImageColorSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SampledImageIntegerSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.sampledImageIntegerSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SampledImageDepthSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.sampledImageDepthSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SampledImageStencilSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.sampledImageStencilSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.StorageImageSampleCounts)))
			temp_in_VkSampleCountFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.storageImageSampleCounts = C.VkSampleCountFlags(temp_in_VkSampleCountFlags)
	}
	c.maxSampleMaskWords = C.uint32_t(g.MaxSampleMaskWords)
	if g.TimestampComputeAndGraphics {
		c.timestampComputeAndGraphics = 1
	} else {
		c.timestampComputeAndGraphics = 0
	}
	c.timestampPeriod = C.float(g.TimestampPeriod)
	c.maxClipDistances = C.uint32_t(g.MaxClipDistances)
	c.maxCullDistances = C.uint32_t(g.MaxCullDistances)
	c.maxCombinedClipAndCullDistances = C.uint32_t(g.MaxCombinedClipAndCullDistances)
	c.discreteQueuePriorities = C.uint32_t(g.DiscreteQueuePriorities)
	for i, _ := range g.PointSizeRange {
		c.pointSizeRange[i] = C.float(g.PointSizeRange[i])
	}
	for i, _ := range g.LineWidthRange {
		c.lineWidthRange[i] = C.float(g.LineWidthRange[i])
	}
	c.pointSizeGranularity = C.float(g.PointSizeGranularity)
	c.lineWidthGranularity = C.float(g.LineWidthGranularity)
	if g.StrictLines {
		c.strictLines = 1
	} else {
		c.strictLines = 0
	}
	if g.StandardSampleLocations {
		c.standardSampleLocations = 1
	} else {
		c.standardSampleLocations = 0
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.OptimalBufferCopyOffsetAlignment))
		c.optimalBufferCopyOffsetAlignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.OptimalBufferCopyRowPitchAlignment))
		c.optimalBufferCopyRowPitchAlignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.NonCoherentAtomSize))
		c.nonCoherentAtomSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *PhysicalDeviceLimits) fromC(c *C.VkPhysicalDeviceLimits) {
	g.MaxImageDimension1D = uint32(c.maxImageDimension1D)
	g.MaxImageDimension2D = uint32(c.maxImageDimension2D)
	g.MaxImageDimension3D = uint32(c.maxImageDimension3D)
	g.MaxImageDimensionCube = uint32(c.maxImageDimensionCube)
	g.MaxImageArrayLayers = uint32(c.maxImageArrayLayers)
	g.MaxTexelBufferElements = uint32(c.maxTexelBufferElements)
	g.MaxUniformBufferRange = uint32(c.maxUniformBufferRange)
	g.MaxStorageBufferRange = uint32(c.maxStorageBufferRange)
	g.MaxPushConstantsSize = uint32(c.maxPushConstantsSize)
	g.MaxMemoryAllocationCount = uint32(c.maxMemoryAllocationCount)
	g.MaxSamplerAllocationCount = uint32(c.maxSamplerAllocationCount)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.bufferImageGranularity))
		g.BufferImageGranularity = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.sparseAddressSpaceSize))
		g.SparseAddressSpaceSize = DeviceSize(temp_in_VkDeviceSize)
	}
	g.MaxBoundDescriptorSets = uint32(c.maxBoundDescriptorSets)
	g.MaxPerStageDescriptorSamplers = uint32(c.maxPerStageDescriptorSamplers)
	g.MaxPerStageDescriptorUniformBuffers = uint32(c.maxPerStageDescriptorUniformBuffers)
	g.MaxPerStageDescriptorStorageBuffers = uint32(c.maxPerStageDescriptorStorageBuffers)
	g.MaxPerStageDescriptorSampledImages = uint32(c.maxPerStageDescriptorSampledImages)
	g.MaxPerStageDescriptorStorageImages = uint32(c.maxPerStageDescriptorStorageImages)
	g.MaxPerStageDescriptorInputAttachments = uint32(c.maxPerStageDescriptorInputAttachments)
	g.MaxPerStageResources = uint32(c.maxPerStageResources)
	g.MaxDescriptorSetSamplers = uint32(c.maxDescriptorSetSamplers)
	g.MaxDescriptorSetUniformBuffers = uint32(c.maxDescriptorSetUniformBuffers)
	g.MaxDescriptorSetUniformBuffersDynamic = uint32(c.maxDescriptorSetUniformBuffersDynamic)
	g.MaxDescriptorSetStorageBuffers = uint32(c.maxDescriptorSetStorageBuffers)
	g.MaxDescriptorSetStorageBuffersDynamic = uint32(c.maxDescriptorSetStorageBuffersDynamic)
	g.MaxDescriptorSetSampledImages = uint32(c.maxDescriptorSetSampledImages)
	g.MaxDescriptorSetStorageImages = uint32(c.maxDescriptorSetStorageImages)
	g.MaxDescriptorSetInputAttachments = uint32(c.maxDescriptorSetInputAttachments)
	g.MaxVertexInputAttributes = uint32(c.maxVertexInputAttributes)
	g.MaxVertexInputBindings = uint32(c.maxVertexInputBindings)
	g.MaxVertexInputAttributeOffset = uint32(c.maxVertexInputAttributeOffset)
	g.MaxVertexInputBindingStride = uint32(c.maxVertexInputBindingStride)
	g.MaxVertexOutputComponents = uint32(c.maxVertexOutputComponents)
	g.MaxTessellationGenerationLevel = uint32(c.maxTessellationGenerationLevel)
	g.MaxTessellationPatchSize = uint32(c.maxTessellationPatchSize)
	g.MaxTessellationControlPerVertexInputComponents = uint32(c.maxTessellationControlPerVertexInputComponents)
	g.MaxTessellationControlPerVertexOutputComponents = uint32(c.maxTessellationControlPerVertexOutputComponents)
	g.MaxTessellationControlPerPatchOutputComponents = uint32(c.maxTessellationControlPerPatchOutputComponents)
	g.MaxTessellationControlTotalOutputComponents = uint32(c.maxTessellationControlTotalOutputComponents)
	g.MaxTessellationEvaluationInputComponents = uint32(c.maxTessellationEvaluationInputComponents)
	g.MaxTessellationEvaluationOutputComponents = uint32(c.maxTessellationEvaluationOutputComponents)
	g.MaxGeometryShaderInvocations = uint32(c.maxGeometryShaderInvocations)
	g.MaxGeometryInputComponents = uint32(c.maxGeometryInputComponents)
	g.MaxGeometryOutputComponents = uint32(c.maxGeometryOutputComponents)
	g.MaxGeometryOutputVertices = uint32(c.maxGeometryOutputVertices)
	g.MaxGeometryTotalOutputComponents = uint32(c.maxGeometryTotalOutputComponents)
	g.MaxFragmentInputComponents = uint32(c.maxFragmentInputComponents)
	g.MaxFragmentOutputAttachments = uint32(c.maxFragmentOutputAttachments)
	g.MaxFragmentDualSrcAttachments = uint32(c.maxFragmentDualSrcAttachments)
	g.MaxFragmentCombinedOutputResources = uint32(c.maxFragmentCombinedOutputResources)
	g.MaxComputeSharedMemorySize = uint32(c.maxComputeSharedMemorySize)
	for i, _ := range g.MaxComputeWorkGroupCount {
		g.MaxComputeWorkGroupCount[i] = uint32(c.maxComputeWorkGroupCount[i])
	}
	g.MaxComputeWorkGroupInvocations = uint32(c.maxComputeWorkGroupInvocations)
	for i, _ := range g.MaxComputeWorkGroupSize {
		g.MaxComputeWorkGroupSize[i] = uint32(c.maxComputeWorkGroupSize[i])
	}
	g.SubPixelPrecisionBits = uint32(c.subPixelPrecisionBits)
	g.SubTexelPrecisionBits = uint32(c.subTexelPrecisionBits)
	g.MipmapPrecisionBits = uint32(c.mipmapPrecisionBits)
	g.MaxDrawIndexedIndexValue = uint32(c.maxDrawIndexedIndexValue)
	g.MaxDrawIndirectCount = uint32(c.maxDrawIndirectCount)
	g.MaxSamplerLodBias = float32(c.maxSamplerLodBias)
	g.MaxSamplerAnisotropy = float32(c.maxSamplerAnisotropy)
	g.MaxViewports = uint32(c.maxViewports)
	for i, _ := range g.MaxViewportDimensions {
		g.MaxViewportDimensions[i] = uint32(c.maxViewportDimensions[i])
	}
	for i, _ := range g.ViewportBoundsRange {
		g.ViewportBoundsRange[i] = float32(c.viewportBoundsRange[i])
	}
	g.ViewportSubPixelBits = uint32(c.viewportSubPixelBits)
	g.MinMemoryMapAlignment = uint(c.minMemoryMapAlignment)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.minTexelBufferOffsetAlignment))
		g.MinTexelBufferOffsetAlignment = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.minUniformBufferOffsetAlignment))
		g.MinUniformBufferOffsetAlignment = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.minStorageBufferOffsetAlignment))
		g.MinStorageBufferOffsetAlignment = DeviceSize(temp_in_VkDeviceSize)
	}
	g.MinTexelOffset = int32(c.minTexelOffset)
	g.MaxTexelOffset = uint32(c.maxTexelOffset)
	g.MinTexelGatherOffset = int32(c.minTexelGatherOffset)
	g.MaxTexelGatherOffset = uint32(c.maxTexelGatherOffset)
	g.MinInterpolationOffset = float32(c.minInterpolationOffset)
	g.MaxInterpolationOffset = float32(c.maxInterpolationOffset)
	g.SubPixelInterpolationOffsetBits = uint32(c.subPixelInterpolationOffsetBits)
	g.MaxFramebufferWidth = uint32(c.maxFramebufferWidth)
	g.MaxFramebufferHeight = uint32(c.maxFramebufferHeight)
	g.MaxFramebufferLayers = uint32(c.maxFramebufferLayers)
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.framebufferColorSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.FramebufferColorSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.framebufferDepthSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.FramebufferDepthSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.framebufferStencilSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.FramebufferStencilSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.framebufferNoAttachmentsSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.FramebufferNoAttachmentsSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	g.MaxColorAttachments = uint32(c.maxColorAttachments)
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageColorSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.SampledImageColorSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageIntegerSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.SampledImageIntegerSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageDepthSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.SampledImageDepthSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageStencilSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.SampledImageStencilSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	{
		var temp_in_VkSampleCountFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.storageImageSampleCounts)))
			temp_in_VkSampleCountFlags = Flags(temp_in_VkFlags)
		}
		g.StorageImageSampleCounts = SampleCountFlags(temp_in_VkSampleCountFlags)
	}
	g.MaxSampleMaskWords = uint32(c.maxSampleMaskWords)
	g.TimestampComputeAndGraphics = c.timestampComputeAndGraphics != 0
	g.TimestampPeriod = float32(c.timestampPeriod)
	g.MaxClipDistances = uint32(c.maxClipDistances)
	g.MaxCullDistances = uint32(c.maxCullDistances)
	g.MaxCombinedClipAndCullDistances = uint32(c.maxCombinedClipAndCullDistances)
	g.DiscreteQueuePriorities = uint32(c.discreteQueuePriorities)
	for i, _ := range g.PointSizeRange {
		g.PointSizeRange[i] = float32(c.pointSizeRange[i])
	}
	for i, _ := range g.LineWidthRange {
		g.LineWidthRange[i] = float32(c.lineWidthRange[i])
	}
	g.PointSizeGranularity = float32(c.pointSizeGranularity)
	g.LineWidthGranularity = float32(c.lineWidthGranularity)
	g.StrictLines = c.strictLines != 0
	g.StandardSampleLocations = c.standardSampleLocations != 0
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.optimalBufferCopyOffsetAlignment))
		g.OptimalBufferCopyOffsetAlignment = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.optimalBufferCopyRowPitchAlignment))
		g.OptimalBufferCopyRowPitchAlignment = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.nonCoherentAtomSize))
		g.NonCoherentAtomSize = DeviceSize(temp_in_VkDeviceSize)
	}
}

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

func (g *PhysicalDeviceSparseProperties) toC(c *C.VkPhysicalDeviceSparseProperties) {
	if g.ResidencyStandard2DBlockShape {
		c.residencyStandard2DBlockShape = 1
	} else {
		c.residencyStandard2DBlockShape = 0
	}
	if g.ResidencyStandard2DMultisampleBlockShape {
		c.residencyStandard2DMultisampleBlockShape = 1
	} else {
		c.residencyStandard2DMultisampleBlockShape = 0
	}
	if g.ResidencyStandard3DBlockShape {
		c.residencyStandard3DBlockShape = 1
	} else {
		c.residencyStandard3DBlockShape = 0
	}
	if g.ResidencyAlignedMipSize {
		c.residencyAlignedMipSize = 1
	} else {
		c.residencyAlignedMipSize = 0
	}
	if g.ResidencyNonResidentStrict {
		c.residencyNonResidentStrict = 1
	} else {
		c.residencyNonResidentStrict = 0
	}
}
func (g *PhysicalDeviceSparseProperties) fromC(c *C.VkPhysicalDeviceSparseProperties) {
	g.ResidencyStandard2DBlockShape = c.residencyStandard2DBlockShape != 0
	g.ResidencyStandard2DMultisampleBlockShape = c.residencyStandard2DMultisampleBlockShape != 0
	g.ResidencyStandard3DBlockShape = c.residencyStandard3DBlockShape != 0
	g.ResidencyAlignedMipSize = c.residencyAlignedMipSize != 0
	g.ResidencyNonResidentStrict = c.residencyNonResidentStrict != 0
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

func (g *PhysicalDeviceProperties) toC(c *C.VkPhysicalDeviceProperties) {
	c.apiVersion = C.uint32_t(g.ApiVersion)
	c.driverVersion = C.uint32_t(g.DriverVersion)
	c.vendorID = C.uint32_t(g.VendorID)
	c.deviceID = C.uint32_t(g.DeviceID)
	c.deviceType = C.VkPhysicalDeviceType(g.DeviceType)
	for i, _ := range g.DeviceName {
		c.deviceName[i] = C.char(g.DeviceName[i])
	}
	for i, _ := range g.PipelineCacheUUID {
		c.pipelineCacheUUID[i] = C.uint8_t(g.PipelineCacheUUID[i])
	}
	g.Limits.toC(&c.limits)
	g.SparseProperties.toC(&c.sparseProperties)
}
func (g *PhysicalDeviceProperties) fromC(c *C.VkPhysicalDeviceProperties) {
	g.ApiVersion = uint32(c.apiVersion)
	g.DriverVersion = uint32(c.driverVersion)
	g.VendorID = uint32(c.vendorID)
	g.DeviceID = uint32(c.deviceID)
	g.DeviceType = PhysicalDeviceType(c.deviceType)
	for i, _ := range g.DeviceName {
		g.DeviceName[i] = byte(c.deviceName[i])
	}
	for i, _ := range g.PipelineCacheUUID {
		g.PipelineCacheUUID[i] = uint8(c.pipelineCacheUUID[i])
	}
	g.Limits.fromC(&c.limits)
	g.SparseProperties.fromC(&c.sparseProperties)
}
func GetPhysicalDeviceProperties(physicalDevice PhysicalDevice, properties []PhysicalDeviceProperties) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pProperties    *C.VkPhysicalDeviceProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pProperties = (*C.VkPhysicalDeviceProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceProperties(c.physicalDevice, c.pProperties)
}

type QueueFlags Flags
type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

func (g *QueueFamilyProperties) toC(c *C.VkQueueFamilyProperties) {
	{
		var temp_in_VkQueueFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.QueueFlags)))
			temp_in_VkQueueFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.queueFlags = C.VkQueueFlags(temp_in_VkQueueFlags)
	}
	c.queueCount = C.uint32_t(g.QueueCount)
	c.timestampValidBits = C.uint32_t(g.TimestampValidBits)
	g.MinImageTransferGranularity.toC(&c.minImageTransferGranularity)
}
func (g *QueueFamilyProperties) fromC(c *C.VkQueueFamilyProperties) {
	{
		var temp_in_VkQueueFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.queueFlags)))
			temp_in_VkQueueFlags = Flags(temp_in_VkFlags)
		}
		g.QueueFlags = QueueFlags(temp_in_VkQueueFlags)
	}
	g.QueueCount = uint32(c.queueCount)
	g.TimestampValidBits = uint32(c.timestampValidBits)
	g.MinImageTransferGranularity.fromC(&c.minImageTransferGranularity)
}
func GetPhysicalDeviceQueueFamilyProperties(physicalDevice PhysicalDevice, queueFamilyPropertyCount *uint32, queueFamilyProperties []QueueFamilyProperties) {
	var c struct {
		physicalDevice            C.VkPhysicalDevice
		pQueueFamilyPropertyCount *C.uint32_t
		pQueueFamilyProperties    *C.VkQueueFamilyProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pQueueFamilyPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pQueueFamilyPropertyCount = C.uint32_t(*queueFamilyPropertyCount)
	}
	{
		c.pQueueFamilyProperties = (*C.VkQueueFamilyProperties)(_sa.alloc(C.sizeof_VkQueueFamilyProperties * uint(len(queueFamilyProperties))))
		slice3 := (*[1 << 31]C.VkQueueFamilyProperties)(unsafe.Pointer(c.pQueueFamilyProperties))[:len(queueFamilyProperties):len(queueFamilyProperties)]
		for i3, _ := range queueFamilyProperties {
			queueFamilyProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceQueueFamilyProperties(c.physicalDevice, c.pQueueFamilyPropertyCount, c.pQueueFamilyProperties)
	*queueFamilyPropertyCount = uint32(*c.pQueueFamilyPropertyCount)
}

type MemoryPropertyFlags Flags
type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}

func (g *MemoryType) toC(c *C.VkMemoryType) {
	{
		var temp_in_VkMemoryPropertyFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.PropertyFlags)))
			temp_in_VkMemoryPropertyFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.propertyFlags = C.VkMemoryPropertyFlags(temp_in_VkMemoryPropertyFlags)
	}
	c.heapIndex = C.uint32_t(g.HeapIndex)
}
func (g *MemoryType) fromC(c *C.VkMemoryType) {
	{
		var temp_in_VkMemoryPropertyFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.propertyFlags)))
			temp_in_VkMemoryPropertyFlags = Flags(temp_in_VkFlags)
		}
		g.PropertyFlags = MemoryPropertyFlags(temp_in_VkMemoryPropertyFlags)
	}
	g.HeapIndex = uint32(c.heapIndex)
}

type MemoryHeapFlags Flags
type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags
}

func (g *MemoryHeap) toC(c *C.VkMemoryHeap) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkMemoryHeapFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkMemoryHeapFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkMemoryHeapFlags(temp_in_VkMemoryHeapFlags)
	}
}
func (g *MemoryHeap) fromC(c *C.VkMemoryHeap) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkMemoryHeapFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkMemoryHeapFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = MemoryHeapFlags(temp_in_VkMemoryHeapFlags)
	}
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]MemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]MemoryHeap
}

func (g *PhysicalDeviceMemoryProperties) toC(c *C.VkPhysicalDeviceMemoryProperties) {
	c.memoryTypeCount = C.uint32_t(g.MemoryTypeCount)
	for i, _ := range g.MemoryTypes {
		g.MemoryTypes[i].toC(&c.memoryTypes[i])
	}
	c.memoryHeapCount = C.uint32_t(g.MemoryHeapCount)
	for i, _ := range g.MemoryHeaps {
		g.MemoryHeaps[i].toC(&c.memoryHeaps[i])
	}
}
func (g *PhysicalDeviceMemoryProperties) fromC(c *C.VkPhysicalDeviceMemoryProperties) {
	g.MemoryTypeCount = uint32(c.memoryTypeCount)
	for i, _ := range g.MemoryTypes {
		g.MemoryTypes[i].fromC(&c.memoryTypes[i])
	}
	g.MemoryHeapCount = uint32(c.memoryHeapCount)
	for i, _ := range g.MemoryHeaps {
		g.MemoryHeaps[i].fromC(&c.memoryHeaps[i])
	}
}
func GetPhysicalDeviceMemoryProperties(physicalDevice PhysicalDevice, memoryProperties []PhysicalDeviceMemoryProperties) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		pMemoryProperties *C.VkPhysicalDeviceMemoryProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pMemoryProperties = (*C.VkPhysicalDeviceMemoryProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceMemoryProperties * uint(len(memoryProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceMemoryProperties)(unsafe.Pointer(c.pMemoryProperties))[:len(memoryProperties):len(memoryProperties)]
		for i3, _ := range memoryProperties {
			memoryProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceMemoryProperties(c.physicalDevice, c.pMemoryProperties)
}

type PFN_vkVoidFunction struct {
	Raw C.PFN_vkVoidFunction
}

func (p PFN_vkVoidFunction) Call() {
	C.callPFN_vkVoidFunction(p.Raw)
}
func GetInstanceProcAddr(instance Instance, name string) (_ret PFN_vkVoidFunction) {
	var c struct {
		instance C.VkInstance
		pName    *C.char
		_ret     C.PFN_vkVoidFunction
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	c.pName = toCString(name, _sa)
	c._ret = C.vkGetInstanceProcAddr(c.instance, c.pName)
	_ret.Raw = c._ret
	return
}

type Device C.VkDevice

func GetDeviceProcAddr(device Device, name string) (_ret PFN_vkVoidFunction) {
	var c struct {
		device C.VkDevice
		pName  *C.char
		_ret   C.PFN_vkVoidFunction
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pName = toCString(name, _sa)
	c._ret = C.vkGetDeviceProcAddr(c.device, c.pName)
	_ret.Raw = c._ret
	return
}

type DeviceCreateFlags Flags
type DeviceQueueCreateFlags Flags
type DeviceQueueCreateInfo struct {
	Type             StructureType
	Next             unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

func (g *DeviceQueueCreateInfo) toC(c *C.VkDeviceQueueCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDeviceQueueCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDeviceQueueCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDeviceQueueCreateFlags(temp_in_VkDeviceQueueCreateFlags)
	}
	c.queueFamilyIndex = C.uint32_t(g.QueueFamilyIndex)
	c.queueCount = C.uint32_t(len(g.QueuePriorities))
	{
		c.pQueuePriorities = (*C.float)(_sa.alloc(C.sizeof_float * uint(len(g.QueuePriorities))))
		slice2 := (*[1 << 31]C.float)(unsafe.Pointer(c.pQueuePriorities))[:len(g.QueuePriorities):len(g.QueuePriorities)]
		for i2, _ := range g.QueuePriorities {
			slice2[i2] = C.float(g.QueuePriorities[i2])
		}
	}
}
func (g *DeviceQueueCreateInfo) fromC(c *C.VkDeviceQueueCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDeviceQueueCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDeviceQueueCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DeviceQueueCreateFlags(temp_in_VkDeviceQueueCreateFlags)
	}
	g.QueueFamilyIndex = uint32(c.queueFamilyIndex)
	g.QueuePriorities = make([]float32, int(c.queueCount))
	{
		slice2 := (*[1 << 31]C.float)(unsafe.Pointer(c.pQueuePriorities))[:len(g.QueuePriorities):len(g.QueuePriorities)]
		for i2, _ := range g.QueuePriorities {
			g.QueuePriorities[i2] = float32(slice2[i2])
		}
	}
}

type DeviceCreateInfo struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	Flags                 DeviceCreateFlags
	QueueCreateInfos      []DeviceQueueCreateInfo
	EnabledLayerNames     []string
	EnabledExtensionNames []string
	EnabledFeatures       []PhysicalDeviceFeatures
}

func (g *DeviceCreateInfo) toC(c *C.VkDeviceCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDeviceCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDeviceCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDeviceCreateFlags(temp_in_VkDeviceCreateFlags)
	}
	c.queueCreateInfoCount = C.uint32_t(len(g.QueueCreateInfos))
	{
		c.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(_sa.alloc(C.sizeof_VkDeviceQueueCreateInfo * uint(len(g.QueueCreateInfos))))
		slice2 := (*[1 << 31]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(c.pQueueCreateInfos))[:len(g.QueueCreateInfos):len(g.QueueCreateInfos)]
		for i2, _ := range g.QueueCreateInfos {
			g.QueueCreateInfos[i2].toC(&slice2[i2], _sa)
		}
	}
	c.enabledLayerCount = C.uint32_t(len(g.EnabledLayerNames))
	{
		c.ppEnabledLayerNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.EnabledLayerNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.EnabledLayerNames):len(g.EnabledLayerNames)]
		for i2, _ := range g.EnabledLayerNames {
			slice2[i2] = toCString(g.EnabledLayerNames[i2], _sa)
		}
	}
	c.enabledExtensionCount = C.uint32_t(len(g.EnabledExtensionNames))
	{
		c.ppEnabledExtensionNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.EnabledExtensionNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.EnabledExtensionNames):len(g.EnabledExtensionNames)]
		for i2, _ := range g.EnabledExtensionNames {
			slice2[i2] = toCString(g.EnabledExtensionNames[i2], _sa)
		}
	}
	{
		c.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures * uint(len(g.EnabledFeatures))))
		slice2 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pEnabledFeatures))[:len(g.EnabledFeatures):len(g.EnabledFeatures)]
		for i2, _ := range g.EnabledFeatures {
			g.EnabledFeatures[i2].toC(&slice2[i2])
		}
	}
}
func (g *DeviceCreateInfo) fromC(c *C.VkDeviceCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDeviceCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDeviceCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DeviceCreateFlags(temp_in_VkDeviceCreateFlags)
	}
	g.QueueCreateInfos = make([]DeviceQueueCreateInfo, int(c.queueCreateInfoCount))
	{
		slice2 := (*[1 << 31]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(c.pQueueCreateInfos))[:len(g.QueueCreateInfos):len(g.QueueCreateInfos)]
		for i2, _ := range g.QueueCreateInfos {
			g.QueueCreateInfos[i2].fromC(&slice2[i2])
		}
	}
	g.EnabledLayerNames = make([]string, int(c.enabledLayerCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.EnabledLayerNames):len(g.EnabledLayerNames)]
		for i2, _ := range g.EnabledLayerNames {
			g.EnabledLayerNames[i2] = toGoString(slice2[i2])
		}
	}
	g.EnabledExtensionNames = make([]string, int(c.enabledExtensionCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.EnabledExtensionNames):len(g.EnabledExtensionNames)]
		for i2, _ := range g.EnabledExtensionNames {
			g.EnabledExtensionNames[i2] = toGoString(slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pEnabledFeatures))[:len(g.EnabledFeatures):len(g.EnabledFeatures)]
		for i2, _ := range g.EnabledFeatures {
			g.EnabledFeatures[i2].fromC(&slice2[i2])
		}
	}
}
func CreateDevice(physicalDevice PhysicalDevice, createInfo *DeviceCreateInfo, allocator *AllocationCallbacks, device *Device) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pCreateInfo    *C.VkDeviceCreateInfo
		pAllocator     *C.VkAllocationCallbacks
		pDevice        *C.VkDevice
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pCreateInfo = (*C.VkDeviceCreateInfo)(_sa.alloc(C.sizeof_VkDeviceCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDevice = (*C.VkDevice)(_sa.alloc(C.sizeof_VkDevice))
		*c.pDevice = C.VkDevice(*device)
	}
	c._ret = C.vkCreateDevice(c.physicalDevice, c.pCreateInfo, c.pAllocator, c.pDevice)
	_ret = Result(c._ret)
	*device = Device(*c.pDevice)
	return
}
func DestroyDevice(device Device, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDevice(c.device, c.pAllocator)
}

type ExtensionProperties struct {
	ExtensionName [256]byte
	SpecVersion   uint32
}

func (g *ExtensionProperties) toC(c *C.VkExtensionProperties) {
	for i, _ := range g.ExtensionName {
		c.extensionName[i] = C.char(g.ExtensionName[i])
	}
	c.specVersion = C.uint32_t(g.SpecVersion)
}
func (g *ExtensionProperties) fromC(c *C.VkExtensionProperties) {
	for i, _ := range g.ExtensionName {
		g.ExtensionName[i] = byte(c.extensionName[i])
	}
	g.SpecVersion = uint32(c.specVersion)
}
func EnumerateInstanceExtensionProperties(layerName string, propertyCount *uint32, properties []ExtensionProperties) (_ret Result) {
	var c struct {
		pLayerName     *C.char
		pPropertyCount *C.uint32_t
		pProperties    *C.VkExtensionProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.pLayerName = toCString(layerName, _sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkExtensionProperties)(_sa.alloc(C.sizeof_VkExtensionProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkExtensionProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateInstanceExtensionProperties(c.pLayerName, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}
func EnumerateDeviceExtensionProperties(physicalDevice PhysicalDevice, layerName string, propertyCount *uint32, properties []ExtensionProperties) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pLayerName     *C.char
		pPropertyCount *C.uint32_t
		pProperties    *C.VkExtensionProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.pLayerName = toCString(layerName, _sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkExtensionProperties)(_sa.alloc(C.sizeof_VkExtensionProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkExtensionProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateDeviceExtensionProperties(c.physicalDevice, c.pLayerName, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type LayerProperties struct {
	LayerName             [256]byte
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [256]byte
}

func (g *LayerProperties) toC(c *C.VkLayerProperties) {
	for i, _ := range g.LayerName {
		c.layerName[i] = C.char(g.LayerName[i])
	}
	c.specVersion = C.uint32_t(g.SpecVersion)
	c.implementationVersion = C.uint32_t(g.ImplementationVersion)
	for i, _ := range g.Description {
		c.description[i] = C.char(g.Description[i])
	}
}
func (g *LayerProperties) fromC(c *C.VkLayerProperties) {
	for i, _ := range g.LayerName {
		g.LayerName[i] = byte(c.layerName[i])
	}
	g.SpecVersion = uint32(c.specVersion)
	g.ImplementationVersion = uint32(c.implementationVersion)
	for i, _ := range g.Description {
		g.Description[i] = byte(c.description[i])
	}
}
func EnumerateInstanceLayerProperties(propertyCount *uint32, properties []LayerProperties) (_ret Result) {
	var c struct {
		pPropertyCount *C.uint32_t
		pProperties    *C.VkLayerProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkLayerProperties)(_sa.alloc(C.sizeof_VkLayerProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkLayerProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateInstanceLayerProperties(c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}
func EnumerateDeviceLayerProperties(physicalDevice PhysicalDevice, propertyCount *uint32, properties []LayerProperties) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pPropertyCount *C.uint32_t
		pProperties    *C.VkLayerProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkLayerProperties)(_sa.alloc(C.sizeof_VkLayerProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkLayerProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateDeviceLayerProperties(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type Queue C.VkQueue

func GetDeviceQueue(device Device, queueFamilyIndex uint32, queueIndex uint32, queue *Queue) {
	var c struct {
		device           C.VkDevice
		queueFamilyIndex C.uint32_t
		queueIndex       C.uint32_t
		pQueue           *C.VkQueue
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.queueFamilyIndex = C.uint32_t(queueFamilyIndex)
	c.queueIndex = C.uint32_t(queueIndex)
	{
		c.pQueue = (*C.VkQueue)(_sa.alloc(C.sizeof_VkQueue))
		*c.pQueue = C.VkQueue(*queue)
	}
	C.vkGetDeviceQueue(c.device, c.queueFamilyIndex, c.queueIndex, c.pQueue)
	*queue = Queue(*c.pQueue)
}

type Semaphore C.VkSemaphore
type PipelineStageFlags Flags
type CommandBuffer C.VkCommandBuffer
type SubmitInfo struct {
	Type             StructureType
	Next             unsafe.Pointer
	WaitSemaphores   []Semaphore
	WaitDstStageMask *PipelineStageFlags
	CommandBuffers   []CommandBuffer
	SignalSemaphores []Semaphore
}

func (g *SubmitInfo) toC(c *C.VkSubmitInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.waitSemaphoreCount = C.uint32_t(len(g.WaitSemaphores))
	{
		c.pWaitSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.WaitSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			slice2[i2] = C.VkSemaphore(g.WaitSemaphores[i2])
		}
	}
	{
		c.pWaitDstStageMask = (*C.VkPipelineStageFlags)(_sa.alloc(C.sizeof_VkPipelineStageFlags))
		{
			var temp_in_VkPipelineStageFlags C.VkFlags
			{
				var temp_in_VkFlags C.uint32_t
				temp_in_VkFlags = C.uint32_t((uint32)((Flags)(*g.WaitDstStageMask)))
				temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
			}
			*c.pWaitDstStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
		}
	}
	c.commandBufferCount = C.uint32_t(len(g.CommandBuffers))
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(g.CommandBuffers))))
		slice2 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(g.CommandBuffers):len(g.CommandBuffers)]
		for i2, _ := range g.CommandBuffers {
			slice2[i2] = C.VkCommandBuffer(g.CommandBuffers[i2])
		}
	}
	c.signalSemaphoreCount = C.uint32_t(len(g.SignalSemaphores))
	{
		c.pSignalSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.SignalSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.SignalSemaphores):len(g.SignalSemaphores)]
		for i2, _ := range g.SignalSemaphores {
			slice2[i2] = C.VkSemaphore(g.SignalSemaphores[i2])
		}
	}
}
func (g *SubmitInfo) fromC(c *C.VkSubmitInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.WaitSemaphores = make([]Semaphore, int(c.waitSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			g.WaitSemaphores[i2] = Semaphore(slice2[i2])
		}
	}
	{
		if g.WaitDstStageMask == nil {
			g.WaitDstStageMask = new(PipelineStageFlags)
		}
		{
			var temp_in_VkPipelineStageFlags Flags
			{
				var temp_in_VkFlags uint32
				temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(*c.pWaitDstStageMask)))
				temp_in_VkPipelineStageFlags = Flags(temp_in_VkFlags)
			}
			*g.WaitDstStageMask = PipelineStageFlags(temp_in_VkPipelineStageFlags)
		}
	}
	g.CommandBuffers = make([]CommandBuffer, int(c.commandBufferCount))
	{
		slice2 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(g.CommandBuffers):len(g.CommandBuffers)]
		for i2, _ := range g.CommandBuffers {
			g.CommandBuffers[i2] = CommandBuffer(slice2[i2])
		}
	}
	g.SignalSemaphores = make([]Semaphore, int(c.signalSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.SignalSemaphores):len(g.SignalSemaphores)]
		for i2, _ := range g.SignalSemaphores {
			g.SignalSemaphores[i2] = Semaphore(slice2[i2])
		}
	}
}

type Fence C.VkFence

func QueueSubmit(queue Queue, submits []SubmitInfo, fence Fence) (_ret Result) {
	var c struct {
		queue       C.VkQueue
		submitCount C.uint32_t
		pSubmits    *C.VkSubmitInfo
		fence       C.VkFence
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.queue = C.VkQueue(queue)
	c.submitCount = C.uint32_t(len(submits))
	{
		c.pSubmits = (*C.VkSubmitInfo)(_sa.alloc(C.sizeof_VkSubmitInfo * uint(len(submits))))
		slice3 := (*[1 << 31]C.VkSubmitInfo)(unsafe.Pointer(c.pSubmits))[:len(submits):len(submits)]
		for i3, _ := range submits {
			submits[i3].toC(&slice3[i3], _sa)
		}
	}
	c.fence = C.VkFence(fence)
	c._ret = C.vkQueueSubmit(c.queue, c.submitCount, c.pSubmits, c.fence)
	_ret = Result(c._ret)
	return
}
func QueueWaitIdle(queue Queue) (_ret Result) {
	var c struct {
		queue C.VkQueue
		_ret  C.VkResult
	}
	c.queue = C.VkQueue(queue)
	c._ret = C.vkQueueWaitIdle(c.queue)
	_ret = Result(c._ret)
	return
}
func DeviceWaitIdle(device Device) (_ret Result) {
	var c struct {
		device C.VkDevice
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c._ret = C.vkDeviceWaitIdle(c.device)
	_ret = Result(c._ret)
	return
}

type MemoryAllocateInfo struct {
	Type            StructureType
	Next            unsafe.Pointer
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

func (g *MemoryAllocateInfo) toC(c *C.VkMemoryAllocateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.AllocationSize))
		c.allocationSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.memoryTypeIndex = C.uint32_t(g.MemoryTypeIndex)
}
func (g *MemoryAllocateInfo) fromC(c *C.VkMemoryAllocateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.allocationSize))
		g.AllocationSize = DeviceSize(temp_in_VkDeviceSize)
	}
	g.MemoryTypeIndex = uint32(c.memoryTypeIndex)
}

type DeviceMemory C.VkDeviceMemory

func AllocateMemory(device Device, allocateInfo *MemoryAllocateInfo, allocator *AllocationCallbacks, memory *DeviceMemory) (_ret Result) {
	var c struct {
		device        C.VkDevice
		pAllocateInfo *C.VkMemoryAllocateInfo
		pAllocator    *C.VkAllocationCallbacks
		pMemory       *C.VkDeviceMemory
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAllocateInfo = (*C.VkMemoryAllocateInfo)(_sa.alloc(C.sizeof_VkMemoryAllocateInfo))
		allocateInfo.toC(c.pAllocateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pMemory = (*C.VkDeviceMemory)(_sa.alloc(C.sizeof_VkDeviceMemory))
		*c.pMemory = C.VkDeviceMemory(*memory)
	}
	c._ret = C.vkAllocateMemory(c.device, c.pAllocateInfo, c.pAllocator, c.pMemory)
	_ret = Result(c._ret)
	*memory = DeviceMemory(*c.pMemory)
	return
}
func FreeMemory(device Device, memory DeviceMemory, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		memory     C.VkDeviceMemory
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memory = C.VkDeviceMemory(memory)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkFreeMemory(c.device, c.memory, c.pAllocator)
}

type MemoryMapFlags Flags

func MapMemory(device Device, memory DeviceMemory, offset DeviceSize, size DeviceSize, flags MemoryMapFlags, data []unsafe.Pointer) (_ret Result) {
	var c struct {
		device C.VkDevice
		memory C.VkDeviceMemory
		offset C.VkDeviceSize
		size   C.VkDeviceSize
		flags  C.VkMemoryMapFlags
		ppData *unsafe.Pointer
		_ret   C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memory = C.VkDeviceMemory(memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkMemoryMapFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkMemoryMapFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkMemoryMapFlags(temp_in_VkMemoryMapFlags)
	}
	{
		c.ppData = (*unsafe.Pointer)(_sa.alloc(C.sizeof_void_pointer * uint(len(data))))
		slice3 := (*[1 << 31]unsafe.Pointer)(unsafe.Pointer(c.ppData))[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	c._ret = C.vkMapMemory(c.device, c.memory, c.offset, c.size, c.flags, c.ppData)
	_ret = Result(c._ret)
	return
}
func UnmapMemory(device Device, memory DeviceMemory) {
	var c struct {
		device C.VkDevice
		memory C.VkDeviceMemory
	}
	c.device = C.VkDevice(device)
	c.memory = C.VkDeviceMemory(memory)
	C.vkUnmapMemory(c.device, c.memory)
}

type MappedMemoryRange struct {
	Type   StructureType
	Next   unsafe.Pointer
	Memory DeviceMemory
	Offset DeviceSize
	Size   DeviceSize
}

func (g *MappedMemoryRange) toC(c *C.VkMappedMemoryRange) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.memory = C.VkDeviceMemory(g.Memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *MappedMemoryRange) fromC(c *C.VkMappedMemoryRange) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Memory = DeviceMemory(c.memory)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.offset))
		g.Offset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
}
func FlushMappedMemoryRanges(device Device, memoryRanges []MappedMemoryRange) (_ret Result) {
	var c struct {
		device           C.VkDevice
		memoryRangeCount C.uint32_t
		pMemoryRanges    *C.VkMappedMemoryRange
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memoryRangeCount = C.uint32_t(len(memoryRanges))
	{
		c.pMemoryRanges = (*C.VkMappedMemoryRange)(_sa.alloc(C.sizeof_VkMappedMemoryRange * uint(len(memoryRanges))))
		slice3 := (*[1 << 31]C.VkMappedMemoryRange)(unsafe.Pointer(c.pMemoryRanges))[:len(memoryRanges):len(memoryRanges)]
		for i3, _ := range memoryRanges {
			memoryRanges[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkFlushMappedMemoryRanges(c.device, c.memoryRangeCount, c.pMemoryRanges)
	_ret = Result(c._ret)
	return
}
func InvalidateMappedMemoryRanges(device Device, memoryRanges []MappedMemoryRange) (_ret Result) {
	var c struct {
		device           C.VkDevice
		memoryRangeCount C.uint32_t
		pMemoryRanges    *C.VkMappedMemoryRange
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memoryRangeCount = C.uint32_t(len(memoryRanges))
	{
		c.pMemoryRanges = (*C.VkMappedMemoryRange)(_sa.alloc(C.sizeof_VkMappedMemoryRange * uint(len(memoryRanges))))
		slice3 := (*[1 << 31]C.VkMappedMemoryRange)(unsafe.Pointer(c.pMemoryRanges))[:len(memoryRanges):len(memoryRanges)]
		for i3, _ := range memoryRanges {
			memoryRanges[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkInvalidateMappedMemoryRanges(c.device, c.memoryRangeCount, c.pMemoryRanges)
	_ret = Result(c._ret)
	return
}
func GetDeviceMemoryCommitment(device Device, memory DeviceMemory, committedMemoryInBytes []DeviceSize) {
	var c struct {
		device                  C.VkDevice
		memory                  C.VkDeviceMemory
		pCommittedMemoryInBytes *C.VkDeviceSize
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memory = C.VkDeviceMemory(memory)
	{
		c.pCommittedMemoryInBytes = (*C.VkDeviceSize)(_sa.alloc(C.sizeof_VkDeviceSize * uint(len(committedMemoryInBytes))))
		slice3 := (*[1 << 31]C.VkDeviceSize)(unsafe.Pointer(c.pCommittedMemoryInBytes))[:len(committedMemoryInBytes):len(committedMemoryInBytes)]
		for i3, _ := range committedMemoryInBytes {
			{
				var temp_in_VkDeviceSize C.uint64_t
				temp_in_VkDeviceSize = C.uint64_t((uint64)(committedMemoryInBytes[i3]))
				slice3[i3] = C.VkDeviceSize(temp_in_VkDeviceSize)
			}
		}
	}
	C.vkGetDeviceMemoryCommitment(c.device, c.memory, c.pCommittedMemoryInBytes)
}

type Buffer C.VkBuffer

func BindBufferMemory(device Device, buffer Buffer, memory DeviceMemory, memoryOffset DeviceSize) (_ret Result) {
	var c struct {
		device       C.VkDevice
		buffer       C.VkBuffer
		memory       C.VkDeviceMemory
		memoryOffset C.VkDeviceSize
		_ret         C.VkResult
	}
	c.device = C.VkDevice(device)
	c.buffer = C.VkBuffer(buffer)
	c.memory = C.VkDeviceMemory(memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(memoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c._ret = C.vkBindBufferMemory(c.device, c.buffer, c.memory, c.memoryOffset)
	_ret = Result(c._ret)
	return
}

type Image C.VkImage

func BindImageMemory(device Device, image Image, memory DeviceMemory, memoryOffset DeviceSize) (_ret Result) {
	var c struct {
		device       C.VkDevice
		image        C.VkImage
		memory       C.VkDeviceMemory
		memoryOffset C.VkDeviceSize
		_ret         C.VkResult
	}
	c.device = C.VkDevice(device)
	c.image = C.VkImage(image)
	c.memory = C.VkDeviceMemory(memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(memoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c._ret = C.vkBindImageMemory(c.device, c.image, c.memory, c.memoryOffset)
	_ret = Result(c._ret)
	return
}

type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}

func (g *MemoryRequirements) toC(c *C.VkMemoryRequirements) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Alignment))
		c.alignment = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.memoryTypeBits = C.uint32_t(g.MemoryTypeBits)
}
func (g *MemoryRequirements) fromC(c *C.VkMemoryRequirements) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.alignment))
		g.Alignment = DeviceSize(temp_in_VkDeviceSize)
	}
	g.MemoryTypeBits = uint32(c.memoryTypeBits)
}
func GetBufferMemoryRequirements(device Device, buffer Buffer, memoryRequirements []MemoryRequirements) {
	var c struct {
		device              C.VkDevice
		buffer              C.VkBuffer
		pMemoryRequirements *C.VkMemoryRequirements
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.buffer = C.VkBuffer(buffer)
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements)(_sa.alloc(C.sizeof_VkMemoryRequirements * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetBufferMemoryRequirements(c.device, c.buffer, c.pMemoryRequirements)
}
func GetImageMemoryRequirements(device Device, image Image, memoryRequirements []MemoryRequirements) {
	var c struct {
		device              C.VkDevice
		image               C.VkImage
		pMemoryRequirements *C.VkMemoryRequirements
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.image = C.VkImage(image)
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements)(_sa.alloc(C.sizeof_VkMemoryRequirements * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageMemoryRequirements(c.device, c.image, c.pMemoryRequirements)
}

type ImageAspectFlags Flags
type SparseImageFormatFlags Flags
type SparseImageFormatProperties struct {
	AspectMask       ImageAspectFlags
	ImageGranularity Extent3D
	Flags            SparseImageFormatFlags
}

func (g *SparseImageFormatProperties) toC(c *C.VkSparseImageFormatProperties) {
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.ImageGranularity.toC(&c.imageGranularity)
	{
		var temp_in_VkSparseImageFormatFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSparseImageFormatFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSparseImageFormatFlags(temp_in_VkSparseImageFormatFlags)
	}
}
func (g *SparseImageFormatProperties) fromC(c *C.VkSparseImageFormatProperties) {
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.ImageGranularity.fromC(&c.imageGranularity)
	{
		var temp_in_VkSparseImageFormatFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSparseImageFormatFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SparseImageFormatFlags(temp_in_VkSparseImageFormatFlags)
	}
}

type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     DeviceSize
	ImageMipTailOffset   DeviceSize
	ImageMipTailStride   DeviceSize
}

func (g *SparseImageMemoryRequirements) toC(c *C.VkSparseImageMemoryRequirements) {
	g.FormatProperties.toC(&c.formatProperties)
	c.imageMipTailFirstLod = C.uint32_t(g.ImageMipTailFirstLod)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.ImageMipTailSize))
		c.imageMipTailSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.ImageMipTailOffset))
		c.imageMipTailOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.ImageMipTailStride))
		c.imageMipTailStride = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *SparseImageMemoryRequirements) fromC(c *C.VkSparseImageMemoryRequirements) {
	g.FormatProperties.fromC(&c.formatProperties)
	g.ImageMipTailFirstLod = uint32(c.imageMipTailFirstLod)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.imageMipTailSize))
		g.ImageMipTailSize = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.imageMipTailOffset))
		g.ImageMipTailOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.imageMipTailStride))
		g.ImageMipTailStride = DeviceSize(temp_in_VkDeviceSize)
	}
}
func GetImageSparseMemoryRequirements(device Device, image Image, sparseMemoryRequirementCount *uint32, sparseMemoryRequirements []SparseImageMemoryRequirements) {
	var c struct {
		device                        C.VkDevice
		image                         C.VkImage
		pSparseMemoryRequirementCount *C.uint32_t
		pSparseMemoryRequirements     *C.VkSparseImageMemoryRequirements
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.image = C.VkImage(image)
	{
		c.pSparseMemoryRequirementCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSparseMemoryRequirementCount = C.uint32_t(*sparseMemoryRequirementCount)
	}
	{
		c.pSparseMemoryRequirements = (*C.VkSparseImageMemoryRequirements)(_sa.alloc(C.sizeof_VkSparseImageMemoryRequirements * uint(len(sparseMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkSparseImageMemoryRequirements)(unsafe.Pointer(c.pSparseMemoryRequirements))[:len(sparseMemoryRequirements):len(sparseMemoryRequirements)]
		for i3, _ := range sparseMemoryRequirements {
			sparseMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageSparseMemoryRequirements(c.device, c.image, c.pSparseMemoryRequirementCount, c.pSparseMemoryRequirements)
	*sparseMemoryRequirementCount = uint32(*c.pSparseMemoryRequirementCount)
}

type SampleCountFlagBits int

const (
	SAMPLE_COUNT_1_BIT              SampleCountFlagBits = 1
	SAMPLE_COUNT_2_BIT              SampleCountFlagBits = 2
	SAMPLE_COUNT_4_BIT              SampleCountFlagBits = 4
	SAMPLE_COUNT_8_BIT              SampleCountFlagBits = 8
	SAMPLE_COUNT_16_BIT             SampleCountFlagBits = 16
	SAMPLE_COUNT_32_BIT             SampleCountFlagBits = 32
	SAMPLE_COUNT_64_BIT             SampleCountFlagBits = 64
	SAMPLE_COUNT_FLAG_BITS_MAX_ENUM SampleCountFlagBits = 2147483647
)

func GetPhysicalDeviceSparseImageFormatProperties(physicalDevice PhysicalDevice, format Format, _type ImageType, samples SampleCountFlagBits, usage ImageUsageFlags, tiling ImageTiling, propertyCount *uint32, properties []SparseImageFormatProperties) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		format         C.VkFormat
		_type          C.VkImageType
		samples        C.VkSampleCountFlagBits
		usage          C.VkImageUsageFlags
		tiling         C.VkImageTiling
		pPropertyCount *C.uint32_t
		pProperties    *C.VkSparseImageFormatProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.format = C.VkFormat(format)
	c._type = C.VkImageType(_type)
	c.samples = C.VkSampleCountFlagBits(samples)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(usage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	c.tiling = C.VkImageTiling(tiling)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkSparseImageFormatProperties)(_sa.alloc(C.sizeof_VkSparseImageFormatProperties * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkSparseImageFormatProperties)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceSparseImageFormatProperties(c.physicalDevice, c.format, c._type, c.samples, c.usage, c.tiling, c.pPropertyCount, c.pProperties)
	*propertyCount = uint32(*c.pPropertyCount)
}

type SparseMemoryBindFlags Flags
type SparseMemoryBind struct {
	ResourceOffset DeviceSize
	Size           DeviceSize
	Memory         DeviceMemory
	MemoryOffset   DeviceSize
	Flags          SparseMemoryBindFlags
}

func (g *SparseMemoryBind) toC(c *C.VkSparseMemoryBind) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.ResourceOffset))
		c.resourceOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.memory = C.VkDeviceMemory(g.Memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MemoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkSparseMemoryBindFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSparseMemoryBindFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSparseMemoryBindFlags(temp_in_VkSparseMemoryBindFlags)
	}
}
func (g *SparseMemoryBind) fromC(c *C.VkSparseMemoryBind) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.resourceOffset))
		g.ResourceOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
	g.Memory = DeviceMemory(c.memory)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.memoryOffset))
		g.MemoryOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkSparseMemoryBindFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSparseMemoryBindFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SparseMemoryBindFlags(temp_in_VkSparseMemoryBindFlags)
	}
}

type SparseBufferMemoryBindInfo struct {
	Buffer Buffer
	Binds  []SparseMemoryBind
}

func (g *SparseBufferMemoryBindInfo) toC(c *C.VkSparseBufferMemoryBindInfo, _sa *stackAllocator) {
	c.buffer = C.VkBuffer(g.Buffer)
	c.bindCount = C.uint32_t(len(g.Binds))
	{
		c.pBinds = (*C.VkSparseMemoryBind)(_sa.alloc(C.sizeof_VkSparseMemoryBind * uint(len(g.Binds))))
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].toC(&slice2[i2])
		}
	}
}
func (g *SparseBufferMemoryBindInfo) fromC(c *C.VkSparseBufferMemoryBindInfo) {
	g.Buffer = Buffer(c.buffer)
	g.Binds = make([]SparseMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].fromC(&slice2[i2])
		}
	}
}

type SparseImageOpaqueMemoryBindInfo struct {
	Image Image
	Binds []SparseMemoryBind
}

func (g *SparseImageOpaqueMemoryBindInfo) toC(c *C.VkSparseImageOpaqueMemoryBindInfo, _sa *stackAllocator) {
	c.image = C.VkImage(g.Image)
	c.bindCount = C.uint32_t(len(g.Binds))
	{
		c.pBinds = (*C.VkSparseMemoryBind)(_sa.alloc(C.sizeof_VkSparseMemoryBind * uint(len(g.Binds))))
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].toC(&slice2[i2])
		}
	}
}
func (g *SparseImageOpaqueMemoryBindInfo) fromC(c *C.VkSparseImageOpaqueMemoryBindInfo) {
	g.Image = Image(c.image)
	g.Binds = make([]SparseMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].fromC(&slice2[i2])
		}
	}
}

type ImageSubresource struct {
	AspectMask ImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}

func (g *ImageSubresource) toC(c *C.VkImageSubresource) {
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	c.mipLevel = C.uint32_t(g.MipLevel)
	c.arrayLayer = C.uint32_t(g.ArrayLayer)
}
func (g *ImageSubresource) fromC(c *C.VkImageSubresource) {
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.MipLevel = uint32(c.mipLevel)
	g.ArrayLayer = uint32(c.arrayLayer)
}

type Offset3D struct {
	X int32
	Y int32
	Z int32
}

func (g *Offset3D) toC(c *C.VkOffset3D) {
	c.x = C.int32_t(g.X)
	c.y = C.int32_t(g.Y)
	c.z = C.int32_t(g.Z)
}
func (g *Offset3D) fromC(c *C.VkOffset3D) {
	g.X = int32(c.x)
	g.Y = int32(c.y)
	g.Z = int32(c.z)
}

type SparseImageMemoryBind struct {
	Subresource  ImageSubresource
	Offset       Offset3D
	Extent       Extent3D
	Memory       DeviceMemory
	MemoryOffset DeviceSize
	Flags        SparseMemoryBindFlags
}

func (g *SparseImageMemoryBind) toC(c *C.VkSparseImageMemoryBind) {
	g.Subresource.toC(&c.subresource)
	g.Offset.toC(&c.offset)
	g.Extent.toC(&c.extent)
	c.memory = C.VkDeviceMemory(g.Memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MemoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkSparseMemoryBindFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSparseMemoryBindFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSparseMemoryBindFlags(temp_in_VkSparseMemoryBindFlags)
	}
}
func (g *SparseImageMemoryBind) fromC(c *C.VkSparseImageMemoryBind) {
	g.Subresource.fromC(&c.subresource)
	g.Offset.fromC(&c.offset)
	g.Extent.fromC(&c.extent)
	g.Memory = DeviceMemory(c.memory)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.memoryOffset))
		g.MemoryOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkSparseMemoryBindFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSparseMemoryBindFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SparseMemoryBindFlags(temp_in_VkSparseMemoryBindFlags)
	}
}

type SparseImageMemoryBindInfo struct {
	Image Image
	Binds []SparseImageMemoryBind
}

func (g *SparseImageMemoryBindInfo) toC(c *C.VkSparseImageMemoryBindInfo, _sa *stackAllocator) {
	c.image = C.VkImage(g.Image)
	c.bindCount = C.uint32_t(len(g.Binds))
	{
		c.pBinds = (*C.VkSparseImageMemoryBind)(_sa.alloc(C.sizeof_VkSparseImageMemoryBind * uint(len(g.Binds))))
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].toC(&slice2[i2])
		}
	}
}
func (g *SparseImageMemoryBindInfo) fromC(c *C.VkSparseImageMemoryBindInfo) {
	g.Image = Image(c.image)
	g.Binds = make([]SparseImageMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.Binds):len(g.Binds)]
		for i2, _ := range g.Binds {
			g.Binds[i2].fromC(&slice2[i2])
		}
	}
}

type BindSparseInfo struct {
	Type             StructureType
	Next             unsafe.Pointer
	WaitSemaphores   []Semaphore
	BufferBinds      []SparseBufferMemoryBindInfo
	ImageOpaqueBinds []SparseImageOpaqueMemoryBindInfo
	ImageBinds       []SparseImageMemoryBindInfo
	SignalSemaphores []Semaphore
}

func (g *BindSparseInfo) toC(c *C.VkBindSparseInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.waitSemaphoreCount = C.uint32_t(len(g.WaitSemaphores))
	{
		c.pWaitSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.WaitSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			slice2[i2] = C.VkSemaphore(g.WaitSemaphores[i2])
		}
	}
	c.bufferBindCount = C.uint32_t(len(g.BufferBinds))
	{
		c.pBufferBinds = (*C.VkSparseBufferMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseBufferMemoryBindInfo * uint(len(g.BufferBinds))))
		slice2 := (*[1 << 31]C.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(c.pBufferBinds))[:len(g.BufferBinds):len(g.BufferBinds)]
		for i2, _ := range g.BufferBinds {
			g.BufferBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.imageOpaqueBindCount = C.uint32_t(len(g.ImageOpaqueBinds))
	{
		c.pImageOpaqueBinds = (*C.VkSparseImageOpaqueMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseImageOpaqueMemoryBindInfo * uint(len(g.ImageOpaqueBinds))))
		slice2 := (*[1 << 31]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(c.pImageOpaqueBinds))[:len(g.ImageOpaqueBinds):len(g.ImageOpaqueBinds)]
		for i2, _ := range g.ImageOpaqueBinds {
			g.ImageOpaqueBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.imageBindCount = C.uint32_t(len(g.ImageBinds))
	{
		c.pImageBinds = (*C.VkSparseImageMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseImageMemoryBindInfo * uint(len(g.ImageBinds))))
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBindInfo)(unsafe.Pointer(c.pImageBinds))[:len(g.ImageBinds):len(g.ImageBinds)]
		for i2, _ := range g.ImageBinds {
			g.ImageBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.signalSemaphoreCount = C.uint32_t(len(g.SignalSemaphores))
	{
		c.pSignalSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.SignalSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.SignalSemaphores):len(g.SignalSemaphores)]
		for i2, _ := range g.SignalSemaphores {
			slice2[i2] = C.VkSemaphore(g.SignalSemaphores[i2])
		}
	}
}
func (g *BindSparseInfo) fromC(c *C.VkBindSparseInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.WaitSemaphores = make([]Semaphore, int(c.waitSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			g.WaitSemaphores[i2] = Semaphore(slice2[i2])
		}
	}
	g.BufferBinds = make([]SparseBufferMemoryBindInfo, int(c.bufferBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(c.pBufferBinds))[:len(g.BufferBinds):len(g.BufferBinds)]
		for i2, _ := range g.BufferBinds {
			g.BufferBinds[i2].fromC(&slice2[i2])
		}
	}
	g.ImageOpaqueBinds = make([]SparseImageOpaqueMemoryBindInfo, int(c.imageOpaqueBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(c.pImageOpaqueBinds))[:len(g.ImageOpaqueBinds):len(g.ImageOpaqueBinds)]
		for i2, _ := range g.ImageOpaqueBinds {
			g.ImageOpaqueBinds[i2].fromC(&slice2[i2])
		}
	}
	g.ImageBinds = make([]SparseImageMemoryBindInfo, int(c.imageBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBindInfo)(unsafe.Pointer(c.pImageBinds))[:len(g.ImageBinds):len(g.ImageBinds)]
		for i2, _ := range g.ImageBinds {
			g.ImageBinds[i2].fromC(&slice2[i2])
		}
	}
	g.SignalSemaphores = make([]Semaphore, int(c.signalSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.SignalSemaphores):len(g.SignalSemaphores)]
		for i2, _ := range g.SignalSemaphores {
			g.SignalSemaphores[i2] = Semaphore(slice2[i2])
		}
	}
}
func QueueBindSparse(queue Queue, bindInfoCount uint32, bindInfo *BindSparseInfo, fence Fence) (_ret Result) {
	var c struct {
		queue         C.VkQueue
		bindInfoCount C.uint32_t
		pBindInfo     *C.VkBindSparseInfo
		fence         C.VkFence
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.queue = C.VkQueue(queue)
	c.bindInfoCount = C.uint32_t(bindInfoCount)
	{
		c.pBindInfo = (*C.VkBindSparseInfo)(_sa.alloc(C.sizeof_VkBindSparseInfo))
		bindInfo.toC(c.pBindInfo, _sa)
	}
	c.fence = C.VkFence(fence)
	c._ret = C.vkQueueBindSparse(c.queue, c.bindInfoCount, c.pBindInfo, c.fence)
	_ret = Result(c._ret)
	return
}

type FenceCreateFlags Flags
type FenceCreateInfo struct {
	Type  StructureType
	Next  unsafe.Pointer
	Flags FenceCreateFlags
}

func (g *FenceCreateInfo) toC(c *C.VkFenceCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkFenceCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkFenceCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkFenceCreateFlags(temp_in_VkFenceCreateFlags)
	}
}
func (g *FenceCreateInfo) fromC(c *C.VkFenceCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkFenceCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkFenceCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = FenceCreateFlags(temp_in_VkFenceCreateFlags)
	}
}
func CreateFence(device Device, createInfo *FenceCreateInfo, allocator *AllocationCallbacks, fence *Fence) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkFenceCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pFence      *C.VkFence
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkFenceCreateInfo)(_sa.alloc(C.sizeof_VkFenceCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pFence = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence))
		*c.pFence = C.VkFence(*fence)
	}
	c._ret = C.vkCreateFence(c.device, c.pCreateInfo, c.pAllocator, c.pFence)
	_ret = Result(c._ret)
	*fence = Fence(*c.pFence)
	return
}
func DestroyFence(device Device, fence Fence, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		fence      C.VkFence
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.fence = C.VkFence(fence)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyFence(c.device, c.fence, c.pAllocator)
}
func ResetFences(device Device, fences []Fence) (_ret Result) {
	var c struct {
		device     C.VkDevice
		fenceCount C.uint32_t
		pFences    *C.VkFence
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.fenceCount = C.uint32_t(len(fences))
	{
		c.pFences = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence * uint(len(fences))))
		slice3 := (*[1 << 31]C.VkFence)(unsafe.Pointer(c.pFences))[:len(fences):len(fences)]
		for i3, _ := range fences {
			slice3[i3] = C.VkFence(fences[i3])
		}
	}
	c._ret = C.vkResetFences(c.device, c.fenceCount, c.pFences)
	_ret = Result(c._ret)
	return
}
func GetFenceStatus(device Device, fence Fence) (_ret Result) {
	var c struct {
		device C.VkDevice
		fence  C.VkFence
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.fence = C.VkFence(fence)
	c._ret = C.vkGetFenceStatus(c.device, c.fence)
	_ret = Result(c._ret)
	return
}
func WaitForFences(device Device, fences []Fence, waitAll bool, timeout uint64) (_ret Result) {
	var c struct {
		device     C.VkDevice
		fenceCount C.uint32_t
		pFences    *C.VkFence
		waitAll    C.VkBool32
		timeout    C.uint64_t
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.fenceCount = C.uint32_t(len(fences))
	{
		c.pFences = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence * uint(len(fences))))
		slice3 := (*[1 << 31]C.VkFence)(unsafe.Pointer(c.pFences))[:len(fences):len(fences)]
		for i3, _ := range fences {
			slice3[i3] = C.VkFence(fences[i3])
		}
	}
	if waitAll {
		c.waitAll = 1
	} else {
		c.waitAll = 0
	}
	c.timeout = C.uint64_t(timeout)
	c._ret = C.vkWaitForFences(c.device, c.fenceCount, c.pFences, c.waitAll, c.timeout)
	_ret = Result(c._ret)
	return
}

type SemaphoreCreateFlags Flags
type SemaphoreCreateInfo struct {
	Type  StructureType
	Next  unsafe.Pointer
	Flags SemaphoreCreateFlags
}

func (g *SemaphoreCreateInfo) toC(c *C.VkSemaphoreCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkSemaphoreCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSemaphoreCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSemaphoreCreateFlags(temp_in_VkSemaphoreCreateFlags)
	}
}
func (g *SemaphoreCreateInfo) fromC(c *C.VkSemaphoreCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkSemaphoreCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSemaphoreCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SemaphoreCreateFlags(temp_in_VkSemaphoreCreateFlags)
	}
}
func CreateSemaphore(device Device, createInfo *SemaphoreCreateInfo, allocator *AllocationCallbacks, semaphore *Semaphore) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkSemaphoreCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pSemaphore  *C.VkSemaphore
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkSemaphoreCreateInfo)(_sa.alloc(C.sizeof_VkSemaphoreCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSemaphore = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore))
		*c.pSemaphore = C.VkSemaphore(*semaphore)
	}
	c._ret = C.vkCreateSemaphore(c.device, c.pCreateInfo, c.pAllocator, c.pSemaphore)
	_ret = Result(c._ret)
	*semaphore = Semaphore(*c.pSemaphore)
	return
}
func DestroySemaphore(device Device, semaphore Semaphore, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		semaphore  C.VkSemaphore
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.semaphore = C.VkSemaphore(semaphore)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySemaphore(c.device, c.semaphore, c.pAllocator)
}

type EventCreateFlags Flags
type EventCreateInfo struct {
	Type  StructureType
	Next  unsafe.Pointer
	Flags EventCreateFlags
}

func (g *EventCreateInfo) toC(c *C.VkEventCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkEventCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkEventCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkEventCreateFlags(temp_in_VkEventCreateFlags)
	}
}
func (g *EventCreateInfo) fromC(c *C.VkEventCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkEventCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkEventCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = EventCreateFlags(temp_in_VkEventCreateFlags)
	}
}

type Event C.VkEvent

func CreateEvent(device Device, createInfo *EventCreateInfo, allocator *AllocationCallbacks, event *Event) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkEventCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pEvent      *C.VkEvent
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkEventCreateInfo)(_sa.alloc(C.sizeof_VkEventCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pEvent = (*C.VkEvent)(_sa.alloc(C.sizeof_VkEvent))
		*c.pEvent = C.VkEvent(*event)
	}
	c._ret = C.vkCreateEvent(c.device, c.pCreateInfo, c.pAllocator, c.pEvent)
	_ret = Result(c._ret)
	*event = Event(*c.pEvent)
	return
}
func DestroyEvent(device Device, event Event, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		event      C.VkEvent
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyEvent(c.device, c.event, c.pAllocator)
}
func GetEventStatus(device Device, event Event) (_ret Result) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkGetEventStatus(c.device, c.event)
	_ret = Result(c._ret)
	return
}
func SetEvent(device Device, event Event) (_ret Result) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkSetEvent(c.device, c.event)
	_ret = Result(c._ret)
	return
}
func ResetEvent(device Device, event Event) (_ret Result) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkResetEvent(c.device, c.event)
	_ret = Result(c._ret)
	return
}

type QueryPoolCreateFlags Flags
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

type QueryPipelineStatisticFlags Flags
type QueryPoolCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              QueryPoolCreateFlags
	QueryType          QueryType
	QueryCount         uint32
	PipelineStatistics QueryPipelineStatisticFlags
}

func (g *QueryPoolCreateInfo) toC(c *C.VkQueryPoolCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkQueryPoolCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkQueryPoolCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkQueryPoolCreateFlags(temp_in_VkQueryPoolCreateFlags)
	}
	c.queryType = C.VkQueryType(g.QueryType)
	c.queryCount = C.uint32_t(g.QueryCount)
	{
		var temp_in_VkQueryPipelineStatisticFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.PipelineStatistics)))
			temp_in_VkQueryPipelineStatisticFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.pipelineStatistics = C.VkQueryPipelineStatisticFlags(temp_in_VkQueryPipelineStatisticFlags)
	}
}
func (g *QueryPoolCreateInfo) fromC(c *C.VkQueryPoolCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkQueryPoolCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkQueryPoolCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = QueryPoolCreateFlags(temp_in_VkQueryPoolCreateFlags)
	}
	g.QueryType = QueryType(c.queryType)
	g.QueryCount = uint32(c.queryCount)
	{
		var temp_in_VkQueryPipelineStatisticFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.pipelineStatistics)))
			temp_in_VkQueryPipelineStatisticFlags = Flags(temp_in_VkFlags)
		}
		g.PipelineStatistics = QueryPipelineStatisticFlags(temp_in_VkQueryPipelineStatisticFlags)
	}
}

type QueryPool C.VkQueryPool

func CreateQueryPool(device Device, createInfo *QueryPoolCreateInfo, allocator *AllocationCallbacks, queryPool *QueryPool) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkQueryPoolCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pQueryPool  *C.VkQueryPool
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkQueryPoolCreateInfo)(_sa.alloc(C.sizeof_VkQueryPoolCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pQueryPool = (*C.VkQueryPool)(_sa.alloc(C.sizeof_VkQueryPool))
		*c.pQueryPool = C.VkQueryPool(*queryPool)
	}
	c._ret = C.vkCreateQueryPool(c.device, c.pCreateInfo, c.pAllocator, c.pQueryPool)
	_ret = Result(c._ret)
	*queryPool = QueryPool(*c.pQueryPool)
	return
}
func DestroyQueryPool(device Device, queryPool QueryPool, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		queryPool  C.VkQueryPool
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.queryPool = C.VkQueryPool(queryPool)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyQueryPool(c.device, c.queryPool, c.pAllocator)
}

type QueryResultFlags Flags

func GetQueryPoolResults(device Device, queryPool QueryPool, firstQuery uint32, queryCount uint32, data []byte, stride DeviceSize, flags QueryResultFlags) (_ret Result) {
	var c struct {
		device     C.VkDevice
		queryPool  C.VkQueryPool
		firstQuery C.uint32_t
		queryCount C.uint32_t
		dataSize   C.size_t
		pData      unsafe.Pointer
		stride     C.VkDeviceSize
		flags      C.VkQueryResultFlags
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.queryPool = C.VkQueryPool(queryPool)
	c.firstQuery = C.uint32_t(firstQuery)
	c.queryCount = C.uint32_t(queryCount)
	c.dataSize = C.size_t(len(data))
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(stride))
		c.stride = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkQueryResultFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkQueryResultFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkQueryResultFlags(temp_in_VkQueryResultFlags)
	}
	c._ret = C.vkGetQueryPoolResults(c.device, c.queryPool, c.firstQuery, c.queryCount, c.dataSize, c.pData, c.stride, c.flags)
	_ret = Result(c._ret)
	return
}

type BufferCreateFlags Flags
type BufferUsageFlags Flags
type SharingMode int

const (
	SHARING_MODE_EXCLUSIVE   SharingMode = 0
	SHARING_MODE_CONCURRENT  SharingMode = 1
	SHARING_MODE_BEGIN_RANGE SharingMode = SHARING_MODE_EXCLUSIVE
	SHARING_MODE_END_RANGE   SharingMode = SHARING_MODE_CONCURRENT
	SHARING_MODE_RANGE_SIZE  SharingMode = (SHARING_MODE_CONCURRENT - SHARING_MODE_EXCLUSIVE + 1)
	SHARING_MODE_MAX_ENUM    SharingMode = 2147483647
)

type BufferCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              BufferCreateFlags
	Size               DeviceSize
	Usage              BufferUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
}

func (g *BufferCreateInfo) toC(c *C.VkBufferCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkBufferCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkBufferCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkBufferCreateFlags(temp_in_VkBufferCreateFlags)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkBufferUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Usage)))
			temp_in_VkBufferUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkBufferUsageFlags(temp_in_VkBufferUsageFlags)
	}
	c.sharingMode = C.VkSharingMode(g.SharingMode)
	c.queueFamilyIndexCount = C.uint32_t(len(g.QueueFamilyIndices))
	{
		c.pQueueFamilyIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.QueueFamilyIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			slice2[i2] = C.uint32_t(g.QueueFamilyIndices[i2])
		}
	}
}
func (g *BufferCreateInfo) fromC(c *C.VkBufferCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkBufferCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkBufferCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = BufferCreateFlags(temp_in_VkBufferCreateFlags)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkBufferUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			temp_in_VkBufferUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Usage = BufferUsageFlags(temp_in_VkBufferUsageFlags)
	}
	g.SharingMode = SharingMode(c.sharingMode)
	g.QueueFamilyIndices = make([]uint32, int(c.queueFamilyIndexCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			g.QueueFamilyIndices[i2] = uint32(slice2[i2])
		}
	}
}
func CreateBuffer(device Device, createInfo *BufferCreateInfo, allocator *AllocationCallbacks, buffer *Buffer) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkBufferCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pBuffer     *C.VkBuffer
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkBufferCreateInfo)(_sa.alloc(C.sizeof_VkBufferCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pBuffer = (*C.VkBuffer)(_sa.alloc(C.sizeof_VkBuffer))
		*c.pBuffer = C.VkBuffer(*buffer)
	}
	c._ret = C.vkCreateBuffer(c.device, c.pCreateInfo, c.pAllocator, c.pBuffer)
	_ret = Result(c._ret)
	*buffer = Buffer(*c.pBuffer)
	return
}
func DestroyBuffer(device Device, buffer Buffer, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		buffer     C.VkBuffer
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.buffer = C.VkBuffer(buffer)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyBuffer(c.device, c.buffer, c.pAllocator)
}

type BufferViewCreateFlags Flags
type BufferViewCreateInfo struct {
	Type   StructureType
	Next   unsafe.Pointer
	Flags  BufferViewCreateFlags
	Buffer Buffer
	Format Format
	Offset DeviceSize
	Range  DeviceSize
}

func (g *BufferViewCreateInfo) toC(c *C.VkBufferViewCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkBufferViewCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkBufferViewCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkBufferViewCreateFlags(temp_in_VkBufferViewCreateFlags)
	}
	c.buffer = C.VkBuffer(g.Buffer)
	c.format = C.VkFormat(g.Format)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Range))
		c._range = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *BufferViewCreateInfo) fromC(c *C.VkBufferViewCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkBufferViewCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkBufferViewCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = BufferViewCreateFlags(temp_in_VkBufferViewCreateFlags)
	}
	g.Buffer = Buffer(c.buffer)
	g.Format = Format(c.format)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.offset))
		g.Offset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c._range))
		g.Range = DeviceSize(temp_in_VkDeviceSize)
	}
}

type BufferView C.VkBufferView

func CreateBufferView(device Device, createInfo *BufferViewCreateInfo, allocator *AllocationCallbacks, view *BufferView) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkBufferViewCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pView       *C.VkBufferView
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkBufferViewCreateInfo)(_sa.alloc(C.sizeof_VkBufferViewCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pView = (*C.VkBufferView)(_sa.alloc(C.sizeof_VkBufferView))
		*c.pView = C.VkBufferView(*view)
	}
	c._ret = C.vkCreateBufferView(c.device, c.pCreateInfo, c.pAllocator, c.pView)
	_ret = Result(c._ret)
	*view = BufferView(*c.pView)
	return
}
func DestroyBufferView(device Device, bufferView BufferView, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		bufferView C.VkBufferView
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.bufferView = C.VkBufferView(bufferView)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyBufferView(c.device, c.bufferView, c.pAllocator)
}

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

type ImageCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              ImageCreateFlags
	ImageType          ImageType
	Format             Format
	Extent             Extent3D
	MipLevels          uint32
	ArrayLayers        uint32
	Samples            SampleCountFlagBits
	Tiling             ImageTiling
	Usage              ImageUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
	InitialLayout      ImageLayout
}

func (g *ImageCreateInfo) toC(c *C.VkImageCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkImageCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkImageCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkImageCreateFlags(temp_in_VkImageCreateFlags)
	}
	c.imageType = C.VkImageType(g.ImageType)
	c.format = C.VkFormat(g.Format)
	g.Extent.toC(&c.extent)
	c.mipLevels = C.uint32_t(g.MipLevels)
	c.arrayLayers = C.uint32_t(g.ArrayLayers)
	c.samples = C.VkSampleCountFlagBits(g.Samples)
	c.tiling = C.VkImageTiling(g.Tiling)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Usage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	c.sharingMode = C.VkSharingMode(g.SharingMode)
	c.queueFamilyIndexCount = C.uint32_t(len(g.QueueFamilyIndices))
	{
		c.pQueueFamilyIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.QueueFamilyIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			slice2[i2] = C.uint32_t(g.QueueFamilyIndices[i2])
		}
	}
	c.initialLayout = C.VkImageLayout(g.InitialLayout)
}
func (g *ImageCreateInfo) fromC(c *C.VkImageCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkImageCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkImageCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = ImageCreateFlags(temp_in_VkImageCreateFlags)
	}
	g.ImageType = ImageType(c.imageType)
	g.Format = Format(c.format)
	g.Extent.fromC(&c.extent)
	g.MipLevels = uint32(c.mipLevels)
	g.ArrayLayers = uint32(c.arrayLayers)
	g.Samples = SampleCountFlagBits(c.samples)
	g.Tiling = ImageTiling(c.tiling)
	{
		var temp_in_VkImageUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			temp_in_VkImageUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Usage = ImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	g.SharingMode = SharingMode(c.sharingMode)
	g.QueueFamilyIndices = make([]uint32, int(c.queueFamilyIndexCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			g.QueueFamilyIndices[i2] = uint32(slice2[i2])
		}
	}
	g.InitialLayout = ImageLayout(c.initialLayout)
}
func CreateImage(device Device, createInfo *ImageCreateInfo, allocator *AllocationCallbacks, image *Image) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkImageCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pImage      *C.VkImage
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkImageCreateInfo)(_sa.alloc(C.sizeof_VkImageCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pImage = (*C.VkImage)(_sa.alloc(C.sizeof_VkImage))
		*c.pImage = C.VkImage(*image)
	}
	c._ret = C.vkCreateImage(c.device, c.pCreateInfo, c.pAllocator, c.pImage)
	_ret = Result(c._ret)
	*image = Image(*c.pImage)
	return
}
func DestroyImage(device Device, image Image, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		image      C.VkImage
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.image = C.VkImage(image)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyImage(c.device, c.image, c.pAllocator)
}

type SubresourceLayout struct {
	Offset     DeviceSize
	Size       DeviceSize
	RowPitch   DeviceSize
	ArrayPitch DeviceSize
	DepthPitch DeviceSize
}

func (g *SubresourceLayout) toC(c *C.VkSubresourceLayout) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.RowPitch))
		c.rowPitch = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.ArrayPitch))
		c.arrayPitch = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.DepthPitch))
		c.depthPitch = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *SubresourceLayout) fromC(c *C.VkSubresourceLayout) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.offset))
		g.Offset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.rowPitch))
		g.RowPitch = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.arrayPitch))
		g.ArrayPitch = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.depthPitch))
		g.DepthPitch = DeviceSize(temp_in_VkDeviceSize)
	}
}
func GetImageSubresourceLayout(device Device, image Image, subresource *ImageSubresource, layout *SubresourceLayout) {
	var c struct {
		device       C.VkDevice
		image        C.VkImage
		pSubresource *C.VkImageSubresource
		pLayout      *C.VkSubresourceLayout
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.image = C.VkImage(image)
	{
		c.pSubresource = (*C.VkImageSubresource)(_sa.alloc(C.sizeof_VkImageSubresource))
		subresource.toC(c.pSubresource)
	}
	{
		c.pLayout = (*C.VkSubresourceLayout)(_sa.alloc(C.sizeof_VkSubresourceLayout))
		layout.toC(c.pLayout)
	}
	C.vkGetImageSubresourceLayout(c.device, c.image, c.pSubresource, c.pLayout)
	layout.fromC(c.pLayout)
}

type ImageViewCreateFlags Flags
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

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

func (g *ComponentMapping) toC(c *C.VkComponentMapping) {
	c.r = C.VkComponentSwizzle(g.R)
	c.g = C.VkComponentSwizzle(g.G)
	c.b = C.VkComponentSwizzle(g.B)
	c.a = C.VkComponentSwizzle(g.A)
}
func (g *ComponentMapping) fromC(c *C.VkComponentMapping) {
	g.R = ComponentSwizzle(c.r)
	g.G = ComponentSwizzle(c.g)
	g.B = ComponentSwizzle(c.b)
	g.A = ComponentSwizzle(c.a)
}

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (g *ImageSubresourceRange) toC(c *C.VkImageSubresourceRange) {
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	c.baseMipLevel = C.uint32_t(g.BaseMipLevel)
	c.levelCount = C.uint32_t(g.LevelCount)
	c.baseArrayLayer = C.uint32_t(g.BaseArrayLayer)
	c.layerCount = C.uint32_t(g.LayerCount)
}
func (g *ImageSubresourceRange) fromC(c *C.VkImageSubresourceRange) {
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.BaseMipLevel = uint32(c.baseMipLevel)
	g.LevelCount = uint32(c.levelCount)
	g.BaseArrayLayer = uint32(c.baseArrayLayer)
	g.LayerCount = uint32(c.layerCount)
}

type ImageViewCreateInfo struct {
	Type             StructureType
	Next             unsafe.Pointer
	Flags            ImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

func (g *ImageViewCreateInfo) toC(c *C.VkImageViewCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkImageViewCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkImageViewCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkImageViewCreateFlags(temp_in_VkImageViewCreateFlags)
	}
	c.image = C.VkImage(g.Image)
	c.viewType = C.VkImageViewType(g.ViewType)
	c.format = C.VkFormat(g.Format)
	g.Components.toC(&c.components)
	g.SubresourceRange.toC(&c.subresourceRange)
}
func (g *ImageViewCreateInfo) fromC(c *C.VkImageViewCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkImageViewCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkImageViewCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = ImageViewCreateFlags(temp_in_VkImageViewCreateFlags)
	}
	g.Image = Image(c.image)
	g.ViewType = ImageViewType(c.viewType)
	g.Format = Format(c.format)
	g.Components.fromC(&c.components)
	g.SubresourceRange.fromC(&c.subresourceRange)
}

type ImageView C.VkImageView

func CreateImageView(device Device, createInfo *ImageViewCreateInfo, allocator *AllocationCallbacks, view *ImageView) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkImageViewCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pView       *C.VkImageView
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkImageViewCreateInfo)(_sa.alloc(C.sizeof_VkImageViewCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pView = (*C.VkImageView)(_sa.alloc(C.sizeof_VkImageView))
		*c.pView = C.VkImageView(*view)
	}
	c._ret = C.vkCreateImageView(c.device, c.pCreateInfo, c.pAllocator, c.pView)
	_ret = Result(c._ret)
	*view = ImageView(*c.pView)
	return
}
func DestroyImageView(device Device, imageView ImageView, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		imageView  C.VkImageView
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.imageView = C.VkImageView(imageView)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyImageView(c.device, c.imageView, c.pAllocator)
}

type ShaderModuleCreateFlags Flags
type ShaderModuleCreateInfo struct {
	Type     StructureType
	Next     unsafe.Pointer
	Flags    ShaderModuleCreateFlags
	CodeSize uint
	Code     *uint32
}

func (g *ShaderModuleCreateInfo) toC(c *C.VkShaderModuleCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkShaderModuleCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkShaderModuleCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkShaderModuleCreateFlags(temp_in_VkShaderModuleCreateFlags)
	}
	c.codeSize = C.size_t(g.CodeSize)
	{
		c.pCode = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pCode = C.uint32_t(*g.Code)
	}
}
func (g *ShaderModuleCreateInfo) fromC(c *C.VkShaderModuleCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkShaderModuleCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkShaderModuleCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = ShaderModuleCreateFlags(temp_in_VkShaderModuleCreateFlags)
	}
	g.CodeSize = uint(c.codeSize)
	{
		if g.Code == nil {
			g.Code = new(uint32)
		}
		*g.Code = uint32(*c.pCode)
	}
}

type ShaderModule C.VkShaderModule

func CreateShaderModule(device Device, createInfo *ShaderModuleCreateInfo, allocator *AllocationCallbacks, shaderModule *ShaderModule) (_ret Result) {
	var c struct {
		device        C.VkDevice
		pCreateInfo   *C.VkShaderModuleCreateInfo
		pAllocator    *C.VkAllocationCallbacks
		pShaderModule *C.VkShaderModule
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkShaderModuleCreateInfo)(_sa.alloc(C.sizeof_VkShaderModuleCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pShaderModule = (*C.VkShaderModule)(_sa.alloc(C.sizeof_VkShaderModule))
		*c.pShaderModule = C.VkShaderModule(*shaderModule)
	}
	c._ret = C.vkCreateShaderModule(c.device, c.pCreateInfo, c.pAllocator, c.pShaderModule)
	_ret = Result(c._ret)
	*shaderModule = ShaderModule(*c.pShaderModule)
	return
}
func DestroyShaderModule(device Device, shaderModule ShaderModule, allocator *AllocationCallbacks) {
	var c struct {
		device       C.VkDevice
		shaderModule C.VkShaderModule
		pAllocator   *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.shaderModule = C.VkShaderModule(shaderModule)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyShaderModule(c.device, c.shaderModule, c.pAllocator)
}

type PipelineCacheCreateFlags Flags
type PipelineCacheCreateInfo struct {
	Type        StructureType
	Next        unsafe.Pointer
	Flags       PipelineCacheCreateFlags
	InitialData []byte
}

func (g *PipelineCacheCreateInfo) toC(c *C.VkPipelineCacheCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineCacheCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineCacheCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineCacheCreateFlags(temp_in_VkPipelineCacheCreateFlags)
	}
	c.initialDataSize = C.size_t(len(g.InitialData))
	{
		c.pInitialData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.InitialData)))
		slice2 := (*[1 << 31]byte)(c.pInitialData)[:len(g.InitialData):len(g.InitialData)]
		for i2, _ := range g.InitialData {
			slice2[i2] = g.InitialData[i2]
		}
	}
}
func (g *PipelineCacheCreateInfo) fromC(c *C.VkPipelineCacheCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineCacheCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineCacheCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineCacheCreateFlags(temp_in_VkPipelineCacheCreateFlags)
	}
	g.InitialData = make([]byte, int(c.initialDataSize))
	{
		slice2 := (*[1 << 31]byte)(c.pInitialData)[:len(g.InitialData):len(g.InitialData)]
		for i2, _ := range g.InitialData {
			g.InitialData[i2] = slice2[i2]
		}
	}
}

type PipelineCache C.VkPipelineCache

func CreatePipelineCache(device Device, createInfo *PipelineCacheCreateInfo, allocator *AllocationCallbacks, pipelineCache *PipelineCache) (_ret Result) {
	var c struct {
		device         C.VkDevice
		pCreateInfo    *C.VkPipelineCacheCreateInfo
		pAllocator     *C.VkAllocationCallbacks
		pPipelineCache *C.VkPipelineCache
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkPipelineCacheCreateInfo)(_sa.alloc(C.sizeof_VkPipelineCacheCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelineCache = (*C.VkPipelineCache)(_sa.alloc(C.sizeof_VkPipelineCache))
		*c.pPipelineCache = C.VkPipelineCache(*pipelineCache)
	}
	c._ret = C.vkCreatePipelineCache(c.device, c.pCreateInfo, c.pAllocator, c.pPipelineCache)
	_ret = Result(c._ret)
	*pipelineCache = PipelineCache(*c.pPipelineCache)
	return
}
func DestroyPipelineCache(device Device, pipelineCache PipelineCache, allocator *AllocationCallbacks) {
	var c struct {
		device        C.VkDevice
		pipelineCache C.VkPipelineCache
		pAllocator    *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipelineCache = C.VkPipelineCache(pipelineCache)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipelineCache(c.device, c.pipelineCache, c.pAllocator)
}
func GetPipelineCacheData(device Device, pipelineCache PipelineCache, dataSize *uint, data []byte) (_ret Result) {
	var c struct {
		device        C.VkDevice
		pipelineCache C.VkPipelineCache
		pDataSize     *C.size_t
		pData         unsafe.Pointer
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipelineCache = C.VkPipelineCache(pipelineCache)
	{
		c.pDataSize = (*C.size_t)(_sa.alloc(C.sizeof_size_t))
		*c.pDataSize = C.size_t(*dataSize)
	}
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	c._ret = C.vkGetPipelineCacheData(c.device, c.pipelineCache, c.pDataSize, c.pData)
	_ret = Result(c._ret)
	*dataSize = uint(*c.pDataSize)
	return
}
func MergePipelineCaches(device Device, dstCache PipelineCache, srcCaches []PipelineCache) (_ret Result) {
	var c struct {
		device        C.VkDevice
		dstCache      C.VkPipelineCache
		srcCacheCount C.uint32_t
		pSrcCaches    *C.VkPipelineCache
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.dstCache = C.VkPipelineCache(dstCache)
	c.srcCacheCount = C.uint32_t(len(srcCaches))
	{
		c.pSrcCaches = (*C.VkPipelineCache)(_sa.alloc(C.sizeof_VkPipelineCache * uint(len(srcCaches))))
		slice3 := (*[1 << 31]C.VkPipelineCache)(unsafe.Pointer(c.pSrcCaches))[:len(srcCaches):len(srcCaches)]
		for i3, _ := range srcCaches {
			slice3[i3] = C.VkPipelineCache(srcCaches[i3])
		}
	}
	c._ret = C.vkMergePipelineCaches(c.device, c.dstCache, c.srcCacheCount, c.pSrcCaches)
	_ret = Result(c._ret)
	return
}

type PipelineCreateFlags Flags
type PipelineShaderStageCreateFlags Flags
type ShaderStageFlagBits int

const (
	SHADER_STAGE_VERTEX_BIT                  ShaderStageFlagBits = 1
	SHADER_STAGE_TESSELLATION_CONTROL_BIT    ShaderStageFlagBits = 2
	SHADER_STAGE_TESSELLATION_EVALUATION_BIT ShaderStageFlagBits = 4
	SHADER_STAGE_GEOMETRY_BIT                ShaderStageFlagBits = 8
	SHADER_STAGE_FRAGMENT_BIT                ShaderStageFlagBits = 16
	SHADER_STAGE_COMPUTE_BIT                 ShaderStageFlagBits = 32
	SHADER_STAGE_ALL_GRAPHICS                ShaderStageFlagBits = 31
	SHADER_STAGE_ALL                         ShaderStageFlagBits = 2147483647
	SHADER_STAGE_FLAG_BITS_MAX_ENUM          ShaderStageFlagBits = 2147483647
)

type SpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uint
}

func (g *SpecializationMapEntry) toC(c *C.VkSpecializationMapEntry) {
	c.constantID = C.uint32_t(g.ConstantID)
	c.offset = C.uint32_t(g.Offset)
	c.size = C.size_t(g.Size)
}
func (g *SpecializationMapEntry) fromC(c *C.VkSpecializationMapEntry) {
	g.ConstantID = uint32(c.constantID)
	g.Offset = uint32(c.offset)
	g.Size = uint(c.size)
}

type SpecializationInfo struct {
	MapEntries []SpecializationMapEntry
	Data       []byte
}

func (g *SpecializationInfo) toC(c *C.VkSpecializationInfo, _sa *stackAllocator) {
	c.mapEntryCount = C.uint32_t(len(g.MapEntries))
	{
		c.pMapEntries = (*C.VkSpecializationMapEntry)(_sa.alloc(C.sizeof_VkSpecializationMapEntry * uint(len(g.MapEntries))))
		slice2 := (*[1 << 31]C.VkSpecializationMapEntry)(unsafe.Pointer(c.pMapEntries))[:len(g.MapEntries):len(g.MapEntries)]
		for i2, _ := range g.MapEntries {
			g.MapEntries[i2].toC(&slice2[i2])
		}
	}
	c.dataSize = C.size_t(len(g.Data))
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.Data)))
		slice2 := (*[1 << 31]byte)(c.pData)[:len(g.Data):len(g.Data)]
		for i2, _ := range g.Data {
			slice2[i2] = g.Data[i2]
		}
	}
}
func (g *SpecializationInfo) fromC(c *C.VkSpecializationInfo) {
	g.MapEntries = make([]SpecializationMapEntry, int(c.mapEntryCount))
	{
		slice2 := (*[1 << 31]C.VkSpecializationMapEntry)(unsafe.Pointer(c.pMapEntries))[:len(g.MapEntries):len(g.MapEntries)]
		for i2, _ := range g.MapEntries {
			g.MapEntries[i2].fromC(&slice2[i2])
		}
	}
	g.Data = make([]byte, int(c.dataSize))
	{
		slice2 := (*[1 << 31]byte)(c.pData)[:len(g.Data):len(g.Data)]
		for i2, _ := range g.Data {
			g.Data[i2] = slice2[i2]
		}
	}
}

type PipelineShaderStageCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              PipelineShaderStageCreateFlags
	Stage              ShaderStageFlagBits
	Module             ShaderModule
	Name               string
	SpecializationInfo *SpecializationInfo
}

func (g *PipelineShaderStageCreateInfo) toC(c *C.VkPipelineShaderStageCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineShaderStageCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineShaderStageCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineShaderStageCreateFlags(temp_in_VkPipelineShaderStageCreateFlags)
	}
	c.stage = C.VkShaderStageFlagBits(g.Stage)
	c.module = C.VkShaderModule(g.Module)
	c.pName = toCString(g.Name, _sa)
	{
		c.pSpecializationInfo = (*C.VkSpecializationInfo)(_sa.alloc(C.sizeof_VkSpecializationInfo))
		g.SpecializationInfo.toC(c.pSpecializationInfo, _sa)
	}
}
func (g *PipelineShaderStageCreateInfo) fromC(c *C.VkPipelineShaderStageCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineShaderStageCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineShaderStageCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineShaderStageCreateFlags(temp_in_VkPipelineShaderStageCreateFlags)
	}
	g.Stage = ShaderStageFlagBits(c.stage)
	g.Module = ShaderModule(c.module)
	g.Name = toGoString(c.pName)
	{
		if g.SpecializationInfo == nil {
			g.SpecializationInfo = new(SpecializationInfo)
		}
		g.SpecializationInfo.fromC(c.pSpecializationInfo)
	}
}

type PipelineVertexInputStateCreateFlags Flags
type VertexInputRate int

const (
	VERTEX_INPUT_RATE_VERTEX      VertexInputRate = 0
	VERTEX_INPUT_RATE_INSTANCE    VertexInputRate = 1
	VERTEX_INPUT_RATE_BEGIN_RANGE VertexInputRate = VERTEX_INPUT_RATE_VERTEX
	VERTEX_INPUT_RATE_END_RANGE   VertexInputRate = VERTEX_INPUT_RATE_INSTANCE
	VERTEX_INPUT_RATE_RANGE_SIZE  VertexInputRate = (VERTEX_INPUT_RATE_INSTANCE - VERTEX_INPUT_RATE_VERTEX + 1)
	VERTEX_INPUT_RATE_MAX_ENUM    VertexInputRate = 2147483647
)

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

func (g *VertexInputBindingDescription) toC(c *C.VkVertexInputBindingDescription) {
	c.binding = C.uint32_t(g.Binding)
	c.stride = C.uint32_t(g.Stride)
	c.inputRate = C.VkVertexInputRate(g.InputRate)
}
func (g *VertexInputBindingDescription) fromC(c *C.VkVertexInputBindingDescription) {
	g.Binding = uint32(c.binding)
	g.Stride = uint32(c.stride)
	g.InputRate = VertexInputRate(c.inputRate)
}

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

func (g *VertexInputAttributeDescription) toC(c *C.VkVertexInputAttributeDescription) {
	c.location = C.uint32_t(g.Location)
	c.binding = C.uint32_t(g.Binding)
	c.format = C.VkFormat(g.Format)
	c.offset = C.uint32_t(g.Offset)
}
func (g *VertexInputAttributeDescription) fromC(c *C.VkVertexInputAttributeDescription) {
	g.Location = uint32(c.location)
	g.Binding = uint32(c.binding)
	g.Format = Format(c.format)
	g.Offset = uint32(c.offset)
}

type PipelineVertexInputStateCreateInfo struct {
	Type                        StructureType
	Next                        unsafe.Pointer
	Flags                       PipelineVertexInputStateCreateFlags
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (g *PipelineVertexInputStateCreateInfo) toC(c *C.VkPipelineVertexInputStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineVertexInputStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineVertexInputStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineVertexInputStateCreateFlags(temp_in_VkPipelineVertexInputStateCreateFlags)
	}
	c.vertexBindingDescriptionCount = C.uint32_t(len(g.VertexBindingDescriptions))
	{
		c.pVertexBindingDescriptions = (*C.VkVertexInputBindingDescription)(_sa.alloc(C.sizeof_VkVertexInputBindingDescription * uint(len(g.VertexBindingDescriptions))))
		slice2 := (*[1 << 31]C.VkVertexInputBindingDescription)(unsafe.Pointer(c.pVertexBindingDescriptions))[:len(g.VertexBindingDescriptions):len(g.VertexBindingDescriptions)]
		for i2, _ := range g.VertexBindingDescriptions {
			g.VertexBindingDescriptions[i2].toC(&slice2[i2])
		}
	}
	c.vertexAttributeDescriptionCount = C.uint32_t(len(g.VertexAttributeDescriptions))
	{
		c.pVertexAttributeDescriptions = (*C.VkVertexInputAttributeDescription)(_sa.alloc(C.sizeof_VkVertexInputAttributeDescription * uint(len(g.VertexAttributeDescriptions))))
		slice2 := (*[1 << 31]C.VkVertexInputAttributeDescription)(unsafe.Pointer(c.pVertexAttributeDescriptions))[:len(g.VertexAttributeDescriptions):len(g.VertexAttributeDescriptions)]
		for i2, _ := range g.VertexAttributeDescriptions {
			g.VertexAttributeDescriptions[i2].toC(&slice2[i2])
		}
	}
}
func (g *PipelineVertexInputStateCreateInfo) fromC(c *C.VkPipelineVertexInputStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineVertexInputStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineVertexInputStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineVertexInputStateCreateFlags(temp_in_VkPipelineVertexInputStateCreateFlags)
	}
	g.VertexBindingDescriptions = make([]VertexInputBindingDescription, int(c.vertexBindingDescriptionCount))
	{
		slice2 := (*[1 << 31]C.VkVertexInputBindingDescription)(unsafe.Pointer(c.pVertexBindingDescriptions))[:len(g.VertexBindingDescriptions):len(g.VertexBindingDescriptions)]
		for i2, _ := range g.VertexBindingDescriptions {
			g.VertexBindingDescriptions[i2].fromC(&slice2[i2])
		}
	}
	g.VertexAttributeDescriptions = make([]VertexInputAttributeDescription, int(c.vertexAttributeDescriptionCount))
	{
		slice2 := (*[1 << 31]C.VkVertexInputAttributeDescription)(unsafe.Pointer(c.pVertexAttributeDescriptions))[:len(g.VertexAttributeDescriptions):len(g.VertexAttributeDescriptions)]
		for i2, _ := range g.VertexAttributeDescriptions {
			g.VertexAttributeDescriptions[i2].fromC(&slice2[i2])
		}
	}
}

type PipelineInputAssemblyStateCreateFlags Flags
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

type PipelineInputAssemblyStateCreateInfo struct {
	Type                   StructureType
	Next                   unsafe.Pointer
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
}

func (g *PipelineInputAssemblyStateCreateInfo) toC(c *C.VkPipelineInputAssemblyStateCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineInputAssemblyStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineInputAssemblyStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineInputAssemblyStateCreateFlags(temp_in_VkPipelineInputAssemblyStateCreateFlags)
	}
	c.topology = C.VkPrimitiveTopology(g.Topology)
	if g.PrimitiveRestartEnable {
		c.primitiveRestartEnable = 1
	} else {
		c.primitiveRestartEnable = 0
	}
}
func (g *PipelineInputAssemblyStateCreateInfo) fromC(c *C.VkPipelineInputAssemblyStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineInputAssemblyStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineInputAssemblyStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineInputAssemblyStateCreateFlags(temp_in_VkPipelineInputAssemblyStateCreateFlags)
	}
	g.Topology = PrimitiveTopology(c.topology)
	g.PrimitiveRestartEnable = c.primitiveRestartEnable != 0
}

type PipelineTessellationStateCreateFlags Flags
type PipelineTessellationStateCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              PipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

func (g *PipelineTessellationStateCreateInfo) toC(c *C.VkPipelineTessellationStateCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineTessellationStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineTessellationStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineTessellationStateCreateFlags(temp_in_VkPipelineTessellationStateCreateFlags)
	}
	c.patchControlPoints = C.uint32_t(g.PatchControlPoints)
}
func (g *PipelineTessellationStateCreateInfo) fromC(c *C.VkPipelineTessellationStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineTessellationStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineTessellationStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineTessellationStateCreateFlags(temp_in_VkPipelineTessellationStateCreateFlags)
	}
	g.PatchControlPoints = uint32(c.patchControlPoints)
}

type PipelineViewportStateCreateFlags Flags
type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

func (g *Viewport) toC(c *C.VkViewport) {
	c.x = C.float(g.X)
	c.y = C.float(g.Y)
	c.width = C.float(g.Width)
	c.height = C.float(g.Height)
	c.minDepth = C.float(g.MinDepth)
	c.maxDepth = C.float(g.MaxDepth)
}
func (g *Viewport) fromC(c *C.VkViewport) {
	g.X = float32(c.x)
	g.Y = float32(c.y)
	g.Width = float32(c.width)
	g.Height = float32(c.height)
	g.MinDepth = float32(c.minDepth)
	g.MaxDepth = float32(c.maxDepth)
}

type Offset2D struct {
	X int32
	Y int32
}

func (g *Offset2D) toC(c *C.VkOffset2D) {
	c.x = C.int32_t(g.X)
	c.y = C.int32_t(g.Y)
}
func (g *Offset2D) fromC(c *C.VkOffset2D) {
	g.X = int32(c.x)
	g.Y = int32(c.y)
}

type Extent2D struct {
	Width  uint32
	Height uint32
}

func (g *Extent2D) toC(c *C.VkExtent2D) {
	c.width = C.uint32_t(g.Width)
	c.height = C.uint32_t(g.Height)
}
func (g *Extent2D) fromC(c *C.VkExtent2D) {
	g.Width = uint32(c.width)
	g.Height = uint32(c.height)
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

func (g *Rect2D) toC(c *C.VkRect2D) {
	g.Offset.toC(&c.offset)
	g.Extent.toC(&c.extent)
}
func (g *Rect2D) fromC(c *C.VkRect2D) {
	g.Offset.fromC(&c.offset)
	g.Extent.fromC(&c.extent)
}

type PipelineViewportStateCreateInfo struct {
	Type      StructureType
	Next      unsafe.Pointer
	Flags     PipelineViewportStateCreateFlags
	Viewports []Viewport
	Scissors  []Rect2D
}

func (g *PipelineViewportStateCreateInfo) toC(c *C.VkPipelineViewportStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineViewportStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineViewportStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineViewportStateCreateFlags(temp_in_VkPipelineViewportStateCreateFlags)
	}
	c.viewportCount = C.uint32_t(len(g.Viewports))
	{
		c.pViewports = (*C.VkViewport)(_sa.alloc(C.sizeof_VkViewport * uint(len(g.Viewports))))
		slice2 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(g.Viewports):len(g.Viewports)]
		for i2, _ := range g.Viewports {
			g.Viewports[i2].toC(&slice2[i2])
		}
	}
	c.scissorCount = C.uint32_t(len(g.Scissors))
	{
		c.pScissors = (*C.VkRect2D)(_sa.alloc(C.sizeof_VkRect2D * uint(len(g.Scissors))))
		slice2 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(g.Scissors):len(g.Scissors)]
		for i2, _ := range g.Scissors {
			g.Scissors[i2].toC(&slice2[i2])
		}
	}
}
func (g *PipelineViewportStateCreateInfo) fromC(c *C.VkPipelineViewportStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineViewportStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineViewportStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineViewportStateCreateFlags(temp_in_VkPipelineViewportStateCreateFlags)
	}
	g.Viewports = make([]Viewport, int(c.viewportCount))
	{
		slice2 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(g.Viewports):len(g.Viewports)]
		for i2, _ := range g.Viewports {
			g.Viewports[i2].fromC(&slice2[i2])
		}
	}
	g.Scissors = make([]Rect2D, int(c.scissorCount))
	{
		slice2 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(g.Scissors):len(g.Scissors)]
		for i2, _ := range g.Scissors {
			g.Scissors[i2].fromC(&slice2[i2])
		}
	}
}

type PipelineRasterizationStateCreateFlags Flags
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

type CullModeFlags Flags
type FrontFace int

const (
	FRONT_FACE_COUNTER_CLOCKWISE FrontFace = 0
	FRONT_FACE_CLOCKWISE         FrontFace = 1
	FRONT_FACE_BEGIN_RANGE       FrontFace = FRONT_FACE_COUNTER_CLOCKWISE
	FRONT_FACE_END_RANGE         FrontFace = FRONT_FACE_CLOCKWISE
	FRONT_FACE_RANGE_SIZE        FrontFace = (FRONT_FACE_CLOCKWISE - FRONT_FACE_COUNTER_CLOCKWISE + 1)
	FRONT_FACE_MAX_ENUM          FrontFace = 2147483647
)

type PipelineRasterizationStateCreateInfo struct {
	Type                    StructureType
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

func (g *PipelineRasterizationStateCreateInfo) toC(c *C.VkPipelineRasterizationStateCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineRasterizationStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineRasterizationStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineRasterizationStateCreateFlags(temp_in_VkPipelineRasterizationStateCreateFlags)
	}
	if g.DepthClampEnable {
		c.depthClampEnable = 1
	} else {
		c.depthClampEnable = 0
	}
	if g.RasterizerDiscardEnable {
		c.rasterizerDiscardEnable = 1
	} else {
		c.rasterizerDiscardEnable = 0
	}
	c.polygonMode = C.VkPolygonMode(g.PolygonMode)
	{
		var temp_in_VkCullModeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.CullMode)))
			temp_in_VkCullModeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.cullMode = C.VkCullModeFlags(temp_in_VkCullModeFlags)
	}
	c.frontFace = C.VkFrontFace(g.FrontFace)
	if g.DepthBiasEnable {
		c.depthBiasEnable = 1
	} else {
		c.depthBiasEnable = 0
	}
	c.depthBiasConstantFactor = C.float(g.DepthBiasConstantFactor)
	c.depthBiasClamp = C.float(g.DepthBiasClamp)
	c.depthBiasSlopeFactor = C.float(g.DepthBiasSlopeFactor)
	c.lineWidth = C.float(g.LineWidth)
}
func (g *PipelineRasterizationStateCreateInfo) fromC(c *C.VkPipelineRasterizationStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineRasterizationStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineRasterizationStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineRasterizationStateCreateFlags(temp_in_VkPipelineRasterizationStateCreateFlags)
	}
	g.DepthClampEnable = c.depthClampEnable != 0
	g.RasterizerDiscardEnable = c.rasterizerDiscardEnable != 0
	g.PolygonMode = PolygonMode(c.polygonMode)
	{
		var temp_in_VkCullModeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.cullMode)))
			temp_in_VkCullModeFlags = Flags(temp_in_VkFlags)
		}
		g.CullMode = CullModeFlags(temp_in_VkCullModeFlags)
	}
	g.FrontFace = FrontFace(c.frontFace)
	g.DepthBiasEnable = c.depthBiasEnable != 0
	g.DepthBiasConstantFactor = float32(c.depthBiasConstantFactor)
	g.DepthBiasClamp = float32(c.depthBiasClamp)
	g.DepthBiasSlopeFactor = float32(c.depthBiasSlopeFactor)
	g.LineWidth = float32(c.lineWidth)
}

type PipelineMultisampleStateCreateFlags Flags
type SampleMask uint32
type PipelineMultisampleStateCreateInfo struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	Flags                 PipelineMultisampleStateCreateFlags
	RasterizationSamples  SampleCountFlagBits
	SampleShadingEnable   bool
	MinSampleShading      float32
	SampleMask            *SampleMask
	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool
}

func (g *PipelineMultisampleStateCreateInfo) toC(c *C.VkPipelineMultisampleStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineMultisampleStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineMultisampleStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineMultisampleStateCreateFlags(temp_in_VkPipelineMultisampleStateCreateFlags)
	}
	c.rasterizationSamples = C.VkSampleCountFlagBits(g.RasterizationSamples)
	if g.SampleShadingEnable {
		c.sampleShadingEnable = 1
	} else {
		c.sampleShadingEnable = 0
	}
	c.minSampleShading = C.float(g.MinSampleShading)
	{
		c.pSampleMask = (*C.VkSampleMask)(_sa.alloc(C.sizeof_VkSampleMask))
		{
			var temp_in_VkSampleMask C.uint32_t
			temp_in_VkSampleMask = C.uint32_t((uint32)(*g.SampleMask))
			*c.pSampleMask = C.VkSampleMask(temp_in_VkSampleMask)
		}
	}
	if g.AlphaToCoverageEnable {
		c.alphaToCoverageEnable = 1
	} else {
		c.alphaToCoverageEnable = 0
	}
	if g.AlphaToOneEnable {
		c.alphaToOneEnable = 1
	} else {
		c.alphaToOneEnable = 0
	}
}
func (g *PipelineMultisampleStateCreateInfo) fromC(c *C.VkPipelineMultisampleStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineMultisampleStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineMultisampleStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineMultisampleStateCreateFlags(temp_in_VkPipelineMultisampleStateCreateFlags)
	}
	g.RasterizationSamples = SampleCountFlagBits(c.rasterizationSamples)
	g.SampleShadingEnable = c.sampleShadingEnable != 0
	g.MinSampleShading = float32(c.minSampleShading)
	{
		if g.SampleMask == nil {
			g.SampleMask = new(SampleMask)
		}
		{
			var temp_in_VkSampleMask uint32
			temp_in_VkSampleMask = uint32((C.uint32_t)(*c.pSampleMask))
			*g.SampleMask = SampleMask(temp_in_VkSampleMask)
		}
	}
	g.AlphaToCoverageEnable = c.alphaToCoverageEnable != 0
	g.AlphaToOneEnable = c.alphaToOneEnable != 0
}

type PipelineDepthStencilStateCreateFlags Flags
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

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

func (g *StencilOpState) toC(c *C.VkStencilOpState) {
	c.failOp = C.VkStencilOp(g.FailOp)
	c.passOp = C.VkStencilOp(g.PassOp)
	c.depthFailOp = C.VkStencilOp(g.DepthFailOp)
	c.compareOp = C.VkCompareOp(g.CompareOp)
	c.compareMask = C.uint32_t(g.CompareMask)
	c.writeMask = C.uint32_t(g.WriteMask)
	c.reference = C.uint32_t(g.Reference)
}
func (g *StencilOpState) fromC(c *C.VkStencilOpState) {
	g.FailOp = StencilOp(c.failOp)
	g.PassOp = StencilOp(c.passOp)
	g.DepthFailOp = StencilOp(c.depthFailOp)
	g.CompareOp = CompareOp(c.compareOp)
	g.CompareMask = uint32(c.compareMask)
	g.WriteMask = uint32(c.writeMask)
	g.Reference = uint32(c.reference)
}

type PipelineDepthStencilStateCreateInfo struct {
	Type                  StructureType
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

func (g *PipelineDepthStencilStateCreateInfo) toC(c *C.VkPipelineDepthStencilStateCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineDepthStencilStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineDepthStencilStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineDepthStencilStateCreateFlags(temp_in_VkPipelineDepthStencilStateCreateFlags)
	}
	if g.DepthTestEnable {
		c.depthTestEnable = 1
	} else {
		c.depthTestEnable = 0
	}
	if g.DepthWriteEnable {
		c.depthWriteEnable = 1
	} else {
		c.depthWriteEnable = 0
	}
	c.depthCompareOp = C.VkCompareOp(g.DepthCompareOp)
	if g.DepthBoundsTestEnable {
		c.depthBoundsTestEnable = 1
	} else {
		c.depthBoundsTestEnable = 0
	}
	if g.StencilTestEnable {
		c.stencilTestEnable = 1
	} else {
		c.stencilTestEnable = 0
	}
	g.Front.toC(&c.front)
	g.Back.toC(&c.back)
	c.minDepthBounds = C.float(g.MinDepthBounds)
	c.maxDepthBounds = C.float(g.MaxDepthBounds)
}
func (g *PipelineDepthStencilStateCreateInfo) fromC(c *C.VkPipelineDepthStencilStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineDepthStencilStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineDepthStencilStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineDepthStencilStateCreateFlags(temp_in_VkPipelineDepthStencilStateCreateFlags)
	}
	g.DepthTestEnable = c.depthTestEnable != 0
	g.DepthWriteEnable = c.depthWriteEnable != 0
	g.DepthCompareOp = CompareOp(c.depthCompareOp)
	g.DepthBoundsTestEnable = c.depthBoundsTestEnable != 0
	g.StencilTestEnable = c.stencilTestEnable != 0
	g.Front.fromC(&c.front)
	g.Back.fromC(&c.back)
	g.MinDepthBounds = float32(c.minDepthBounds)
	g.MaxDepthBounds = float32(c.maxDepthBounds)
}

type PipelineColorBlendStateCreateFlags Flags
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

type ColorComponentFlags Flags
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

func (g *PipelineColorBlendAttachmentState) toC(c *C.VkPipelineColorBlendAttachmentState) {
	if g.BlendEnable {
		c.blendEnable = 1
	} else {
		c.blendEnable = 0
	}
	c.srcColorBlendFactor = C.VkBlendFactor(g.SrcColorBlendFactor)
	c.dstColorBlendFactor = C.VkBlendFactor(g.DstColorBlendFactor)
	c.colorBlendOp = C.VkBlendOp(g.ColorBlendOp)
	c.srcAlphaBlendFactor = C.VkBlendFactor(g.SrcAlphaBlendFactor)
	c.dstAlphaBlendFactor = C.VkBlendFactor(g.DstAlphaBlendFactor)
	c.alphaBlendOp = C.VkBlendOp(g.AlphaBlendOp)
	{
		var temp_in_VkColorComponentFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ColorWriteMask)))
			temp_in_VkColorComponentFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.colorWriteMask = C.VkColorComponentFlags(temp_in_VkColorComponentFlags)
	}
}
func (g *PipelineColorBlendAttachmentState) fromC(c *C.VkPipelineColorBlendAttachmentState) {
	g.BlendEnable = c.blendEnable != 0
	g.SrcColorBlendFactor = BlendFactor(c.srcColorBlendFactor)
	g.DstColorBlendFactor = BlendFactor(c.dstColorBlendFactor)
	g.ColorBlendOp = BlendOp(c.colorBlendOp)
	g.SrcAlphaBlendFactor = BlendFactor(c.srcAlphaBlendFactor)
	g.DstAlphaBlendFactor = BlendFactor(c.dstAlphaBlendFactor)
	g.AlphaBlendOp = BlendOp(c.alphaBlendOp)
	{
		var temp_in_VkColorComponentFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.colorWriteMask)))
			temp_in_VkColorComponentFlags = Flags(temp_in_VkFlags)
		}
		g.ColorWriteMask = ColorComponentFlags(temp_in_VkColorComponentFlags)
	}
}

type PipelineColorBlendStateCreateInfo struct {
	Type           StructureType
	Next           unsafe.Pointer
	Flags          PipelineColorBlendStateCreateFlags
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (g *PipelineColorBlendStateCreateInfo) toC(c *C.VkPipelineColorBlendStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineColorBlendStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineColorBlendStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineColorBlendStateCreateFlags(temp_in_VkPipelineColorBlendStateCreateFlags)
	}
	if g.LogicOpEnable {
		c.logicOpEnable = 1
	} else {
		c.logicOpEnable = 0
	}
	c.logicOp = C.VkLogicOp(g.LogicOp)
	c.attachmentCount = C.uint32_t(len(g.Attachments))
	{
		c.pAttachments = (*C.VkPipelineColorBlendAttachmentState)(_sa.alloc(C.sizeof_VkPipelineColorBlendAttachmentState * uint(len(g.Attachments))))
		slice2 := (*[1 << 31]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].toC(&slice2[i2])
		}
	}
	for i, _ := range g.BlendConstants {
		c.blendConstants[i] = C.float(g.BlendConstants[i])
	}
}
func (g *PipelineColorBlendStateCreateInfo) fromC(c *C.VkPipelineColorBlendStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineColorBlendStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineColorBlendStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineColorBlendStateCreateFlags(temp_in_VkPipelineColorBlendStateCreateFlags)
	}
	g.LogicOpEnable = c.logicOpEnable != 0
	g.LogicOp = LogicOp(c.logicOp)
	g.Attachments = make([]PipelineColorBlendAttachmentState, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].fromC(&slice2[i2])
		}
	}
	for i, _ := range g.BlendConstants {
		g.BlendConstants[i] = float32(c.blendConstants[i])
	}
}

type PipelineDynamicStateCreateFlags Flags
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

type PipelineDynamicStateCreateInfo struct {
	Type          StructureType
	Next          unsafe.Pointer
	Flags         PipelineDynamicStateCreateFlags
	DynamicStates []DynamicState
}

func (g *PipelineDynamicStateCreateInfo) toC(c *C.VkPipelineDynamicStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineDynamicStateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineDynamicStateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineDynamicStateCreateFlags(temp_in_VkPipelineDynamicStateCreateFlags)
	}
	c.dynamicStateCount = C.uint32_t(len(g.DynamicStates))
	{
		c.pDynamicStates = (*C.VkDynamicState)(_sa.alloc(C.sizeof_VkDynamicState * uint(len(g.DynamicStates))))
		slice2 := (*[1 << 31]C.VkDynamicState)(unsafe.Pointer(c.pDynamicStates))[:len(g.DynamicStates):len(g.DynamicStates)]
		for i2, _ := range g.DynamicStates {
			slice2[i2] = C.VkDynamicState(g.DynamicStates[i2])
		}
	}
}
func (g *PipelineDynamicStateCreateInfo) fromC(c *C.VkPipelineDynamicStateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineDynamicStateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineDynamicStateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineDynamicStateCreateFlags(temp_in_VkPipelineDynamicStateCreateFlags)
	}
	g.DynamicStates = make([]DynamicState, int(c.dynamicStateCount))
	{
		slice2 := (*[1 << 31]C.VkDynamicState)(unsafe.Pointer(c.pDynamicStates))[:len(g.DynamicStates):len(g.DynamicStates)]
		for i2, _ := range g.DynamicStates {
			g.DynamicStates[i2] = DynamicState(slice2[i2])
		}
	}
}

type PipelineLayout C.VkPipelineLayout
type RenderPass C.VkRenderPass
type Pipeline C.VkPipeline
type GraphicsPipelineCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              PipelineCreateFlags
	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func (g *GraphicsPipelineCreateInfo) toC(c *C.VkGraphicsPipelineCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineCreateFlags(temp_in_VkPipelineCreateFlags)
	}
	c.stageCount = C.uint32_t(len(g.Stages))
	{
		c.pStages = (*C.VkPipelineShaderStageCreateInfo)(_sa.alloc(C.sizeof_VkPipelineShaderStageCreateInfo * uint(len(g.Stages))))
		slice2 := (*[1 << 31]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(c.pStages))[:len(g.Stages):len(g.Stages)]
		for i2, _ := range g.Stages {
			g.Stages[i2].toC(&slice2[i2], _sa)
		}
	}
	{
		c.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineVertexInputStateCreateInfo))
		g.VertexInputState.toC(c.pVertexInputState, _sa)
	}
	{
		c.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineInputAssemblyStateCreateInfo))
		g.InputAssemblyState.toC(c.pInputAssemblyState)
	}
	{
		c.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineTessellationStateCreateInfo))
		g.TessellationState.toC(c.pTessellationState)
	}
	{
		c.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineViewportStateCreateInfo))
		g.ViewportState.toC(c.pViewportState, _sa)
	}
	{
		c.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineRasterizationStateCreateInfo))
		g.RasterizationState.toC(c.pRasterizationState)
	}
	{
		c.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineMultisampleStateCreateInfo))
		g.MultisampleState.toC(c.pMultisampleState, _sa)
	}
	{
		c.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineDepthStencilStateCreateInfo))
		g.DepthStencilState.toC(c.pDepthStencilState)
	}
	{
		c.pColorBlendState = (*C.VkPipelineColorBlendStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineColorBlendStateCreateInfo))
		g.ColorBlendState.toC(c.pColorBlendState, _sa)
	}
	{
		c.pDynamicState = (*C.VkPipelineDynamicStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineDynamicStateCreateInfo))
		g.DynamicState.toC(c.pDynamicState, _sa)
	}
	c.layout = C.VkPipelineLayout(g.Layout)
	c.renderPass = C.VkRenderPass(g.RenderPass)
	c.subpass = C.uint32_t(g.Subpass)
	c.basePipelineHandle = C.VkPipeline(g.BasePipelineHandle)
	c.basePipelineIndex = C.int32_t(g.BasePipelineIndex)
}
func (g *GraphicsPipelineCreateInfo) fromC(c *C.VkGraphicsPipelineCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineCreateFlags(temp_in_VkPipelineCreateFlags)
	}
	g.Stages = make([]PipelineShaderStageCreateInfo, int(c.stageCount))
	{
		slice2 := (*[1 << 31]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(c.pStages))[:len(g.Stages):len(g.Stages)]
		for i2, _ := range g.Stages {
			g.Stages[i2].fromC(&slice2[i2])
		}
	}
	{
		if g.VertexInputState == nil {
			g.VertexInputState = new(PipelineVertexInputStateCreateInfo)
		}
		g.VertexInputState.fromC(c.pVertexInputState)
	}
	{
		if g.InputAssemblyState == nil {
			g.InputAssemblyState = new(PipelineInputAssemblyStateCreateInfo)
		}
		g.InputAssemblyState.fromC(c.pInputAssemblyState)
	}
	{
		if g.TessellationState == nil {
			g.TessellationState = new(PipelineTessellationStateCreateInfo)
		}
		g.TessellationState.fromC(c.pTessellationState)
	}
	{
		if g.ViewportState == nil {
			g.ViewportState = new(PipelineViewportStateCreateInfo)
		}
		g.ViewportState.fromC(c.pViewportState)
	}
	{
		if g.RasterizationState == nil {
			g.RasterizationState = new(PipelineRasterizationStateCreateInfo)
		}
		g.RasterizationState.fromC(c.pRasterizationState)
	}
	{
		if g.MultisampleState == nil {
			g.MultisampleState = new(PipelineMultisampleStateCreateInfo)
		}
		g.MultisampleState.fromC(c.pMultisampleState)
	}
	{
		if g.DepthStencilState == nil {
			g.DepthStencilState = new(PipelineDepthStencilStateCreateInfo)
		}
		g.DepthStencilState.fromC(c.pDepthStencilState)
	}
	{
		if g.ColorBlendState == nil {
			g.ColorBlendState = new(PipelineColorBlendStateCreateInfo)
		}
		g.ColorBlendState.fromC(c.pColorBlendState)
	}
	{
		if g.DynamicState == nil {
			g.DynamicState = new(PipelineDynamicStateCreateInfo)
		}
		g.DynamicState.fromC(c.pDynamicState)
	}
	g.Layout = PipelineLayout(c.layout)
	g.RenderPass = RenderPass(c.renderPass)
	g.Subpass = uint32(c.subpass)
	g.BasePipelineHandle = Pipeline(c.basePipelineHandle)
	g.BasePipelineIndex = int32(c.basePipelineIndex)
}
func CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, createInfos []GraphicsPipelineCreateInfo, allocator *AllocationCallbacks, pipelines []Pipeline) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pipelineCache   C.VkPipelineCache
		createInfoCount C.uint32_t
		pCreateInfos    *C.VkGraphicsPipelineCreateInfo
		pAllocator      *C.VkAllocationCallbacks
		pPipelines      *C.VkPipeline
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipelineCache = C.VkPipelineCache(pipelineCache)
	c.createInfoCount = C.uint32_t(len(createInfos))
	{
		c.pCreateInfos = (*C.VkGraphicsPipelineCreateInfo)(_sa.alloc(C.sizeof_VkGraphicsPipelineCreateInfo * uint(len(createInfos))))
		slice3 := (*[1 << 31]C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(c.pCreateInfos))[:len(createInfos):len(createInfos)]
		for i3, _ := range createInfos {
			createInfos[i3].toC(&slice3[i3], _sa)
		}
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelines = (*C.VkPipeline)(_sa.alloc(C.sizeof_VkPipeline * uint(len(pipelines))))
		slice3 := (*[1 << 31]C.VkPipeline)(unsafe.Pointer(c.pPipelines))[:len(pipelines):len(pipelines)]
		for i3, _ := range pipelines {
			slice3[i3] = C.VkPipeline(pipelines[i3])
		}
	}
	c._ret = C.vkCreateGraphicsPipelines(c.device, c.pipelineCache, c.createInfoCount, c.pCreateInfos, c.pAllocator, c.pPipelines)
	_ret = Result(c._ret)
	return
}

type ComputePipelineCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              PipelineCreateFlags
	Stage              PipelineShaderStageCreateInfo
	Layout             PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func (g *ComputePipelineCreateInfo) toC(c *C.VkComputePipelineCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineCreateFlags(temp_in_VkPipelineCreateFlags)
	}
	g.Stage.toC(&c.stage, _sa)
	c.layout = C.VkPipelineLayout(g.Layout)
	c.basePipelineHandle = C.VkPipeline(g.BasePipelineHandle)
	c.basePipelineIndex = C.int32_t(g.BasePipelineIndex)
}
func (g *ComputePipelineCreateInfo) fromC(c *C.VkComputePipelineCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineCreateFlags(temp_in_VkPipelineCreateFlags)
	}
	g.Stage.fromC(&c.stage)
	g.Layout = PipelineLayout(c.layout)
	g.BasePipelineHandle = Pipeline(c.basePipelineHandle)
	g.BasePipelineIndex = int32(c.basePipelineIndex)
}
func CreateComputePipelines(device Device, pipelineCache PipelineCache, createInfos []ComputePipelineCreateInfo, allocator *AllocationCallbacks, pipelines []Pipeline) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pipelineCache   C.VkPipelineCache
		createInfoCount C.uint32_t
		pCreateInfos    *C.VkComputePipelineCreateInfo
		pAllocator      *C.VkAllocationCallbacks
		pPipelines      *C.VkPipeline
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipelineCache = C.VkPipelineCache(pipelineCache)
	c.createInfoCount = C.uint32_t(len(createInfos))
	{
		c.pCreateInfos = (*C.VkComputePipelineCreateInfo)(_sa.alloc(C.sizeof_VkComputePipelineCreateInfo * uint(len(createInfos))))
		slice3 := (*[1 << 31]C.VkComputePipelineCreateInfo)(unsafe.Pointer(c.pCreateInfos))[:len(createInfos):len(createInfos)]
		for i3, _ := range createInfos {
			createInfos[i3].toC(&slice3[i3], _sa)
		}
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelines = (*C.VkPipeline)(_sa.alloc(C.sizeof_VkPipeline * uint(len(pipelines))))
		slice3 := (*[1 << 31]C.VkPipeline)(unsafe.Pointer(c.pPipelines))[:len(pipelines):len(pipelines)]
		for i3, _ := range pipelines {
			slice3[i3] = C.VkPipeline(pipelines[i3])
		}
	}
	c._ret = C.vkCreateComputePipelines(c.device, c.pipelineCache, c.createInfoCount, c.pCreateInfos, c.pAllocator, c.pPipelines)
	_ret = Result(c._ret)
	return
}
func DestroyPipeline(device Device, pipeline Pipeline, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		pipeline   C.VkPipeline
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipeline = C.VkPipeline(pipeline)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipeline(c.device, c.pipeline, c.pAllocator)
}

type PipelineLayoutCreateFlags Flags
type DescriptorSetLayout C.VkDescriptorSetLayout
type ShaderStageFlags Flags
type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}

func (g *PushConstantRange) toC(c *C.VkPushConstantRange) {
	{
		var temp_in_VkShaderStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.StageFlags)))
			temp_in_VkShaderStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.stageFlags = C.VkShaderStageFlags(temp_in_VkShaderStageFlags)
	}
	c.offset = C.uint32_t(g.Offset)
	c.size = C.uint32_t(g.Size)
}
func (g *PushConstantRange) fromC(c *C.VkPushConstantRange) {
	{
		var temp_in_VkShaderStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.stageFlags)))
			temp_in_VkShaderStageFlags = Flags(temp_in_VkFlags)
		}
		g.StageFlags = ShaderStageFlags(temp_in_VkShaderStageFlags)
	}
	g.Offset = uint32(c.offset)
	g.Size = uint32(c.size)
}

type PipelineLayoutCreateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              PipelineLayoutCreateFlags
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (g *PipelineLayoutCreateInfo) toC(c *C.VkPipelineLayoutCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkPipelineLayoutCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkPipelineLayoutCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkPipelineLayoutCreateFlags(temp_in_VkPipelineLayoutCreateFlags)
	}
	c.setLayoutCount = C.uint32_t(len(g.SetLayouts))
	{
		c.pSetLayouts = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout * uint(len(g.SetLayouts))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.SetLayouts):len(g.SetLayouts)]
		for i2, _ := range g.SetLayouts {
			slice2[i2] = C.VkDescriptorSetLayout(g.SetLayouts[i2])
		}
	}
	c.pushConstantRangeCount = C.uint32_t(len(g.PushConstantRanges))
	{
		c.pPushConstantRanges = (*C.VkPushConstantRange)(_sa.alloc(C.sizeof_VkPushConstantRange * uint(len(g.PushConstantRanges))))
		slice2 := (*[1 << 31]C.VkPushConstantRange)(unsafe.Pointer(c.pPushConstantRanges))[:len(g.PushConstantRanges):len(g.PushConstantRanges)]
		for i2, _ := range g.PushConstantRanges {
			g.PushConstantRanges[i2].toC(&slice2[i2])
		}
	}
}
func (g *PipelineLayoutCreateInfo) fromC(c *C.VkPipelineLayoutCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkPipelineLayoutCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkPipelineLayoutCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = PipelineLayoutCreateFlags(temp_in_VkPipelineLayoutCreateFlags)
	}
	g.SetLayouts = make([]DescriptorSetLayout, int(c.setLayoutCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.SetLayouts):len(g.SetLayouts)]
		for i2, _ := range g.SetLayouts {
			g.SetLayouts[i2] = DescriptorSetLayout(slice2[i2])
		}
	}
	g.PushConstantRanges = make([]PushConstantRange, int(c.pushConstantRangeCount))
	{
		slice2 := (*[1 << 31]C.VkPushConstantRange)(unsafe.Pointer(c.pPushConstantRanges))[:len(g.PushConstantRanges):len(g.PushConstantRanges)]
		for i2, _ := range g.PushConstantRanges {
			g.PushConstantRanges[i2].fromC(&slice2[i2])
		}
	}
}
func CreatePipelineLayout(device Device, createInfo *PipelineLayoutCreateInfo, allocator *AllocationCallbacks, pipelineLayout *PipelineLayout) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pCreateInfo     *C.VkPipelineLayoutCreateInfo
		pAllocator      *C.VkAllocationCallbacks
		pPipelineLayout *C.VkPipelineLayout
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkPipelineLayoutCreateInfo)(_sa.alloc(C.sizeof_VkPipelineLayoutCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelineLayout = (*C.VkPipelineLayout)(_sa.alloc(C.sizeof_VkPipelineLayout))
		*c.pPipelineLayout = C.VkPipelineLayout(*pipelineLayout)
	}
	c._ret = C.vkCreatePipelineLayout(c.device, c.pCreateInfo, c.pAllocator, c.pPipelineLayout)
	_ret = Result(c._ret)
	*pipelineLayout = PipelineLayout(*c.pPipelineLayout)
	return
}
func DestroyPipelineLayout(device Device, pipelineLayout PipelineLayout, allocator *AllocationCallbacks) {
	var c struct {
		device         C.VkDevice
		pipelineLayout C.VkPipelineLayout
		pAllocator     *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pipelineLayout = C.VkPipelineLayout(pipelineLayout)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipelineLayout(c.device, c.pipelineLayout, c.pAllocator)
}

type SamplerCreateFlags Flags
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

type SamplerCreateInfo struct {
	Type                    StructureType
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

func (g *SamplerCreateInfo) toC(c *C.VkSamplerCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkSamplerCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSamplerCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSamplerCreateFlags(temp_in_VkSamplerCreateFlags)
	}
	c.magFilter = C.VkFilter(g.MagFilter)
	c.minFilter = C.VkFilter(g.MinFilter)
	c.mipmapMode = C.VkSamplerMipmapMode(g.MipmapMode)
	c.addressModeU = C.VkSamplerAddressMode(g.AddressModeU)
	c.addressModeV = C.VkSamplerAddressMode(g.AddressModeV)
	c.addressModeW = C.VkSamplerAddressMode(g.AddressModeW)
	c.mipLodBias = C.float(g.MipLodBias)
	if g.AnisotropyEnable {
		c.anisotropyEnable = 1
	} else {
		c.anisotropyEnable = 0
	}
	c.maxAnisotropy = C.float(g.MaxAnisotropy)
	if g.CompareEnable {
		c.compareEnable = 1
	} else {
		c.compareEnable = 0
	}
	c.compareOp = C.VkCompareOp(g.CompareOp)
	c.minLod = C.float(g.MinLod)
	c.maxLod = C.float(g.MaxLod)
	c.borderColor = C.VkBorderColor(g.BorderColor)
	if g.UnnormalizedCoordinates {
		c.unnormalizedCoordinates = 1
	} else {
		c.unnormalizedCoordinates = 0
	}
}
func (g *SamplerCreateInfo) fromC(c *C.VkSamplerCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkSamplerCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSamplerCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SamplerCreateFlags(temp_in_VkSamplerCreateFlags)
	}
	g.MagFilter = Filter(c.magFilter)
	g.MinFilter = Filter(c.minFilter)
	g.MipmapMode = SamplerMipmapMode(c.mipmapMode)
	g.AddressModeU = SamplerAddressMode(c.addressModeU)
	g.AddressModeV = SamplerAddressMode(c.addressModeV)
	g.AddressModeW = SamplerAddressMode(c.addressModeW)
	g.MipLodBias = float32(c.mipLodBias)
	g.AnisotropyEnable = c.anisotropyEnable != 0
	g.MaxAnisotropy = float32(c.maxAnisotropy)
	g.CompareEnable = c.compareEnable != 0
	g.CompareOp = CompareOp(c.compareOp)
	g.MinLod = float32(c.minLod)
	g.MaxLod = float32(c.maxLod)
	g.BorderColor = BorderColor(c.borderColor)
	g.UnnormalizedCoordinates = c.unnormalizedCoordinates != 0
}

type Sampler C.VkSampler

func CreateSampler(device Device, createInfo *SamplerCreateInfo, allocator *AllocationCallbacks, sampler *Sampler) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkSamplerCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pSampler    *C.VkSampler
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkSamplerCreateInfo)(_sa.alloc(C.sizeof_VkSamplerCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSampler = (*C.VkSampler)(_sa.alloc(C.sizeof_VkSampler))
		*c.pSampler = C.VkSampler(*sampler)
	}
	c._ret = C.vkCreateSampler(c.device, c.pCreateInfo, c.pAllocator, c.pSampler)
	_ret = Result(c._ret)
	*sampler = Sampler(*c.pSampler)
	return
}
func DestroySampler(device Device, sampler Sampler, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		sampler    C.VkSampler
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.sampler = C.VkSampler(sampler)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySampler(c.device, c.sampler, c.pAllocator)
}

type DescriptorSetLayoutCreateFlags Flags
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

type DescriptorSetLayoutBinding struct {
	Binding           uint32
	DescriptorType    DescriptorType
	DescriptorCount   uint32
	StageFlags        ShaderStageFlags
	ImmutableSamplers []Sampler
}

func (g *DescriptorSetLayoutBinding) toC(c *C.VkDescriptorSetLayoutBinding, _sa *stackAllocator) {
	c.binding = C.uint32_t(g.Binding)
	c.descriptorType = C.VkDescriptorType(g.DescriptorType)
	c.descriptorCount = C.uint32_t(g.DescriptorCount)
	{
		var temp_in_VkShaderStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.StageFlags)))
			temp_in_VkShaderStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.stageFlags = C.VkShaderStageFlags(temp_in_VkShaderStageFlags)
	}
	{
		c.pImmutableSamplers = (*C.VkSampler)(_sa.alloc(C.sizeof_VkSampler * uint(len(g.ImmutableSamplers))))
		slice2 := (*[1 << 31]C.VkSampler)(unsafe.Pointer(c.pImmutableSamplers))[:len(g.ImmutableSamplers):len(g.ImmutableSamplers)]
		for i2, _ := range g.ImmutableSamplers {
			slice2[i2] = C.VkSampler(g.ImmutableSamplers[i2])
		}
	}
}
func (g *DescriptorSetLayoutBinding) fromC(c *C.VkDescriptorSetLayoutBinding) {
	g.Binding = uint32(c.binding)
	g.DescriptorType = DescriptorType(c.descriptorType)
	g.DescriptorCount = uint32(c.descriptorCount)
	{
		var temp_in_VkShaderStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.stageFlags)))
			temp_in_VkShaderStageFlags = Flags(temp_in_VkFlags)
		}
		g.StageFlags = ShaderStageFlags(temp_in_VkShaderStageFlags)
	}
	{
		slice2 := (*[1 << 31]C.VkSampler)(unsafe.Pointer(c.pImmutableSamplers))[:len(g.ImmutableSamplers):len(g.ImmutableSamplers)]
		for i2, _ := range g.ImmutableSamplers {
			g.ImmutableSamplers[i2] = Sampler(slice2[i2])
		}
	}
}

type DescriptorSetLayoutCreateInfo struct {
	Type     StructureType
	Next     unsafe.Pointer
	Flags    DescriptorSetLayoutCreateFlags
	Bindings []DescriptorSetLayoutBinding
}

func (g *DescriptorSetLayoutCreateInfo) toC(c *C.VkDescriptorSetLayoutCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDescriptorSetLayoutCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDescriptorSetLayoutCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDescriptorSetLayoutCreateFlags(temp_in_VkDescriptorSetLayoutCreateFlags)
	}
	c.bindingCount = C.uint32_t(len(g.Bindings))
	{
		c.pBindings = (*C.VkDescriptorSetLayoutBinding)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutBinding * uint(len(g.Bindings))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayoutBinding)(unsafe.Pointer(c.pBindings))[:len(g.Bindings):len(g.Bindings)]
		for i2, _ := range g.Bindings {
			g.Bindings[i2].toC(&slice2[i2], _sa)
		}
	}
}
func (g *DescriptorSetLayoutCreateInfo) fromC(c *C.VkDescriptorSetLayoutCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDescriptorSetLayoutCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDescriptorSetLayoutCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DescriptorSetLayoutCreateFlags(temp_in_VkDescriptorSetLayoutCreateFlags)
	}
	g.Bindings = make([]DescriptorSetLayoutBinding, int(c.bindingCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayoutBinding)(unsafe.Pointer(c.pBindings))[:len(g.Bindings):len(g.Bindings)]
		for i2, _ := range g.Bindings {
			g.Bindings[i2].fromC(&slice2[i2])
		}
	}
}
func CreateDescriptorSetLayout(device Device, createInfo *DescriptorSetLayoutCreateInfo, allocator *AllocationCallbacks, setLayout *DescriptorSetLayout) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkDescriptorSetLayoutCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pSetLayout  *C.VkDescriptorSetLayout
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorSetLayoutCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSetLayout = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout))
		*c.pSetLayout = C.VkDescriptorSetLayout(*setLayout)
	}
	c._ret = C.vkCreateDescriptorSetLayout(c.device, c.pCreateInfo, c.pAllocator, c.pSetLayout)
	_ret = Result(c._ret)
	*setLayout = DescriptorSetLayout(*c.pSetLayout)
	return
}
func DestroyDescriptorSetLayout(device Device, descriptorSetLayout DescriptorSetLayout, allocator *AllocationCallbacks) {
	var c struct {
		device              C.VkDevice
		descriptorSetLayout C.VkDescriptorSetLayout
		pAllocator          *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorSetLayout = C.VkDescriptorSetLayout(descriptorSetLayout)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorSetLayout(c.device, c.descriptorSetLayout, c.pAllocator)
}

type DescriptorPoolCreateFlags Flags
type DescriptorPoolSize struct {
	Type            DescriptorType
	DescriptorCount uint32
}

func (g *DescriptorPoolSize) toC(c *C.VkDescriptorPoolSize) {
	c._type = C.VkDescriptorType(g.Type)
	c.descriptorCount = C.uint32_t(g.DescriptorCount)
}
func (g *DescriptorPoolSize) fromC(c *C.VkDescriptorPoolSize) {
	g.Type = DescriptorType(c._type)
	g.DescriptorCount = uint32(c.descriptorCount)
}

type DescriptorPoolCreateInfo struct {
	Type      StructureType
	Next      unsafe.Pointer
	Flags     DescriptorPoolCreateFlags
	MaxSets   uint32
	PoolSizes []DescriptorPoolSize
}

func (g *DescriptorPoolCreateInfo) toC(c *C.VkDescriptorPoolCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDescriptorPoolCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDescriptorPoolCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDescriptorPoolCreateFlags(temp_in_VkDescriptorPoolCreateFlags)
	}
	c.maxSets = C.uint32_t(g.MaxSets)
	c.poolSizeCount = C.uint32_t(len(g.PoolSizes))
	{
		c.pPoolSizes = (*C.VkDescriptorPoolSize)(_sa.alloc(C.sizeof_VkDescriptorPoolSize * uint(len(g.PoolSizes))))
		slice2 := (*[1 << 31]C.VkDescriptorPoolSize)(unsafe.Pointer(c.pPoolSizes))[:len(g.PoolSizes):len(g.PoolSizes)]
		for i2, _ := range g.PoolSizes {
			g.PoolSizes[i2].toC(&slice2[i2])
		}
	}
}
func (g *DescriptorPoolCreateInfo) fromC(c *C.VkDescriptorPoolCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDescriptorPoolCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDescriptorPoolCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DescriptorPoolCreateFlags(temp_in_VkDescriptorPoolCreateFlags)
	}
	g.MaxSets = uint32(c.maxSets)
	g.PoolSizes = make([]DescriptorPoolSize, int(c.poolSizeCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorPoolSize)(unsafe.Pointer(c.pPoolSizes))[:len(g.PoolSizes):len(g.PoolSizes)]
		for i2, _ := range g.PoolSizes {
			g.PoolSizes[i2].fromC(&slice2[i2])
		}
	}
}

type DescriptorPool C.VkDescriptorPool

func CreateDescriptorPool(device Device, createInfo *DescriptorPoolCreateInfo, allocator *AllocationCallbacks, descriptorPool *DescriptorPool) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pCreateInfo     *C.VkDescriptorPoolCreateInfo
		pAllocator      *C.VkAllocationCallbacks
		pDescriptorPool *C.VkDescriptorPool
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorPoolCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorPoolCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDescriptorPool = (*C.VkDescriptorPool)(_sa.alloc(C.sizeof_VkDescriptorPool))
		*c.pDescriptorPool = C.VkDescriptorPool(*descriptorPool)
	}
	c._ret = C.vkCreateDescriptorPool(c.device, c.pCreateInfo, c.pAllocator, c.pDescriptorPool)
	_ret = Result(c._ret)
	*descriptorPool = DescriptorPool(*c.pDescriptorPool)
	return
}
func DestroyDescriptorPool(device Device, descriptorPool DescriptorPool, allocator *AllocationCallbacks) {
	var c struct {
		device         C.VkDevice
		descriptorPool C.VkDescriptorPool
		pAllocator     *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorPool = C.VkDescriptorPool(descriptorPool)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorPool(c.device, c.descriptorPool, c.pAllocator)
}

type DescriptorPoolResetFlags Flags

func ResetDescriptorPool(device Device, descriptorPool DescriptorPool, flags DescriptorPoolResetFlags) (_ret Result) {
	var c struct {
		device         C.VkDevice
		descriptorPool C.VkDescriptorPool
		flags          C.VkDescriptorPoolResetFlags
		_ret           C.VkResult
	}
	c.device = C.VkDevice(device)
	c.descriptorPool = C.VkDescriptorPool(descriptorPool)
	{
		var temp_in_VkDescriptorPoolResetFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkDescriptorPoolResetFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDescriptorPoolResetFlags(temp_in_VkDescriptorPoolResetFlags)
	}
	c._ret = C.vkResetDescriptorPool(c.device, c.descriptorPool, c.flags)
	_ret = Result(c._ret)
	return
}

type DescriptorSetAllocateInfo struct {
	Type           StructureType
	Next           unsafe.Pointer
	DescriptorPool DescriptorPool
	SetLayouts     []DescriptorSetLayout
}

func (g *DescriptorSetAllocateInfo) toC(c *C.VkDescriptorSetAllocateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.descriptorPool = C.VkDescriptorPool(g.DescriptorPool)
	c.descriptorSetCount = C.uint32_t(len(g.SetLayouts))
	{
		c.pSetLayouts = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout * uint(len(g.SetLayouts))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.SetLayouts):len(g.SetLayouts)]
		for i2, _ := range g.SetLayouts {
			slice2[i2] = C.VkDescriptorSetLayout(g.SetLayouts[i2])
		}
	}
}
func (g *DescriptorSetAllocateInfo) fromC(c *C.VkDescriptorSetAllocateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.DescriptorPool = DescriptorPool(c.descriptorPool)
	g.SetLayouts = make([]DescriptorSetLayout, int(c.descriptorSetCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.SetLayouts):len(g.SetLayouts)]
		for i2, _ := range g.SetLayouts {
			g.SetLayouts[i2] = DescriptorSetLayout(slice2[i2])
		}
	}
}

type DescriptorSet C.VkDescriptorSet

func AllocateDescriptorSets(device Device, allocateInfo *DescriptorSetAllocateInfo, descriptorSets []DescriptorSet) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pAllocateInfo   *C.VkDescriptorSetAllocateInfo
		pDescriptorSets *C.VkDescriptorSet
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAllocateInfo = (*C.VkDescriptorSetAllocateInfo)(_sa.alloc(C.sizeof_VkDescriptorSetAllocateInfo))
		allocateInfo.toC(c.pAllocateInfo, _sa)
	}
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(descriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(descriptorSets):len(descriptorSets)]
		for i3, _ := range descriptorSets {
			slice3[i3] = C.VkDescriptorSet(descriptorSets[i3])
		}
	}
	c._ret = C.vkAllocateDescriptorSets(c.device, c.pAllocateInfo, c.pDescriptorSets)
	_ret = Result(c._ret)
	return
}
func FreeDescriptorSets(device Device, descriptorPool DescriptorPool, descriptorSets []DescriptorSet) (_ret Result) {
	var c struct {
		device             C.VkDevice
		descriptorPool     C.VkDescriptorPool
		descriptorSetCount C.uint32_t
		pDescriptorSets    *C.VkDescriptorSet
		_ret               C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorPool = C.VkDescriptorPool(descriptorPool)
	c.descriptorSetCount = C.uint32_t(len(descriptorSets))
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(descriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(descriptorSets):len(descriptorSets)]
		for i3, _ := range descriptorSets {
			slice3[i3] = C.VkDescriptorSet(descriptorSets[i3])
		}
	}
	c._ret = C.vkFreeDescriptorSets(c.device, c.descriptorPool, c.descriptorSetCount, c.pDescriptorSets)
	_ret = Result(c._ret)
	return
}

type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}

func (g *DescriptorImageInfo) toC(c *C.VkDescriptorImageInfo) {
	c.sampler = C.VkSampler(g.Sampler)
	c.imageView = C.VkImageView(g.ImageView)
	c.imageLayout = C.VkImageLayout(g.ImageLayout)
}
func (g *DescriptorImageInfo) fromC(c *C.VkDescriptorImageInfo) {
	g.Sampler = Sampler(c.sampler)
	g.ImageView = ImageView(c.imageView)
	g.ImageLayout = ImageLayout(c.imageLayout)
}

type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset DeviceSize
	Range  DeviceSize
}

func (g *DescriptorBufferInfo) toC(c *C.VkDescriptorBufferInfo) {
	c.buffer = C.VkBuffer(g.Buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Range))
		c._range = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *DescriptorBufferInfo) fromC(c *C.VkDescriptorBufferInfo) {
	g.Buffer = Buffer(c.buffer)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.offset))
		g.Offset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c._range))
		g.Range = DeviceSize(temp_in_VkDeviceSize)
	}
}

type WriteDescriptorSet struct {
	Type            StructureType
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

func (g *WriteDescriptorSet) toC(c *C.VkWriteDescriptorSet, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.dstSet = C.VkDescriptorSet(g.DstSet)
	c.dstBinding = C.uint32_t(g.DstBinding)
	c.dstArrayElement = C.uint32_t(g.DstArrayElement)
	c.descriptorCount = C.uint32_t(g.DescriptorCount)
	c.descriptorType = C.VkDescriptorType(g.DescriptorType)
	{
		c.pImageInfo = (*C.VkDescriptorImageInfo)(_sa.alloc(C.sizeof_VkDescriptorImageInfo))
		g.ImageInfo.toC(c.pImageInfo)
	}
	{
		c.pBufferInfo = (*C.VkDescriptorBufferInfo)(_sa.alloc(C.sizeof_VkDescriptorBufferInfo))
		g.BufferInfo.toC(c.pBufferInfo)
	}
	{
		c.pTexelBufferView = (*C.VkBufferView)(_sa.alloc(C.sizeof_VkBufferView))
		*c.pTexelBufferView = C.VkBufferView(*g.TexelBufferView)
	}
}
func (g *WriteDescriptorSet) fromC(c *C.VkWriteDescriptorSet) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.DstSet = DescriptorSet(c.dstSet)
	g.DstBinding = uint32(c.dstBinding)
	g.DstArrayElement = uint32(c.dstArrayElement)
	g.DescriptorCount = uint32(c.descriptorCount)
	g.DescriptorType = DescriptorType(c.descriptorType)
	{
		if g.ImageInfo == nil {
			g.ImageInfo = new(DescriptorImageInfo)
		}
		g.ImageInfo.fromC(c.pImageInfo)
	}
	{
		if g.BufferInfo == nil {
			g.BufferInfo = new(DescriptorBufferInfo)
		}
		g.BufferInfo.fromC(c.pBufferInfo)
	}
	{
		if g.TexelBufferView == nil {
			g.TexelBufferView = new(BufferView)
		}
		*g.TexelBufferView = BufferView(*c.pTexelBufferView)
	}
}

type CopyDescriptorSet struct {
	Type            StructureType
	Next            unsafe.Pointer
	SrcSet          DescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

func (g *CopyDescriptorSet) toC(c *C.VkCopyDescriptorSet) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.srcSet = C.VkDescriptorSet(g.SrcSet)
	c.srcBinding = C.uint32_t(g.SrcBinding)
	c.srcArrayElement = C.uint32_t(g.SrcArrayElement)
	c.dstSet = C.VkDescriptorSet(g.DstSet)
	c.dstBinding = C.uint32_t(g.DstBinding)
	c.dstArrayElement = C.uint32_t(g.DstArrayElement)
	c.descriptorCount = C.uint32_t(g.DescriptorCount)
}
func (g *CopyDescriptorSet) fromC(c *C.VkCopyDescriptorSet) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.SrcSet = DescriptorSet(c.srcSet)
	g.SrcBinding = uint32(c.srcBinding)
	g.SrcArrayElement = uint32(c.srcArrayElement)
	g.DstSet = DescriptorSet(c.dstSet)
	g.DstBinding = uint32(c.dstBinding)
	g.DstArrayElement = uint32(c.dstArrayElement)
	g.DescriptorCount = uint32(c.descriptorCount)
}
func UpdateDescriptorSets(device Device, descriptorWrites []WriteDescriptorSet, descriptorCopies []CopyDescriptorSet) {
	var c struct {
		device               C.VkDevice
		descriptorWriteCount C.uint32_t
		pDescriptorWrites    *C.VkWriteDescriptorSet
		descriptorCopyCount  C.uint32_t
		pDescriptorCopies    *C.VkCopyDescriptorSet
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorWriteCount = C.uint32_t(len(descriptorWrites))
	{
		c.pDescriptorWrites = (*C.VkWriteDescriptorSet)(_sa.alloc(C.sizeof_VkWriteDescriptorSet * uint(len(descriptorWrites))))
		slice3 := (*[1 << 31]C.VkWriteDescriptorSet)(unsafe.Pointer(c.pDescriptorWrites))[:len(descriptorWrites):len(descriptorWrites)]
		for i3, _ := range descriptorWrites {
			descriptorWrites[i3].toC(&slice3[i3], _sa)
		}
	}
	c.descriptorCopyCount = C.uint32_t(len(descriptorCopies))
	{
		c.pDescriptorCopies = (*C.VkCopyDescriptorSet)(_sa.alloc(C.sizeof_VkCopyDescriptorSet * uint(len(descriptorCopies))))
		slice3 := (*[1 << 31]C.VkCopyDescriptorSet)(unsafe.Pointer(c.pDescriptorCopies))[:len(descriptorCopies):len(descriptorCopies)]
		for i3, _ := range descriptorCopies {
			descriptorCopies[i3].toC(&slice3[i3])
		}
	}
	C.vkUpdateDescriptorSets(c.device, c.descriptorWriteCount, c.pDescriptorWrites, c.descriptorCopyCount, c.pDescriptorCopies)
}

type FramebufferCreateFlags Flags
type FramebufferCreateInfo struct {
	Type        StructureType
	Next        unsafe.Pointer
	Flags       FramebufferCreateFlags
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

func (g *FramebufferCreateInfo) toC(c *C.VkFramebufferCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkFramebufferCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkFramebufferCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkFramebufferCreateFlags(temp_in_VkFramebufferCreateFlags)
	}
	c.renderPass = C.VkRenderPass(g.RenderPass)
	c.attachmentCount = C.uint32_t(len(g.Attachments))
	{
		c.pAttachments = (*C.VkImageView)(_sa.alloc(C.sizeof_VkImageView * uint(len(g.Attachments))))
		slice2 := (*[1 << 31]C.VkImageView)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			slice2[i2] = C.VkImageView(g.Attachments[i2])
		}
	}
	c.width = C.uint32_t(g.Width)
	c.height = C.uint32_t(g.Height)
	c.layers = C.uint32_t(g.Layers)
}
func (g *FramebufferCreateInfo) fromC(c *C.VkFramebufferCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkFramebufferCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkFramebufferCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = FramebufferCreateFlags(temp_in_VkFramebufferCreateFlags)
	}
	g.RenderPass = RenderPass(c.renderPass)
	g.Attachments = make([]ImageView, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkImageView)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2] = ImageView(slice2[i2])
		}
	}
	g.Width = uint32(c.width)
	g.Height = uint32(c.height)
	g.Layers = uint32(c.layers)
}

type Framebuffer C.VkFramebuffer

func CreateFramebuffer(device Device, createInfo *FramebufferCreateInfo, allocator *AllocationCallbacks, framebuffer *Framebuffer) (_ret Result) {
	var c struct {
		device       C.VkDevice
		pCreateInfo  *C.VkFramebufferCreateInfo
		pAllocator   *C.VkAllocationCallbacks
		pFramebuffer *C.VkFramebuffer
		_ret         C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkFramebufferCreateInfo)(_sa.alloc(C.sizeof_VkFramebufferCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pFramebuffer = (*C.VkFramebuffer)(_sa.alloc(C.sizeof_VkFramebuffer))
		*c.pFramebuffer = C.VkFramebuffer(*framebuffer)
	}
	c._ret = C.vkCreateFramebuffer(c.device, c.pCreateInfo, c.pAllocator, c.pFramebuffer)
	_ret = Result(c._ret)
	*framebuffer = Framebuffer(*c.pFramebuffer)
	return
}
func DestroyFramebuffer(device Device, framebuffer Framebuffer, allocator *AllocationCallbacks) {
	var c struct {
		device      C.VkDevice
		framebuffer C.VkFramebuffer
		pAllocator  *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.framebuffer = C.VkFramebuffer(framebuffer)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyFramebuffer(c.device, c.framebuffer, c.pAllocator)
}

type RenderPassCreateFlags Flags
type AttachmentDescriptionFlags Flags
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

type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlagBits
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func (g *AttachmentDescription) toC(c *C.VkAttachmentDescription) {
	{
		var temp_in_VkAttachmentDescriptionFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkAttachmentDescriptionFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkAttachmentDescriptionFlags(temp_in_VkAttachmentDescriptionFlags)
	}
	c.format = C.VkFormat(g.Format)
	c.samples = C.VkSampleCountFlagBits(g.Samples)
	c.loadOp = C.VkAttachmentLoadOp(g.LoadOp)
	c.storeOp = C.VkAttachmentStoreOp(g.StoreOp)
	c.stencilLoadOp = C.VkAttachmentLoadOp(g.StencilLoadOp)
	c.stencilStoreOp = C.VkAttachmentStoreOp(g.StencilStoreOp)
	c.initialLayout = C.VkImageLayout(g.InitialLayout)
	c.finalLayout = C.VkImageLayout(g.FinalLayout)
}
func (g *AttachmentDescription) fromC(c *C.VkAttachmentDescription) {
	{
		var temp_in_VkAttachmentDescriptionFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkAttachmentDescriptionFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = AttachmentDescriptionFlags(temp_in_VkAttachmentDescriptionFlags)
	}
	g.Format = Format(c.format)
	g.Samples = SampleCountFlagBits(c.samples)
	g.LoadOp = AttachmentLoadOp(c.loadOp)
	g.StoreOp = AttachmentStoreOp(c.storeOp)
	g.StencilLoadOp = AttachmentLoadOp(c.stencilLoadOp)
	g.StencilStoreOp = AttachmentStoreOp(c.stencilStoreOp)
	g.InitialLayout = ImageLayout(c.initialLayout)
	g.FinalLayout = ImageLayout(c.finalLayout)
}

type SubpassDescriptionFlags Flags
type PipelineBindPoint int

const (
	PIPELINE_BIND_POINT_GRAPHICS    PipelineBindPoint = 0
	PIPELINE_BIND_POINT_COMPUTE     PipelineBindPoint = 1
	PIPELINE_BIND_POINT_BEGIN_RANGE PipelineBindPoint = PIPELINE_BIND_POINT_GRAPHICS
	PIPELINE_BIND_POINT_END_RANGE   PipelineBindPoint = PIPELINE_BIND_POINT_COMPUTE
	PIPELINE_BIND_POINT_RANGE_SIZE  PipelineBindPoint = (PIPELINE_BIND_POINT_COMPUTE - PIPELINE_BIND_POINT_GRAPHICS + 1)
	PIPELINE_BIND_POINT_MAX_ENUM    PipelineBindPoint = 2147483647
)

type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
}

func (g *AttachmentReference) toC(c *C.VkAttachmentReference) {
	c.attachment = C.uint32_t(g.Attachment)
	c.layout = C.VkImageLayout(g.Layout)
}
func (g *AttachmentReference) fromC(c *C.VkAttachmentReference) {
	g.Attachment = uint32(c.attachment)
	g.Layout = ImageLayout(c.layout)
}

type SubpassDescription struct {
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment *AttachmentReference
	PreserveAttachments    []uint32
}

func (g *SubpassDescription) toC(c *C.VkSubpassDescription, _sa *stackAllocator) {
	{
		var temp_in_VkSubpassDescriptionFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSubpassDescriptionFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSubpassDescriptionFlags(temp_in_VkSubpassDescriptionFlags)
	}
	c.pipelineBindPoint = C.VkPipelineBindPoint(g.PipelineBindPoint)
	c.inputAttachmentCount = C.uint32_t(len(g.InputAttachments))
	{
		c.pInputAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.InputAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pInputAttachments))[:len(g.InputAttachments):len(g.InputAttachments)]
		for i2, _ := range g.InputAttachments {
			g.InputAttachments[i2].toC(&slice2[i2])
		}
	}
	c.colorAttachmentCount = C.uint32_t(len(g.ColorAttachments))
	{
		c.pColorAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.ColorAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pColorAttachments))[:len(g.ColorAttachments):len(g.ColorAttachments)]
		for i2, _ := range g.ColorAttachments {
			g.ColorAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pResolveAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.ResolveAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pResolveAttachments))[:len(g.ResolveAttachments):len(g.ResolveAttachments)]
		for i2, _ := range g.ResolveAttachments {
			g.ResolveAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pDepthStencilAttachment = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference))
		g.DepthStencilAttachment.toC(c.pDepthStencilAttachment)
	}
	c.preserveAttachmentCount = C.uint32_t(len(g.PreserveAttachments))
	{
		c.pPreserveAttachments = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.PreserveAttachments))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.PreserveAttachments):len(g.PreserveAttachments)]
		for i2, _ := range g.PreserveAttachments {
			slice2[i2] = C.uint32_t(g.PreserveAttachments[i2])
		}
	}
}
func (g *SubpassDescription) fromC(c *C.VkSubpassDescription) {
	{
		var temp_in_VkSubpassDescriptionFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSubpassDescriptionFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SubpassDescriptionFlags(temp_in_VkSubpassDescriptionFlags)
	}
	g.PipelineBindPoint = PipelineBindPoint(c.pipelineBindPoint)
	g.InputAttachments = make([]AttachmentReference, int(c.inputAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pInputAttachments))[:len(g.InputAttachments):len(g.InputAttachments)]
		for i2, _ := range g.InputAttachments {
			g.InputAttachments[i2].fromC(&slice2[i2])
		}
	}
	g.ColorAttachments = make([]AttachmentReference, int(c.colorAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pColorAttachments))[:len(g.ColorAttachments):len(g.ColorAttachments)]
		for i2, _ := range g.ColorAttachments {
			g.ColorAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pResolveAttachments))[:len(g.ResolveAttachments):len(g.ResolveAttachments)]
		for i2, _ := range g.ResolveAttachments {
			g.ResolveAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		if g.DepthStencilAttachment == nil {
			g.DepthStencilAttachment = new(AttachmentReference)
		}
		g.DepthStencilAttachment.fromC(c.pDepthStencilAttachment)
	}
	g.PreserveAttachments = make([]uint32, int(c.preserveAttachmentCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.PreserveAttachments):len(g.PreserveAttachments)]
		for i2, _ := range g.PreserveAttachments {
			g.PreserveAttachments[i2] = uint32(slice2[i2])
		}
	}
}

type AccessFlags Flags
type DependencyFlags Flags
type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}

func (g *SubpassDependency) toC(c *C.VkSubpassDependency) {
	c.srcSubpass = C.uint32_t(g.SrcSubpass)
	c.dstSubpass = C.uint32_t(g.DstSubpass)
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkDependencyFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DependencyFlags)))
			temp_in_VkDependencyFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dependencyFlags = C.VkDependencyFlags(temp_in_VkDependencyFlags)
	}
}
func (g *SubpassDependency) fromC(c *C.VkSubpassDependency) {
	g.SrcSubpass = uint32(c.srcSubpass)
	g.DstSubpass = uint32(c.dstSubpass)
	{
		var temp_in_VkPipelineStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcStageMask)))
			temp_in_VkPipelineStageFlags = Flags(temp_in_VkFlags)
		}
		g.SrcStageMask = PipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstStageMask)))
			temp_in_VkPipelineStageFlags = Flags(temp_in_VkFlags)
		}
		g.DstStageMask = PipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.SrcAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.DstAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkDependencyFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dependencyFlags)))
			temp_in_VkDependencyFlags = Flags(temp_in_VkFlags)
		}
		g.DependencyFlags = DependencyFlags(temp_in_VkDependencyFlags)
	}
}

type RenderPassCreateInfo struct {
	Type         StructureType
	Next         unsafe.Pointer
	Flags        RenderPassCreateFlags
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

func (g *RenderPassCreateInfo) toC(c *C.VkRenderPassCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkRenderPassCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkRenderPassCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkRenderPassCreateFlags(temp_in_VkRenderPassCreateFlags)
	}
	c.attachmentCount = C.uint32_t(len(g.Attachments))
	{
		c.pAttachments = (*C.VkAttachmentDescription)(_sa.alloc(C.sizeof_VkAttachmentDescription * uint(len(g.Attachments))))
		slice2 := (*[1 << 31]C.VkAttachmentDescription)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].toC(&slice2[i2])
		}
	}
	c.subpassCount = C.uint32_t(len(g.Subpasses))
	{
		c.pSubpasses = (*C.VkSubpassDescription)(_sa.alloc(C.sizeof_VkSubpassDescription * uint(len(g.Subpasses))))
		slice2 := (*[1 << 31]C.VkSubpassDescription)(unsafe.Pointer(c.pSubpasses))[:len(g.Subpasses):len(g.Subpasses)]
		for i2, _ := range g.Subpasses {
			g.Subpasses[i2].toC(&slice2[i2], _sa)
		}
	}
	c.dependencyCount = C.uint32_t(len(g.Dependencies))
	{
		c.pDependencies = (*C.VkSubpassDependency)(_sa.alloc(C.sizeof_VkSubpassDependency * uint(len(g.Dependencies))))
		slice2 := (*[1 << 31]C.VkSubpassDependency)(unsafe.Pointer(c.pDependencies))[:len(g.Dependencies):len(g.Dependencies)]
		for i2, _ := range g.Dependencies {
			g.Dependencies[i2].toC(&slice2[i2])
		}
	}
}
func (g *RenderPassCreateInfo) fromC(c *C.VkRenderPassCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkRenderPassCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkRenderPassCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = RenderPassCreateFlags(temp_in_VkRenderPassCreateFlags)
	}
	g.Attachments = make([]AttachmentDescription, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentDescription)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].fromC(&slice2[i2])
		}
	}
	g.Subpasses = make([]SubpassDescription, int(c.subpassCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDescription)(unsafe.Pointer(c.pSubpasses))[:len(g.Subpasses):len(g.Subpasses)]
		for i2, _ := range g.Subpasses {
			g.Subpasses[i2].fromC(&slice2[i2])
		}
	}
	g.Dependencies = make([]SubpassDependency, int(c.dependencyCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDependency)(unsafe.Pointer(c.pDependencies))[:len(g.Dependencies):len(g.Dependencies)]
		for i2, _ := range g.Dependencies {
			g.Dependencies[i2].fromC(&slice2[i2])
		}
	}
}
func CreateRenderPass(device Device, createInfo *RenderPassCreateInfo, allocator *AllocationCallbacks, renderPass *RenderPass) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkRenderPassCreateInfo
		pAllocator  *C.VkAllocationCallbacks
		pRenderPass *C.VkRenderPass
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkRenderPassCreateInfo)(_sa.alloc(C.sizeof_VkRenderPassCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pRenderPass = (*C.VkRenderPass)(_sa.alloc(C.sizeof_VkRenderPass))
		*c.pRenderPass = C.VkRenderPass(*renderPass)
	}
	c._ret = C.vkCreateRenderPass(c.device, c.pCreateInfo, c.pAllocator, c.pRenderPass)
	_ret = Result(c._ret)
	*renderPass = RenderPass(*c.pRenderPass)
	return
}
func DestroyRenderPass(device Device, renderPass RenderPass, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		renderPass C.VkRenderPass
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.renderPass = C.VkRenderPass(renderPass)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyRenderPass(c.device, c.renderPass, c.pAllocator)
}
func GetRenderAreaGranularity(device Device, renderPass RenderPass, granularity *Extent2D) {
	var c struct {
		device       C.VkDevice
		renderPass   C.VkRenderPass
		pGranularity *C.VkExtent2D
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.renderPass = C.VkRenderPass(renderPass)
	{
		c.pGranularity = (*C.VkExtent2D)(_sa.alloc(C.sizeof_VkExtent2D))
		granularity.toC(c.pGranularity)
	}
	C.vkGetRenderAreaGranularity(c.device, c.renderPass, c.pGranularity)
	granularity.fromC(c.pGranularity)
}

type CommandPoolCreateFlags Flags
type CommandPoolCreateInfo struct {
	Type             StructureType
	Next             unsafe.Pointer
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

func (g *CommandPoolCreateInfo) toC(c *C.VkCommandPoolCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkCommandPoolCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkCommandPoolCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandPoolCreateFlags(temp_in_VkCommandPoolCreateFlags)
	}
	c.queueFamilyIndex = C.uint32_t(g.QueueFamilyIndex)
}
func (g *CommandPoolCreateInfo) fromC(c *C.VkCommandPoolCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkCommandPoolCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkCommandPoolCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = CommandPoolCreateFlags(temp_in_VkCommandPoolCreateFlags)
	}
	g.QueueFamilyIndex = uint32(c.queueFamilyIndex)
}

type CommandPool C.VkCommandPool

func CreateCommandPool(device Device, createInfo *CommandPoolCreateInfo, allocator *AllocationCallbacks, commandPool *CommandPool) (_ret Result) {
	var c struct {
		device       C.VkDevice
		pCreateInfo  *C.VkCommandPoolCreateInfo
		pAllocator   *C.VkAllocationCallbacks
		pCommandPool *C.VkCommandPool
		_ret         C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkCommandPoolCreateInfo)(_sa.alloc(C.sizeof_VkCommandPoolCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pCommandPool = (*C.VkCommandPool)(_sa.alloc(C.sizeof_VkCommandPool))
		*c.pCommandPool = C.VkCommandPool(*commandPool)
	}
	c._ret = C.vkCreateCommandPool(c.device, c.pCreateInfo, c.pAllocator, c.pCommandPool)
	_ret = Result(c._ret)
	*commandPool = CommandPool(*c.pCommandPool)
	return
}
func DestroyCommandPool(device Device, commandPool CommandPool, allocator *AllocationCallbacks) {
	var c struct {
		device      C.VkDevice
		commandPool C.VkCommandPool
		pAllocator  *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyCommandPool(c.device, c.commandPool, c.pAllocator)
}

type CommandPoolResetFlags Flags

func ResetCommandPool(device Device, commandPool CommandPool, flags CommandPoolResetFlags) (_ret Result) {
	var c struct {
		device      C.VkDevice
		commandPool C.VkCommandPool
		flags       C.VkCommandPoolResetFlags
		_ret        C.VkResult
	}
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	{
		var temp_in_VkCommandPoolResetFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkCommandPoolResetFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandPoolResetFlags(temp_in_VkCommandPoolResetFlags)
	}
	c._ret = C.vkResetCommandPool(c.device, c.commandPool, c.flags)
	_ret = Result(c._ret)
	return
}

type CommandBufferLevel int

const (
	COMMAND_BUFFER_LEVEL_PRIMARY     CommandBufferLevel = 0
	COMMAND_BUFFER_LEVEL_SECONDARY   CommandBufferLevel = 1
	COMMAND_BUFFER_LEVEL_BEGIN_RANGE CommandBufferLevel = COMMAND_BUFFER_LEVEL_PRIMARY
	COMMAND_BUFFER_LEVEL_END_RANGE   CommandBufferLevel = COMMAND_BUFFER_LEVEL_SECONDARY
	COMMAND_BUFFER_LEVEL_RANGE_SIZE  CommandBufferLevel = (COMMAND_BUFFER_LEVEL_SECONDARY - COMMAND_BUFFER_LEVEL_PRIMARY + 1)
	COMMAND_BUFFER_LEVEL_MAX_ENUM    CommandBufferLevel = 2147483647
)

type CommandBufferAllocateInfo struct {
	Type               StructureType
	Next               unsafe.Pointer
	CommandPool        CommandPool
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

func (g *CommandBufferAllocateInfo) toC(c *C.VkCommandBufferAllocateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.commandPool = C.VkCommandPool(g.CommandPool)
	c.level = C.VkCommandBufferLevel(g.Level)
	c.commandBufferCount = C.uint32_t(g.CommandBufferCount)
}
func (g *CommandBufferAllocateInfo) fromC(c *C.VkCommandBufferAllocateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.CommandPool = CommandPool(c.commandPool)
	g.Level = CommandBufferLevel(c.level)
	g.CommandBufferCount = uint32(c.commandBufferCount)
}
func AllocateCommandBuffers(device Device, allocateInfo *CommandBufferAllocateInfo, commandBuffers []CommandBuffer) (_ret Result) {
	var c struct {
		device          C.VkDevice
		pAllocateInfo   *C.VkCommandBufferAllocateInfo
		pCommandBuffers *C.VkCommandBuffer
		_ret            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAllocateInfo = (*C.VkCommandBufferAllocateInfo)(_sa.alloc(C.sizeof_VkCommandBufferAllocateInfo))
		allocateInfo.toC(c.pAllocateInfo)
	}
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(commandBuffers))))
		slice3 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(commandBuffers):len(commandBuffers)]
		for i3, _ := range commandBuffers {
			slice3[i3] = C.VkCommandBuffer(commandBuffers[i3])
		}
	}
	c._ret = C.vkAllocateCommandBuffers(c.device, c.pAllocateInfo, c.pCommandBuffers)
	_ret = Result(c._ret)
	return
}
func FreeCommandBuffers(device Device, commandPool CommandPool, commandBuffers []CommandBuffer) {
	var c struct {
		device             C.VkDevice
		commandPool        C.VkCommandPool
		commandBufferCount C.uint32_t
		pCommandBuffers    *C.VkCommandBuffer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	c.commandBufferCount = C.uint32_t(len(commandBuffers))
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(commandBuffers))))
		slice3 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(commandBuffers):len(commandBuffers)]
		for i3, _ := range commandBuffers {
			slice3[i3] = C.VkCommandBuffer(commandBuffers[i3])
		}
	}
	C.vkFreeCommandBuffers(c.device, c.commandPool, c.commandBufferCount, c.pCommandBuffers)
}

type CommandBufferUsageFlags Flags
type QueryControlFlags Flags
type CommandBufferInheritanceInfo struct {
	Type                 StructureType
	Next                 unsafe.Pointer
	RenderPass           RenderPass
	Subpass              uint32
	Framebuffer          Framebuffer
	OcclusionQueryEnable bool
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags
}

func (g *CommandBufferInheritanceInfo) toC(c *C.VkCommandBufferInheritanceInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.renderPass = C.VkRenderPass(g.RenderPass)
	c.subpass = C.uint32_t(g.Subpass)
	c.framebuffer = C.VkFramebuffer(g.Framebuffer)
	if g.OcclusionQueryEnable {
		c.occlusionQueryEnable = 1
	} else {
		c.occlusionQueryEnable = 0
	}
	{
		var temp_in_VkQueryControlFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.QueryFlags)))
			temp_in_VkQueryControlFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.queryFlags = C.VkQueryControlFlags(temp_in_VkQueryControlFlags)
	}
	{
		var temp_in_VkQueryPipelineStatisticFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.PipelineStatistics)))
			temp_in_VkQueryPipelineStatisticFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.pipelineStatistics = C.VkQueryPipelineStatisticFlags(temp_in_VkQueryPipelineStatisticFlags)
	}
}
func (g *CommandBufferInheritanceInfo) fromC(c *C.VkCommandBufferInheritanceInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.RenderPass = RenderPass(c.renderPass)
	g.Subpass = uint32(c.subpass)
	g.Framebuffer = Framebuffer(c.framebuffer)
	g.OcclusionQueryEnable = c.occlusionQueryEnable != 0
	{
		var temp_in_VkQueryControlFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.queryFlags)))
			temp_in_VkQueryControlFlags = Flags(temp_in_VkFlags)
		}
		g.QueryFlags = QueryControlFlags(temp_in_VkQueryControlFlags)
	}
	{
		var temp_in_VkQueryPipelineStatisticFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.pipelineStatistics)))
			temp_in_VkQueryPipelineStatisticFlags = Flags(temp_in_VkFlags)
		}
		g.PipelineStatistics = QueryPipelineStatisticFlags(temp_in_VkQueryPipelineStatisticFlags)
	}
}

type CommandBufferBeginInfo struct {
	Type            StructureType
	Next            unsafe.Pointer
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

func (g *CommandBufferBeginInfo) toC(c *C.VkCommandBufferBeginInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkCommandBufferUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkCommandBufferUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandBufferUsageFlags(temp_in_VkCommandBufferUsageFlags)
	}
	{
		c.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(_sa.alloc(C.sizeof_VkCommandBufferInheritanceInfo))
		g.InheritanceInfo.toC(c.pInheritanceInfo)
	}
}
func (g *CommandBufferBeginInfo) fromC(c *C.VkCommandBufferBeginInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkCommandBufferUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkCommandBufferUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = CommandBufferUsageFlags(temp_in_VkCommandBufferUsageFlags)
	}
	{
		if g.InheritanceInfo == nil {
			g.InheritanceInfo = new(CommandBufferInheritanceInfo)
		}
		g.InheritanceInfo.fromC(c.pInheritanceInfo)
	}
}
func BeginCommandBuffer(commandBuffer CommandBuffer, beginInfo *CommandBufferBeginInfo) (_ret Result) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		pBeginInfo    *C.VkCommandBufferBeginInfo
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.pBeginInfo = (*C.VkCommandBufferBeginInfo)(_sa.alloc(C.sizeof_VkCommandBufferBeginInfo))
		beginInfo.toC(c.pBeginInfo, _sa)
	}
	c._ret = C.vkBeginCommandBuffer(c.commandBuffer, c.pBeginInfo)
	_ret = Result(c._ret)
	return
}
func EndCommandBuffer(commandBuffer CommandBuffer) (_ret Result) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		_ret          C.VkResult
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c._ret = C.vkEndCommandBuffer(c.commandBuffer)
	_ret = Result(c._ret)
	return
}

type CommandBufferResetFlags Flags

func ResetCommandBuffer(commandBuffer CommandBuffer, flags CommandBufferResetFlags) (_ret Result) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		flags         C.VkCommandBufferResetFlags
		_ret          C.VkResult
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var temp_in_VkCommandBufferResetFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkCommandBufferResetFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandBufferResetFlags(temp_in_VkCommandBufferResetFlags)
	}
	c._ret = C.vkResetCommandBuffer(c.commandBuffer, c.flags)
	_ret = Result(c._ret)
	return
}
func CmdBindPipeline(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	var c struct {
		commandBuffer     C.VkCommandBuffer
		pipelineBindPoint C.VkPipelineBindPoint
		pipeline          C.VkPipeline
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.pipelineBindPoint = C.VkPipelineBindPoint(pipelineBindPoint)
	c.pipeline = C.VkPipeline(pipeline)
	C.vkCmdBindPipeline(c.commandBuffer, c.pipelineBindPoint, c.pipeline)
}
func CmdSetViewport(commandBuffer CommandBuffer, firstViewport uint32, viewports []Viewport) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		firstViewport C.uint32_t
		viewportCount C.uint32_t
		pViewports    *C.VkViewport
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.firstViewport = C.uint32_t(firstViewport)
	c.viewportCount = C.uint32_t(len(viewports))
	{
		c.pViewports = (*C.VkViewport)(_sa.alloc(C.sizeof_VkViewport * uint(len(viewports))))
		slice3 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(viewports):len(viewports)]
		for i3, _ := range viewports {
			viewports[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdSetViewport(c.commandBuffer, c.firstViewport, c.viewportCount, c.pViewports)
}
func CmdSetScissor(commandBuffer CommandBuffer, firstScissor uint32, scissors []Rect2D) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		firstScissor  C.uint32_t
		scissorCount  C.uint32_t
		pScissors     *C.VkRect2D
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.firstScissor = C.uint32_t(firstScissor)
	c.scissorCount = C.uint32_t(len(scissors))
	{
		c.pScissors = (*C.VkRect2D)(_sa.alloc(C.sizeof_VkRect2D * uint(len(scissors))))
		slice3 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(scissors):len(scissors)]
		for i3, _ := range scissors {
			scissors[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdSetScissor(c.commandBuffer, c.firstScissor, c.scissorCount, c.pScissors)
}
func CmdSetLineWidth(commandBuffer CommandBuffer, lineWidth float32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		lineWidth     C.float
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.lineWidth = C.float(lineWidth)
	C.vkCmdSetLineWidth(c.commandBuffer, c.lineWidth)
}
func CmdSetDepthBias(commandBuffer CommandBuffer, depthBiasConstantFactor float32, depthBiasClamp float32, depthBiasSlopeFactor float32) {
	var c struct {
		commandBuffer           C.VkCommandBuffer
		depthBiasConstantFactor C.float
		depthBiasClamp          C.float
		depthBiasSlopeFactor    C.float
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.depthBiasConstantFactor = C.float(depthBiasConstantFactor)
	c.depthBiasClamp = C.float(depthBiasClamp)
	c.depthBiasSlopeFactor = C.float(depthBiasSlopeFactor)
	C.vkCmdSetDepthBias(c.commandBuffer, c.depthBiasConstantFactor, c.depthBiasClamp, c.depthBiasSlopeFactor)
}
func CmdSetBlendConstants(commandBuffer CommandBuffer, blendConstants []float32) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		blendConstants *C.float
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.blendConstants = (*C.float)(_sa.alloc(C.sizeof_float * uint(len(blendConstants))))
		slice3 := (*[1 << 31]C.float)(unsafe.Pointer(c.blendConstants))[:len(blendConstants):len(blendConstants)]
		for i3, _ := range blendConstants {
			slice3[i3] = C.float(blendConstants[i3])
		}
	}
	C.vkCmdSetBlendConstants(c.commandBuffer, c.blendConstants)
}
func CmdSetDepthBounds(commandBuffer CommandBuffer, minDepthBounds float32, maxDepthBounds float32) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		minDepthBounds C.float
		maxDepthBounds C.float
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.minDepthBounds = C.float(minDepthBounds)
	c.maxDepthBounds = C.float(maxDepthBounds)
	C.vkCmdSetDepthBounds(c.commandBuffer, c.minDepthBounds, c.maxDepthBounds)
}

type StencilFaceFlags Flags

func CmdSetStencilCompareMask(commandBuffer CommandBuffer, faceMask StencilFaceFlags, compareMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		compareMask   C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var temp_in_VkStencilFaceFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(faceMask)))
			temp_in_VkStencilFaceFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.faceMask = C.VkStencilFaceFlags(temp_in_VkStencilFaceFlags)
	}
	c.compareMask = C.uint32_t(compareMask)
	C.vkCmdSetStencilCompareMask(c.commandBuffer, c.faceMask, c.compareMask)
}
func CmdSetStencilWriteMask(commandBuffer CommandBuffer, faceMask StencilFaceFlags, writeMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		writeMask     C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var temp_in_VkStencilFaceFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(faceMask)))
			temp_in_VkStencilFaceFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.faceMask = C.VkStencilFaceFlags(temp_in_VkStencilFaceFlags)
	}
	c.writeMask = C.uint32_t(writeMask)
	C.vkCmdSetStencilWriteMask(c.commandBuffer, c.faceMask, c.writeMask)
}
func CmdSetStencilReference(commandBuffer CommandBuffer, faceMask StencilFaceFlags, reference uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		reference     C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var temp_in_VkStencilFaceFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(faceMask)))
			temp_in_VkStencilFaceFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.faceMask = C.VkStencilFaceFlags(temp_in_VkStencilFaceFlags)
	}
	c.reference = C.uint32_t(reference)
	C.vkCmdSetStencilReference(c.commandBuffer, c.faceMask, c.reference)
}
func CmdBindDescriptorSets(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet uint32, descriptorSets []DescriptorSet, dynamicOffsets []uint32) {
	var c struct {
		commandBuffer      C.VkCommandBuffer
		pipelineBindPoint  C.VkPipelineBindPoint
		layout             C.VkPipelineLayout
		firstSet           C.uint32_t
		descriptorSetCount C.uint32_t
		pDescriptorSets    *C.VkDescriptorSet
		dynamicOffsetCount C.uint32_t
		pDynamicOffsets    *C.uint32_t
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.pipelineBindPoint = C.VkPipelineBindPoint(pipelineBindPoint)
	c.layout = C.VkPipelineLayout(layout)
	c.firstSet = C.uint32_t(firstSet)
	c.descriptorSetCount = C.uint32_t(len(descriptorSets))
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(descriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(descriptorSets):len(descriptorSets)]
		for i3, _ := range descriptorSets {
			slice3[i3] = C.VkDescriptorSet(descriptorSets[i3])
		}
	}
	c.dynamicOffsetCount = C.uint32_t(len(dynamicOffsets))
	{
		c.pDynamicOffsets = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(dynamicOffsets))))
		slice3 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pDynamicOffsets))[:len(dynamicOffsets):len(dynamicOffsets)]
		for i3, _ := range dynamicOffsets {
			slice3[i3] = C.uint32_t(dynamicOffsets[i3])
		}
	}
	C.vkCmdBindDescriptorSets(c.commandBuffer, c.pipelineBindPoint, c.layout, c.firstSet, c.descriptorSetCount, c.pDescriptorSets, c.dynamicOffsetCount, c.pDynamicOffsets)
}

type IndexType int

const (
	INDEX_TYPE_UINT16      IndexType = 0
	INDEX_TYPE_UINT32      IndexType = 1
	INDEX_TYPE_BEGIN_RANGE IndexType = INDEX_TYPE_UINT16
	INDEX_TYPE_END_RANGE   IndexType = INDEX_TYPE_UINT32
	INDEX_TYPE_RANGE_SIZE  IndexType = (INDEX_TYPE_UINT32 - INDEX_TYPE_UINT16 + 1)
	INDEX_TYPE_MAX_ENUM    IndexType = 2147483647
)

func CmdBindIndexBuffer(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, indexType IndexType) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
		indexType     C.VkIndexType
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.indexType = C.VkIndexType(indexType)
	C.vkCmdBindIndexBuffer(c.commandBuffer, c.buffer, c.offset, c.indexType)
}
func CmdBindVertexBuffers(commandBuffer CommandBuffer, firstBinding uint32, buffers []Buffer, offsets []DeviceSize) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		firstBinding  C.uint32_t
		bindingCount  C.uint32_t
		pBuffers      *C.VkBuffer
		pOffsets      *C.VkDeviceSize
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.firstBinding = C.uint32_t(firstBinding)
	c.bindingCount = C.uint32_t(len(buffers))
	{
		c.pBuffers = (*C.VkBuffer)(_sa.alloc(C.sizeof_VkBuffer * uint(len(buffers))))
		slice3 := (*[1 << 31]C.VkBuffer)(unsafe.Pointer(c.pBuffers))[:len(buffers):len(buffers)]
		for i3, _ := range buffers {
			slice3[i3] = C.VkBuffer(buffers[i3])
		}
	}
	{
		c.pOffsets = (*C.VkDeviceSize)(_sa.alloc(C.sizeof_VkDeviceSize * uint(len(offsets))))
		slice3 := (*[1 << 31]C.VkDeviceSize)(unsafe.Pointer(c.pOffsets))[:len(offsets):len(offsets)]
		for i3, _ := range offsets {
			{
				var temp_in_VkDeviceSize C.uint64_t
				temp_in_VkDeviceSize = C.uint64_t((uint64)(offsets[i3]))
				slice3[i3] = C.VkDeviceSize(temp_in_VkDeviceSize)
			}
		}
	}
	C.vkCmdBindVertexBuffers(c.commandBuffer, c.firstBinding, c.bindingCount, c.pBuffers, c.pOffsets)
}
func CmdDraw(commandBuffer CommandBuffer, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		vertexCount   C.uint32_t
		instanceCount C.uint32_t
		firstVertex   C.uint32_t
		firstInstance C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.vertexCount = C.uint32_t(vertexCount)
	c.instanceCount = C.uint32_t(instanceCount)
	c.firstVertex = C.uint32_t(firstVertex)
	c.firstInstance = C.uint32_t(firstInstance)
	C.vkCmdDraw(c.commandBuffer, c.vertexCount, c.instanceCount, c.firstVertex, c.firstInstance)
}
func CmdDrawIndexed(commandBuffer CommandBuffer, indexCount uint32, instanceCount uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		indexCount    C.uint32_t
		instanceCount C.uint32_t
		firstIndex    C.uint32_t
		vertexOffset  C.int32_t
		firstInstance C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.indexCount = C.uint32_t(indexCount)
	c.instanceCount = C.uint32_t(instanceCount)
	c.firstIndex = C.uint32_t(firstIndex)
	c.vertexOffset = C.int32_t(vertexOffset)
	c.firstInstance = C.uint32_t(firstInstance)
	C.vkCmdDrawIndexed(c.commandBuffer, c.indexCount, c.instanceCount, c.firstIndex, c.vertexOffset, c.firstInstance)
}
func CmdDrawIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
		drawCount     C.uint32_t
		stride        C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.drawCount = C.uint32_t(drawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndirect(c.commandBuffer, c.buffer, c.offset, c.drawCount, c.stride)
}
func CmdDrawIndexedIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount uint32, stride uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
		drawCount     C.uint32_t
		stride        C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.drawCount = C.uint32_t(drawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndexedIndirect(c.commandBuffer, c.buffer, c.offset, c.drawCount, c.stride)
}
func CmdDispatch(commandBuffer CommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		groupCountX   C.uint32_t
		groupCountY   C.uint32_t
		groupCountZ   C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.groupCountX = C.uint32_t(groupCountX)
	c.groupCountY = C.uint32_t(groupCountY)
	c.groupCountZ = C.uint32_t(groupCountZ)
	C.vkCmdDispatch(c.commandBuffer, c.groupCountX, c.groupCountY, c.groupCountZ)
}
func CmdDispatchIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	C.vkCmdDispatchIndirect(c.commandBuffer, c.buffer, c.offset)
}

type BufferCopy struct {
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}

func (g *BufferCopy) toC(c *C.VkBufferCopy) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.SrcOffset))
		c.srcOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.DstOffset))
		c.dstOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *BufferCopy) fromC(c *C.VkBufferCopy) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.srcOffset))
		g.SrcOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.dstOffset))
		g.DstOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
}
func CmdCopyBuffer(commandBuffer CommandBuffer, srcBuffer Buffer, dstBuffer Buffer, regions []BufferCopy) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		srcBuffer     C.VkBuffer
		dstBuffer     C.VkBuffer
		regionCount   C.uint32_t
		pRegions      *C.VkBufferCopy
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcBuffer = C.VkBuffer(srcBuffer)
	c.dstBuffer = C.VkBuffer(dstBuffer)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkBufferCopy)(_sa.alloc(C.sizeof_VkBufferCopy * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkBufferCopy)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyBuffer(c.commandBuffer, c.srcBuffer, c.dstBuffer, c.regionCount, c.pRegions)
}

type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (g *ImageSubresourceLayers) toC(c *C.VkImageSubresourceLayers) {
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	c.mipLevel = C.uint32_t(g.MipLevel)
	c.baseArrayLayer = C.uint32_t(g.BaseArrayLayer)
	c.layerCount = C.uint32_t(g.LayerCount)
}
func (g *ImageSubresourceLayers) fromC(c *C.VkImageSubresourceLayers) {
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.MipLevel = uint32(c.mipLevel)
	g.BaseArrayLayer = uint32(c.baseArrayLayer)
	g.LayerCount = uint32(c.layerCount)
}

type ImageCopy struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func (g *ImageCopy) toC(c *C.VkImageCopy) {
	g.SrcSubresource.toC(&c.srcSubresource)
	g.SrcOffset.toC(&c.srcOffset)
	g.DstSubresource.toC(&c.dstSubresource)
	g.DstOffset.toC(&c.dstOffset)
	g.Extent.toC(&c.extent)
}
func (g *ImageCopy) fromC(c *C.VkImageCopy) {
	g.SrcSubresource.fromC(&c.srcSubresource)
	g.SrcOffset.fromC(&c.srcOffset)
	g.DstSubresource.fromC(&c.dstSubresource)
	g.DstOffset.fromC(&c.dstOffset)
	g.Extent.fromC(&c.extent)
}
func CmdCopyImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageCopy) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		srcImage       C.VkImage
		srcImageLayout C.VkImageLayout
		dstImage       C.VkImage
		dstImageLayout C.VkImageLayout
		regionCount    C.uint32_t
		pRegions       *C.VkImageCopy
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcImage = C.VkImage(srcImage)
	c.srcImageLayout = C.VkImageLayout(srcImageLayout)
	c.dstImage = C.VkImage(dstImage)
	c.dstImageLayout = C.VkImageLayout(dstImageLayout)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkImageCopy)(_sa.alloc(C.sizeof_VkImageCopy * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkImageCopy)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyImage(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions)
}

type ImageBlit struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}

func (g *ImageBlit) toC(c *C.VkImageBlit) {
	g.SrcSubresource.toC(&c.srcSubresource)
	for i, _ := range g.SrcOffsets {
		g.SrcOffsets[i].toC(&c.srcOffsets[i])
	}
	g.DstSubresource.toC(&c.dstSubresource)
	for i, _ := range g.DstOffsets {
		g.DstOffsets[i].toC(&c.dstOffsets[i])
	}
}
func (g *ImageBlit) fromC(c *C.VkImageBlit) {
	g.SrcSubresource.fromC(&c.srcSubresource)
	for i, _ := range g.SrcOffsets {
		g.SrcOffsets[i].fromC(&c.srcOffsets[i])
	}
	g.DstSubresource.fromC(&c.dstSubresource)
	for i, _ := range g.DstOffsets {
		g.DstOffsets[i].fromC(&c.dstOffsets[i])
	}
}
func CmdBlitImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageBlit, filter Filter) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		srcImage       C.VkImage
		srcImageLayout C.VkImageLayout
		dstImage       C.VkImage
		dstImageLayout C.VkImageLayout
		regionCount    C.uint32_t
		pRegions       *C.VkImageBlit
		filter         C.VkFilter
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcImage = C.VkImage(srcImage)
	c.srcImageLayout = C.VkImageLayout(srcImageLayout)
	c.dstImage = C.VkImage(dstImage)
	c.dstImageLayout = C.VkImageLayout(dstImageLayout)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkImageBlit)(_sa.alloc(C.sizeof_VkImageBlit * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkImageBlit)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	c.filter = C.VkFilter(filter)
	C.vkCmdBlitImage(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions, c.filter)
}

type BufferImageCopy struct {
	BufferOffset      DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D
}

func (g *BufferImageCopy) toC(c *C.VkBufferImageCopy) {
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.BufferOffset))
		c.bufferOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.bufferRowLength = C.uint32_t(g.BufferRowLength)
	c.bufferImageHeight = C.uint32_t(g.BufferImageHeight)
	g.ImageSubresource.toC(&c.imageSubresource)
	g.ImageOffset.toC(&c.imageOffset)
	g.ImageExtent.toC(&c.imageExtent)
}
func (g *BufferImageCopy) fromC(c *C.VkBufferImageCopy) {
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.bufferOffset))
		g.BufferOffset = DeviceSize(temp_in_VkDeviceSize)
	}
	g.BufferRowLength = uint32(c.bufferRowLength)
	g.BufferImageHeight = uint32(c.bufferImageHeight)
	g.ImageSubresource.fromC(&c.imageSubresource)
	g.ImageOffset.fromC(&c.imageOffset)
	g.ImageExtent.fromC(&c.imageExtent)
}
func CmdCopyBufferToImage(commandBuffer CommandBuffer, srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regions []BufferImageCopy) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		srcBuffer      C.VkBuffer
		dstImage       C.VkImage
		dstImageLayout C.VkImageLayout
		regionCount    C.uint32_t
		pRegions       *C.VkBufferImageCopy
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcBuffer = C.VkBuffer(srcBuffer)
	c.dstImage = C.VkImage(dstImage)
	c.dstImageLayout = C.VkImageLayout(dstImageLayout)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkBufferImageCopy)(_sa.alloc(C.sizeof_VkBufferImageCopy * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkBufferImageCopy)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyBufferToImage(c.commandBuffer, c.srcBuffer, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions)
}
func CmdCopyImageToBuffer(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		srcImage       C.VkImage
		srcImageLayout C.VkImageLayout
		dstBuffer      C.VkBuffer
		regionCount    C.uint32_t
		pRegions       *C.VkBufferImageCopy
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcImage = C.VkImage(srcImage)
	c.srcImageLayout = C.VkImageLayout(srcImageLayout)
	c.dstBuffer = C.VkBuffer(dstBuffer)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkBufferImageCopy)(_sa.alloc(C.sizeof_VkBufferImageCopy * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkBufferImageCopy)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyImageToBuffer(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstBuffer, c.regionCount, c.pRegions)
}
func CmdUpdateBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, dataSize DeviceSize, data []byte) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		dstBuffer     C.VkBuffer
		dstOffset     C.VkDeviceSize
		dataSize      C.VkDeviceSize
		pData         unsafe.Pointer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.dstBuffer = C.VkBuffer(dstBuffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(dstOffset))
		c.dstOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(dataSize))
		c.dataSize = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	C.vkCmdUpdateBuffer(c.commandBuffer, c.dstBuffer, c.dstOffset, c.dataSize, c.pData)
}
func CmdFillBuffer(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset DeviceSize, size DeviceSize, data uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		dstBuffer     C.VkBuffer
		dstOffset     C.VkDeviceSize
		size          C.VkDeviceSize
		data          C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.dstBuffer = C.VkBuffer(dstBuffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(dstOffset))
		c.dstOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.data = C.uint32_t(data)
	C.vkCmdFillBuffer(c.commandBuffer, c.dstBuffer, c.dstOffset, c.size, c.data)
}

type ClearColorValue struct{ Raw C.VkClearColorValue }

func (g *ClearColorValue) AssginFloat32(v [4]float32) {
	var cv [4]C.float
	for i, _ := range v {
		cv[i] = C.float(v[i])
	}
	*(*[4]C.float)(unsafe.Pointer(&g.Raw)) = cv
}
func (g *ClearColorValue) Float32() (v [4]float32) {
	cv := *(*[4]C.float)(unsafe.Pointer(&g.Raw))
	for i, _ := range v {
		v[i] = float32(cv[i])
	}
	return
}
func (g *ClearColorValue) AssginInt32(v [4]int32) {
	var cv [4]C.int32_t
	for i, _ := range v {
		cv[i] = C.int32_t(v[i])
	}
	*(*[4]C.int32_t)(unsafe.Pointer(&g.Raw)) = cv
}
func (g *ClearColorValue) Int32() (v [4]int32) {
	cv := *(*[4]C.int32_t)(unsafe.Pointer(&g.Raw))
	for i, _ := range v {
		v[i] = int32(cv[i])
	}
	return
}
func (g *ClearColorValue) AssginUint32(v [4]uint32) {
	var cv [4]C.uint32_t
	for i, _ := range v {
		cv[i] = C.uint32_t(v[i])
	}
	*(*[4]C.uint32_t)(unsafe.Pointer(&g.Raw)) = cv
}
func (g *ClearColorValue) Uint32() (v [4]uint32) {
	cv := *(*[4]C.uint32_t)(unsafe.Pointer(&g.Raw))
	for i, _ := range v {
		v[i] = uint32(cv[i])
	}
	return
}
func CmdClearColorImage(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, color *ClearColorValue, ranges []ImageSubresourceRange) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		image         C.VkImage
		imageLayout   C.VkImageLayout
		pColor        *C.VkClearColorValue
		rangeCount    C.uint32_t
		pRanges       *C.VkImageSubresourceRange
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.image = C.VkImage(image)
	c.imageLayout = C.VkImageLayout(imageLayout)
	{
		c.pColor = (*C.VkClearColorValue)(_sa.alloc(C.sizeof_VkClearColorValue))
		*c.pColor = g.Raw
	}
	c.rangeCount = C.uint32_t(len(ranges))
	{
		c.pRanges = (*C.VkImageSubresourceRange)(_sa.alloc(C.sizeof_VkImageSubresourceRange * uint(len(ranges))))
		slice3 := (*[1 << 31]C.VkImageSubresourceRange)(unsafe.Pointer(c.pRanges))[:len(ranges):len(ranges)]
		for i3, _ := range ranges {
			ranges[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdClearColorImage(c.commandBuffer, c.image, c.imageLayout, c.pColor, c.rangeCount, c.pRanges)
}

type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

func (g *ClearDepthStencilValue) toC(c *C.VkClearDepthStencilValue) {
	c.depth = C.float(g.Depth)
	c.stencil = C.uint32_t(g.Stencil)
}
func (g *ClearDepthStencilValue) fromC(c *C.VkClearDepthStencilValue) {
	g.Depth = float32(c.depth)
	g.Stencil = uint32(c.stencil)
}
func CmdClearDepthStencilImage(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, depthStencil *ClearDepthStencilValue, ranges []ImageSubresourceRange) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		image         C.VkImage
		imageLayout   C.VkImageLayout
		pDepthStencil *C.VkClearDepthStencilValue
		rangeCount    C.uint32_t
		pRanges       *C.VkImageSubresourceRange
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.image = C.VkImage(image)
	c.imageLayout = C.VkImageLayout(imageLayout)
	{
		c.pDepthStencil = (*C.VkClearDepthStencilValue)(_sa.alloc(C.sizeof_VkClearDepthStencilValue))
		depthStencil.toC(c.pDepthStencil)
	}
	c.rangeCount = C.uint32_t(len(ranges))
	{
		c.pRanges = (*C.VkImageSubresourceRange)(_sa.alloc(C.sizeof_VkImageSubresourceRange * uint(len(ranges))))
		slice3 := (*[1 << 31]C.VkImageSubresourceRange)(unsafe.Pointer(c.pRanges))[:len(ranges):len(ranges)]
		for i3, _ := range ranges {
			ranges[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdClearDepthStencilImage(c.commandBuffer, c.image, c.imageLayout, c.pDepthStencil, c.rangeCount, c.pRanges)
}

type ClearValue struct{ Raw C.VkClearValue }

func (g *ClearValue) AssginColor(v ClearColorValue) {
	var cv C.VkClearColorValue
	cv = g.Raw
	*(*C.VkClearColorValue)(unsafe.Pointer(&g.Raw)) = cv
}
func (g *ClearValue) Color() (v ClearColorValue) {
	cv := *(*C.VkClearColorValue)(unsafe.Pointer(&g.Raw))
	g.Raw = cv
	return
}
func (g *ClearValue) AssginDepthStencil(v ClearDepthStencilValue) {
	var cv C.VkClearDepthStencilValue
	v.toC(&cv)
	*(*C.VkClearDepthStencilValue)(unsafe.Pointer(&g.Raw)) = cv
}
func (g *ClearValue) DepthStencil() (v ClearDepthStencilValue) {
	cv := *(*C.VkClearDepthStencilValue)(unsafe.Pointer(&g.Raw))
	v.fromC(&cv)
	return
}

type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}

func (g *ClearAttachment) toC(c *C.VkClearAttachment) {
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	c.colorAttachment = C.uint32_t(g.ColorAttachment)
	c.clearValue = g.Raw
}
func (g *ClearAttachment) fromC(c *C.VkClearAttachment) {
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
	g.ColorAttachment = uint32(c.colorAttachment)
	g.Raw = c.clearValue
}

type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

func (g *ClearRect) toC(c *C.VkClearRect) {
	g.Rect.toC(&c.rect)
	c.baseArrayLayer = C.uint32_t(g.BaseArrayLayer)
	c.layerCount = C.uint32_t(g.LayerCount)
}
func (g *ClearRect) fromC(c *C.VkClearRect) {
	g.Rect.fromC(&c.rect)
	g.BaseArrayLayer = uint32(c.baseArrayLayer)
	g.LayerCount = uint32(c.layerCount)
}
func CmdClearAttachments(commandBuffer CommandBuffer, attachments []ClearAttachment, rects []ClearRect) {
	var c struct {
		commandBuffer   C.VkCommandBuffer
		attachmentCount C.uint32_t
		pAttachments    *C.VkClearAttachment
		rectCount       C.uint32_t
		pRects          *C.VkClearRect
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.attachmentCount = C.uint32_t(len(attachments))
	{
		c.pAttachments = (*C.VkClearAttachment)(_sa.alloc(C.sizeof_VkClearAttachment * uint(len(attachments))))
		slice3 := (*[1 << 31]C.VkClearAttachment)(unsafe.Pointer(c.pAttachments))[:len(attachments):len(attachments)]
		for i3, _ := range attachments {
			attachments[i3].toC(&slice3[i3])
		}
	}
	c.rectCount = C.uint32_t(len(rects))
	{
		c.pRects = (*C.VkClearRect)(_sa.alloc(C.sizeof_VkClearRect * uint(len(rects))))
		slice3 := (*[1 << 31]C.VkClearRect)(unsafe.Pointer(c.pRects))[:len(rects):len(rects)]
		for i3, _ := range rects {
			rects[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdClearAttachments(c.commandBuffer, c.attachmentCount, c.pAttachments, c.rectCount, c.pRects)
}

type ImageResolve struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func (g *ImageResolve) toC(c *C.VkImageResolve) {
	g.SrcSubresource.toC(&c.srcSubresource)
	g.SrcOffset.toC(&c.srcOffset)
	g.DstSubresource.toC(&c.dstSubresource)
	g.DstOffset.toC(&c.dstOffset)
	g.Extent.toC(&c.extent)
}
func (g *ImageResolve) fromC(c *C.VkImageResolve) {
	g.SrcSubresource.fromC(&c.srcSubresource)
	g.SrcOffset.fromC(&c.srcOffset)
	g.DstSubresource.fromC(&c.dstSubresource)
	g.DstOffset.fromC(&c.dstOffset)
	g.Extent.fromC(&c.extent)
}
func CmdResolveImage(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageResolve) {
	var c struct {
		commandBuffer  C.VkCommandBuffer
		srcImage       C.VkImage
		srcImageLayout C.VkImageLayout
		dstImage       C.VkImage
		dstImageLayout C.VkImageLayout
		regionCount    C.uint32_t
		pRegions       *C.VkImageResolve
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.srcImage = C.VkImage(srcImage)
	c.srcImageLayout = C.VkImageLayout(srcImageLayout)
	c.dstImage = C.VkImage(dstImage)
	c.dstImageLayout = C.VkImageLayout(dstImageLayout)
	c.regionCount = C.uint32_t(len(regions))
	{
		c.pRegions = (*C.VkImageResolve)(_sa.alloc(C.sizeof_VkImageResolve * uint(len(regions))))
		slice3 := (*[1 << 31]C.VkImageResolve)(unsafe.Pointer(c.pRegions))[:len(regions):len(regions)]
		for i3, _ := range regions {
			regions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdResolveImage(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions)
}
func CmdSetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		event         C.VkEvent
		stageMask     C.VkPipelineStageFlags
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.event = C.VkEvent(event)
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(stageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.stageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	C.vkCmdSetEvent(c.commandBuffer, c.event, c.stageMask)
}
func CmdResetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		event         C.VkEvent
		stageMask     C.VkPipelineStageFlags
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.event = C.VkEvent(event)
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(stageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.stageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	C.vkCmdResetEvent(c.commandBuffer, c.event, c.stageMask)
}

type MemoryBarrier struct {
	Type          StructureType
	Next          unsafe.Pointer
	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

func (g *MemoryBarrier) toC(c *C.VkMemoryBarrier) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
}
func (g *MemoryBarrier) fromC(c *C.VkMemoryBarrier) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.SrcAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.DstAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
}

type BufferMemoryBarrier struct {
	Type                StructureType
	Next                unsafe.Pointer
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              Buffer
	Offset              DeviceSize
	Size                DeviceSize
}

func (g *BufferMemoryBarrier) toC(c *C.VkBufferMemoryBarrier) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	c.srcQueueFamilyIndex = C.uint32_t(g.SrcQueueFamilyIndex)
	c.dstQueueFamilyIndex = C.uint32_t(g.DstQueueFamilyIndex)
	c.buffer = C.VkBuffer(g.Buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.Size))
		c.size = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *BufferMemoryBarrier) fromC(c *C.VkBufferMemoryBarrier) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.SrcAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.DstAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	g.SrcQueueFamilyIndex = uint32(c.srcQueueFamilyIndex)
	g.DstQueueFamilyIndex = uint32(c.dstQueueFamilyIndex)
	g.Buffer = Buffer(c.buffer)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.offset))
		g.Offset = DeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.size))
		g.Size = DeviceSize(temp_in_VkDeviceSize)
	}
}

type ImageMemoryBarrier struct {
	Type                StructureType
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

func (g *ImageMemoryBarrier) toC(c *C.VkImageMemoryBarrier) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	c.oldLayout = C.VkImageLayout(g.OldLayout)
	c.newLayout = C.VkImageLayout(g.NewLayout)
	c.srcQueueFamilyIndex = C.uint32_t(g.SrcQueueFamilyIndex)
	c.dstQueueFamilyIndex = C.uint32_t(g.DstQueueFamilyIndex)
	c.image = C.VkImage(g.Image)
	g.SubresourceRange.toC(&c.subresourceRange)
}
func (g *ImageMemoryBarrier) fromC(c *C.VkImageMemoryBarrier) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.SrcAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.DstAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	g.OldLayout = ImageLayout(c.oldLayout)
	g.NewLayout = ImageLayout(c.newLayout)
	g.SrcQueueFamilyIndex = uint32(c.srcQueueFamilyIndex)
	g.DstQueueFamilyIndex = uint32(c.dstQueueFamilyIndex)
	g.Image = Image(c.image)
	g.SubresourceRange.fromC(&c.subresourceRange)
}
func CmdWaitEvents(commandBuffer CommandBuffer, events []Event, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {
	var c struct {
		commandBuffer            C.VkCommandBuffer
		eventCount               C.uint32_t
		pEvents                  *C.VkEvent
		srcStageMask             C.VkPipelineStageFlags
		dstStageMask             C.VkPipelineStageFlags
		memoryBarrierCount       C.uint32_t
		pMemoryBarriers          *C.VkMemoryBarrier
		bufferMemoryBarrierCount C.uint32_t
		pBufferMemoryBarriers    *C.VkBufferMemoryBarrier
		imageMemoryBarrierCount  C.uint32_t
		pImageMemoryBarriers     *C.VkImageMemoryBarrier
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.eventCount = C.uint32_t(len(events))
	{
		c.pEvents = (*C.VkEvent)(_sa.alloc(C.sizeof_VkEvent * uint(len(events))))
		slice3 := (*[1 << 31]C.VkEvent)(unsafe.Pointer(c.pEvents))[:len(events):len(events)]
		for i3, _ := range events {
			slice3[i3] = C.VkEvent(events[i3])
		}
	}
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(srcStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(dstStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	c.memoryBarrierCount = C.uint32_t(len(memoryBarriers))
	{
		c.pMemoryBarriers = (*C.VkMemoryBarrier)(_sa.alloc(C.sizeof_VkMemoryBarrier * uint(len(memoryBarriers))))
		slice3 := (*[1 << 31]C.VkMemoryBarrier)(unsafe.Pointer(c.pMemoryBarriers))[:len(memoryBarriers):len(memoryBarriers)]
		for i3, _ := range memoryBarriers {
			memoryBarriers[i3].toC(&slice3[i3])
		}
	}
	c.bufferMemoryBarrierCount = C.uint32_t(len(bufferMemoryBarriers))
	{
		c.pBufferMemoryBarriers = (*C.VkBufferMemoryBarrier)(_sa.alloc(C.sizeof_VkBufferMemoryBarrier * uint(len(bufferMemoryBarriers))))
		slice3 := (*[1 << 31]C.VkBufferMemoryBarrier)(unsafe.Pointer(c.pBufferMemoryBarriers))[:len(bufferMemoryBarriers):len(bufferMemoryBarriers)]
		for i3, _ := range bufferMemoryBarriers {
			bufferMemoryBarriers[i3].toC(&slice3[i3])
		}
	}
	c.imageMemoryBarrierCount = C.uint32_t(len(imageMemoryBarriers))
	{
		c.pImageMemoryBarriers = (*C.VkImageMemoryBarrier)(_sa.alloc(C.sizeof_VkImageMemoryBarrier * uint(len(imageMemoryBarriers))))
		slice3 := (*[1 << 31]C.VkImageMemoryBarrier)(unsafe.Pointer(c.pImageMemoryBarriers))[:len(imageMemoryBarriers):len(imageMemoryBarriers)]
		for i3, _ := range imageMemoryBarriers {
			imageMemoryBarriers[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdWaitEvents(c.commandBuffer, c.eventCount, c.pEvents, c.srcStageMask, c.dstStageMask, c.memoryBarrierCount, c.pMemoryBarriers, c.bufferMemoryBarrierCount, c.pBufferMemoryBarriers, c.imageMemoryBarrierCount, c.pImageMemoryBarriers)
}
func CmdPipelineBarrier(commandBuffer CommandBuffer, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {
	var c struct {
		commandBuffer            C.VkCommandBuffer
		srcStageMask             C.VkPipelineStageFlags
		dstStageMask             C.VkPipelineStageFlags
		dependencyFlags          C.VkDependencyFlags
		memoryBarrierCount       C.uint32_t
		pMemoryBarriers          *C.VkMemoryBarrier
		bufferMemoryBarrierCount C.uint32_t
		pBufferMemoryBarriers    *C.VkBufferMemoryBarrier
		imageMemoryBarrierCount  C.uint32_t
		pImageMemoryBarriers     *C.VkImageMemoryBarrier
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(srcStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(dstStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkDependencyFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(dependencyFlags)))
			temp_in_VkDependencyFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dependencyFlags = C.VkDependencyFlags(temp_in_VkDependencyFlags)
	}
	c.memoryBarrierCount = C.uint32_t(len(memoryBarriers))
	{
		c.pMemoryBarriers = (*C.VkMemoryBarrier)(_sa.alloc(C.sizeof_VkMemoryBarrier * uint(len(memoryBarriers))))
		slice3 := (*[1 << 31]C.VkMemoryBarrier)(unsafe.Pointer(c.pMemoryBarriers))[:len(memoryBarriers):len(memoryBarriers)]
		for i3, _ := range memoryBarriers {
			memoryBarriers[i3].toC(&slice3[i3])
		}
	}
	c.bufferMemoryBarrierCount = C.uint32_t(len(bufferMemoryBarriers))
	{
		c.pBufferMemoryBarriers = (*C.VkBufferMemoryBarrier)(_sa.alloc(C.sizeof_VkBufferMemoryBarrier * uint(len(bufferMemoryBarriers))))
		slice3 := (*[1 << 31]C.VkBufferMemoryBarrier)(unsafe.Pointer(c.pBufferMemoryBarriers))[:len(bufferMemoryBarriers):len(bufferMemoryBarriers)]
		for i3, _ := range bufferMemoryBarriers {
			bufferMemoryBarriers[i3].toC(&slice3[i3])
		}
	}
	c.imageMemoryBarrierCount = C.uint32_t(len(imageMemoryBarriers))
	{
		c.pImageMemoryBarriers = (*C.VkImageMemoryBarrier)(_sa.alloc(C.sizeof_VkImageMemoryBarrier * uint(len(imageMemoryBarriers))))
		slice3 := (*[1 << 31]C.VkImageMemoryBarrier)(unsafe.Pointer(c.pImageMemoryBarriers))[:len(imageMemoryBarriers):len(imageMemoryBarriers)]
		for i3, _ := range imageMemoryBarriers {
			imageMemoryBarriers[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdPipelineBarrier(c.commandBuffer, c.srcStageMask, c.dstStageMask, c.dependencyFlags, c.memoryBarrierCount, c.pMemoryBarriers, c.bufferMemoryBarrierCount, c.pBufferMemoryBarriers, c.imageMemoryBarrierCount, c.pImageMemoryBarriers)
}
func CmdBeginQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		queryPool     C.VkQueryPool
		query         C.uint32_t
		flags         C.VkQueryControlFlags
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.queryPool = C.VkQueryPool(queryPool)
	c.query = C.uint32_t(query)
	{
		var temp_in_VkQueryControlFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkQueryControlFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkQueryControlFlags(temp_in_VkQueryControlFlags)
	}
	C.vkCmdBeginQuery(c.commandBuffer, c.queryPool, c.query, c.flags)
}
func CmdEndQuery(commandBuffer CommandBuffer, queryPool QueryPool, query uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		queryPool     C.VkQueryPool
		query         C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.queryPool = C.VkQueryPool(queryPool)
	c.query = C.uint32_t(query)
	C.vkCmdEndQuery(c.commandBuffer, c.queryPool, c.query)
}
func CmdResetQueryPool(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery uint32, queryCount uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		queryPool     C.VkQueryPool
		firstQuery    C.uint32_t
		queryCount    C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.queryPool = C.VkQueryPool(queryPool)
	c.firstQuery = C.uint32_t(firstQuery)
	c.queryCount = C.uint32_t(queryCount)
	C.vkCmdResetQueryPool(c.commandBuffer, c.queryPool, c.firstQuery, c.queryCount)
}

type PipelineStageFlagBits int

const (
	PIPELINE_STAGE_TOP_OF_PIPE_BIT                    PipelineStageFlagBits = 1
	PIPELINE_STAGE_DRAW_INDIRECT_BIT                  PipelineStageFlagBits = 2
	PIPELINE_STAGE_VERTEX_INPUT_BIT                   PipelineStageFlagBits = 4
	PIPELINE_STAGE_VERTEX_SHADER_BIT                  PipelineStageFlagBits = 8
	PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT    PipelineStageFlagBits = 16
	PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT PipelineStageFlagBits = 32
	PIPELINE_STAGE_GEOMETRY_SHADER_BIT                PipelineStageFlagBits = 64
	PIPELINE_STAGE_FRAGMENT_SHADER_BIT                PipelineStageFlagBits = 128
	PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT           PipelineStageFlagBits = 256
	PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT            PipelineStageFlagBits = 512
	PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT        PipelineStageFlagBits = 1024
	PIPELINE_STAGE_COMPUTE_SHADER_BIT                 PipelineStageFlagBits = 2048
	PIPELINE_STAGE_TRANSFER_BIT                       PipelineStageFlagBits = 4096
	PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT                 PipelineStageFlagBits = 8192
	PIPELINE_STAGE_HOST_BIT                           PipelineStageFlagBits = 16384
	PIPELINE_STAGE_ALL_GRAPHICS_BIT                   PipelineStageFlagBits = 32768
	PIPELINE_STAGE_ALL_COMMANDS_BIT                   PipelineStageFlagBits = 65536
	PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT      PipelineStageFlagBits = 262144
	PIPELINE_STAGE_COMMAND_PROCESS_BIT_NVX            PipelineStageFlagBits = 131072
	PIPELINE_STAGE_FLAG_BITS_MAX_ENUM                 PipelineStageFlagBits = 2147483647
)

func CmdWriteTimestamp(commandBuffer CommandBuffer, pipelineStage PipelineStageFlagBits, queryPool QueryPool, query uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		pipelineStage C.VkPipelineStageFlagBits
		queryPool     C.VkQueryPool
		query         C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.pipelineStage = C.VkPipelineStageFlagBits(pipelineStage)
	c.queryPool = C.VkQueryPool(queryPool)
	c.query = C.uint32_t(query)
	C.vkCmdWriteTimestamp(c.commandBuffer, c.pipelineStage, c.queryPool, c.query)
}
func CmdCopyQueryPoolResults(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery uint32, queryCount uint32, dstBuffer Buffer, dstOffset DeviceSize, stride DeviceSize, flags QueryResultFlags) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		queryPool     C.VkQueryPool
		firstQuery    C.uint32_t
		queryCount    C.uint32_t
		dstBuffer     C.VkBuffer
		dstOffset     C.VkDeviceSize
		stride        C.VkDeviceSize
		flags         C.VkQueryResultFlags
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.queryPool = C.VkQueryPool(queryPool)
	c.firstQuery = C.uint32_t(firstQuery)
	c.queryCount = C.uint32_t(queryCount)
	c.dstBuffer = C.VkBuffer(dstBuffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(dstOffset))
		c.dstOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(stride))
		c.stride = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	{
		var temp_in_VkQueryResultFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkQueryResultFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkQueryResultFlags(temp_in_VkQueryResultFlags)
	}
	C.vkCmdCopyQueryPoolResults(c.commandBuffer, c.queryPool, c.firstQuery, c.queryCount, c.dstBuffer, c.dstOffset, c.stride, c.flags)
}
func CmdPushConstants(commandBuffer CommandBuffer, layout PipelineLayout, stageFlags ShaderStageFlags, offset uint32, values []byte) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		layout        C.VkPipelineLayout
		stageFlags    C.VkShaderStageFlags
		offset        C.uint32_t
		size          C.uint32_t
		pValues       unsafe.Pointer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.layout = C.VkPipelineLayout(layout)
	{
		var temp_in_VkShaderStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(stageFlags)))
			temp_in_VkShaderStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.stageFlags = C.VkShaderStageFlags(temp_in_VkShaderStageFlags)
	}
	c.offset = C.uint32_t(offset)
	c.size = C.uint32_t(len(values))
	{
		c.pValues = _sa.alloc(C.sizeof_void_pointer * uint(len(values)))
		slice3 := (*[1 << 31]byte)(c.pValues)[:len(values):len(values)]
		for i3, _ := range values {
			slice3[i3] = values[i3]
		}
	}
	C.vkCmdPushConstants(c.commandBuffer, c.layout, c.stageFlags, c.offset, c.size, c.pValues)
}

type RenderPassBeginInfo struct {
	Type        StructureType
	Next        unsafe.Pointer
	RenderPass  RenderPass
	Framebuffer Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

func (g *RenderPassBeginInfo) toC(c *C.VkRenderPassBeginInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.renderPass = C.VkRenderPass(g.RenderPass)
	c.framebuffer = C.VkFramebuffer(g.Framebuffer)
	g.RenderArea.toC(&c.renderArea)
	c.clearValueCount = C.uint32_t(len(g.ClearValues))
	{
		c.pClearValues = (*C.VkClearValue)(_sa.alloc(C.sizeof_VkClearValue * uint(len(g.ClearValues))))
		slice2 := (*[1 << 31]C.VkClearValue)(unsafe.Pointer(c.pClearValues))[:len(g.ClearValues):len(g.ClearValues)]
		for i2, _ := range g.ClearValues {
			slice2[i2] = g.Raw
		}
	}
}
func (g *RenderPassBeginInfo) fromC(c *C.VkRenderPassBeginInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.RenderPass = RenderPass(c.renderPass)
	g.Framebuffer = Framebuffer(c.framebuffer)
	g.RenderArea.fromC(&c.renderArea)
	g.ClearValues = make([]ClearValue, int(c.clearValueCount))
	{
		slice2 := (*[1 << 31]C.VkClearValue)(unsafe.Pointer(c.pClearValues))[:len(g.ClearValues):len(g.ClearValues)]
		for i2, _ := range g.ClearValues {
			g.Raw = slice2[i2]
		}
	}
}

type SubpassContents int

const (
	SUBPASS_CONTENTS_INLINE                    SubpassContents = 0
	SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS SubpassContents = 1
	SUBPASS_CONTENTS_BEGIN_RANGE               SubpassContents = SUBPASS_CONTENTS_INLINE
	SUBPASS_CONTENTS_END_RANGE                 SubpassContents = SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	SUBPASS_CONTENTS_RANGE_SIZE                SubpassContents = (SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS - SUBPASS_CONTENTS_INLINE + 1)
	SUBPASS_CONTENTS_MAX_ENUM                  SubpassContents = 2147483647
)

func CmdBeginRenderPass(commandBuffer CommandBuffer, renderPassBegin *RenderPassBeginInfo, contents SubpassContents) {
	var c struct {
		commandBuffer    C.VkCommandBuffer
		pRenderPassBegin *C.VkRenderPassBeginInfo
		contents         C.VkSubpassContents
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.pRenderPassBegin = (*C.VkRenderPassBeginInfo)(_sa.alloc(C.sizeof_VkRenderPassBeginInfo))
		renderPassBegin.toC(c.pRenderPassBegin, _sa)
	}
	c.contents = C.VkSubpassContents(contents)
	C.vkCmdBeginRenderPass(c.commandBuffer, c.pRenderPassBegin, c.contents)
}
func CmdNextSubpass(commandBuffer CommandBuffer, contents SubpassContents) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		contents      C.VkSubpassContents
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.contents = C.VkSubpassContents(contents)
	C.vkCmdNextSubpass(c.commandBuffer, c.contents)
}
func CmdEndRenderPass(commandBuffer CommandBuffer) {
	var c struct{ commandBuffer C.VkCommandBuffer }
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	C.vkCmdEndRenderPass(c.commandBuffer)
}
func CmdExecuteCommands(commandBuffer CommandBuffer, commandBuffers []CommandBuffer) {
	var c struct {
		commandBuffer      C.VkCommandBuffer
		commandBufferCount C.uint32_t
		pCommandBuffers    *C.VkCommandBuffer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.commandBufferCount = C.uint32_t(len(commandBuffers))
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(commandBuffers))))
		slice3 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(commandBuffers):len(commandBuffers)]
		for i3, _ := range commandBuffers {
			slice3[i3] = C.VkCommandBuffer(commandBuffers[i3])
		}
	}
	C.vkCmdExecuteCommands(c.commandBuffer, c.commandBufferCount, c.pCommandBuffers)
}
func EnumerateInstanceVersion(apiVersion *uint32) (_ret Result) {
	var c struct {
		pApiVersion *C.uint32_t
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pApiVersion = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pApiVersion = C.uint32_t(*apiVersion)
	}
	c._ret = C.vkEnumerateInstanceVersion(c.pApiVersion)
	_ret = Result(c._ret)
	*apiVersion = uint32(*c.pApiVersion)
	return
}

type BindBufferMemoryInfo struct {
	Type         StructureType
	Next         unsafe.Pointer
	Buffer       Buffer
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func (g *BindBufferMemoryInfo) toC(c *C.VkBindBufferMemoryInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.buffer = C.VkBuffer(g.Buffer)
	c.memory = C.VkDeviceMemory(g.Memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MemoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *BindBufferMemoryInfo) fromC(c *C.VkBindBufferMemoryInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Buffer = Buffer(c.buffer)
	g.Memory = DeviceMemory(c.memory)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.memoryOffset))
		g.MemoryOffset = DeviceSize(temp_in_VkDeviceSize)
	}
}
func BindBufferMemory2(device Device, bindInfos []BindBufferMemoryInfo) (_ret Result) {
	var c struct {
		device        C.VkDevice
		bindInfoCount C.uint32_t
		pBindInfos    *C.VkBindBufferMemoryInfo
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.bindInfoCount = C.uint32_t(len(bindInfos))
	{
		c.pBindInfos = (*C.VkBindBufferMemoryInfo)(_sa.alloc(C.sizeof_VkBindBufferMemoryInfo * uint(len(bindInfos))))
		slice3 := (*[1 << 31]C.VkBindBufferMemoryInfo)(unsafe.Pointer(c.pBindInfos))[:len(bindInfos):len(bindInfos)]
		for i3, _ := range bindInfos {
			bindInfos[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkBindBufferMemory2(c.device, c.bindInfoCount, c.pBindInfos)
	_ret = Result(c._ret)
	return
}

type BindImageMemoryInfo struct {
	Type         StructureType
	Next         unsafe.Pointer
	Image        Image
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func (g *BindImageMemoryInfo) toC(c *C.VkBindImageMemoryInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.image = C.VkImage(g.Image)
	c.memory = C.VkDeviceMemory(g.Memory)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(g.MemoryOffset))
		c.memoryOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
}
func (g *BindImageMemoryInfo) fromC(c *C.VkBindImageMemoryInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Image = Image(c.image)
	g.Memory = DeviceMemory(c.memory)
	{
		var temp_in_VkDeviceSize uint64
		temp_in_VkDeviceSize = uint64((C.uint64_t)(c.memoryOffset))
		g.MemoryOffset = DeviceSize(temp_in_VkDeviceSize)
	}
}
func BindImageMemory2(device Device, bindInfos []BindImageMemoryInfo) (_ret Result) {
	var c struct {
		device        C.VkDevice
		bindInfoCount C.uint32_t
		pBindInfos    *C.VkBindImageMemoryInfo
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.bindInfoCount = C.uint32_t(len(bindInfos))
	{
		c.pBindInfos = (*C.VkBindImageMemoryInfo)(_sa.alloc(C.sizeof_VkBindImageMemoryInfo * uint(len(bindInfos))))
		slice3 := (*[1 << 31]C.VkBindImageMemoryInfo)(unsafe.Pointer(c.pBindInfos))[:len(bindInfos):len(bindInfos)]
		for i3, _ := range bindInfos {
			bindInfos[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkBindImageMemory2(c.device, c.bindInfoCount, c.pBindInfos)
	_ret = Result(c._ret)
	return
}

type PeerMemoryFeatureFlags Flags

func GetDeviceGroupPeerMemoryFeatures(device Device, heapIndex uint32, localDeviceIndex uint32, peerMemoryFeatures []PeerMemoryFeatureFlags) {
	var c struct {
		device              C.VkDevice
		heapIndex           C.uint32_t
		localDeviceIndex    C.uint32_t
		remoteDeviceIndex   C.uint32_t
		pPeerMemoryFeatures *C.VkPeerMemoryFeatureFlags
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.heapIndex = C.uint32_t(heapIndex)
	c.localDeviceIndex = C.uint32_t(localDeviceIndex)
	c.remoteDeviceIndex = C.uint32_t(len(peerMemoryFeatures))
	{
		c.pPeerMemoryFeatures = (*C.VkPeerMemoryFeatureFlags)(_sa.alloc(C.sizeof_VkPeerMemoryFeatureFlags * uint(len(peerMemoryFeatures))))
		slice3 := (*[1 << 31]C.VkPeerMemoryFeatureFlags)(unsafe.Pointer(c.pPeerMemoryFeatures))[:len(peerMemoryFeatures):len(peerMemoryFeatures)]
		for i3, _ := range peerMemoryFeatures {
			{
				var temp_in_VkPeerMemoryFeatureFlags C.VkFlags
				{
					var temp_in_VkFlags C.uint32_t
					temp_in_VkFlags = C.uint32_t((uint32)((Flags)(peerMemoryFeatures[i3])))
					temp_in_VkPeerMemoryFeatureFlags = C.VkFlags(temp_in_VkFlags)
				}
				slice3[i3] = C.VkPeerMemoryFeatureFlags(temp_in_VkPeerMemoryFeatureFlags)
			}
		}
	}
	C.vkGetDeviceGroupPeerMemoryFeatures(c.device, c.heapIndex, c.localDeviceIndex, c.remoteDeviceIndex, c.pPeerMemoryFeatures)
}
func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		deviceMask    C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.deviceMask = C.uint32_t(deviceMask)
	C.vkCmdSetDeviceMask(c.commandBuffer, c.deviceMask)
}
func CmdDispatchBase(commandBuffer CommandBuffer, baseGroupX uint32, baseGroupY uint32, baseGroupZ uint32, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		baseGroupX    C.uint32_t
		baseGroupY    C.uint32_t
		baseGroupZ    C.uint32_t
		groupCountX   C.uint32_t
		groupCountY   C.uint32_t
		groupCountZ   C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.baseGroupX = C.uint32_t(baseGroupX)
	c.baseGroupY = C.uint32_t(baseGroupY)
	c.baseGroupZ = C.uint32_t(baseGroupZ)
	c.groupCountX = C.uint32_t(groupCountX)
	c.groupCountY = C.uint32_t(groupCountY)
	c.groupCountZ = C.uint32_t(groupCountZ)
	C.vkCmdDispatchBase(c.commandBuffer, c.baseGroupX, c.baseGroupY, c.baseGroupZ, c.groupCountX, c.groupCountY, c.groupCountZ)
}

type PhysicalDeviceGroupProperties struct {
	Type                StructureType
	Next                unsafe.Pointer
	PhysicalDeviceCount uint32
	PhysicalDevices     [32]PhysicalDevice
	SubsetAllocation    bool
}

func (g *PhysicalDeviceGroupProperties) toC(c *C.VkPhysicalDeviceGroupProperties) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.physicalDeviceCount = C.uint32_t(g.PhysicalDeviceCount)
	for i, _ := range g.PhysicalDevices {
		c.physicalDevices[i] = C.VkPhysicalDevice(g.PhysicalDevices[i])
	}
	if g.SubsetAllocation {
		c.subsetAllocation = 1
	} else {
		c.subsetAllocation = 0
	}
}
func (g *PhysicalDeviceGroupProperties) fromC(c *C.VkPhysicalDeviceGroupProperties) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.PhysicalDeviceCount = uint32(c.physicalDeviceCount)
	for i, _ := range g.PhysicalDevices {
		g.PhysicalDevices[i] = PhysicalDevice(c.physicalDevices[i])
	}
	g.SubsetAllocation = c.subsetAllocation != 0
}
func EnumeratePhysicalDeviceGroups(instance Instance, physicalDeviceGroupCount *uint32, physicalDeviceGroupProperties []PhysicalDeviceGroupProperties) (_ret Result) {
	var c struct {
		instance                       C.VkInstance
		pPhysicalDeviceGroupCount      *C.uint32_t
		pPhysicalDeviceGroupProperties *C.VkPhysicalDeviceGroupProperties
		_ret                           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pPhysicalDeviceGroupCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPhysicalDeviceGroupCount = C.uint32_t(*physicalDeviceGroupCount)
	}
	{
		c.pPhysicalDeviceGroupProperties = (*C.VkPhysicalDeviceGroupProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceGroupProperties * uint(len(physicalDeviceGroupProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceGroupProperties)(unsafe.Pointer(c.pPhysicalDeviceGroupProperties))[:len(physicalDeviceGroupProperties):len(physicalDeviceGroupProperties)]
		for i3, _ := range physicalDeviceGroupProperties {
			physicalDeviceGroupProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumeratePhysicalDeviceGroups(c.instance, c.pPhysicalDeviceGroupCount, c.pPhysicalDeviceGroupProperties)
	_ret = Result(c._ret)
	*physicalDeviceGroupCount = uint32(*c.pPhysicalDeviceGroupCount)
	return
}

type ImageMemoryRequirementsInfo2 struct {
	Type  StructureType
	Next  unsafe.Pointer
	Image Image
}

func (g *ImageMemoryRequirementsInfo2) toC(c *C.VkImageMemoryRequirementsInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.image = C.VkImage(g.Image)
}
func (g *ImageMemoryRequirementsInfo2) fromC(c *C.VkImageMemoryRequirementsInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Image = Image(c.image)
}

type MemoryRequirements2 struct {
	Type               StructureType
	Next               unsafe.Pointer
	MemoryRequirements MemoryRequirements
}

func (g *MemoryRequirements2) toC(c *C.VkMemoryRequirements2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.MemoryRequirements.toC(&c.memoryRequirements)
}
func (g *MemoryRequirements2) fromC(c *C.VkMemoryRequirements2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.MemoryRequirements.fromC(&c.memoryRequirements)
}
func GetImageMemoryRequirements2(device Device, info *ImageMemoryRequirementsInfo2, memoryRequirements []MemoryRequirements2) {
	var c struct {
		device              C.VkDevice
		pInfo               *C.VkImageMemoryRequirementsInfo2
		pMemoryRequirements *C.VkMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkImageMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkImageMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements2)(_sa.alloc(C.sizeof_VkMemoryRequirements2 * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements2)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageMemoryRequirements2(c.device, c.pInfo, c.pMemoryRequirements)
}

type BufferMemoryRequirementsInfo2 struct {
	Type   StructureType
	Next   unsafe.Pointer
	Buffer Buffer
}

func (g *BufferMemoryRequirementsInfo2) toC(c *C.VkBufferMemoryRequirementsInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.buffer = C.VkBuffer(g.Buffer)
}
func (g *BufferMemoryRequirementsInfo2) fromC(c *C.VkBufferMemoryRequirementsInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Buffer = Buffer(c.buffer)
}
func GetBufferMemoryRequirements2(device Device, info *BufferMemoryRequirementsInfo2, memoryRequirements []MemoryRequirements2) {
	var c struct {
		device              C.VkDevice
		pInfo               *C.VkBufferMemoryRequirementsInfo2
		pMemoryRequirements *C.VkMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkBufferMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkBufferMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements2)(_sa.alloc(C.sizeof_VkMemoryRequirements2 * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements2)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetBufferMemoryRequirements2(c.device, c.pInfo, c.pMemoryRequirements)
}

type ImageSparseMemoryRequirementsInfo2 struct {
	Type  StructureType
	Next  unsafe.Pointer
	Image Image
}

func (g *ImageSparseMemoryRequirementsInfo2) toC(c *C.VkImageSparseMemoryRequirementsInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.image = C.VkImage(g.Image)
}
func (g *ImageSparseMemoryRequirementsInfo2) fromC(c *C.VkImageSparseMemoryRequirementsInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Image = Image(c.image)
}

type SparseImageMemoryRequirements2 struct {
	Type               StructureType
	Next               unsafe.Pointer
	MemoryRequirements SparseImageMemoryRequirements
}

func (g *SparseImageMemoryRequirements2) toC(c *C.VkSparseImageMemoryRequirements2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.MemoryRequirements.toC(&c.memoryRequirements)
}
func (g *SparseImageMemoryRequirements2) fromC(c *C.VkSparseImageMemoryRequirements2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.MemoryRequirements.fromC(&c.memoryRequirements)
}
func GetImageSparseMemoryRequirements2(device Device, info *ImageSparseMemoryRequirementsInfo2, sparseMemoryRequirementCount *uint32, sparseMemoryRequirements []SparseImageMemoryRequirements2) {
	var c struct {
		device                        C.VkDevice
		pInfo                         *C.VkImageSparseMemoryRequirementsInfo2
		pSparseMemoryRequirementCount *C.uint32_t
		pSparseMemoryRequirements     *C.VkSparseImageMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkImageSparseMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkImageSparseMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pSparseMemoryRequirementCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSparseMemoryRequirementCount = C.uint32_t(*sparseMemoryRequirementCount)
	}
	{
		c.pSparseMemoryRequirements = (*C.VkSparseImageMemoryRequirements2)(_sa.alloc(C.sizeof_VkSparseImageMemoryRequirements2 * uint(len(sparseMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkSparseImageMemoryRequirements2)(unsafe.Pointer(c.pSparseMemoryRequirements))[:len(sparseMemoryRequirements):len(sparseMemoryRequirements)]
		for i3, _ := range sparseMemoryRequirements {
			sparseMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageSparseMemoryRequirements2(c.device, c.pInfo, c.pSparseMemoryRequirementCount, c.pSparseMemoryRequirements)
	*sparseMemoryRequirementCount = uint32(*c.pSparseMemoryRequirementCount)
}

type PhysicalDeviceFeatures2 struct {
	Type     StructureType
	Next     unsafe.Pointer
	Features PhysicalDeviceFeatures
}

func (g *PhysicalDeviceFeatures2) toC(c *C.VkPhysicalDeviceFeatures2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.Features.toC(&c.features)
}
func (g *PhysicalDeviceFeatures2) fromC(c *C.VkPhysicalDeviceFeatures2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Features.fromC(&c.features)
}
func GetPhysicalDeviceFeatures2(physicalDevice PhysicalDevice, features []PhysicalDeviceFeatures2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFeatures      *C.VkPhysicalDeviceFeatures2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFeatures = (*C.VkPhysicalDeviceFeatures2)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures2 * uint(len(features))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceFeatures2)(unsafe.Pointer(c.pFeatures))[:len(features):len(features)]
		for i3, _ := range features {
			features[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFeatures2(c.physicalDevice, c.pFeatures)
}

type PhysicalDeviceProperties2 struct {
	Type       StructureType
	Next       unsafe.Pointer
	Properties PhysicalDeviceProperties
}

func (g *PhysicalDeviceProperties2) toC(c *C.VkPhysicalDeviceProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.Properties.toC(&c.properties)
}
func (g *PhysicalDeviceProperties2) fromC(c *C.VkPhysicalDeviceProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Properties.fromC(&c.properties)
}
func GetPhysicalDeviceProperties2(physicalDevice PhysicalDevice, properties []PhysicalDeviceProperties2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pProperties    *C.VkPhysicalDeviceProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pProperties = (*C.VkPhysicalDeviceProperties2)(_sa.alloc(C.sizeof_VkPhysicalDeviceProperties2 * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceProperties2)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceProperties2(c.physicalDevice, c.pProperties)
}

type FormatProperties2 struct {
	Type             StructureType
	Next             unsafe.Pointer
	FormatProperties FormatProperties
}

func (g *FormatProperties2) toC(c *C.VkFormatProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.FormatProperties.toC(&c.formatProperties)
}
func (g *FormatProperties2) fromC(c *C.VkFormatProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.FormatProperties.fromC(&c.formatProperties)
}
func GetPhysicalDeviceFormatProperties2(physicalDevice PhysicalDevice, format Format, formatProperties []FormatProperties2) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		format            C.VkFormat
		pFormatProperties *C.VkFormatProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.format = C.VkFormat(format)
	{
		c.pFormatProperties = (*C.VkFormatProperties2)(_sa.alloc(C.sizeof_VkFormatProperties2 * uint(len(formatProperties))))
		slice3 := (*[1 << 31]C.VkFormatProperties2)(unsafe.Pointer(c.pFormatProperties))[:len(formatProperties):len(formatProperties)]
		for i3, _ := range formatProperties {
			formatProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFormatProperties2(c.physicalDevice, c.format, c.pFormatProperties)
}

type PhysicalDeviceImageFormatInfo2 struct {
	Type   StructureType
	Next   unsafe.Pointer
	Format Format
	Type   ImageType
	Tiling ImageTiling
	Usage  ImageUsageFlags
	Flags  ImageCreateFlags
}

func (g *PhysicalDeviceImageFormatInfo2) toC(c *C.VkPhysicalDeviceImageFormatInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.format = C.VkFormat(g.Format)
	c._type = C.VkImageType(g.Type)
	c.tiling = C.VkImageTiling(g.Tiling)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Usage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	{
		var temp_in_VkImageCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkImageCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkImageCreateFlags(temp_in_VkImageCreateFlags)
	}
}
func (g *PhysicalDeviceImageFormatInfo2) fromC(c *C.VkPhysicalDeviceImageFormatInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Format = Format(c.format)
	g.Type = ImageType(c._type)
	g.Tiling = ImageTiling(c.tiling)
	{
		var temp_in_VkImageUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			temp_in_VkImageUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Usage = ImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	{
		var temp_in_VkImageCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkImageCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = ImageCreateFlags(temp_in_VkImageCreateFlags)
	}
}

type ImageFormatProperties2 struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	ImageFormatProperties ImageFormatProperties
}

func (g *ImageFormatProperties2) toC(c *C.VkImageFormatProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.ImageFormatProperties.toC(&c.imageFormatProperties)
}
func (g *ImageFormatProperties2) fromC(c *C.VkImageFormatProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.ImageFormatProperties.fromC(&c.imageFormatProperties)
}
func GetPhysicalDeviceImageFormatProperties2(physicalDevice PhysicalDevice, imageFormatInfo *PhysicalDeviceImageFormatInfo2, imageFormatProperties []ImageFormatProperties2) (_ret Result) {
	var c struct {
		physicalDevice         C.VkPhysicalDevice
		pImageFormatInfo       *C.VkPhysicalDeviceImageFormatInfo2
		pImageFormatProperties *C.VkImageFormatProperties2
		_ret                   C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pImageFormatInfo = (*C.VkPhysicalDeviceImageFormatInfo2)(_sa.alloc(C.sizeof_VkPhysicalDeviceImageFormatInfo2))
		imageFormatInfo.toC(c.pImageFormatInfo)
	}
	{
		c.pImageFormatProperties = (*C.VkImageFormatProperties2)(_sa.alloc(C.sizeof_VkImageFormatProperties2 * uint(len(imageFormatProperties))))
		slice3 := (*[1 << 31]C.VkImageFormatProperties2)(unsafe.Pointer(c.pImageFormatProperties))[:len(imageFormatProperties):len(imageFormatProperties)]
		for i3, _ := range imageFormatProperties {
			imageFormatProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceImageFormatProperties2(c.physicalDevice, c.pImageFormatInfo, c.pImageFormatProperties)
	_ret = Result(c._ret)
	return
}

type QueueFamilyProperties2 struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	QueueFamilyProperties QueueFamilyProperties
}

func (g *QueueFamilyProperties2) toC(c *C.VkQueueFamilyProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.QueueFamilyProperties.toC(&c.queueFamilyProperties)
}
func (g *QueueFamilyProperties2) fromC(c *C.VkQueueFamilyProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.QueueFamilyProperties.fromC(&c.queueFamilyProperties)
}
func GetPhysicalDeviceQueueFamilyProperties2(physicalDevice PhysicalDevice, queueFamilyPropertyCount *uint32, queueFamilyProperties []QueueFamilyProperties2) {
	var c struct {
		physicalDevice            C.VkPhysicalDevice
		pQueueFamilyPropertyCount *C.uint32_t
		pQueueFamilyProperties    *C.VkQueueFamilyProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pQueueFamilyPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pQueueFamilyPropertyCount = C.uint32_t(*queueFamilyPropertyCount)
	}
	{
		c.pQueueFamilyProperties = (*C.VkQueueFamilyProperties2)(_sa.alloc(C.sizeof_VkQueueFamilyProperties2 * uint(len(queueFamilyProperties))))
		slice3 := (*[1 << 31]C.VkQueueFamilyProperties2)(unsafe.Pointer(c.pQueueFamilyProperties))[:len(queueFamilyProperties):len(queueFamilyProperties)]
		for i3, _ := range queueFamilyProperties {
			queueFamilyProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceQueueFamilyProperties2(c.physicalDevice, c.pQueueFamilyPropertyCount, c.pQueueFamilyProperties)
	*queueFamilyPropertyCount = uint32(*c.pQueueFamilyPropertyCount)
}

type PhysicalDeviceMemoryProperties2 struct {
	Type             StructureType
	Next             unsafe.Pointer
	MemoryProperties PhysicalDeviceMemoryProperties
}

func (g *PhysicalDeviceMemoryProperties2) toC(c *C.VkPhysicalDeviceMemoryProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.MemoryProperties.toC(&c.memoryProperties)
}
func (g *PhysicalDeviceMemoryProperties2) fromC(c *C.VkPhysicalDeviceMemoryProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.MemoryProperties.fromC(&c.memoryProperties)
}
func GetPhysicalDeviceMemoryProperties2(physicalDevice PhysicalDevice, memoryProperties []PhysicalDeviceMemoryProperties2) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		pMemoryProperties *C.VkPhysicalDeviceMemoryProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pMemoryProperties = (*C.VkPhysicalDeviceMemoryProperties2)(_sa.alloc(C.sizeof_VkPhysicalDeviceMemoryProperties2 * uint(len(memoryProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceMemoryProperties2)(unsafe.Pointer(c.pMemoryProperties))[:len(memoryProperties):len(memoryProperties)]
		for i3, _ := range memoryProperties {
			memoryProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceMemoryProperties2(c.physicalDevice, c.pMemoryProperties)
}

type PhysicalDeviceSparseImageFormatInfo2 struct {
	Type    StructureType
	Next    unsafe.Pointer
	Format  Format
	Type    ImageType
	Samples SampleCountFlagBits
	Usage   ImageUsageFlags
	Tiling  ImageTiling
}

func (g *PhysicalDeviceSparseImageFormatInfo2) toC(c *C.VkPhysicalDeviceSparseImageFormatInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.format = C.VkFormat(g.Format)
	c._type = C.VkImageType(g.Type)
	c.samples = C.VkSampleCountFlagBits(g.Samples)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Usage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	c.tiling = C.VkImageTiling(g.Tiling)
}
func (g *PhysicalDeviceSparseImageFormatInfo2) fromC(c *C.VkPhysicalDeviceSparseImageFormatInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Format = Format(c.format)
	g.Type = ImageType(c._type)
	g.Samples = SampleCountFlagBits(c.samples)
	{
		var temp_in_VkImageUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			temp_in_VkImageUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Usage = ImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	g.Tiling = ImageTiling(c.tiling)
}

type SparseImageFormatProperties2 struct {
	Type       StructureType
	Next       unsafe.Pointer
	Properties SparseImageFormatProperties
}

func (g *SparseImageFormatProperties2) toC(c *C.VkSparseImageFormatProperties2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.Properties.toC(&c.properties)
}
func (g *SparseImageFormatProperties2) fromC(c *C.VkSparseImageFormatProperties2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Properties.fromC(&c.properties)
}
func GetPhysicalDeviceSparseImageFormatProperties2(physicalDevice PhysicalDevice, formatInfo *PhysicalDeviceSparseImageFormatInfo2, propertyCount *uint32, properties []SparseImageFormatProperties2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFormatInfo    *C.VkPhysicalDeviceSparseImageFormatInfo2
		pPropertyCount *C.uint32_t
		pProperties    *C.VkSparseImageFormatProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFormatInfo = (*C.VkPhysicalDeviceSparseImageFormatInfo2)(_sa.alloc(C.sizeof_VkPhysicalDeviceSparseImageFormatInfo2))
		formatInfo.toC(c.pFormatInfo)
	}
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkSparseImageFormatProperties2)(_sa.alloc(C.sizeof_VkSparseImageFormatProperties2 * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkSparseImageFormatProperties2)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceSparseImageFormatProperties2(c.physicalDevice, c.pFormatInfo, c.pPropertyCount, c.pProperties)
	*propertyCount = uint32(*c.pPropertyCount)
}

type CommandPoolTrimFlags Flags

func TrimCommandPool(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	var c struct {
		device      C.VkDevice
		commandPool C.VkCommandPool
		flags       C.VkCommandPoolTrimFlags
	}
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	{
		var temp_in_VkCommandPoolTrimFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkCommandPoolTrimFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandPoolTrimFlags(temp_in_VkCommandPoolTrimFlags)
	}
	C.vkTrimCommandPool(c.device, c.commandPool, c.flags)
}

type DeviceQueueInfo2 struct {
	Type             StructureType
	Next             unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueIndex       uint32
}

func (g *DeviceQueueInfo2) toC(c *C.VkDeviceQueueInfo2) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDeviceQueueCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDeviceQueueCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDeviceQueueCreateFlags(temp_in_VkDeviceQueueCreateFlags)
	}
	c.queueFamilyIndex = C.uint32_t(g.QueueFamilyIndex)
	c.queueIndex = C.uint32_t(g.QueueIndex)
}
func (g *DeviceQueueInfo2) fromC(c *C.VkDeviceQueueInfo2) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDeviceQueueCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDeviceQueueCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DeviceQueueCreateFlags(temp_in_VkDeviceQueueCreateFlags)
	}
	g.QueueFamilyIndex = uint32(c.queueFamilyIndex)
	g.QueueIndex = uint32(c.queueIndex)
}
func GetDeviceQueue2(device Device, queueInfo *DeviceQueueInfo2, queue *Queue) {
	var c struct {
		device     C.VkDevice
		pQueueInfo *C.VkDeviceQueueInfo2
		pQueue     *C.VkQueue
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pQueueInfo = (*C.VkDeviceQueueInfo2)(_sa.alloc(C.sizeof_VkDeviceQueueInfo2))
		queueInfo.toC(c.pQueueInfo)
	}
	{
		c.pQueue = (*C.VkQueue)(_sa.alloc(C.sizeof_VkQueue))
		*c.pQueue = C.VkQueue(*queue)
	}
	C.vkGetDeviceQueue2(c.device, c.pQueueInfo, c.pQueue)
	*queue = Queue(*c.pQueue)
}

type SamplerYcbcrModelConversion int

const (
	SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY       SamplerYcbcrModelConversion = 0
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY     SamplerYcbcrModelConversion = 1
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709          SamplerYcbcrModelConversion = 2
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601          SamplerYcbcrModelConversion = 3
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020         SamplerYcbcrModelConversion = 4
	SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR   SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR      SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR      SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR     SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	SAMPLER_YCBCR_MODEL_CONVERSION_BEGIN_RANGE        SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_END_RANGE          SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	SAMPLER_YCBCR_MODEL_CONVERSION_RANGE_SIZE         SamplerYcbcrModelConversion = (SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020 - SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY + 1)
	SAMPLER_YCBCR_MODEL_CONVERSION_MAX_ENUM           SamplerYcbcrModelConversion = 2147483647
)

type SamplerYcbcrRange int

const (
	SAMPLER_YCBCR_RANGE_ITU_FULL       SamplerYcbcrRange = 0
	SAMPLER_YCBCR_RANGE_ITU_NARROW     SamplerYcbcrRange = 1
	SAMPLER_YCBCR_RANGE_ITU_FULL_KHR   SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_FULL
	SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_NARROW
	SAMPLER_YCBCR_RANGE_BEGIN_RANGE    SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_FULL
	SAMPLER_YCBCR_RANGE_END_RANGE      SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_NARROW
	SAMPLER_YCBCR_RANGE_RANGE_SIZE     SamplerYcbcrRange = (SAMPLER_YCBCR_RANGE_ITU_NARROW - SAMPLER_YCBCR_RANGE_ITU_FULL + 1)
	SAMPLER_YCBCR_RANGE_MAX_ENUM       SamplerYcbcrRange = 2147483647
)

type ChromaLocation int

const (
	CHROMA_LOCATION_COSITED_EVEN     ChromaLocation = 0
	CHROMA_LOCATION_MIDPOINT         ChromaLocation = 1
	CHROMA_LOCATION_COSITED_EVEN_KHR ChromaLocation = CHROMA_LOCATION_COSITED_EVEN
	CHROMA_LOCATION_MIDPOINT_KHR     ChromaLocation = CHROMA_LOCATION_MIDPOINT
	CHROMA_LOCATION_BEGIN_RANGE      ChromaLocation = CHROMA_LOCATION_COSITED_EVEN
	CHROMA_LOCATION_END_RANGE        ChromaLocation = CHROMA_LOCATION_MIDPOINT
	CHROMA_LOCATION_RANGE_SIZE       ChromaLocation = (CHROMA_LOCATION_MIDPOINT - CHROMA_LOCATION_COSITED_EVEN + 1)
	CHROMA_LOCATION_MAX_ENUM         ChromaLocation = 2147483647
)

type SamplerYcbcrConversionCreateInfo struct {
	Type                        StructureType
	Next                        unsafe.Pointer
	Format                      Format
	YcbcrModel                  SamplerYcbcrModelConversion
	YcbcrRange                  SamplerYcbcrRange
	Components                  ComponentMapping
	XChromaOffset               ChromaLocation
	YChromaOffset               ChromaLocation
	ChromaFilter                Filter
	ForceExplicitReconstruction bool
}

func (g *SamplerYcbcrConversionCreateInfo) toC(c *C.VkSamplerYcbcrConversionCreateInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.format = C.VkFormat(g.Format)
	c.ycbcrModel = C.VkSamplerYcbcrModelConversion(g.YcbcrModel)
	c.ycbcrRange = C.VkSamplerYcbcrRange(g.YcbcrRange)
	g.Components.toC(&c.components)
	c.xChromaOffset = C.VkChromaLocation(g.XChromaOffset)
	c.yChromaOffset = C.VkChromaLocation(g.YChromaOffset)
	c.chromaFilter = C.VkFilter(g.ChromaFilter)
	if g.ForceExplicitReconstruction {
		c.forceExplicitReconstruction = 1
	} else {
		c.forceExplicitReconstruction = 0
	}
}
func (g *SamplerYcbcrConversionCreateInfo) fromC(c *C.VkSamplerYcbcrConversionCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Format = Format(c.format)
	g.YcbcrModel = SamplerYcbcrModelConversion(c.ycbcrModel)
	g.YcbcrRange = SamplerYcbcrRange(c.ycbcrRange)
	g.Components.fromC(&c.components)
	g.XChromaOffset = ChromaLocation(c.xChromaOffset)
	g.YChromaOffset = ChromaLocation(c.yChromaOffset)
	g.ChromaFilter = Filter(c.chromaFilter)
	g.ForceExplicitReconstruction = c.forceExplicitReconstruction != 0
}

type SamplerYcbcrConversion C.VkSamplerYcbcrConversion

func CreateSamplerYcbcrConversion(device Device, createInfo *SamplerYcbcrConversionCreateInfo, allocator *AllocationCallbacks, ycbcrConversion *SamplerYcbcrConversion) (_ret Result) {
	var c struct {
		device           C.VkDevice
		pCreateInfo      *C.VkSamplerYcbcrConversionCreateInfo
		pAllocator       *C.VkAllocationCallbacks
		pYcbcrConversion *C.VkSamplerYcbcrConversion
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkSamplerYcbcrConversionCreateInfo)(_sa.alloc(C.sizeof_VkSamplerYcbcrConversionCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pYcbcrConversion = (*C.VkSamplerYcbcrConversion)(_sa.alloc(C.sizeof_VkSamplerYcbcrConversion))
		*c.pYcbcrConversion = C.VkSamplerYcbcrConversion(*ycbcrConversion)
	}
	c._ret = C.vkCreateSamplerYcbcrConversion(c.device, c.pCreateInfo, c.pAllocator, c.pYcbcrConversion)
	_ret = Result(c._ret)
	*ycbcrConversion = SamplerYcbcrConversion(*c.pYcbcrConversion)
	return
}
func DestroySamplerYcbcrConversion(device Device, ycbcrConversion SamplerYcbcrConversion, allocator *AllocationCallbacks) {
	var c struct {
		device          C.VkDevice
		ycbcrConversion C.VkSamplerYcbcrConversion
		pAllocator      *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.ycbcrConversion = C.VkSamplerYcbcrConversion(ycbcrConversion)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySamplerYcbcrConversion(c.device, c.ycbcrConversion, c.pAllocator)
}

type DescriptorUpdateTemplateCreateFlags Flags
type DescriptorUpdateTemplateEntry struct {
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  DescriptorType
	Offset          uint
	Stride          uint
}

func (g *DescriptorUpdateTemplateEntry) toC(c *C.VkDescriptorUpdateTemplateEntry) {
	c.dstBinding = C.uint32_t(g.DstBinding)
	c.dstArrayElement = C.uint32_t(g.DstArrayElement)
	c.descriptorCount = C.uint32_t(g.DescriptorCount)
	c.descriptorType = C.VkDescriptorType(g.DescriptorType)
	c.offset = C.size_t(g.Offset)
	c.stride = C.size_t(g.Stride)
}
func (g *DescriptorUpdateTemplateEntry) fromC(c *C.VkDescriptorUpdateTemplateEntry) {
	g.DstBinding = uint32(c.dstBinding)
	g.DstArrayElement = uint32(c.dstArrayElement)
	g.DescriptorCount = uint32(c.descriptorCount)
	g.DescriptorType = DescriptorType(c.descriptorType)
	g.Offset = uint(c.offset)
	g.Stride = uint(c.stride)
}

type DescriptorUpdateTemplateType int

const (
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET       DescriptorUpdateTemplateType = 0
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR DescriptorUpdateTemplateType = 1
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR   DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_BEGIN_RANGE          DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_END_RANGE            DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_RANGE_SIZE           DescriptorUpdateTemplateType = (DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET - DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET + 1)
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_MAX_ENUM             DescriptorUpdateTemplateType = 2147483647
)

type DescriptorUpdateTemplateCreateInfo struct {
	Type                    StructureType
	Next                    unsafe.Pointer
	Flags                   DescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntries []DescriptorUpdateTemplateEntry
	TemplateType            DescriptorUpdateTemplateType
	DescriptorSetLayout     DescriptorSetLayout
	PipelineBindPoint       PipelineBindPoint
	PipelineLayout          PipelineLayout
	Set                     uint32
}

func (g *DescriptorUpdateTemplateCreateInfo) toC(c *C.VkDescriptorUpdateTemplateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDescriptorUpdateTemplateCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDescriptorUpdateTemplateCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDescriptorUpdateTemplateCreateFlags(temp_in_VkDescriptorUpdateTemplateCreateFlags)
	}
	c.descriptorUpdateEntryCount = C.uint32_t(len(g.DescriptorUpdateEntries))
	{
		c.pDescriptorUpdateEntries = (*C.VkDescriptorUpdateTemplateEntry)(_sa.alloc(C.sizeof_VkDescriptorUpdateTemplateEntry * uint(len(g.DescriptorUpdateEntries))))
		slice2 := (*[1 << 31]C.VkDescriptorUpdateTemplateEntry)(unsafe.Pointer(c.pDescriptorUpdateEntries))[:len(g.DescriptorUpdateEntries):len(g.DescriptorUpdateEntries)]
		for i2, _ := range g.DescriptorUpdateEntries {
			g.DescriptorUpdateEntries[i2].toC(&slice2[i2])
		}
	}
	c.templateType = C.VkDescriptorUpdateTemplateType(g.TemplateType)
	c.descriptorSetLayout = C.VkDescriptorSetLayout(g.DescriptorSetLayout)
	c.pipelineBindPoint = C.VkPipelineBindPoint(g.PipelineBindPoint)
	c.pipelineLayout = C.VkPipelineLayout(g.PipelineLayout)
	c.set = C.uint32_t(g.Set)
}
func (g *DescriptorUpdateTemplateCreateInfo) fromC(c *C.VkDescriptorUpdateTemplateCreateInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDescriptorUpdateTemplateCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDescriptorUpdateTemplateCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = DescriptorUpdateTemplateCreateFlags(temp_in_VkDescriptorUpdateTemplateCreateFlags)
	}
	g.DescriptorUpdateEntries = make([]DescriptorUpdateTemplateEntry, int(c.descriptorUpdateEntryCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorUpdateTemplateEntry)(unsafe.Pointer(c.pDescriptorUpdateEntries))[:len(g.DescriptorUpdateEntries):len(g.DescriptorUpdateEntries)]
		for i2, _ := range g.DescriptorUpdateEntries {
			g.DescriptorUpdateEntries[i2].fromC(&slice2[i2])
		}
	}
	g.TemplateType = DescriptorUpdateTemplateType(c.templateType)
	g.DescriptorSetLayout = DescriptorSetLayout(c.descriptorSetLayout)
	g.PipelineBindPoint = PipelineBindPoint(c.pipelineBindPoint)
	g.PipelineLayout = PipelineLayout(c.pipelineLayout)
	g.Set = uint32(c.set)
}

type DescriptorUpdateTemplate C.VkDescriptorUpdateTemplate

func CreateDescriptorUpdateTemplate(device Device, createInfo *DescriptorUpdateTemplateCreateInfo, allocator *AllocationCallbacks, descriptorUpdateTemplate *DescriptorUpdateTemplate) (_ret Result) {
	var c struct {
		device                    C.VkDevice
		pCreateInfo               *C.VkDescriptorUpdateTemplateCreateInfo
		pAllocator                *C.VkAllocationCallbacks
		pDescriptorUpdateTemplate *C.VkDescriptorUpdateTemplate
		_ret                      C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorUpdateTemplateCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorUpdateTemplateCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDescriptorUpdateTemplate = (*C.VkDescriptorUpdateTemplate)(_sa.alloc(C.sizeof_VkDescriptorUpdateTemplate))
		*c.pDescriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(*descriptorUpdateTemplate)
	}
	c._ret = C.vkCreateDescriptorUpdateTemplate(c.device, c.pCreateInfo, c.pAllocator, c.pDescriptorUpdateTemplate)
	_ret = Result(c._ret)
	*descriptorUpdateTemplate = DescriptorUpdateTemplate(*c.pDescriptorUpdateTemplate)
	return
}
func DestroyDescriptorUpdateTemplate(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, allocator *AllocationCallbacks) {
	var c struct {
		device                   C.VkDevice
		descriptorUpdateTemplate C.VkDescriptorUpdateTemplate
		pAllocator               *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorUpdateTemplate(c.device, c.descriptorUpdateTemplate, c.pAllocator)
}
func UpdateDescriptorSetWithTemplate(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, data []byte) {
	var c struct {
		device                   C.VkDevice
		descriptorSet            C.VkDescriptorSet
		descriptorUpdateTemplate C.VkDescriptorUpdateTemplate
		pData                    unsafe.Pointer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorSet = C.VkDescriptorSet(descriptorSet)
	c.descriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate)
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	C.vkUpdateDescriptorSetWithTemplate(c.device, c.descriptorSet, c.descriptorUpdateTemplate, c.pData)
}

type ExternalMemoryHandleTypeFlagBits int

const (
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT                       ExternalMemoryHandleTypeFlagBits = 1
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT                    ExternalMemoryHandleTypeFlagBits = 2
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT                ExternalMemoryHandleTypeFlagBits = 4
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT                   ExternalMemoryHandleTypeFlagBits = 8
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT               ExternalMemoryHandleTypeFlagBits = 16
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT                      ExternalMemoryHandleTypeFlagBits = 32
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT                  ExternalMemoryHandleTypeFlagBits = 64
	EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT                     ExternalMemoryHandleTypeFlagBits = 512
	EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID ExternalMemoryHandleTypeFlagBits = 1024
	EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT             ExternalMemoryHandleTypeFlagBits = 128
	EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT  ExternalMemoryHandleTypeFlagBits = 256
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR                   ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR                ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR            ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR               ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR           ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR                  ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR              ExternalMemoryHandleTypeFlagBits = EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_FLAG_BITS_MAX_ENUM                  ExternalMemoryHandleTypeFlagBits = 2147483647
)

type PhysicalDeviceExternalBufferInfo struct {
	Type       StructureType
	Next       unsafe.Pointer
	Flags      BufferCreateFlags
	Usage      BufferUsageFlags
	HandleType ExternalMemoryHandleTypeFlagBits
}

func (g *PhysicalDeviceExternalBufferInfo) toC(c *C.VkPhysicalDeviceExternalBufferInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkBufferCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkBufferCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkBufferCreateFlags(temp_in_VkBufferCreateFlags)
	}
	{
		var temp_in_VkBufferUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Usage)))
			temp_in_VkBufferUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.usage = C.VkBufferUsageFlags(temp_in_VkBufferUsageFlags)
	}
	c.handleType = C.VkExternalMemoryHandleTypeFlagBits(g.HandleType)
}
func (g *PhysicalDeviceExternalBufferInfo) fromC(c *C.VkPhysicalDeviceExternalBufferInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkBufferCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkBufferCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = BufferCreateFlags(temp_in_VkBufferCreateFlags)
	}
	{
		var temp_in_VkBufferUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			temp_in_VkBufferUsageFlags = Flags(temp_in_VkFlags)
		}
		g.Usage = BufferUsageFlags(temp_in_VkBufferUsageFlags)
	}
	g.HandleType = ExternalMemoryHandleTypeFlagBits(c.handleType)
}

type ExternalMemoryFeatureFlags Flags
type ExternalMemoryHandleTypeFlags Flags
type ExternalMemoryProperties struct {
	ExternalMemoryFeatures        ExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         ExternalMemoryHandleTypeFlags
}

func (g *ExternalMemoryProperties) toC(c *C.VkExternalMemoryProperties) {
	{
		var temp_in_VkExternalMemoryFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExternalMemoryFeatures)))
			temp_in_VkExternalMemoryFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.externalMemoryFeatures = C.VkExternalMemoryFeatureFlags(temp_in_VkExternalMemoryFeatureFlags)
	}
	{
		var temp_in_VkExternalMemoryHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExportFromImportedHandleTypes)))
			temp_in_VkExternalMemoryHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.exportFromImportedHandleTypes = C.VkExternalMemoryHandleTypeFlags(temp_in_VkExternalMemoryHandleTypeFlags)
	}
	{
		var temp_in_VkExternalMemoryHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.CompatibleHandleTypes)))
			temp_in_VkExternalMemoryHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.compatibleHandleTypes = C.VkExternalMemoryHandleTypeFlags(temp_in_VkExternalMemoryHandleTypeFlags)
	}
}
func (g *ExternalMemoryProperties) fromC(c *C.VkExternalMemoryProperties) {
	{
		var temp_in_VkExternalMemoryFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.externalMemoryFeatures)))
			temp_in_VkExternalMemoryFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.ExternalMemoryFeatures = ExternalMemoryFeatureFlags(temp_in_VkExternalMemoryFeatureFlags)
	}
	{
		var temp_in_VkExternalMemoryHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.exportFromImportedHandleTypes)))
			temp_in_VkExternalMemoryHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.ExportFromImportedHandleTypes = ExternalMemoryHandleTypeFlags(temp_in_VkExternalMemoryHandleTypeFlags)
	}
	{
		var temp_in_VkExternalMemoryHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.compatibleHandleTypes)))
			temp_in_VkExternalMemoryHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.CompatibleHandleTypes = ExternalMemoryHandleTypeFlags(temp_in_VkExternalMemoryHandleTypeFlags)
	}
}

type ExternalBufferProperties struct {
	Type                     StructureType
	Next                     unsafe.Pointer
	ExternalMemoryProperties ExternalMemoryProperties
}

func (g *ExternalBufferProperties) toC(c *C.VkExternalBufferProperties) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.ExternalMemoryProperties.toC(&c.externalMemoryProperties)
}
func (g *ExternalBufferProperties) fromC(c *C.VkExternalBufferProperties) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.ExternalMemoryProperties.fromC(&c.externalMemoryProperties)
}
func GetPhysicalDeviceExternalBufferProperties(physicalDevice PhysicalDevice, externalBufferInfo *PhysicalDeviceExternalBufferInfo, externalBufferProperties []ExternalBufferProperties) {
	var c struct {
		physicalDevice            C.VkPhysicalDevice
		pExternalBufferInfo       *C.VkPhysicalDeviceExternalBufferInfo
		pExternalBufferProperties *C.VkExternalBufferProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalBufferInfo = (*C.VkPhysicalDeviceExternalBufferInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalBufferInfo))
		externalBufferInfo.toC(c.pExternalBufferInfo)
	}
	{
		c.pExternalBufferProperties = (*C.VkExternalBufferProperties)(_sa.alloc(C.sizeof_VkExternalBufferProperties * uint(len(externalBufferProperties))))
		slice3 := (*[1 << 31]C.VkExternalBufferProperties)(unsafe.Pointer(c.pExternalBufferProperties))[:len(externalBufferProperties):len(externalBufferProperties)]
		for i3, _ := range externalBufferProperties {
			externalBufferProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalBufferProperties(c.physicalDevice, c.pExternalBufferInfo, c.pExternalBufferProperties)
}

type ExternalFenceHandleTypeFlagBits int

const (
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT            ExternalFenceHandleTypeFlagBits = 1
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT         ExternalFenceHandleTypeFlagBits = 2
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     ExternalFenceHandleTypeFlagBits = 4
	EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT              ExternalFenceHandleTypeFlagBits = 8
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        ExternalFenceHandleTypeFlagBits = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     ExternalFenceHandleTypeFlagBits = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR ExternalFenceHandleTypeFlagBits = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR          ExternalFenceHandleTypeFlagBits = EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       ExternalFenceHandleTypeFlagBits = 2147483647
)

type PhysicalDeviceExternalFenceInfo struct {
	Type       StructureType
	Next       unsafe.Pointer
	HandleType ExternalFenceHandleTypeFlagBits
}

func (g *PhysicalDeviceExternalFenceInfo) toC(c *C.VkPhysicalDeviceExternalFenceInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.handleType = C.VkExternalFenceHandleTypeFlagBits(g.HandleType)
}
func (g *PhysicalDeviceExternalFenceInfo) fromC(c *C.VkPhysicalDeviceExternalFenceInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.HandleType = ExternalFenceHandleTypeFlagBits(c.handleType)
}

type ExternalFenceHandleTypeFlags Flags
type ExternalFenceFeatureFlags Flags
type ExternalFenceProperties struct {
	Type                          StructureType
	Next                          unsafe.Pointer
	ExportFromImportedHandleTypes ExternalFenceHandleTypeFlags
	CompatibleHandleTypes         ExternalFenceHandleTypeFlags
	ExternalFenceFeatures         ExternalFenceFeatureFlags
}

func (g *ExternalFenceProperties) toC(c *C.VkExternalFenceProperties) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkExternalFenceHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExportFromImportedHandleTypes)))
			temp_in_VkExternalFenceHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.exportFromImportedHandleTypes = C.VkExternalFenceHandleTypeFlags(temp_in_VkExternalFenceHandleTypeFlags)
	}
	{
		var temp_in_VkExternalFenceHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.CompatibleHandleTypes)))
			temp_in_VkExternalFenceHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.compatibleHandleTypes = C.VkExternalFenceHandleTypeFlags(temp_in_VkExternalFenceHandleTypeFlags)
	}
	{
		var temp_in_VkExternalFenceFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExternalFenceFeatures)))
			temp_in_VkExternalFenceFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.externalFenceFeatures = C.VkExternalFenceFeatureFlags(temp_in_VkExternalFenceFeatureFlags)
	}
}
func (g *ExternalFenceProperties) fromC(c *C.VkExternalFenceProperties) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkExternalFenceHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.exportFromImportedHandleTypes)))
			temp_in_VkExternalFenceHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.ExportFromImportedHandleTypes = ExternalFenceHandleTypeFlags(temp_in_VkExternalFenceHandleTypeFlags)
	}
	{
		var temp_in_VkExternalFenceHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.compatibleHandleTypes)))
			temp_in_VkExternalFenceHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.CompatibleHandleTypes = ExternalFenceHandleTypeFlags(temp_in_VkExternalFenceHandleTypeFlags)
	}
	{
		var temp_in_VkExternalFenceFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.externalFenceFeatures)))
			temp_in_VkExternalFenceFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.ExternalFenceFeatures = ExternalFenceFeatureFlags(temp_in_VkExternalFenceFeatureFlags)
	}
}
func GetPhysicalDeviceExternalFenceProperties(physicalDevice PhysicalDevice, externalFenceInfo *PhysicalDeviceExternalFenceInfo, externalFenceProperties []ExternalFenceProperties) {
	var c struct {
		physicalDevice           C.VkPhysicalDevice
		pExternalFenceInfo       *C.VkPhysicalDeviceExternalFenceInfo
		pExternalFenceProperties *C.VkExternalFenceProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalFenceInfo = (*C.VkPhysicalDeviceExternalFenceInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalFenceInfo))
		externalFenceInfo.toC(c.pExternalFenceInfo)
	}
	{
		c.pExternalFenceProperties = (*C.VkExternalFenceProperties)(_sa.alloc(C.sizeof_VkExternalFenceProperties * uint(len(externalFenceProperties))))
		slice3 := (*[1 << 31]C.VkExternalFenceProperties)(unsafe.Pointer(c.pExternalFenceProperties))[:len(externalFenceProperties):len(externalFenceProperties)]
		for i3, _ := range externalFenceProperties {
			externalFenceProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalFenceProperties(c.physicalDevice, c.pExternalFenceInfo, c.pExternalFenceProperties)
}

type ExternalSemaphoreHandleTypeFlagBits int

const (
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT            ExternalSemaphoreHandleTypeFlagBits = 1
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT         ExternalSemaphoreHandleTypeFlagBits = 2
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     ExternalSemaphoreHandleTypeFlagBits = 4
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT          ExternalSemaphoreHandleTypeFlagBits = 8
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT              ExternalSemaphoreHandleTypeFlagBits = 16
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        ExternalSemaphoreHandleTypeFlagBits = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     ExternalSemaphoreHandleTypeFlagBits = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR ExternalSemaphoreHandleTypeFlagBits = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR      ExternalSemaphoreHandleTypeFlagBits = EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR          ExternalSemaphoreHandleTypeFlagBits = EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       ExternalSemaphoreHandleTypeFlagBits = 2147483647
)

type PhysicalDeviceExternalSemaphoreInfo struct {
	Type       StructureType
	Next       unsafe.Pointer
	HandleType ExternalSemaphoreHandleTypeFlagBits
}

func (g *PhysicalDeviceExternalSemaphoreInfo) toC(c *C.VkPhysicalDeviceExternalSemaphoreInfo) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(g.HandleType)
}
func (g *PhysicalDeviceExternalSemaphoreInfo) fromC(c *C.VkPhysicalDeviceExternalSemaphoreInfo) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.HandleType = ExternalSemaphoreHandleTypeFlagBits(c.handleType)
}

type ExternalSemaphoreHandleTypeFlags Flags
type ExternalSemaphoreFeatureFlags Flags
type ExternalSemaphoreProperties struct {
	Type                          StructureType
	Next                          unsafe.Pointer
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         ExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatureFlags
}

func (g *ExternalSemaphoreProperties) toC(c *C.VkExternalSemaphoreProperties) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkExternalSemaphoreHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExportFromImportedHandleTypes)))
			temp_in_VkExternalSemaphoreHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.exportFromImportedHandleTypes = C.VkExternalSemaphoreHandleTypeFlags(temp_in_VkExternalSemaphoreHandleTypeFlags)
	}
	{
		var temp_in_VkExternalSemaphoreHandleTypeFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.CompatibleHandleTypes)))
			temp_in_VkExternalSemaphoreHandleTypeFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.compatibleHandleTypes = C.VkExternalSemaphoreHandleTypeFlags(temp_in_VkExternalSemaphoreHandleTypeFlags)
	}
	{
		var temp_in_VkExternalSemaphoreFeatureFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ExternalSemaphoreFeatures)))
			temp_in_VkExternalSemaphoreFeatureFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.externalSemaphoreFeatures = C.VkExternalSemaphoreFeatureFlags(temp_in_VkExternalSemaphoreFeatureFlags)
	}
}
func (g *ExternalSemaphoreProperties) fromC(c *C.VkExternalSemaphoreProperties) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkExternalSemaphoreHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.exportFromImportedHandleTypes)))
			temp_in_VkExternalSemaphoreHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypeFlags(temp_in_VkExternalSemaphoreHandleTypeFlags)
	}
	{
		var temp_in_VkExternalSemaphoreHandleTypeFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.compatibleHandleTypes)))
			temp_in_VkExternalSemaphoreHandleTypeFlags = Flags(temp_in_VkFlags)
		}
		g.CompatibleHandleTypes = ExternalSemaphoreHandleTypeFlags(temp_in_VkExternalSemaphoreHandleTypeFlags)
	}
	{
		var temp_in_VkExternalSemaphoreFeatureFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.externalSemaphoreFeatures)))
			temp_in_VkExternalSemaphoreFeatureFlags = Flags(temp_in_VkFlags)
		}
		g.ExternalSemaphoreFeatures = ExternalSemaphoreFeatureFlags(temp_in_VkExternalSemaphoreFeatureFlags)
	}
}
func GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice PhysicalDevice, externalSemaphoreInfo *PhysicalDeviceExternalSemaphoreInfo, externalSemaphoreProperties []ExternalSemaphoreProperties) {
	var c struct {
		physicalDevice               C.VkPhysicalDevice
		pExternalSemaphoreInfo       *C.VkPhysicalDeviceExternalSemaphoreInfo
		pExternalSemaphoreProperties *C.VkExternalSemaphoreProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalSemaphoreInfo = (*C.VkPhysicalDeviceExternalSemaphoreInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalSemaphoreInfo))
		externalSemaphoreInfo.toC(c.pExternalSemaphoreInfo)
	}
	{
		c.pExternalSemaphoreProperties = (*C.VkExternalSemaphoreProperties)(_sa.alloc(C.sizeof_VkExternalSemaphoreProperties * uint(len(externalSemaphoreProperties))))
		slice3 := (*[1 << 31]C.VkExternalSemaphoreProperties)(unsafe.Pointer(c.pExternalSemaphoreProperties))[:len(externalSemaphoreProperties):len(externalSemaphoreProperties)]
		for i3, _ := range externalSemaphoreProperties {
			externalSemaphoreProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalSemaphoreProperties(c.physicalDevice, c.pExternalSemaphoreInfo, c.pExternalSemaphoreProperties)
}

type DescriptorSetLayoutSupport struct {
	Type      StructureType
	Next      unsafe.Pointer
	Supported bool
}

func (g *DescriptorSetLayoutSupport) toC(c *C.VkDescriptorSetLayoutSupport) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	if g.Supported {
		c.supported = 1
	} else {
		c.supported = 0
	}
}
func (g *DescriptorSetLayoutSupport) fromC(c *C.VkDescriptorSetLayoutSupport) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Supported = c.supported != 0
}
func GetDescriptorSetLayoutSupport(device Device, createInfo *DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkDescriptorSetLayoutCreateInfo
		pSupport    *C.VkDescriptorSetLayoutSupport
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorSetLayoutCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pSupport = (*C.VkDescriptorSetLayoutSupport)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutSupport))
		support.toC(c.pSupport)
	}
	C.vkGetDescriptorSetLayoutSupport(c.device, c.pCreateInfo, c.pSupport)
	support.fromC(c.pSupport)
}

type SurfaceKHR C.VkSurfaceKHR

func DestroySurfaceKHR(instance Instance, surface SurfaceKHR, allocator *AllocationCallbacks) {
	var c struct {
		instance   C.VkInstance
		surface    C.VkSurfaceKHR
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySurfaceKHR(c.instance, c.surface, c.pAllocator)
}
func GetPhysicalDeviceSurfaceSupportKHR(physicalDevice PhysicalDevice, queueFamilyIndex uint32, surface SurfaceKHR, supported *bool) (_ret Result) {
	var c struct {
		physicalDevice   C.VkPhysicalDevice
		queueFamilyIndex C.uint32_t
		surface          C.VkSurfaceKHR
		pSupported       *C.VkBool32
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.queueFamilyIndex = C.uint32_t(queueFamilyIndex)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pSupported = (*C.VkBool32)(_sa.alloc(C.sizeof_VkBool32))
		if *supported {
			*c.pSupported = 1
		} else {
			*c.pSupported = 0
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfaceSupportKHR(c.physicalDevice, c.queueFamilyIndex, c.surface, c.pSupported)
	_ret = Result(c._ret)
	*supported = *c.pSupported != 0
	return
}

type SurfaceTransformFlagsKHR Flags
type SurfaceTransformFlagBitsKHR int

const (
	SURFACE_TRANSFORM_IDENTITY_BIT_KHR                     SurfaceTransformFlagBitsKHR = 1
	SURFACE_TRANSFORM_ROTATE_90_BIT_KHR                    SurfaceTransformFlagBitsKHR = 2
	SURFACE_TRANSFORM_ROTATE_180_BIT_KHR                   SurfaceTransformFlagBitsKHR = 4
	SURFACE_TRANSFORM_ROTATE_270_BIT_KHR                   SurfaceTransformFlagBitsKHR = 8
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR            SurfaceTransformFlagBitsKHR = 16
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR  SurfaceTransformFlagBitsKHR = 32
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR SurfaceTransformFlagBitsKHR = 64
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR SurfaceTransformFlagBitsKHR = 128
	SURFACE_TRANSFORM_INHERIT_BIT_KHR                      SurfaceTransformFlagBitsKHR = 256
	SURFACE_TRANSFORM_FLAG_BITS_MAX_ENUM_KHR               SurfaceTransformFlagBitsKHR = 2147483647
)

type CompositeAlphaFlagsKHR Flags
type SurfaceCapabilitiesKHR struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     SurfaceTransformFlagsKHR
	CurrentTransform        SurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha CompositeAlphaFlagsKHR
	SupportedUsageFlags     ImageUsageFlags
}

func (g *SurfaceCapabilitiesKHR) toC(c *C.VkSurfaceCapabilitiesKHR) {
	c.minImageCount = C.uint32_t(g.MinImageCount)
	c.maxImageCount = C.uint32_t(g.MaxImageCount)
	g.CurrentExtent.toC(&c.currentExtent)
	g.MinImageExtent.toC(&c.minImageExtent)
	g.MaxImageExtent.toC(&c.maxImageExtent)
	c.maxImageArrayLayers = C.uint32_t(g.MaxImageArrayLayers)
	{
		var temp_in_VkSurfaceTransformFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SupportedTransforms)))
			temp_in_VkSurfaceTransformFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.supportedTransforms = C.VkSurfaceTransformFlagsKHR(temp_in_VkSurfaceTransformFlagsKHR)
	}
	c.currentTransform = C.VkSurfaceTransformFlagBitsKHR(g.CurrentTransform)
	{
		var temp_in_VkCompositeAlphaFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SupportedCompositeAlpha)))
			temp_in_VkCompositeAlphaFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.supportedCompositeAlpha = C.VkCompositeAlphaFlagsKHR(temp_in_VkCompositeAlphaFlagsKHR)
	}
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SupportedUsageFlags)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.supportedUsageFlags = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
}
func (g *SurfaceCapabilitiesKHR) fromC(c *C.VkSurfaceCapabilitiesKHR) {
	g.MinImageCount = uint32(c.minImageCount)
	g.MaxImageCount = uint32(c.maxImageCount)
	g.CurrentExtent.fromC(&c.currentExtent)
	g.MinImageExtent.fromC(&c.minImageExtent)
	g.MaxImageExtent.fromC(&c.maxImageExtent)
	g.MaxImageArrayLayers = uint32(c.maxImageArrayLayers)
	{
		var temp_in_VkSurfaceTransformFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.supportedTransforms)))
			temp_in_VkSurfaceTransformFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.SupportedTransforms = SurfaceTransformFlagsKHR(temp_in_VkSurfaceTransformFlagsKHR)
	}
	g.CurrentTransform = SurfaceTransformFlagBitsKHR(c.currentTransform)
	{
		var temp_in_VkCompositeAlphaFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.supportedCompositeAlpha)))
			temp_in_VkCompositeAlphaFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.SupportedCompositeAlpha = CompositeAlphaFlagsKHR(temp_in_VkCompositeAlphaFlagsKHR)
	}
	{
		var temp_in_VkImageUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.supportedUsageFlags)))
			temp_in_VkImageUsageFlags = Flags(temp_in_VkFlags)
		}
		g.SupportedUsageFlags = ImageUsageFlags(temp_in_VkImageUsageFlags)
	}
}
func GetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, surfaceCapabilities []SurfaceCapabilitiesKHR) (_ret Result) {
	var c struct {
		physicalDevice       C.VkPhysicalDevice
		surface              C.VkSurfaceKHR
		pSurfaceCapabilities *C.VkSurfaceCapabilitiesKHR
		_ret                 C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pSurfaceCapabilities = (*C.VkSurfaceCapabilitiesKHR)(_sa.alloc(C.sizeof_VkSurfaceCapabilitiesKHR * uint(len(surfaceCapabilities))))
		slice3 := (*[1 << 31]C.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(c.pSurfaceCapabilities))[:len(surfaceCapabilities):len(surfaceCapabilities)]
		for i3, _ := range surfaceCapabilities {
			surfaceCapabilities[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(c.physicalDevice, c.surface, c.pSurfaceCapabilities)
	_ret = Result(c._ret)
	return
}

type ColorSpaceKHR int

const (
	COLOR_SPACE_SRGB_NONLINEAR_KHR          ColorSpaceKHR = 0
	COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT    ColorSpaceKHR = 1000104001
	COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT    ColorSpaceKHR = 1000104002
	COLOR_SPACE_DCI_P3_LINEAR_EXT           ColorSpaceKHR = 1000104003
	COLOR_SPACE_DCI_P3_NONLINEAR_EXT        ColorSpaceKHR = 1000104004
	COLOR_SPACE_BT709_LINEAR_EXT            ColorSpaceKHR = 1000104005
	COLOR_SPACE_BT709_NONLINEAR_EXT         ColorSpaceKHR = 1000104006
	COLOR_SPACE_BT2020_LINEAR_EXT           ColorSpaceKHR = 1000104007
	COLOR_SPACE_HDR10_ST2084_EXT            ColorSpaceKHR = 1000104008
	COLOR_SPACE_DOLBYVISION_EXT             ColorSpaceKHR = 1000104009
	COLOR_SPACE_HDR10_HLG_EXT               ColorSpaceKHR = 1000104010
	COLOR_SPACE_ADOBERGB_LINEAR_EXT         ColorSpaceKHR = 1000104011
	COLOR_SPACE_ADOBERGB_NONLINEAR_EXT      ColorSpaceKHR = 1000104012
	COLOR_SPACE_PASS_THROUGH_EXT            ColorSpaceKHR = 1000104013
	COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT ColorSpaceKHR = 1000104014
	COLOR_SPACE_BEGIN_RANGE_KHR             ColorSpaceKHR = COLOR_SPACE_SRGB_NONLINEAR_KHR
	COLOR_SPACE_END_RANGE_KHR               ColorSpaceKHR = COLOR_SPACE_SRGB_NONLINEAR_KHR
	COLOR_SPACE_RANGE_SIZE_KHR              ColorSpaceKHR = (COLOR_SPACE_SRGB_NONLINEAR_KHR - COLOR_SPACE_SRGB_NONLINEAR_KHR + 1)
	COLOR_SPACE_MAX_ENUM_KHR                ColorSpaceKHR = 2147483647
)

type SurfaceFormatKHR struct {
	Format     Format
	ColorSpace ColorSpaceKHR
}

func (g *SurfaceFormatKHR) toC(c *C.VkSurfaceFormatKHR) {
	c.format = C.VkFormat(g.Format)
	c.colorSpace = C.VkColorSpaceKHR(g.ColorSpace)
}
func (g *SurfaceFormatKHR) fromC(c *C.VkSurfaceFormatKHR) {
	g.Format = Format(c.format)
	g.ColorSpace = ColorSpaceKHR(c.colorSpace)
}
func GetPhysicalDeviceSurfaceFormatsKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, surfaceFormatCount *uint32, surfaceFormats []SurfaceFormatKHR) (_ret Result) {
	var c struct {
		physicalDevice      C.VkPhysicalDevice
		surface             C.VkSurfaceKHR
		pSurfaceFormatCount *C.uint32_t
		pSurfaceFormats     *C.VkSurfaceFormatKHR
		_ret                C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pSurfaceFormatCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSurfaceFormatCount = C.uint32_t(*surfaceFormatCount)
	}
	{
		c.pSurfaceFormats = (*C.VkSurfaceFormatKHR)(_sa.alloc(C.sizeof_VkSurfaceFormatKHR * uint(len(surfaceFormats))))
		slice3 := (*[1 << 31]C.VkSurfaceFormatKHR)(unsafe.Pointer(c.pSurfaceFormats))[:len(surfaceFormats):len(surfaceFormats)]
		for i3, _ := range surfaceFormats {
			surfaceFormats[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfaceFormatsKHR(c.physicalDevice, c.surface, c.pSurfaceFormatCount, c.pSurfaceFormats)
	_ret = Result(c._ret)
	*surfaceFormatCount = uint32(*c.pSurfaceFormatCount)
	return
}

type PresentModeKHR int

const (
	PRESENT_MODE_IMMEDIATE_KHR                 PresentModeKHR = 0
	PRESENT_MODE_MAILBOX_KHR                   PresentModeKHR = 1
	PRESENT_MODE_FIFO_KHR                      PresentModeKHR = 2
	PRESENT_MODE_FIFO_RELAXED_KHR              PresentModeKHR = 3
	PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR     PresentModeKHR = 1000111000
	PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR PresentModeKHR = 1000111001
	PRESENT_MODE_BEGIN_RANGE_KHR               PresentModeKHR = PRESENT_MODE_IMMEDIATE_KHR
	PRESENT_MODE_END_RANGE_KHR                 PresentModeKHR = PRESENT_MODE_FIFO_RELAXED_KHR
	PRESENT_MODE_RANGE_SIZE_KHR                PresentModeKHR = (PRESENT_MODE_FIFO_RELAXED_KHR - PRESENT_MODE_IMMEDIATE_KHR + 1)
	PRESENT_MODE_MAX_ENUM_KHR                  PresentModeKHR = 2147483647
)

func GetPhysicalDeviceSurfacePresentModesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, presentModeCount *uint32, presentModes []PresentModeKHR) (_ret Result) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		surface           C.VkSurfaceKHR
		pPresentModeCount *C.uint32_t
		pPresentModes     *C.VkPresentModeKHR
		_ret              C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pPresentModeCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPresentModeCount = C.uint32_t(*presentModeCount)
	}
	{
		c.pPresentModes = (*C.VkPresentModeKHR)(_sa.alloc(C.sizeof_VkPresentModeKHR * uint(len(presentModes))))
		slice3 := (*[1 << 31]C.VkPresentModeKHR)(unsafe.Pointer(c.pPresentModes))[:len(presentModes):len(presentModes)]
		for i3, _ := range presentModes {
			slice3[i3] = C.VkPresentModeKHR(presentModes[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfacePresentModesKHR(c.physicalDevice, c.surface, c.pPresentModeCount, c.pPresentModes)
	_ret = Result(c._ret)
	*presentModeCount = uint32(*c.pPresentModeCount)
	return
}

type SwapchainCreateFlagsKHR Flags
type CompositeAlphaFlagBitsKHR int

const (
	COMPOSITE_ALPHA_OPAQUE_BIT_KHR          CompositeAlphaFlagBitsKHR = 1
	COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR  CompositeAlphaFlagBitsKHR = 2
	COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR CompositeAlphaFlagBitsKHR = 4
	COMPOSITE_ALPHA_INHERIT_BIT_KHR         CompositeAlphaFlagBitsKHR = 8
	COMPOSITE_ALPHA_FLAG_BITS_MAX_ENUM_KHR  CompositeAlphaFlagBitsKHR = 2147483647
)

type SwapchainKHR C.VkSwapchainKHR
type SwapchainCreateInfoKHR struct {
	Type               StructureType
	Next               unsafe.Pointer
	Flags              SwapchainCreateFlagsKHR
	Surface            SurfaceKHR
	MinImageCount      uint32
	ImageFormat        Format
	ImageColorSpace    ColorSpaceKHR
	ImageExtent        Extent2D
	ImageArrayLayers   uint32
	ImageUsage         ImageUsageFlags
	ImageSharingMode   SharingMode
	QueueFamilyIndices []uint32
	PreTransform       SurfaceTransformFlagBitsKHR
	CompositeAlpha     CompositeAlphaFlagBitsKHR
	PresentMode        PresentModeKHR
	Clipped            bool
	OldSwapchain       SwapchainKHR
}

func (g *SwapchainCreateInfoKHR) toC(c *C.VkSwapchainCreateInfoKHR, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkSwapchainCreateFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSwapchainCreateFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSwapchainCreateFlagsKHR(temp_in_VkSwapchainCreateFlagsKHR)
	}
	c.surface = C.VkSurfaceKHR(g.Surface)
	c.minImageCount = C.uint32_t(g.MinImageCount)
	c.imageFormat = C.VkFormat(g.ImageFormat)
	c.imageColorSpace = C.VkColorSpaceKHR(g.ImageColorSpace)
	g.ImageExtent.toC(&c.imageExtent)
	c.imageArrayLayers = C.uint32_t(g.ImageArrayLayers)
	{
		var temp_in_VkImageUsageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.ImageUsage)))
			temp_in_VkImageUsageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.imageUsage = C.VkImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	c.imageSharingMode = C.VkSharingMode(g.ImageSharingMode)
	c.queueFamilyIndexCount = C.uint32_t(len(g.QueueFamilyIndices))
	{
		c.pQueueFamilyIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.QueueFamilyIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			slice2[i2] = C.uint32_t(g.QueueFamilyIndices[i2])
		}
	}
	c.preTransform = C.VkSurfaceTransformFlagBitsKHR(g.PreTransform)
	c.compositeAlpha = C.VkCompositeAlphaFlagBitsKHR(g.CompositeAlpha)
	c.presentMode = C.VkPresentModeKHR(g.PresentMode)
	if g.Clipped {
		c.clipped = 1
	} else {
		c.clipped = 0
	}
	c.oldSwapchain = C.VkSwapchainKHR(g.OldSwapchain)
}
func (g *SwapchainCreateInfoKHR) fromC(c *C.VkSwapchainCreateInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkSwapchainCreateFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSwapchainCreateFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.Flags = SwapchainCreateFlagsKHR(temp_in_VkSwapchainCreateFlagsKHR)
	}
	g.Surface = SurfaceKHR(c.surface)
	g.MinImageCount = uint32(c.minImageCount)
	g.ImageFormat = Format(c.imageFormat)
	g.ImageColorSpace = ColorSpaceKHR(c.imageColorSpace)
	g.ImageExtent.fromC(&c.imageExtent)
	g.ImageArrayLayers = uint32(c.imageArrayLayers)
	{
		var temp_in_VkImageUsageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.imageUsage)))
			temp_in_VkImageUsageFlags = Flags(temp_in_VkFlags)
		}
		g.ImageUsage = ImageUsageFlags(temp_in_VkImageUsageFlags)
	}
	g.ImageSharingMode = SharingMode(c.imageSharingMode)
	g.QueueFamilyIndices = make([]uint32, int(c.queueFamilyIndexCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.QueueFamilyIndices):len(g.QueueFamilyIndices)]
		for i2, _ := range g.QueueFamilyIndices {
			g.QueueFamilyIndices[i2] = uint32(slice2[i2])
		}
	}
	g.PreTransform = SurfaceTransformFlagBitsKHR(c.preTransform)
	g.CompositeAlpha = CompositeAlphaFlagBitsKHR(c.compositeAlpha)
	g.PresentMode = PresentModeKHR(c.presentMode)
	g.Clipped = c.clipped != 0
	g.OldSwapchain = SwapchainKHR(c.oldSwapchain)
}
func CreateSwapchainKHR(device Device, createInfo *SwapchainCreateInfoKHR, allocator *AllocationCallbacks, swapchain *SwapchainKHR) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkSwapchainCreateInfoKHR
		pAllocator  *C.VkAllocationCallbacks
		pSwapchain  *C.VkSwapchainKHR
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkSwapchainCreateInfoKHR)(_sa.alloc(C.sizeof_VkSwapchainCreateInfoKHR))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSwapchain = (*C.VkSwapchainKHR)(_sa.alloc(C.sizeof_VkSwapchainKHR))
		*c.pSwapchain = C.VkSwapchainKHR(*swapchain)
	}
	c._ret = C.vkCreateSwapchainKHR(c.device, c.pCreateInfo, c.pAllocator, c.pSwapchain)
	_ret = Result(c._ret)
	*swapchain = SwapchainKHR(*c.pSwapchain)
	return
}
func DestroySwapchainKHR(device Device, swapchain SwapchainKHR, allocator *AllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		swapchain  C.VkSwapchainKHR
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.swapchain = C.VkSwapchainKHR(swapchain)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySwapchainKHR(c.device, c.swapchain, c.pAllocator)
}
func GetSwapchainImagesKHR(device Device, swapchain SwapchainKHR, swapchainImageCount *uint32, swapchainImages []Image) (_ret Result) {
	var c struct {
		device               C.VkDevice
		swapchain            C.VkSwapchainKHR
		pSwapchainImageCount *C.uint32_t
		pSwapchainImages     *C.VkImage
		_ret                 C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.swapchain = C.VkSwapchainKHR(swapchain)
	{
		c.pSwapchainImageCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSwapchainImageCount = C.uint32_t(*swapchainImageCount)
	}
	{
		c.pSwapchainImages = (*C.VkImage)(_sa.alloc(C.sizeof_VkImage * uint(len(swapchainImages))))
		slice3 := (*[1 << 31]C.VkImage)(unsafe.Pointer(c.pSwapchainImages))[:len(swapchainImages):len(swapchainImages)]
		for i3, _ := range swapchainImages {
			slice3[i3] = C.VkImage(swapchainImages[i3])
		}
	}
	c._ret = C.vkGetSwapchainImagesKHR(c.device, c.swapchain, c.pSwapchainImageCount, c.pSwapchainImages)
	_ret = Result(c._ret)
	*swapchainImageCount = uint32(*c.pSwapchainImageCount)
	return
}
func AcquireNextImageKHR(device Device, swapchain SwapchainKHR, timeout uint64, semaphore Semaphore, fence Fence, imageIndex *uint32) (_ret Result) {
	var c struct {
		device      C.VkDevice
		swapchain   C.VkSwapchainKHR
		timeout     C.uint64_t
		semaphore   C.VkSemaphore
		fence       C.VkFence
		pImageIndex *C.uint32_t
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.swapchain = C.VkSwapchainKHR(swapchain)
	c.timeout = C.uint64_t(timeout)
	c.semaphore = C.VkSemaphore(semaphore)
	c.fence = C.VkFence(fence)
	{
		c.pImageIndex = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pImageIndex = C.uint32_t(*imageIndex)
	}
	c._ret = C.vkAcquireNextImageKHR(c.device, c.swapchain, c.timeout, c.semaphore, c.fence, c.pImageIndex)
	_ret = Result(c._ret)
	*imageIndex = uint32(*c.pImageIndex)
	return
}

type PresentInfoKHR struct {
	Type           StructureType
	Next           unsafe.Pointer
	WaitSemaphores []Semaphore
	Swapchains     []SwapchainKHR
	ImageIndices   []uint32
	Results        []Result
}

func (g *PresentInfoKHR) toC(c *C.VkPresentInfoKHR, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.waitSemaphoreCount = C.uint32_t(len(g.WaitSemaphores))
	{
		c.pWaitSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.WaitSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			slice2[i2] = C.VkSemaphore(g.WaitSemaphores[i2])
		}
	}
	c.swapchainCount = C.uint32_t(len(g.Swapchains))
	{
		c.pSwapchains = (*C.VkSwapchainKHR)(_sa.alloc(C.sizeof_VkSwapchainKHR * uint(len(g.Swapchains))))
		slice2 := (*[1 << 31]C.VkSwapchainKHR)(unsafe.Pointer(c.pSwapchains))[:len(g.Swapchains):len(g.Swapchains)]
		for i2, _ := range g.Swapchains {
			slice2[i2] = C.VkSwapchainKHR(g.Swapchains[i2])
		}
	}
	{
		c.pImageIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.ImageIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pImageIndices))[:len(g.ImageIndices):len(g.ImageIndices)]
		for i2, _ := range g.ImageIndices {
			slice2[i2] = C.uint32_t(g.ImageIndices[i2])
		}
	}
	{
		c.pResults = (*C.VkResult)(_sa.alloc(C.sizeof_VkResult * uint(len(g.Results))))
		slice2 := (*[1 << 31]C.VkResult)(unsafe.Pointer(c.pResults))[:len(g.Results):len(g.Results)]
		for i2, _ := range g.Results {
			slice2[i2] = C.VkResult(g.Results[i2])
		}
	}
}
func (g *PresentInfoKHR) fromC(c *C.VkPresentInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.WaitSemaphores = make([]Semaphore, int(c.waitSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.WaitSemaphores):len(g.WaitSemaphores)]
		for i2, _ := range g.WaitSemaphores {
			g.WaitSemaphores[i2] = Semaphore(slice2[i2])
		}
	}
	g.Swapchains = make([]SwapchainKHR, int(c.swapchainCount))
	{
		slice2 := (*[1 << 31]C.VkSwapchainKHR)(unsafe.Pointer(c.pSwapchains))[:len(g.Swapchains):len(g.Swapchains)]
		for i2, _ := range g.Swapchains {
			g.Swapchains[i2] = SwapchainKHR(slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pImageIndices))[:len(g.ImageIndices):len(g.ImageIndices)]
		for i2, _ := range g.ImageIndices {
			g.ImageIndices[i2] = uint32(slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkResult)(unsafe.Pointer(c.pResults))[:len(g.Results):len(g.Results)]
		for i2, _ := range g.Results {
			g.Results[i2] = Result(slice2[i2])
		}
	}
}
func QueuePresentKHR(queue Queue, presentInfo *PresentInfoKHR) (_ret Result) {
	var c struct {
		queue        C.VkQueue
		pPresentInfo *C.VkPresentInfoKHR
		_ret         C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.queue = C.VkQueue(queue)
	{
		c.pPresentInfo = (*C.VkPresentInfoKHR)(_sa.alloc(C.sizeof_VkPresentInfoKHR))
		presentInfo.toC(c.pPresentInfo, _sa)
	}
	c._ret = C.vkQueuePresentKHR(c.queue, c.pPresentInfo)
	_ret = Result(c._ret)
	return
}

type DeviceGroupPresentModeFlagsKHR Flags
type DeviceGroupPresentCapabilitiesKHR struct {
	Type        StructureType
	Next        unsafe.Pointer
	PresentMask [32]uint32
	Modes       DeviceGroupPresentModeFlagsKHR
}

func (g *DeviceGroupPresentCapabilitiesKHR) toC(c *C.VkDeviceGroupPresentCapabilitiesKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	for i, _ := range g.PresentMask {
		c.presentMask[i] = C.uint32_t(g.PresentMask[i])
	}
	{
		var temp_in_VkDeviceGroupPresentModeFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Modes)))
			temp_in_VkDeviceGroupPresentModeFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.modes = C.VkDeviceGroupPresentModeFlagsKHR(temp_in_VkDeviceGroupPresentModeFlagsKHR)
	}
}
func (g *DeviceGroupPresentCapabilitiesKHR) fromC(c *C.VkDeviceGroupPresentCapabilitiesKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	for i, _ := range g.PresentMask {
		g.PresentMask[i] = uint32(c.presentMask[i])
	}
	{
		var temp_in_VkDeviceGroupPresentModeFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.modes)))
			temp_in_VkDeviceGroupPresentModeFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.Modes = DeviceGroupPresentModeFlagsKHR(temp_in_VkDeviceGroupPresentModeFlagsKHR)
	}
}
func GetDeviceGroupPresentCapabilitiesKHR(device Device, deviceGroupPresentCapabilities []DeviceGroupPresentCapabilitiesKHR) (_ret Result) {
	var c struct {
		device                          C.VkDevice
		pDeviceGroupPresentCapabilities *C.VkDeviceGroupPresentCapabilitiesKHR
		_ret                            C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pDeviceGroupPresentCapabilities = (*C.VkDeviceGroupPresentCapabilitiesKHR)(_sa.alloc(C.sizeof_VkDeviceGroupPresentCapabilitiesKHR * uint(len(deviceGroupPresentCapabilities))))
		slice3 := (*[1 << 31]C.VkDeviceGroupPresentCapabilitiesKHR)(unsafe.Pointer(c.pDeviceGroupPresentCapabilities))[:len(deviceGroupPresentCapabilities):len(deviceGroupPresentCapabilities)]
		for i3, _ := range deviceGroupPresentCapabilities {
			deviceGroupPresentCapabilities[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetDeviceGroupPresentCapabilitiesKHR(c.device, c.pDeviceGroupPresentCapabilities)
	_ret = Result(c._ret)
	return
}
func GetDeviceGroupSurfacePresentModesKHR(device Device, surface SurfaceKHR, modes []DeviceGroupPresentModeFlagsKHR) (_ret Result) {
	var c struct {
		device  C.VkDevice
		surface C.VkSurfaceKHR
		pModes  *C.VkDeviceGroupPresentModeFlagsKHR
		_ret    C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pModes = (*C.VkDeviceGroupPresentModeFlagsKHR)(_sa.alloc(C.sizeof_VkDeviceGroupPresentModeFlagsKHR * uint(len(modes))))
		slice3 := (*[1 << 31]C.VkDeviceGroupPresentModeFlagsKHR)(unsafe.Pointer(c.pModes))[:len(modes):len(modes)]
		for i3, _ := range modes {
			{
				var temp_in_VkDeviceGroupPresentModeFlagsKHR C.VkFlags
				{
					var temp_in_VkFlags C.uint32_t
					temp_in_VkFlags = C.uint32_t((uint32)((Flags)(modes[i3])))
					temp_in_VkDeviceGroupPresentModeFlagsKHR = C.VkFlags(temp_in_VkFlags)
				}
				slice3[i3] = C.VkDeviceGroupPresentModeFlagsKHR(temp_in_VkDeviceGroupPresentModeFlagsKHR)
			}
		}
	}
	c._ret = C.vkGetDeviceGroupSurfacePresentModesKHR(c.device, c.surface, c.pModes)
	_ret = Result(c._ret)
	return
}
func GetPhysicalDevicePresentRectanglesKHR(physicalDevice PhysicalDevice, surface SurfaceKHR, rectCount *uint32, rects []Rect2D) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		surface        C.VkSurfaceKHR
		pRectCount     *C.uint32_t
		pRects         *C.VkRect2D
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.surface = C.VkSurfaceKHR(surface)
	{
		c.pRectCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pRectCount = C.uint32_t(*rectCount)
	}
	{
		c.pRects = (*C.VkRect2D)(_sa.alloc(C.sizeof_VkRect2D * uint(len(rects))))
		slice3 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pRects))[:len(rects):len(rects)]
		for i3, _ := range rects {
			rects[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDevicePresentRectanglesKHR(c.physicalDevice, c.surface, c.pRectCount, c.pRects)
	_ret = Result(c._ret)
	*rectCount = uint32(*c.pRectCount)
	return
}

type AcquireNextImageInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Swapchain  SwapchainKHR
	Timeout    uint64
	Semaphore  Semaphore
	Fence      Fence
	DeviceMask uint32
}

func (g *AcquireNextImageInfoKHR) toC(c *C.VkAcquireNextImageInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.swapchain = C.VkSwapchainKHR(g.Swapchain)
	c.timeout = C.uint64_t(g.Timeout)
	c.semaphore = C.VkSemaphore(g.Semaphore)
	c.fence = C.VkFence(g.Fence)
	c.deviceMask = C.uint32_t(g.DeviceMask)
}
func (g *AcquireNextImageInfoKHR) fromC(c *C.VkAcquireNextImageInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Swapchain = SwapchainKHR(c.swapchain)
	g.Timeout = uint64(c.timeout)
	g.Semaphore = Semaphore(c.semaphore)
	g.Fence = Fence(c.fence)
	g.DeviceMask = uint32(c.deviceMask)
}
func AcquireNextImage2KHR(device Device, acquireInfo *AcquireNextImageInfoKHR, imageIndex *uint32) (_ret Result) {
	var c struct {
		device       C.VkDevice
		pAcquireInfo *C.VkAcquireNextImageInfoKHR
		pImageIndex  *C.uint32_t
		_ret         C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAcquireInfo = (*C.VkAcquireNextImageInfoKHR)(_sa.alloc(C.sizeof_VkAcquireNextImageInfoKHR))
		acquireInfo.toC(c.pAcquireInfo)
	}
	{
		c.pImageIndex = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pImageIndex = C.uint32_t(*imageIndex)
	}
	c._ret = C.vkAcquireNextImage2KHR(c.device, c.pAcquireInfo, c.pImageIndex)
	_ret = Result(c._ret)
	*imageIndex = uint32(*c.pImageIndex)
	return
}

type DisplayKHR C.VkDisplayKHR
type DisplayPropertiesKHR struct {
	Display              DisplayKHR
	DisplayName          string
	PhysicalDimensions   Extent2D
	PhysicalResolution   Extent2D
	SupportedTransforms  SurfaceTransformFlagsKHR
	PlaneReorderPossible bool
	PersistentContent    bool
}

func (g *DisplayPropertiesKHR) toC(c *C.VkDisplayPropertiesKHR, _sa *stackAllocator) {
	c.display = C.VkDisplayKHR(g.Display)
	c.displayName = toCString(g.DisplayName, _sa)
	g.PhysicalDimensions.toC(&c.physicalDimensions)
	g.PhysicalResolution.toC(&c.physicalResolution)
	{
		var temp_in_VkSurfaceTransformFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SupportedTransforms)))
			temp_in_VkSurfaceTransformFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.supportedTransforms = C.VkSurfaceTransformFlagsKHR(temp_in_VkSurfaceTransformFlagsKHR)
	}
	if g.PlaneReorderPossible {
		c.planeReorderPossible = 1
	} else {
		c.planeReorderPossible = 0
	}
	if g.PersistentContent {
		c.persistentContent = 1
	} else {
		c.persistentContent = 0
	}
}
func (g *DisplayPropertiesKHR) fromC(c *C.VkDisplayPropertiesKHR) {
	g.Display = DisplayKHR(c.display)
	g.DisplayName = toGoString(c.displayName)
	g.PhysicalDimensions.fromC(&c.physicalDimensions)
	g.PhysicalResolution.fromC(&c.physicalResolution)
	{
		var temp_in_VkSurfaceTransformFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.supportedTransforms)))
			temp_in_VkSurfaceTransformFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.SupportedTransforms = SurfaceTransformFlagsKHR(temp_in_VkSurfaceTransformFlagsKHR)
	}
	g.PlaneReorderPossible = c.planeReorderPossible != 0
	g.PersistentContent = c.persistentContent != 0
}
func GetPhysicalDeviceDisplayPropertiesKHR(physicalDevice PhysicalDevice, propertyCount *uint32, properties []DisplayPropertiesKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayPropertiesKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayPropertiesKHR)(_sa.alloc(C.sizeof_VkDisplayPropertiesKHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayPropertiesKHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3], _sa)
		}
	}
	c._ret = C.vkGetPhysicalDeviceDisplayPropertiesKHR(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type DisplayPlanePropertiesKHR struct {
	CurrentDisplay    DisplayKHR
	CurrentStackIndex uint32
}

func (g *DisplayPlanePropertiesKHR) toC(c *C.VkDisplayPlanePropertiesKHR) {
	c.currentDisplay = C.VkDisplayKHR(g.CurrentDisplay)
	c.currentStackIndex = C.uint32_t(g.CurrentStackIndex)
}
func (g *DisplayPlanePropertiesKHR) fromC(c *C.VkDisplayPlanePropertiesKHR) {
	g.CurrentDisplay = DisplayKHR(c.currentDisplay)
	g.CurrentStackIndex = uint32(c.currentStackIndex)
}
func GetPhysicalDeviceDisplayPlanePropertiesKHR(physicalDevice PhysicalDevice, propertyCount *uint32, properties []DisplayPlanePropertiesKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayPlanePropertiesKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayPlanePropertiesKHR)(_sa.alloc(C.sizeof_VkDisplayPlanePropertiesKHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayPlanePropertiesKHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceDisplayPlanePropertiesKHR(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}
func GetDisplayPlaneSupportedDisplaysKHR(physicalDevice PhysicalDevice, planeIndex uint32, displayCount *uint32, displays []DisplayKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		planeIndex     C.uint32_t
		pDisplayCount  *C.uint32_t
		pDisplays      *C.VkDisplayKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.planeIndex = C.uint32_t(planeIndex)
	{
		c.pDisplayCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pDisplayCount = C.uint32_t(*displayCount)
	}
	{
		c.pDisplays = (*C.VkDisplayKHR)(_sa.alloc(C.sizeof_VkDisplayKHR * uint(len(displays))))
		slice3 := (*[1 << 31]C.VkDisplayKHR)(unsafe.Pointer(c.pDisplays))[:len(displays):len(displays)]
		for i3, _ := range displays {
			slice3[i3] = C.VkDisplayKHR(displays[i3])
		}
	}
	c._ret = C.vkGetDisplayPlaneSupportedDisplaysKHR(c.physicalDevice, c.planeIndex, c.pDisplayCount, c.pDisplays)
	_ret = Result(c._ret)
	*displayCount = uint32(*c.pDisplayCount)
	return
}

type DisplayModeKHR C.VkDisplayModeKHR
type DisplayModeParametersKHR struct {
	VisibleRegion Extent2D
	RefreshRate   uint32
}

func (g *DisplayModeParametersKHR) toC(c *C.VkDisplayModeParametersKHR) {
	g.VisibleRegion.toC(&c.visibleRegion)
	c.refreshRate = C.uint32_t(g.RefreshRate)
}
func (g *DisplayModeParametersKHR) fromC(c *C.VkDisplayModeParametersKHR) {
	g.VisibleRegion.fromC(&c.visibleRegion)
	g.RefreshRate = uint32(c.refreshRate)
}

type DisplayModePropertiesKHR struct {
	DisplayMode DisplayModeKHR
	Parameters  DisplayModeParametersKHR
}

func (g *DisplayModePropertiesKHR) toC(c *C.VkDisplayModePropertiesKHR) {
	c.displayMode = C.VkDisplayModeKHR(g.DisplayMode)
	g.Parameters.toC(&c.parameters)
}
func (g *DisplayModePropertiesKHR) fromC(c *C.VkDisplayModePropertiesKHR) {
	g.DisplayMode = DisplayModeKHR(c.displayMode)
	g.Parameters.fromC(&c.parameters)
}
func GetDisplayModePropertiesKHR(physicalDevice PhysicalDevice, display DisplayKHR, propertyCount *uint32, properties []DisplayModePropertiesKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		display        C.VkDisplayKHR
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayModePropertiesKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.display = C.VkDisplayKHR(display)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayModePropertiesKHR)(_sa.alloc(C.sizeof_VkDisplayModePropertiesKHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayModePropertiesKHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetDisplayModePropertiesKHR(c.physicalDevice, c.display, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type DisplayModeCreateFlagsKHR Flags
type DisplayModeCreateInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Flags      DisplayModeCreateFlagsKHR
	Parameters DisplayModeParametersKHR
}

func (g *DisplayModeCreateInfoKHR) toC(c *C.VkDisplayModeCreateInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDisplayModeCreateFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDisplayModeCreateFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDisplayModeCreateFlagsKHR(temp_in_VkDisplayModeCreateFlagsKHR)
	}
	g.Parameters.toC(&c.parameters)
}
func (g *DisplayModeCreateInfoKHR) fromC(c *C.VkDisplayModeCreateInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDisplayModeCreateFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDisplayModeCreateFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.Flags = DisplayModeCreateFlagsKHR(temp_in_VkDisplayModeCreateFlagsKHR)
	}
	g.Parameters.fromC(&c.parameters)
}
func CreateDisplayModeKHR(physicalDevice PhysicalDevice, display DisplayKHR, createInfo *DisplayModeCreateInfoKHR, allocator *AllocationCallbacks, mode *DisplayModeKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		display        C.VkDisplayKHR
		pCreateInfo    *C.VkDisplayModeCreateInfoKHR
		pAllocator     *C.VkAllocationCallbacks
		pMode          *C.VkDisplayModeKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.display = C.VkDisplayKHR(display)
	{
		c.pCreateInfo = (*C.VkDisplayModeCreateInfoKHR)(_sa.alloc(C.sizeof_VkDisplayModeCreateInfoKHR))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pMode = (*C.VkDisplayModeKHR)(_sa.alloc(C.sizeof_VkDisplayModeKHR))
		*c.pMode = C.VkDisplayModeKHR(*mode)
	}
	c._ret = C.vkCreateDisplayModeKHR(c.physicalDevice, c.display, c.pCreateInfo, c.pAllocator, c.pMode)
	_ret = Result(c._ret)
	*mode = DisplayModeKHR(*c.pMode)
	return
}

type DisplayPlaneAlphaFlagsKHR Flags
type DisplayPlaneCapabilitiesKHR struct {
	SupportedAlpha DisplayPlaneAlphaFlagsKHR
	MinSrcPosition Offset2D
	MaxSrcPosition Offset2D
	MinSrcExtent   Extent2D
	MaxSrcExtent   Extent2D
	MinDstPosition Offset2D
	MaxDstPosition Offset2D
	MinDstExtent   Extent2D
	MaxDstExtent   Extent2D
}

func (g *DisplayPlaneCapabilitiesKHR) toC(c *C.VkDisplayPlaneCapabilitiesKHR) {
	{
		var temp_in_VkDisplayPlaneAlphaFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SupportedAlpha)))
			temp_in_VkDisplayPlaneAlphaFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.supportedAlpha = C.VkDisplayPlaneAlphaFlagsKHR(temp_in_VkDisplayPlaneAlphaFlagsKHR)
	}
	g.MinSrcPosition.toC(&c.minSrcPosition)
	g.MaxSrcPosition.toC(&c.maxSrcPosition)
	g.MinSrcExtent.toC(&c.minSrcExtent)
	g.MaxSrcExtent.toC(&c.maxSrcExtent)
	g.MinDstPosition.toC(&c.minDstPosition)
	g.MaxDstPosition.toC(&c.maxDstPosition)
	g.MinDstExtent.toC(&c.minDstExtent)
	g.MaxDstExtent.toC(&c.maxDstExtent)
}
func (g *DisplayPlaneCapabilitiesKHR) fromC(c *C.VkDisplayPlaneCapabilitiesKHR) {
	{
		var temp_in_VkDisplayPlaneAlphaFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.supportedAlpha)))
			temp_in_VkDisplayPlaneAlphaFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.SupportedAlpha = DisplayPlaneAlphaFlagsKHR(temp_in_VkDisplayPlaneAlphaFlagsKHR)
	}
	g.MinSrcPosition.fromC(&c.minSrcPosition)
	g.MaxSrcPosition.fromC(&c.maxSrcPosition)
	g.MinSrcExtent.fromC(&c.minSrcExtent)
	g.MaxSrcExtent.fromC(&c.maxSrcExtent)
	g.MinDstPosition.fromC(&c.minDstPosition)
	g.MaxDstPosition.fromC(&c.maxDstPosition)
	g.MinDstExtent.fromC(&c.minDstExtent)
	g.MaxDstExtent.fromC(&c.maxDstExtent)
}
func GetDisplayPlaneCapabilitiesKHR(physicalDevice PhysicalDevice, mode DisplayModeKHR, capabilities []DisplayPlaneCapabilitiesKHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		mode           C.VkDisplayModeKHR
		planeIndex     C.uint32_t
		pCapabilities  *C.VkDisplayPlaneCapabilitiesKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.mode = C.VkDisplayModeKHR(mode)
	c.planeIndex = C.uint32_t(len(capabilities))
	{
		c.pCapabilities = (*C.VkDisplayPlaneCapabilitiesKHR)(_sa.alloc(C.sizeof_VkDisplayPlaneCapabilitiesKHR * uint(len(capabilities))))
		slice3 := (*[1 << 31]C.VkDisplayPlaneCapabilitiesKHR)(unsafe.Pointer(c.pCapabilities))[:len(capabilities):len(capabilities)]
		for i3, _ := range capabilities {
			capabilities[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetDisplayPlaneCapabilitiesKHR(c.physicalDevice, c.mode, c.planeIndex, c.pCapabilities)
	_ret = Result(c._ret)
	return
}

type DisplaySurfaceCreateFlagsKHR Flags
type DisplayPlaneAlphaFlagBitsKHR int

const (
	DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR                  DisplayPlaneAlphaFlagBitsKHR = 1
	DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR                  DisplayPlaneAlphaFlagBitsKHR = 2
	DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR               DisplayPlaneAlphaFlagBitsKHR = 4
	DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR DisplayPlaneAlphaFlagBitsKHR = 8
	DISPLAY_PLANE_ALPHA_FLAG_BITS_MAX_ENUM_KHR          DisplayPlaneAlphaFlagBitsKHR = 2147483647
)

type DisplaySurfaceCreateInfoKHR struct {
	Type            StructureType
	Next            unsafe.Pointer
	Flags           DisplaySurfaceCreateFlagsKHR
	DisplayMode     DisplayModeKHR
	PlaneIndex      uint32
	PlaneStackIndex uint32
	Transform       SurfaceTransformFlagBitsKHR
	GlobalAlpha     float32
	AlphaMode       DisplayPlaneAlphaFlagBitsKHR
	ImageExtent     Extent2D
}

func (g *DisplaySurfaceCreateInfoKHR) toC(c *C.VkDisplaySurfaceCreateInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkDisplaySurfaceCreateFlagsKHR C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkDisplaySurfaceCreateFlagsKHR = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkDisplaySurfaceCreateFlagsKHR(temp_in_VkDisplaySurfaceCreateFlagsKHR)
	}
	c.displayMode = C.VkDisplayModeKHR(g.DisplayMode)
	c.planeIndex = C.uint32_t(g.PlaneIndex)
	c.planeStackIndex = C.uint32_t(g.PlaneStackIndex)
	c.transform = C.VkSurfaceTransformFlagBitsKHR(g.Transform)
	c.globalAlpha = C.float(g.GlobalAlpha)
	c.alphaMode = C.VkDisplayPlaneAlphaFlagBitsKHR(g.AlphaMode)
	g.ImageExtent.toC(&c.imageExtent)
}
func (g *DisplaySurfaceCreateInfoKHR) fromC(c *C.VkDisplaySurfaceCreateInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkDisplaySurfaceCreateFlagsKHR Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkDisplaySurfaceCreateFlagsKHR = Flags(temp_in_VkFlags)
		}
		g.Flags = DisplaySurfaceCreateFlagsKHR(temp_in_VkDisplaySurfaceCreateFlagsKHR)
	}
	g.DisplayMode = DisplayModeKHR(c.displayMode)
	g.PlaneIndex = uint32(c.planeIndex)
	g.PlaneStackIndex = uint32(c.planeStackIndex)
	g.Transform = SurfaceTransformFlagBitsKHR(c.transform)
	g.GlobalAlpha = float32(c.globalAlpha)
	g.AlphaMode = DisplayPlaneAlphaFlagBitsKHR(c.alphaMode)
	g.ImageExtent.fromC(&c.imageExtent)
}
func CreateDisplayPlaneSurfaceKHR(instance Instance, createInfo *DisplaySurfaceCreateInfoKHR, allocator *AllocationCallbacks, surface *SurfaceKHR) (_ret Result) {
	var c struct {
		instance    C.VkInstance
		pCreateInfo *C.VkDisplaySurfaceCreateInfoKHR
		pAllocator  *C.VkAllocationCallbacks
		pSurface    *C.VkSurfaceKHR
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pCreateInfo = (*C.VkDisplaySurfaceCreateInfoKHR)(_sa.alloc(C.sizeof_VkDisplaySurfaceCreateInfoKHR))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSurface = (*C.VkSurfaceKHR)(_sa.alloc(C.sizeof_VkSurfaceKHR))
		*c.pSurface = C.VkSurfaceKHR(*surface)
	}
	c._ret = C.vkCreateDisplayPlaneSurfaceKHR(c.instance, c.pCreateInfo, c.pAllocator, c.pSurface)
	_ret = Result(c._ret)
	*surface = SurfaceKHR(*c.pSurface)
	return
}
func CreateSharedSwapchainsKHR(device Device, createInfos []SwapchainCreateInfoKHR, allocator *AllocationCallbacks, swapchains []SwapchainKHR) (_ret Result) {
	var c struct {
		device         C.VkDevice
		swapchainCount C.uint32_t
		pCreateInfos   *C.VkSwapchainCreateInfoKHR
		pAllocator     *C.VkAllocationCallbacks
		pSwapchains    *C.VkSwapchainKHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.swapchainCount = C.uint32_t(len(createInfos))
	{
		c.pCreateInfos = (*C.VkSwapchainCreateInfoKHR)(_sa.alloc(C.sizeof_VkSwapchainCreateInfoKHR * uint(len(createInfos))))
		slice3 := (*[1 << 31]C.VkSwapchainCreateInfoKHR)(unsafe.Pointer(c.pCreateInfos))[:len(createInfos):len(createInfos)]
		for i3, _ := range createInfos {
			createInfos[i3].toC(&slice3[i3], _sa)
		}
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSwapchains = (*C.VkSwapchainKHR)(_sa.alloc(C.sizeof_VkSwapchainKHR * uint(len(swapchains))))
		slice3 := (*[1 << 31]C.VkSwapchainKHR)(unsafe.Pointer(c.pSwapchains))[:len(swapchains):len(swapchains)]
		for i3, _ := range swapchains {
			slice3[i3] = C.VkSwapchainKHR(swapchains[i3])
		}
	}
	c._ret = C.vkCreateSharedSwapchainsKHR(c.device, c.swapchainCount, c.pCreateInfos, c.pAllocator, c.pSwapchains)
	_ret = Result(c._ret)
	return
}
func GetPhysicalDeviceFeatures2KHR(physicalDevice PhysicalDevice, features []PhysicalDeviceFeatures2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFeatures      *C.VkPhysicalDeviceFeatures2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFeatures = (*C.VkPhysicalDeviceFeatures2)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures2 * uint(len(features))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceFeatures2)(unsafe.Pointer(c.pFeatures))[:len(features):len(features)]
		for i3, _ := range features {
			features[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFeatures2KHR(c.physicalDevice, c.pFeatures)
}
func GetPhysicalDeviceProperties2KHR(physicalDevice PhysicalDevice, properties []PhysicalDeviceProperties2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pProperties    *C.VkPhysicalDeviceProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pProperties = (*C.VkPhysicalDeviceProperties2)(_sa.alloc(C.sizeof_VkPhysicalDeviceProperties2 * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceProperties2)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceProperties2KHR(c.physicalDevice, c.pProperties)
}
func GetPhysicalDeviceFormatProperties2KHR(physicalDevice PhysicalDevice, format Format, formatProperties []FormatProperties2) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		format            C.VkFormat
		pFormatProperties *C.VkFormatProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.format = C.VkFormat(format)
	{
		c.pFormatProperties = (*C.VkFormatProperties2)(_sa.alloc(C.sizeof_VkFormatProperties2 * uint(len(formatProperties))))
		slice3 := (*[1 << 31]C.VkFormatProperties2)(unsafe.Pointer(c.pFormatProperties))[:len(formatProperties):len(formatProperties)]
		for i3, _ := range formatProperties {
			formatProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFormatProperties2KHR(c.physicalDevice, c.format, c.pFormatProperties)
}
func GetPhysicalDeviceImageFormatProperties2KHR(physicalDevice PhysicalDevice, imageFormatInfo *PhysicalDeviceImageFormatInfo2, imageFormatProperties []ImageFormatProperties2) (_ret Result) {
	var c struct {
		physicalDevice         C.VkPhysicalDevice
		pImageFormatInfo       *C.VkPhysicalDeviceImageFormatInfo2
		pImageFormatProperties *C.VkImageFormatProperties2
		_ret                   C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pImageFormatInfo = (*C.VkPhysicalDeviceImageFormatInfo2)(_sa.alloc(C.sizeof_VkPhysicalDeviceImageFormatInfo2))
		imageFormatInfo.toC(c.pImageFormatInfo)
	}
	{
		c.pImageFormatProperties = (*C.VkImageFormatProperties2)(_sa.alloc(C.sizeof_VkImageFormatProperties2 * uint(len(imageFormatProperties))))
		slice3 := (*[1 << 31]C.VkImageFormatProperties2)(unsafe.Pointer(c.pImageFormatProperties))[:len(imageFormatProperties):len(imageFormatProperties)]
		for i3, _ := range imageFormatProperties {
			imageFormatProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceImageFormatProperties2KHR(c.physicalDevice, c.pImageFormatInfo, c.pImageFormatProperties)
	_ret = Result(c._ret)
	return
}
func GetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice PhysicalDevice, queueFamilyPropertyCount *uint32, queueFamilyProperties []QueueFamilyProperties2) {
	var c struct {
		physicalDevice            C.VkPhysicalDevice
		pQueueFamilyPropertyCount *C.uint32_t
		pQueueFamilyProperties    *C.VkQueueFamilyProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pQueueFamilyPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pQueueFamilyPropertyCount = C.uint32_t(*queueFamilyPropertyCount)
	}
	{
		c.pQueueFamilyProperties = (*C.VkQueueFamilyProperties2)(_sa.alloc(C.sizeof_VkQueueFamilyProperties2 * uint(len(queueFamilyProperties))))
		slice3 := (*[1 << 31]C.VkQueueFamilyProperties2)(unsafe.Pointer(c.pQueueFamilyProperties))[:len(queueFamilyProperties):len(queueFamilyProperties)]
		for i3, _ := range queueFamilyProperties {
			queueFamilyProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceQueueFamilyProperties2KHR(c.physicalDevice, c.pQueueFamilyPropertyCount, c.pQueueFamilyProperties)
	*queueFamilyPropertyCount = uint32(*c.pQueueFamilyPropertyCount)
}
func GetPhysicalDeviceMemoryProperties2KHR(physicalDevice PhysicalDevice, memoryProperties []PhysicalDeviceMemoryProperties2) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		pMemoryProperties *C.VkPhysicalDeviceMemoryProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pMemoryProperties = (*C.VkPhysicalDeviceMemoryProperties2)(_sa.alloc(C.sizeof_VkPhysicalDeviceMemoryProperties2 * uint(len(memoryProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceMemoryProperties2)(unsafe.Pointer(c.pMemoryProperties))[:len(memoryProperties):len(memoryProperties)]
		for i3, _ := range memoryProperties {
			memoryProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceMemoryProperties2KHR(c.physicalDevice, c.pMemoryProperties)
}
func GetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice PhysicalDevice, formatInfo *PhysicalDeviceSparseImageFormatInfo2, propertyCount *uint32, properties []SparseImageFormatProperties2) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFormatInfo    *C.VkPhysicalDeviceSparseImageFormatInfo2
		pPropertyCount *C.uint32_t
		pProperties    *C.VkSparseImageFormatProperties2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFormatInfo = (*C.VkPhysicalDeviceSparseImageFormatInfo2)(_sa.alloc(C.sizeof_VkPhysicalDeviceSparseImageFormatInfo2))
		formatInfo.toC(c.pFormatInfo)
	}
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkSparseImageFormatProperties2)(_sa.alloc(C.sizeof_VkSparseImageFormatProperties2 * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkSparseImageFormatProperties2)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceSparseImageFormatProperties2KHR(c.physicalDevice, c.pFormatInfo, c.pPropertyCount, c.pProperties)
	*propertyCount = uint32(*c.pPropertyCount)
}
func GetDeviceGroupPeerMemoryFeaturesKHR(device Device, heapIndex uint32, localDeviceIndex uint32, peerMemoryFeatures []PeerMemoryFeatureFlags) {
	var c struct {
		device              C.VkDevice
		heapIndex           C.uint32_t
		localDeviceIndex    C.uint32_t
		remoteDeviceIndex   C.uint32_t
		pPeerMemoryFeatures *C.VkPeerMemoryFeatureFlags
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.heapIndex = C.uint32_t(heapIndex)
	c.localDeviceIndex = C.uint32_t(localDeviceIndex)
	c.remoteDeviceIndex = C.uint32_t(len(peerMemoryFeatures))
	{
		c.pPeerMemoryFeatures = (*C.VkPeerMemoryFeatureFlags)(_sa.alloc(C.sizeof_VkPeerMemoryFeatureFlags * uint(len(peerMemoryFeatures))))
		slice3 := (*[1 << 31]C.VkPeerMemoryFeatureFlags)(unsafe.Pointer(c.pPeerMemoryFeatures))[:len(peerMemoryFeatures):len(peerMemoryFeatures)]
		for i3, _ := range peerMemoryFeatures {
			{
				var temp_in_VkPeerMemoryFeatureFlags C.VkFlags
				{
					var temp_in_VkFlags C.uint32_t
					temp_in_VkFlags = C.uint32_t((uint32)((Flags)(peerMemoryFeatures[i3])))
					temp_in_VkPeerMemoryFeatureFlags = C.VkFlags(temp_in_VkFlags)
				}
				slice3[i3] = C.VkPeerMemoryFeatureFlags(temp_in_VkPeerMemoryFeatureFlags)
			}
		}
	}
	C.vkGetDeviceGroupPeerMemoryFeaturesKHR(c.device, c.heapIndex, c.localDeviceIndex, c.remoteDeviceIndex, c.pPeerMemoryFeatures)
}
func CmdSetDeviceMaskKHR(commandBuffer CommandBuffer, deviceMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		deviceMask    C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.deviceMask = C.uint32_t(deviceMask)
	C.vkCmdSetDeviceMaskKHR(c.commandBuffer, c.deviceMask)
}
func CmdDispatchBaseKHR(commandBuffer CommandBuffer, baseGroupX uint32, baseGroupY uint32, baseGroupZ uint32, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		baseGroupX    C.uint32_t
		baseGroupY    C.uint32_t
		baseGroupZ    C.uint32_t
		groupCountX   C.uint32_t
		groupCountY   C.uint32_t
		groupCountZ   C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.baseGroupX = C.uint32_t(baseGroupX)
	c.baseGroupY = C.uint32_t(baseGroupY)
	c.baseGroupZ = C.uint32_t(baseGroupZ)
	c.groupCountX = C.uint32_t(groupCountX)
	c.groupCountY = C.uint32_t(groupCountY)
	c.groupCountZ = C.uint32_t(groupCountZ)
	C.vkCmdDispatchBaseKHR(c.commandBuffer, c.baseGroupX, c.baseGroupY, c.baseGroupZ, c.groupCountX, c.groupCountY, c.groupCountZ)
}
func TrimCommandPoolKHR(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	var c struct {
		device      C.VkDevice
		commandPool C.VkCommandPool
		flags       C.VkCommandPoolTrimFlags
	}
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	{
		var temp_in_VkCommandPoolTrimFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(flags)))
			temp_in_VkCommandPoolTrimFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkCommandPoolTrimFlags(temp_in_VkCommandPoolTrimFlags)
	}
	C.vkTrimCommandPoolKHR(c.device, c.commandPool, c.flags)
}
func EnumeratePhysicalDeviceGroupsKHR(instance Instance, physicalDeviceGroupCount *uint32, physicalDeviceGroupProperties []PhysicalDeviceGroupProperties) (_ret Result) {
	var c struct {
		instance                       C.VkInstance
		pPhysicalDeviceGroupCount      *C.uint32_t
		pPhysicalDeviceGroupProperties *C.VkPhysicalDeviceGroupProperties
		_ret                           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pPhysicalDeviceGroupCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPhysicalDeviceGroupCount = C.uint32_t(*physicalDeviceGroupCount)
	}
	{
		c.pPhysicalDeviceGroupProperties = (*C.VkPhysicalDeviceGroupProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceGroupProperties * uint(len(physicalDeviceGroupProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceGroupProperties)(unsafe.Pointer(c.pPhysicalDeviceGroupProperties))[:len(physicalDeviceGroupProperties):len(physicalDeviceGroupProperties)]
		for i3, _ := range physicalDeviceGroupProperties {
			physicalDeviceGroupProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumeratePhysicalDeviceGroupsKHR(c.instance, c.pPhysicalDeviceGroupCount, c.pPhysicalDeviceGroupProperties)
	_ret = Result(c._ret)
	*physicalDeviceGroupCount = uint32(*c.pPhysicalDeviceGroupCount)
	return
}
func GetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice PhysicalDevice, externalBufferInfo *PhysicalDeviceExternalBufferInfo, externalBufferProperties []ExternalBufferProperties) {
	var c struct {
		physicalDevice            C.VkPhysicalDevice
		pExternalBufferInfo       *C.VkPhysicalDeviceExternalBufferInfo
		pExternalBufferProperties *C.VkExternalBufferProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalBufferInfo = (*C.VkPhysicalDeviceExternalBufferInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalBufferInfo))
		externalBufferInfo.toC(c.pExternalBufferInfo)
	}
	{
		c.pExternalBufferProperties = (*C.VkExternalBufferProperties)(_sa.alloc(C.sizeof_VkExternalBufferProperties * uint(len(externalBufferProperties))))
		slice3 := (*[1 << 31]C.VkExternalBufferProperties)(unsafe.Pointer(c.pExternalBufferProperties))[:len(externalBufferProperties):len(externalBufferProperties)]
		for i3, _ := range externalBufferProperties {
			externalBufferProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalBufferPropertiesKHR(c.physicalDevice, c.pExternalBufferInfo, c.pExternalBufferProperties)
}

type MemoryGetFdInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Memory     DeviceMemory
	HandleType ExternalMemoryHandleTypeFlagBits
}

func (g *MemoryGetFdInfoKHR) toC(c *C.VkMemoryGetFdInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.memory = C.VkDeviceMemory(g.Memory)
	c.handleType = C.VkExternalMemoryHandleTypeFlagBits(g.HandleType)
}
func (g *MemoryGetFdInfoKHR) fromC(c *C.VkMemoryGetFdInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Memory = DeviceMemory(c.memory)
	g.HandleType = ExternalMemoryHandleTypeFlagBits(c.handleType)
}
func GetMemoryFdKHR(device Device, getFdInfo *MemoryGetFdInfoKHR, fd *int32) (_ret Result) {
	var c struct {
		device     C.VkDevice
		pGetFdInfo *C.VkMemoryGetFdInfoKHR
		pFd        *C.int
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pGetFdInfo = (*C.VkMemoryGetFdInfoKHR)(_sa.alloc(C.sizeof_VkMemoryGetFdInfoKHR))
		getFdInfo.toC(c.pGetFdInfo)
	}
	{
		c.pFd = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.pFd = C.int(*fd)
	}
	c._ret = C.vkGetMemoryFdKHR(c.device, c.pGetFdInfo, c.pFd)
	_ret = Result(c._ret)
	*fd = int32(*c.pFd)
	return
}

type MemoryFdPropertiesKHR struct {
	Type           StructureType
	Next           unsafe.Pointer
	MemoryTypeBits uint32
}

func (g *MemoryFdPropertiesKHR) toC(c *C.VkMemoryFdPropertiesKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.memoryTypeBits = C.uint32_t(g.MemoryTypeBits)
}
func (g *MemoryFdPropertiesKHR) fromC(c *C.VkMemoryFdPropertiesKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.MemoryTypeBits = uint32(c.memoryTypeBits)
}
func GetMemoryFdPropertiesKHR(device Device, handleType ExternalMemoryHandleTypeFlagBits, fd int32, memoryFdProperties []MemoryFdPropertiesKHR) (_ret Result) {
	var c struct {
		device              C.VkDevice
		handleType          C.VkExternalMemoryHandleTypeFlagBits
		fd                  C.int
		pMemoryFdProperties *C.VkMemoryFdPropertiesKHR
		_ret                C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.handleType = C.VkExternalMemoryHandleTypeFlagBits(handleType)
	c.fd = C.int(fd)
	{
		c.pMemoryFdProperties = (*C.VkMemoryFdPropertiesKHR)(_sa.alloc(C.sizeof_VkMemoryFdPropertiesKHR * uint(len(memoryFdProperties))))
		slice3 := (*[1 << 31]C.VkMemoryFdPropertiesKHR)(unsafe.Pointer(c.pMemoryFdProperties))[:len(memoryFdProperties):len(memoryFdProperties)]
		for i3, _ := range memoryFdProperties {
			memoryFdProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetMemoryFdPropertiesKHR(c.device, c.handleType, c.fd, c.pMemoryFdProperties)
	_ret = Result(c._ret)
	return
}
func GetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice PhysicalDevice, externalSemaphoreInfo *PhysicalDeviceExternalSemaphoreInfo, externalSemaphoreProperties []ExternalSemaphoreProperties) {
	var c struct {
		physicalDevice               C.VkPhysicalDevice
		pExternalSemaphoreInfo       *C.VkPhysicalDeviceExternalSemaphoreInfo
		pExternalSemaphoreProperties *C.VkExternalSemaphoreProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalSemaphoreInfo = (*C.VkPhysicalDeviceExternalSemaphoreInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalSemaphoreInfo))
		externalSemaphoreInfo.toC(c.pExternalSemaphoreInfo)
	}
	{
		c.pExternalSemaphoreProperties = (*C.VkExternalSemaphoreProperties)(_sa.alloc(C.sizeof_VkExternalSemaphoreProperties * uint(len(externalSemaphoreProperties))))
		slice3 := (*[1 << 31]C.VkExternalSemaphoreProperties)(unsafe.Pointer(c.pExternalSemaphoreProperties))[:len(externalSemaphoreProperties):len(externalSemaphoreProperties)]
		for i3, _ := range externalSemaphoreProperties {
			externalSemaphoreProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalSemaphorePropertiesKHR(c.physicalDevice, c.pExternalSemaphoreInfo, c.pExternalSemaphoreProperties)
}

type SemaphoreImportFlags Flags
type ImportSemaphoreFdInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Semaphore  Semaphore
	Flags      SemaphoreImportFlags
	HandleType ExternalSemaphoreHandleTypeFlagBits
	Fd         int32
}

func (g *ImportSemaphoreFdInfoKHR) toC(c *C.VkImportSemaphoreFdInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.semaphore = C.VkSemaphore(g.Semaphore)
	{
		var temp_in_VkSemaphoreImportFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSemaphoreImportFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSemaphoreImportFlags(temp_in_VkSemaphoreImportFlags)
	}
	c.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(g.HandleType)
	c.fd = C.int(g.Fd)
}
func (g *ImportSemaphoreFdInfoKHR) fromC(c *C.VkImportSemaphoreFdInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Semaphore = Semaphore(c.semaphore)
	{
		var temp_in_VkSemaphoreImportFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSemaphoreImportFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SemaphoreImportFlags(temp_in_VkSemaphoreImportFlags)
	}
	g.HandleType = ExternalSemaphoreHandleTypeFlagBits(c.handleType)
	g.Fd = int32(c.fd)
}
func ImportSemaphoreFdKHR(device Device, importSemaphoreFdInfo *ImportSemaphoreFdInfoKHR) (_ret Result) {
	var c struct {
		device                 C.VkDevice
		pImportSemaphoreFdInfo *C.VkImportSemaphoreFdInfoKHR
		_ret                   C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pImportSemaphoreFdInfo = (*C.VkImportSemaphoreFdInfoKHR)(_sa.alloc(C.sizeof_VkImportSemaphoreFdInfoKHR))
		importSemaphoreFdInfo.toC(c.pImportSemaphoreFdInfo)
	}
	c._ret = C.vkImportSemaphoreFdKHR(c.device, c.pImportSemaphoreFdInfo)
	_ret = Result(c._ret)
	return
}

type SemaphoreGetFdInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Semaphore  Semaphore
	HandleType ExternalSemaphoreHandleTypeFlagBits
}

func (g *SemaphoreGetFdInfoKHR) toC(c *C.VkSemaphoreGetFdInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.semaphore = C.VkSemaphore(g.Semaphore)
	c.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(g.HandleType)
}
func (g *SemaphoreGetFdInfoKHR) fromC(c *C.VkSemaphoreGetFdInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Semaphore = Semaphore(c.semaphore)
	g.HandleType = ExternalSemaphoreHandleTypeFlagBits(c.handleType)
}
func GetSemaphoreFdKHR(device Device, getFdInfo *SemaphoreGetFdInfoKHR, fd *int32) (_ret Result) {
	var c struct {
		device     C.VkDevice
		pGetFdInfo *C.VkSemaphoreGetFdInfoKHR
		pFd        *C.int
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pGetFdInfo = (*C.VkSemaphoreGetFdInfoKHR)(_sa.alloc(C.sizeof_VkSemaphoreGetFdInfoKHR))
		getFdInfo.toC(c.pGetFdInfo)
	}
	{
		c.pFd = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.pFd = C.int(*fd)
	}
	c._ret = C.vkGetSemaphoreFdKHR(c.device, c.pGetFdInfo, c.pFd)
	_ret = Result(c._ret)
	*fd = int32(*c.pFd)
	return
}
func CmdPushDescriptorSetKHR(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, set uint32, descriptorWrites []WriteDescriptorSet) {
	var c struct {
		commandBuffer        C.VkCommandBuffer
		pipelineBindPoint    C.VkPipelineBindPoint
		layout               C.VkPipelineLayout
		set                  C.uint32_t
		descriptorWriteCount C.uint32_t
		pDescriptorWrites    *C.VkWriteDescriptorSet
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.pipelineBindPoint = C.VkPipelineBindPoint(pipelineBindPoint)
	c.layout = C.VkPipelineLayout(layout)
	c.set = C.uint32_t(set)
	c.descriptorWriteCount = C.uint32_t(len(descriptorWrites))
	{
		c.pDescriptorWrites = (*C.VkWriteDescriptorSet)(_sa.alloc(C.sizeof_VkWriteDescriptorSet * uint(len(descriptorWrites))))
		slice3 := (*[1 << 31]C.VkWriteDescriptorSet)(unsafe.Pointer(c.pDescriptorWrites))[:len(descriptorWrites):len(descriptorWrites)]
		for i3, _ := range descriptorWrites {
			descriptorWrites[i3].toC(&slice3[i3], _sa)
		}
	}
	C.vkCmdPushDescriptorSetKHR(c.commandBuffer, c.pipelineBindPoint, c.layout, c.set, c.descriptorWriteCount, c.pDescriptorWrites)
}
func CmdPushDescriptorSetWithTemplateKHR(commandBuffer CommandBuffer, descriptorUpdateTemplate DescriptorUpdateTemplate, layout PipelineLayout, data []byte) {
	var c struct {
		commandBuffer            C.VkCommandBuffer
		descriptorUpdateTemplate C.VkDescriptorUpdateTemplate
		layout                   C.VkPipelineLayout
		set                      C.uint32_t
		pData                    unsafe.Pointer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.descriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate)
	c.layout = C.VkPipelineLayout(layout)
	c.set = C.uint32_t(len(data))
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	C.vkCmdPushDescriptorSetWithTemplateKHR(c.commandBuffer, c.descriptorUpdateTemplate, c.layout, c.set, c.pData)
}
func CreateDescriptorUpdateTemplateKHR(device Device, createInfo *DescriptorUpdateTemplateCreateInfo, allocator *AllocationCallbacks, descriptorUpdateTemplate *DescriptorUpdateTemplate) (_ret Result) {
	var c struct {
		device                    C.VkDevice
		pCreateInfo               *C.VkDescriptorUpdateTemplateCreateInfo
		pAllocator                *C.VkAllocationCallbacks
		pDescriptorUpdateTemplate *C.VkDescriptorUpdateTemplate
		_ret                      C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorUpdateTemplateCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorUpdateTemplateCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDescriptorUpdateTemplate = (*C.VkDescriptorUpdateTemplate)(_sa.alloc(C.sizeof_VkDescriptorUpdateTemplate))
		*c.pDescriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(*descriptorUpdateTemplate)
	}
	c._ret = C.vkCreateDescriptorUpdateTemplateKHR(c.device, c.pCreateInfo, c.pAllocator, c.pDescriptorUpdateTemplate)
	_ret = Result(c._ret)
	*descriptorUpdateTemplate = DescriptorUpdateTemplate(*c.pDescriptorUpdateTemplate)
	return
}
func DestroyDescriptorUpdateTemplateKHR(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, allocator *AllocationCallbacks) {
	var c struct {
		device                   C.VkDevice
		descriptorUpdateTemplate C.VkDescriptorUpdateTemplate
		pAllocator               *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorUpdateTemplateKHR(c.device, c.descriptorUpdateTemplate, c.pAllocator)
}
func UpdateDescriptorSetWithTemplateKHR(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, data []byte) {
	var c struct {
		device                   C.VkDevice
		descriptorSet            C.VkDescriptorSet
		descriptorUpdateTemplate C.VkDescriptorUpdateTemplate
		pData                    unsafe.Pointer
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.descriptorSet = C.VkDescriptorSet(descriptorSet)
	c.descriptorUpdateTemplate = C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate)
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(data)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(data):len(data)]
		for i3, _ := range data {
			slice3[i3] = data[i3]
		}
	}
	C.vkUpdateDescriptorSetWithTemplateKHR(c.device, c.descriptorSet, c.descriptorUpdateTemplate, c.pData)
}

type AttachmentDescription2KHR struct {
	Type           StructureType
	Next           unsafe.Pointer
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlagBits
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func (g *AttachmentDescription2KHR) toC(c *C.VkAttachmentDescription2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkAttachmentDescriptionFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkAttachmentDescriptionFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkAttachmentDescriptionFlags(temp_in_VkAttachmentDescriptionFlags)
	}
	c.format = C.VkFormat(g.Format)
	c.samples = C.VkSampleCountFlagBits(g.Samples)
	c.loadOp = C.VkAttachmentLoadOp(g.LoadOp)
	c.storeOp = C.VkAttachmentStoreOp(g.StoreOp)
	c.stencilLoadOp = C.VkAttachmentLoadOp(g.StencilLoadOp)
	c.stencilStoreOp = C.VkAttachmentStoreOp(g.StencilStoreOp)
	c.initialLayout = C.VkImageLayout(g.InitialLayout)
	c.finalLayout = C.VkImageLayout(g.FinalLayout)
}
func (g *AttachmentDescription2KHR) fromC(c *C.VkAttachmentDescription2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkAttachmentDescriptionFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkAttachmentDescriptionFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = AttachmentDescriptionFlags(temp_in_VkAttachmentDescriptionFlags)
	}
	g.Format = Format(c.format)
	g.Samples = SampleCountFlagBits(c.samples)
	g.LoadOp = AttachmentLoadOp(c.loadOp)
	g.StoreOp = AttachmentStoreOp(c.storeOp)
	g.StencilLoadOp = AttachmentLoadOp(c.stencilLoadOp)
	g.StencilStoreOp = AttachmentStoreOp(c.stencilStoreOp)
	g.InitialLayout = ImageLayout(c.initialLayout)
	g.FinalLayout = ImageLayout(c.finalLayout)
}

type AttachmentReference2KHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Attachment uint32
	Layout     ImageLayout
	AspectMask ImageAspectFlags
}

func (g *AttachmentReference2KHR) toC(c *C.VkAttachmentReference2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.attachment = C.uint32_t(g.Attachment)
	c.layout = C.VkImageLayout(g.Layout)
	{
		var temp_in_VkImageAspectFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.AspectMask)))
			temp_in_VkImageAspectFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.aspectMask = C.VkImageAspectFlags(temp_in_VkImageAspectFlags)
	}
}
func (g *AttachmentReference2KHR) fromC(c *C.VkAttachmentReference2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Attachment = uint32(c.attachment)
	g.Layout = ImageLayout(c.layout)
	{
		var temp_in_VkImageAspectFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			temp_in_VkImageAspectFlags = Flags(temp_in_VkFlags)
		}
		g.AspectMask = ImageAspectFlags(temp_in_VkImageAspectFlags)
	}
}

type SubpassDescription2KHR struct {
	Type                   StructureType
	Next                   unsafe.Pointer
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	ViewMask               uint32
	InputAttachments       []AttachmentReference2KHR
	ColorAttachments       []AttachmentReference2KHR
	ResolveAttachments     []AttachmentReference2KHR
	DepthStencilAttachment *AttachmentReference2KHR
	PreserveAttachments    []uint32
}

func (g *SubpassDescription2KHR) toC(c *C.VkSubpassDescription2KHR, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkSubpassDescriptionFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkSubpassDescriptionFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkSubpassDescriptionFlags(temp_in_VkSubpassDescriptionFlags)
	}
	c.pipelineBindPoint = C.VkPipelineBindPoint(g.PipelineBindPoint)
	c.viewMask = C.uint32_t(g.ViewMask)
	c.inputAttachmentCount = C.uint32_t(len(g.InputAttachments))
	{
		c.pInputAttachments = (*C.VkAttachmentReference2KHR)(_sa.alloc(C.sizeof_VkAttachmentReference2KHR * uint(len(g.InputAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pInputAttachments))[:len(g.InputAttachments):len(g.InputAttachments)]
		for i2, _ := range g.InputAttachments {
			g.InputAttachments[i2].toC(&slice2[i2])
		}
	}
	c.colorAttachmentCount = C.uint32_t(len(g.ColorAttachments))
	{
		c.pColorAttachments = (*C.VkAttachmentReference2KHR)(_sa.alloc(C.sizeof_VkAttachmentReference2KHR * uint(len(g.ColorAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pColorAttachments))[:len(g.ColorAttachments):len(g.ColorAttachments)]
		for i2, _ := range g.ColorAttachments {
			g.ColorAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pResolveAttachments = (*C.VkAttachmentReference2KHR)(_sa.alloc(C.sizeof_VkAttachmentReference2KHR * uint(len(g.ResolveAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pResolveAttachments))[:len(g.ResolveAttachments):len(g.ResolveAttachments)]
		for i2, _ := range g.ResolveAttachments {
			g.ResolveAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pDepthStencilAttachment = (*C.VkAttachmentReference2KHR)(_sa.alloc(C.sizeof_VkAttachmentReference2KHR))
		g.DepthStencilAttachment.toC(c.pDepthStencilAttachment)
	}
	c.preserveAttachmentCount = C.uint32_t(len(g.PreserveAttachments))
	{
		c.pPreserveAttachments = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.PreserveAttachments))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.PreserveAttachments):len(g.PreserveAttachments)]
		for i2, _ := range g.PreserveAttachments {
			slice2[i2] = C.uint32_t(g.PreserveAttachments[i2])
		}
	}
}
func (g *SubpassDescription2KHR) fromC(c *C.VkSubpassDescription2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkSubpassDescriptionFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkSubpassDescriptionFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = SubpassDescriptionFlags(temp_in_VkSubpassDescriptionFlags)
	}
	g.PipelineBindPoint = PipelineBindPoint(c.pipelineBindPoint)
	g.ViewMask = uint32(c.viewMask)
	g.InputAttachments = make([]AttachmentReference2KHR, int(c.inputAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pInputAttachments))[:len(g.InputAttachments):len(g.InputAttachments)]
		for i2, _ := range g.InputAttachments {
			g.InputAttachments[i2].fromC(&slice2[i2])
		}
	}
	g.ColorAttachments = make([]AttachmentReference2KHR, int(c.colorAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pColorAttachments))[:len(g.ColorAttachments):len(g.ColorAttachments)]
		for i2, _ := range g.ColorAttachments {
			g.ColorAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference2KHR)(unsafe.Pointer(c.pResolveAttachments))[:len(g.ResolveAttachments):len(g.ResolveAttachments)]
		for i2, _ := range g.ResolveAttachments {
			g.ResolveAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		if g.DepthStencilAttachment == nil {
			g.DepthStencilAttachment = new(AttachmentReference2KHR)
		}
		g.DepthStencilAttachment.fromC(c.pDepthStencilAttachment)
	}
	g.PreserveAttachments = make([]uint32, int(c.preserveAttachmentCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.PreserveAttachments):len(g.PreserveAttachments)]
		for i2, _ := range g.PreserveAttachments {
			g.PreserveAttachments[i2] = uint32(slice2[i2])
		}
	}
}

type SubpassDependency2KHR struct {
	Type            StructureType
	Next            unsafe.Pointer
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
	ViewOffset      int32
}

func (g *SubpassDependency2KHR) toC(c *C.VkSubpassDependency2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.srcSubpass = C.uint32_t(g.SrcSubpass)
	c.dstSubpass = C.uint32_t(g.DstSubpass)
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstStageMask)))
			temp_in_VkPipelineStageFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstStageMask = C.VkPipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.SrcAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.srcAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DstAccessMask)))
			temp_in_VkAccessFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dstAccessMask = C.VkAccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkDependencyFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.DependencyFlags)))
			temp_in_VkDependencyFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.dependencyFlags = C.VkDependencyFlags(temp_in_VkDependencyFlags)
	}
	c.viewOffset = C.int32_t(g.ViewOffset)
}
func (g *SubpassDependency2KHR) fromC(c *C.VkSubpassDependency2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.SrcSubpass = uint32(c.srcSubpass)
	g.DstSubpass = uint32(c.dstSubpass)
	{
		var temp_in_VkPipelineStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcStageMask)))
			temp_in_VkPipelineStageFlags = Flags(temp_in_VkFlags)
		}
		g.SrcStageMask = PipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkPipelineStageFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstStageMask)))
			temp_in_VkPipelineStageFlags = Flags(temp_in_VkFlags)
		}
		g.DstStageMask = PipelineStageFlags(temp_in_VkPipelineStageFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.SrcAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkAccessFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			temp_in_VkAccessFlags = Flags(temp_in_VkFlags)
		}
		g.DstAccessMask = AccessFlags(temp_in_VkAccessFlags)
	}
	{
		var temp_in_VkDependencyFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.dependencyFlags)))
			temp_in_VkDependencyFlags = Flags(temp_in_VkFlags)
		}
		g.DependencyFlags = DependencyFlags(temp_in_VkDependencyFlags)
	}
	g.ViewOffset = int32(c.viewOffset)
}

type RenderPassCreateInfo2KHR struct {
	Type                StructureType
	Next                unsafe.Pointer
	Flags               RenderPassCreateFlags
	Attachments         []AttachmentDescription2KHR
	Subpasses           []SubpassDescription2KHR
	Dependencies        []SubpassDependency2KHR
	CorrelatedViewMasks []uint32
}

func (g *RenderPassCreateInfo2KHR) toC(c *C.VkRenderPassCreateInfo2KHR, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	{
		var temp_in_VkRenderPassCreateFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkRenderPassCreateFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkRenderPassCreateFlags(temp_in_VkRenderPassCreateFlags)
	}
	c.attachmentCount = C.uint32_t(len(g.Attachments))
	{
		c.pAttachments = (*C.VkAttachmentDescription2KHR)(_sa.alloc(C.sizeof_VkAttachmentDescription2KHR * uint(len(g.Attachments))))
		slice2 := (*[1 << 31]C.VkAttachmentDescription2KHR)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].toC(&slice2[i2])
		}
	}
	c.subpassCount = C.uint32_t(len(g.Subpasses))
	{
		c.pSubpasses = (*C.VkSubpassDescription2KHR)(_sa.alloc(C.sizeof_VkSubpassDescription2KHR * uint(len(g.Subpasses))))
		slice2 := (*[1 << 31]C.VkSubpassDescription2KHR)(unsafe.Pointer(c.pSubpasses))[:len(g.Subpasses):len(g.Subpasses)]
		for i2, _ := range g.Subpasses {
			g.Subpasses[i2].toC(&slice2[i2], _sa)
		}
	}
	c.dependencyCount = C.uint32_t(len(g.Dependencies))
	{
		c.pDependencies = (*C.VkSubpassDependency2KHR)(_sa.alloc(C.sizeof_VkSubpassDependency2KHR * uint(len(g.Dependencies))))
		slice2 := (*[1 << 31]C.VkSubpassDependency2KHR)(unsafe.Pointer(c.pDependencies))[:len(g.Dependencies):len(g.Dependencies)]
		for i2, _ := range g.Dependencies {
			g.Dependencies[i2].toC(&slice2[i2])
		}
	}
	c.correlatedViewMaskCount = C.uint32_t(len(g.CorrelatedViewMasks))
	{
		c.pCorrelatedViewMasks = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.CorrelatedViewMasks))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pCorrelatedViewMasks))[:len(g.CorrelatedViewMasks):len(g.CorrelatedViewMasks)]
		for i2, _ := range g.CorrelatedViewMasks {
			slice2[i2] = C.uint32_t(g.CorrelatedViewMasks[i2])
		}
	}
}
func (g *RenderPassCreateInfo2KHR) fromC(c *C.VkRenderPassCreateInfo2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	{
		var temp_in_VkRenderPassCreateFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkRenderPassCreateFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = RenderPassCreateFlags(temp_in_VkRenderPassCreateFlags)
	}
	g.Attachments = make([]AttachmentDescription2KHR, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentDescription2KHR)(unsafe.Pointer(c.pAttachments))[:len(g.Attachments):len(g.Attachments)]
		for i2, _ := range g.Attachments {
			g.Attachments[i2].fromC(&slice2[i2])
		}
	}
	g.Subpasses = make([]SubpassDescription2KHR, int(c.subpassCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDescription2KHR)(unsafe.Pointer(c.pSubpasses))[:len(g.Subpasses):len(g.Subpasses)]
		for i2, _ := range g.Subpasses {
			g.Subpasses[i2].fromC(&slice2[i2])
		}
	}
	g.Dependencies = make([]SubpassDependency2KHR, int(c.dependencyCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDependency2KHR)(unsafe.Pointer(c.pDependencies))[:len(g.Dependencies):len(g.Dependencies)]
		for i2, _ := range g.Dependencies {
			g.Dependencies[i2].fromC(&slice2[i2])
		}
	}
	g.CorrelatedViewMasks = make([]uint32, int(c.correlatedViewMaskCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pCorrelatedViewMasks))[:len(g.CorrelatedViewMasks):len(g.CorrelatedViewMasks)]
		for i2, _ := range g.CorrelatedViewMasks {
			g.CorrelatedViewMasks[i2] = uint32(slice2[i2])
		}
	}
}
func CreateRenderPass2KHR(device Device, createInfo *RenderPassCreateInfo2KHR, allocator *AllocationCallbacks, renderPass *RenderPass) (_ret Result) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkRenderPassCreateInfo2KHR
		pAllocator  *C.VkAllocationCallbacks
		pRenderPass *C.VkRenderPass
		_ret        C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkRenderPassCreateInfo2KHR)(_sa.alloc(C.sizeof_VkRenderPassCreateInfo2KHR))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pRenderPass = (*C.VkRenderPass)(_sa.alloc(C.sizeof_VkRenderPass))
		*c.pRenderPass = C.VkRenderPass(*renderPass)
	}
	c._ret = C.vkCreateRenderPass2KHR(c.device, c.pCreateInfo, c.pAllocator, c.pRenderPass)
	_ret = Result(c._ret)
	*renderPass = RenderPass(*c.pRenderPass)
	return
}

type SubpassBeginInfoKHR struct {
	Type     StructureType
	Next     unsafe.Pointer
	Contents SubpassContents
}

func (g *SubpassBeginInfoKHR) toC(c *C.VkSubpassBeginInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.contents = C.VkSubpassContents(g.Contents)
}
func (g *SubpassBeginInfoKHR) fromC(c *C.VkSubpassBeginInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Contents = SubpassContents(c.contents)
}
func CmdBeginRenderPass2KHR(commandBuffer CommandBuffer, renderPassBegin *RenderPassBeginInfo, subpassBeginInfo *SubpassBeginInfoKHR) {
	var c struct {
		commandBuffer     C.VkCommandBuffer
		pRenderPassBegin  *C.VkRenderPassBeginInfo
		pSubpassBeginInfo *C.VkSubpassBeginInfoKHR
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.pRenderPassBegin = (*C.VkRenderPassBeginInfo)(_sa.alloc(C.sizeof_VkRenderPassBeginInfo))
		renderPassBegin.toC(c.pRenderPassBegin, _sa)
	}
	{
		c.pSubpassBeginInfo = (*C.VkSubpassBeginInfoKHR)(_sa.alloc(C.sizeof_VkSubpassBeginInfoKHR))
		subpassBeginInfo.toC(c.pSubpassBeginInfo)
	}
	C.vkCmdBeginRenderPass2KHR(c.commandBuffer, c.pRenderPassBegin, c.pSubpassBeginInfo)
}

type SubpassEndInfoKHR struct {
	Type StructureType
	Next unsafe.Pointer
}

func (g *SubpassEndInfoKHR) toC(c *C.VkSubpassEndInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
}
func (g *SubpassEndInfoKHR) fromC(c *C.VkSubpassEndInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
}
func CmdNextSubpass2KHR(commandBuffer CommandBuffer, subpassBeginInfo *SubpassBeginInfoKHR, subpassEndInfo *SubpassEndInfoKHR) {
	var c struct {
		commandBuffer     C.VkCommandBuffer
		pSubpassBeginInfo *C.VkSubpassBeginInfoKHR
		pSubpassEndInfo   *C.VkSubpassEndInfoKHR
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.pSubpassBeginInfo = (*C.VkSubpassBeginInfoKHR)(_sa.alloc(C.sizeof_VkSubpassBeginInfoKHR))
		subpassBeginInfo.toC(c.pSubpassBeginInfo)
	}
	{
		c.pSubpassEndInfo = (*C.VkSubpassEndInfoKHR)(_sa.alloc(C.sizeof_VkSubpassEndInfoKHR))
		subpassEndInfo.toC(c.pSubpassEndInfo)
	}
	C.vkCmdNextSubpass2KHR(c.commandBuffer, c.pSubpassBeginInfo, c.pSubpassEndInfo)
}
func CmdEndRenderPass2KHR(commandBuffer CommandBuffer, subpassEndInfo *SubpassEndInfoKHR) {
	var c struct {
		commandBuffer   C.VkCommandBuffer
		pSubpassEndInfo *C.VkSubpassEndInfoKHR
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		c.pSubpassEndInfo = (*C.VkSubpassEndInfoKHR)(_sa.alloc(C.sizeof_VkSubpassEndInfoKHR))
		subpassEndInfo.toC(c.pSubpassEndInfo)
	}
	C.vkCmdEndRenderPass2KHR(c.commandBuffer, c.pSubpassEndInfo)
}
func GetSwapchainStatusKHR(device Device, swapchain SwapchainKHR) (_ret Result) {
	var c struct {
		device    C.VkDevice
		swapchain C.VkSwapchainKHR
		_ret      C.VkResult
	}
	c.device = C.VkDevice(device)
	c.swapchain = C.VkSwapchainKHR(swapchain)
	c._ret = C.vkGetSwapchainStatusKHR(c.device, c.swapchain)
	_ret = Result(c._ret)
	return
}
func GetPhysicalDeviceExternalFencePropertiesKHR(physicalDevice PhysicalDevice, externalFenceInfo *PhysicalDeviceExternalFenceInfo, externalFenceProperties []ExternalFenceProperties) {
	var c struct {
		physicalDevice           C.VkPhysicalDevice
		pExternalFenceInfo       *C.VkPhysicalDeviceExternalFenceInfo
		pExternalFenceProperties *C.VkExternalFenceProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pExternalFenceInfo = (*C.VkPhysicalDeviceExternalFenceInfo)(_sa.alloc(C.sizeof_VkPhysicalDeviceExternalFenceInfo))
		externalFenceInfo.toC(c.pExternalFenceInfo)
	}
	{
		c.pExternalFenceProperties = (*C.VkExternalFenceProperties)(_sa.alloc(C.sizeof_VkExternalFenceProperties * uint(len(externalFenceProperties))))
		slice3 := (*[1 << 31]C.VkExternalFenceProperties)(unsafe.Pointer(c.pExternalFenceProperties))[:len(externalFenceProperties):len(externalFenceProperties)]
		for i3, _ := range externalFenceProperties {
			externalFenceProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceExternalFencePropertiesKHR(c.physicalDevice, c.pExternalFenceInfo, c.pExternalFenceProperties)
}

type FenceImportFlags Flags
type ImportFenceFdInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Fence      Fence
	Flags      FenceImportFlags
	HandleType ExternalFenceHandleTypeFlagBits
	Fd         int32
}

func (g *ImportFenceFdInfoKHR) toC(c *C.VkImportFenceFdInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.fence = C.VkFence(g.Fence)
	{
		var temp_in_VkFenceImportFlags C.VkFlags
		{
			var temp_in_VkFlags C.uint32_t
			temp_in_VkFlags = C.uint32_t((uint32)((Flags)(g.Flags)))
			temp_in_VkFenceImportFlags = C.VkFlags(temp_in_VkFlags)
		}
		c.flags = C.VkFenceImportFlags(temp_in_VkFenceImportFlags)
	}
	c.handleType = C.VkExternalFenceHandleTypeFlagBits(g.HandleType)
	c.fd = C.int(g.Fd)
}
func (g *ImportFenceFdInfoKHR) fromC(c *C.VkImportFenceFdInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Fence = Fence(c.fence)
	{
		var temp_in_VkFenceImportFlags Flags
		{
			var temp_in_VkFlags uint32
			temp_in_VkFlags = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			temp_in_VkFenceImportFlags = Flags(temp_in_VkFlags)
		}
		g.Flags = FenceImportFlags(temp_in_VkFenceImportFlags)
	}
	g.HandleType = ExternalFenceHandleTypeFlagBits(c.handleType)
	g.Fd = int32(c.fd)
}
func ImportFenceFdKHR(device Device, importFenceFdInfo *ImportFenceFdInfoKHR) (_ret Result) {
	var c struct {
		device             C.VkDevice
		pImportFenceFdInfo *C.VkImportFenceFdInfoKHR
		_ret               C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pImportFenceFdInfo = (*C.VkImportFenceFdInfoKHR)(_sa.alloc(C.sizeof_VkImportFenceFdInfoKHR))
		importFenceFdInfo.toC(c.pImportFenceFdInfo)
	}
	c._ret = C.vkImportFenceFdKHR(c.device, c.pImportFenceFdInfo)
	_ret = Result(c._ret)
	return
}

type FenceGetFdInfoKHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Fence      Fence
	HandleType ExternalFenceHandleTypeFlagBits
}

func (g *FenceGetFdInfoKHR) toC(c *C.VkFenceGetFdInfoKHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.fence = C.VkFence(g.Fence)
	c.handleType = C.VkExternalFenceHandleTypeFlagBits(g.HandleType)
}
func (g *FenceGetFdInfoKHR) fromC(c *C.VkFenceGetFdInfoKHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Fence = Fence(c.fence)
	g.HandleType = ExternalFenceHandleTypeFlagBits(c.handleType)
}
func GetFenceFdKHR(device Device, getFdInfo *FenceGetFdInfoKHR, fd *int32) (_ret Result) {
	var c struct {
		device     C.VkDevice
		pGetFdInfo *C.VkFenceGetFdInfoKHR
		pFd        *C.int
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pGetFdInfo = (*C.VkFenceGetFdInfoKHR)(_sa.alloc(C.sizeof_VkFenceGetFdInfoKHR))
		getFdInfo.toC(c.pGetFdInfo)
	}
	{
		c.pFd = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.pFd = C.int(*fd)
	}
	c._ret = C.vkGetFenceFdKHR(c.device, c.pGetFdInfo, c.pFd)
	_ret = Result(c._ret)
	*fd = int32(*c.pFd)
	return
}

type PhysicalDeviceSurfaceInfo2KHR struct {
	Type    StructureType
	Next    unsafe.Pointer
	Surface SurfaceKHR
}

func (g *PhysicalDeviceSurfaceInfo2KHR) toC(c *C.VkPhysicalDeviceSurfaceInfo2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.surface = C.VkSurfaceKHR(g.Surface)
}
func (g *PhysicalDeviceSurfaceInfo2KHR) fromC(c *C.VkPhysicalDeviceSurfaceInfo2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Surface = SurfaceKHR(c.surface)
}

type SurfaceCapabilities2KHR struct {
	Type                StructureType
	Next                unsafe.Pointer
	SurfaceCapabilities SurfaceCapabilitiesKHR
}

func (g *SurfaceCapabilities2KHR) toC(c *C.VkSurfaceCapabilities2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.SurfaceCapabilities.toC(&c.surfaceCapabilities)
}
func (g *SurfaceCapabilities2KHR) fromC(c *C.VkSurfaceCapabilities2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.SurfaceCapabilities.fromC(&c.surfaceCapabilities)
}
func GetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice PhysicalDevice, surfaceInfo *PhysicalDeviceSurfaceInfo2KHR, surfaceCapabilities []SurfaceCapabilities2KHR) (_ret Result) {
	var c struct {
		physicalDevice       C.VkPhysicalDevice
		pSurfaceInfo         *C.VkPhysicalDeviceSurfaceInfo2KHR
		pSurfaceCapabilities *C.VkSurfaceCapabilities2KHR
		_ret                 C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pSurfaceInfo = (*C.VkPhysicalDeviceSurfaceInfo2KHR)(_sa.alloc(C.sizeof_VkPhysicalDeviceSurfaceInfo2KHR))
		surfaceInfo.toC(c.pSurfaceInfo)
	}
	{
		c.pSurfaceCapabilities = (*C.VkSurfaceCapabilities2KHR)(_sa.alloc(C.sizeof_VkSurfaceCapabilities2KHR * uint(len(surfaceCapabilities))))
		slice3 := (*[1 << 31]C.VkSurfaceCapabilities2KHR)(unsafe.Pointer(c.pSurfaceCapabilities))[:len(surfaceCapabilities):len(surfaceCapabilities)]
		for i3, _ := range surfaceCapabilities {
			surfaceCapabilities[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfaceCapabilities2KHR(c.physicalDevice, c.pSurfaceInfo, c.pSurfaceCapabilities)
	_ret = Result(c._ret)
	return
}

type SurfaceFormat2KHR struct {
	Type          StructureType
	Next          unsafe.Pointer
	SurfaceFormat SurfaceFormatKHR
}

func (g *SurfaceFormat2KHR) toC(c *C.VkSurfaceFormat2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.SurfaceFormat.toC(&c.surfaceFormat)
}
func (g *SurfaceFormat2KHR) fromC(c *C.VkSurfaceFormat2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.SurfaceFormat.fromC(&c.surfaceFormat)
}
func GetPhysicalDeviceSurfaceFormats2KHR(physicalDevice PhysicalDevice, surfaceInfo *PhysicalDeviceSurfaceInfo2KHR, surfaceFormatCount *uint32, surfaceFormats []SurfaceFormat2KHR) (_ret Result) {
	var c struct {
		physicalDevice      C.VkPhysicalDevice
		pSurfaceInfo        *C.VkPhysicalDeviceSurfaceInfo2KHR
		pSurfaceFormatCount *C.uint32_t
		pSurfaceFormats     *C.VkSurfaceFormat2KHR
		_ret                C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pSurfaceInfo = (*C.VkPhysicalDeviceSurfaceInfo2KHR)(_sa.alloc(C.sizeof_VkPhysicalDeviceSurfaceInfo2KHR))
		surfaceInfo.toC(c.pSurfaceInfo)
	}
	{
		c.pSurfaceFormatCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSurfaceFormatCount = C.uint32_t(*surfaceFormatCount)
	}
	{
		c.pSurfaceFormats = (*C.VkSurfaceFormat2KHR)(_sa.alloc(C.sizeof_VkSurfaceFormat2KHR * uint(len(surfaceFormats))))
		slice3 := (*[1 << 31]C.VkSurfaceFormat2KHR)(unsafe.Pointer(c.pSurfaceFormats))[:len(surfaceFormats):len(surfaceFormats)]
		for i3, _ := range surfaceFormats {
			surfaceFormats[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceSurfaceFormats2KHR(c.physicalDevice, c.pSurfaceInfo, c.pSurfaceFormatCount, c.pSurfaceFormats)
	_ret = Result(c._ret)
	*surfaceFormatCount = uint32(*c.pSurfaceFormatCount)
	return
}

type DisplayProperties2KHR struct {
	Type              StructureType
	Next              unsafe.Pointer
	DisplayProperties DisplayPropertiesKHR
}

func (g *DisplayProperties2KHR) toC(c *C.VkDisplayProperties2KHR, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.DisplayProperties.toC(&c.displayProperties, _sa)
}
func (g *DisplayProperties2KHR) fromC(c *C.VkDisplayProperties2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.DisplayProperties.fromC(&c.displayProperties)
}
func GetPhysicalDeviceDisplayProperties2KHR(physicalDevice PhysicalDevice, propertyCount *uint32, properties []DisplayProperties2KHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayProperties2KHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayProperties2KHR)(_sa.alloc(C.sizeof_VkDisplayProperties2KHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayProperties2KHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3], _sa)
		}
	}
	c._ret = C.vkGetPhysicalDeviceDisplayProperties2KHR(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type DisplayPlaneProperties2KHR struct {
	Type                   StructureType
	Next                   unsafe.Pointer
	DisplayPlaneProperties DisplayPlanePropertiesKHR
}

func (g *DisplayPlaneProperties2KHR) toC(c *C.VkDisplayPlaneProperties2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.DisplayPlaneProperties.toC(&c.displayPlaneProperties)
}
func (g *DisplayPlaneProperties2KHR) fromC(c *C.VkDisplayPlaneProperties2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.DisplayPlaneProperties.fromC(&c.displayPlaneProperties)
}
func GetPhysicalDeviceDisplayPlaneProperties2KHR(physicalDevice PhysicalDevice, propertyCount *uint32, properties []DisplayPlaneProperties2KHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayPlaneProperties2KHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayPlaneProperties2KHR)(_sa.alloc(C.sizeof_VkDisplayPlaneProperties2KHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayPlaneProperties2KHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceDisplayPlaneProperties2KHR(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type DisplayModeProperties2KHR struct {
	Type                  StructureType
	Next                  unsafe.Pointer
	DisplayModeProperties DisplayModePropertiesKHR
}

func (g *DisplayModeProperties2KHR) toC(c *C.VkDisplayModeProperties2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.DisplayModeProperties.toC(&c.displayModeProperties)
}
func (g *DisplayModeProperties2KHR) fromC(c *C.VkDisplayModeProperties2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.DisplayModeProperties.fromC(&c.displayModeProperties)
}
func GetDisplayModeProperties2KHR(physicalDevice PhysicalDevice, display DisplayKHR, propertyCount *uint32, properties []DisplayModeProperties2KHR) (_ret Result) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		display        C.VkDisplayKHR
		pPropertyCount *C.uint32_t
		pProperties    *C.VkDisplayModeProperties2KHR
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	c.display = C.VkDisplayKHR(display)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*propertyCount)
	}
	{
		c.pProperties = (*C.VkDisplayModeProperties2KHR)(_sa.alloc(C.sizeof_VkDisplayModeProperties2KHR * uint(len(properties))))
		slice3 := (*[1 << 31]C.VkDisplayModeProperties2KHR)(unsafe.Pointer(c.pProperties))[:len(properties):len(properties)]
		for i3, _ := range properties {
			properties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetDisplayModeProperties2KHR(c.physicalDevice, c.display, c.pPropertyCount, c.pProperties)
	_ret = Result(c._ret)
	*propertyCount = uint32(*c.pPropertyCount)
	return
}

type DisplayPlaneInfo2KHR struct {
	Type       StructureType
	Next       unsafe.Pointer
	Mode       DisplayModeKHR
	PlaneIndex uint32
}

func (g *DisplayPlaneInfo2KHR) toC(c *C.VkDisplayPlaneInfo2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	c.mode = C.VkDisplayModeKHR(g.Mode)
	c.planeIndex = C.uint32_t(g.PlaneIndex)
}
func (g *DisplayPlaneInfo2KHR) fromC(c *C.VkDisplayPlaneInfo2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Mode = DisplayModeKHR(c.mode)
	g.PlaneIndex = uint32(c.planeIndex)
}

type DisplayPlaneCapabilities2KHR struct {
	Type         StructureType
	Next         unsafe.Pointer
	Capabilities DisplayPlaneCapabilitiesKHR
}

func (g *DisplayPlaneCapabilities2KHR) toC(c *C.VkDisplayPlaneCapabilities2KHR) {
	c.sType = C.VkStructureType(g.Type)
	c.pNext = g.Next
	g.Capabilities.toC(&c.capabilities)
}
func (g *DisplayPlaneCapabilities2KHR) fromC(c *C.VkDisplayPlaneCapabilities2KHR) {
	g.Type = StructureType(c.sType)
	g.Next = c.pNext
	g.Capabilities.fromC(&c.capabilities)
}
func GetDisplayPlaneCapabilities2KHR(physicalDevice PhysicalDevice, displayPlaneInfo *DisplayPlaneInfo2KHR, capabilities []DisplayPlaneCapabilities2KHR) (_ret Result) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		pDisplayPlaneInfo *C.VkDisplayPlaneInfo2KHR
		pCapabilities     *C.VkDisplayPlaneCapabilities2KHR
		_ret              C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pDisplayPlaneInfo = (*C.VkDisplayPlaneInfo2KHR)(_sa.alloc(C.sizeof_VkDisplayPlaneInfo2KHR))
		displayPlaneInfo.toC(c.pDisplayPlaneInfo)
	}
	{
		c.pCapabilities = (*C.VkDisplayPlaneCapabilities2KHR)(_sa.alloc(C.sizeof_VkDisplayPlaneCapabilities2KHR * uint(len(capabilities))))
		slice3 := (*[1 << 31]C.VkDisplayPlaneCapabilities2KHR)(unsafe.Pointer(c.pCapabilities))[:len(capabilities):len(capabilities)]
		for i3, _ := range capabilities {
			capabilities[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetDisplayPlaneCapabilities2KHR(c.physicalDevice, c.pDisplayPlaneInfo, c.pCapabilities)
	_ret = Result(c._ret)
	return
}
func GetImageMemoryRequirements2KHR(device Device, info *ImageMemoryRequirementsInfo2, memoryRequirements []MemoryRequirements2) {
	var c struct {
		device              C.VkDevice
		pInfo               *C.VkImageMemoryRequirementsInfo2
		pMemoryRequirements *C.VkMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkImageMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkImageMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements2)(_sa.alloc(C.sizeof_VkMemoryRequirements2 * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements2)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageMemoryRequirements2KHR(c.device, c.pInfo, c.pMemoryRequirements)
}
func GetBufferMemoryRequirements2KHR(device Device, info *BufferMemoryRequirementsInfo2, memoryRequirements []MemoryRequirements2) {
	var c struct {
		device              C.VkDevice
		pInfo               *C.VkBufferMemoryRequirementsInfo2
		pMemoryRequirements *C.VkMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkBufferMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkBufferMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pMemoryRequirements = (*C.VkMemoryRequirements2)(_sa.alloc(C.sizeof_VkMemoryRequirements2 * uint(len(memoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements2)(unsafe.Pointer(c.pMemoryRequirements))[:len(memoryRequirements):len(memoryRequirements)]
		for i3, _ := range memoryRequirements {
			memoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetBufferMemoryRequirements2KHR(c.device, c.pInfo, c.pMemoryRequirements)
}
func GetImageSparseMemoryRequirements2KHR(device Device, info *ImageSparseMemoryRequirementsInfo2, sparseMemoryRequirementCount *uint32, sparseMemoryRequirements []SparseImageMemoryRequirements2) {
	var c struct {
		device                        C.VkDevice
		pInfo                         *C.VkImageSparseMemoryRequirementsInfo2
		pSparseMemoryRequirementCount *C.uint32_t
		pSparseMemoryRequirements     *C.VkSparseImageMemoryRequirements2
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pInfo = (*C.VkImageSparseMemoryRequirementsInfo2)(_sa.alloc(C.sizeof_VkImageSparseMemoryRequirementsInfo2))
		info.toC(c.pInfo)
	}
	{
		c.pSparseMemoryRequirementCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pSparseMemoryRequirementCount = C.uint32_t(*sparseMemoryRequirementCount)
	}
	{
		c.pSparseMemoryRequirements = (*C.VkSparseImageMemoryRequirements2)(_sa.alloc(C.sizeof_VkSparseImageMemoryRequirements2 * uint(len(sparseMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkSparseImageMemoryRequirements2)(unsafe.Pointer(c.pSparseMemoryRequirements))[:len(sparseMemoryRequirements):len(sparseMemoryRequirements)]
		for i3, _ := range sparseMemoryRequirements {
			sparseMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageSparseMemoryRequirements2KHR(c.device, c.pInfo, c.pSparseMemoryRequirementCount, c.pSparseMemoryRequirements)
	*sparseMemoryRequirementCount = uint32(*c.pSparseMemoryRequirementCount)
}
func CreateSamplerYcbcrConversionKHR(device Device, createInfo *SamplerYcbcrConversionCreateInfo, allocator *AllocationCallbacks, ycbcrConversion *SamplerYcbcrConversion) (_ret Result) {
	var c struct {
		device           C.VkDevice
		pCreateInfo      *C.VkSamplerYcbcrConversionCreateInfo
		pAllocator       *C.VkAllocationCallbacks
		pYcbcrConversion *C.VkSamplerYcbcrConversion
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkSamplerYcbcrConversionCreateInfo)(_sa.alloc(C.sizeof_VkSamplerYcbcrConversionCreateInfo))
		createInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	{
		c.pYcbcrConversion = (*C.VkSamplerYcbcrConversion)(_sa.alloc(C.sizeof_VkSamplerYcbcrConversion))
		*c.pYcbcrConversion = C.VkSamplerYcbcrConversion(*ycbcrConversion)
	}
	c._ret = C.vkCreateSamplerYcbcrConversionKHR(c.device, c.pCreateInfo, c.pAllocator, c.pYcbcrConversion)
	_ret = Result(c._ret)
	*ycbcrConversion = SamplerYcbcrConversion(*c.pYcbcrConversion)
	return
}
func DestroySamplerYcbcrConversionKHR(device Device, ycbcrConversion SamplerYcbcrConversion, allocator *AllocationCallbacks) {
	var c struct {
		device          C.VkDevice
		ycbcrConversion C.VkSamplerYcbcrConversion
		pAllocator      *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.ycbcrConversion = C.VkSamplerYcbcrConversion(ycbcrConversion)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		allocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySamplerYcbcrConversionKHR(c.device, c.ycbcrConversion, c.pAllocator)
}
func BindBufferMemory2KHR(device Device, bindInfos []BindBufferMemoryInfo) (_ret Result) {
	var c struct {
		device        C.VkDevice
		bindInfoCount C.uint32_t
		pBindInfos    *C.VkBindBufferMemoryInfo
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.bindInfoCount = C.uint32_t(len(bindInfos))
	{
		c.pBindInfos = (*C.VkBindBufferMemoryInfo)(_sa.alloc(C.sizeof_VkBindBufferMemoryInfo * uint(len(bindInfos))))
		slice3 := (*[1 << 31]C.VkBindBufferMemoryInfo)(unsafe.Pointer(c.pBindInfos))[:len(bindInfos):len(bindInfos)]
		for i3, _ := range bindInfos {
			bindInfos[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkBindBufferMemory2KHR(c.device, c.bindInfoCount, c.pBindInfos)
	_ret = Result(c._ret)
	return
}
func BindImageMemory2KHR(device Device, bindInfos []BindImageMemoryInfo) (_ret Result) {
	var c struct {
		device        C.VkDevice
		bindInfoCount C.uint32_t
		pBindInfos    *C.VkBindImageMemoryInfo
		_ret          C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.bindInfoCount = C.uint32_t(len(bindInfos))
	{
		c.pBindInfos = (*C.VkBindImageMemoryInfo)(_sa.alloc(C.sizeof_VkBindImageMemoryInfo * uint(len(bindInfos))))
		slice3 := (*[1 << 31]C.VkBindImageMemoryInfo)(unsafe.Pointer(c.pBindInfos))[:len(bindInfos):len(bindInfos)]
		for i3, _ := range bindInfos {
			bindInfos[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkBindImageMemory2KHR(c.device, c.bindInfoCount, c.pBindInfos)
	_ret = Result(c._ret)
	return
}
func GetDescriptorSetLayoutSupportKHR(device Device, createInfo *DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) {
	var c struct {
		device      C.VkDevice
		pCreateInfo *C.VkDescriptorSetLayoutCreateInfo
		pSupport    *C.VkDescriptorSetLayoutSupport
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pCreateInfo = (*C.VkDescriptorSetLayoutCreateInfo)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutCreateInfo))
		createInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pSupport = (*C.VkDescriptorSetLayoutSupport)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutSupport))
		support.toC(c.pSupport)
	}
	C.vkGetDescriptorSetLayoutSupportKHR(c.device, c.pCreateInfo, c.pSupport)
	support.fromC(c.pSupport)
}
func CmdDrawIndirectCountKHR(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	var c struct {
		commandBuffer     C.VkCommandBuffer
		buffer            C.VkBuffer
		offset            C.VkDeviceSize
		countBuffer       C.VkBuffer
		countBufferOffset C.VkDeviceSize
		maxDrawCount      C.uint32_t
		stride            C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.countBuffer = C.VkBuffer(countBuffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(countBufferOffset))
		c.countBufferOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.maxDrawCount = C.uint32_t(maxDrawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndirectCountKHR(c.commandBuffer, c.buffer, c.offset, c.countBuffer, c.countBufferOffset, c.maxDrawCount, c.stride)
}
func CmdDrawIndexedIndirectCountKHR(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount uint32, stride uint32) {
	var c struct {
		commandBuffer     C.VkCommandBuffer
		buffer            C.VkBuffer
		offset            C.VkDeviceSize
		countBuffer       C.VkBuffer
		countBufferOffset C.VkDeviceSize
		maxDrawCount      C.uint32_t
		stride            C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.countBuffer = C.VkBuffer(countBuffer)
	{
		var temp_in_VkDeviceSize C.uint64_t
		temp_in_VkDeviceSize = C.uint64_t((uint64)(countBufferOffset))
		c.countBufferOffset = C.VkDeviceSize(temp_in_VkDeviceSize)
	}
	c.maxDrawCount = C.uint32_t(maxDrawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndexedIndirectCountKHR(c.commandBuffer, c.buffer, c.offset, c.countBuffer, c.countBufferOffset, c.maxDrawCount, c.stride)
}
