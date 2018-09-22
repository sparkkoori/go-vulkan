#pragma once

void print();

int num();

int** movePointer(int *a, void *b);

void setArray(int a[]);

typedef unsigned int bigN;

void setbigN(bigN n);

void setFn(int (*fn)(void));

typedef void (*FUNC)(void);

FUNC changeFunc(FUNC fun);

typedef struct Abc{
  int a;
  int b;
  int c;
} Abc;

void setAbc(Abc abc);

void setPAbc(Abc *pAbc);

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
