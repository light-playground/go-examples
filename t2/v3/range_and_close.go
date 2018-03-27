package main

import "fmt"

// a sender can close a channel to indicate that no more values
// will be sent.
// Receivers can test whether a channel has been closed by assigning a
// second parameter
// v, ok := <-ch
// ok is `false` if there are no more value to receives and the channel is closed.

// Note: only the sender should close a channel.

func fibnacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibnacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
