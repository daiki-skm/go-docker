package main

import "fmt"

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
	createSlice()
}