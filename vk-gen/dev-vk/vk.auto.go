package vk
//#include "vulkan/vulkan.h"
//#include "bridges.auto.h"
import "C"

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

type PFN_print struct {
	Raw C.PFN_print
}

func (p PFN_print) Call() {
	C.callPFN_print(p.Raw)
}
func changeFunc(fun PFN_print) (_ret PFN_print) {
	var c struct {
		fun  C.PFN_print
		_ret C.PFN_print
	}
	c.fun = fun.Raw
	c._ret = C.changeFunc(c.fun)
	_ret.Raw = c._ret
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
func setPAbc(pAbc *Abc, pcAbc *Abc) {
	var c struct {
		pAbc  *C.Abc
		pcAbc *C.Abc
	}
	_sa := pool.take()
	defer pool.give(_sa)
	{
		c.pAbc = (*C.Abc)(_sa.alloc(C.sizeof_Abc))
		pAbc.toC(c.pAbc)
	}
	{
		c.pcAbc = (*C.Abc)(_sa.alloc(C.sizeof_Abc))
		pcAbc.toC(c.pcAbc)
	}
	C.setPAbc(c.pAbc, c.pcAbc)
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
func readString(str string) (_ret string) {
	var c struct {
		str  *C.char
		_ret *C.char
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.str = toCString(str, _sa)
	c._ret = C.readString(c.str)
	_ret = toGoString(c._ret)
	return
}
func writeString(count uint, str *byte) {
	var c struct {
		count C.size_t
		str   *C.char
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.count = C.size_t(count)
	{
		c.str = (*C.char)(_sa.alloc(C.sizeof_char))
		*c.str = C.char(*str)
	}
	C.writeString(c.count, c.str)
	*str = byte(*c.str)
}
func fooFunc(data unsafe.Pointer, pers float32) (_ret int32) {
	var c struct {
		data unsafe.Pointer
		pers C.float
		_ret C.int
	}
	c.data = data
	c.pers = C.float(pers)
	c._ret = C.fooFunc(c.data, c.pers)
	_ret = int32(c._ret)
	return
}

type PFN_fooFunc struct {
	Raw C.PFN_fooFunc
}

func (p PFN_fooFunc) Call(arg0 unsafe.Pointer, arg1 float32) (_ret int32) {
	var c struct {
		arg0 unsafe.Pointer
		arg1 C.float
		_ret C.int
	}
	c.arg0 = arg0
	c.arg1 = C.float(arg1)
	c._ret = C.callPFN_fooFunc(p.Raw, c.arg0, c.arg1)
	_ret = int32(c._ret)
	return
}
func getFooFunc() (_ret PFN_fooFunc) {
	var c struct{ _ret C.PFN_fooFunc }
	c._ret = C.getFooFunc()
	_ret.Raw = c._ret
	return
}
func setArr(imageCount uint, images *int32) {
	var c struct {
		imageCount C.size_t
		images     *C.int
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.imageCount = C.size_t(imageCount)
	{
		c.images = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.images = C.int(*images)
	}
	C.setArr(c.imageCount, c.images)
	*images = int32(*c.images)
}

type PFN_setArr struct {
	Raw C.PFN_setArr
}

func (p PFN_setArr) Call(arg0 uint, arg1 *int32) {
	var c struct {
		arg0 C.size_t
		arg1 *C.int
	}
	_sa := pool.take()
	defer pool.give(_sa)
	c.arg0 = C.size_t(arg0)
	{
		c.arg1 = (*C.int)(_sa.alloc(C.sizeof_int))
		*c.arg1 = C.int(*arg1)
	}
	C.callPFN_setArr(p.Raw, c.arg0, c.arg1)
	*arg1 = int32(*c.arg1)
}
func getSetArrFunc() (_ret PFN_setArr) {
	var c struct{ _ret C.PFN_setArr }
	c._ret = C.getSetArrFunc()
	_ret.Raw = c._ret
	return
}
