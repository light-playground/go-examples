package main

import (
	"time"
	"math/rand"
	"fmt"
	"strings"
)

type Result struct {
	res string
}

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		return Result{
			res: fmt.Sprintf("Kind: %s, query %q", kind, query),
		}
	}
}

//to avoid discarding results from slow severs?
//Replicate servers, send request to multiple replicas, and use the first response
func First(query string, replicas ...Search) Result {
	c := make(chan Result)

	for num := range replicas {
		go func(i int) {
			c <- replicas[i](query)
		}(num)
	}

	return <-c
}

//search one by one
func Google(query string) (results []Result) {

	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

	return
}

//simple goroutines
func GoogleV1(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()

	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}

	return
}

func GoogleV1WithTimeout(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

func GoogleV2(query string) (results []Result)  {
	c := make(chan Result)

	go func() {
		c <- First(query, fakeSearch("Web1"), fakeSearch("Web2"))
	}()
	go func() {
		c <- First(query, fakeSearch("Image1"), fakeSearch("Image2"))
	}()
	go func() {
		c <- First(query, fakeSearch("Video1"), fakeSearch("Video1"))
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(strings.Repeat("*", 30))
	start := time.Now()
	res := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	fmt.Println(strings.Repeat("*", 30))
	start = time.Now()
	res = GoogleV1("golang")
	elapsed = time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	fmt.Println(strings.Repeat("*", 30))
	start = time.Now()
	res = GoogleV1WithTimeout("golang")
	elapsed = time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	fmt.Println(strings.Repeat("*", 30))
	start = time.Now()
	res = GoogleV2("golang")
	elapsed = time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)
}
