package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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

func main() {
	// methods()
	// methodsCon()
	// pointerReceiver()
	indirection()
}