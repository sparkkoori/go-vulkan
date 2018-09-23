#pragma once
#include "stddef.h"

void print(void);

int num();

int** movePointer(int *a, void *b);

void setArray(int a[]);

typedef unsigned int bigN;

void setbigN(bigN n);

void setFn(int (*fn)(void));

typedef void (*PFN_print)(void);

PFN_print changeFunc(PFN_print fun);

typedef struct Abc{
  int a;
  int b;
  int c;
} Abc;

void setAbc(Abc abc);

void setPAbc(Abc *pAbc, const Abc *pcAbc);

typedef struct ComplexAbc{
  Abc abc;
  Abc *pAbc;
} ComplexAbc;

void setComplexAbc(ComplexAbc xabc);

typedef enum Color{
  Red = 1,
  Blue = 2,
  Yellow = 3,
  Green = (Red + Blue) - Yellow
} Color;

void setColor(Color color);

typedef struct Handle_T *Handle;

void setHandle(Handle h);

void read(const int* con);

typedef int int4[4];

void setInt4(int4 n);

typedef struct Code{
  int4 n0;
  float n1[2];
  long* n3[8];
} Code;

void setCode(Code code);

const char * readString(const char *str);

void writeString(size_t count, char *str);

typedef int (*PFN_fooFunc)(void* data, float pers);

int fooFunc (void* data, float pers);

PFN_fooFunc getFooFunc();

typedef void (*PFN_setArr)(size_t imageCount, int *images);

void setArr (size_t imageCount, int *images);

PFN_setArr getSetArrFunc();
