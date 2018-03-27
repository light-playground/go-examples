package main

import (
	"sync"
	"time"
	"fmt"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) inc(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.v[key]++
}

func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 100; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("somekey"))
}
