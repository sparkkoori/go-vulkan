package vk
//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
//typedef void * void_pointer;
import "C"

type VkStructureType int

const (
	VK_STRUCTURE_TYPE_APPLICATION_INFO                                             VkStructureType = 0
	VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO                                         VkStructureType = 1
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO                                     VkStructureType = 2
	VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO                                           VkStructureType = 3
	VK_STRUCTURE_TYPE_SUBMIT_INFO                                                  VkStructureType = 4
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO                                         VkStructureType = 5
	VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE                                          VkStructureType = 6
	VK_STRUCTURE_TYPE_BIND_SPARSE_INFO                                             VkStructureType = 7
	VK_STRUCTURE_TYPE_FENCE_CREATE_INFO                                            VkStructureType = 8
	VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO                                        VkStructureType = 9
	VK_STRUCTURE_TYPE_EVENT_CREATE_INFO                                            VkStructureType = 10
	VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO                                       VkStructureType = 11
	VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO                                           VkStructureType = 12
	VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO                                      VkStructureType = 13
	VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO                                            VkStructureType = 14
	VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO                                       VkStructureType = 15
	VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO                                    VkStructureType = 16
	VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO                                   VkStructureType = 17
	VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO                            VkStructureType = 18
	VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO                      VkStructureType = 19
	VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO                    VkStructureType = 20
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO                      VkStructureType = 21
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO                          VkStructureType = 22
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO                     VkStructureType = 23
	VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO                       VkStructureType = 24
	VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO                     VkStructureType = 25
	VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO                       VkStructureType = 26
	VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO                           VkStructureType = 27
	VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO                                VkStructureType = 28
	VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO                                 VkStructureType = 29
	VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO                                  VkStructureType = 30
	VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO                                          VkStructureType = 31
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO                            VkStructureType = 32
	VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO                                  VkStructureType = 33
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO                                 VkStructureType = 34
	VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET                                         VkStructureType = 35
	VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET                                          VkStructureType = 36
	VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO                                      VkStructureType = 37
	VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO                                      VkStructureType = 38
	VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO                                     VkStructureType = 39
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO                                 VkStructureType = 40
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO                              VkStructureType = 41
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO                                    VkStructureType = 42
	VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO                                       VkStructureType = 43
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER                                        VkStructureType = 44
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER                                         VkStructureType = 45
	VK_STRUCTURE_TYPE_MEMORY_BARRIER                                               VkStructureType = 46
	VK_STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO                                  VkStructureType = 47
	VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO                                    VkStructureType = 48
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES                          VkStructureType = 1000094000
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO                                      VkStructureType = 1000157000
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO                                       VkStructureType = 1000157001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES                       VkStructureType = 1000083000
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS                                VkStructureType = 1000127000
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO                               VkStructureType = 1000127001
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO                                   VkStructureType = 1000060000
	VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO                          VkStructureType = 1000060003
	VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO                       VkStructureType = 1000060004
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO                                     VkStructureType = 1000060005
	VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO                                VkStructureType = 1000060006
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO                         VkStructureType = 1000060013
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO                          VkStructureType = 1000060014
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES                             VkStructureType = 1000070000
	VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO                              VkStructureType = 1000070001
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2                            VkStructureType = 1000146000
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2                             VkStructureType = 1000146001
	VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2                      VkStructureType = 1000146002
	VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2                                        VkStructureType = 1000146003
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2                           VkStructureType = 1000146004
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2                                   VkStructureType = 1000059000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2                                 VkStructureType = 1000059001
	VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2                                          VkStructureType = 1000059002
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2                                    VkStructureType = 1000059003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2                          VkStructureType = 1000059004
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2                                    VkStructureType = 1000059005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2                          VkStructureType = 1000059006
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2                             VkStructureType = 1000059007
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2                   VkStructureType = 1000059008
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES                    VkStructureType = 1000117000
	VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO              VkStructureType = 1000117001
	VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO                                 VkStructureType = 1000117002
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO        VkStructureType = 1000117003
	VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO                            VkStructureType = 1000053000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES                           VkStructureType = 1000053001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES                         VkStructureType = 1000053002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES                    VkStructureType = 1000120000
	VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO                                        VkStructureType = 1000145000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES                    VkStructureType = 1000145001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES                  VkStructureType = 1000145002
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2                                          VkStructureType = 1000145003
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO                         VkStructureType = 1000156000
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO                                VkStructureType = 1000156001
	VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO                                 VkStructureType = 1000156002
	VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO                         VkStructureType = 1000156003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES            VkStructureType = 1000156004
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES             VkStructureType = 1000156005
	VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO                       VkStructureType = 1000085000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO                   VkStructureType = 1000071000
	VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES                             VkStructureType = 1000071001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO                         VkStructureType = 1000071002
	VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES                                   VkStructureType = 1000071003
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES                                VkStructureType = 1000071004
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO                           VkStructureType = 1000072000
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO                            VkStructureType = 1000072001
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO                                  VkStructureType = 1000072002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO                          VkStructureType = 1000112000
	VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES                                    VkStructureType = 1000112001
	VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO                                     VkStructureType = 1000113000
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO                                 VkStructureType = 1000077000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO                      VkStructureType = 1000076000
	VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES                                VkStructureType = 1000076001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES                     VkStructureType = 1000168000
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT                                VkStructureType = 1000168001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES               VkStructureType = 1000063000
	VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR                                    VkStructureType = 1000001000
	VK_STRUCTURE_TYPE_PRESENT_INFO_KHR                                             VkStructureType = 1000001001
	VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR                        VkStructureType = 1000060007
	VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR                              VkStructureType = 1000060008
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR                         VkStructureType = 1000060009
	VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR                                  VkStructureType = 1000060010
	VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR                                VkStructureType = 1000060011
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR                       VkStructureType = 1000060012
	VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR                                 VkStructureType = 1000002000
	VK_STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR                              VkStructureType = 1000002001
	VK_STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR                                     VkStructureType = 1000003000
	VK_STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR                                 VkStructureType = 1000004000
	VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR                                  VkStructureType = 1000005000
	VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR                              VkStructureType = 1000006000
	VK_STRUCTURE_TYPE_MIR_SURFACE_CREATE_INFO_KHR                                  VkStructureType = 1000007000
	VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR                              VkStructureType = 1000008000
	VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR                                VkStructureType = 1000009000
	VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT                        VkStructureType = 1000011000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD         VkStructureType = 1000018000
	VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT                            VkStructureType = 1000022000
	VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT                             VkStructureType = 1000022001
	VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT                                 VkStructureType = 1000022002
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV                    VkStructureType = 1000026000
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV                   VkStructureType = 1000026001
	VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV                 VkStructureType = 1000026002
	VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD                     VkStructureType = 1000041000
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV                         VkStructureType = 1000056000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV                               VkStructureType = 1000056001
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV                           VkStructureType = 1000057000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV                           VkStructureType = 1000057001
	VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV                    VkStructureType = 1000058000
	VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT                                         VkStructureType = 1000061000
	VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN                                    VkStructureType = 1000062000
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR                          VkStructureType = 1000073000
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR                          VkStructureType = 1000073001
	VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR                           VkStructureType = 1000073002
	VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR                             VkStructureType = 1000073003
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR                                    VkStructureType = 1000074000
	VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR                                     VkStructureType = 1000074001
	VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR                                       VkStructureType = 1000074002
	VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR                   VkStructureType = 1000075000
	VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                       VkStructureType = 1000078000
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                       VkStructureType = 1000078001
	VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR                                  VkStructureType = 1000078002
	VK_STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR                          VkStructureType = 1000078003
	VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR                                 VkStructureType = 1000079000
	VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR                                    VkStructureType = 1000079001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR               VkStructureType = 1000080000
	VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT    VkStructureType = 1000081000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT           VkStructureType = 1000081001
	VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT                         VkStructureType = 1000081002
	VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR                                          VkStructureType = 1000084000
	VK_STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX                                 VkStructureType = 1000086000
	VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX                     VkStructureType = 1000086001
	VK_STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX                                VkStructureType = 1000086002
	VK_STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX                      VkStructureType = 1000086003
	VK_STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX                         VkStructureType = 1000086004
	VK_STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX                       VkStructureType = 1000086005
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV             VkStructureType = 1000087000
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT                                   VkStructureType = 1000090000
	VK_STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT                                       VkStructureType = 1000091000
	VK_STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT                                        VkStructureType = 1000091001
	VK_STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT                                       VkStructureType = 1000091002
	VK_STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT                            VkStructureType = 1000091003
	VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE                                    VkStructureType = 1000092000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX VkStructureType = 1000097000
	VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV               VkStructureType = 1000098000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT             VkStructureType = 1000099000
	VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT             VkStructureType = 1000099001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT    VkStructureType = 1000101000
	VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT    VkStructureType = 1000101001
	VK_STRUCTURE_TYPE_HDR_METADATA_EXT                                             VkStructureType = 1000105000
	VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR                                 VkStructureType = 1000109000
	VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR                                   VkStructureType = 1000109001
	VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR                                    VkStructureType = 1000109002
	VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR                                     VkStructureType = 1000109003
	VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR                                VkStructureType = 1000109004
	VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR                                       VkStructureType = 1000109005
	VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR                                         VkStructureType = 1000109006
	VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR                      VkStructureType = 1000111000
	VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR                           VkStructureType = 1000114000
	VK_STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR                           VkStructureType = 1000114001
	VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR                              VkStructureType = 1000114002
	VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR                                     VkStructureType = 1000115000
	VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR                                        VkStructureType = 1000115001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR                           VkStructureType = 1000119000
	VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR                                   VkStructureType = 1000119001
	VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR                                         VkStructureType = 1000119002
	VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR                                     VkStructureType = 1000121000
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR                               VkStructureType = 1000121001
	VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR                                VkStructureType = 1000121002
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR                                     VkStructureType = 1000121003
	VK_STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR                             VkStructureType = 1000121004
	VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK                                  VkStructureType = 1000122000
	VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK                                VkStructureType = 1000123000
	VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT                             VkStructureType = 1000128000
	VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT                              VkStructureType = 1000128001
	VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT                                        VkStructureType = 1000128002
	VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT                      VkStructureType = 1000128003
	VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT                        VkStructureType = 1000128004
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID                        VkStructureType = 1000129000
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID                   VkStructureType = 1000129001
	VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID            VkStructureType = 1000129002
	VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                  VkStructureType = 1000129003
	VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID              VkStructureType = 1000129004
	VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID                                      VkStructureType = 1000129005
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT         VkStructureType = 1000130000
	VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT                       VkStructureType = 1000130001
	VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT                                    VkStructureType = 1000143000
	VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT                  VkStructureType = 1000143001
	VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT              VkStructureType = 1000143002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT              VkStructureType = 1000143003
	VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT                                   VkStructureType = 1000143004
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR                            VkStructureType = 1000147000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT        VkStructureType = 1000148000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT      VkStructureType = 1000148001
	VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT          VkStructureType = 1000148002
	VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV              VkStructureType = 1000149000
	VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV            VkStructureType = 1000152000
	VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT                             VkStructureType = 1000160000
	VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT               VkStructureType = 1000160001
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT          VkStructureType = 1000161000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT             VkStructureType = 1000161001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT           VkStructureType = 1000161002
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT   VkStructureType = 1000161003
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT  VkStructureType = 1000161004
	VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT                 VkStructureType = 1000174000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR                    VkStructureType = 1000177000
	VK_STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT                          VkStructureType = 1000178000
	VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT                           VkStructureType = 1000178001
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT          VkStructureType = 1000178002
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD                   VkStructureType = 1000185000
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT      VkStructureType = 1000190000
	VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT          VkStructureType = 1000190001
	VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV                                           VkStructureType = 1000206000
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV                        VkStructureType = 1000206001
	VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR                        VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR                       VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR                               VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR                             VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR                                      VkStructureType = VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR                                VkStructureType = VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR                      VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR                                VkStructureType = VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR                      VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR                         VkStructureType = VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR               VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR                               VkStructureType = VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR                      VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR                   VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR                            VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR                      VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR                         VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR                          VkStructureType = VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR               VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR                         VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR                               VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR                            VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR                       VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR                        VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR                              VkStructureType = VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR                  VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR                            VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR                   VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR                   VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR                      VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR                                VkStructureType = VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR                                 VkStructureType = VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR                VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR          VkStructureType = VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR    VkStructureType = VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR                VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR                            VkStructureType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR                           VkStructureType = VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR                        VkStructureType = VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR                         VkStructureType = VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR                  VkStructureType = VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR                                    VkStructureType = VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR                       VkStructureType = VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR                            VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR                             VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR                     VkStructureType = VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR        VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR         VkStructureType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR                                  VkStructureType = VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR                                   VkStructureType = VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR                 VkStructureType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR                            VkStructureType = VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	VK_STRUCTURE_TYPE_BEGIN_RANGE                                                  VkStructureType = VK_STRUCTURE_TYPE_APPLICATION_INFO
	VK_STRUCTURE_TYPE_END_RANGE                                                    VkStructureType = VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO
	VK_STRUCTURE_TYPE_RANGE_SIZE                                                   VkStructureType = (VK_STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO - VK_STRUCTURE_TYPE_APPLICATION_INFO + 1)
	VK_STRUCTURE_TYPE_MAX_ENUM                                                     VkStructureType = 2147483647
)

type VkFlags uint32
type VkInstanceCreateFlags VkFlags
type VkApplicationInfo struct {
	sType              VkStructureType
	pNext              *[0]byte
	pApplicationName   string
	applicationVersion uint32
	pEngineName        string
	engineVersion      uint32
	apiVersion         uint32
}

func (g *VkApplicationInfo) toC(c *C.VkApplicationInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.pApplicationName = toCString(g.pApplicationName, _sa)
	c.applicationVersion = C.uint32_t(g.applicationVersion)
	c.pEngineName = toCString(g.pEngineName, _sa)
	c.engineVersion = C.uint32_t(g.engineVersion)
	c.apiVersion = C.uint32_t(g.apiVersion)
}
func (g *VkApplicationInfo) fromC(c *C.VkApplicationInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.pApplicationName = toGoString(c.pApplicationName)
	g.applicationVersion = uint32(c.applicationVersion)
	g.pEngineName = toGoString(c.pEngineName)
	g.engineVersion = uint32(c.engineVersion)
	g.apiVersion = uint32(c.apiVersion)
}

type VkInstanceCreateInfo struct {
	sType                   VkStructureType
	pNext                   *[0]byte
	flags                   VkInstanceCreateFlags
	pApplicationInfo        *VkApplicationInfo
	ppEnabledLayerNames     []string
	ppEnabledExtensionNames []string
}

func (g *VkInstanceCreateInfo) toC(c *C.VkInstanceCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkInstanceCreateFlags(_temp)
	}
	{
		c.pApplicationInfo = (*C.VkApplicationInfo)(_sa.alloc(C.sizeof_VkApplicationInfo))
		g.pApplicationInfo.toC(c.pApplicationInfo, _sa)
	}
	c.enabledLayerCount = C.uint32_t(len(g.ppEnabledLayerNames))
	{
		c.ppEnabledLayerNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.ppEnabledLayerNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.ppEnabledLayerNames):len(g.ppEnabledLayerNames)]
		for i2, _ := range g.ppEnabledLayerNames {
			slice2[i2] = toCString(g.ppEnabledLayerNames[i2], _sa)
		}
	}
	c.enabledExtensionCount = C.uint32_t(len(g.ppEnabledExtensionNames))
	{
		c.ppEnabledExtensionNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.ppEnabledExtensionNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.ppEnabledExtensionNames):len(g.ppEnabledExtensionNames)]
		for i2, _ := range g.ppEnabledExtensionNames {
			slice2[i2] = toCString(g.ppEnabledExtensionNames[i2], _sa)
		}
	}
}
func (g *VkInstanceCreateInfo) fromC(c *C.VkInstanceCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkInstanceCreateFlags(_temp)
	}
	{
		if g.pApplicationInfo == nil {
			g.pApplicationInfo = new(VkApplicationInfo)
		}
		g.pApplicationInfo.fromC(c.pApplicationInfo)
	}
	g.ppEnabledLayerNames = make([]string, int(c.enabledLayerCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.ppEnabledLayerNames):len(g.ppEnabledLayerNames)]
		for i2, _ := range g.ppEnabledLayerNames {
			g.ppEnabledLayerNames[i2] = toGoString(slice2[i2])
		}
	}
	g.ppEnabledExtensionNames = make([]string, int(c.enabledExtensionCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.ppEnabledExtensionNames):len(g.ppEnabledExtensionNames)]
		for i2, _ := range g.ppEnabledExtensionNames {
			g.ppEnabledExtensionNames[i2] = toGoString(slice2[i2])
		}
	}
}

type PFN_vkAllocationFunction struct {
	Raw C.PFN_vkAllocationFunction
}
type VkSystemAllocationScope int

const (
	VK_SYSTEM_ALLOCATION_SCOPE_COMMAND     VkSystemAllocationScope = 0
	VK_SYSTEM_ALLOCATION_SCOPE_OBJECT      VkSystemAllocationScope = 1
	VK_SYSTEM_ALLOCATION_SCOPE_CACHE       VkSystemAllocationScope = 2
	VK_SYSTEM_ALLOCATION_SCOPE_DEVICE      VkSystemAllocationScope = 3
	VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE    VkSystemAllocationScope = 4
	VK_SYSTEM_ALLOCATION_SCOPE_BEGIN_RANGE VkSystemAllocationScope = VK_SYSTEM_ALLOCATION_SCOPE_COMMAND
	VK_SYSTEM_ALLOCATION_SCOPE_END_RANGE   VkSystemAllocationScope = VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE
	VK_SYSTEM_ALLOCATION_SCOPE_RANGE_SIZE  VkSystemAllocationScope = (VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE - VK_SYSTEM_ALLOCATION_SCOPE_COMMAND + 1)
	VK_SYSTEM_ALLOCATION_SCOPE_MAX_ENUM    VkSystemAllocationScope = 2147483647
)

func (p PFN_vkAllocationFunction) Call( unsafe.Pointer,  uint,  uint,  VkSystemAllocationScope) (_ret unsafe.Pointer) {
	var c struct {
			unsafe.Pointer
			C.size_t
			C.size_t
			C.VkSystemAllocationScope
		_ret unsafe.Pointer
	}
	c. =
	c. = C.size_t()
	c. = C.size_t()
	c. = C.VkSystemAllocationScope()
	c._ret = C.callPFN_vkAllocationFunction(p.Raw, c., c., c., c.)
	_ret = c._ret
	return
}

type PFN_vkReallocationFunction struct {
	Raw C.PFN_vkReallocationFunction
}

func (p PFN_vkReallocationFunction) Call( unsafe.Pointer,  unsafe.Pointer,  uint,  uint,  VkSystemAllocationScope) (_ret unsafe.Pointer) {
	var c struct {
			unsafe.Pointer
			unsafe.Pointer
			C.size_t
			C.size_t
			C.VkSystemAllocationScope
		_ret unsafe.Pointer
	}
	c. =
	c. =
	c. = C.size_t()
	c. = C.size_t()
	c. = C.VkSystemAllocationScope()
	c._ret = C.callPFN_vkReallocationFunction(p.Raw, c., c., c., c., c.)
	_ret = c._ret
	return
}

type PFN_vkFreeFunction struct {
	Raw C.PFN_vkFreeFunction
}

func (p PFN_vkFreeFunction) Call( unsafe.Pointer,  unsafe.Pointer) {
	var c struct {
		unsafe.Pointer
		unsafe.Pointer
	}
	c. =
	c. =
	C.callPFN_vkFreeFunction(p.Raw, c., c.)
}

type PFN_vkInternalAllocationNotification struct {
	Raw C.PFN_vkInternalAllocationNotification
}
type VkInternalAllocationType int

const (
	VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE  VkInternalAllocationType = 0
	VK_INTERNAL_ALLOCATION_TYPE_BEGIN_RANGE VkInternalAllocationType = VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	VK_INTERNAL_ALLOCATION_TYPE_END_RANGE   VkInternalAllocationType = VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	VK_INTERNAL_ALLOCATION_TYPE_RANGE_SIZE  VkInternalAllocationType = (VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE - VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE + 1)
	VK_INTERNAL_ALLOCATION_TYPE_MAX_ENUM    VkInternalAllocationType = 2147483647
)

func (p PFN_vkInternalAllocationNotification) Call( unsafe.Pointer,  uint,  VkInternalAllocationType,  VkSystemAllocationScope) {
	var c struct {
		unsafe.Pointer
		C.size_t
		C.VkInternalAllocationType
		C.VkSystemAllocationScope
	}
	c. =
	c. = C.size_t()
	c. = C.VkInternalAllocationType()
	c. = C.VkSystemAllocationScope()
	C.callPFN_vkInternalAllocationNotification(p.Raw, c., c., c., c.)
}

type PFN_vkInternalFreeNotification struct {
	Raw C.PFN_vkInternalFreeNotification
}

func (p PFN_vkInternalFreeNotification) Call( unsafe.Pointer,  uint,  VkInternalAllocationType,  VkSystemAllocationScope) {
	var c struct {
		unsafe.Pointer
		C.size_t
		C.VkInternalAllocationType
		C.VkSystemAllocationScope
	}
	c. =
	c. = C.size_t()
	c. = C.VkInternalAllocationType()
	c. = C.VkSystemAllocationScope()
	C.callPFN_vkInternalFreeNotification(p.Raw, c., c., c., c.)
}

type VkAllocationCallbacks struct {
	pUserData             []byte
	pfnAllocation         PFN_vkAllocationFunction
	pfnReallocation       PFN_vkReallocationFunction
	pfnFree               PFN_vkFreeFunction
	pfnInternalAllocation PFN_vkInternalAllocationNotification
	pfnInternalFree       PFN_vkInternalFreeNotification
}

func (g *VkAllocationCallbacks) toC(c *C.VkAllocationCallbacks, _sa *stackAllocator) {
	{
		c.pUserData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.pUserData)))
		slice2 := (*[1 << 31]byte)(c.pUserData)[:len(g.pUserData):len(g.pUserData)]
		for i2, _ := range g.pUserData {
			slice2[i2] = g.pUserData[i2]
		}
	}
	c.pfnAllocation = g.pfnAllocation.Raw
	c.pfnReallocation = g.pfnReallocation.Raw
	c.pfnFree = g.pfnFree.Raw
	c.pfnInternalAllocation = g.pfnInternalAllocation.Raw
	c.pfnInternalFree = g.pfnInternalFree.Raw
}
func (g *VkAllocationCallbacks) fromC(c *C.VkAllocationCallbacks) {
	{
		slice2 := (*[1 << 31]byte)(c.pUserData)[:len(g.pUserData):len(g.pUserData)]
		for i2, _ := range g.pUserData {
			g.pUserData[i2] = slice2[i2]
		}
	}
	g.pfnAllocation.Raw = c.pfnAllocation
	g.pfnReallocation.Raw = c.pfnReallocation
	g.pfnFree.Raw = c.pfnFree
	g.pfnInternalAllocation.Raw = c.pfnInternalAllocation
	g.pfnInternalFree.Raw = c.pfnInternalFree
}

type VkInstance C.VkInstance
type VkResult int

const (
	VK_SUCCESS                           VkResult = 0
	VK_NOT_READY                         VkResult = 1
	VK_TIMEOUT                           VkResult = 2
	VK_EVENT_SET                         VkResult = 3
	VK_EVENT_RESET                       VkResult = 4
	VK_INCOMPLETE                        VkResult = 5
	VK_ERROR_OUT_OF_HOST_MEMORY          VkResult = -1
	VK_ERROR_OUT_OF_DEVICE_MEMORY        VkResult = -2
	VK_ERROR_INITIALIZATION_FAILED       VkResult = -3
	VK_ERROR_DEVICE_LOST                 VkResult = -4
	VK_ERROR_MEMORY_MAP_FAILED           VkResult = -5
	VK_ERROR_LAYER_NOT_PRESENT           VkResult = -6
	VK_ERROR_EXTENSION_NOT_PRESENT       VkResult = -7
	VK_ERROR_FEATURE_NOT_PRESENT         VkResult = -8
	VK_ERROR_INCOMPATIBLE_DRIVER         VkResult = -9
	VK_ERROR_TOO_MANY_OBJECTS            VkResult = -10
	VK_ERROR_FORMAT_NOT_SUPPORTED        VkResult = -11
	VK_ERROR_FRAGMENTED_POOL             VkResult = -12
	VK_ERROR_OUT_OF_POOL_MEMORY          VkResult = -1000069000
	VK_ERROR_INVALID_EXTERNAL_HANDLE     VkResult = -1000072003
	VK_ERROR_SURFACE_LOST_KHR            VkResult = -1000000000
	VK_ERROR_NATIVE_WINDOW_IN_USE_KHR    VkResult = -1000000001
	VK_SUBOPTIMAL_KHR                    VkResult = 1000001003
	VK_ERROR_OUT_OF_DATE_KHR             VkResult = -1000001004
	VK_ERROR_INCOMPATIBLE_DISPLAY_KHR    VkResult = -1000003001
	VK_ERROR_VALIDATION_FAILED_EXT       VkResult = -1000011001
	VK_ERROR_INVALID_SHADER_NV           VkResult = -1000012000
	VK_ERROR_FRAGMENTATION_EXT           VkResult = -1000161000
	VK_ERROR_NOT_PERMITTED_EXT           VkResult = -1000174001
	VK_ERROR_OUT_OF_POOL_MEMORY_KHR      VkResult = VK_ERROR_OUT_OF_POOL_MEMORY
	VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR VkResult = VK_ERROR_INVALID_EXTERNAL_HANDLE
	VK_RESULT_BEGIN_RANGE                VkResult = VK_ERROR_FRAGMENTED_POOL
	VK_RESULT_END_RANGE                  VkResult = VK_INCOMPLETE
	VK_RESULT_RANGE_SIZE                 VkResult = (VK_INCOMPLETE - VK_ERROR_FRAGMENTED_POOL + 1)
	VK_RESULT_MAX_ENUM                   VkResult = 2147483647
)

func vkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pInstance = (*C.VkInstance)(_sa.alloc(C.sizeof_VkInstance))
		*c.pInstance = C.VkInstance(*pInstance)
	}
	c._ret = C.vkCreateInstance(c.pCreateInfo, c.pAllocator, c.pInstance)
	_ret = VkResult(c._ret)
	*pInstance = VkInstance(*c.pInstance)
	return
}
func vkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) {
	var c struct {
		instance   C.VkInstance
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyInstance(c.instance, c.pAllocator)
}

type VkPhysicalDevice C.VkPhysicalDevice

func vkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *uint32, pPhysicalDevices []VkPhysicalDevice) (_ret VkResult) {
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
		*c.pPhysicalDeviceCount = C.uint32_t(*pPhysicalDeviceCount)
	}
	{
		c.pPhysicalDevices = (*C.VkPhysicalDevice)(_sa.alloc(C.sizeof_VkPhysicalDevice * uint(len(pPhysicalDevices))))
		slice3 := (*[1 << 31]C.VkPhysicalDevice)(unsafe.Pointer(c.pPhysicalDevices))[:len(pPhysicalDevices):len(pPhysicalDevices)]
		for i3, _ := range pPhysicalDevices {
			slice3[i3] = C.VkPhysicalDevice(pPhysicalDevices[i3])
		}
	}
	c._ret = C.vkEnumeratePhysicalDevices(c.instance, c.pPhysicalDeviceCount, c.pPhysicalDevices)
	_ret = VkResult(c._ret)
	*pPhysicalDeviceCount = uint32(*c.pPhysicalDeviceCount)
	return
}

type VkBool32 uint32
type VkPhysicalDeviceFeatures struct {
	robustBufferAccess                      VkBool32
	fullDrawIndexUint32                     VkBool32
	imageCubeArray                          VkBool32
	independentBlend                        VkBool32
	geometryShader                          VkBool32
	tessellationShader                      VkBool32
	sampleRateShading                       VkBool32
	dualSrcBlend                            VkBool32
	logicOp                                 VkBool32
	multiDrawIndirect                       VkBool32
	drawIndirectFirstInstance               VkBool32
	depthClamp                              VkBool32
	depthBiasClamp                          VkBool32
	fillModeNonSolid                        VkBool32
	depthBounds                             VkBool32
	wideLines                               VkBool32
	largePoints                             VkBool32
	alphaToOne                              VkBool32
	multiViewport                           VkBool32
	samplerAnisotropy                       VkBool32
	textureCompressionETC2                  VkBool32
	textureCompressionASTC_LDR              VkBool32
	textureCompressionBC                    VkBool32
	occlusionQueryPrecise                   VkBool32
	pipelineStatisticsQuery                 VkBool32
	vertexPipelineStoresAndAtomics          VkBool32
	fragmentStoresAndAtomics                VkBool32
	shaderTessellationAndGeometryPointSize  VkBool32
	shaderImageGatherExtended               VkBool32
	shaderStorageImageExtendedFormats       VkBool32
	shaderStorageImageMultisample           VkBool32
	shaderStorageImageReadWithoutFormat     VkBool32
	shaderStorageImageWriteWithoutFormat    VkBool32
	shaderUniformBufferArrayDynamicIndexing VkBool32
	shaderSampledImageArrayDynamicIndexing  VkBool32
	shaderStorageBufferArrayDynamicIndexing VkBool32
	shaderStorageImageArrayDynamicIndexing  VkBool32
	shaderClipDistance                      VkBool32
	shaderCullDistance                      VkBool32
	shaderFloat64                           VkBool32
	shaderInt64                             VkBool32
	shaderInt16                             VkBool32
	shaderResourceResidency                 VkBool32
	shaderResourceMinLod                    VkBool32
	sparseBinding                           VkBool32
	sparseResidencyBuffer                   VkBool32
	sparseResidencyImage2D                  VkBool32
	sparseResidencyImage3D                  VkBool32
	sparseResidency2Samples                 VkBool32
	sparseResidency4Samples                 VkBool32
	sparseResidency8Samples                 VkBool32
	sparseResidency16Samples                VkBool32
	sparseResidencyAliased                  VkBool32
	variableMultisampleRate                 VkBool32
	inheritedQueries                        VkBool32
}

