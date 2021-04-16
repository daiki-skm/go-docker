// export GOPATH=${pwd}

package main

import (
	"fmt"
	// "math/cmplx"
	// "time"
	// "math/rand"
	"math"
)

// func add(x, y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// var c, python, java bool

// var i, j int = 1, 2

// var (
// 	ToBe bool = false
// 	MaxInt uint64 = 1<<64 - 1
// 	z complex128 = cmplx.Sqrt(-5 + 12i)
// )

// const Pi = 3.14

// const (
// 	Big = 1 << 100
// 	Small = Big >> 99
// )

// func needInt(x int) int { return x*10 + 1 }

// func needFloat(x float64) float64 { return x*0.1 }

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// fmt.Println(sqrt(2), sqrt(-4))
	// sum := 1
	// for sum < 100 {
	// 	sum += sum
	// }
	// for ; sum < 100; {
	// 	sum += sum
	// }
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += 1
	// }
	// fmt.Println(sum)
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Small))
	// const World = "world"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")
	// const Truth = true
	// fmt.Println("Go rules?", Truth)
	// v := 42
	// fmt.Printf("v is of type %T\n", v)
	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	// fmt.Println("Hello, 世界")
	// fmt.Println("The time is", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(110))
	// fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// fmt.Println(math.Pi)
	// fmt.Println(add(42, 13))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(split(17))
	// var i int
	// fmt.Println(i, c, python, java)
	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)
	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, k, c, python, java)
}
