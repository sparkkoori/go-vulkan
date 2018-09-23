#include "vulkan/vulkan.h"

void print(void){};

int num(){return 0;};

int** movePointer(int *a, void *b){return 0;};

void setArray(int a[]){};

void setbigN(bigN n){};

void setFn(int (*fn)(void)){};

PFN_print changeFunc(PFN_print fun){return 0;};


void setAbc(Abc abc){};

void setPAbc(Abc *pAbc, const Abc *pcAbc){};

void setComplexAbc(ComplexAbc xabc){};

void setColor(Color color){};

void setHandle(Handle h){};

void read(const int* con){};

void setInt4(int4 n){};

void setCode(Code c){};

const char * readString(const char *str){return 0;};

void writeString(size_t count, char *str){};

int fooFunc (void* data, float pers){return 0};

PFN_FooFunc getFooFunc(){return 0;};

void setArr (size_t imageCount, int *images){};

PFN_setArr getSetArrFunc(){};