func (g *VkPhysicalDeviceFeatures) toC(c *C.VkPhysicalDeviceFeatures) {
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.robustBufferAccess))
		c.robustBufferAccess = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.fullDrawIndexUint32))
		c.fullDrawIndexUint32 = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.imageCubeArray))
		c.imageCubeArray = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.independentBlend))
		c.independentBlend = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.geometryShader))
		c.geometryShader = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.tessellationShader))
		c.tessellationShader = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sampleRateShading))
		c.sampleRateShading = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.dualSrcBlend))
		c.dualSrcBlend = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.logicOp))
		c.logicOp = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.multiDrawIndirect))
		c.multiDrawIndirect = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.drawIndirectFirstInstance))
		c.drawIndirectFirstInstance = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthClamp))
		c.depthClamp = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthBiasClamp))
		c.depthBiasClamp = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.fillModeNonSolid))
		c.fillModeNonSolid = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthBounds))
		c.depthBounds = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.wideLines))
		c.wideLines = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.largePoints))
		c.largePoints = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.alphaToOne))
		c.alphaToOne = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.multiViewport))
		c.multiViewport = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.samplerAnisotropy))
		c.samplerAnisotropy = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.textureCompressionETC2))
		c.textureCompressionETC2 = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.textureCompressionASTC_LDR))
		c.textureCompressionASTC_LDR = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.textureCompressionBC))
		c.textureCompressionBC = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.occlusionQueryPrecise))
		c.occlusionQueryPrecise = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.pipelineStatisticsQuery))
		c.pipelineStatisticsQuery = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.vertexPipelineStoresAndAtomics))
		c.vertexPipelineStoresAndAtomics = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.fragmentStoresAndAtomics))
		c.fragmentStoresAndAtomics = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderTessellationAndGeometryPointSize))
		c.shaderTessellationAndGeometryPointSize = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderImageGatherExtended))
		c.shaderImageGatherExtended = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageImageExtendedFormats))
		c.shaderStorageImageExtendedFormats = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageImageMultisample))
		c.shaderStorageImageMultisample = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageImageReadWithoutFormat))
		c.shaderStorageImageReadWithoutFormat = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageImageWriteWithoutFormat))
		c.shaderStorageImageWriteWithoutFormat = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderUniformBufferArrayDynamicIndexing))
		c.shaderUniformBufferArrayDynamicIndexing = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderSampledImageArrayDynamicIndexing))
		c.shaderSampledImageArrayDynamicIndexing = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageBufferArrayDynamicIndexing))
		c.shaderStorageBufferArrayDynamicIndexing = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderStorageImageArrayDynamicIndexing))
		c.shaderStorageImageArrayDynamicIndexing = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderClipDistance))
		c.shaderClipDistance = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderCullDistance))
		c.shaderCullDistance = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderFloat64))
		c.shaderFloat64 = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderInt64))
		c.shaderInt64 = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderInt16))
		c.shaderInt16 = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderResourceResidency))
		c.shaderResourceResidency = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.shaderResourceMinLod))
		c.shaderResourceMinLod = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseBinding))
		c.sparseBinding = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidencyBuffer))
		c.sparseResidencyBuffer = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidencyImage2D))
		c.sparseResidencyImage2D = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidencyImage3D))
		c.sparseResidencyImage3D = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidency2Samples))
		c.sparseResidency2Samples = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidency4Samples))
		c.sparseResidency4Samples = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidency8Samples))
		c.sparseResidency8Samples = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidency16Samples))
		c.sparseResidency16Samples = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sparseResidencyAliased))
		c.sparseResidencyAliased = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.variableMultisampleRate))
		c.variableMultisampleRate = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.inheritedQueries))
		c.inheritedQueries = C.VkBool32(_temp)
	}
}
func (g *VkPhysicalDeviceFeatures) fromC(c *C.VkPhysicalDeviceFeatures) {
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.robustBufferAccess))
		g.robustBufferAccess = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.fullDrawIndexUint32))
		g.fullDrawIndexUint32 = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.imageCubeArray))
		g.imageCubeArray = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.independentBlend))
		g.independentBlend = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.geometryShader))
		g.geometryShader = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.tessellationShader))
		g.tessellationShader = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sampleRateShading))
		g.sampleRateShading = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.dualSrcBlend))
		g.dualSrcBlend = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.logicOp))
		g.logicOp = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.multiDrawIndirect))
		g.multiDrawIndirect = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.drawIndirectFirstInstance))
		g.drawIndirectFirstInstance = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthClamp))
		g.depthClamp = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthBiasClamp))
		g.depthBiasClamp = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.fillModeNonSolid))
		g.fillModeNonSolid = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthBounds))
		g.depthBounds = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.wideLines))
		g.wideLines = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.largePoints))
		g.largePoints = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.alphaToOne))
		g.alphaToOne = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.multiViewport))
		g.multiViewport = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.samplerAnisotropy))
		g.samplerAnisotropy = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.textureCompressionETC2))
		g.textureCompressionETC2 = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.textureCompressionASTC_LDR))
		g.textureCompressionASTC_LDR = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.textureCompressionBC))
		g.textureCompressionBC = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.occlusionQueryPrecise))
		g.occlusionQueryPrecise = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.pipelineStatisticsQuery))
		g.pipelineStatisticsQuery = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.vertexPipelineStoresAndAtomics))
		g.vertexPipelineStoresAndAtomics = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.fragmentStoresAndAtomics))
		g.fragmentStoresAndAtomics = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderTessellationAndGeometryPointSize))
		g.shaderTessellationAndGeometryPointSize = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderImageGatherExtended))
		g.shaderImageGatherExtended = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageImageExtendedFormats))
		g.shaderStorageImageExtendedFormats = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageImageMultisample))
		g.shaderStorageImageMultisample = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageImageReadWithoutFormat))
		g.shaderStorageImageReadWithoutFormat = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageImageWriteWithoutFormat))
		g.shaderStorageImageWriteWithoutFormat = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderUniformBufferArrayDynamicIndexing))
		g.shaderUniformBufferArrayDynamicIndexing = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderSampledImageArrayDynamicIndexing))
		g.shaderSampledImageArrayDynamicIndexing = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageBufferArrayDynamicIndexing))
		g.shaderStorageBufferArrayDynamicIndexing = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderStorageImageArrayDynamicIndexing))
		g.shaderStorageImageArrayDynamicIndexing = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderClipDistance))
		g.shaderClipDistance = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderCullDistance))
		g.shaderCullDistance = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderFloat64))
		g.shaderFloat64 = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderInt64))
		g.shaderInt64 = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderInt16))
		g.shaderInt16 = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderResourceResidency))
		g.shaderResourceResidency = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.shaderResourceMinLod))
		g.shaderResourceMinLod = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseBinding))
		g.sparseBinding = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidencyBuffer))
		g.sparseResidencyBuffer = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidencyImage2D))
		g.sparseResidencyImage2D = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidencyImage3D))
		g.sparseResidencyImage3D = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidency2Samples))
		g.sparseResidency2Samples = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidency4Samples))
		g.sparseResidency4Samples = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidency8Samples))
		g.sparseResidency8Samples = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidency16Samples))
		g.sparseResidency16Samples = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sparseResidencyAliased))
		g.sparseResidencyAliased = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.variableMultisampleRate))
		g.variableMultisampleRate = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.inheritedQueries))
		g.inheritedQueries = VkBool32(_temp)
	}
}
func vkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures []VkPhysicalDeviceFeatures) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pFeatures      *C.VkPhysicalDeviceFeatures
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pFeatures = (*C.VkPhysicalDeviceFeatures)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures * uint(len(pFeatures))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pFeatures))[:len(pFeatures):len(pFeatures)]
		for i3, _ := range pFeatures {
			pFeatures[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFeatures(c.physicalDevice, c.pFeatures)
}

type VkFormat int

const (
	VK_FORMAT_UNDEFINED                                      VkFormat = 0
	VK_FORMAT_R4G4_UNORM_PACK8                               VkFormat = 1
	VK_FORMAT_R4G4B4A4_UNORM_PACK16                          VkFormat = 2
	VK_FORMAT_B4G4R4A4_UNORM_PACK16                          VkFormat = 3
	VK_FORMAT_R5G6B5_UNORM_PACK16                            VkFormat = 4
	VK_FORMAT_B5G6R5_UNORM_PACK16                            VkFormat = 5
	VK_FORMAT_R5G5B5A1_UNORM_PACK16                          VkFormat = 6
	VK_FORMAT_B5G5R5A1_UNORM_PACK16                          VkFormat = 7
	VK_FORMAT_A1R5G5B5_UNORM_PACK16                          VkFormat = 8
	VK_FORMAT_R8_UNORM                                       VkFormat = 9
	VK_FORMAT_R8_SNORM                                       VkFormat = 10
	VK_FORMAT_R8_USCALED                                     VkFormat = 11
	VK_FORMAT_R8_SSCALED                                     VkFormat = 12
	VK_FORMAT_R8_UINT                                        VkFormat = 13
	VK_FORMAT_R8_SINT                                        VkFormat = 14
	VK_FORMAT_R8_SRGB                                        VkFormat = 15
	VK_FORMAT_R8G8_UNORM                                     VkFormat = 16
	VK_FORMAT_R8G8_SNORM                                     VkFormat = 17
	VK_FORMAT_R8G8_USCALED                                   VkFormat = 18
	VK_FORMAT_R8G8_SSCALED                                   VkFormat = 19
	VK_FORMAT_R8G8_UINT                                      VkFormat = 20
	VK_FORMAT_R8G8_SINT                                      VkFormat = 21
	VK_FORMAT_R8G8_SRGB                                      VkFormat = 22
	VK_FORMAT_R8G8B8_UNORM                                   VkFormat = 23
	VK_FORMAT_R8G8B8_SNORM                                   VkFormat = 24
	VK_FORMAT_R8G8B8_USCALED                                 VkFormat = 25
	VK_FORMAT_R8G8B8_SSCALED                                 VkFormat = 26
	VK_FORMAT_R8G8B8_UINT                                    VkFormat = 27
	VK_FORMAT_R8G8B8_SINT                                    VkFormat = 28
	VK_FORMAT_R8G8B8_SRGB                                    VkFormat = 29
	VK_FORMAT_B8G8R8_UNORM                                   VkFormat = 30
	VK_FORMAT_B8G8R8_SNORM                                   VkFormat = 31
	VK_FORMAT_B8G8R8_USCALED                                 VkFormat = 32
	VK_FORMAT_B8G8R8_SSCALED                                 VkFormat = 33
	VK_FORMAT_B8G8R8_UINT                                    VkFormat = 34
	VK_FORMAT_B8G8R8_SINT                                    VkFormat = 35
	VK_FORMAT_B8G8R8_SRGB                                    VkFormat = 36
	VK_FORMAT_R8G8B8A8_UNORM                                 VkFormat = 37
	VK_FORMAT_R8G8B8A8_SNORM                                 VkFormat = 38
	VK_FORMAT_R8G8B8A8_USCALED                               VkFormat = 39
	VK_FORMAT_R8G8B8A8_SSCALED                               VkFormat = 40
	VK_FORMAT_R8G8B8A8_UINT                                  VkFormat = 41
	VK_FORMAT_R8G8B8A8_SINT                                  VkFormat = 42
	VK_FORMAT_R8G8B8A8_SRGB                                  VkFormat = 43
	VK_FORMAT_B8G8R8A8_UNORM                                 VkFormat = 44
	VK_FORMAT_B8G8R8A8_SNORM                                 VkFormat = 45
	VK_FORMAT_B8G8R8A8_USCALED                               VkFormat = 46
	VK_FORMAT_B8G8R8A8_SSCALED                               VkFormat = 47
	VK_FORMAT_B8G8R8A8_UINT                                  VkFormat = 48
	VK_FORMAT_B8G8R8A8_SINT                                  VkFormat = 49
	VK_FORMAT_B8G8R8A8_SRGB                                  VkFormat = 50
	VK_FORMAT_A8B8G8R8_UNORM_PACK32                          VkFormat = 51
	VK_FORMAT_A8B8G8R8_SNORM_PACK32                          VkFormat = 52
	VK_FORMAT_A8B8G8R8_USCALED_PACK32                        VkFormat = 53
	VK_FORMAT_A8B8G8R8_SSCALED_PACK32                        VkFormat = 54
	VK_FORMAT_A8B8G8R8_UINT_PACK32                           VkFormat = 55
	VK_FORMAT_A8B8G8R8_SINT_PACK32                           VkFormat = 56
	VK_FORMAT_A8B8G8R8_SRGB_PACK32                           VkFormat = 57
	VK_FORMAT_A2R10G10B10_UNORM_PACK32                       VkFormat = 58
	VK_FORMAT_A2R10G10B10_SNORM_PACK32                       VkFormat = 59
	VK_FORMAT_A2R10G10B10_USCALED_PACK32                     VkFormat = 60
	VK_FORMAT_A2R10G10B10_SSCALED_PACK32                     VkFormat = 61
	VK_FORMAT_A2R10G10B10_UINT_PACK32                        VkFormat = 62
	VK_FORMAT_A2R10G10B10_SINT_PACK32                        VkFormat = 63
	VK_FORMAT_A2B10G10R10_UNORM_PACK32                       VkFormat = 64
	VK_FORMAT_A2B10G10R10_SNORM_PACK32                       VkFormat = 65
	VK_FORMAT_A2B10G10R10_USCALED_PACK32                     VkFormat = 66
	VK_FORMAT_A2B10G10R10_SSCALED_PACK32                     VkFormat = 67
	VK_FORMAT_A2B10G10R10_UINT_PACK32                        VkFormat = 68
	VK_FORMAT_A2B10G10R10_SINT_PACK32                        VkFormat = 69
	VK_FORMAT_R16_UNORM                                      VkFormat = 70
	VK_FORMAT_R16_SNORM                                      VkFormat = 71
	VK_FORMAT_R16_USCALED                                    VkFormat = 72
	VK_FORMAT_R16_SSCALED                                    VkFormat = 73
	VK_FORMAT_R16_UINT                                       VkFormat = 74
	VK_FORMAT_R16_SINT                                       VkFormat = 75
	VK_FORMAT_R16_SFLOAT                                     VkFormat = 76
	VK_FORMAT_R16G16_UNORM                                   VkFormat = 77
	VK_FORMAT_R16G16_SNORM                                   VkFormat = 78
	VK_FORMAT_R16G16_USCALED                                 VkFormat = 79
	VK_FORMAT_R16G16_SSCALED                                 VkFormat = 80
	VK_FORMAT_R16G16_UINT                                    VkFormat = 81
	VK_FORMAT_R16G16_SINT                                    VkFormat = 82
	VK_FORMAT_R16G16_SFLOAT                                  VkFormat = 83
	VK_FORMAT_R16G16B16_UNORM                                VkFormat = 84
	VK_FORMAT_R16G16B16_SNORM                                VkFormat = 85
	VK_FORMAT_R16G16B16_USCALED                              VkFormat = 86
	VK_FORMAT_R16G16B16_SSCALED                              VkFormat = 87
	VK_FORMAT_R16G16B16_UINT                                 VkFormat = 88
	VK_FORMAT_R16G16B16_SINT                                 VkFormat = 89
	VK_FORMAT_R16G16B16_SFLOAT                               VkFormat = 90
	VK_FORMAT_R16G16B16A16_UNORM                             VkFormat = 91
	VK_FORMAT_R16G16B16A16_SNORM                             VkFormat = 92
	VK_FORMAT_R16G16B16A16_USCALED                           VkFormat = 93
	VK_FORMAT_R16G16B16A16_SSCALED                           VkFormat = 94
	VK_FORMAT_R16G16B16A16_UINT                              VkFormat = 95
	VK_FORMAT_R16G16B16A16_SINT                              VkFormat = 96
	VK_FORMAT_R16G16B16A16_SFLOAT                            VkFormat = 97
	VK_FORMAT_R32_UINT                                       VkFormat = 98
	VK_FORMAT_R32_SINT                                       VkFormat = 99
	VK_FORMAT_R32_SFLOAT                                     VkFormat = 100
	VK_FORMAT_R32G32_UINT                                    VkFormat = 101
	VK_FORMAT_R32G32_SINT                                    VkFormat = 102
	VK_FORMAT_R32G32_SFLOAT                                  VkFormat = 103
	VK_FORMAT_R32G32B32_UINT                                 VkFormat = 104
	VK_FORMAT_R32G32B32_SINT                                 VkFormat = 105
	VK_FORMAT_R32G32B32_SFLOAT                               VkFormat = 106
	VK_FORMAT_R32G32B32A32_UINT                              VkFormat = 107
	VK_FORMAT_R32G32B32A32_SINT                              VkFormat = 108
	VK_FORMAT_R32G32B32A32_SFLOAT                            VkFormat = 109
	VK_FORMAT_R64_UINT                                       VkFormat = 110
	VK_FORMAT_R64_SINT                                       VkFormat = 111
	VK_FORMAT_R64_SFLOAT                                     VkFormat = 112
	VK_FORMAT_R64G64_UINT                                    VkFormat = 113
	VK_FORMAT_R64G64_SINT                                    VkFormat = 114
	VK_FORMAT_R64G64_SFLOAT                                  VkFormat = 115
	VK_FORMAT_R64G64B64_UINT                                 VkFormat = 116
	VK_FORMAT_R64G64B64_SINT                                 VkFormat = 117
	VK_FORMAT_R64G64B64_SFLOAT                               VkFormat = 118
	VK_FORMAT_R64G64B64A64_UINT                              VkFormat = 119
	VK_FORMAT_R64G64B64A64_SINT                              VkFormat = 120
	VK_FORMAT_R64G64B64A64_SFLOAT                            VkFormat = 121
	VK_FORMAT_B10G11R11_UFLOAT_PACK32                        VkFormat = 122
	VK_FORMAT_E5B9G9R9_UFLOAT_PACK32                         VkFormat = 123
	VK_FORMAT_D16_UNORM                                      VkFormat = 124
	VK_FORMAT_X8_D24_UNORM_PACK32                            VkFormat = 125
	VK_FORMAT_D32_SFLOAT                                     VkFormat = 126
	VK_FORMAT_S8_UINT                                        VkFormat = 127
	VK_FORMAT_D16_UNORM_S8_UINT                              VkFormat = 128
	VK_FORMAT_D24_UNORM_S8_UINT                              VkFormat = 129
	VK_FORMAT_D32_SFLOAT_S8_UINT                             VkFormat = 130
	VK_FORMAT_BC1_RGB_UNORM_BLOCK                            VkFormat = 131
	VK_FORMAT_BC1_RGB_SRGB_BLOCK                             VkFormat = 132
	VK_FORMAT_BC1_RGBA_UNORM_BLOCK                           VkFormat = 133
	VK_FORMAT_BC1_RGBA_SRGB_BLOCK                            VkFormat = 134
	VK_FORMAT_BC2_UNORM_BLOCK                                VkFormat = 135
	VK_FORMAT_BC2_SRGB_BLOCK                                 VkFormat = 136
	VK_FORMAT_BC3_UNORM_BLOCK                                VkFormat = 137
	VK_FORMAT_BC3_SRGB_BLOCK                                 VkFormat = 138
	VK_FORMAT_BC4_UNORM_BLOCK                                VkFormat = 139
	VK_FORMAT_BC4_SNORM_BLOCK                                VkFormat = 140
	VK_FORMAT_BC5_UNORM_BLOCK                                VkFormat = 141
	VK_FORMAT_BC5_SNORM_BLOCK                                VkFormat = 142
	VK_FORMAT_BC6H_UFLOAT_BLOCK                              VkFormat = 143
	VK_FORMAT_BC6H_SFLOAT_BLOCK                              VkFormat = 144
	VK_FORMAT_BC7_UNORM_BLOCK                                VkFormat = 145
	VK_FORMAT_BC7_SRGB_BLOCK                                 VkFormat = 146
	VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK                        VkFormat = 147
	VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK                         VkFormat = 148
	VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK                      VkFormat = 149
	VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK                       VkFormat = 150
	VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK                      VkFormat = 151
	VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK                       VkFormat = 152
	VK_FORMAT_EAC_R11_UNORM_BLOCK                            VkFormat = 153
	VK_FORMAT_EAC_R11_SNORM_BLOCK                            VkFormat = 154
	VK_FORMAT_EAC_R11G11_UNORM_BLOCK                         VkFormat = 155
	VK_FORMAT_EAC_R11G11_SNORM_BLOCK                         VkFormat = 156
	VK_FORMAT_ASTC_4x4_UNORM_BLOCK                           VkFormat = 157
	VK_FORMAT_ASTC_4x4_SRGB_BLOCK                            VkFormat = 158
	VK_FORMAT_ASTC_5x4_UNORM_BLOCK                           VkFormat = 159
	VK_FORMAT_ASTC_5x4_SRGB_BLOCK                            VkFormat = 160
	VK_FORMAT_ASTC_5x5_UNORM_BLOCK                           VkFormat = 161
	VK_FORMAT_ASTC_5x5_SRGB_BLOCK                            VkFormat = 162
	VK_FORMAT_ASTC_6x5_UNORM_BLOCK                           VkFormat = 163
	VK_FORMAT_ASTC_6x5_SRGB_BLOCK                            VkFormat = 164
	VK_FORMAT_ASTC_6x6_UNORM_BLOCK                           VkFormat = 165
	VK_FORMAT_ASTC_6x6_SRGB_BLOCK                            VkFormat = 166
	VK_FORMAT_ASTC_8x5_UNORM_BLOCK                           VkFormat = 167
	VK_FORMAT_ASTC_8x5_SRGB_BLOCK                            VkFormat = 168
	VK_FORMAT_ASTC_8x6_UNORM_BLOCK                           VkFormat = 169
	VK_FORMAT_ASTC_8x6_SRGB_BLOCK                            VkFormat = 170
	VK_FORMAT_ASTC_8x8_UNORM_BLOCK                           VkFormat = 171
	VK_FORMAT_ASTC_8x8_SRGB_BLOCK                            VkFormat = 172
	VK_FORMAT_ASTC_10x5_UNORM_BLOCK                          VkFormat = 173
	VK_FORMAT_ASTC_10x5_SRGB_BLOCK                           VkFormat = 174
	VK_FORMAT_ASTC_10x6_UNORM_BLOCK                          VkFormat = 175
	VK_FORMAT_ASTC_10x6_SRGB_BLOCK                           VkFormat = 176
	VK_FORMAT_ASTC_10x8_UNORM_BLOCK                          VkFormat = 177
	VK_FORMAT_ASTC_10x8_SRGB_BLOCK                           VkFormat = 178
	VK_FORMAT_ASTC_10x10_UNORM_BLOCK                         VkFormat = 179
	VK_FORMAT_ASTC_10x10_SRGB_BLOCK                          VkFormat = 180
	VK_FORMAT_ASTC_12x10_UNORM_BLOCK                         VkFormat = 181
	VK_FORMAT_ASTC_12x10_SRGB_BLOCK                          VkFormat = 182
	VK_FORMAT_ASTC_12x12_UNORM_BLOCK                         VkFormat = 183
	VK_FORMAT_ASTC_12x12_SRGB_BLOCK                          VkFormat = 184
	VK_FORMAT_G8B8G8R8_422_UNORM                             VkFormat = 1000156000
	VK_FORMAT_B8G8R8G8_422_UNORM                             VkFormat = 1000156001
	VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM                      VkFormat = 1000156002
	VK_FORMAT_G8_B8R8_2PLANE_420_UNORM                       VkFormat = 1000156003
	VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM                      VkFormat = 1000156004
	VK_FORMAT_G8_B8R8_2PLANE_422_UNORM                       VkFormat = 1000156005
	VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM                      VkFormat = 1000156006
	VK_FORMAT_R10X6_UNORM_PACK16                             VkFormat = 1000156007
	VK_FORMAT_R10X6G10X6_UNORM_2PACK16                       VkFormat = 1000156008
	VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16             VkFormat = 1000156009
	VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16         VkFormat = 1000156010
	VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16         VkFormat = 1000156011
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16     VkFormat = 1000156012
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16      VkFormat = 1000156013
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16     VkFormat = 1000156014
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16      VkFormat = 1000156015
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16     VkFormat = 1000156016
	VK_FORMAT_R12X4_UNORM_PACK16                             VkFormat = 1000156017
	VK_FORMAT_R12X4G12X4_UNORM_2PACK16                       VkFormat = 1000156018
	VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16             VkFormat = 1000156019
	VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16         VkFormat = 1000156020
	VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16         VkFormat = 1000156021
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16     VkFormat = 1000156022
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16      VkFormat = 1000156023
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16     VkFormat = 1000156024
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16      VkFormat = 1000156025
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16     VkFormat = 1000156026
	VK_FORMAT_G16B16G16R16_422_UNORM                         VkFormat = 1000156027
	VK_FORMAT_B16G16R16G16_422_UNORM                         VkFormat = 1000156028
	VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM                   VkFormat = 1000156029
	VK_FORMAT_G16_B16R16_2PLANE_420_UNORM                    VkFormat = 1000156030
	VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM                   VkFormat = 1000156031
	VK_FORMAT_G16_B16R16_2PLANE_422_UNORM                    VkFormat = 1000156032
	VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM                   VkFormat = 1000156033
	VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054000
	VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054001
	VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054002
	VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG                    VkFormat = 1000054003
	VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054004
	VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054005
	VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054006
	VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG                     VkFormat = 1000054007
	VK_FORMAT_G8B8G8R8_422_UNORM_KHR                         VkFormat = VK_FORMAT_G8B8G8R8_422_UNORM
	VK_FORMAT_B8G8R8G8_422_UNORM_KHR                         VkFormat = VK_FORMAT_B8G8R8G8_422_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM
	VK_FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR                   VkFormat = VK_FORMAT_G8_B8R8_2PLANE_420_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM
	VK_FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR                   VkFormat = VK_FORMAT_G8_B8R8_2PLANE_422_UNORM
	VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR                  VkFormat = VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM
	VK_FORMAT_R10X6_UNORM_PACK16_KHR                         VkFormat = VK_FORMAT_R10X6_UNORM_PACK16
	VK_FORMAT_R10X6G10X6_UNORM_2PACK16_KHR                   VkFormat = VK_FORMAT_R10X6G10X6_UNORM_2PACK16
	VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR         VkFormat = VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	VK_FORMAT_R12X4_UNORM_PACK16_KHR                         VkFormat = VK_FORMAT_R12X4_UNORM_PACK16
	VK_FORMAT_R12X4G12X4_UNORM_2PACK16_KHR                   VkFormat = VK_FORMAT_R12X4G12X4_UNORM_2PACK16
	VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR         VkFormat = VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR     VkFormat = VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR  VkFormat = VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR VkFormat = VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	VK_FORMAT_G16B16G16R16_422_UNORM_KHR                     VkFormat = VK_FORMAT_G16B16G16R16_422_UNORM
	VK_FORMAT_B16G16R16G16_422_UNORM_KHR                     VkFormat = VK_FORMAT_B16G16R16G16_422_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM
	VK_FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR                VkFormat = VK_FORMAT_G16_B16R16_2PLANE_420_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM
	VK_FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR                VkFormat = VK_FORMAT_G16_B16R16_2PLANE_422_UNORM
	VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR               VkFormat = VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM
	VK_FORMAT_BEGIN_RANGE                                    VkFormat = VK_FORMAT_UNDEFINED
	VK_FORMAT_END_RANGE                                      VkFormat = VK_FORMAT_ASTC_12x12_SRGB_BLOCK
	VK_FORMAT_RANGE_SIZE                                     VkFormat = (VK_FORMAT_ASTC_12x12_SRGB_BLOCK - VK_FORMAT_UNDEFINED + 1)
	VK_FORMAT_MAX_ENUM                                       VkFormat = 2147483647
)

type VkFormatFeatureFlags VkFlags
type VkFormatProperties struct {
	linearTilingFeatures  VkFormatFeatureFlags
	optimalTilingFeatures VkFormatFeatureFlags
	bufferFeatures        VkFormatFeatureFlags
}

func (g *VkFormatProperties) toC(c *C.VkFormatProperties) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.linearTilingFeatures)))
			_temp = C.VkFlags(_temp)
		}
		c.linearTilingFeatures = C.VkFormatFeatureFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.optimalTilingFeatures)))
			_temp = C.VkFlags(_temp)
		}
		c.optimalTilingFeatures = C.VkFormatFeatureFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.bufferFeatures)))
			_temp = C.VkFlags(_temp)
		}
		c.bufferFeatures = C.VkFormatFeatureFlags(_temp)
	}
}
func (g *VkFormatProperties) fromC(c *C.VkFormatProperties) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.linearTilingFeatures)))
			_temp = VkFlags(_temp)
		}
		g.linearTilingFeatures = VkFormatFeatureFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.optimalTilingFeatures)))
			_temp = VkFlags(_temp)
		}
		g.optimalTilingFeatures = VkFormatFeatureFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.bufferFeatures)))
			_temp = VkFlags(_temp)
		}
		g.bufferFeatures = VkFormatFeatureFlags(_temp)
	}
}
func vkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties []VkFormatProperties) {
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
		c.pFormatProperties = (*C.VkFormatProperties)(_sa.alloc(C.sizeof_VkFormatProperties * uint(len(pFormatProperties))))
		slice3 := (*[1 << 31]C.VkFormatProperties)(unsafe.Pointer(c.pFormatProperties))[:len(pFormatProperties):len(pFormatProperties)]
		for i3, _ := range pFormatProperties {
			pFormatProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceFormatProperties(c.physicalDevice, c.format, c.pFormatProperties)
}

type VkImageType int

const (
	VK_IMAGE_TYPE_1D          VkImageType = 0
	VK_IMAGE_TYPE_2D          VkImageType = 1
	VK_IMAGE_TYPE_3D          VkImageType = 2
	VK_IMAGE_TYPE_BEGIN_RANGE VkImageType = VK_IMAGE_TYPE_1D
	VK_IMAGE_TYPE_END_RANGE   VkImageType = VK_IMAGE_TYPE_3D
	VK_IMAGE_TYPE_RANGE_SIZE  VkImageType = (VK_IMAGE_TYPE_3D - VK_IMAGE_TYPE_1D + 1)
	VK_IMAGE_TYPE_MAX_ENUM    VkImageType = 2147483647
)

type VkImageTiling int

const (
	VK_IMAGE_TILING_OPTIMAL     VkImageTiling = 0
	VK_IMAGE_TILING_LINEAR      VkImageTiling = 1
	VK_IMAGE_TILING_BEGIN_RANGE VkImageTiling = VK_IMAGE_TILING_OPTIMAL
	VK_IMAGE_TILING_END_RANGE   VkImageTiling = VK_IMAGE_TILING_LINEAR
	VK_IMAGE_TILING_RANGE_SIZE  VkImageTiling = (VK_IMAGE_TILING_LINEAR - VK_IMAGE_TILING_OPTIMAL + 1)
	VK_IMAGE_TILING_MAX_ENUM    VkImageTiling = 2147483647
)

type VkImageUsageFlags VkFlags
type VkImageCreateFlags VkFlags
type VkExtent3D struct {
	width  uint32
	height uint32
	depth  uint32
}

func (g *VkExtent3D) toC(c *C.VkExtent3D) {
	c.width = C.uint32_t(g.width)
	c.height = C.uint32_t(g.height)
	c.depth = C.uint32_t(g.depth)
}
func (g *VkExtent3D) fromC(c *C.VkExtent3D) {
	g.width = uint32(c.width)
	g.height = uint32(c.height)
	g.depth = uint32(c.depth)
}

type VkSampleCountFlags VkFlags
type VkDeviceSize uint64
type VkImageFormatProperties struct {
	maxExtent       VkExtent3D
	maxMipLevels    uint32
	maxArrayLayers  uint32
	sampleCounts    VkSampleCountFlags
	maxResourceSize VkDeviceSize
}

func (g *VkImageFormatProperties) toC(c *C.VkImageFormatProperties) {
	g.maxExtent.toC(&c.maxExtent)
	c.maxMipLevels = C.uint32_t(g.maxMipLevels)
	c.maxArrayLayers = C.uint32_t(g.maxArrayLayers)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.sampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.sampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.maxResourceSize))
		c.maxResourceSize = C.VkDeviceSize(_temp)
	}
}
func (g *VkImageFormatProperties) fromC(c *C.VkImageFormatProperties) {
	g.maxExtent.fromC(&c.maxExtent)
	g.maxMipLevels = uint32(c.maxMipLevels)
	g.maxArrayLayers = uint32(c.maxArrayLayers)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.sampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.sampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.maxResourceSize))
		g.maxResourceSize = VkDeviceSize(_temp)
	}
}
func vkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, type VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties []VkImageFormatProperties) (_ret VkResult) {
	var c struct {
		physicalDevice         C.VkPhysicalDevice
		format                 C.VkFormat
		type                   C.VkImageType
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
	c.type = C.VkImageType(type)
	c.tiling = C.VkImageTiling(tiling)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(usage)))
			_temp = C.VkFlags(_temp)
		}
		c.usage = C.VkImageUsageFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkImageCreateFlags(_temp)
	}
	{
		c.pImageFormatProperties = (*C.VkImageFormatProperties)(_sa.alloc(C.sizeof_VkImageFormatProperties * uint(len(pImageFormatProperties))))
		slice3 := (*[1 << 31]C.VkImageFormatProperties)(unsafe.Pointer(c.pImageFormatProperties))[:len(pImageFormatProperties):len(pImageFormatProperties)]
		for i3, _ := range pImageFormatProperties {
			pImageFormatProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkGetPhysicalDeviceImageFormatProperties(c.physicalDevice, c.format, c.type, c.tiling, c.usage, c.flags, c.pImageFormatProperties)
	_ret = VkResult(c._ret)
	return
}

type VkPhysicalDeviceType int

const (
	VK_PHYSICAL_DEVICE_TYPE_OTHER          VkPhysicalDeviceType = 0
	VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU VkPhysicalDeviceType = 1
	VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU   VkPhysicalDeviceType = 2
	VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU    VkPhysicalDeviceType = 3
	VK_PHYSICAL_DEVICE_TYPE_CPU            VkPhysicalDeviceType = 4
	VK_PHYSICAL_DEVICE_TYPE_BEGIN_RANGE    VkPhysicalDeviceType = VK_PHYSICAL_DEVICE_TYPE_OTHER
	VK_PHYSICAL_DEVICE_TYPE_END_RANGE      VkPhysicalDeviceType = VK_PHYSICAL_DEVICE_TYPE_CPU
	VK_PHYSICAL_DEVICE_TYPE_RANGE_SIZE     VkPhysicalDeviceType = (VK_PHYSICAL_DEVICE_TYPE_CPU - VK_PHYSICAL_DEVICE_TYPE_OTHER + 1)
	VK_PHYSICAL_DEVICE_TYPE_MAX_ENUM       VkPhysicalDeviceType = 2147483647
)

type VkPhysicalDeviceLimits struct {
	maxImageDimension1D                             uint32
	maxImageDimension2D                             uint32
	maxImageDimension3D                             uint32
	maxImageDimensionCube                           uint32
	maxImageArrayLayers                             uint32
	maxTexelBufferElements                          uint32
	maxUniformBufferRange                           uint32
	maxStorageBufferRange                           uint32
	maxPushConstantsSize                            uint32
	maxMemoryAllocationCount                        uint32
	maxSamplerAllocationCount                       uint32
	bufferImageGranularity                          VkDeviceSize
	sparseAddressSpaceSize                          VkDeviceSize
	maxBoundDescriptorSets                          uint32
	maxPerStageDescriptorSamplers                   uint32
	maxPerStageDescriptorUniformBuffers             uint32
	maxPerStageDescriptorStorageBuffers             uint32
	maxPerStageDescriptorSampledImages              uint32
	maxPerStageDescriptorStorageImages              uint32
	maxPerStageDescriptorInputAttachments           uint32
	maxPerStageResources                            uint32
	maxDescriptorSetSamplers                        uint32
	maxDescriptorSetUniformBuffers                  uint32
	maxDescriptorSetUniformBuffersDynamic           uint32
	maxDescriptorSetStorageBuffers                  uint32
	maxDescriptorSetStorageBuffersDynamic           uint32
	maxDescriptorSetSampledImages                   uint32
	maxDescriptorSetStorageImages                   uint32
	maxDescriptorSetInputAttachments                uint32
	maxVertexInputAttributes                        uint32
	maxVertexInputBindings                          uint32
	maxVertexInputAttributeOffset                   uint32
	maxVertexInputBindingStride                     uint32
	maxVertexOutputComponents                       uint32
	maxTessellationGenerationLevel                  uint32
	maxTessellationPatchSize                        uint32
	maxTessellationControlPerVertexInputComponents  uint32
	maxTessellationControlPerVertexOutputComponents uint32
	maxTessellationControlPerPatchOutputComponents  uint32
	maxTessellationControlTotalOutputComponents     uint32
	maxTessellationEvaluationInputComponents        uint32
	maxTessellationEvaluationOutputComponents       uint32
	maxGeometryShaderInvocations                    uint32
	maxGeometryInputComponents                      uint32
	maxGeometryOutputComponents                     uint32
	maxGeometryOutputVertices                       uint32
	maxGeometryTotalOutputComponents                uint32
	maxFragmentInputComponents                      uint32
	maxFragmentOutputAttachments                    uint32
	maxFragmentDualSrcAttachments                   uint32
	maxFragmentCombinedOutputResources              uint32
	maxComputeSharedMemorySize                      uint32
	maxComputeWorkGroupCount                        [3]uint32
	maxComputeWorkGroupInvocations                  uint32
	maxComputeWorkGroupSize                         [3]uint32
	subPixelPrecisionBits                           uint32
	subTexelPrecisionBits                           uint32
	mipmapPrecisionBits                             uint32
	maxDrawIndexedIndexValue                        uint32
	maxDrawIndirectCount                            uint32
	maxSamplerLodBias                               float32
	maxSamplerAnisotropy                            float32
	maxViewports                                    uint32
	maxViewportDimensions                           [2]uint32
	viewportBoundsRange                             [2]float32
	viewportSubPixelBits                            uint32
	minMemoryMapAlignment                           uint
	minTexelBufferOffsetAlignment                   VkDeviceSize
	minUniformBufferOffsetAlignment                 VkDeviceSize
	minStorageBufferOffsetAlignment                 VkDeviceSize
	minTexelOffset                                  int32
	maxTexelOffset                                  uint32
	minTexelGatherOffset                            int32
	maxTexelGatherOffset                            uint32
	minInterpolationOffset                          float32
	maxInterpolationOffset                          float32
	subPixelInterpolationOffsetBits                 uint32
	maxFramebufferWidth                             uint32
	maxFramebufferHeight                            uint32
	maxFramebufferLayers                            uint32
	framebufferColorSampleCounts                    VkSampleCountFlags
	framebufferDepthSampleCounts                    VkSampleCountFlags
	framebufferStencilSampleCounts                  VkSampleCountFlags
	framebufferNoAttachmentsSampleCounts            VkSampleCountFlags
	maxColorAttachments                             uint32
	sampledImageColorSampleCounts                   VkSampleCountFlags
	sampledImageIntegerSampleCounts                 VkSampleCountFlags
	sampledImageDepthSampleCounts                   VkSampleCountFlags
	sampledImageStencilSampleCounts                 VkSampleCountFlags
	storageImageSampleCounts                        VkSampleCountFlags
	maxSampleMaskWords                              uint32
	timestampComputeAndGraphics                     VkBool32
	timestampPeriod                                 float32
	maxClipDistances                                uint32
	maxCullDistances                                uint32
	maxCombinedClipAndCullDistances                 uint32
	discreteQueuePriorities                         uint32
	pointSizeRange                                  [2]float32
	lineWidthRange                                  [2]float32
	pointSizeGranularity                            float32
	lineWidthGranularity                            float32
	strictLines                                     VkBool32
	standardSampleLocations                         VkBool32
	optimalBufferCopyOffsetAlignment                VkDeviceSize
	optimalBufferCopyRowPitchAlignment              VkDeviceSize
	nonCoherentAtomSize                             VkDeviceSize
}

func (g *VkPhysicalDeviceLimits) toC(c *C.VkPhysicalDeviceLimits) {
	c.maxImageDimension1D = C.uint32_t(g.maxImageDimension1D)
	c.maxImageDimension2D = C.uint32_t(g.maxImageDimension2D)
	c.maxImageDimension3D = C.uint32_t(g.maxImageDimension3D)
	c.maxImageDimensionCube = C.uint32_t(g.maxImageDimensionCube)
	c.maxImageArrayLayers = C.uint32_t(g.maxImageArrayLayers)
	c.maxTexelBufferElements = C.uint32_t(g.maxTexelBufferElements)
	c.maxUniformBufferRange = C.uint32_t(g.maxUniformBufferRange)
	c.maxStorageBufferRange = C.uint32_t(g.maxStorageBufferRange)
	c.maxPushConstantsSize = C.uint32_t(g.maxPushConstantsSize)
	c.maxMemoryAllocationCount = C.uint32_t(g.maxMemoryAllocationCount)
	c.maxSamplerAllocationCount = C.uint32_t(g.maxSamplerAllocationCount)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.bufferImageGranularity))
		c.bufferImageGranularity = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.sparseAddressSpaceSize))
		c.sparseAddressSpaceSize = C.VkDeviceSize(_temp)
	}
	c.maxBoundDescriptorSets = C.uint32_t(g.maxBoundDescriptorSets)
	c.maxPerStageDescriptorSamplers = C.uint32_t(g.maxPerStageDescriptorSamplers)
	c.maxPerStageDescriptorUniformBuffers = C.uint32_t(g.maxPerStageDescriptorUniformBuffers)
	c.maxPerStageDescriptorStorageBuffers = C.uint32_t(g.maxPerStageDescriptorStorageBuffers)
	c.maxPerStageDescriptorSampledImages = C.uint32_t(g.maxPerStageDescriptorSampledImages)
	c.maxPerStageDescriptorStorageImages = C.uint32_t(g.maxPerStageDescriptorStorageImages)
	c.maxPerStageDescriptorInputAttachments = C.uint32_t(g.maxPerStageDescriptorInputAttachments)
	c.maxPerStageResources = C.uint32_t(g.maxPerStageResources)
	c.maxDescriptorSetSamplers = C.uint32_t(g.maxDescriptorSetSamplers)
	c.maxDescriptorSetUniformBuffers = C.uint32_t(g.maxDescriptorSetUniformBuffers)
	c.maxDescriptorSetUniformBuffersDynamic = C.uint32_t(g.maxDescriptorSetUniformBuffersDynamic)
	c.maxDescriptorSetStorageBuffers = C.uint32_t(g.maxDescriptorSetStorageBuffers)
	c.maxDescriptorSetStorageBuffersDynamic = C.uint32_t(g.maxDescriptorSetStorageBuffersDynamic)
	c.maxDescriptorSetSampledImages = C.uint32_t(g.maxDescriptorSetSampledImages)
	c.maxDescriptorSetStorageImages = C.uint32_t(g.maxDescriptorSetStorageImages)
	c.maxDescriptorSetInputAttachments = C.uint32_t(g.maxDescriptorSetInputAttachments)
	c.maxVertexInputAttributes = C.uint32_t(g.maxVertexInputAttributes)
	c.maxVertexInputBindings = C.uint32_t(g.maxVertexInputBindings)
	c.maxVertexInputAttributeOffset = C.uint32_t(g.maxVertexInputAttributeOffset)
	c.maxVertexInputBindingStride = C.uint32_t(g.maxVertexInputBindingStride)
	c.maxVertexOutputComponents = C.uint32_t(g.maxVertexOutputComponents)
	c.maxTessellationGenerationLevel = C.uint32_t(g.maxTessellationGenerationLevel)
	c.maxTessellationPatchSize = C.uint32_t(g.maxTessellationPatchSize)
	c.maxTessellationControlPerVertexInputComponents = C.uint32_t(g.maxTessellationControlPerVertexInputComponents)
	c.maxTessellationControlPerVertexOutputComponents = C.uint32_t(g.maxTessellationControlPerVertexOutputComponents)
	c.maxTessellationControlPerPatchOutputComponents = C.uint32_t(g.maxTessellationControlPerPatchOutputComponents)
	c.maxTessellationControlTotalOutputComponents = C.uint32_t(g.maxTessellationControlTotalOutputComponents)
	c.maxTessellationEvaluationInputComponents = C.uint32_t(g.maxTessellationEvaluationInputComponents)
	c.maxTessellationEvaluationOutputComponents = C.uint32_t(g.maxTessellationEvaluationOutputComponents)
	c.maxGeometryShaderInvocations = C.uint32_t(g.maxGeometryShaderInvocations)
	c.maxGeometryInputComponents = C.uint32_t(g.maxGeometryInputComponents)
	c.maxGeometryOutputComponents = C.uint32_t(g.maxGeometryOutputComponents)
	c.maxGeometryOutputVertices = C.uint32_t(g.maxGeometryOutputVertices)
	c.maxGeometryTotalOutputComponents = C.uint32_t(g.maxGeometryTotalOutputComponents)
	c.maxFragmentInputComponents = C.uint32_t(g.maxFragmentInputComponents)
	c.maxFragmentOutputAttachments = C.uint32_t(g.maxFragmentOutputAttachments)
	c.maxFragmentDualSrcAttachments = C.uint32_t(g.maxFragmentDualSrcAttachments)
	c.maxFragmentCombinedOutputResources = C.uint32_t(g.maxFragmentCombinedOutputResources)
	c.maxComputeSharedMemorySize = C.uint32_t(g.maxComputeSharedMemorySize)
	for i, _ := range g.maxComputeWorkGroupCount {
		c.maxComputeWorkGroupCount[i] = C.uint32_t(g.maxComputeWorkGroupCount[i])
	}
	c.maxComputeWorkGroupInvocations = C.uint32_t(g.maxComputeWorkGroupInvocations)
	for i, _ := range g.maxComputeWorkGroupSize {
		c.maxComputeWorkGroupSize[i] = C.uint32_t(g.maxComputeWorkGroupSize[i])
	}
	c.subPixelPrecisionBits = C.uint32_t(g.subPixelPrecisionBits)
	c.subTexelPrecisionBits = C.uint32_t(g.subTexelPrecisionBits)
	c.mipmapPrecisionBits = C.uint32_t(g.mipmapPrecisionBits)
	c.maxDrawIndexedIndexValue = C.uint32_t(g.maxDrawIndexedIndexValue)
	c.maxDrawIndirectCount = C.uint32_t(g.maxDrawIndirectCount)
	c.maxSamplerLodBias = C.float(g.maxSamplerLodBias)
	c.maxSamplerAnisotropy = C.float(g.maxSamplerAnisotropy)
	c.maxViewports = C.uint32_t(g.maxViewports)
	for i, _ := range g.maxViewportDimensions {
		c.maxViewportDimensions[i] = C.uint32_t(g.maxViewportDimensions[i])
	}
	for i, _ := range g.viewportBoundsRange {
		c.viewportBoundsRange[i] = C.float(g.viewportBoundsRange[i])
	}
	c.viewportSubPixelBits = C.uint32_t(g.viewportSubPixelBits)
	c.minMemoryMapAlignment = C.size_t(g.minMemoryMapAlignment)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.minTexelBufferOffsetAlignment))
		c.minTexelBufferOffsetAlignment = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.minUniformBufferOffsetAlignment))
		c.minUniformBufferOffsetAlignment = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.minStorageBufferOffsetAlignment))
		c.minStorageBufferOffsetAlignment = C.VkDeviceSize(_temp)
	}
	c.minTexelOffset = C.int32_t(g.minTexelOffset)
	c.maxTexelOffset = C.uint32_t(g.maxTexelOffset)
	c.minTexelGatherOffset = C.int32_t(g.minTexelGatherOffset)
	c.maxTexelGatherOffset = C.uint32_t(g.maxTexelGatherOffset)
	c.minInterpolationOffset = C.float(g.minInterpolationOffset)
	c.maxInterpolationOffset = C.float(g.maxInterpolationOffset)
	c.subPixelInterpolationOffsetBits = C.uint32_t(g.subPixelInterpolationOffsetBits)
	c.maxFramebufferWidth = C.uint32_t(g.maxFramebufferWidth)
	c.maxFramebufferHeight = C.uint32_t(g.maxFramebufferHeight)
	c.maxFramebufferLayers = C.uint32_t(g.maxFramebufferLayers)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.framebufferColorSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.framebufferColorSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.framebufferDepthSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.framebufferDepthSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.framebufferStencilSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.framebufferStencilSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.framebufferNoAttachmentsSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.framebufferNoAttachmentsSampleCounts = C.VkSampleCountFlags(_temp)
	}
	c.maxColorAttachments = C.uint32_t(g.maxColorAttachments)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.sampledImageColorSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.sampledImageColorSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.sampledImageIntegerSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.sampledImageIntegerSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.sampledImageDepthSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.sampledImageDepthSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.sampledImageStencilSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.sampledImageStencilSampleCounts = C.VkSampleCountFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.storageImageSampleCounts)))
			_temp = C.VkFlags(_temp)
		}
		c.storageImageSampleCounts = C.VkSampleCountFlags(_temp)
	}
	c.maxSampleMaskWords = C.uint32_t(g.maxSampleMaskWords)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.timestampComputeAndGraphics))
		c.timestampComputeAndGraphics = C.VkBool32(_temp)
	}
	c.timestampPeriod = C.float(g.timestampPeriod)
	c.maxClipDistances = C.uint32_t(g.maxClipDistances)
	c.maxCullDistances = C.uint32_t(g.maxCullDistances)
	c.maxCombinedClipAndCullDistances = C.uint32_t(g.maxCombinedClipAndCullDistances)
	c.discreteQueuePriorities = C.uint32_t(g.discreteQueuePriorities)
	for i, _ := range g.pointSizeRange {
		c.pointSizeRange[i] = C.float(g.pointSizeRange[i])
	}
	for i, _ := range g.lineWidthRange {
		c.lineWidthRange[i] = C.float(g.lineWidthRange[i])
	}
	c.pointSizeGranularity = C.float(g.pointSizeGranularity)
	c.lineWidthGranularity = C.float(g.lineWidthGranularity)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.strictLines))
		c.strictLines = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.standardSampleLocations))
		c.standardSampleLocations = C.VkBool32(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.optimalBufferCopyOffsetAlignment))
		c.optimalBufferCopyOffsetAlignment = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.optimalBufferCopyRowPitchAlignment))
		c.optimalBufferCopyRowPitchAlignment = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.nonCoherentAtomSize))
		c.nonCoherentAtomSize = C.VkDeviceSize(_temp)
	}
}
func (g *VkPhysicalDeviceLimits) fromC(c *C.VkPhysicalDeviceLimits) {
	g.maxImageDimension1D = uint32(c.maxImageDimension1D)
	g.maxImageDimension2D = uint32(c.maxImageDimension2D)
	g.maxImageDimension3D = uint32(c.maxImageDimension3D)
	g.maxImageDimensionCube = uint32(c.maxImageDimensionCube)
	g.maxImageArrayLayers = uint32(c.maxImageArrayLayers)
	g.maxTexelBufferElements = uint32(c.maxTexelBufferElements)
	g.maxUniformBufferRange = uint32(c.maxUniformBufferRange)
	g.maxStorageBufferRange = uint32(c.maxStorageBufferRange)
	g.maxPushConstantsSize = uint32(c.maxPushConstantsSize)
	g.maxMemoryAllocationCount = uint32(c.maxMemoryAllocationCount)
	g.maxSamplerAllocationCount = uint32(c.maxSamplerAllocationCount)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.bufferImageGranularity))
		g.bufferImageGranularity = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.sparseAddressSpaceSize))
		g.sparseAddressSpaceSize = VkDeviceSize(_temp)
	}
	g.maxBoundDescriptorSets = uint32(c.maxBoundDescriptorSets)
	g.maxPerStageDescriptorSamplers = uint32(c.maxPerStageDescriptorSamplers)
	g.maxPerStageDescriptorUniformBuffers = uint32(c.maxPerStageDescriptorUniformBuffers)
	g.maxPerStageDescriptorStorageBuffers = uint32(c.maxPerStageDescriptorStorageBuffers)
	g.maxPerStageDescriptorSampledImages = uint32(c.maxPerStageDescriptorSampledImages)
	g.maxPerStageDescriptorStorageImages = uint32(c.maxPerStageDescriptorStorageImages)
	g.maxPerStageDescriptorInputAttachments = uint32(c.maxPerStageDescriptorInputAttachments)
	g.maxPerStageResources = uint32(c.maxPerStageResources)
	g.maxDescriptorSetSamplers = uint32(c.maxDescriptorSetSamplers)
	g.maxDescriptorSetUniformBuffers = uint32(c.maxDescriptorSetUniformBuffers)
	g.maxDescriptorSetUniformBuffersDynamic = uint32(c.maxDescriptorSetUniformBuffersDynamic)
	g.maxDescriptorSetStorageBuffers = uint32(c.maxDescriptorSetStorageBuffers)
	g.maxDescriptorSetStorageBuffersDynamic = uint32(c.maxDescriptorSetStorageBuffersDynamic)
	g.maxDescriptorSetSampledImages = uint32(c.maxDescriptorSetSampledImages)
	g.maxDescriptorSetStorageImages = uint32(c.maxDescriptorSetStorageImages)
	g.maxDescriptorSetInputAttachments = uint32(c.maxDescriptorSetInputAttachments)
	g.maxVertexInputAttributes = uint32(c.maxVertexInputAttributes)
	g.maxVertexInputBindings = uint32(c.maxVertexInputBindings)
	g.maxVertexInputAttributeOffset = uint32(c.maxVertexInputAttributeOffset)
	g.maxVertexInputBindingStride = uint32(c.maxVertexInputBindingStride)
	g.maxVertexOutputComponents = uint32(c.maxVertexOutputComponents)
	g.maxTessellationGenerationLevel = uint32(c.maxTessellationGenerationLevel)
	g.maxTessellationPatchSize = uint32(c.maxTessellationPatchSize)
	g.maxTessellationControlPerVertexInputComponents = uint32(c.maxTessellationControlPerVertexInputComponents)
	g.maxTessellationControlPerVertexOutputComponents = uint32(c.maxTessellationControlPerVertexOutputComponents)
	g.maxTessellationControlPerPatchOutputComponents = uint32(c.maxTessellationControlPerPatchOutputComponents)
	g.maxTessellationControlTotalOutputComponents = uint32(c.maxTessellationControlTotalOutputComponents)
	g.maxTessellationEvaluationInputComponents = uint32(c.maxTessellationEvaluationInputComponents)
	g.maxTessellationEvaluationOutputComponents = uint32(c.maxTessellationEvaluationOutputComponents)
	g.maxGeometryShaderInvocations = uint32(c.maxGeometryShaderInvocations)
	g.maxGeometryInputComponents = uint32(c.maxGeometryInputComponents)
	g.maxGeometryOutputComponents = uint32(c.maxGeometryOutputComponents)
	g.maxGeometryOutputVertices = uint32(c.maxGeometryOutputVertices)
	g.maxGeometryTotalOutputComponents = uint32(c.maxGeometryTotalOutputComponents)
	g.maxFragmentInputComponents = uint32(c.maxFragmentInputComponents)
	g.maxFragmentOutputAttachments = uint32(c.maxFragmentOutputAttachments)
	g.maxFragmentDualSrcAttachments = uint32(c.maxFragmentDualSrcAttachments)
	g.maxFragmentCombinedOutputResources = uint32(c.maxFragmentCombinedOutputResources)
	g.maxComputeSharedMemorySize = uint32(c.maxComputeSharedMemorySize)
	for i, _ := range g.maxComputeWorkGroupCount {
		g.maxComputeWorkGroupCount[i] = uint32(c.maxComputeWorkGroupCount[i])
	}
	g.maxComputeWorkGroupInvocations = uint32(c.maxComputeWorkGroupInvocations)
	for i, _ := range g.maxComputeWorkGroupSize {
		g.maxComputeWorkGroupSize[i] = uint32(c.maxComputeWorkGroupSize[i])
	}
	g.subPixelPrecisionBits = uint32(c.subPixelPrecisionBits)
	g.subTexelPrecisionBits = uint32(c.subTexelPrecisionBits)
	g.mipmapPrecisionBits = uint32(c.mipmapPrecisionBits)
	g.maxDrawIndexedIndexValue = uint32(c.maxDrawIndexedIndexValue)
	g.maxDrawIndirectCount = uint32(c.maxDrawIndirectCount)
	g.maxSamplerLodBias = float32(c.maxSamplerLodBias)
	g.maxSamplerAnisotropy = float32(c.maxSamplerAnisotropy)
	g.maxViewports = uint32(c.maxViewports)
	for i, _ := range g.maxViewportDimensions {
		g.maxViewportDimensions[i] = uint32(c.maxViewportDimensions[i])
	}
	for i, _ := range g.viewportBoundsRange {
		g.viewportBoundsRange[i] = float32(c.viewportBoundsRange[i])
	}
	g.viewportSubPixelBits = uint32(c.viewportSubPixelBits)
	g.minMemoryMapAlignment = uint(c.minMemoryMapAlignment)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.minTexelBufferOffsetAlignment))
		g.minTexelBufferOffsetAlignment = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.minUniformBufferOffsetAlignment))
		g.minUniformBufferOffsetAlignment = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.minStorageBufferOffsetAlignment))
		g.minStorageBufferOffsetAlignment = VkDeviceSize(_temp)
	}
	g.minTexelOffset = int32(c.minTexelOffset)
	g.maxTexelOffset = uint32(c.maxTexelOffset)
	g.minTexelGatherOffset = int32(c.minTexelGatherOffset)
	g.maxTexelGatherOffset = uint32(c.maxTexelGatherOffset)
	g.minInterpolationOffset = float32(c.minInterpolationOffset)
	g.maxInterpolationOffset = float32(c.maxInterpolationOffset)
	g.subPixelInterpolationOffsetBits = uint32(c.subPixelInterpolationOffsetBits)
	g.maxFramebufferWidth = uint32(c.maxFramebufferWidth)
	g.maxFramebufferHeight = uint32(c.maxFramebufferHeight)
	g.maxFramebufferLayers = uint32(c.maxFramebufferLayers)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.framebufferColorSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.framebufferColorSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.framebufferDepthSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.framebufferDepthSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.framebufferStencilSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.framebufferStencilSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.framebufferNoAttachmentsSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.framebufferNoAttachmentsSampleCounts = VkSampleCountFlags(_temp)
	}
	g.maxColorAttachments = uint32(c.maxColorAttachments)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageColorSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.sampledImageColorSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageIntegerSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.sampledImageIntegerSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageDepthSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.sampledImageDepthSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.sampledImageStencilSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.sampledImageStencilSampleCounts = VkSampleCountFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.storageImageSampleCounts)))
			_temp = VkFlags(_temp)
		}
		g.storageImageSampleCounts = VkSampleCountFlags(_temp)
	}
	g.maxSampleMaskWords = uint32(c.maxSampleMaskWords)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.timestampComputeAndGraphics))
		g.timestampComputeAndGraphics = VkBool32(_temp)
	}
	g.timestampPeriod = float32(c.timestampPeriod)
	g.maxClipDistances = uint32(c.maxClipDistances)
	g.maxCullDistances = uint32(c.maxCullDistances)
	g.maxCombinedClipAndCullDistances = uint32(c.maxCombinedClipAndCullDistances)
	g.discreteQueuePriorities = uint32(c.discreteQueuePriorities)
	for i, _ := range g.pointSizeRange {
		g.pointSizeRange[i] = float32(c.pointSizeRange[i])
	}
	for i, _ := range g.lineWidthRange {
		g.lineWidthRange[i] = float32(c.lineWidthRange[i])
	}
	g.pointSizeGranularity = float32(c.pointSizeGranularity)
	g.lineWidthGranularity = float32(c.lineWidthGranularity)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.strictLines))
		g.strictLines = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.standardSampleLocations))
		g.standardSampleLocations = VkBool32(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.optimalBufferCopyOffsetAlignment))
		g.optimalBufferCopyOffsetAlignment = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.optimalBufferCopyRowPitchAlignment))
		g.optimalBufferCopyRowPitchAlignment = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.nonCoherentAtomSize))
		g.nonCoherentAtomSize = VkDeviceSize(_temp)
	}
}

