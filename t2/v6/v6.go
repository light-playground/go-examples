package main

import (
	"math/rand"
	"fmt"
)

//Tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}

	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}

	s += fmt.Sprint(t.Value)

	if t.Right != nil {
		s += " " + t.Right.String()
	}

	return "(" + s + ")"
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}

	return t
}

func New(k int) *Tree {
	var t *Tree

	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

//Walk the tree
func Walk(t *Tree, ch chan int) {
	var walker func(t *Tree)

	walker = func(t *Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)

	close(ch)
}

//whether two tree contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println("1 and 1 same: ", Same(New(1), New(1)))
	fmt.Println("1 and 2 same: ", Same(New(1), New(2)))
}
