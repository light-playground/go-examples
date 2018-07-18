//并发控制：waitGroup
//适用于多个goroutine同时做一件事情
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1 Done")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 Done")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("all task finished")
}