type VkPhysicalDeviceSparseProperties struct {
	residencyStandard2DBlockShape            VkBool32
	residencyStandard2DMultisampleBlockShape VkBool32
	residencyStandard3DBlockShape            VkBool32
	residencyAlignedMipSize                  VkBool32
	residencyNonResidentStrict               VkBool32
}

func (g *VkPhysicalDeviceSparseProperties) toC(c *C.VkPhysicalDeviceSparseProperties) {
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.residencyStandard2DBlockShape))
		c.residencyStandard2DBlockShape = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.residencyStandard2DMultisampleBlockShape))
		c.residencyStandard2DMultisampleBlockShape = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.residencyStandard3DBlockShape))
		c.residencyStandard3DBlockShape = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.residencyAlignedMipSize))
		c.residencyAlignedMipSize = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.residencyNonResidentStrict))
		c.residencyNonResidentStrict = C.VkBool32(_temp)
	}
}
func (g *VkPhysicalDeviceSparseProperties) fromC(c *C.VkPhysicalDeviceSparseProperties) {
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.residencyStandard2DBlockShape))
		g.residencyStandard2DBlockShape = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.residencyStandard2DMultisampleBlockShape))
		g.residencyStandard2DMultisampleBlockShape = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.residencyStandard3DBlockShape))
		g.residencyStandard3DBlockShape = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.residencyAlignedMipSize))
		g.residencyAlignedMipSize = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.residencyNonResidentStrict))
		g.residencyNonResidentStrict = VkBool32(_temp)
	}
}

type VkPhysicalDeviceProperties struct {
	apiVersion        uint32
	driverVersion     uint32
	vendorID          uint32
	deviceID          uint32
	deviceType        VkPhysicalDeviceType
	deviceName        [256]byte
	pipelineCacheUUID [16]uint8
	limits            VkPhysicalDeviceLimits
	sparseProperties  VkPhysicalDeviceSparseProperties
}

func (g *VkPhysicalDeviceProperties) toC(c *C.VkPhysicalDeviceProperties) {
	c.apiVersion = C.uint32_t(g.apiVersion)
	c.driverVersion = C.uint32_t(g.driverVersion)
	c.vendorID = C.uint32_t(g.vendorID)
	c.deviceID = C.uint32_t(g.deviceID)
	c.deviceType = C.VkPhysicalDeviceType(g.deviceType)
	for i, _ := range g.deviceName {
		c.deviceName[i] = C.char(g.deviceName[i])
	}
	for i, _ := range g.pipelineCacheUUID {
		c.pipelineCacheUUID[i] = C.uint8_t(g.pipelineCacheUUID[i])
	}
	g.limits.toC(&c.limits)
	g.sparseProperties.toC(&c.sparseProperties)
}
func (g *VkPhysicalDeviceProperties) fromC(c *C.VkPhysicalDeviceProperties) {
	g.apiVersion = uint32(c.apiVersion)
	g.driverVersion = uint32(c.driverVersion)
	g.vendorID = uint32(c.vendorID)
	g.deviceID = uint32(c.deviceID)
	g.deviceType = VkPhysicalDeviceType(c.deviceType)
	for i, _ := range g.deviceName {
		g.deviceName[i] = byte(c.deviceName[i])
	}
	for i, _ := range g.pipelineCacheUUID {
		g.pipelineCacheUUID[i] = uint8(c.pipelineCacheUUID[i])
	}
	g.limits.fromC(&c.limits)
	g.sparseProperties.fromC(&c.sparseProperties)
}
func vkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties []VkPhysicalDeviceProperties) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		pProperties    *C.VkPhysicalDeviceProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pProperties = (*C.VkPhysicalDeviceProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceProperties(c.physicalDevice, c.pProperties)
}

type VkQueueFlags VkFlags
type VkQueueFamilyProperties struct {
	queueFlags                  VkQueueFlags
	queueCount                  uint32
	timestampValidBits          uint32
	minImageTransferGranularity VkExtent3D
}

func (g *VkQueueFamilyProperties) toC(c *C.VkQueueFamilyProperties) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.queueFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.queueFlags = C.VkQueueFlags(_temp)
	}
	c.queueCount = C.uint32_t(g.queueCount)
	c.timestampValidBits = C.uint32_t(g.timestampValidBits)
	g.minImageTransferGranularity.toC(&c.minImageTransferGranularity)
}
func (g *VkQueueFamilyProperties) fromC(c *C.VkQueueFamilyProperties) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.queueFlags)))
			_temp = VkFlags(_temp)
		}
		g.queueFlags = VkQueueFlags(_temp)
	}
	g.queueCount = uint32(c.queueCount)
	g.timestampValidBits = uint32(c.timestampValidBits)
	g.minImageTransferGranularity.fromC(&c.minImageTransferGranularity)
}
func vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties []VkQueueFamilyProperties) {
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
		*c.pQueueFamilyPropertyCount = C.uint32_t(*pQueueFamilyPropertyCount)
	}
	{
		c.pQueueFamilyProperties = (*C.VkQueueFamilyProperties)(_sa.alloc(C.sizeof_VkQueueFamilyProperties * uint(len(pQueueFamilyProperties))))
		slice3 := (*[1 << 31]C.VkQueueFamilyProperties)(unsafe.Pointer(c.pQueueFamilyProperties))[:len(pQueueFamilyProperties):len(pQueueFamilyProperties)]
		for i3, _ := range pQueueFamilyProperties {
			pQueueFamilyProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceQueueFamilyProperties(c.physicalDevice, c.pQueueFamilyPropertyCount, c.pQueueFamilyProperties)
	*pQueueFamilyPropertyCount = uint32(*c.pQueueFamilyPropertyCount)
}

type VkMemoryPropertyFlags VkFlags
type VkMemoryType struct {
	propertyFlags VkMemoryPropertyFlags
	heapIndex     uint32
}

func (g *VkMemoryType) toC(c *C.VkMemoryType) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.propertyFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.propertyFlags = C.VkMemoryPropertyFlags(_temp)
	}
	c.heapIndex = C.uint32_t(g.heapIndex)
}
func (g *VkMemoryType) fromC(c *C.VkMemoryType) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.propertyFlags)))
			_temp = VkFlags(_temp)
		}
		g.propertyFlags = VkMemoryPropertyFlags(_temp)
	}
	g.heapIndex = uint32(c.heapIndex)
}

type VkMemoryHeapFlags VkFlags
type VkMemoryHeap struct {
	size  VkDeviceSize
	flags VkMemoryHeapFlags
}

func (g *VkMemoryHeap) toC(c *C.VkMemoryHeap) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkMemoryHeapFlags(_temp)
	}
}
func (g *VkMemoryHeap) fromC(c *C.VkMemoryHeap) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkMemoryHeapFlags(_temp)
	}
}

type VkPhysicalDeviceMemoryProperties struct {
	memoryTypeCount uint32
	memoryTypes     [32]VkMemoryType
	memoryHeapCount uint32
	memoryHeaps     [16]VkMemoryHeap
}

func (g *VkPhysicalDeviceMemoryProperties) toC(c *C.VkPhysicalDeviceMemoryProperties) {
	c.memoryTypeCount = C.uint32_t(g.memoryTypeCount)
	for i, _ := range g.memoryTypes {
		g.memoryTypes[i].toC(&c.memoryTypes[i])
	}
	c.memoryHeapCount = C.uint32_t(g.memoryHeapCount)
	for i, _ := range g.memoryHeaps {
		g.memoryHeaps[i].toC(&c.memoryHeaps[i])
	}
}
func (g *VkPhysicalDeviceMemoryProperties) fromC(c *C.VkPhysicalDeviceMemoryProperties) {
	g.memoryTypeCount = uint32(c.memoryTypeCount)
	for i, _ := range g.memoryTypes {
		g.memoryTypes[i].fromC(&c.memoryTypes[i])
	}
	g.memoryHeapCount = uint32(c.memoryHeapCount)
	for i, _ := range g.memoryHeaps {
		g.memoryHeaps[i].fromC(&c.memoryHeaps[i])
	}
}
func vkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties []VkPhysicalDeviceMemoryProperties) {
	var c struct {
		physicalDevice    C.VkPhysicalDevice
		pMemoryProperties *C.VkPhysicalDeviceMemoryProperties
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.physicalDevice = C.VkPhysicalDevice(physicalDevice)
	{
		c.pMemoryProperties = (*C.VkPhysicalDeviceMemoryProperties)(_sa.alloc(C.sizeof_VkPhysicalDeviceMemoryProperties * uint(len(pMemoryProperties))))
		slice3 := (*[1 << 31]C.VkPhysicalDeviceMemoryProperties)(unsafe.Pointer(c.pMemoryProperties))[:len(pMemoryProperties):len(pMemoryProperties)]
		for i3, _ := range pMemoryProperties {
			pMemoryProperties[i3].toC(&slice3[i3])
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
func vkGetInstanceProcAddr(instance VkInstance, pName string) (_ret PFN_vkVoidFunction) {
	var c struct {
		instance C.VkInstance
		pName    *C.char
		_ret     C.PFN_vkVoidFunction
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.instance = C.VkInstance(instance)
	c.pName = toCString(pName, _sa)
	c._ret = C.vkGetInstanceProcAddr(c.instance, c.pName)
	_ret.Raw = c._ret
	return
}

type VkDevice C.VkDevice

func vkGetDeviceProcAddr(device VkDevice, pName string) (_ret PFN_vkVoidFunction) {
	var c struct {
		device C.VkDevice
		pName  *C.char
		_ret   C.PFN_vkVoidFunction
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.pName = toCString(pName, _sa)
	c._ret = C.vkGetDeviceProcAddr(c.device, c.pName)
	_ret.Raw = c._ret
	return
}

type VkDeviceCreateFlags VkFlags
type VkDeviceQueueCreateFlags VkFlags
type VkDeviceQueueCreateInfo struct {
	sType            VkStructureType
	pNext            *[0]byte
	flags            VkDeviceQueueCreateFlags
	queueFamilyIndex uint32
	pQueuePriorities []float32
}

func (g *VkDeviceQueueCreateInfo) toC(c *C.VkDeviceQueueCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkDeviceQueueCreateFlags(_temp)
	}
	c.queueFamilyIndex = C.uint32_t(g.queueFamilyIndex)
	c.queueCount = C.uint32_t(len(g.pQueuePriorities))
	{
		c.pQueuePriorities = (*C.float)(_sa.alloc(C.sizeof_float * uint(len(g.pQueuePriorities))))
		slice2 := (*[1 << 31]C.float)(unsafe.Pointer(c.pQueuePriorities))[:len(g.pQueuePriorities):len(g.pQueuePriorities)]
		for i2, _ := range g.pQueuePriorities {
			slice2[i2] = C.float(g.pQueuePriorities[i2])
		}
	}
}
func (g *VkDeviceQueueCreateInfo) fromC(c *C.VkDeviceQueueCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkDeviceQueueCreateFlags(_temp)
	}
	g.queueFamilyIndex = uint32(c.queueFamilyIndex)
	g.pQueuePriorities = make([]float32, int(c.queueCount))
	{
		slice2 := (*[1 << 31]C.float)(unsafe.Pointer(c.pQueuePriorities))[:len(g.pQueuePriorities):len(g.pQueuePriorities)]
		for i2, _ := range g.pQueuePriorities {
			g.pQueuePriorities[i2] = float32(slice2[i2])
		}
	}
}

type VkDeviceCreateInfo struct {
	sType                   VkStructureType
	pNext                   *[0]byte
	flags                   VkDeviceCreateFlags
	pQueueCreateInfos       []VkDeviceQueueCreateInfo
	ppEnabledLayerNames     []string
	ppEnabledExtensionNames []string
	pEnabledFeatures        []VkPhysicalDeviceFeatures
}

func (g *VkDeviceCreateInfo) toC(c *C.VkDeviceCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkDeviceCreateFlags(_temp)
	}
	c.queueCreateInfoCount = C.uint32_t(len(g.pQueueCreateInfos))
	{
		c.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(_sa.alloc(C.sizeof_VkDeviceQueueCreateInfo * uint(len(g.pQueueCreateInfos))))
		slice2 := (*[1 << 31]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(c.pQueueCreateInfos))[:len(g.pQueueCreateInfos):len(g.pQueueCreateInfos)]
		for i2, _ := range g.pQueueCreateInfos {
			g.pQueueCreateInfos[i2].toC(&slice2[i2], _sa)
		}
	}
	c.enabledLayerCount = C.uint32_t(len(g.ppEnabledLayerNames))
	{
		c.ppEnabledLayerNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.ppEnabledLayerNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.ppEnabledLayerNames):len(g.ppEnabledLayerNames)]
		for i2, _ := range g.ppEnabledLayerNames {
			slice2[i2] = toCString(g.ppEnabledLayerNames[i2], _sa)
		}
	}
	c.enabledExtensionCount = C.uint32_t(len(g.ppEnabledExtensionNames))
	{
		c.ppEnabledExtensionNames = (**C.char)(_sa.alloc(C.sizeof_void_pointer * uint(len(g.ppEnabledExtensionNames))))
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.ppEnabledExtensionNames):len(g.ppEnabledExtensionNames)]
		for i2, _ := range g.ppEnabledExtensionNames {
			slice2[i2] = toCString(g.ppEnabledExtensionNames[i2], _sa)
		}
	}
	{
		c.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(_sa.alloc(C.sizeof_VkPhysicalDeviceFeatures * uint(len(g.pEnabledFeatures))))
		slice2 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pEnabledFeatures))[:len(g.pEnabledFeatures):len(g.pEnabledFeatures)]
		for i2, _ := range g.pEnabledFeatures {
			g.pEnabledFeatures[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkDeviceCreateInfo) fromC(c *C.VkDeviceCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkDeviceCreateFlags(_temp)
	}
	g.pQueueCreateInfos = make([]VkDeviceQueueCreateInfo, int(c.queueCreateInfoCount))
	{
		slice2 := (*[1 << 31]C.VkDeviceQueueCreateInfo)(unsafe.Pointer(c.pQueueCreateInfos))[:len(g.pQueueCreateInfos):len(g.pQueueCreateInfos)]
		for i2, _ := range g.pQueueCreateInfos {
			g.pQueueCreateInfos[i2].fromC(&slice2[i2])
		}
	}
	g.ppEnabledLayerNames = make([]string, int(c.enabledLayerCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledLayerNames))[:len(g.ppEnabledLayerNames):len(g.ppEnabledLayerNames)]
		for i2, _ := range g.ppEnabledLayerNames {
			g.ppEnabledLayerNames[i2] = toGoString(slice2[i2])
		}
	}
	g.ppEnabledExtensionNames = make([]string, int(c.enabledExtensionCount))
	{
		slice2 := (*[1 << 31]*C.char)(unsafe.Pointer(c.ppEnabledExtensionNames))[:len(g.ppEnabledExtensionNames):len(g.ppEnabledExtensionNames)]
		for i2, _ := range g.ppEnabledExtensionNames {
			g.ppEnabledExtensionNames[i2] = toGoString(slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkPhysicalDeviceFeatures)(unsafe.Pointer(c.pEnabledFeatures))[:len(g.pEnabledFeatures):len(g.pEnabledFeatures)]
		for i2, _ := range g.pEnabledFeatures {
			g.pEnabledFeatures[i2].fromC(&slice2[i2])
		}
	}
}
func vkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDevice = (*C.VkDevice)(_sa.alloc(C.sizeof_VkDevice))
		*c.pDevice = C.VkDevice(*pDevice)
	}
	c._ret = C.vkCreateDevice(c.physicalDevice, c.pCreateInfo, c.pAllocator, c.pDevice)
	_ret = VkResult(c._ret)
	*pDevice = VkDevice(*c.pDevice)
	return
}
func vkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) {
	var c struct {
		device     C.VkDevice
		pAllocator *C.VkAllocationCallbacks
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDevice(c.device, c.pAllocator)
}

type VkExtensionProperties struct {
	extensionName [256]byte
	specVersion   uint32
}

func (g *VkExtensionProperties) toC(c *C.VkExtensionProperties) {
	for i, _ := range g.extensionName {
		c.extensionName[i] = C.char(g.extensionName[i])
	}
	c.specVersion = C.uint32_t(g.specVersion)
}
func (g *VkExtensionProperties) fromC(c *C.VkExtensionProperties) {
	for i, _ := range g.extensionName {
		g.extensionName[i] = byte(c.extensionName[i])
	}
	g.specVersion = uint32(c.specVersion)
}
func vkEnumerateInstanceExtensionProperties(pLayerName string, pPropertyCount *uint32, pProperties []VkExtensionProperties) (_ret VkResult) {
	var c struct {
		pLayerName     *C.char
		pPropertyCount *C.uint32_t
		pProperties    *C.VkExtensionProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.pLayerName = toCString(pLayerName, _sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*pPropertyCount)
	}
	{
		c.pProperties = (*C.VkExtensionProperties)(_sa.alloc(C.sizeof_VkExtensionProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkExtensionProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateInstanceExtensionProperties(c.pLayerName, c.pPropertyCount, c.pProperties)
	_ret = VkResult(c._ret)
	*pPropertyCount = uint32(*c.pPropertyCount)
	return
}
func vkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName string, pPropertyCount *uint32, pProperties []VkExtensionProperties) (_ret VkResult) {
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
	c.pLayerName = toCString(pLayerName, _sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*pPropertyCount)
	}
	{
		c.pProperties = (*C.VkExtensionProperties)(_sa.alloc(C.sizeof_VkExtensionProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkExtensionProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateDeviceExtensionProperties(c.physicalDevice, c.pLayerName, c.pPropertyCount, c.pProperties)
	_ret = VkResult(c._ret)
	*pPropertyCount = uint32(*c.pPropertyCount)
	return
}

type VkLayerProperties struct {
	layerName             [256]byte
	specVersion           uint32
	implementationVersion uint32
	description           [256]byte
}

func (g *VkLayerProperties) toC(c *C.VkLayerProperties) {
	for i, _ := range g.layerName {
		c.layerName[i] = C.char(g.layerName[i])
	}
	c.specVersion = C.uint32_t(g.specVersion)
	c.implementationVersion = C.uint32_t(g.implementationVersion)
	for i, _ := range g.description {
		c.description[i] = C.char(g.description[i])
	}
}
func (g *VkLayerProperties) fromC(c *C.VkLayerProperties) {
	for i, _ := range g.layerName {
		g.layerName[i] = byte(c.layerName[i])
	}
	g.specVersion = uint32(c.specVersion)
	g.implementationVersion = uint32(c.implementationVersion)
	for i, _ := range g.description {
		g.description[i] = byte(c.description[i])
	}
}
func vkEnumerateInstanceLayerProperties(pPropertyCount *uint32, pProperties []VkLayerProperties) (_ret VkResult) {
	var c struct {
		pPropertyCount *C.uint32_t
		pProperties    *C.VkLayerProperties
		_ret           C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*pPropertyCount)
	}
	{
		c.pProperties = (*C.VkLayerProperties)(_sa.alloc(C.sizeof_VkLayerProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkLayerProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateInstanceLayerProperties(c.pPropertyCount, c.pProperties)
	_ret = VkResult(c._ret)
	*pPropertyCount = uint32(*c.pPropertyCount)
	return
}
func vkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *uint32, pProperties []VkLayerProperties) (_ret VkResult) {
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
		*c.pPropertyCount = C.uint32_t(*pPropertyCount)
	}
	{
		c.pProperties = (*C.VkLayerProperties)(_sa.alloc(C.sizeof_VkLayerProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkLayerProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkEnumerateDeviceLayerProperties(c.physicalDevice, c.pPropertyCount, c.pProperties)
	_ret = VkResult(c._ret)
	*pPropertyCount = uint32(*c.pPropertyCount)
	return
}

type VkQueue C.VkQueue

func vkGetDeviceQueue(device VkDevice, queueFamilyIndex uint32, queueIndex uint32, pQueue *VkQueue) {
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
		*c.pQueue = C.VkQueue(*pQueue)
	}
	C.vkGetDeviceQueue(c.device, c.queueFamilyIndex, c.queueIndex, c.pQueue)
	*pQueue = VkQueue(*c.pQueue)
}

type VkSemaphore C.VkSemaphore
type VkPipelineStageFlags VkFlags
type VkCommandBuffer C.VkCommandBuffer
type VkSubmitInfo struct {
	sType             VkStructureType
	pNext             *[0]byte
	pWaitSemaphores   []VkSemaphore
	pWaitDstStageMask *VkPipelineStageFlags
	pCommandBuffers   []VkCommandBuffer
	pSignalSemaphores []VkSemaphore
}

func (g *VkSubmitInfo) toC(c *C.VkSubmitInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.waitSemaphoreCount = C.uint32_t(len(g.pWaitSemaphores))
	{
		c.pWaitSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.pWaitSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.pWaitSemaphores):len(g.pWaitSemaphores)]
		for i2, _ := range g.pWaitSemaphores {
			slice2[i2] = C.VkSemaphore(g.pWaitSemaphores[i2])
		}
	}
	{
		c.pWaitDstStageMask = (*C.VkPipelineStageFlags)(_sa.alloc(C.sizeof_VkPipelineStageFlags))
		{
			var _temp C.VkFlags
			{
				var _temp C.uint32_t
				_temp = C.uint32_t((uint32)((VkFlags)(*g.pWaitDstStageMask)))
				_temp = C.VkFlags(_temp)
			}
			*c.pWaitDstStageMask = C.VkPipelineStageFlags(_temp)
		}
	}
	c.commandBufferCount = C.uint32_t(len(g.pCommandBuffers))
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(g.pCommandBuffers))))
		slice2 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(g.pCommandBuffers):len(g.pCommandBuffers)]
		for i2, _ := range g.pCommandBuffers {
			slice2[i2] = C.VkCommandBuffer(g.pCommandBuffers[i2])
		}
	}
	c.signalSemaphoreCount = C.uint32_t(len(g.pSignalSemaphores))
	{
		c.pSignalSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.pSignalSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.pSignalSemaphores):len(g.pSignalSemaphores)]
		for i2, _ := range g.pSignalSemaphores {
			slice2[i2] = C.VkSemaphore(g.pSignalSemaphores[i2])
		}
	}
}
func (g *VkSubmitInfo) fromC(c *C.VkSubmitInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.pWaitSemaphores = make([]VkSemaphore, int(c.waitSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.pWaitSemaphores):len(g.pWaitSemaphores)]
		for i2, _ := range g.pWaitSemaphores {
			g.pWaitSemaphores[i2] = VkSemaphore(slice2[i2])
		}
	}
	{
		if g.pWaitDstStageMask == nil {
			g.pWaitDstStageMask = new(VkPipelineStageFlags)
		}
		{
			var _temp VkFlags
			{
				var _temp uint32
				_temp = uint32((C.uint32_t)((C.VkFlags)(*c.pWaitDstStageMask)))
				_temp = VkFlags(_temp)
			}
			*g.pWaitDstStageMask = VkPipelineStageFlags(_temp)
		}
	}
	g.pCommandBuffers = make([]VkCommandBuffer, int(c.commandBufferCount))
	{
		slice2 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(g.pCommandBuffers):len(g.pCommandBuffers)]
		for i2, _ := range g.pCommandBuffers {
			g.pCommandBuffers[i2] = VkCommandBuffer(slice2[i2])
		}
	}
	g.pSignalSemaphores = make([]VkSemaphore, int(c.signalSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.pSignalSemaphores):len(g.pSignalSemaphores)]
		for i2, _ := range g.pSignalSemaphores {
			g.pSignalSemaphores[i2] = VkSemaphore(slice2[i2])
		}
	}
}

