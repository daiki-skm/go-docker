package main

import (
	"fmt"
	"strings"
	"math"
)

func pointers(i, j int) {
	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X, Y int
}

func structs(i, j int) {
	v := Vertex{i,j}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)
}

func pToS(i, j int) {
	v := Vertex{i,j}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

func slicesArray() {
	names := [4]string{
		"john",
		"daiki",
		"sakuma",
		"my",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{1,2,3,4,4}
	fmt.Println(q)

	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)
}

func sliceDefault() {
	s := []int{2,3,4,44,45,5}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLengCapa() {
	s := []int{2,3,4,5,5,6,33}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilslice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func printSliceEx(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func createSlice() {
	a := make([]int, 5)
	printSliceEx("a", a)

	b := make([]int, 0, 5)
	printSliceEx("b", b)

	c := b[:2]
	printSliceEx("c", c)

	d := c[2:5]
	printSliceEx("d", d)
}

func ticTac() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "X"
	board[1][2] = "X"
	board[1][0] = "X"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceappend() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 3)
	printSlice(s)

	s = append(s, 34)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func slicerange() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i ,v)
	}
}

func rangeCon() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}
}

type VertexEx struct {
	Lat, Long float64
}

func maps() {
	var m map[string]VertexEx
	m = make(map[string]VertexEx)
	m["daiki sakuma"] = VertexEx{
		40.68433, -74.39967,
	}
	fmt.Println(m["daiki sakuma"])
}

func mapsLiteral() {
	m := map[string]VertexEx{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mutateMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcVal(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func main() {
	// pointers(42, 2701)
	// structs(1, 2)
	// pToS(1, 2)
	// fmt.Println(v1, v2, v3, p)
	// arrays()
	// slices()
	// slicesArray()
	// sliceLiterals()
	// sliceDefault()
	// sliceLengCapa()
	// nilslice()
	// createSlice()
	// ticTac()
	// sliceappend()
	// slicerange()
	// rangeCon()
	// maps()
	// mapsLiteral()
	// mutateMaps()
	// funcVal()
	closures()
}