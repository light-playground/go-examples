package main

import (
	"io"
	"fmt"
	"strings"
)

//io

//Reader
/*
所有实现了Read方法的类型都实现了io.Reader接口
~~~
type Reader interface {
	Read(p []byte) (n int, err error)
}
~~~
 */

//可以从任意地方读取数据, 只要实现了 io.Reader 接口
func readFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)

	n, err := reader.Read(p)

	if n > 0 {
		return p, nil
	}

	return p, err
}

//Writer
/*
所有实现了Write方法的类型都实现了io.Writer接口
~~~
type Writer interface {
	Write(p []byte) (n int, err error)
}
~~~
 */

func main() {
	reader := strings.NewReader("from string")
	data, err := readFrom(reader, 10)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))


	p := make([]byte, 6)

	n, err := reader.ReadAt(p, 2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %d\n", p, n)
}
