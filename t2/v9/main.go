package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	//c := boring("boring!")
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("You say: %q\n", <-c)
	//}

	//joe := boring("Joe")
	//ann := boring("Ann")
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-joe)
	//	fmt.Println(<-ann)
	//}

	b := fanIn(boring("Joe"), boring("Ann"))

	for j := 0; j < 10; j++ {
		fmt.Println(<-b)
	}
}
