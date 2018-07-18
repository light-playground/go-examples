//通过反射来获取struct的tag
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" validate:"min=1,max=12"`
	Age  int    `json:"age"`
}

func main() {
	var u User

	r := reflect.TypeOf(u)
	for i := 0; i < r.NumField(); i++ {
		sf := r.Field(i)
		fmt.Println(sf.Tag.Get("json"), sf.Tag.Get("validate"))
	}
}
