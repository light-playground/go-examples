package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)

	fmt.Println("type of p:", p.Type())
	fmt.Println("settabillity of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settabillity of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
