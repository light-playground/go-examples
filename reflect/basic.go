package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.4

	v := reflect.ValueOf(f)

	fmt.Println("type: ", v.Type())
	fmt.Println("Kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
