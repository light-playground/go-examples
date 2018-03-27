package main

import "fmt"

// select statement lets a goroutines wait on multiple communication operations.
// a select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.

func fibnacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return //break
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibnacci(c, quit)
}
