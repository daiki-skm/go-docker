package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) AbsCon() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methodsCon() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.AbsCon())
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func pointerReceiver() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func indirection() {
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4,3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func indirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

func choose() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

type Abser interface {
	Abs() float64
}

func interfaces() {
	var a Abser
	// f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	// a = f
	a = &v
	// a = v

	fmt.Println(a.Abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func interfacesVal() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func interfaces10() {
// 	var i I = T{"hello"}
// 	i.M()
// }

func withNil() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"&hello"}
	describe(i)
	i.M()
}

func empty() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func assertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

func main() {
	// methods()
	// methodsCon()
	// pointerReceiver()
	// indirection()
	// indirection2()
	// choose()
	// interfaces()
	// interfaces10()
	// interfacesVal()
	// withNil()
	// empty()
	assertions()
}