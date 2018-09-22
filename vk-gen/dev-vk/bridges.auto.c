#include "bridges.auto.h"

void callFUNC(FUNC f)
{
  return f();
};

int callFooFunc(FooFunc f, void * arg0,  float arg1)
{
  return f(arg0, arg1);
};

