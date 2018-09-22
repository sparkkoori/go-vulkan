package vk

//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"
import "unsafe"

func print() {
	C.print()
}
func num() (_ret int32) {
	var c struct{ _ret C.int }
	c._ret = C.num()
	_ret = int32(c._ret)
	return
}
func movePointer(a *int32, b unsafe.Pointer) (_ret **int32) {
	var c struct {
		a    *C.int
		b    unsafe.Pointer
		_ret **C.int
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.a = C.int(*a)
	}
	c.b = b
	c._ret = C.movePointer(c.a, c.b)
	{
		if _ret == nil {
			_ret = new(*int32)
		}
		{
			if *_ret == nil {
				*_ret = new(int32)
			}
			**_ret = int32(**c._ret)
		}
	}
	*a = int32(*c.a)
	return
}
func setArray(a *int32) {
	var c struct{ a *C.int }
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.a = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.a = C.int(*a)
	}
	C.setArray(c.a)
	*a = int32(*c.a)
}

type bigN uint32

func setbigN(n bigN) {
	var c struct{ n C.bigN }
	{
		var _temp C.uint
		_temp = C.uint((uint32)(n))
		c.n = C.bigN(_temp)
	}
	C.setbigN(c.n)
}
func setFn(fn *[0]byte) {
	var c struct{ fn *[0]byte }
	c.fn = fn
	C.setFn(c.fn)
}

type FUNC C.FUNC

func changeFunc(fun FUNC) (_ret FUNC) {
	var c struct {
		fun  C.FUNC
		_ret C.FUNC
	}
	c.fun = C.FUNC(fun)
	c._ret = C.changeFunc(c.fun)
	_ret = FUNC(c._ret)
	return
}

type Abc struct {
	a int32
	b int32
	c int32
}

func (g *Abc) toC(c *C.Abc) {
	c.a = C.int(g.a)
	c.b = C.int(g.b)
	c.c = C.int(g.c)
}
func (g *Abc) fromC(c *C.Abc) {
	g.a = int32(c.a)
	g.b = int32(c.b)
	g.c = int32(c.c)
}
func setAbc(abc Abc) {
	var c struct{ abc C.Abc }
	abc.toC(&c.abc)
	C.setAbc(c.abc)
}
func setPAbc(pAbc *Abc) {
	var c struct{ pAbc *C.Abc }
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pAbc = (*C.Abc)(_sa.alloc(C.sizeof_Abc))
		pAbc.toC(c.pAbc)
	}
	C.setPAbc(c.pAbc)
	pAbc.fromC(c.pAbc)
}

type ComplexAbc struct {
	abc  Abc
	pAbc *Abc
}

func (g *ComplexAbc) toC(c *C.ComplexAbc, _sa *stackAllocator) {
	g.abc.toC(&c.abc)
	{
		c.pAbc = (*C.Abc)(_sa.alloc(C.sizeof_Abc))
		g.pAbc.toC(c.pAbc)
	}
}
func (g *ComplexAbc) fromC(c *C.ComplexAbc) {
	g.abc.fromC(&c.abc)
	{
		if g.pAbc == nil {
			g.pAbc = new(Abc)
		}
		g.pAbc.fromC(c.pAbc)
	}
}
func setComplexAbc(xabc ComplexAbc) {
	var c struct{ xabc C.ComplexAbc }
	_sa := pool.take()
	defer pool.give(_sa)
	xabc.toC(&c.xabc, _sa)
	C.setComplexAbc(c.xabc)
	{
		xabc.pAbc.fromC(c.xabc.pAbc)
	}
}

type Color int

const (
	Red    int = 1
	Blue   int = 2
	Yellow int = 3
	Green  int = (Red + Blue) - Yellow
)

func setColor(color int) {
	var c struct{ color C.Color }
	c.color = C.Color(color)
	C.setColor(c.color)
}

type Handle C.Handle

func setHandle(h Handle) {
	var c struct{ h C.Handle }
	c.h = C.Handle(h)
	C.setHandle(c.h)
}
func read(con *int32) {
	var c struct{ con *C.int }
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.con = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.con = C.int(*con)
	}
	C.read(c.con)
	*con = int32(*c.con)
}
func setInt4(n *int32) {
	var c struct{ n *C.int }
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.n = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.n = C.int(*n)
	}
	C.setInt4(c.n)
	*n = int32(*c.n)
}

type int4 [4]int32
type Code struct {
	n0 int4
	n1 [2]float32
	n3 [8]*int32
}

func (g *Code) toC(c *C.Code, _sa *stackAllocator) {
	{
		var _temp [4]C.int
		for i, _ := range ([4]int32)(g.n0) {
			_temp[i] = C.int(([4]int32)(g.n0)[i])
		}
		c.n0 = C.int4(_temp)
	}
	for i, _ := range g.n1 {
		c.n1[i] = C.float(g.n1[i])
	}
	for i, _ := range g.n3 {
		{
			c.n3[i] = (*C.long)(_sa.alloc(C.sizeof_long))
			*c.n3[i] = C.long(*g.n3[i])
		}
	}
}
func (g *Code) fromC(c *C.Code) {
	{
		var _temp [4]int32
		for i, _ := range _temp {
			_temp[i] = int32(([4]C.int)(c.n0)[i])
		}
		g.n0 = int4(_temp)
	}
	for i, _ := range g.n1 {
		g.n1[i] = float32(c.n1[i])
	}
	for i, _ := range g.n3 {
		{
			if g.n3[i] == nil {
				g.n3[i] = new(int32)
			}
			*g.n3[i] = int32(*c.n3[i])
		}
	}
}
func setCode(code Code) {
	var c struct{ code C.Code }
	_sa := pool.take()
	defer pool.give(_sa)
	code.toC(&c.code, _sa)
	C.setCode(c.code)
	{
		for i, _ := range code.n3 {
			*code.n3[i] = int32(*c.code.n3[i])
		}
	}
}
