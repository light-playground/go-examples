package main

import (
	"os"
	"reflect"
	"strconv"
)

func print(args ...interface{}) {
	for _, arg := range args {
		switch val := reflect.ValueOf(arg); val.Kind() {
		case reflect.String:
			os.Stdout.WriteString(val.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(val.Int(), 10))
		}
	}
}

func main() {
	print("dsadsa\n", 1212)
}
