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

func main() {
	// pointers(42, 2701)
	// structs(1, 2)
	// pToS(1, 2)
	// fmt.Println(v1, v2, v3, p)
	// arrays()
	slices()
}