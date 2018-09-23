#pragma once
#include "vulkan/vulkan.h"

void callPFN_print(PFN_print f);

int callPFN_fooFunc(PFN_fooFunc f, void * arg0,  float arg1);

void callPFN_setArr(PFN_setArr f, size_t arg0,  int * arg1);