type VkFence C.VkFence

func vkQueueSubmit(queue VkQueue, pSubmits []VkSubmitInfo, fence VkFence) (_ret VkResult) {
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
	c.submitCount = C.uint32_t(len(pSubmits))
	{
		c.pSubmits = (*C.VkSubmitInfo)(_sa.alloc(C.sizeof_VkSubmitInfo * uint(len(pSubmits))))
		slice3 := (*[1 << 31]C.VkSubmitInfo)(unsafe.Pointer(c.pSubmits))[:len(pSubmits):len(pSubmits)]
		for i3, _ := range pSubmits {
			pSubmits[i3].toC(&slice3[i3], _sa)
		}
	}
	c.fence = C.VkFence(fence)
	c._ret = C.vkQueueSubmit(c.queue, c.submitCount, c.pSubmits, c.fence)
	_ret = VkResult(c._ret)
	return
}
func vkQueueWaitIdle(queue VkQueue) (_ret VkResult) {
	var c struct {
		queue C.VkQueue
		_ret  C.VkResult
	}
	c.queue = C.VkQueue(queue)
	c._ret = C.vkQueueWaitIdle(c.queue)
	_ret = VkResult(c._ret)
	return
}
func vkDeviceWaitIdle(device VkDevice) (_ret VkResult) {
	var c struct {
		device C.VkDevice
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c._ret = C.vkDeviceWaitIdle(c.device)
	_ret = VkResult(c._ret)
	return
}

type VkMemoryAllocateInfo struct {
	sType           VkStructureType
	pNext           *[0]byte
	allocationSize  VkDeviceSize
	memoryTypeIndex uint32
}

func (g *VkMemoryAllocateInfo) toC(c *C.VkMemoryAllocateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.allocationSize))
		c.allocationSize = C.VkDeviceSize(_temp)
	}
	c.memoryTypeIndex = C.uint32_t(g.memoryTypeIndex)
}
func (g *VkMemoryAllocateInfo) fromC(c *C.VkMemoryAllocateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.allocationSize))
		g.allocationSize = VkDeviceSize(_temp)
	}
	g.memoryTypeIndex = uint32(c.memoryTypeIndex)
}

type VkDeviceMemory C.VkDeviceMemory

func vkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (_ret VkResult) {
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
		pAllocateInfo.toC(c.pAllocateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pMemory = (*C.VkDeviceMemory)(_sa.alloc(C.sizeof_VkDeviceMemory))
		*c.pMemory = C.VkDeviceMemory(*pMemory)
	}
	c._ret = C.vkAllocateMemory(c.device, c.pAllocateInfo, c.pAllocator, c.pMemory)
	_ret = VkResult(c._ret)
	*pMemory = VkDeviceMemory(*c.pMemory)
	return
}
func vkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkFreeMemory(c.device, c.memory, c.pAllocator)
}

type VkMemoryMapFlags VkFlags

func vkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData []unsafe.Pointer) (_ret VkResult) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(size))
		c.size = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkMemoryMapFlags(_temp)
	}
	{
		c.ppData = (*unsafe.Pointer)(_sa.alloc(C.sizeof_void_pointer * uint(len(ppData))))
		slice3 := (*[1 << 31]unsafe.Pointer)(unsafe.Pointer(c.ppData))[:len(ppData):len(ppData)]
		for i3, _ := range ppData {
			slice3[i3] = ppData[i3]
		}
	}
	c._ret = C.vkMapMemory(c.device, c.memory, c.offset, c.size, c.flags, c.ppData)
	_ret = VkResult(c._ret)
	return
}
func vkUnmapMemory(device VkDevice, memory VkDeviceMemory) {
	var c struct {
		device C.VkDevice
		memory C.VkDeviceMemory
	}
	c.device = C.VkDevice(device)
	c.memory = C.VkDeviceMemory(memory)
	C.vkUnmapMemory(c.device, c.memory)
}

type VkMappedMemoryRange struct {
	sType  VkStructureType
	pNext  *[0]byte
	memory VkDeviceMemory
	offset VkDeviceSize
	size   VkDeviceSize
}

func (g *VkMappedMemoryRange) toC(c *C.VkMappedMemoryRange) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.memory = C.VkDeviceMemory(g.memory)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
}
func (g *VkMappedMemoryRange) fromC(c *C.VkMappedMemoryRange) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.memory = VkDeviceMemory(c.memory)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.offset))
		g.offset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
}
func vkFlushMappedMemoryRanges(device VkDevice, pMemoryRanges []VkMappedMemoryRange) (_ret VkResult) {
	var c struct {
		device           C.VkDevice
		memoryRangeCount C.uint32_t
		pMemoryRanges    *C.VkMappedMemoryRange
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memoryRangeCount = C.uint32_t(len(pMemoryRanges))
	{
		c.pMemoryRanges = (*C.VkMappedMemoryRange)(_sa.alloc(C.sizeof_VkMappedMemoryRange * uint(len(pMemoryRanges))))
		slice3 := (*[1 << 31]C.VkMappedMemoryRange)(unsafe.Pointer(c.pMemoryRanges))[:len(pMemoryRanges):len(pMemoryRanges)]
		for i3, _ := range pMemoryRanges {
			pMemoryRanges[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkFlushMappedMemoryRanges(c.device, c.memoryRangeCount, c.pMemoryRanges)
	_ret = VkResult(c._ret)
	return
}
func vkInvalidateMappedMemoryRanges(device VkDevice, pMemoryRanges []VkMappedMemoryRange) (_ret VkResult) {
	var c struct {
		device           C.VkDevice
		memoryRangeCount C.uint32_t
		pMemoryRanges    *C.VkMappedMemoryRange
		_ret             C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.memoryRangeCount = C.uint32_t(len(pMemoryRanges))
	{
		c.pMemoryRanges = (*C.VkMappedMemoryRange)(_sa.alloc(C.sizeof_VkMappedMemoryRange * uint(len(pMemoryRanges))))
		slice3 := (*[1 << 31]C.VkMappedMemoryRange)(unsafe.Pointer(c.pMemoryRanges))[:len(pMemoryRanges):len(pMemoryRanges)]
		for i3, _ := range pMemoryRanges {
			pMemoryRanges[i3].toC(&slice3[i3])
		}
	}
	c._ret = C.vkInvalidateMappedMemoryRanges(c.device, c.memoryRangeCount, c.pMemoryRanges)
	_ret = VkResult(c._ret)
	return
}
func vkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes []VkDeviceSize) {
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
		c.pCommittedMemoryInBytes = (*C.VkDeviceSize)(_sa.alloc(C.sizeof_VkDeviceSize * uint(len(pCommittedMemoryInBytes))))
		slice3 := (*[1 << 31]C.VkDeviceSize)(unsafe.Pointer(c.pCommittedMemoryInBytes))[:len(pCommittedMemoryInBytes):len(pCommittedMemoryInBytes)]
		for i3, _ := range pCommittedMemoryInBytes {
			{
				var _temp C.uint64_t
				_temp = C.uint64_t((uint64)(pCommittedMemoryInBytes[i3]))
				slice3[i3] = C.VkDeviceSize(_temp)
			}
		}
	}
	C.vkGetDeviceMemoryCommitment(c.device, c.memory, c.pCommittedMemoryInBytes)
}

type VkBuffer C.VkBuffer

func vkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (_ret VkResult) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(memoryOffset))
		c.memoryOffset = C.VkDeviceSize(_temp)
	}
	c._ret = C.vkBindBufferMemory(c.device, c.buffer, c.memory, c.memoryOffset)
	_ret = VkResult(c._ret)
	return
}

type VkImage C.VkImage

func vkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (_ret VkResult) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(memoryOffset))
		c.memoryOffset = C.VkDeviceSize(_temp)
	}
	c._ret = C.vkBindImageMemory(c.device, c.image, c.memory, c.memoryOffset)
	_ret = VkResult(c._ret)
	return
}

type VkMemoryRequirements struct {
	size           VkDeviceSize
	alignment      VkDeviceSize
	memoryTypeBits uint32
}

func (g *VkMemoryRequirements) toC(c *C.VkMemoryRequirements) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.alignment))
		c.alignment = C.VkDeviceSize(_temp)
	}
	c.memoryTypeBits = C.uint32_t(g.memoryTypeBits)
}
func (g *VkMemoryRequirements) fromC(c *C.VkMemoryRequirements) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.alignment))
		g.alignment = VkDeviceSize(_temp)
	}
	g.memoryTypeBits = uint32(c.memoryTypeBits)
}
func vkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements []VkMemoryRequirements) {
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
		c.pMemoryRequirements = (*C.VkMemoryRequirements)(_sa.alloc(C.sizeof_VkMemoryRequirements * uint(len(pMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements)(unsafe.Pointer(c.pMemoryRequirements))[:len(pMemoryRequirements):len(pMemoryRequirements)]
		for i3, _ := range pMemoryRequirements {
			pMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetBufferMemoryRequirements(c.device, c.buffer, c.pMemoryRequirements)
}
func vkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements []VkMemoryRequirements) {
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
		c.pMemoryRequirements = (*C.VkMemoryRequirements)(_sa.alloc(C.sizeof_VkMemoryRequirements * uint(len(pMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkMemoryRequirements)(unsafe.Pointer(c.pMemoryRequirements))[:len(pMemoryRequirements):len(pMemoryRequirements)]
		for i3, _ := range pMemoryRequirements {
			pMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageMemoryRequirements(c.device, c.image, c.pMemoryRequirements)
}

type VkImageAspectFlags VkFlags
type VkSparseImageFormatFlags VkFlags
type VkSparseImageFormatProperties struct {
	aspectMask       VkImageAspectFlags
	imageGranularity VkExtent3D
	flags            VkSparseImageFormatFlags
}

func (g *VkSparseImageFormatProperties) toC(c *C.VkSparseImageFormatProperties) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.aspectMask)))
			_temp = C.VkFlags(_temp)
		}
		c.aspectMask = C.VkImageAspectFlags(_temp)
	}
	g.imageGranularity.toC(&c.imageGranularity)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSparseImageFormatFlags(_temp)
	}
}
func (g *VkSparseImageFormatProperties) fromC(c *C.VkSparseImageFormatProperties) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			_temp = VkFlags(_temp)
		}
		g.aspectMask = VkImageAspectFlags(_temp)
	}
	g.imageGranularity.fromC(&c.imageGranularity)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSparseImageFormatFlags(_temp)
	}
}

type VkSparseImageMemoryRequirements struct {
	formatProperties     VkSparseImageFormatProperties
	imageMipTailFirstLod uint32
	imageMipTailSize     VkDeviceSize
	imageMipTailOffset   VkDeviceSize
	imageMipTailStride   VkDeviceSize
}

func (g *VkSparseImageMemoryRequirements) toC(c *C.VkSparseImageMemoryRequirements) {
	g.formatProperties.toC(&c.formatProperties)
	c.imageMipTailFirstLod = C.uint32_t(g.imageMipTailFirstLod)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.imageMipTailSize))
		c.imageMipTailSize = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.imageMipTailOffset))
		c.imageMipTailOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.imageMipTailStride))
		c.imageMipTailStride = C.VkDeviceSize(_temp)
	}
}
func (g *VkSparseImageMemoryRequirements) fromC(c *C.VkSparseImageMemoryRequirements) {
	g.formatProperties.fromC(&c.formatProperties)
	g.imageMipTailFirstLod = uint32(c.imageMipTailFirstLod)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.imageMipTailSize))
		g.imageMipTailSize = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.imageMipTailOffset))
		g.imageMipTailOffset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.imageMipTailStride))
		g.imageMipTailStride = VkDeviceSize(_temp)
	}
}
func vkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements []VkSparseImageMemoryRequirements) {
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
		*c.pSparseMemoryRequirementCount = C.uint32_t(*pSparseMemoryRequirementCount)
	}
	{
		c.pSparseMemoryRequirements = (*C.VkSparseImageMemoryRequirements)(_sa.alloc(C.sizeof_VkSparseImageMemoryRequirements * uint(len(pSparseMemoryRequirements))))
		slice3 := (*[1 << 31]C.VkSparseImageMemoryRequirements)(unsafe.Pointer(c.pSparseMemoryRequirements))[:len(pSparseMemoryRequirements):len(pSparseMemoryRequirements)]
		for i3, _ := range pSparseMemoryRequirements {
			pSparseMemoryRequirements[i3].toC(&slice3[i3])
		}
	}
	C.vkGetImageSparseMemoryRequirements(c.device, c.image, c.pSparseMemoryRequirementCount, c.pSparseMemoryRequirements)
	*pSparseMemoryRequirementCount = uint32(*c.pSparseMemoryRequirementCount)
}

type VkSampleCountFlagBits int

const (
	VK_SAMPLE_COUNT_1_BIT              VkSampleCountFlagBits = 1
	VK_SAMPLE_COUNT_2_BIT              VkSampleCountFlagBits = 2
	VK_SAMPLE_COUNT_4_BIT              VkSampleCountFlagBits = 4
	VK_SAMPLE_COUNT_8_BIT              VkSampleCountFlagBits = 8
	VK_SAMPLE_COUNT_16_BIT             VkSampleCountFlagBits = 16
	VK_SAMPLE_COUNT_32_BIT             VkSampleCountFlagBits = 32
	VK_SAMPLE_COUNT_64_BIT             VkSampleCountFlagBits = 64
	VK_SAMPLE_COUNT_FLAG_BITS_MAX_ENUM VkSampleCountFlagBits = 2147483647
)

func vkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, type VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *uint32, pProperties []VkSparseImageFormatProperties) {
	var c struct {
		physicalDevice C.VkPhysicalDevice
		format         C.VkFormat
		type           C.VkImageType
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
	c.type = C.VkImageType(type)
	c.samples = C.VkSampleCountFlagBits(samples)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(usage)))
			_temp = C.VkFlags(_temp)
		}
		c.usage = C.VkImageUsageFlags(_temp)
	}
	c.tiling = C.VkImageTiling(tiling)
	{
		c.pPropertyCount = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pPropertyCount = C.uint32_t(*pPropertyCount)
	}
	{
		c.pProperties = (*C.VkSparseImageFormatProperties)(_sa.alloc(C.sizeof_VkSparseImageFormatProperties * uint(len(pProperties))))
		slice3 := (*[1 << 31]C.VkSparseImageFormatProperties)(unsafe.Pointer(c.pProperties))[:len(pProperties):len(pProperties)]
		for i3, _ := range pProperties {
			pProperties[i3].toC(&slice3[i3])
		}
	}
	C.vkGetPhysicalDeviceSparseImageFormatProperties(c.physicalDevice, c.format, c.type, c.samples, c.usage, c.tiling, c.pPropertyCount, c.pProperties)
	*pPropertyCount = uint32(*c.pPropertyCount)
}

type VkSparseMemoryBindFlags VkFlags
type VkSparseMemoryBind struct {
	resourceOffset VkDeviceSize
	size           VkDeviceSize
	memory         VkDeviceMemory
	memoryOffset   VkDeviceSize
	flags          VkSparseMemoryBindFlags
}

func (g *VkSparseMemoryBind) toC(c *C.VkSparseMemoryBind) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.resourceOffset))
		c.resourceOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
	c.memory = C.VkDeviceMemory(g.memory)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.memoryOffset))
		c.memoryOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSparseMemoryBindFlags(_temp)
	}
}
func (g *VkSparseMemoryBind) fromC(c *C.VkSparseMemoryBind) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.resourceOffset))
		g.resourceOffset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
	g.memory = VkDeviceMemory(c.memory)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.memoryOffset))
		g.memoryOffset = VkDeviceSize(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSparseMemoryBindFlags(_temp)
	}
}

type VkSparseBufferMemoryBindInfo struct {
	buffer VkBuffer
	pBinds []VkSparseMemoryBind
}

func (g *VkSparseBufferMemoryBindInfo) toC(c *C.VkSparseBufferMemoryBindInfo, _sa *stackAllocator) {
	c.buffer = C.VkBuffer(g.buffer)
	c.bindCount = C.uint32_t(len(g.pBinds))
	{
		c.pBinds = (*C.VkSparseMemoryBind)(_sa.alloc(C.sizeof_VkSparseMemoryBind * uint(len(g.pBinds))))
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkSparseBufferMemoryBindInfo) fromC(c *C.VkSparseBufferMemoryBindInfo) {
	g.buffer = VkBuffer(c.buffer)
	g.pBinds = make([]VkSparseMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].fromC(&slice2[i2])
		}
	}
}

type VkSparseImageOpaqueMemoryBindInfo struct {
	image  VkImage
	pBinds []VkSparseMemoryBind
}

func (g *VkSparseImageOpaqueMemoryBindInfo) toC(c *C.VkSparseImageOpaqueMemoryBindInfo, _sa *stackAllocator) {
	c.image = C.VkImage(g.image)
	c.bindCount = C.uint32_t(len(g.pBinds))
	{
		c.pBinds = (*C.VkSparseMemoryBind)(_sa.alloc(C.sizeof_VkSparseMemoryBind * uint(len(g.pBinds))))
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkSparseImageOpaqueMemoryBindInfo) fromC(c *C.VkSparseImageOpaqueMemoryBindInfo) {
	g.image = VkImage(c.image)
	g.pBinds = make([]VkSparseMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].fromC(&slice2[i2])
		}
	}
}

type VkImageSubresource struct {
	aspectMask VkImageAspectFlags
	mipLevel   uint32
	arrayLayer uint32
}

func (g *VkImageSubresource) toC(c *C.VkImageSubresource) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.aspectMask)))
			_temp = C.VkFlags(_temp)
		}
		c.aspectMask = C.VkImageAspectFlags(_temp)
	}
	c.mipLevel = C.uint32_t(g.mipLevel)
	c.arrayLayer = C.uint32_t(g.arrayLayer)
}
func (g *VkImageSubresource) fromC(c *C.VkImageSubresource) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			_temp = VkFlags(_temp)
		}
		g.aspectMask = VkImageAspectFlags(_temp)
	}
	g.mipLevel = uint32(c.mipLevel)
	g.arrayLayer = uint32(c.arrayLayer)
}

type VkOffset3D struct {
	x int32
	y int32
	z int32
}

func (g *VkOffset3D) toC(c *C.VkOffset3D) {
	c.x = C.int32_t(g.x)
	c.y = C.int32_t(g.y)
	c.z = C.int32_t(g.z)
}
func (g *VkOffset3D) fromC(c *C.VkOffset3D) {
	g.x = int32(c.x)
	g.y = int32(c.y)
	g.z = int32(c.z)
}

type VkSparseImageMemoryBind struct {
	subresource  VkImageSubresource
	offset       VkOffset3D
	extent       VkExtent3D
	memory       VkDeviceMemory
	memoryOffset VkDeviceSize
	flags        VkSparseMemoryBindFlags
}

func (g *VkSparseImageMemoryBind) toC(c *C.VkSparseImageMemoryBind) {
	g.subresource.toC(&c.subresource)
	g.offset.toC(&c.offset)
	g.extent.toC(&c.extent)
	c.memory = C.VkDeviceMemory(g.memory)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.memoryOffset))
		c.memoryOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSparseMemoryBindFlags(_temp)
	}
}
func (g *VkSparseImageMemoryBind) fromC(c *C.VkSparseImageMemoryBind) {
	g.subresource.fromC(&c.subresource)
	g.offset.fromC(&c.offset)
	g.extent.fromC(&c.extent)
	g.memory = VkDeviceMemory(c.memory)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.memoryOffset))
		g.memoryOffset = VkDeviceSize(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSparseMemoryBindFlags(_temp)
	}
}

type VkSparseImageMemoryBindInfo struct {
	image  VkImage
	pBinds []VkSparseImageMemoryBind
}

func (g *VkSparseImageMemoryBindInfo) toC(c *C.VkSparseImageMemoryBindInfo, _sa *stackAllocator) {
	c.image = C.VkImage(g.image)
	c.bindCount = C.uint32_t(len(g.pBinds))
	{
		c.pBinds = (*C.VkSparseImageMemoryBind)(_sa.alloc(C.sizeof_VkSparseImageMemoryBind * uint(len(g.pBinds))))
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkSparseImageMemoryBindInfo) fromC(c *C.VkSparseImageMemoryBindInfo) {
	g.image = VkImage(c.image)
	g.pBinds = make([]VkSparseImageMemoryBind, int(c.bindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBind)(unsafe.Pointer(c.pBinds))[:len(g.pBinds):len(g.pBinds)]
		for i2, _ := range g.pBinds {
			g.pBinds[i2].fromC(&slice2[i2])
		}
	}
}

type VkBindSparseInfo struct {
	sType             VkStructureType
	pNext             *[0]byte
	pWaitSemaphores   []VkSemaphore
	pBufferBinds      []VkSparseBufferMemoryBindInfo
	pImageOpaqueBinds []VkSparseImageOpaqueMemoryBindInfo
	pImageBinds       []VkSparseImageMemoryBindInfo
	pSignalSemaphores []VkSemaphore
}

func (g *VkBindSparseInfo) toC(c *C.VkBindSparseInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.waitSemaphoreCount = C.uint32_t(len(g.pWaitSemaphores))
	{
		c.pWaitSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.pWaitSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.pWaitSemaphores):len(g.pWaitSemaphores)]
		for i2, _ := range g.pWaitSemaphores {
			slice2[i2] = C.VkSemaphore(g.pWaitSemaphores[i2])
		}
	}
	c.bufferBindCount = C.uint32_t(len(g.pBufferBinds))
	{
		c.pBufferBinds = (*C.VkSparseBufferMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseBufferMemoryBindInfo * uint(len(g.pBufferBinds))))
		slice2 := (*[1 << 31]C.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(c.pBufferBinds))[:len(g.pBufferBinds):len(g.pBufferBinds)]
		for i2, _ := range g.pBufferBinds {
			g.pBufferBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.imageOpaqueBindCount = C.uint32_t(len(g.pImageOpaqueBinds))
	{
		c.pImageOpaqueBinds = (*C.VkSparseImageOpaqueMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseImageOpaqueMemoryBindInfo * uint(len(g.pImageOpaqueBinds))))
		slice2 := (*[1 << 31]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(c.pImageOpaqueBinds))[:len(g.pImageOpaqueBinds):len(g.pImageOpaqueBinds)]
		for i2, _ := range g.pImageOpaqueBinds {
			g.pImageOpaqueBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.imageBindCount = C.uint32_t(len(g.pImageBinds))
	{
		c.pImageBinds = (*C.VkSparseImageMemoryBindInfo)(_sa.alloc(C.sizeof_VkSparseImageMemoryBindInfo * uint(len(g.pImageBinds))))
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBindInfo)(unsafe.Pointer(c.pImageBinds))[:len(g.pImageBinds):len(g.pImageBinds)]
		for i2, _ := range g.pImageBinds {
			g.pImageBinds[i2].toC(&slice2[i2], _sa)
		}
	}
	c.signalSemaphoreCount = C.uint32_t(len(g.pSignalSemaphores))
	{
		c.pSignalSemaphores = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore * uint(len(g.pSignalSemaphores))))
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.pSignalSemaphores):len(g.pSignalSemaphores)]
		for i2, _ := range g.pSignalSemaphores {
			slice2[i2] = C.VkSemaphore(g.pSignalSemaphores[i2])
		}
	}
}
func (g *VkBindSparseInfo) fromC(c *C.VkBindSparseInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.pWaitSemaphores = make([]VkSemaphore, int(c.waitSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pWaitSemaphores))[:len(g.pWaitSemaphores):len(g.pWaitSemaphores)]
		for i2, _ := range g.pWaitSemaphores {
			g.pWaitSemaphores[i2] = VkSemaphore(slice2[i2])
		}
	}
	g.pBufferBinds = make([]VkSparseBufferMemoryBindInfo, int(c.bufferBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(c.pBufferBinds))[:len(g.pBufferBinds):len(g.pBufferBinds)]
		for i2, _ := range g.pBufferBinds {
			g.pBufferBinds[i2].fromC(&slice2[i2])
		}
	}
	g.pImageOpaqueBinds = make([]VkSparseImageOpaqueMemoryBindInfo, int(c.imageOpaqueBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(c.pImageOpaqueBinds))[:len(g.pImageOpaqueBinds):len(g.pImageOpaqueBinds)]
		for i2, _ := range g.pImageOpaqueBinds {
			g.pImageOpaqueBinds[i2].fromC(&slice2[i2])
		}
	}
	g.pImageBinds = make([]VkSparseImageMemoryBindInfo, int(c.imageBindCount))
	{
		slice2 := (*[1 << 31]C.VkSparseImageMemoryBindInfo)(unsafe.Pointer(c.pImageBinds))[:len(g.pImageBinds):len(g.pImageBinds)]
		for i2, _ := range g.pImageBinds {
			g.pImageBinds[i2].fromC(&slice2[i2])
		}
	}
	g.pSignalSemaphores = make([]VkSemaphore, int(c.signalSemaphoreCount))
	{
		slice2 := (*[1 << 31]C.VkSemaphore)(unsafe.Pointer(c.pSignalSemaphores))[:len(g.pSignalSemaphores):len(g.pSignalSemaphores)]
		for i2, _ := range g.pSignalSemaphores {
			g.pSignalSemaphores[i2] = VkSemaphore(slice2[i2])
		}
	}
}
func vkQueueBindSparse(queue VkQueue, bindInfoCount uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (_ret VkResult) {
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
		pBindInfo.toC(c.pBindInfo, _sa)
	}
	c.fence = C.VkFence(fence)
	c._ret = C.vkQueueBindSparse(c.queue, c.bindInfoCount, c.pBindInfo, c.fence)
	_ret = VkResult(c._ret)
	return
}

type VkFenceCreateFlags VkFlags
type VkFenceCreateInfo struct {
	sType VkStructureType
	pNext *[0]byte
	flags VkFenceCreateFlags
}

func (g *VkFenceCreateInfo) toC(c *C.VkFenceCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkFenceCreateFlags(_temp)
	}
}
func (g *VkFenceCreateInfo) fromC(c *C.VkFenceCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkFenceCreateFlags(_temp)
	}
}
func vkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pFence = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence))
		*c.pFence = C.VkFence(*pFence)
	}
	c._ret = C.vkCreateFence(c.device, c.pCreateInfo, c.pAllocator, c.pFence)
	_ret = VkResult(c._ret)
	*pFence = VkFence(*c.pFence)
	return
}
func vkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyFence(c.device, c.fence, c.pAllocator)
}
func vkResetFences(device VkDevice, pFences []VkFence) (_ret VkResult) {
	var c struct {
		device     C.VkDevice
		fenceCount C.uint32_t
		pFences    *C.VkFence
		_ret       C.VkResult
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.device = C.VkDevice(device)
	c.fenceCount = C.uint32_t(len(pFences))
	{
		c.pFences = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence * uint(len(pFences))))
		slice3 := (*[1 << 31]C.VkFence)(unsafe.Pointer(c.pFences))[:len(pFences):len(pFences)]
		for i3, _ := range pFences {
			slice3[i3] = C.VkFence(pFences[i3])
		}
	}
	c._ret = C.vkResetFences(c.device, c.fenceCount, c.pFences)
	_ret = VkResult(c._ret)
	return
}
func vkGetFenceStatus(device VkDevice, fence VkFence) (_ret VkResult) {
	var c struct {
		device C.VkDevice
		fence  C.VkFence
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.fence = C.VkFence(fence)
	c._ret = C.vkGetFenceStatus(c.device, c.fence)
	_ret = VkResult(c._ret)
	return
}
func vkWaitForFences(device VkDevice, pFences []VkFence, waitAll VkBool32, timeout uint64) (_ret VkResult) {
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
	c.fenceCount = C.uint32_t(len(pFences))
	{
		c.pFences = (*C.VkFence)(_sa.alloc(C.sizeof_VkFence * uint(len(pFences))))
		slice3 := (*[1 << 31]C.VkFence)(unsafe.Pointer(c.pFences))[:len(pFences):len(pFences)]
		for i3, _ := range pFences {
			slice3[i3] = C.VkFence(pFences[i3])
		}
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(waitAll))
		c.waitAll = C.VkBool32(_temp)
	}
	c.timeout = C.uint64_t(timeout)
	c._ret = C.vkWaitForFences(c.device, c.fenceCount, c.pFences, c.waitAll, c.timeout)
	_ret = VkResult(c._ret)
	return
}

type VkSemaphoreCreateFlags VkFlags
type VkSemaphoreCreateInfo struct {
	sType VkStructureType
	pNext *[0]byte
	flags VkSemaphoreCreateFlags
}

func (g *VkSemaphoreCreateInfo) toC(c *C.VkSemaphoreCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSemaphoreCreateFlags(_temp)
	}
}
func (g *VkSemaphoreCreateInfo) fromC(c *C.VkSemaphoreCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSemaphoreCreateFlags(_temp)
	}
}
func vkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSemaphore = (*C.VkSemaphore)(_sa.alloc(C.sizeof_VkSemaphore))
		*c.pSemaphore = C.VkSemaphore(*pSemaphore)
	}
	c._ret = C.vkCreateSemaphore(c.device, c.pCreateInfo, c.pAllocator, c.pSemaphore)
	_ret = VkResult(c._ret)
	*pSemaphore = VkSemaphore(*c.pSemaphore)
	return
}
func vkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySemaphore(c.device, c.semaphore, c.pAllocator)
}

type VkEventCreateFlags VkFlags
type VkEventCreateInfo struct {
	sType VkStructureType
	pNext *[0]byte
	flags VkEventCreateFlags
}

func (g *VkEventCreateInfo) toC(c *C.VkEventCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkEventCreateFlags(_temp)
	}
}
func (g *VkEventCreateInfo) fromC(c *C.VkEventCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkEventCreateFlags(_temp)
	}
}

type VkEvent C.VkEvent

func vkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pEvent = (*C.VkEvent)(_sa.alloc(C.sizeof_VkEvent))
		*c.pEvent = C.VkEvent(*pEvent)
	}
	c._ret = C.vkCreateEvent(c.device, c.pCreateInfo, c.pAllocator, c.pEvent)
	_ret = VkResult(c._ret)
	*pEvent = VkEvent(*c.pEvent)
	return
}
func vkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyEvent(c.device, c.event, c.pAllocator)
}
func vkGetEventStatus(device VkDevice, event VkEvent) (_ret VkResult) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkGetEventStatus(c.device, c.event)
	_ret = VkResult(c._ret)
	return
}
func vkSetEvent(device VkDevice, event VkEvent) (_ret VkResult) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkSetEvent(c.device, c.event)
	_ret = VkResult(c._ret)
	return
}
func vkResetEvent(device VkDevice, event VkEvent) (_ret VkResult) {
	var c struct {
		device C.VkDevice
		event  C.VkEvent
		_ret   C.VkResult
	}
	c.device = C.VkDevice(device)
	c.event = C.VkEvent(event)
	c._ret = C.vkResetEvent(c.device, c.event)
	_ret = VkResult(c._ret)
	return
}

type VkQueryPoolCreateFlags VkFlags
type VkQueryType int

const (
	VK_QUERY_TYPE_OCCLUSION           VkQueryType = 0
	VK_QUERY_TYPE_PIPELINE_STATISTICS VkQueryType = 1
	VK_QUERY_TYPE_TIMESTAMP           VkQueryType = 2
	VK_QUERY_TYPE_BEGIN_RANGE         VkQueryType = VK_QUERY_TYPE_OCCLUSION
	VK_QUERY_TYPE_END_RANGE           VkQueryType = VK_QUERY_TYPE_TIMESTAMP
	VK_QUERY_TYPE_RANGE_SIZE          VkQueryType = (VK_QUERY_TYPE_TIMESTAMP - VK_QUERY_TYPE_OCCLUSION + 1)
	VK_QUERY_TYPE_MAX_ENUM            VkQueryType = 2147483647
)

type VkQueryPipelineStatisticFlags VkFlags
type VkQueryPoolCreateInfo struct {
	sType              VkStructureType
	pNext              *[0]byte
	flags              VkQueryPoolCreateFlags
	queryType          VkQueryType
	queryCount         uint32
	pipelineStatistics VkQueryPipelineStatisticFlags
}

func (g *VkQueryPoolCreateInfo) toC(c *C.VkQueryPoolCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkQueryPoolCreateFlags(_temp)
	}
	c.queryType = C.VkQueryType(g.queryType)
	c.queryCount = C.uint32_t(g.queryCount)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.pipelineStatistics)))
			_temp = C.VkFlags(_temp)
		}
		c.pipelineStatistics = C.VkQueryPipelineStatisticFlags(_temp)
	}
}
func (g *VkQueryPoolCreateInfo) fromC(c *C.VkQueryPoolCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkQueryPoolCreateFlags(_temp)
	}
	g.queryType = VkQueryType(c.queryType)
	g.queryCount = uint32(c.queryCount)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.pipelineStatistics)))
			_temp = VkFlags(_temp)
		}
		g.pipelineStatistics = VkQueryPipelineStatisticFlags(_temp)
	}
}

type VkQueryPool C.VkQueryPool

func vkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pQueryPool = (*C.VkQueryPool)(_sa.alloc(C.sizeof_VkQueryPool))
		*c.pQueryPool = C.VkQueryPool(*pQueryPool)
	}
	c._ret = C.vkCreateQueryPool(c.device, c.pCreateInfo, c.pAllocator, c.pQueryPool)
	_ret = VkResult(c._ret)
	*pQueryPool = VkQueryPool(*c.pQueryPool)
	return
}
func vkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyQueryPool(c.device, c.queryPool, c.pAllocator)
}

