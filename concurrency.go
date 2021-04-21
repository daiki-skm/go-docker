package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channels() {
	s := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x, y := <-c, <-c

	fmt.Println(x,y,x+y)
}

func bufferChan() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	// goroutines()
	// channels()
	bufferChan()
}