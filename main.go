package main

import (
	"fmt"
	// "time"
	// "math/rand"
	// "math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, 世界")
	// fmt.Println("The time is", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(110))
	// fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// fmt.Println(math.Pi)
	// fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