type VkQueryResultFlags VkFlags

func vkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery uint32, queryCount uint32, pData []byte, stride VkDeviceSize, flags VkQueryResultFlags) (_ret VkResult) {
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
	c.dataSize = C.size_t(len(pData))
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(pData)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(pData):len(pData)]
		for i3, _ := range pData {
			slice3[i3] = pData[i3]
		}
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(stride))
		c.stride = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkQueryResultFlags(_temp)
	}
	c._ret = C.vkGetQueryPoolResults(c.device, c.queryPool, c.firstQuery, c.queryCount, c.dataSize, c.pData, c.stride, c.flags)
	_ret = VkResult(c._ret)
	return
}

type VkBufferCreateFlags VkFlags
type VkBufferUsageFlags VkFlags
type VkSharingMode int

const (
	VK_SHARING_MODE_EXCLUSIVE   VkSharingMode = 0
	VK_SHARING_MODE_CONCURRENT  VkSharingMode = 1
	VK_SHARING_MODE_BEGIN_RANGE VkSharingMode = VK_SHARING_MODE_EXCLUSIVE
	VK_SHARING_MODE_END_RANGE   VkSharingMode = VK_SHARING_MODE_CONCURRENT
	VK_SHARING_MODE_RANGE_SIZE  VkSharingMode = (VK_SHARING_MODE_CONCURRENT - VK_SHARING_MODE_EXCLUSIVE + 1)
	VK_SHARING_MODE_MAX_ENUM    VkSharingMode = 2147483647
)

type VkBufferCreateInfo struct {
	sType               VkStructureType
	pNext               *[0]byte
	flags               VkBufferCreateFlags
	size                VkDeviceSize
	usage               VkBufferUsageFlags
	sharingMode         VkSharingMode
	pQueueFamilyIndices []uint32
}

func (g *VkBufferCreateInfo) toC(c *C.VkBufferCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkBufferCreateFlags(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.usage)))
			_temp = C.VkFlags(_temp)
		}
		c.usage = C.VkBufferUsageFlags(_temp)
	}
	c.sharingMode = C.VkSharingMode(g.sharingMode)
	c.queueFamilyIndexCount = C.uint32_t(len(g.pQueueFamilyIndices))
	{
		c.pQueueFamilyIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.pQueueFamilyIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.pQueueFamilyIndices):len(g.pQueueFamilyIndices)]
		for i2, _ := range g.pQueueFamilyIndices {
			slice2[i2] = C.uint32_t(g.pQueueFamilyIndices[i2])
		}
	}
}
func (g *VkBufferCreateInfo) fromC(c *C.VkBufferCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkBufferCreateFlags(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			_temp = VkFlags(_temp)
		}
		g.usage = VkBufferUsageFlags(_temp)
	}
	g.sharingMode = VkSharingMode(c.sharingMode)
	g.pQueueFamilyIndices = make([]uint32, int(c.queueFamilyIndexCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.pQueueFamilyIndices):len(g.pQueueFamilyIndices)]
		for i2, _ := range g.pQueueFamilyIndices {
			g.pQueueFamilyIndices[i2] = uint32(slice2[i2])
		}
	}
}
func vkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pBuffer = (*C.VkBuffer)(_sa.alloc(C.sizeof_VkBuffer))
		*c.pBuffer = C.VkBuffer(*pBuffer)
	}
	c._ret = C.vkCreateBuffer(c.device, c.pCreateInfo, c.pAllocator, c.pBuffer)
	_ret = VkResult(c._ret)
	*pBuffer = VkBuffer(*c.pBuffer)
	return
}
func vkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyBuffer(c.device, c.buffer, c.pAllocator)
}

type VkBufferViewCreateFlags VkFlags
type VkBufferViewCreateInfo struct {
	sType  VkStructureType
	pNext  *[0]byte
	flags  VkBufferViewCreateFlags
	buffer VkBuffer
	format VkFormat
	offset VkDeviceSize
	range  VkDeviceSize
}

func (g *VkBufferViewCreateInfo) toC(c *C.VkBufferViewCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkBufferViewCreateFlags(_temp)
	}
	c.buffer = C.VkBuffer(g.buffer)
	c.format = C.VkFormat(g.format)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.range))
		c.range = C.VkDeviceSize(_temp)
	}
}
func (g *VkBufferViewCreateInfo) fromC(c *C.VkBufferViewCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkBufferViewCreateFlags(_temp)
	}
	g.buffer = VkBuffer(c.buffer)
	g.format = VkFormat(c.format)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.offset))
		g.offset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.range))
		g.range = VkDeviceSize(_temp)
	}
}

type VkBufferView C.VkBufferView

func vkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pView = (*C.VkBufferView)(_sa.alloc(C.sizeof_VkBufferView))
		*c.pView = C.VkBufferView(*pView)
	}
	c._ret = C.vkCreateBufferView(c.device, c.pCreateInfo, c.pAllocator, c.pView)
	_ret = VkResult(c._ret)
	*pView = VkBufferView(*c.pView)
	return
}
func vkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyBufferView(c.device, c.bufferView, c.pAllocator)
}

type VkImageLayout int

const (
	VK_IMAGE_LAYOUT_UNDEFINED                                      VkImageLayout = 0
	VK_IMAGE_LAYOUT_GENERAL                                        VkImageLayout = 1
	VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL                       VkImageLayout = 2
	VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL               VkImageLayout = 3
	VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL                VkImageLayout = 4
	VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL                       VkImageLayout = 5
	VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL                           VkImageLayout = 6
	VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL                           VkImageLayout = 7
	VK_IMAGE_LAYOUT_PREINITIALIZED                                 VkImageLayout = 8
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL     VkImageLayout = 1000117000
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL     VkImageLayout = 1000117001
	VK_IMAGE_LAYOUT_PRESENT_SRC_KHR                                VkImageLayout = 1000001002
	VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR                             VkImageLayout = 1000111000
	VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
	VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR VkImageLayout = VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	VK_IMAGE_LAYOUT_BEGIN_RANGE                                    VkImageLayout = VK_IMAGE_LAYOUT_UNDEFINED
	VK_IMAGE_LAYOUT_END_RANGE                                      VkImageLayout = VK_IMAGE_LAYOUT_PREINITIALIZED
	VK_IMAGE_LAYOUT_RANGE_SIZE                                     VkImageLayout = (VK_IMAGE_LAYOUT_PREINITIALIZED - VK_IMAGE_LAYOUT_UNDEFINED + 1)
	VK_IMAGE_LAYOUT_MAX_ENUM                                       VkImageLayout = 2147483647
)

type VkImageCreateInfo struct {
	sType               VkStructureType
	pNext               *[0]byte
	flags               VkImageCreateFlags
	imageType           VkImageType
	format              VkFormat
	extent              VkExtent3D
	mipLevels           uint32
	arrayLayers         uint32
	samples             VkSampleCountFlagBits
	tiling              VkImageTiling
	usage               VkImageUsageFlags
	sharingMode         VkSharingMode
	pQueueFamilyIndices []uint32
	initialLayout       VkImageLayout
}

func (g *VkImageCreateInfo) toC(c *C.VkImageCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkImageCreateFlags(_temp)
	}
	c.imageType = C.VkImageType(g.imageType)
	c.format = C.VkFormat(g.format)
	g.extent.toC(&c.extent)
	c.mipLevels = C.uint32_t(g.mipLevels)
	c.arrayLayers = C.uint32_t(g.arrayLayers)
	c.samples = C.VkSampleCountFlagBits(g.samples)
	c.tiling = C.VkImageTiling(g.tiling)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.usage)))
			_temp = C.VkFlags(_temp)
		}
		c.usage = C.VkImageUsageFlags(_temp)
	}
	c.sharingMode = C.VkSharingMode(g.sharingMode)
	c.queueFamilyIndexCount = C.uint32_t(len(g.pQueueFamilyIndices))
	{
		c.pQueueFamilyIndices = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.pQueueFamilyIndices))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.pQueueFamilyIndices):len(g.pQueueFamilyIndices)]
		for i2, _ := range g.pQueueFamilyIndices {
			slice2[i2] = C.uint32_t(g.pQueueFamilyIndices[i2])
		}
	}
	c.initialLayout = C.VkImageLayout(g.initialLayout)
}
func (g *VkImageCreateInfo) fromC(c *C.VkImageCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkImageCreateFlags(_temp)
	}
	g.imageType = VkImageType(c.imageType)
	g.format = VkFormat(c.format)
	g.extent.fromC(&c.extent)
	g.mipLevels = uint32(c.mipLevels)
	g.arrayLayers = uint32(c.arrayLayers)
	g.samples = VkSampleCountFlagBits(c.samples)
	g.tiling = VkImageTiling(c.tiling)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.usage)))
			_temp = VkFlags(_temp)
		}
		g.usage = VkImageUsageFlags(_temp)
	}
	g.sharingMode = VkSharingMode(c.sharingMode)
	g.pQueueFamilyIndices = make([]uint32, int(c.queueFamilyIndexCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pQueueFamilyIndices))[:len(g.pQueueFamilyIndices):len(g.pQueueFamilyIndices)]
		for i2, _ := range g.pQueueFamilyIndices {
			g.pQueueFamilyIndices[i2] = uint32(slice2[i2])
		}
	}
	g.initialLayout = VkImageLayout(c.initialLayout)
}
func vkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pImage = (*C.VkImage)(_sa.alloc(C.sizeof_VkImage))
		*c.pImage = C.VkImage(*pImage)
	}
	c._ret = C.vkCreateImage(c.device, c.pCreateInfo, c.pAllocator, c.pImage)
	_ret = VkResult(c._ret)
	*pImage = VkImage(*c.pImage)
	return
}
func vkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyImage(c.device, c.image, c.pAllocator)
}

type VkSubresourceLayout struct {
	offset     VkDeviceSize
	size       VkDeviceSize
	rowPitch   VkDeviceSize
	arrayPitch VkDeviceSize
	depthPitch VkDeviceSize
}

func (g *VkSubresourceLayout) toC(c *C.VkSubresourceLayout) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.rowPitch))
		c.rowPitch = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.arrayPitch))
		c.arrayPitch = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.depthPitch))
		c.depthPitch = C.VkDeviceSize(_temp)
	}
}
func (g *VkSubresourceLayout) fromC(c *C.VkSubresourceLayout) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.offset))
		g.offset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.rowPitch))
		g.rowPitch = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.arrayPitch))
		g.arrayPitch = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.depthPitch))
		g.depthPitch = VkDeviceSize(_temp)
	}
}
func vkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) {
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
		pSubresource.toC(c.pSubresource)
	}
	{
		c.pLayout = (*C.VkSubresourceLayout)(_sa.alloc(C.sizeof_VkSubresourceLayout))
		pLayout.toC(c.pLayout)
	}
	C.vkGetImageSubresourceLayout(c.device, c.image, c.pSubresource, c.pLayout)
	pLayout.fromC(c.pLayout)
}

type VkImageViewCreateFlags VkFlags
type VkImageViewType int

const (
	VK_IMAGE_VIEW_TYPE_1D          VkImageViewType = 0
	VK_IMAGE_VIEW_TYPE_2D          VkImageViewType = 1
	VK_IMAGE_VIEW_TYPE_3D          VkImageViewType = 2
	VK_IMAGE_VIEW_TYPE_CUBE        VkImageViewType = 3
	VK_IMAGE_VIEW_TYPE_1D_ARRAY    VkImageViewType = 4
	VK_IMAGE_VIEW_TYPE_2D_ARRAY    VkImageViewType = 5
	VK_IMAGE_VIEW_TYPE_CUBE_ARRAY  VkImageViewType = 6
	VK_IMAGE_VIEW_TYPE_BEGIN_RANGE VkImageViewType = VK_IMAGE_VIEW_TYPE_1D
	VK_IMAGE_VIEW_TYPE_END_RANGE   VkImageViewType = VK_IMAGE_VIEW_TYPE_CUBE_ARRAY
	VK_IMAGE_VIEW_TYPE_RANGE_SIZE  VkImageViewType = (VK_IMAGE_VIEW_TYPE_CUBE_ARRAY - VK_IMAGE_VIEW_TYPE_1D + 1)
	VK_IMAGE_VIEW_TYPE_MAX_ENUM    VkImageViewType = 2147483647
)

type VkComponentSwizzle int

const (
	VK_COMPONENT_SWIZZLE_IDENTITY    VkComponentSwizzle = 0
	VK_COMPONENT_SWIZZLE_ZERO        VkComponentSwizzle = 1
	VK_COMPONENT_SWIZZLE_ONE         VkComponentSwizzle = 2
	VK_COMPONENT_SWIZZLE_R           VkComponentSwizzle = 3
	VK_COMPONENT_SWIZZLE_G           VkComponentSwizzle = 4
	VK_COMPONENT_SWIZZLE_B           VkComponentSwizzle = 5
	VK_COMPONENT_SWIZZLE_A           VkComponentSwizzle = 6
	VK_COMPONENT_SWIZZLE_BEGIN_RANGE VkComponentSwizzle = VK_COMPONENT_SWIZZLE_IDENTITY
	VK_COMPONENT_SWIZZLE_END_RANGE   VkComponentSwizzle = VK_COMPONENT_SWIZZLE_A
	VK_COMPONENT_SWIZZLE_RANGE_SIZE  VkComponentSwizzle = (VK_COMPONENT_SWIZZLE_A - VK_COMPONENT_SWIZZLE_IDENTITY + 1)
	VK_COMPONENT_SWIZZLE_MAX_ENUM    VkComponentSwizzle = 2147483647
)

type VkComponentMapping struct {
	r VkComponentSwizzle
	g VkComponentSwizzle
	b VkComponentSwizzle
	a VkComponentSwizzle
}

func (g *VkComponentMapping) toC(c *C.VkComponentMapping) {
	c.r = C.VkComponentSwizzle(g.r)
	c.g = C.VkComponentSwizzle(g.g)
	c.b = C.VkComponentSwizzle(g.b)
	c.a = C.VkComponentSwizzle(g.a)
}
func (g *VkComponentMapping) fromC(c *C.VkComponentMapping) {
	g.r = VkComponentSwizzle(c.r)
	g.g = VkComponentSwizzle(c.g)
	g.b = VkComponentSwizzle(c.b)
	g.a = VkComponentSwizzle(c.a)
}

type VkImageSubresourceRange struct {
	aspectMask     VkImageAspectFlags
	baseMipLevel   uint32
	levelCount     uint32
	baseArrayLayer uint32
	layerCount     uint32
}

func (g *VkImageSubresourceRange) toC(c *C.VkImageSubresourceRange) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.aspectMask)))
			_temp = C.VkFlags(_temp)
		}
		c.aspectMask = C.VkImageAspectFlags(_temp)
	}
	c.baseMipLevel = C.uint32_t(g.baseMipLevel)
	c.levelCount = C.uint32_t(g.levelCount)
	c.baseArrayLayer = C.uint32_t(g.baseArrayLayer)
	c.layerCount = C.uint32_t(g.layerCount)
}
func (g *VkImageSubresourceRange) fromC(c *C.VkImageSubresourceRange) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			_temp = VkFlags(_temp)
		}
		g.aspectMask = VkImageAspectFlags(_temp)
	}
	g.baseMipLevel = uint32(c.baseMipLevel)
	g.levelCount = uint32(c.levelCount)
	g.baseArrayLayer = uint32(c.baseArrayLayer)
	g.layerCount = uint32(c.layerCount)
}

type VkImageViewCreateInfo struct {
	sType            VkStructureType
	pNext            *[0]byte
	flags            VkImageViewCreateFlags
	image            VkImage
	viewType         VkImageViewType
	format           VkFormat
	components       VkComponentMapping
	subresourceRange VkImageSubresourceRange
}

func (g *VkImageViewCreateInfo) toC(c *C.VkImageViewCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkImageViewCreateFlags(_temp)
	}
	c.image = C.VkImage(g.image)
	c.viewType = C.VkImageViewType(g.viewType)
	c.format = C.VkFormat(g.format)
	g.components.toC(&c.components)
	g.subresourceRange.toC(&c.subresourceRange)
}
func (g *VkImageViewCreateInfo) fromC(c *C.VkImageViewCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkImageViewCreateFlags(_temp)
	}
	g.image = VkImage(c.image)
	g.viewType = VkImageViewType(c.viewType)
	g.format = VkFormat(c.format)
	g.components.fromC(&c.components)
	g.subresourceRange.fromC(&c.subresourceRange)
}

type VkImageView C.VkImageView

func vkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pView = (*C.VkImageView)(_sa.alloc(C.sizeof_VkImageView))
		*c.pView = C.VkImageView(*pView)
	}
	c._ret = C.vkCreateImageView(c.device, c.pCreateInfo, c.pAllocator, c.pView)
	_ret = VkResult(c._ret)
	*pView = VkImageView(*c.pView)
	return
}
func vkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyImageView(c.device, c.imageView, c.pAllocator)
}

type VkShaderModuleCreateFlags VkFlags
type VkShaderModuleCreateInfo struct {
	sType    VkStructureType
	pNext    *[0]byte
	flags    VkShaderModuleCreateFlags
	codeSize uint
	pCode    *uint32
}

func (g *VkShaderModuleCreateInfo) toC(c *C.VkShaderModuleCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkShaderModuleCreateFlags(_temp)
	}
	c.codeSize = C.size_t(g.codeSize)
	{
		c.pCode = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t))
		*c.pCode = C.uint32_t(*g.pCode)
	}
}
func (g *VkShaderModuleCreateInfo) fromC(c *C.VkShaderModuleCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkShaderModuleCreateFlags(_temp)
	}
	g.codeSize = uint(c.codeSize)
	{
		if g.pCode == nil {
			g.pCode = new(uint32)
		}
		*g.pCode = uint32(*c.pCode)
	}
}

type VkShaderModule C.VkShaderModule

func vkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pShaderModule = (*C.VkShaderModule)(_sa.alloc(C.sizeof_VkShaderModule))
		*c.pShaderModule = C.VkShaderModule(*pShaderModule)
	}
	c._ret = C.vkCreateShaderModule(c.device, c.pCreateInfo, c.pAllocator, c.pShaderModule)
	_ret = VkResult(c._ret)
	*pShaderModule = VkShaderModule(*c.pShaderModule)
	return
}
func vkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyShaderModule(c.device, c.shaderModule, c.pAllocator)
}

type VkPipelineCacheCreateFlags VkFlags
type VkPipelineCacheCreateInfo struct {
	sType        VkStructureType
	pNext        *[0]byte
	flags        VkPipelineCacheCreateFlags
	pInitialData []byte
}

func (g *VkPipelineCacheCreateInfo) toC(c *C.VkPipelineCacheCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineCacheCreateFlags(_temp)
	}
	c.initialDataSize = C.size_t(len(g.pInitialData))
	{
		c.pInitialData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.pInitialData)))
		slice2 := (*[1 << 31]byte)(c.pInitialData)[:len(g.pInitialData):len(g.pInitialData)]
		for i2, _ := range g.pInitialData {
			slice2[i2] = g.pInitialData[i2]
		}
	}
}
func (g *VkPipelineCacheCreateInfo) fromC(c *C.VkPipelineCacheCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineCacheCreateFlags(_temp)
	}
	g.pInitialData = make([]byte, int(c.initialDataSize))
	{
		slice2 := (*[1 << 31]byte)(c.pInitialData)[:len(g.pInitialData):len(g.pInitialData)]
		for i2, _ := range g.pInitialData {
			g.pInitialData[i2] = slice2[i2]
		}
	}
}

type VkPipelineCache C.VkPipelineCache

func vkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelineCache = (*C.VkPipelineCache)(_sa.alloc(C.sizeof_VkPipelineCache))
		*c.pPipelineCache = C.VkPipelineCache(*pPipelineCache)
	}
	c._ret = C.vkCreatePipelineCache(c.device, c.pCreateInfo, c.pAllocator, c.pPipelineCache)
	_ret = VkResult(c._ret)
	*pPipelineCache = VkPipelineCache(*c.pPipelineCache)
	return
}
func vkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipelineCache(c.device, c.pipelineCache, c.pAllocator)
}
func vkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *uint, pData []byte) (_ret VkResult) {
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
		*c.pDataSize = C.size_t(*pDataSize)
	}
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(pData)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(pData):len(pData)]
		for i3, _ := range pData {
			slice3[i3] = pData[i3]
		}
	}
	c._ret = C.vkGetPipelineCacheData(c.device, c.pipelineCache, c.pDataSize, c.pData)
	_ret = VkResult(c._ret)
	*pDataSize = uint(*c.pDataSize)
	return
}
func vkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, pSrcCaches []VkPipelineCache) (_ret VkResult) {
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
	c.srcCacheCount = C.uint32_t(len(pSrcCaches))
	{
		c.pSrcCaches = (*C.VkPipelineCache)(_sa.alloc(C.sizeof_VkPipelineCache * uint(len(pSrcCaches))))
		slice3 := (*[1 << 31]C.VkPipelineCache)(unsafe.Pointer(c.pSrcCaches))[:len(pSrcCaches):len(pSrcCaches)]
		for i3, _ := range pSrcCaches {
			slice3[i3] = C.VkPipelineCache(pSrcCaches[i3])
		}
	}
	c._ret = C.vkMergePipelineCaches(c.device, c.dstCache, c.srcCacheCount, c.pSrcCaches)
	_ret = VkResult(c._ret)
	return
}

type VkPipelineCreateFlags VkFlags
type VkPipelineShaderStageCreateFlags VkFlags
type VkShaderStageFlagBits int

const (
	VK_SHADER_STAGE_VERTEX_BIT                  VkShaderStageFlagBits = 1
	VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT    VkShaderStageFlagBits = 2
	VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT VkShaderStageFlagBits = 4
	VK_SHADER_STAGE_GEOMETRY_BIT                VkShaderStageFlagBits = 8
	VK_SHADER_STAGE_FRAGMENT_BIT                VkShaderStageFlagBits = 16
	VK_SHADER_STAGE_COMPUTE_BIT                 VkShaderStageFlagBits = 32
	VK_SHADER_STAGE_ALL_GRAPHICS                VkShaderStageFlagBits = 31
	VK_SHADER_STAGE_ALL                         VkShaderStageFlagBits = 2147483647
	VK_SHADER_STAGE_FLAG_BITS_MAX_ENUM          VkShaderStageFlagBits = 2147483647
)

type VkSpecializationMapEntry struct {
	constantID uint32
	offset     uint32
	size       uint
}

func (g *VkSpecializationMapEntry) toC(c *C.VkSpecializationMapEntry) {
	c.constantID = C.uint32_t(g.constantID)
	c.offset = C.uint32_t(g.offset)
	c.size = C.size_t(g.size)
}
func (g *VkSpecializationMapEntry) fromC(c *C.VkSpecializationMapEntry) {
	g.constantID = uint32(c.constantID)
	g.offset = uint32(c.offset)
	g.size = uint(c.size)
}

type VkSpecializationInfo struct {
	pMapEntries []VkSpecializationMapEntry
	pData       []byte
}

func (g *VkSpecializationInfo) toC(c *C.VkSpecializationInfo, _sa *stackAllocator) {
	c.mapEntryCount = C.uint32_t(len(g.pMapEntries))
	{
		c.pMapEntries = (*C.VkSpecializationMapEntry)(_sa.alloc(C.sizeof_VkSpecializationMapEntry * uint(len(g.pMapEntries))))
		slice2 := (*[1 << 31]C.VkSpecializationMapEntry)(unsafe.Pointer(c.pMapEntries))[:len(g.pMapEntries):len(g.pMapEntries)]
		for i2, _ := range g.pMapEntries {
			g.pMapEntries[i2].toC(&slice2[i2])
		}
	}
	c.dataSize = C.size_t(len(g.pData))
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(g.pData)))
		slice2 := (*[1 << 31]byte)(c.pData)[:len(g.pData):len(g.pData)]
		for i2, _ := range g.pData {
			slice2[i2] = g.pData[i2]
		}
	}
}
func (g *VkSpecializationInfo) fromC(c *C.VkSpecializationInfo) {
	g.pMapEntries = make([]VkSpecializationMapEntry, int(c.mapEntryCount))
	{
		slice2 := (*[1 << 31]C.VkSpecializationMapEntry)(unsafe.Pointer(c.pMapEntries))[:len(g.pMapEntries):len(g.pMapEntries)]
		for i2, _ := range g.pMapEntries {
			g.pMapEntries[i2].fromC(&slice2[i2])
		}
	}
	g.pData = make([]byte, int(c.dataSize))
	{
		slice2 := (*[1 << 31]byte)(c.pData)[:len(g.pData):len(g.pData)]
		for i2, _ := range g.pData {
			g.pData[i2] = slice2[i2]
		}
	}
}

type VkPipelineShaderStageCreateInfo struct {
	sType               VkStructureType
	pNext               *[0]byte
	flags               VkPipelineShaderStageCreateFlags
	stage               VkShaderStageFlagBits
	module              VkShaderModule
	pName               string
	pSpecializationInfo *VkSpecializationInfo
}

func (g *VkPipelineShaderStageCreateInfo) toC(c *C.VkPipelineShaderStageCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineShaderStageCreateFlags(_temp)
	}
	c.stage = C.VkShaderStageFlagBits(g.stage)
	c.module = C.VkShaderModule(g.module)
	c.pName = toCString(g.pName, _sa)
	{
		c.pSpecializationInfo = (*C.VkSpecializationInfo)(_sa.alloc(C.sizeof_VkSpecializationInfo))
		g.pSpecializationInfo.toC(c.pSpecializationInfo, _sa)
	}
}
func (g *VkPipelineShaderStageCreateInfo) fromC(c *C.VkPipelineShaderStageCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineShaderStageCreateFlags(_temp)
	}
	g.stage = VkShaderStageFlagBits(c.stage)
	g.module = VkShaderModule(c.module)
	g.pName = toGoString(c.pName)
	{
		if g.pSpecializationInfo == nil {
			g.pSpecializationInfo = new(VkSpecializationInfo)
		}
		g.pSpecializationInfo.fromC(c.pSpecializationInfo)
	}
}

type VkPipelineVertexInputStateCreateFlags VkFlags
type VkVertexInputRate int

const (
	VK_VERTEX_INPUT_RATE_VERTEX      VkVertexInputRate = 0
	VK_VERTEX_INPUT_RATE_INSTANCE    VkVertexInputRate = 1
	VK_VERTEX_INPUT_RATE_BEGIN_RANGE VkVertexInputRate = VK_VERTEX_INPUT_RATE_VERTEX
	VK_VERTEX_INPUT_RATE_END_RANGE   VkVertexInputRate = VK_VERTEX_INPUT_RATE_INSTANCE
	VK_VERTEX_INPUT_RATE_RANGE_SIZE  VkVertexInputRate = (VK_VERTEX_INPUT_RATE_INSTANCE - VK_VERTEX_INPUT_RATE_VERTEX + 1)
	VK_VERTEX_INPUT_RATE_MAX_ENUM    VkVertexInputRate = 2147483647
)

type VkVertexInputBindingDescription struct {
	binding   uint32
	stride    uint32
	inputRate VkVertexInputRate
}

func (g *VkVertexInputBindingDescription) toC(c *C.VkVertexInputBindingDescription) {
	c.binding = C.uint32_t(g.binding)
	c.stride = C.uint32_t(g.stride)
	c.inputRate = C.VkVertexInputRate(g.inputRate)
}
func (g *VkVertexInputBindingDescription) fromC(c *C.VkVertexInputBindingDescription) {
	g.binding = uint32(c.binding)
	g.stride = uint32(c.stride)
	g.inputRate = VkVertexInputRate(c.inputRate)
}

type VkVertexInputAttributeDescription struct {
	location uint32
	binding  uint32
	format   VkFormat
	offset   uint32
}

func (g *VkVertexInputAttributeDescription) toC(c *C.VkVertexInputAttributeDescription) {
	c.location = C.uint32_t(g.location)
	c.binding = C.uint32_t(g.binding)
	c.format = C.VkFormat(g.format)
	c.offset = C.uint32_t(g.offset)
}
func (g *VkVertexInputAttributeDescription) fromC(c *C.VkVertexInputAttributeDescription) {
	g.location = uint32(c.location)
	g.binding = uint32(c.binding)
	g.format = VkFormat(c.format)
	g.offset = uint32(c.offset)
}

type VkPipelineVertexInputStateCreateInfo struct {
	sType                        VkStructureType
	pNext                        *[0]byte
	flags                        VkPipelineVertexInputStateCreateFlags
	pVertexBindingDescriptions   []VkVertexInputBindingDescription
	pVertexAttributeDescriptions []VkVertexInputAttributeDescription
}

func (g *VkPipelineVertexInputStateCreateInfo) toC(c *C.VkPipelineVertexInputStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineVertexInputStateCreateFlags(_temp)
	}
	c.vertexBindingDescriptionCount = C.uint32_t(len(g.pVertexBindingDescriptions))
	{
		c.pVertexBindingDescriptions = (*C.VkVertexInputBindingDescription)(_sa.alloc(C.sizeof_VkVertexInputBindingDescription * uint(len(g.pVertexBindingDescriptions))))
		slice2 := (*[1 << 31]C.VkVertexInputBindingDescription)(unsafe.Pointer(c.pVertexBindingDescriptions))[:len(g.pVertexBindingDescriptions):len(g.pVertexBindingDescriptions)]
		for i2, _ := range g.pVertexBindingDescriptions {
			g.pVertexBindingDescriptions[i2].toC(&slice2[i2])
		}
	}
	c.vertexAttributeDescriptionCount = C.uint32_t(len(g.pVertexAttributeDescriptions))
	{
		c.pVertexAttributeDescriptions = (*C.VkVertexInputAttributeDescription)(_sa.alloc(C.sizeof_VkVertexInputAttributeDescription * uint(len(g.pVertexAttributeDescriptions))))
		slice2 := (*[1 << 31]C.VkVertexInputAttributeDescription)(unsafe.Pointer(c.pVertexAttributeDescriptions))[:len(g.pVertexAttributeDescriptions):len(g.pVertexAttributeDescriptions)]
		for i2, _ := range g.pVertexAttributeDescriptions {
			g.pVertexAttributeDescriptions[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkPipelineVertexInputStateCreateInfo) fromC(c *C.VkPipelineVertexInputStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineVertexInputStateCreateFlags(_temp)
	}
	g.pVertexBindingDescriptions = make([]VkVertexInputBindingDescription, int(c.vertexBindingDescriptionCount))
	{
		slice2 := (*[1 << 31]C.VkVertexInputBindingDescription)(unsafe.Pointer(c.pVertexBindingDescriptions))[:len(g.pVertexBindingDescriptions):len(g.pVertexBindingDescriptions)]
		for i2, _ := range g.pVertexBindingDescriptions {
			g.pVertexBindingDescriptions[i2].fromC(&slice2[i2])
		}
	}
	g.pVertexAttributeDescriptions = make([]VkVertexInputAttributeDescription, int(c.vertexAttributeDescriptionCount))
	{
		slice2 := (*[1 << 31]C.VkVertexInputAttributeDescription)(unsafe.Pointer(c.pVertexAttributeDescriptions))[:len(g.pVertexAttributeDescriptions):len(g.pVertexAttributeDescriptions)]
		for i2, _ := range g.pVertexAttributeDescriptions {
			g.pVertexAttributeDescriptions[i2].fromC(&slice2[i2])
		}
	}
}

type VkPipelineInputAssemblyStateCreateFlags VkFlags
type VkPrimitiveTopology int

const (
	VK_PRIMITIVE_TOPOLOGY_POINT_LIST                    VkPrimitiveTopology = 0
	VK_PRIMITIVE_TOPOLOGY_LINE_LIST                     VkPrimitiveTopology = 1
	VK_PRIMITIVE_TOPOLOGY_LINE_STRIP                    VkPrimitiveTopology = 2
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST                 VkPrimitiveTopology = 3
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP                VkPrimitiveTopology = 4
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN                  VkPrimitiveTopology = 5
	VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY      VkPrimitiveTopology = 6
	VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY     VkPrimitiveTopology = 7
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY  VkPrimitiveTopology = 8
	VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY VkPrimitiveTopology = 9
	VK_PRIMITIVE_TOPOLOGY_PATCH_LIST                    VkPrimitiveTopology = 10
	VK_PRIMITIVE_TOPOLOGY_BEGIN_RANGE                   VkPrimitiveTopology = VK_PRIMITIVE_TOPOLOGY_POINT_LIST
	VK_PRIMITIVE_TOPOLOGY_END_RANGE                     VkPrimitiveTopology = VK_PRIMITIVE_TOPOLOGY_PATCH_LIST
	VK_PRIMITIVE_TOPOLOGY_RANGE_SIZE                    VkPrimitiveTopology = (VK_PRIMITIVE_TOPOLOGY_PATCH_LIST - VK_PRIMITIVE_TOPOLOGY_POINT_LIST + 1)
	VK_PRIMITIVE_TOPOLOGY_MAX_ENUM                      VkPrimitiveTopology = 2147483647
)

type VkPipelineInputAssemblyStateCreateInfo struct {
	sType                  VkStructureType
	pNext                  *[0]byte
	flags                  VkPipelineInputAssemblyStateCreateFlags
	topology               VkPrimitiveTopology
	primitiveRestartEnable VkBool32
}

func (g *VkPipelineInputAssemblyStateCreateInfo) toC(c *C.VkPipelineInputAssemblyStateCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineInputAssemblyStateCreateFlags(_temp)
	}
	c.topology = C.VkPrimitiveTopology(g.topology)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.primitiveRestartEnable))
		c.primitiveRestartEnable = C.VkBool32(_temp)
	}
}
func (g *VkPipelineInputAssemblyStateCreateInfo) fromC(c *C.VkPipelineInputAssemblyStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineInputAssemblyStateCreateFlags(_temp)
	}
	g.topology = VkPrimitiveTopology(c.topology)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.primitiveRestartEnable))
		g.primitiveRestartEnable = VkBool32(_temp)
	}
}

type VkPipelineTessellationStateCreateFlags VkFlags
type VkPipelineTessellationStateCreateInfo struct {
	sType              VkStructureType
	pNext              *[0]byte
	flags              VkPipelineTessellationStateCreateFlags
	patchControlPoints uint32
}

func (g *VkPipelineTessellationStateCreateInfo) toC(c *C.VkPipelineTessellationStateCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineTessellationStateCreateFlags(_temp)
	}
	c.patchControlPoints = C.uint32_t(g.patchControlPoints)
}
func (g *VkPipelineTessellationStateCreateInfo) fromC(c *C.VkPipelineTessellationStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineTessellationStateCreateFlags(_temp)
	}
	g.patchControlPoints = uint32(c.patchControlPoints)
}

type VkPipelineViewportStateCreateFlags VkFlags
type VkViewport struct {
	x        float32
	y        float32
	width    float32
	height   float32
	minDepth float32
	maxDepth float32
}

func (g *VkViewport) toC(c *C.VkViewport) {
	c.x = C.float(g.x)
	c.y = C.float(g.y)
	c.width = C.float(g.width)
	c.height = C.float(g.height)
	c.minDepth = C.float(g.minDepth)
	c.maxDepth = C.float(g.maxDepth)
}
func (g *VkViewport) fromC(c *C.VkViewport) {
	g.x = float32(c.x)
	g.y = float32(c.y)
	g.width = float32(c.width)
	g.height = float32(c.height)
	g.minDepth = float32(c.minDepth)
	g.maxDepth = float32(c.maxDepth)
}

type VkOffset2D struct {
	x int32
	y int32
}

func (g *VkOffset2D) toC(c *C.VkOffset2D) {
	c.x = C.int32_t(g.x)
	c.y = C.int32_t(g.y)
}
func (g *VkOffset2D) fromC(c *C.VkOffset2D) {
	g.x = int32(c.x)
	g.y = int32(c.y)
}

type VkExtent2D struct {
	width  uint32
	height uint32
}

func (g *VkExtent2D) toC(c *C.VkExtent2D) {
	c.width = C.uint32_t(g.width)
	c.height = C.uint32_t(g.height)
}
func (g *VkExtent2D) fromC(c *C.VkExtent2D) {
	g.width = uint32(c.width)
	g.height = uint32(c.height)
}

type VkRect2D struct {
	offset VkOffset2D
	extent VkExtent2D
}

func (g *VkRect2D) toC(c *C.VkRect2D) {
	g.offset.toC(&c.offset)
	g.extent.toC(&c.extent)
}
func (g *VkRect2D) fromC(c *C.VkRect2D) {
	g.offset.fromC(&c.offset)
	g.extent.fromC(&c.extent)
}

type VkPipelineViewportStateCreateInfo struct {
	sType      VkStructureType
	pNext      *[0]byte
	flags      VkPipelineViewportStateCreateFlags
	pViewports []VkViewport
	pScissors  []VkRect2D
}

