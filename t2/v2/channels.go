package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	sep := len(s) / 2
	//chan
	c := make(chan int)

	go sum(s[:sep], c)
	go sum(s[sep:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
