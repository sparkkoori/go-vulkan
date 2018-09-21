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

void sumAbc(Abc abc, Abc *pAbc);