func (g *VkPipelineViewportStateCreateInfo) toC(c *C.VkPipelineViewportStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineViewportStateCreateFlags(_temp)
	}
	c.viewportCount = C.uint32_t(len(g.pViewports))
	{
		c.pViewports = (*C.VkViewport)(_sa.alloc(C.sizeof_VkViewport * uint(len(g.pViewports))))
		slice2 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(g.pViewports):len(g.pViewports)]
		for i2, _ := range g.pViewports {
			g.pViewports[i2].toC(&slice2[i2])
		}
	}
	c.scissorCount = C.uint32_t(len(g.pScissors))
	{
		c.pScissors = (*C.VkRect2D)(_sa.alloc(C.sizeof_VkRect2D * uint(len(g.pScissors))))
		slice2 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(g.pScissors):len(g.pScissors)]
		for i2, _ := range g.pScissors {
			g.pScissors[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkPipelineViewportStateCreateInfo) fromC(c *C.VkPipelineViewportStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineViewportStateCreateFlags(_temp)
	}
	g.pViewports = make([]VkViewport, int(c.viewportCount))
	{
		slice2 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(g.pViewports):len(g.pViewports)]
		for i2, _ := range g.pViewports {
			g.pViewports[i2].fromC(&slice2[i2])
		}
	}
	g.pScissors = make([]VkRect2D, int(c.scissorCount))
	{
		slice2 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(g.pScissors):len(g.pScissors)]
		for i2, _ := range g.pScissors {
			g.pScissors[i2].fromC(&slice2[i2])
		}
	}
}

type VkPipelineRasterizationStateCreateFlags VkFlags
type VkPolygonMode int

const (
	VK_POLYGON_MODE_FILL              VkPolygonMode = 0
	VK_POLYGON_MODE_LINE              VkPolygonMode = 1
	VK_POLYGON_MODE_POINT             VkPolygonMode = 2
	VK_POLYGON_MODE_FILL_RECTANGLE_NV VkPolygonMode = 1000153000
	VK_POLYGON_MODE_BEGIN_RANGE       VkPolygonMode = VK_POLYGON_MODE_FILL
	VK_POLYGON_MODE_END_RANGE         VkPolygonMode = VK_POLYGON_MODE_POINT
	VK_POLYGON_MODE_RANGE_SIZE        VkPolygonMode = (VK_POLYGON_MODE_POINT - VK_POLYGON_MODE_FILL + 1)
	VK_POLYGON_MODE_MAX_ENUM          VkPolygonMode = 2147483647
)

type VkCullModeFlags VkFlags
type VkFrontFace int

const (
	VK_FRONT_FACE_COUNTER_CLOCKWISE VkFrontFace = 0
	VK_FRONT_FACE_CLOCKWISE         VkFrontFace = 1
	VK_FRONT_FACE_BEGIN_RANGE       VkFrontFace = VK_FRONT_FACE_COUNTER_CLOCKWISE
	VK_FRONT_FACE_END_RANGE         VkFrontFace = VK_FRONT_FACE_CLOCKWISE
	VK_FRONT_FACE_RANGE_SIZE        VkFrontFace = (VK_FRONT_FACE_CLOCKWISE - VK_FRONT_FACE_COUNTER_CLOCKWISE + 1)
	VK_FRONT_FACE_MAX_ENUM          VkFrontFace = 2147483647
)

type VkPipelineRasterizationStateCreateInfo struct {
	sType                   VkStructureType
	pNext                   *[0]byte
	flags                   VkPipelineRasterizationStateCreateFlags
	depthClampEnable        VkBool32
	rasterizerDiscardEnable VkBool32
	polygonMode             VkPolygonMode
	cullMode                VkCullModeFlags
	frontFace               VkFrontFace
	depthBiasEnable         VkBool32
	depthBiasConstantFactor float32
	depthBiasClamp          float32
	depthBiasSlopeFactor    float32
	lineWidth               float32
}

func (g *VkPipelineRasterizationStateCreateInfo) toC(c *C.VkPipelineRasterizationStateCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineRasterizationStateCreateFlags(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthClampEnable))
		c.depthClampEnable = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.rasterizerDiscardEnable))
		c.rasterizerDiscardEnable = C.VkBool32(_temp)
	}
	c.polygonMode = C.VkPolygonMode(g.polygonMode)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.cullMode)))
			_temp = C.VkFlags(_temp)
		}
		c.cullMode = C.VkCullModeFlags(_temp)
	}
	c.frontFace = C.VkFrontFace(g.frontFace)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthBiasEnable))
		c.depthBiasEnable = C.VkBool32(_temp)
	}
	c.depthBiasConstantFactor = C.float(g.depthBiasConstantFactor)
	c.depthBiasClamp = C.float(g.depthBiasClamp)
	c.depthBiasSlopeFactor = C.float(g.depthBiasSlopeFactor)
	c.lineWidth = C.float(g.lineWidth)
}
func (g *VkPipelineRasterizationStateCreateInfo) fromC(c *C.VkPipelineRasterizationStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineRasterizationStateCreateFlags(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthClampEnable))
		g.depthClampEnable = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.rasterizerDiscardEnable))
		g.rasterizerDiscardEnable = VkBool32(_temp)
	}
	g.polygonMode = VkPolygonMode(c.polygonMode)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.cullMode)))
			_temp = VkFlags(_temp)
		}
		g.cullMode = VkCullModeFlags(_temp)
	}
	g.frontFace = VkFrontFace(c.frontFace)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthBiasEnable))
		g.depthBiasEnable = VkBool32(_temp)
	}
	g.depthBiasConstantFactor = float32(c.depthBiasConstantFactor)
	g.depthBiasClamp = float32(c.depthBiasClamp)
	g.depthBiasSlopeFactor = float32(c.depthBiasSlopeFactor)
	g.lineWidth = float32(c.lineWidth)
}

type VkPipelineMultisampleStateCreateFlags VkFlags
type VkSampleMask uint32
type VkPipelineMultisampleStateCreateInfo struct {
	sType                 VkStructureType
	pNext                 *[0]byte
	flags                 VkPipelineMultisampleStateCreateFlags
	rasterizationSamples  VkSampleCountFlagBits
	sampleShadingEnable   VkBool32
	minSampleShading      float32
	pSampleMask           *VkSampleMask
	alphaToCoverageEnable VkBool32
	alphaToOneEnable      VkBool32
}

func (g *VkPipelineMultisampleStateCreateInfo) toC(c *C.VkPipelineMultisampleStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineMultisampleStateCreateFlags(_temp)
	}
	c.rasterizationSamples = C.VkSampleCountFlagBits(g.rasterizationSamples)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.sampleShadingEnable))
		c.sampleShadingEnable = C.VkBool32(_temp)
	}
	c.minSampleShading = C.float(g.minSampleShading)
	{
		c.pSampleMask = (*C.VkSampleMask)(_sa.alloc(C.sizeof_VkSampleMask))
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)(*g.pSampleMask))
			*c.pSampleMask = C.VkSampleMask(_temp)
		}
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.alphaToCoverageEnable))
		c.alphaToCoverageEnable = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.alphaToOneEnable))
		c.alphaToOneEnable = C.VkBool32(_temp)
	}
}
func (g *VkPipelineMultisampleStateCreateInfo) fromC(c *C.VkPipelineMultisampleStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineMultisampleStateCreateFlags(_temp)
	}
	g.rasterizationSamples = VkSampleCountFlagBits(c.rasterizationSamples)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.sampleShadingEnable))
		g.sampleShadingEnable = VkBool32(_temp)
	}
	g.minSampleShading = float32(c.minSampleShading)
	{
		if g.pSampleMask == nil {
			g.pSampleMask = new(VkSampleMask)
		}
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)(*c.pSampleMask))
			*g.pSampleMask = VkSampleMask(_temp)
		}
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.alphaToCoverageEnable))
		g.alphaToCoverageEnable = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.alphaToOneEnable))
		g.alphaToOneEnable = VkBool32(_temp)
	}
}

type VkPipelineDepthStencilStateCreateFlags VkFlags
type VkCompareOp int

const (
	VK_COMPARE_OP_NEVER            VkCompareOp = 0
	VK_COMPARE_OP_LESS             VkCompareOp = 1
	VK_COMPARE_OP_EQUAL            VkCompareOp = 2
	VK_COMPARE_OP_LESS_OR_EQUAL    VkCompareOp = 3
	VK_COMPARE_OP_GREATER          VkCompareOp = 4
	VK_COMPARE_OP_NOT_EQUAL        VkCompareOp = 5
	VK_COMPARE_OP_GREATER_OR_EQUAL VkCompareOp = 6
	VK_COMPARE_OP_ALWAYS           VkCompareOp = 7
	VK_COMPARE_OP_BEGIN_RANGE      VkCompareOp = VK_COMPARE_OP_NEVER
	VK_COMPARE_OP_END_RANGE        VkCompareOp = VK_COMPARE_OP_ALWAYS
	VK_COMPARE_OP_RANGE_SIZE       VkCompareOp = (VK_COMPARE_OP_ALWAYS - VK_COMPARE_OP_NEVER + 1)
	VK_COMPARE_OP_MAX_ENUM         VkCompareOp = 2147483647
)

type VkStencilOp int

const (
	VK_STENCIL_OP_KEEP                VkStencilOp = 0
	VK_STENCIL_OP_ZERO                VkStencilOp = 1
	VK_STENCIL_OP_REPLACE             VkStencilOp = 2
	VK_STENCIL_OP_INCREMENT_AND_CLAMP VkStencilOp = 3
	VK_STENCIL_OP_DECREMENT_AND_CLAMP VkStencilOp = 4
	VK_STENCIL_OP_INVERT              VkStencilOp = 5
	VK_STENCIL_OP_INCREMENT_AND_WRAP  VkStencilOp = 6
	VK_STENCIL_OP_DECREMENT_AND_WRAP  VkStencilOp = 7
	VK_STENCIL_OP_BEGIN_RANGE         VkStencilOp = VK_STENCIL_OP_KEEP
	VK_STENCIL_OP_END_RANGE           VkStencilOp = VK_STENCIL_OP_DECREMENT_AND_WRAP
	VK_STENCIL_OP_RANGE_SIZE          VkStencilOp = (VK_STENCIL_OP_DECREMENT_AND_WRAP - VK_STENCIL_OP_KEEP + 1)
	VK_STENCIL_OP_MAX_ENUM            VkStencilOp = 2147483647
)

type VkStencilOpState struct {
	failOp      VkStencilOp
	passOp      VkStencilOp
	depthFailOp VkStencilOp
	compareOp   VkCompareOp
	compareMask uint32
	writeMask   uint32
	reference   uint32
}

func (g *VkStencilOpState) toC(c *C.VkStencilOpState) {
	c.failOp = C.VkStencilOp(g.failOp)
	c.passOp = C.VkStencilOp(g.passOp)
	c.depthFailOp = C.VkStencilOp(g.depthFailOp)
	c.compareOp = C.VkCompareOp(g.compareOp)
	c.compareMask = C.uint32_t(g.compareMask)
	c.writeMask = C.uint32_t(g.writeMask)
	c.reference = C.uint32_t(g.reference)
}
func (g *VkStencilOpState) fromC(c *C.VkStencilOpState) {
	g.failOp = VkStencilOp(c.failOp)
	g.passOp = VkStencilOp(c.passOp)
	g.depthFailOp = VkStencilOp(c.depthFailOp)
	g.compareOp = VkCompareOp(c.compareOp)
	g.compareMask = uint32(c.compareMask)
	g.writeMask = uint32(c.writeMask)
	g.reference = uint32(c.reference)
}

type VkPipelineDepthStencilStateCreateInfo struct {
	sType                 VkStructureType
	pNext                 *[0]byte
	flags                 VkPipelineDepthStencilStateCreateFlags
	depthTestEnable       VkBool32
	depthWriteEnable      VkBool32
	depthCompareOp        VkCompareOp
	depthBoundsTestEnable VkBool32
	stencilTestEnable     VkBool32
	front                 VkStencilOpState
	back                  VkStencilOpState
	minDepthBounds        float32
	maxDepthBounds        float32
}

func (g *VkPipelineDepthStencilStateCreateInfo) toC(c *C.VkPipelineDepthStencilStateCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineDepthStencilStateCreateFlags(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthTestEnable))
		c.depthTestEnable = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthWriteEnable))
		c.depthWriteEnable = C.VkBool32(_temp)
	}
	c.depthCompareOp = C.VkCompareOp(g.depthCompareOp)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.depthBoundsTestEnable))
		c.depthBoundsTestEnable = C.VkBool32(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.stencilTestEnable))
		c.stencilTestEnable = C.VkBool32(_temp)
	}
	g.front.toC(&c.front)
	g.back.toC(&c.back)
	c.minDepthBounds = C.float(g.minDepthBounds)
	c.maxDepthBounds = C.float(g.maxDepthBounds)
}
func (g *VkPipelineDepthStencilStateCreateInfo) fromC(c *C.VkPipelineDepthStencilStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineDepthStencilStateCreateFlags(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthTestEnable))
		g.depthTestEnable = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthWriteEnable))
		g.depthWriteEnable = VkBool32(_temp)
	}
	g.depthCompareOp = VkCompareOp(c.depthCompareOp)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.depthBoundsTestEnable))
		g.depthBoundsTestEnable = VkBool32(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.stencilTestEnable))
		g.stencilTestEnable = VkBool32(_temp)
	}
	g.front.fromC(&c.front)
	g.back.fromC(&c.back)
	g.minDepthBounds = float32(c.minDepthBounds)
	g.maxDepthBounds = float32(c.maxDepthBounds)
}

type VkPipelineColorBlendStateCreateFlags VkFlags
type VkLogicOp int

const (
	VK_LOGIC_OP_CLEAR         VkLogicOp = 0
	VK_LOGIC_OP_AND           VkLogicOp = 1
	VK_LOGIC_OP_AND_REVERSE   VkLogicOp = 2
	VK_LOGIC_OP_COPY          VkLogicOp = 3
	VK_LOGIC_OP_AND_INVERTED  VkLogicOp = 4
	VK_LOGIC_OP_NO_OP         VkLogicOp = 5
	VK_LOGIC_OP_XOR           VkLogicOp = 6
	VK_LOGIC_OP_OR            VkLogicOp = 7
	VK_LOGIC_OP_NOR           VkLogicOp = 8
	VK_LOGIC_OP_EQUIVALENT    VkLogicOp = 9
	VK_LOGIC_OP_INVERT        VkLogicOp = 10
	VK_LOGIC_OP_OR_REVERSE    VkLogicOp = 11
	VK_LOGIC_OP_COPY_INVERTED VkLogicOp = 12
	VK_LOGIC_OP_OR_INVERTED   VkLogicOp = 13
	VK_LOGIC_OP_NAND          VkLogicOp = 14
	VK_LOGIC_OP_SET           VkLogicOp = 15
	VK_LOGIC_OP_BEGIN_RANGE   VkLogicOp = VK_LOGIC_OP_CLEAR
	VK_LOGIC_OP_END_RANGE     VkLogicOp = VK_LOGIC_OP_SET
	VK_LOGIC_OP_RANGE_SIZE    VkLogicOp = (VK_LOGIC_OP_SET - VK_LOGIC_OP_CLEAR + 1)
	VK_LOGIC_OP_MAX_ENUM      VkLogicOp = 2147483647
)

type VkBlendFactor int

const (
	VK_BLEND_FACTOR_ZERO                     VkBlendFactor = 0
	VK_BLEND_FACTOR_ONE                      VkBlendFactor = 1
	VK_BLEND_FACTOR_SRC_COLOR                VkBlendFactor = 2
	VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR      VkBlendFactor = 3
	VK_BLEND_FACTOR_DST_COLOR                VkBlendFactor = 4
	VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR      VkBlendFactor = 5
	VK_BLEND_FACTOR_SRC_ALPHA                VkBlendFactor = 6
	VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA      VkBlendFactor = 7
	VK_BLEND_FACTOR_DST_ALPHA                VkBlendFactor = 8
	VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA      VkBlendFactor = 9
	VK_BLEND_FACTOR_CONSTANT_COLOR           VkBlendFactor = 10
	VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR VkBlendFactor = 11
	VK_BLEND_FACTOR_CONSTANT_ALPHA           VkBlendFactor = 12
	VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA VkBlendFactor = 13
	VK_BLEND_FACTOR_SRC_ALPHA_SATURATE       VkBlendFactor = 14
	VK_BLEND_FACTOR_SRC1_COLOR               VkBlendFactor = 15
	VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR     VkBlendFactor = 16
	VK_BLEND_FACTOR_SRC1_ALPHA               VkBlendFactor = 17
	VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA     VkBlendFactor = 18
	VK_BLEND_FACTOR_BEGIN_RANGE              VkBlendFactor = VK_BLEND_FACTOR_ZERO
	VK_BLEND_FACTOR_END_RANGE                VkBlendFactor = VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
	VK_BLEND_FACTOR_RANGE_SIZE               VkBlendFactor = (VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA - VK_BLEND_FACTOR_ZERO + 1)
	VK_BLEND_FACTOR_MAX_ENUM                 VkBlendFactor = 2147483647
)

type VkBlendOp int

const (
	VK_BLEND_OP_ADD                    VkBlendOp = 0
	VK_BLEND_OP_SUBTRACT               VkBlendOp = 1
	VK_BLEND_OP_REVERSE_SUBTRACT       VkBlendOp = 2
	VK_BLEND_OP_MIN                    VkBlendOp = 3
	VK_BLEND_OP_MAX                    VkBlendOp = 4
	VK_BLEND_OP_ZERO_EXT               VkBlendOp = 1000148000
	VK_BLEND_OP_SRC_EXT                VkBlendOp = 1000148001
	VK_BLEND_OP_DST_EXT                VkBlendOp = 1000148002
	VK_BLEND_OP_SRC_OVER_EXT           VkBlendOp = 1000148003
	VK_BLEND_OP_DST_OVER_EXT           VkBlendOp = 1000148004
	VK_BLEND_OP_SRC_IN_EXT             VkBlendOp = 1000148005
	VK_BLEND_OP_DST_IN_EXT             VkBlendOp = 1000148006
	VK_BLEND_OP_SRC_OUT_EXT            VkBlendOp = 1000148007
	VK_BLEND_OP_DST_OUT_EXT            VkBlendOp = 1000148008
	VK_BLEND_OP_SRC_ATOP_EXT           VkBlendOp = 1000148009
	VK_BLEND_OP_DST_ATOP_EXT           VkBlendOp = 1000148010
	VK_BLEND_OP_XOR_EXT                VkBlendOp = 1000148011
	VK_BLEND_OP_MULTIPLY_EXT           VkBlendOp = 1000148012
	VK_BLEND_OP_SCREEN_EXT             VkBlendOp = 1000148013
	VK_BLEND_OP_OVERLAY_EXT            VkBlendOp = 1000148014
	VK_BLEND_OP_DARKEN_EXT             VkBlendOp = 1000148015
	VK_BLEND_OP_LIGHTEN_EXT            VkBlendOp = 1000148016
	VK_BLEND_OP_COLORDODGE_EXT         VkBlendOp = 1000148017
	VK_BLEND_OP_COLORBURN_EXT          VkBlendOp = 1000148018
	VK_BLEND_OP_HARDLIGHT_EXT          VkBlendOp = 1000148019
	VK_BLEND_OP_SOFTLIGHT_EXT          VkBlendOp = 1000148020
	VK_BLEND_OP_DIFFERENCE_EXT         VkBlendOp = 1000148021
	VK_BLEND_OP_EXCLUSION_EXT          VkBlendOp = 1000148022
	VK_BLEND_OP_INVERT_EXT             VkBlendOp = 1000148023
	VK_BLEND_OP_INVERT_RGB_EXT         VkBlendOp = 1000148024
	VK_BLEND_OP_LINEARDODGE_EXT        VkBlendOp = 1000148025
	VK_BLEND_OP_LINEARBURN_EXT         VkBlendOp = 1000148026
	VK_BLEND_OP_VIVIDLIGHT_EXT         VkBlendOp = 1000148027
	VK_BLEND_OP_LINEARLIGHT_EXT        VkBlendOp = 1000148028
	VK_BLEND_OP_PINLIGHT_EXT           VkBlendOp = 1000148029
	VK_BLEND_OP_HARDMIX_EXT            VkBlendOp = 1000148030
	VK_BLEND_OP_HSL_HUE_EXT            VkBlendOp = 1000148031
	VK_BLEND_OP_HSL_SATURATION_EXT     VkBlendOp = 1000148032
	VK_BLEND_OP_HSL_COLOR_EXT          VkBlendOp = 1000148033
	VK_BLEND_OP_HSL_LUMINOSITY_EXT     VkBlendOp = 1000148034
	VK_BLEND_OP_PLUS_EXT               VkBlendOp = 1000148035
	VK_BLEND_OP_PLUS_CLAMPED_EXT       VkBlendOp = 1000148036
	VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT VkBlendOp = 1000148037
	VK_BLEND_OP_PLUS_DARKER_EXT        VkBlendOp = 1000148038
	VK_BLEND_OP_MINUS_EXT              VkBlendOp = 1000148039
	VK_BLEND_OP_MINUS_CLAMPED_EXT      VkBlendOp = 1000148040
	VK_BLEND_OP_CONTRAST_EXT           VkBlendOp = 1000148041
	VK_BLEND_OP_INVERT_OVG_EXT         VkBlendOp = 1000148042
	VK_BLEND_OP_RED_EXT                VkBlendOp = 1000148043
	VK_BLEND_OP_GREEN_EXT              VkBlendOp = 1000148044
	VK_BLEND_OP_BLUE_EXT               VkBlendOp = 1000148045
	VK_BLEND_OP_BEGIN_RANGE            VkBlendOp = VK_BLEND_OP_ADD
	VK_BLEND_OP_END_RANGE              VkBlendOp = VK_BLEND_OP_MAX
	VK_BLEND_OP_RANGE_SIZE             VkBlendOp = (VK_BLEND_OP_MAX - VK_BLEND_OP_ADD + 1)
	VK_BLEND_OP_MAX_ENUM               VkBlendOp = 2147483647
)

type VkColorComponentFlags VkFlags
type VkPipelineColorBlendAttachmentState struct {
	blendEnable         VkBool32
	srcColorBlendFactor VkBlendFactor
	dstColorBlendFactor VkBlendFactor
	colorBlendOp        VkBlendOp
	srcAlphaBlendFactor VkBlendFactor
	dstAlphaBlendFactor VkBlendFactor
	alphaBlendOp        VkBlendOp
	colorWriteMask      VkColorComponentFlags
}

func (g *VkPipelineColorBlendAttachmentState) toC(c *C.VkPipelineColorBlendAttachmentState) {
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.blendEnable))
		c.blendEnable = C.VkBool32(_temp)
	}
	c.srcColorBlendFactor = C.VkBlendFactor(g.srcColorBlendFactor)
	c.dstColorBlendFactor = C.VkBlendFactor(g.dstColorBlendFactor)
	c.colorBlendOp = C.VkBlendOp(g.colorBlendOp)
	c.srcAlphaBlendFactor = C.VkBlendFactor(g.srcAlphaBlendFactor)
	c.dstAlphaBlendFactor = C.VkBlendFactor(g.dstAlphaBlendFactor)
	c.alphaBlendOp = C.VkBlendOp(g.alphaBlendOp)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.colorWriteMask)))
			_temp = C.VkFlags(_temp)
		}
		c.colorWriteMask = C.VkColorComponentFlags(_temp)
	}
}
func (g *VkPipelineColorBlendAttachmentState) fromC(c *C.VkPipelineColorBlendAttachmentState) {
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.blendEnable))
		g.blendEnable = VkBool32(_temp)
	}
	g.srcColorBlendFactor = VkBlendFactor(c.srcColorBlendFactor)
	g.dstColorBlendFactor = VkBlendFactor(c.dstColorBlendFactor)
	g.colorBlendOp = VkBlendOp(c.colorBlendOp)
	g.srcAlphaBlendFactor = VkBlendFactor(c.srcAlphaBlendFactor)
	g.dstAlphaBlendFactor = VkBlendFactor(c.dstAlphaBlendFactor)
	g.alphaBlendOp = VkBlendOp(c.alphaBlendOp)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.colorWriteMask)))
			_temp = VkFlags(_temp)
		}
		g.colorWriteMask = VkColorComponentFlags(_temp)
	}
}

type VkPipelineColorBlendStateCreateInfo struct {
	sType          VkStructureType
	pNext          *[0]byte
	flags          VkPipelineColorBlendStateCreateFlags
	logicOpEnable  VkBool32
	logicOp        VkLogicOp
	pAttachments   []VkPipelineColorBlendAttachmentState
	blendConstants [4]float32
}

func (g *VkPipelineColorBlendStateCreateInfo) toC(c *C.VkPipelineColorBlendStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineColorBlendStateCreateFlags(_temp)
	}
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.logicOpEnable))
		c.logicOpEnable = C.VkBool32(_temp)
	}
	c.logicOp = C.VkLogicOp(g.logicOp)
	c.attachmentCount = C.uint32_t(len(g.pAttachments))
	{
		c.pAttachments = (*C.VkPipelineColorBlendAttachmentState)(_sa.alloc(C.sizeof_VkPipelineColorBlendAttachmentState * uint(len(g.pAttachments))))
		slice2 := (*[1 << 31]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			g.pAttachments[i2].toC(&slice2[i2])
		}
	}
	for i, _ := range g.blendConstants {
		c.blendConstants[i] = C.float(g.blendConstants[i])
	}
}
func (g *VkPipelineColorBlendStateCreateInfo) fromC(c *C.VkPipelineColorBlendStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineColorBlendStateCreateFlags(_temp)
	}
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.logicOpEnable))
		g.logicOpEnable = VkBool32(_temp)
	}
	g.logicOp = VkLogicOp(c.logicOp)
	g.pAttachments = make([]VkPipelineColorBlendAttachmentState, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkPipelineColorBlendAttachmentState)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			g.pAttachments[i2].fromC(&slice2[i2])
		}
	}
	for i, _ := range g.blendConstants {
		g.blendConstants[i] = float32(c.blendConstants[i])
	}
}

type VkPipelineDynamicStateCreateFlags VkFlags
type VkDynamicState int

const (
	VK_DYNAMIC_STATE_VIEWPORT              VkDynamicState = 0
	VK_DYNAMIC_STATE_SCISSOR               VkDynamicState = 1
	VK_DYNAMIC_STATE_LINE_WIDTH            VkDynamicState = 2
	VK_DYNAMIC_STATE_DEPTH_BIAS            VkDynamicState = 3
	VK_DYNAMIC_STATE_BLEND_CONSTANTS       VkDynamicState = 4
	VK_DYNAMIC_STATE_DEPTH_BOUNDS          VkDynamicState = 5
	VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK  VkDynamicState = 6
	VK_DYNAMIC_STATE_STENCIL_WRITE_MASK    VkDynamicState = 7
	VK_DYNAMIC_STATE_STENCIL_REFERENCE     VkDynamicState = 8
	VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV VkDynamicState = 1000087000
	VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT VkDynamicState = 1000099000
	VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT  VkDynamicState = 1000143000
	VK_DYNAMIC_STATE_BEGIN_RANGE           VkDynamicState = VK_DYNAMIC_STATE_VIEWPORT
	VK_DYNAMIC_STATE_END_RANGE             VkDynamicState = VK_DYNAMIC_STATE_STENCIL_REFERENCE
	VK_DYNAMIC_STATE_RANGE_SIZE            VkDynamicState = (VK_DYNAMIC_STATE_STENCIL_REFERENCE - VK_DYNAMIC_STATE_VIEWPORT + 1)
	VK_DYNAMIC_STATE_MAX_ENUM              VkDynamicState = 2147483647
)

type VkPipelineDynamicStateCreateInfo struct {
	sType          VkStructureType
	pNext          *[0]byte
	flags          VkPipelineDynamicStateCreateFlags
	pDynamicStates []VkDynamicState
}

func (g *VkPipelineDynamicStateCreateInfo) toC(c *C.VkPipelineDynamicStateCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineDynamicStateCreateFlags(_temp)
	}
	c.dynamicStateCount = C.uint32_t(len(g.pDynamicStates))
	{
		c.pDynamicStates = (*C.VkDynamicState)(_sa.alloc(C.sizeof_VkDynamicState * uint(len(g.pDynamicStates))))
		slice2 := (*[1 << 31]C.VkDynamicState)(unsafe.Pointer(c.pDynamicStates))[:len(g.pDynamicStates):len(g.pDynamicStates)]
		for i2, _ := range g.pDynamicStates {
			slice2[i2] = C.VkDynamicState(g.pDynamicStates[i2])
		}
	}
}
func (g *VkPipelineDynamicStateCreateInfo) fromC(c *C.VkPipelineDynamicStateCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineDynamicStateCreateFlags(_temp)
	}
	g.pDynamicStates = make([]VkDynamicState, int(c.dynamicStateCount))
	{
		slice2 := (*[1 << 31]C.VkDynamicState)(unsafe.Pointer(c.pDynamicStates))[:len(g.pDynamicStates):len(g.pDynamicStates)]
		for i2, _ := range g.pDynamicStates {
			g.pDynamicStates[i2] = VkDynamicState(slice2[i2])
		}
	}
}

type VkPipelineLayout C.VkPipelineLayout
type VkRenderPass C.VkRenderPass
type VkPipeline C.VkPipeline
type VkGraphicsPipelineCreateInfo struct {
	sType               VkStructureType
	pNext               *[0]byte
	flags               VkPipelineCreateFlags
	pStages             []VkPipelineShaderStageCreateInfo
	pVertexInputState   *VkPipelineVertexInputStateCreateInfo
	pInputAssemblyState *VkPipelineInputAssemblyStateCreateInfo
	pTessellationState  *VkPipelineTessellationStateCreateInfo
	pViewportState      *VkPipelineViewportStateCreateInfo
	pRasterizationState *VkPipelineRasterizationStateCreateInfo
	pMultisampleState   *VkPipelineMultisampleStateCreateInfo
	pDepthStencilState  *VkPipelineDepthStencilStateCreateInfo
	pColorBlendState    *VkPipelineColorBlendStateCreateInfo
	pDynamicState       *VkPipelineDynamicStateCreateInfo
	layout              VkPipelineLayout
	renderPass          VkRenderPass
	subpass             uint32
	basePipelineHandle  VkPipeline
	basePipelineIndex   int32
}

func (g *VkGraphicsPipelineCreateInfo) toC(c *C.VkGraphicsPipelineCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineCreateFlags(_temp)
	}
	c.stageCount = C.uint32_t(len(g.pStages))
	{
		c.pStages = (*C.VkPipelineShaderStageCreateInfo)(_sa.alloc(C.sizeof_VkPipelineShaderStageCreateInfo * uint(len(g.pStages))))
		slice2 := (*[1 << 31]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(c.pStages))[:len(g.pStages):len(g.pStages)]
		for i2, _ := range g.pStages {
			g.pStages[i2].toC(&slice2[i2], _sa)
		}
	}
	{
		c.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineVertexInputStateCreateInfo))
		g.pVertexInputState.toC(c.pVertexInputState, _sa)
	}
	{
		c.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineInputAssemblyStateCreateInfo))
		g.pInputAssemblyState.toC(c.pInputAssemblyState)
	}
	{
		c.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineTessellationStateCreateInfo))
		g.pTessellationState.toC(c.pTessellationState)
	}
	{
		c.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineViewportStateCreateInfo))
		g.pViewportState.toC(c.pViewportState, _sa)
	}
	{
		c.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineRasterizationStateCreateInfo))
		g.pRasterizationState.toC(c.pRasterizationState)
	}
	{
		c.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineMultisampleStateCreateInfo))
		g.pMultisampleState.toC(c.pMultisampleState, _sa)
	}
	{
		c.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineDepthStencilStateCreateInfo))
		g.pDepthStencilState.toC(c.pDepthStencilState)
	}
	{
		c.pColorBlendState = (*C.VkPipelineColorBlendStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineColorBlendStateCreateInfo))
		g.pColorBlendState.toC(c.pColorBlendState, _sa)
	}
	{
		c.pDynamicState = (*C.VkPipelineDynamicStateCreateInfo)(_sa.alloc(C.sizeof_VkPipelineDynamicStateCreateInfo))
		g.pDynamicState.toC(c.pDynamicState, _sa)
	}
	c.layout = C.VkPipelineLayout(g.layout)
	c.renderPass = C.VkRenderPass(g.renderPass)
	c.subpass = C.uint32_t(g.subpass)
	c.basePipelineHandle = C.VkPipeline(g.basePipelineHandle)
	c.basePipelineIndex = C.int32_t(g.basePipelineIndex)
}
func (g *VkGraphicsPipelineCreateInfo) fromC(c *C.VkGraphicsPipelineCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineCreateFlags(_temp)
	}
	g.pStages = make([]VkPipelineShaderStageCreateInfo, int(c.stageCount))
	{
		slice2 := (*[1 << 31]C.VkPipelineShaderStageCreateInfo)(unsafe.Pointer(c.pStages))[:len(g.pStages):len(g.pStages)]
		for i2, _ := range g.pStages {
			g.pStages[i2].fromC(&slice2[i2])
		}
	}
	{
		if g.pVertexInputState == nil {
			g.pVertexInputState = new(VkPipelineVertexInputStateCreateInfo)
		}
		g.pVertexInputState.fromC(c.pVertexInputState)
	}
	{
		if g.pInputAssemblyState == nil {
			g.pInputAssemblyState = new(VkPipelineInputAssemblyStateCreateInfo)
		}
		g.pInputAssemblyState.fromC(c.pInputAssemblyState)
	}
	{
		if g.pTessellationState == nil {
			g.pTessellationState = new(VkPipelineTessellationStateCreateInfo)
		}
		g.pTessellationState.fromC(c.pTessellationState)
	}
	{
		if g.pViewportState == nil {
			g.pViewportState = new(VkPipelineViewportStateCreateInfo)
		}
		g.pViewportState.fromC(c.pViewportState)
	}
	{
		if g.pRasterizationState == nil {
			g.pRasterizationState = new(VkPipelineRasterizationStateCreateInfo)
		}
		g.pRasterizationState.fromC(c.pRasterizationState)
	}
	{
		if g.pMultisampleState == nil {
			g.pMultisampleState = new(VkPipelineMultisampleStateCreateInfo)
		}
		g.pMultisampleState.fromC(c.pMultisampleState)
	}
	{
		if g.pDepthStencilState == nil {
			g.pDepthStencilState = new(VkPipelineDepthStencilStateCreateInfo)
		}
		g.pDepthStencilState.fromC(c.pDepthStencilState)
	}
	{
		if g.pColorBlendState == nil {
			g.pColorBlendState = new(VkPipelineColorBlendStateCreateInfo)
		}
		g.pColorBlendState.fromC(c.pColorBlendState)
	}
	{
		if g.pDynamicState == nil {
			g.pDynamicState = new(VkPipelineDynamicStateCreateInfo)
		}
		g.pDynamicState.fromC(c.pDynamicState)
	}
	g.layout = VkPipelineLayout(c.layout)
	g.renderPass = VkRenderPass(c.renderPass)
	g.subpass = uint32(c.subpass)
	g.basePipelineHandle = VkPipeline(c.basePipelineHandle)
	g.basePipelineIndex = int32(c.basePipelineIndex)
}
func vkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, pCreateInfos []VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines []VkPipeline) (_ret VkResult) {
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
	c.createInfoCount = C.uint32_t(len(pCreateInfos))
	{
		c.pCreateInfos = (*C.VkGraphicsPipelineCreateInfo)(_sa.alloc(C.sizeof_VkGraphicsPipelineCreateInfo * uint(len(pCreateInfos))))
		slice3 := (*[1 << 31]C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(c.pCreateInfos))[:len(pCreateInfos):len(pCreateInfos)]
		for i3, _ := range pCreateInfos {
			pCreateInfos[i3].toC(&slice3[i3], _sa)
		}
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelines = (*C.VkPipeline)(_sa.alloc(C.sizeof_VkPipeline * uint(len(pPipelines))))
		slice3 := (*[1 << 31]C.VkPipeline)(unsafe.Pointer(c.pPipelines))[:len(pPipelines):len(pPipelines)]
		for i3, _ := range pPipelines {
			slice3[i3] = C.VkPipeline(pPipelines[i3])
		}
	}
	c._ret = C.vkCreateGraphicsPipelines(c.device, c.pipelineCache, c.createInfoCount, c.pCreateInfos, c.pAllocator, c.pPipelines)
	_ret = VkResult(c._ret)
	return
}

type VkComputePipelineCreateInfo struct {
	sType              VkStructureType
	pNext              *[0]byte
	flags              VkPipelineCreateFlags
	stage              VkPipelineShaderStageCreateInfo
	layout             VkPipelineLayout
	basePipelineHandle VkPipeline
	basePipelineIndex  int32
}

