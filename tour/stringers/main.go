package main

import (
	"crypto/md5"
	"fmt"
)

type Office int

const (
	Boston Office = iota
	NewYork
)

var officePlaces = [2]string{
	"test",
	"aaaa",
}

func (o Office) String() string {
	return "Google, " + officePlaces[o]
}

func main() {
	fmt.Printf("Hello, %s\n", Boston)

	m := md5.New()
	m.Write([]byte(`this is a test`))

	fmt.Println(string(m.Sum([]byte(`a`))))

}