func (g *VkComputePipelineCreateInfo) toC(c *C.VkComputePipelineCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineCreateFlags(_temp)
	}
	g.stage.toC(&c.stage, _sa)
	c.layout = C.VkPipelineLayout(g.layout)
	c.basePipelineHandle = C.VkPipeline(g.basePipelineHandle)
	c.basePipelineIndex = C.int32_t(g.basePipelineIndex)
}
func (g *VkComputePipelineCreateInfo) fromC(c *C.VkComputePipelineCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineCreateFlags(_temp)
	}
	g.stage.fromC(&c.stage)
	g.layout = VkPipelineLayout(c.layout)
	g.basePipelineHandle = VkPipeline(c.basePipelineHandle)
	g.basePipelineIndex = int32(c.basePipelineIndex)
}
func vkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, pCreateInfos []VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines []VkPipeline) (_ret VkResult) {
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
	c.createInfoCount = C.uint32_t(len(pCreateInfos))
	{
		c.pCreateInfos = (*C.VkComputePipelineCreateInfo)(_sa.alloc(C.sizeof_VkComputePipelineCreateInfo * uint(len(pCreateInfos))))
		slice3 := (*[1 << 31]C.VkComputePipelineCreateInfo)(unsafe.Pointer(c.pCreateInfos))[:len(pCreateInfos):len(pCreateInfos)]
		for i3, _ := range pCreateInfos {
			pCreateInfos[i3].toC(&slice3[i3], _sa)
		}
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelines = (*C.VkPipeline)(_sa.alloc(C.sizeof_VkPipeline * uint(len(pPipelines))))
		slice3 := (*[1 << 31]C.VkPipeline)(unsafe.Pointer(c.pPipelines))[:len(pPipelines):len(pPipelines)]
		for i3, _ := range pPipelines {
			slice3[i3] = C.VkPipeline(pPipelines[i3])
		}
	}
	c._ret = C.vkCreateComputePipelines(c.device, c.pipelineCache, c.createInfoCount, c.pCreateInfos, c.pAllocator, c.pPipelines)
	_ret = VkResult(c._ret)
	return
}
func vkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipeline(c.device, c.pipeline, c.pAllocator)
}

type VkPipelineLayoutCreateFlags VkFlags
type VkDescriptorSetLayout C.VkDescriptorSetLayout
type VkShaderStageFlags VkFlags
type VkPushConstantRange struct {
	stageFlags VkShaderStageFlags
	offset     uint32
	size       uint32
}

func (g *VkPushConstantRange) toC(c *C.VkPushConstantRange) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.stageFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.stageFlags = C.VkShaderStageFlags(_temp)
	}
	c.offset = C.uint32_t(g.offset)
	c.size = C.uint32_t(g.size)
}
func (g *VkPushConstantRange) fromC(c *C.VkPushConstantRange) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.stageFlags)))
			_temp = VkFlags(_temp)
		}
		g.stageFlags = VkShaderStageFlags(_temp)
	}
	g.offset = uint32(c.offset)
	g.size = uint32(c.size)
}

type VkPipelineLayoutCreateInfo struct {
	sType               VkStructureType
	pNext               *[0]byte
	flags               VkPipelineLayoutCreateFlags
	pSetLayouts         []VkDescriptorSetLayout
	pPushConstantRanges []VkPushConstantRange
}

func (g *VkPipelineLayoutCreateInfo) toC(c *C.VkPipelineLayoutCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkPipelineLayoutCreateFlags(_temp)
	}
	c.setLayoutCount = C.uint32_t(len(g.pSetLayouts))
	{
		c.pSetLayouts = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout * uint(len(g.pSetLayouts))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.pSetLayouts):len(g.pSetLayouts)]
		for i2, _ := range g.pSetLayouts {
			slice2[i2] = C.VkDescriptorSetLayout(g.pSetLayouts[i2])
		}
	}
	c.pushConstantRangeCount = C.uint32_t(len(g.pPushConstantRanges))
	{
		c.pPushConstantRanges = (*C.VkPushConstantRange)(_sa.alloc(C.sizeof_VkPushConstantRange * uint(len(g.pPushConstantRanges))))
		slice2 := (*[1 << 31]C.VkPushConstantRange)(unsafe.Pointer(c.pPushConstantRanges))[:len(g.pPushConstantRanges):len(g.pPushConstantRanges)]
		for i2, _ := range g.pPushConstantRanges {
			g.pPushConstantRanges[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkPipelineLayoutCreateInfo) fromC(c *C.VkPipelineLayoutCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkPipelineLayoutCreateFlags(_temp)
	}
	g.pSetLayouts = make([]VkDescriptorSetLayout, int(c.setLayoutCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.pSetLayouts):len(g.pSetLayouts)]
		for i2, _ := range g.pSetLayouts {
			g.pSetLayouts[i2] = VkDescriptorSetLayout(slice2[i2])
		}
	}
	g.pPushConstantRanges = make([]VkPushConstantRange, int(c.pushConstantRangeCount))
	{
		slice2 := (*[1 << 31]C.VkPushConstantRange)(unsafe.Pointer(c.pPushConstantRanges))[:len(g.pPushConstantRanges):len(g.pPushConstantRanges)]
		for i2, _ := range g.pPushConstantRanges {
			g.pPushConstantRanges[i2].fromC(&slice2[i2])
		}
	}
}
func vkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pPipelineLayout = (*C.VkPipelineLayout)(_sa.alloc(C.sizeof_VkPipelineLayout))
		*c.pPipelineLayout = C.VkPipelineLayout(*pPipelineLayout)
	}
	c._ret = C.vkCreatePipelineLayout(c.device, c.pCreateInfo, c.pAllocator, c.pPipelineLayout)
	_ret = VkResult(c._ret)
	*pPipelineLayout = VkPipelineLayout(*c.pPipelineLayout)
	return
}
func vkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyPipelineLayout(c.device, c.pipelineLayout, c.pAllocator)
}

type VkSamplerCreateFlags VkFlags
type VkFilter int

const (
	VK_FILTER_NEAREST     VkFilter = 0
	VK_FILTER_LINEAR      VkFilter = 1
	VK_FILTER_CUBIC_IMG   VkFilter = 1000015000
	VK_FILTER_BEGIN_RANGE VkFilter = VK_FILTER_NEAREST
	VK_FILTER_END_RANGE   VkFilter = VK_FILTER_LINEAR
	VK_FILTER_RANGE_SIZE  VkFilter = (VK_FILTER_LINEAR - VK_FILTER_NEAREST + 1)
	VK_FILTER_MAX_ENUM    VkFilter = 2147483647
)

type VkSamplerMipmapMode int

const (
	VK_SAMPLER_MIPMAP_MODE_NEAREST     VkSamplerMipmapMode = 0
	VK_SAMPLER_MIPMAP_MODE_LINEAR      VkSamplerMipmapMode = 1
	VK_SAMPLER_MIPMAP_MODE_BEGIN_RANGE VkSamplerMipmapMode = VK_SAMPLER_MIPMAP_MODE_NEAREST
	VK_SAMPLER_MIPMAP_MODE_END_RANGE   VkSamplerMipmapMode = VK_SAMPLER_MIPMAP_MODE_LINEAR
	VK_SAMPLER_MIPMAP_MODE_RANGE_SIZE  VkSamplerMipmapMode = (VK_SAMPLER_MIPMAP_MODE_LINEAR - VK_SAMPLER_MIPMAP_MODE_NEAREST + 1)
	VK_SAMPLER_MIPMAP_MODE_MAX_ENUM    VkSamplerMipmapMode = 2147483647
)

type VkSamplerAddressMode int

const (
	VK_SAMPLER_ADDRESS_MODE_REPEAT               VkSamplerAddressMode = 0
	VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT      VkSamplerAddressMode = 1
	VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE        VkSamplerAddressMode = 2
	VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER      VkSamplerAddressMode = 3
	VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE VkSamplerAddressMode = 4
	VK_SAMPLER_ADDRESS_MODE_BEGIN_RANGE          VkSamplerAddressMode = VK_SAMPLER_ADDRESS_MODE_REPEAT
	VK_SAMPLER_ADDRESS_MODE_END_RANGE            VkSamplerAddressMode = VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
	VK_SAMPLER_ADDRESS_MODE_RANGE_SIZE           VkSamplerAddressMode = (VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER - VK_SAMPLER_ADDRESS_MODE_REPEAT + 1)
	VK_SAMPLER_ADDRESS_MODE_MAX_ENUM             VkSamplerAddressMode = 2147483647
)

type VkBorderColor int

const (
	VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK VkBorderColor = 0
	VK_BORDER_COLOR_INT_TRANSPARENT_BLACK   VkBorderColor = 1
	VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK      VkBorderColor = 2
	VK_BORDER_COLOR_INT_OPAQUE_BLACK        VkBorderColor = 3
	VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE      VkBorderColor = 4
	VK_BORDER_COLOR_INT_OPAQUE_WHITE        VkBorderColor = 5
	VK_BORDER_COLOR_BEGIN_RANGE             VkBorderColor = VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	VK_BORDER_COLOR_END_RANGE               VkBorderColor = VK_BORDER_COLOR_INT_OPAQUE_WHITE
	VK_BORDER_COLOR_RANGE_SIZE              VkBorderColor = (VK_BORDER_COLOR_INT_OPAQUE_WHITE - VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK + 1)
	VK_BORDER_COLOR_MAX_ENUM                VkBorderColor = 2147483647
)

type VkSamplerCreateInfo struct {
	sType                   VkStructureType
	pNext                   *[0]byte
	flags                   VkSamplerCreateFlags
	magFilter               VkFilter
	minFilter               VkFilter
	mipmapMode              VkSamplerMipmapMode
	addressModeU            VkSamplerAddressMode
	addressModeV            VkSamplerAddressMode
	addressModeW            VkSamplerAddressMode
	mipLodBias              float32
	anisotropyEnable        VkBool32
	maxAnisotropy           float32
	compareEnable           VkBool32
	compareOp               VkCompareOp
	minLod                  float32
	maxLod                  float32
	borderColor             VkBorderColor
	unnormalizedCoordinates VkBool32
}

func (g *VkSamplerCreateInfo) toC(c *C.VkSamplerCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSamplerCreateFlags(_temp)
	}
	c.magFilter = C.VkFilter(g.magFilter)
	c.minFilter = C.VkFilter(g.minFilter)
	c.mipmapMode = C.VkSamplerMipmapMode(g.mipmapMode)
	c.addressModeU = C.VkSamplerAddressMode(g.addressModeU)
	c.addressModeV = C.VkSamplerAddressMode(g.addressModeV)
	c.addressModeW = C.VkSamplerAddressMode(g.addressModeW)
	c.mipLodBias = C.float(g.mipLodBias)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.anisotropyEnable))
		c.anisotropyEnable = C.VkBool32(_temp)
	}
	c.maxAnisotropy = C.float(g.maxAnisotropy)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.compareEnable))
		c.compareEnable = C.VkBool32(_temp)
	}
	c.compareOp = C.VkCompareOp(g.compareOp)
	c.minLod = C.float(g.minLod)
	c.maxLod = C.float(g.maxLod)
	c.borderColor = C.VkBorderColor(g.borderColor)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.unnormalizedCoordinates))
		c.unnormalizedCoordinates = C.VkBool32(_temp)
	}
}
func (g *VkSamplerCreateInfo) fromC(c *C.VkSamplerCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSamplerCreateFlags(_temp)
	}
	g.magFilter = VkFilter(c.magFilter)
	g.minFilter = VkFilter(c.minFilter)
	g.mipmapMode = VkSamplerMipmapMode(c.mipmapMode)
	g.addressModeU = VkSamplerAddressMode(c.addressModeU)
	g.addressModeV = VkSamplerAddressMode(c.addressModeV)
	g.addressModeW = VkSamplerAddressMode(c.addressModeW)
	g.mipLodBias = float32(c.mipLodBias)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.anisotropyEnable))
		g.anisotropyEnable = VkBool32(_temp)
	}
	g.maxAnisotropy = float32(c.maxAnisotropy)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.compareEnable))
		g.compareEnable = VkBool32(_temp)
	}
	g.compareOp = VkCompareOp(c.compareOp)
	g.minLod = float32(c.minLod)
	g.maxLod = float32(c.maxLod)
	g.borderColor = VkBorderColor(c.borderColor)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.unnormalizedCoordinates))
		g.unnormalizedCoordinates = VkBool32(_temp)
	}
}

type VkSampler C.VkSampler

func vkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSampler = (*C.VkSampler)(_sa.alloc(C.sizeof_VkSampler))
		*c.pSampler = C.VkSampler(*pSampler)
	}
	c._ret = C.vkCreateSampler(c.device, c.pCreateInfo, c.pAllocator, c.pSampler)
	_ret = VkResult(c._ret)
	*pSampler = VkSampler(*c.pSampler)
	return
}
func vkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroySampler(c.device, c.sampler, c.pAllocator)
}

type VkDescriptorSetLayoutCreateFlags VkFlags
type VkDescriptorType int

const (
	VK_DESCRIPTOR_TYPE_SAMPLER                VkDescriptorType = 0
	VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER VkDescriptorType = 1
	VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE          VkDescriptorType = 2
	VK_DESCRIPTOR_TYPE_STORAGE_IMAGE          VkDescriptorType = 3
	VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER   VkDescriptorType = 4
	VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER   VkDescriptorType = 5
	VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER         VkDescriptorType = 6
	VK_DESCRIPTOR_TYPE_STORAGE_BUFFER         VkDescriptorType = 7
	VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC VkDescriptorType = 8
	VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC VkDescriptorType = 9
	VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT       VkDescriptorType = 10
	VK_DESCRIPTOR_TYPE_BEGIN_RANGE            VkDescriptorType = VK_DESCRIPTOR_TYPE_SAMPLER
	VK_DESCRIPTOR_TYPE_END_RANGE              VkDescriptorType = VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
	VK_DESCRIPTOR_TYPE_RANGE_SIZE             VkDescriptorType = (VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT - VK_DESCRIPTOR_TYPE_SAMPLER + 1)
	VK_DESCRIPTOR_TYPE_MAX_ENUM               VkDescriptorType = 2147483647
)

type VkDescriptorSetLayoutBinding struct {
	binding            uint32
	descriptorType     VkDescriptorType
	descriptorCount    uint32
	stageFlags         VkShaderStageFlags
	pImmutableSamplers []VkSampler
}

func (g *VkDescriptorSetLayoutBinding) toC(c *C.VkDescriptorSetLayoutBinding, _sa *stackAllocator) {
	c.binding = C.uint32_t(g.binding)
	c.descriptorType = C.VkDescriptorType(g.descriptorType)
	c.descriptorCount = C.uint32_t(g.descriptorCount)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.stageFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.stageFlags = C.VkShaderStageFlags(_temp)
	}
	{
		c.pImmutableSamplers = (*C.VkSampler)(_sa.alloc(C.sizeof_VkSampler * uint(len(g.pImmutableSamplers))))
		slice2 := (*[1 << 31]C.VkSampler)(unsafe.Pointer(c.pImmutableSamplers))[:len(g.pImmutableSamplers):len(g.pImmutableSamplers)]
		for i2, _ := range g.pImmutableSamplers {
			slice2[i2] = C.VkSampler(g.pImmutableSamplers[i2])
		}
	}
}
func (g *VkDescriptorSetLayoutBinding) fromC(c *C.VkDescriptorSetLayoutBinding) {
	g.binding = uint32(c.binding)
	g.descriptorType = VkDescriptorType(c.descriptorType)
	g.descriptorCount = uint32(c.descriptorCount)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.stageFlags)))
			_temp = VkFlags(_temp)
		}
		g.stageFlags = VkShaderStageFlags(_temp)
	}
	{
		slice2 := (*[1 << 31]C.VkSampler)(unsafe.Pointer(c.pImmutableSamplers))[:len(g.pImmutableSamplers):len(g.pImmutableSamplers)]
		for i2, _ := range g.pImmutableSamplers {
			g.pImmutableSamplers[i2] = VkSampler(slice2[i2])
		}
	}
}

type VkDescriptorSetLayoutCreateInfo struct {
	sType     VkStructureType
	pNext     *[0]byte
	flags     VkDescriptorSetLayoutCreateFlags
	pBindings []VkDescriptorSetLayoutBinding
}

func (g *VkDescriptorSetLayoutCreateInfo) toC(c *C.VkDescriptorSetLayoutCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkDescriptorSetLayoutCreateFlags(_temp)
	}
	c.bindingCount = C.uint32_t(len(g.pBindings))
	{
		c.pBindings = (*C.VkDescriptorSetLayoutBinding)(_sa.alloc(C.sizeof_VkDescriptorSetLayoutBinding * uint(len(g.pBindings))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayoutBinding)(unsafe.Pointer(c.pBindings))[:len(g.pBindings):len(g.pBindings)]
		for i2, _ := range g.pBindings {
			g.pBindings[i2].toC(&slice2[i2], _sa)
		}
	}
}
func (g *VkDescriptorSetLayoutCreateInfo) fromC(c *C.VkDescriptorSetLayoutCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkDescriptorSetLayoutCreateFlags(_temp)
	}
	g.pBindings = make([]VkDescriptorSetLayoutBinding, int(c.bindingCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayoutBinding)(unsafe.Pointer(c.pBindings))[:len(g.pBindings):len(g.pBindings)]
		for i2, _ := range g.pBindings {
			g.pBindings[i2].fromC(&slice2[i2])
		}
	}
}
func vkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pSetLayout = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout))
		*c.pSetLayout = C.VkDescriptorSetLayout(*pSetLayout)
	}
	c._ret = C.vkCreateDescriptorSetLayout(c.device, c.pCreateInfo, c.pAllocator, c.pSetLayout)
	_ret = VkResult(c._ret)
	*pSetLayout = VkDescriptorSetLayout(*c.pSetLayout)
	return
}
func vkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorSetLayout(c.device, c.descriptorSetLayout, c.pAllocator)
}

type VkDescriptorPoolCreateFlags VkFlags
type VkDescriptorPoolSize struct {
	type            VkDescriptorType
	descriptorCount uint32
}

func (g *VkDescriptorPoolSize) toC(c *C.VkDescriptorPoolSize) {
	c.type = C.VkDescriptorType(g.type)
	c.descriptorCount = C.uint32_t(g.descriptorCount)
}
func (g *VkDescriptorPoolSize) fromC(c *C.VkDescriptorPoolSize) {
	g.type = VkDescriptorType(c.type)
	g.descriptorCount = uint32(c.descriptorCount)
}

type VkDescriptorPoolCreateInfo struct {
	sType      VkStructureType
	pNext      *[0]byte
	flags      VkDescriptorPoolCreateFlags
	maxSets    uint32
	pPoolSizes []VkDescriptorPoolSize
}

func (g *VkDescriptorPoolCreateInfo) toC(c *C.VkDescriptorPoolCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkDescriptorPoolCreateFlags(_temp)
	}
	c.maxSets = C.uint32_t(g.maxSets)
	c.poolSizeCount = C.uint32_t(len(g.pPoolSizes))
	{
		c.pPoolSizes = (*C.VkDescriptorPoolSize)(_sa.alloc(C.sizeof_VkDescriptorPoolSize * uint(len(g.pPoolSizes))))
		slice2 := (*[1 << 31]C.VkDescriptorPoolSize)(unsafe.Pointer(c.pPoolSizes))[:len(g.pPoolSizes):len(g.pPoolSizes)]
		for i2, _ := range g.pPoolSizes {
			g.pPoolSizes[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkDescriptorPoolCreateInfo) fromC(c *C.VkDescriptorPoolCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkDescriptorPoolCreateFlags(_temp)
	}
	g.maxSets = uint32(c.maxSets)
	g.pPoolSizes = make([]VkDescriptorPoolSize, int(c.poolSizeCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorPoolSize)(unsafe.Pointer(c.pPoolSizes))[:len(g.pPoolSizes):len(g.pPoolSizes)]
		for i2, _ := range g.pPoolSizes {
			g.pPoolSizes[i2].fromC(&slice2[i2])
		}
	}
}

type VkDescriptorPool C.VkDescriptorPool

func vkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pDescriptorPool = (*C.VkDescriptorPool)(_sa.alloc(C.sizeof_VkDescriptorPool))
		*c.pDescriptorPool = C.VkDescriptorPool(*pDescriptorPool)
	}
	c._ret = C.vkCreateDescriptorPool(c.device, c.pCreateInfo, c.pAllocator, c.pDescriptorPool)
	_ret = VkResult(c._ret)
	*pDescriptorPool = VkDescriptorPool(*c.pDescriptorPool)
	return
}
func vkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyDescriptorPool(c.device, c.descriptorPool, c.pAllocator)
}

type VkDescriptorPoolResetFlags VkFlags

func vkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (_ret VkResult) {
	var c struct {
		device         C.VkDevice
		descriptorPool C.VkDescriptorPool
		flags          C.VkDescriptorPoolResetFlags
		_ret           C.VkResult
	}
	c.device = C.VkDevice(device)
	c.descriptorPool = C.VkDescriptorPool(descriptorPool)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkDescriptorPoolResetFlags(_temp)
	}
	c._ret = C.vkResetDescriptorPool(c.device, c.descriptorPool, c.flags)
	_ret = VkResult(c._ret)
	return
}

type VkDescriptorSetAllocateInfo struct {
	sType          VkStructureType
	pNext          *[0]byte
	descriptorPool VkDescriptorPool
	pSetLayouts    []VkDescriptorSetLayout
}

func (g *VkDescriptorSetAllocateInfo) toC(c *C.VkDescriptorSetAllocateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.descriptorPool = C.VkDescriptorPool(g.descriptorPool)
	c.descriptorSetCount = C.uint32_t(len(g.pSetLayouts))
	{
		c.pSetLayouts = (*C.VkDescriptorSetLayout)(_sa.alloc(C.sizeof_VkDescriptorSetLayout * uint(len(g.pSetLayouts))))
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.pSetLayouts):len(g.pSetLayouts)]
		for i2, _ := range g.pSetLayouts {
			slice2[i2] = C.VkDescriptorSetLayout(g.pSetLayouts[i2])
		}
	}
}
func (g *VkDescriptorSetAllocateInfo) fromC(c *C.VkDescriptorSetAllocateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.descriptorPool = VkDescriptorPool(c.descriptorPool)
	g.pSetLayouts = make([]VkDescriptorSetLayout, int(c.descriptorSetCount))
	{
		slice2 := (*[1 << 31]C.VkDescriptorSetLayout)(unsafe.Pointer(c.pSetLayouts))[:len(g.pSetLayouts):len(g.pSetLayouts)]
		for i2, _ := range g.pSetLayouts {
			g.pSetLayouts[i2] = VkDescriptorSetLayout(slice2[i2])
		}
	}
}

type VkDescriptorSet C.VkDescriptorSet

func vkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets []VkDescriptorSet) (_ret VkResult) {
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
		pAllocateInfo.toC(c.pAllocateInfo, _sa)
	}
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(pDescriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(pDescriptorSets):len(pDescriptorSets)]
		for i3, _ := range pDescriptorSets {
			slice3[i3] = C.VkDescriptorSet(pDescriptorSets[i3])
		}
	}
	c._ret = C.vkAllocateDescriptorSets(c.device, c.pAllocateInfo, c.pDescriptorSets)
	_ret = VkResult(c._ret)
	return
}
func vkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, pDescriptorSets []VkDescriptorSet) (_ret VkResult) {
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
	c.descriptorSetCount = C.uint32_t(len(pDescriptorSets))
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(pDescriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(pDescriptorSets):len(pDescriptorSets)]
		for i3, _ := range pDescriptorSets {
			slice3[i3] = C.VkDescriptorSet(pDescriptorSets[i3])
		}
	}
	c._ret = C.vkFreeDescriptorSets(c.device, c.descriptorPool, c.descriptorSetCount, c.pDescriptorSets)
	_ret = VkResult(c._ret)
	return
}

type VkDescriptorImageInfo struct {
	sampler     VkSampler
	imageView   VkImageView
	imageLayout VkImageLayout
}

func (g *VkDescriptorImageInfo) toC(c *C.VkDescriptorImageInfo) {
	c.sampler = C.VkSampler(g.sampler)
	c.imageView = C.VkImageView(g.imageView)
	c.imageLayout = C.VkImageLayout(g.imageLayout)
}
func (g *VkDescriptorImageInfo) fromC(c *C.VkDescriptorImageInfo) {
	g.sampler = VkSampler(c.sampler)
	g.imageView = VkImageView(c.imageView)
	g.imageLayout = VkImageLayout(c.imageLayout)
}

type VkDescriptorBufferInfo struct {
	buffer VkBuffer
	offset VkDeviceSize
	range  VkDeviceSize
}

func (g *VkDescriptorBufferInfo) toC(c *C.VkDescriptorBufferInfo) {
	c.buffer = C.VkBuffer(g.buffer)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.range))
		c.range = C.VkDeviceSize(_temp)
	}
}
func (g *VkDescriptorBufferInfo) fromC(c *C.VkDescriptorBufferInfo) {
	g.buffer = VkBuffer(c.buffer)
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.offset))
		g.offset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.range))
		g.range = VkDeviceSize(_temp)
	}
}

type VkWriteDescriptorSet struct {
	sType            VkStructureType
	pNext            *[0]byte
	dstSet           VkDescriptorSet
	dstBinding       uint32
	dstArrayElement  uint32
	descriptorCount  uint32
	descriptorType   VkDescriptorType
	pImageInfo       *VkDescriptorImageInfo
	pBufferInfo      *VkDescriptorBufferInfo
	pTexelBufferView *VkBufferView
}

func (g *VkWriteDescriptorSet) toC(c *C.VkWriteDescriptorSet, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.dstSet = C.VkDescriptorSet(g.dstSet)
	c.dstBinding = C.uint32_t(g.dstBinding)
	c.dstArrayElement = C.uint32_t(g.dstArrayElement)
	c.descriptorCount = C.uint32_t(g.descriptorCount)
	c.descriptorType = C.VkDescriptorType(g.descriptorType)
	{
		c.pImageInfo = (*C.VkDescriptorImageInfo)(_sa.alloc(C.sizeof_VkDescriptorImageInfo))
		g.pImageInfo.toC(c.pImageInfo)
	}
	{
		c.pBufferInfo = (*C.VkDescriptorBufferInfo)(_sa.alloc(C.sizeof_VkDescriptorBufferInfo))
		g.pBufferInfo.toC(c.pBufferInfo)
	}
	{
		c.pTexelBufferView = (*C.VkBufferView)(_sa.alloc(C.sizeof_VkBufferView))
		*c.pTexelBufferView = C.VkBufferView(*g.pTexelBufferView)
	}
}
func (g *VkWriteDescriptorSet) fromC(c *C.VkWriteDescriptorSet) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.dstSet = VkDescriptorSet(c.dstSet)
	g.dstBinding = uint32(c.dstBinding)
	g.dstArrayElement = uint32(c.dstArrayElement)
	g.descriptorCount = uint32(c.descriptorCount)
	g.descriptorType = VkDescriptorType(c.descriptorType)
	{
		if g.pImageInfo == nil {
			g.pImageInfo = new(VkDescriptorImageInfo)
		}
		g.pImageInfo.fromC(c.pImageInfo)
	}
	{
		if g.pBufferInfo == nil {
			g.pBufferInfo = new(VkDescriptorBufferInfo)
		}
		g.pBufferInfo.fromC(c.pBufferInfo)
	}
	{
		if g.pTexelBufferView == nil {
			g.pTexelBufferView = new(VkBufferView)
		}
		*g.pTexelBufferView = VkBufferView(*c.pTexelBufferView)
	}
}

type VkCopyDescriptorSet struct {
	sType           VkStructureType
	pNext           *[0]byte
	srcSet          VkDescriptorSet
	srcBinding      uint32
	srcArrayElement uint32
	dstSet          VkDescriptorSet
	dstBinding      uint32
	dstArrayElement uint32
	descriptorCount uint32
}

func (g *VkCopyDescriptorSet) toC(c *C.VkCopyDescriptorSet) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.srcSet = C.VkDescriptorSet(g.srcSet)
	c.srcBinding = C.uint32_t(g.srcBinding)
	c.srcArrayElement = C.uint32_t(g.srcArrayElement)
	c.dstSet = C.VkDescriptorSet(g.dstSet)
	c.dstBinding = C.uint32_t(g.dstBinding)
	c.dstArrayElement = C.uint32_t(g.dstArrayElement)
	c.descriptorCount = C.uint32_t(g.descriptorCount)
}
func (g *VkCopyDescriptorSet) fromC(c *C.VkCopyDescriptorSet) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.srcSet = VkDescriptorSet(c.srcSet)
	g.srcBinding = uint32(c.srcBinding)
	g.srcArrayElement = uint32(c.srcArrayElement)
	g.dstSet = VkDescriptorSet(c.dstSet)
	g.dstBinding = uint32(c.dstBinding)
	g.dstArrayElement = uint32(c.dstArrayElement)
	g.descriptorCount = uint32(c.descriptorCount)
}
func vkUpdateDescriptorSets(device VkDevice, pDescriptorWrites []VkWriteDescriptorSet, pDescriptorCopies []VkCopyDescriptorSet) {
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
	c.descriptorWriteCount = C.uint32_t(len(pDescriptorWrites))
	{
		c.pDescriptorWrites = (*C.VkWriteDescriptorSet)(_sa.alloc(C.sizeof_VkWriteDescriptorSet * uint(len(pDescriptorWrites))))
		slice3 := (*[1 << 31]C.VkWriteDescriptorSet)(unsafe.Pointer(c.pDescriptorWrites))[:len(pDescriptorWrites):len(pDescriptorWrites)]
		for i3, _ := range pDescriptorWrites {
			pDescriptorWrites[i3].toC(&slice3[i3], _sa)
		}
	}
	c.descriptorCopyCount = C.uint32_t(len(pDescriptorCopies))
	{
		c.pDescriptorCopies = (*C.VkCopyDescriptorSet)(_sa.alloc(C.sizeof_VkCopyDescriptorSet * uint(len(pDescriptorCopies))))
		slice3 := (*[1 << 31]C.VkCopyDescriptorSet)(unsafe.Pointer(c.pDescriptorCopies))[:len(pDescriptorCopies):len(pDescriptorCopies)]
		for i3, _ := range pDescriptorCopies {
			pDescriptorCopies[i3].toC(&slice3[i3])
		}
	}
	C.vkUpdateDescriptorSets(c.device, c.descriptorWriteCount, c.pDescriptorWrites, c.descriptorCopyCount, c.pDescriptorCopies)
}

type VkFramebufferCreateFlags VkFlags
type VkFramebufferCreateInfo struct {
	sType        VkStructureType
	pNext        *[0]byte
	flags        VkFramebufferCreateFlags
	renderPass   VkRenderPass
	pAttachments []VkImageView
	width        uint32
	height       uint32
	layers       uint32
}

func (g *VkFramebufferCreateInfo) toC(c *C.VkFramebufferCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkFramebufferCreateFlags(_temp)
	}
	c.renderPass = C.VkRenderPass(g.renderPass)
	c.attachmentCount = C.uint32_t(len(g.pAttachments))
	{
		c.pAttachments = (*C.VkImageView)(_sa.alloc(C.sizeof_VkImageView * uint(len(g.pAttachments))))
		slice2 := (*[1 << 31]C.VkImageView)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			slice2[i2] = C.VkImageView(g.pAttachments[i2])
		}
	}
	c.width = C.uint32_t(g.width)
	c.height = C.uint32_t(g.height)
	c.layers = C.uint32_t(g.layers)
}
func (g *VkFramebufferCreateInfo) fromC(c *C.VkFramebufferCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkFramebufferCreateFlags(_temp)
	}
	g.renderPass = VkRenderPass(c.renderPass)
	g.pAttachments = make([]VkImageView, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkImageView)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			g.pAttachments[i2] = VkImageView(slice2[i2])
		}
	}
	g.width = uint32(c.width)
	g.height = uint32(c.height)
	g.layers = uint32(c.layers)
}

type VkFramebuffer C.VkFramebuffer

func vkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pFramebuffer = (*C.VkFramebuffer)(_sa.alloc(C.sizeof_VkFramebuffer))
		*c.pFramebuffer = C.VkFramebuffer(*pFramebuffer)
	}
	c._ret = C.vkCreateFramebuffer(c.device, c.pCreateInfo, c.pAllocator, c.pFramebuffer)
	_ret = VkResult(c._ret)
	*pFramebuffer = VkFramebuffer(*c.pFramebuffer)
	return
}
func vkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyFramebuffer(c.device, c.framebuffer, c.pAllocator)
}

type VkRenderPassCreateFlags VkFlags
type VkAttachmentDescriptionFlags VkFlags
type VkAttachmentLoadOp int

const (
	VK_ATTACHMENT_LOAD_OP_LOAD        VkAttachmentLoadOp = 0
	VK_ATTACHMENT_LOAD_OP_CLEAR       VkAttachmentLoadOp = 1
	VK_ATTACHMENT_LOAD_OP_DONT_CARE   VkAttachmentLoadOp = 2
	VK_ATTACHMENT_LOAD_OP_BEGIN_RANGE VkAttachmentLoadOp = VK_ATTACHMENT_LOAD_OP_LOAD
	VK_ATTACHMENT_LOAD_OP_END_RANGE   VkAttachmentLoadOp = VK_ATTACHMENT_LOAD_OP_DONT_CARE
	VK_ATTACHMENT_LOAD_OP_RANGE_SIZE  VkAttachmentLoadOp = (VK_ATTACHMENT_LOAD_OP_DONT_CARE - VK_ATTACHMENT_LOAD_OP_LOAD + 1)
	VK_ATTACHMENT_LOAD_OP_MAX_ENUM    VkAttachmentLoadOp = 2147483647
)

type VkAttachmentStoreOp int

const (
	VK_ATTACHMENT_STORE_OP_STORE       VkAttachmentStoreOp = 0
	VK_ATTACHMENT_STORE_OP_DONT_CARE   VkAttachmentStoreOp = 1
	VK_ATTACHMENT_STORE_OP_BEGIN_RANGE VkAttachmentStoreOp = VK_ATTACHMENT_STORE_OP_STORE
	VK_ATTACHMENT_STORE_OP_END_RANGE   VkAttachmentStoreOp = VK_ATTACHMENT_STORE_OP_DONT_CARE
	VK_ATTACHMENT_STORE_OP_RANGE_SIZE  VkAttachmentStoreOp = (VK_ATTACHMENT_STORE_OP_DONT_CARE - VK_ATTACHMENT_STORE_OP_STORE + 1)
	VK_ATTACHMENT_STORE_OP_MAX_ENUM    VkAttachmentStoreOp = 2147483647
)

type VkAttachmentDescription struct {
	flags          VkAttachmentDescriptionFlags
	format         VkFormat
	samples        VkSampleCountFlagBits
	loadOp         VkAttachmentLoadOp
	storeOp        VkAttachmentStoreOp
	stencilLoadOp  VkAttachmentLoadOp
	stencilStoreOp VkAttachmentStoreOp
	initialLayout  VkImageLayout
	finalLayout    VkImageLayout
}

func (g *VkAttachmentDescription) toC(c *C.VkAttachmentDescription) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkAttachmentDescriptionFlags(_temp)
	}
	c.format = C.VkFormat(g.format)
	c.samples = C.VkSampleCountFlagBits(g.samples)
	c.loadOp = C.VkAttachmentLoadOp(g.loadOp)
	c.storeOp = C.VkAttachmentStoreOp(g.storeOp)
	c.stencilLoadOp = C.VkAttachmentLoadOp(g.stencilLoadOp)
	c.stencilStoreOp = C.VkAttachmentStoreOp(g.stencilStoreOp)
	c.initialLayout = C.VkImageLayout(g.initialLayout)
	c.finalLayout = C.VkImageLayout(g.finalLayout)
}
func (g *VkAttachmentDescription) fromC(c *C.VkAttachmentDescription) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkAttachmentDescriptionFlags(_temp)
	}
	g.format = VkFormat(c.format)
	g.samples = VkSampleCountFlagBits(c.samples)
	g.loadOp = VkAttachmentLoadOp(c.loadOp)
	g.storeOp = VkAttachmentStoreOp(c.storeOp)
	g.stencilLoadOp = VkAttachmentLoadOp(c.stencilLoadOp)
	g.stencilStoreOp = VkAttachmentStoreOp(c.stencilStoreOp)
	g.initialLayout = VkImageLayout(c.initialLayout)
	g.finalLayout = VkImageLayout(c.finalLayout)
}

type VkSubpassDescriptionFlags VkFlags
type VkPipelineBindPoint int

const (
	VK_PIPELINE_BIND_POINT_GRAPHICS    VkPipelineBindPoint = 0
	VK_PIPELINE_BIND_POINT_COMPUTE     VkPipelineBindPoint = 1
	VK_PIPELINE_BIND_POINT_BEGIN_RANGE VkPipelineBindPoint = VK_PIPELINE_BIND_POINT_GRAPHICS
	VK_PIPELINE_BIND_POINT_END_RANGE   VkPipelineBindPoint = VK_PIPELINE_BIND_POINT_COMPUTE
	VK_PIPELINE_BIND_POINT_RANGE_SIZE  VkPipelineBindPoint = (VK_PIPELINE_BIND_POINT_COMPUTE - VK_PIPELINE_BIND_POINT_GRAPHICS + 1)
	VK_PIPELINE_BIND_POINT_MAX_ENUM    VkPipelineBindPoint = 2147483647
)

type VkAttachmentReference struct {
	attachment uint32
	layout     VkImageLayout
}

func (g *VkAttachmentReference) toC(c *C.VkAttachmentReference) {
	c.attachment = C.uint32_t(g.attachment)
	c.layout = C.VkImageLayout(g.layout)
}
func (g *VkAttachmentReference) fromC(c *C.VkAttachmentReference) {
	g.attachment = uint32(c.attachment)
	g.layout = VkImageLayout(c.layout)
}

type VkSubpassDescription struct {
	flags                   VkSubpassDescriptionFlags
	pipelineBindPoint       VkPipelineBindPoint
	pInputAttachments       []VkAttachmentReference
	pColorAttachments       []VkAttachmentReference
	pResolveAttachments     []VkAttachmentReference
	pDepthStencilAttachment *VkAttachmentReference
	pPreserveAttachments    []uint32
}

func (g *VkSubpassDescription) toC(c *C.VkSubpassDescription, _sa *stackAllocator) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkSubpassDescriptionFlags(_temp)
	}
	c.pipelineBindPoint = C.VkPipelineBindPoint(g.pipelineBindPoint)
	c.inputAttachmentCount = C.uint32_t(len(g.pInputAttachments))
	{
		c.pInputAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.pInputAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pInputAttachments))[:len(g.pInputAttachments):len(g.pInputAttachments)]
		for i2, _ := range g.pInputAttachments {
			g.pInputAttachments[i2].toC(&slice2[i2])
		}
	}
	c.colorAttachmentCount = C.uint32_t(len(g.pColorAttachments))
	{
		c.pColorAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.pColorAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pColorAttachments))[:len(g.pColorAttachments):len(g.pColorAttachments)]
		for i2, _ := range g.pColorAttachments {
			g.pColorAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pResolveAttachments = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference * uint(len(g.pResolveAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pResolveAttachments))[:len(g.pResolveAttachments):len(g.pResolveAttachments)]
		for i2, _ := range g.pResolveAttachments {
			g.pResolveAttachments[i2].toC(&slice2[i2])
		}
	}
	{
		c.pDepthStencilAttachment = (*C.VkAttachmentReference)(_sa.alloc(C.sizeof_VkAttachmentReference))
		g.pDepthStencilAttachment.toC(c.pDepthStencilAttachment)
	}
	c.preserveAttachmentCount = C.uint32_t(len(g.pPreserveAttachments))
	{
		c.pPreserveAttachments = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(g.pPreserveAttachments))))
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.pPreserveAttachments):len(g.pPreserveAttachments)]
		for i2, _ := range g.pPreserveAttachments {
			slice2[i2] = C.uint32_t(g.pPreserveAttachments[i2])
		}
	}
}
func (g *VkSubpassDescription) fromC(c *C.VkSubpassDescription) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkSubpassDescriptionFlags(_temp)
	}
	g.pipelineBindPoint = VkPipelineBindPoint(c.pipelineBindPoint)
	g.pInputAttachments = make([]VkAttachmentReference, int(c.inputAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pInputAttachments))[:len(g.pInputAttachments):len(g.pInputAttachments)]
		for i2, _ := range g.pInputAttachments {
			g.pInputAttachments[i2].fromC(&slice2[i2])
		}
	}
	g.pColorAttachments = make([]VkAttachmentReference, int(c.colorAttachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pColorAttachments))[:len(g.pColorAttachments):len(g.pColorAttachments)]
		for i2, _ := range g.pColorAttachments {
			g.pColorAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		slice2 := (*[1 << 31]C.VkAttachmentReference)(unsafe.Pointer(c.pResolveAttachments))[:len(g.pResolveAttachments):len(g.pResolveAttachments)]
		for i2, _ := range g.pResolveAttachments {
			g.pResolveAttachments[i2].fromC(&slice2[i2])
		}
	}
	{
		if g.pDepthStencilAttachment == nil {
			g.pDepthStencilAttachment = new(VkAttachmentReference)
		}
		g.pDepthStencilAttachment.fromC(c.pDepthStencilAttachment)
	}
	g.pPreserveAttachments = make([]uint32, int(c.preserveAttachmentCount))
	{
		slice2 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pPreserveAttachments))[:len(g.pPreserveAttachments):len(g.pPreserveAttachments)]
		for i2, _ := range g.pPreserveAttachments {
			g.pPreserveAttachments[i2] = uint32(slice2[i2])
		}
	}
}

type VkAccessFlags VkFlags
type VkDependencyFlags VkFlags
type VkSubpassDependency struct {
	srcSubpass      uint32
	dstSubpass      uint32
	srcStageMask    VkPipelineStageFlags
	dstStageMask    VkPipelineStageFlags
	srcAccessMask   VkAccessFlags
	dstAccessMask   VkAccessFlags
	dependencyFlags VkDependencyFlags
}

func (g *VkSubpassDependency) toC(c *C.VkSubpassDependency) {
	c.srcSubpass = C.uint32_t(g.srcSubpass)
	c.dstSubpass = C.uint32_t(g.dstSubpass)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.srcStageMask)))
			_temp = C.VkFlags(_temp)
		}
		c.srcStageMask = C.VkPipelineStageFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.dstStageMask)))
			_temp = C.VkFlags(_temp)
		}
		c.dstStageMask = C.VkPipelineStageFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.srcAccessMask)))
			_temp = C.VkFlags(_temp)
		}
		c.srcAccessMask = C.VkAccessFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.dstAccessMask)))
			_temp = C.VkFlags(_temp)
		}
		c.dstAccessMask = C.VkAccessFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.dependencyFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.dependencyFlags = C.VkDependencyFlags(_temp)
	}
}
func (g *VkSubpassDependency) fromC(c *C.VkSubpassDependency) {
	g.srcSubpass = uint32(c.srcSubpass)
	g.dstSubpass = uint32(c.dstSubpass)
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.srcStageMask)))
			_temp = VkFlags(_temp)
		}
		g.srcStageMask = VkPipelineStageFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.dstStageMask)))
			_temp = VkFlags(_temp)
		}
		g.dstStageMask = VkPipelineStageFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.srcAccessMask)))
			_temp = VkFlags(_temp)
		}
		g.srcAccessMask = VkAccessFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.dstAccessMask)))
			_temp = VkFlags(_temp)
		}
		g.dstAccessMask = VkAccessFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.dependencyFlags)))
			_temp = VkFlags(_temp)
		}
		g.dependencyFlags = VkDependencyFlags(_temp)
	}
}

type VkRenderPassCreateInfo struct {
	sType         VkStructureType
	pNext         *[0]byte
	flags         VkRenderPassCreateFlags
	pAttachments  []VkAttachmentDescription
	pSubpasses    []VkSubpassDescription
	pDependencies []VkSubpassDependency
}

func (g *VkRenderPassCreateInfo) toC(c *C.VkRenderPassCreateInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkRenderPassCreateFlags(_temp)
	}
	c.attachmentCount = C.uint32_t(len(g.pAttachments))
	{
		c.pAttachments = (*C.VkAttachmentDescription)(_sa.alloc(C.sizeof_VkAttachmentDescription * uint(len(g.pAttachments))))
		slice2 := (*[1 << 31]C.VkAttachmentDescription)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			g.pAttachments[i2].toC(&slice2[i2])
		}
	}
	c.subpassCount = C.uint32_t(len(g.pSubpasses))
	{
		c.pSubpasses = (*C.VkSubpassDescription)(_sa.alloc(C.sizeof_VkSubpassDescription * uint(len(g.pSubpasses))))
		slice2 := (*[1 << 31]C.VkSubpassDescription)(unsafe.Pointer(c.pSubpasses))[:len(g.pSubpasses):len(g.pSubpasses)]
		for i2, _ := range g.pSubpasses {
			g.pSubpasses[i2].toC(&slice2[i2], _sa)
		}
	}
	c.dependencyCount = C.uint32_t(len(g.pDependencies))
	{
		c.pDependencies = (*C.VkSubpassDependency)(_sa.alloc(C.sizeof_VkSubpassDependency * uint(len(g.pDependencies))))
		slice2 := (*[1 << 31]C.VkSubpassDependency)(unsafe.Pointer(c.pDependencies))[:len(g.pDependencies):len(g.pDependencies)]
		for i2, _ := range g.pDependencies {
			g.pDependencies[i2].toC(&slice2[i2])
		}
	}
}
func (g *VkRenderPassCreateInfo) fromC(c *C.VkRenderPassCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkRenderPassCreateFlags(_temp)
	}
	g.pAttachments = make([]VkAttachmentDescription, int(c.attachmentCount))
	{
		slice2 := (*[1 << 31]C.VkAttachmentDescription)(unsafe.Pointer(c.pAttachments))[:len(g.pAttachments):len(g.pAttachments)]
		for i2, _ := range g.pAttachments {
			g.pAttachments[i2].fromC(&slice2[i2])
		}
	}
	g.pSubpasses = make([]VkSubpassDescription, int(c.subpassCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDescription)(unsafe.Pointer(c.pSubpasses))[:len(g.pSubpasses):len(g.pSubpasses)]
		for i2, _ := range g.pSubpasses {
			g.pSubpasses[i2].fromC(&slice2[i2])
		}
	}
	g.pDependencies = make([]VkSubpassDependency, int(c.dependencyCount))
	{
		slice2 := (*[1 << 31]C.VkSubpassDependency)(unsafe.Pointer(c.pDependencies))[:len(g.pDependencies):len(g.pDependencies)]
		for i2, _ := range g.pDependencies {
			g.pDependencies[i2].fromC(&slice2[i2])
		}
	}
}
func vkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo, _sa)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pRenderPass = (*C.VkRenderPass)(_sa.alloc(C.sizeof_VkRenderPass))
		*c.pRenderPass = C.VkRenderPass(*pRenderPass)
	}
	c._ret = C.vkCreateRenderPass(c.device, c.pCreateInfo, c.pAllocator, c.pRenderPass)
	_ret = VkResult(c._ret)
	*pRenderPass = VkRenderPass(*c.pRenderPass)
	return
}
func vkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyRenderPass(c.device, c.renderPass, c.pAllocator)
}
func vkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) {
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
		pGranularity.toC(c.pGranularity)
	}
	C.vkGetRenderAreaGranularity(c.device, c.renderPass, c.pGranularity)
	pGranularity.fromC(c.pGranularity)
}

type VkCommandPoolCreateFlags VkFlags
type VkCommandPoolCreateInfo struct {
	sType            VkStructureType
	pNext            *[0]byte
	flags            VkCommandPoolCreateFlags
	queueFamilyIndex uint32
}

func (g *VkCommandPoolCreateInfo) toC(c *C.VkCommandPoolCreateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkCommandPoolCreateFlags(_temp)
	}
	c.queueFamilyIndex = C.uint32_t(g.queueFamilyIndex)
}
func (g *VkCommandPoolCreateInfo) fromC(c *C.VkCommandPoolCreateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkCommandPoolCreateFlags(_temp)
	}
	g.queueFamilyIndex = uint32(c.queueFamilyIndex)
}

type VkCommandPool C.VkCommandPool

func vkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (_ret VkResult) {
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
		pCreateInfo.toC(c.pCreateInfo)
	}
	{
		c.pAllocator = (*C.VkAllocationCallbacks)(_sa.alloc(C.sizeof_VkAllocationCallbacks))
		pAllocator.toC(c.pAllocator, _sa)
	}
	{
		c.pCommandPool = (*C.VkCommandPool)(_sa.alloc(C.sizeof_VkCommandPool))
		*c.pCommandPool = C.VkCommandPool(*pCommandPool)
	}
	c._ret = C.vkCreateCommandPool(c.device, c.pCreateInfo, c.pAllocator, c.pCommandPool)
	_ret = VkResult(c._ret)
	*pCommandPool = VkCommandPool(*c.pCommandPool)
	return
}
func vkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) {
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
		pAllocator.toC(c.pAllocator, _sa)
	}
	C.vkDestroyCommandPool(c.device, c.commandPool, c.pAllocator)
}

type VkCommandPoolResetFlags VkFlags

func vkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (_ret VkResult) {
	var c struct {
		device      C.VkDevice
		commandPool C.VkCommandPool
		flags       C.VkCommandPoolResetFlags
		_ret        C.VkResult
	}
	c.device = C.VkDevice(device)
	c.commandPool = C.VkCommandPool(commandPool)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkCommandPoolResetFlags(_temp)
	}
	c._ret = C.vkResetCommandPool(c.device, c.commandPool, c.flags)
	_ret = VkResult(c._ret)
	return
}

type VkCommandBufferLevel int

const (
	VK_COMMAND_BUFFER_LEVEL_PRIMARY     VkCommandBufferLevel = 0
	VK_COMMAND_BUFFER_LEVEL_SECONDARY   VkCommandBufferLevel = 1
	VK_COMMAND_BUFFER_LEVEL_BEGIN_RANGE VkCommandBufferLevel = VK_COMMAND_BUFFER_LEVEL_PRIMARY
	VK_COMMAND_BUFFER_LEVEL_END_RANGE   VkCommandBufferLevel = VK_COMMAND_BUFFER_LEVEL_SECONDARY
	VK_COMMAND_BUFFER_LEVEL_RANGE_SIZE  VkCommandBufferLevel = (VK_COMMAND_BUFFER_LEVEL_SECONDARY - VK_COMMAND_BUFFER_LEVEL_PRIMARY + 1)
	VK_COMMAND_BUFFER_LEVEL_MAX_ENUM    VkCommandBufferLevel = 2147483647
)

type VkCommandBufferAllocateInfo struct {
	sType              VkStructureType
	pNext              *[0]byte
	commandPool        VkCommandPool
	level              VkCommandBufferLevel
	commandBufferCount uint32
}

func (g *VkCommandBufferAllocateInfo) toC(c *C.VkCommandBufferAllocateInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.commandPool = C.VkCommandPool(g.commandPool)
	c.level = C.VkCommandBufferLevel(g.level)
	c.commandBufferCount = C.uint32_t(g.commandBufferCount)
}
func (g *VkCommandBufferAllocateInfo) fromC(c *C.VkCommandBufferAllocateInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.commandPool = VkCommandPool(c.commandPool)
	g.level = VkCommandBufferLevel(c.level)
	g.commandBufferCount = uint32(c.commandBufferCount)
}
func vkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers []VkCommandBuffer) (_ret VkResult) {
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
		pAllocateInfo.toC(c.pAllocateInfo)
	}
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(pCommandBuffers))))
		slice3 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(pCommandBuffers):len(pCommandBuffers)]
		for i3, _ := range pCommandBuffers {
			slice3[i3] = C.VkCommandBuffer(pCommandBuffers[i3])
		}
	}
	c._ret = C.vkAllocateCommandBuffers(c.device, c.pAllocateInfo, c.pCommandBuffers)
	_ret = VkResult(c._ret)
	return
}
func vkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, pCommandBuffers []VkCommandBuffer) {
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
	c.commandBufferCount = C.uint32_t(len(pCommandBuffers))
	{
		c.pCommandBuffers = (*C.VkCommandBuffer)(_sa.alloc(C.sizeof_VkCommandBuffer * uint(len(pCommandBuffers))))
		slice3 := (*[1 << 31]C.VkCommandBuffer)(unsafe.Pointer(c.pCommandBuffers))[:len(pCommandBuffers):len(pCommandBuffers)]
		for i3, _ := range pCommandBuffers {
			slice3[i3] = C.VkCommandBuffer(pCommandBuffers[i3])
		}
	}
	C.vkFreeCommandBuffers(c.device, c.commandPool, c.commandBufferCount, c.pCommandBuffers)
}

type VkCommandBufferUsageFlags VkFlags
type VkQueryControlFlags VkFlags
type VkCommandBufferInheritanceInfo struct {
	sType                VkStructureType
	pNext                *[0]byte
	renderPass           VkRenderPass
	subpass              uint32
	framebuffer          VkFramebuffer
	occlusionQueryEnable VkBool32
	queryFlags           VkQueryControlFlags
	pipelineStatistics   VkQueryPipelineStatisticFlags
}

func (g *VkCommandBufferInheritanceInfo) toC(c *C.VkCommandBufferInheritanceInfo) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	c.renderPass = C.VkRenderPass(g.renderPass)
	c.subpass = C.uint32_t(g.subpass)
	c.framebuffer = C.VkFramebuffer(g.framebuffer)
	{
		var _temp C.uint32_t
		_temp = C.uint32_t((uint32)(g.occlusionQueryEnable))
		c.occlusionQueryEnable = C.VkBool32(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.queryFlags)))
			_temp = C.VkFlags(_temp)
		}
		c.queryFlags = C.VkQueryControlFlags(_temp)
	}
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.pipelineStatistics)))
			_temp = C.VkFlags(_temp)
		}
		c.pipelineStatistics = C.VkQueryPipelineStatisticFlags(_temp)
	}
}
func (g *VkCommandBufferInheritanceInfo) fromC(c *C.VkCommandBufferInheritanceInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	g.renderPass = VkRenderPass(c.renderPass)
	g.subpass = uint32(c.subpass)
	g.framebuffer = VkFramebuffer(c.framebuffer)
	{
		var _temp uint32
		_temp = uint32((C.uint32_t)(c.occlusionQueryEnable))
		g.occlusionQueryEnable = VkBool32(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.queryFlags)))
			_temp = VkFlags(_temp)
		}
		g.queryFlags = VkQueryControlFlags(_temp)
	}
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.pipelineStatistics)))
			_temp = VkFlags(_temp)
		}
		g.pipelineStatistics = VkQueryPipelineStatisticFlags(_temp)
	}
}

type VkCommandBufferBeginInfo struct {
	sType            VkStructureType
	pNext            *[0]byte
	flags            VkCommandBufferUsageFlags
	pInheritanceInfo *VkCommandBufferInheritanceInfo
}

func (g *VkCommandBufferBeginInfo) toC(c *C.VkCommandBufferBeginInfo, _sa *stackAllocator) {
	c.sType = C.VkStructureType(g.sType)
	c.pNext = g.pNext
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkCommandBufferUsageFlags(_temp)
	}
	{
		c.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(_sa.alloc(C.sizeof_VkCommandBufferInheritanceInfo))
		g.pInheritanceInfo.toC(c.pInheritanceInfo)
	}
}
func (g *VkCommandBufferBeginInfo) fromC(c *C.VkCommandBufferBeginInfo) {
	g.sType = VkStructureType(c.sType)
	g.pNext = c.pNext
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.flags)))
			_temp = VkFlags(_temp)
		}
		g.flags = VkCommandBufferUsageFlags(_temp)
	}
	{
		if g.pInheritanceInfo == nil {
			g.pInheritanceInfo = new(VkCommandBufferInheritanceInfo)
		}
		g.pInheritanceInfo.fromC(c.pInheritanceInfo)
	}
}
func vkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (_ret VkResult) {
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
		pBeginInfo.toC(c.pBeginInfo, _sa)
	}
	c._ret = C.vkBeginCommandBuffer(c.commandBuffer, c.pBeginInfo)
	_ret = VkResult(c._ret)
	return
}
func vkEndCommandBuffer(commandBuffer VkCommandBuffer) (_ret VkResult) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		_ret          C.VkResult
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c._ret = C.vkEndCommandBuffer(c.commandBuffer)
	_ret = VkResult(c._ret)
	return
}

type VkCommandBufferResetFlags VkFlags

func vkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (_ret VkResult) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		flags         C.VkCommandBufferResetFlags
		_ret          C.VkResult
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(flags)))
			_temp = C.VkFlags(_temp)
		}
		c.flags = C.VkCommandBufferResetFlags(_temp)
	}
	c._ret = C.vkResetCommandBuffer(c.commandBuffer, c.flags)
	_ret = VkResult(c._ret)
	return
}
func vkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) {
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
func vkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport uint32, pViewports []VkViewport) {
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
	c.viewportCount = C.uint32_t(len(pViewports))
	{
		c.pViewports = (*C.VkViewport)(_sa.alloc(C.sizeof_VkViewport * uint(len(pViewports))))
		slice3 := (*[1 << 31]C.VkViewport)(unsafe.Pointer(c.pViewports))[:len(pViewports):len(pViewports)]
		for i3, _ := range pViewports {
			pViewports[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdSetViewport(c.commandBuffer, c.firstViewport, c.viewportCount, c.pViewports)
}
func vkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor uint32, pScissors []VkRect2D) {
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
	c.scissorCount = C.uint32_t(len(pScissors))
	{
		c.pScissors = (*C.VkRect2D)(_sa.alloc(C.sizeof_VkRect2D * uint(len(pScissors))))
		slice3 := (*[1 << 31]C.VkRect2D)(unsafe.Pointer(c.pScissors))[:len(pScissors):len(pScissors)]
		for i3, _ := range pScissors {
			pScissors[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdSetScissor(c.commandBuffer, c.firstScissor, c.scissorCount, c.pScissors)
}
func vkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth float32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		lineWidth     C.float
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.lineWidth = C.float(lineWidth)
	C.vkCmdSetLineWidth(c.commandBuffer, c.lineWidth)
}
func vkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor float32, depthBiasClamp float32, depthBiasSlopeFactor float32) {
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
func vkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants []float32) {
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
func vkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds float32, maxDepthBounds float32) {
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

type VkStencilFaceFlags VkFlags

func vkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		compareMask   C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(faceMask)))
			_temp = C.VkFlags(_temp)
		}
		c.faceMask = C.VkStencilFaceFlags(_temp)
	}
	c.compareMask = C.uint32_t(compareMask)
	C.vkCmdSetStencilCompareMask(c.commandBuffer, c.faceMask, c.compareMask)
}
func vkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		writeMask     C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(faceMask)))
			_temp = C.VkFlags(_temp)
		}
		c.faceMask = C.VkStencilFaceFlags(_temp)
	}
	c.writeMask = C.uint32_t(writeMask)
	C.vkCmdSetStencilWriteMask(c.commandBuffer, c.faceMask, c.writeMask)
}
func vkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference uint32) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		faceMask      C.VkStencilFaceFlags
		reference     C.uint32_t
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(faceMask)))
			_temp = C.VkFlags(_temp)
		}
		c.faceMask = C.VkStencilFaceFlags(_temp)
	}
	c.reference = C.uint32_t(reference)
	C.vkCmdSetStencilReference(c.commandBuffer, c.faceMask, c.reference)
}
func vkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet uint32, pDescriptorSets []VkDescriptorSet, pDynamicOffsets []uint32) {
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
	c.descriptorSetCount = C.uint32_t(len(pDescriptorSets))
	{
		c.pDescriptorSets = (*C.VkDescriptorSet)(_sa.alloc(C.sizeof_VkDescriptorSet * uint(len(pDescriptorSets))))
		slice3 := (*[1 << 31]C.VkDescriptorSet)(unsafe.Pointer(c.pDescriptorSets))[:len(pDescriptorSets):len(pDescriptorSets)]
		for i3, _ := range pDescriptorSets {
			slice3[i3] = C.VkDescriptorSet(pDescriptorSets[i3])
		}
	}
	c.dynamicOffsetCount = C.uint32_t(len(pDynamicOffsets))
	{
		c.pDynamicOffsets = (*C.uint32_t)(_sa.alloc(C.sizeof_uint32_t * uint(len(pDynamicOffsets))))
		slice3 := (*[1 << 31]C.uint32_t)(unsafe.Pointer(c.pDynamicOffsets))[:len(pDynamicOffsets):len(pDynamicOffsets)]
		for i3, _ := range pDynamicOffsets {
			slice3[i3] = C.uint32_t(pDynamicOffsets[i3])
		}
	}
	C.vkCmdBindDescriptorSets(c.commandBuffer, c.pipelineBindPoint, c.layout, c.firstSet, c.descriptorSetCount, c.pDescriptorSets, c.dynamicOffsetCount, c.pDynamicOffsets)
}

type VkIndexType int

const (
	VK_INDEX_TYPE_UINT16      VkIndexType = 0
	VK_INDEX_TYPE_UINT32      VkIndexType = 1
	VK_INDEX_TYPE_BEGIN_RANGE VkIndexType = VK_INDEX_TYPE_UINT16
	VK_INDEX_TYPE_END_RANGE   VkIndexType = VK_INDEX_TYPE_UINT32
	VK_INDEX_TYPE_RANGE_SIZE  VkIndexType = (VK_INDEX_TYPE_UINT32 - VK_INDEX_TYPE_UINT16 + 1)
	VK_INDEX_TYPE_MAX_ENUM    VkIndexType = 2147483647
)

func vkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
		indexType     C.VkIndexType
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	c.indexType = C.VkIndexType(indexType)
	C.vkCmdBindIndexBuffer(c.commandBuffer, c.buffer, c.offset, c.indexType)
}
func vkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding uint32, pBuffers []VkBuffer, pOffsets []VkDeviceSize) {
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
	c.bindingCount = C.uint32_t(len(pBuffers))
	{
		c.pBuffers = (*C.VkBuffer)(_sa.alloc(C.sizeof_VkBuffer * uint(len(pBuffers))))
		slice3 := (*[1 << 31]C.VkBuffer)(unsafe.Pointer(c.pBuffers))[:len(pBuffers):len(pBuffers)]
		for i3, _ := range pBuffers {
			slice3[i3] = C.VkBuffer(pBuffers[i3])
		}
	}
	{
		c.pOffsets = (*C.VkDeviceSize)(_sa.alloc(C.sizeof_VkDeviceSize * uint(len(pOffsets))))
		slice3 := (*[1 << 31]C.VkDeviceSize)(unsafe.Pointer(c.pOffsets))[:len(pOffsets):len(pOffsets)]
		for i3, _ := range pOffsets {
			{
				var _temp C.uint64_t
				_temp = C.uint64_t((uint64)(pOffsets[i3]))
				slice3[i3] = C.VkDeviceSize(_temp)
			}
		}
	}
	C.vkCmdBindVertexBuffers(c.commandBuffer, c.firstBinding, c.bindingCount, c.pBuffers, c.pOffsets)
}
func vkCmdDraw(commandBuffer VkCommandBuffer, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
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
func vkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount uint32, instanceCount uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
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
func vkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount uint32, stride uint32) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	c.drawCount = C.uint32_t(drawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndirect(c.commandBuffer, c.buffer, c.offset, c.drawCount, c.stride)
}
func vkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount uint32, stride uint32) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	c.drawCount = C.uint32_t(drawCount)
	c.stride = C.uint32_t(stride)
	C.vkCmdDrawIndexedIndirect(c.commandBuffer, c.buffer, c.offset, c.drawCount, c.stride)
}
func vkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
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
func vkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) {
	var c struct {
		commandBuffer C.VkCommandBuffer
		buffer        C.VkBuffer
		offset        C.VkDeviceSize
	}
	c.commandBuffer = C.VkCommandBuffer(commandBuffer)
	c.buffer = C.VkBuffer(buffer)
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(offset))
		c.offset = C.VkDeviceSize(_temp)
	}
	C.vkCmdDispatchIndirect(c.commandBuffer, c.buffer, c.offset)
}

type VkBufferCopy struct {
	srcOffset VkDeviceSize
	dstOffset VkDeviceSize
	size      VkDeviceSize
}

func (g *VkBufferCopy) toC(c *C.VkBufferCopy) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.srcOffset))
		c.srcOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.dstOffset))
		c.dstOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.size))
		c.size = C.VkDeviceSize(_temp)
	}
}
func (g *VkBufferCopy) fromC(c *C.VkBufferCopy) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.srcOffset))
		g.srcOffset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.dstOffset))
		g.dstOffset = VkDeviceSize(_temp)
	}
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.size))
		g.size = VkDeviceSize(_temp)
	}
}
func vkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, pRegions []VkBufferCopy) {
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
	c.regionCount = C.uint32_t(len(pRegions))
	{
		c.pRegions = (*C.VkBufferCopy)(_sa.alloc(C.sizeof_VkBufferCopy * uint(len(pRegions))))
		slice3 := (*[1 << 31]C.VkBufferCopy)(unsafe.Pointer(c.pRegions))[:len(pRegions):len(pRegions)]
		for i3, _ := range pRegions {
			pRegions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyBuffer(c.commandBuffer, c.srcBuffer, c.dstBuffer, c.regionCount, c.pRegions)
}

type VkImageSubresourceLayers struct {
	aspectMask     VkImageAspectFlags
	mipLevel       uint32
	baseArrayLayer uint32
	layerCount     uint32
}

func (g *VkImageSubresourceLayers) toC(c *C.VkImageSubresourceLayers) {
	{
		var _temp C.VkFlags
		{
			var _temp C.uint32_t
			_temp = C.uint32_t((uint32)((VkFlags)(g.aspectMask)))
			_temp = C.VkFlags(_temp)
		}
		c.aspectMask = C.VkImageAspectFlags(_temp)
	}
	c.mipLevel = C.uint32_t(g.mipLevel)
	c.baseArrayLayer = C.uint32_t(g.baseArrayLayer)
	c.layerCount = C.uint32_t(g.layerCount)
}
func (g *VkImageSubresourceLayers) fromC(c *C.VkImageSubresourceLayers) {
	{
		var _temp VkFlags
		{
			var _temp uint32
			_temp = uint32((C.uint32_t)((C.VkFlags)(c.aspectMask)))
			_temp = VkFlags(_temp)
		}
		g.aspectMask = VkImageAspectFlags(_temp)
	}
	g.mipLevel = uint32(c.mipLevel)
	g.baseArrayLayer = uint32(c.baseArrayLayer)
	g.layerCount = uint32(c.layerCount)
}

type VkImageCopy struct {
	srcSubresource VkImageSubresourceLayers
	srcOffset      VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffset      VkOffset3D
	extent         VkExtent3D
}

func (g *VkImageCopy) toC(c *C.VkImageCopy) {
	g.srcSubresource.toC(&c.srcSubresource)
	g.srcOffset.toC(&c.srcOffset)
	g.dstSubresource.toC(&c.dstSubresource)
	g.dstOffset.toC(&c.dstOffset)
	g.extent.toC(&c.extent)
}
func (g *VkImageCopy) fromC(c *C.VkImageCopy) {
	g.srcSubresource.fromC(&c.srcSubresource)
	g.srcOffset.fromC(&c.srcOffset)
	g.dstSubresource.fromC(&c.dstSubresource)
	g.dstOffset.fromC(&c.dstOffset)
	g.extent.fromC(&c.extent)
}
func vkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, pRegions []VkImageCopy) {
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
	c.regionCount = C.uint32_t(len(pRegions))
	{
		c.pRegions = (*C.VkImageCopy)(_sa.alloc(C.sizeof_VkImageCopy * uint(len(pRegions))))
		slice3 := (*[1 << 31]C.VkImageCopy)(unsafe.Pointer(c.pRegions))[:len(pRegions):len(pRegions)]
		for i3, _ := range pRegions {
			pRegions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyImage(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions)
}

type VkImageBlit struct {
	srcSubresource VkImageSubresourceLayers
	srcOffsets     [2]VkOffset3D
	dstSubresource VkImageSubresourceLayers
	dstOffsets     [2]VkOffset3D
}

func (g *VkImageBlit) toC(c *C.VkImageBlit) {
	g.srcSubresource.toC(&c.srcSubresource)
	for i, _ := range g.srcOffsets {
		g.srcOffsets[i].toC(&c.srcOffsets[i])
	}
	g.dstSubresource.toC(&c.dstSubresource)
	for i, _ := range g.dstOffsets {
		g.dstOffsets[i].toC(&c.dstOffsets[i])
	}
}
func (g *VkImageBlit) fromC(c *C.VkImageBlit) {
	g.srcSubresource.fromC(&c.srcSubresource)
	for i, _ := range g.srcOffsets {
		g.srcOffsets[i].fromC(&c.srcOffsets[i])
	}
	g.dstSubresource.fromC(&c.dstSubresource)
	for i, _ := range g.dstOffsets {
		g.dstOffsets[i].fromC(&c.dstOffsets[i])
	}
}
func vkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, pRegions []VkImageBlit, filter VkFilter) {
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
	c.regionCount = C.uint32_t(len(pRegions))
	{
		c.pRegions = (*C.VkImageBlit)(_sa.alloc(C.sizeof_VkImageBlit * uint(len(pRegions))))
		slice3 := (*[1 << 31]C.VkImageBlit)(unsafe.Pointer(c.pRegions))[:len(pRegions):len(pRegions)]
		for i3, _ := range pRegions {
			pRegions[i3].toC(&slice3[i3])
		}
	}
	c.filter = C.VkFilter(filter)
	C.vkCmdBlitImage(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions, c.filter)
}

type VkBufferImageCopy struct {
	bufferOffset      VkDeviceSize
	bufferRowLength   uint32
	bufferImageHeight uint32
	imageSubresource  VkImageSubresourceLayers
	imageOffset       VkOffset3D
	imageExtent       VkExtent3D
}

func (g *VkBufferImageCopy) toC(c *C.VkBufferImageCopy) {
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(g.bufferOffset))
		c.bufferOffset = C.VkDeviceSize(_temp)
	}
	c.bufferRowLength = C.uint32_t(g.bufferRowLength)
	c.bufferImageHeight = C.uint32_t(g.bufferImageHeight)
	g.imageSubresource.toC(&c.imageSubresource)
	g.imageOffset.toC(&c.imageOffset)
	g.imageExtent.toC(&c.imageExtent)
}
func (g *VkBufferImageCopy) fromC(c *C.VkBufferImageCopy) {
	{
		var _temp uint64
		_temp = uint64((C.uint64_t)(c.bufferOffset))
		g.bufferOffset = VkDeviceSize(_temp)
	}
	g.bufferRowLength = uint32(c.bufferRowLength)
	g.bufferImageHeight = uint32(c.bufferImageHeight)
	g.imageSubresource.fromC(&c.imageSubresource)
	g.imageOffset.fromC(&c.imageOffset)
	g.imageExtent.fromC(&c.imageExtent)
}
func vkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, pRegions []VkBufferImageCopy) {
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
	c.regionCount = C.uint32_t(len(pRegions))
	{
		c.pRegions = (*C.VkBufferImageCopy)(_sa.alloc(C.sizeof_VkBufferImageCopy * uint(len(pRegions))))
		slice3 := (*[1 << 31]C.VkBufferImageCopy)(unsafe.Pointer(c.pRegions))[:len(pRegions):len(pRegions)]
		for i3, _ := range pRegions {
			pRegions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyBufferToImage(c.commandBuffer, c.srcBuffer, c.dstImage, c.dstImageLayout, c.regionCount, c.pRegions)
}
func vkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, pRegions []VkBufferImageCopy) {
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
	c.regionCount = C.uint32_t(len(pRegions))
	{
		c.pRegions = (*C.VkBufferImageCopy)(_sa.alloc(C.sizeof_VkBufferImageCopy * uint(len(pRegions))))
		slice3 := (*[1 << 31]C.VkBufferImageCopy)(unsafe.Pointer(c.pRegions))[:len(pRegions):len(pRegions)]
		for i3, _ := range pRegions {
			pRegions[i3].toC(&slice3[i3])
		}
	}
	C.vkCmdCopyImageToBuffer(c.commandBuffer, c.srcImage, c.srcImageLayout, c.dstBuffer, c.regionCount, c.pRegions)
}
func vkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData []byte) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(dstOffset))
		c.dstOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(dataSize))
		c.dataSize = C.VkDeviceSize(_temp)
	}
	{
		c.pData = _sa.alloc(C.sizeof_void_pointer * uint(len(pData)))
		slice3 := (*[1 << 31]byte)(c.pData)[:len(pData):len(pData)]
		for i3, _ := range pData {
			slice3[i3] = pData[i3]
		}
	}
	C.vkCmdUpdateBuffer(c.commandBuffer, c.dstBuffer, c.dstOffset, c.dataSize, c.pData)
}
func vkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data uint32) {
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
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(dstOffset))
		c.dstOffset = C.VkDeviceSize(_temp)
	}
	{
		var _temp C.uint64_t
		_temp = C.uint64_t((uint64)(size))
		c.size = C.VkDeviceSize(_temp)
	}
	c.data = C.uint32_t(data)
	C.vkCmdFillBuffer(c.commandBuffer, c.dstBuffer, c.dstOffset, c.size, c.data)
}
