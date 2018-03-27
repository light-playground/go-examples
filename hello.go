package main

import (
	"fmt"
	"runtime"
	"math"
	"log"
	"os"
)

//variables

/*
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

//equals

i := 42
f := float64(i)
u := uint(f)
*/

//condition

func testCondition(x int) int {
	if x > 0 {
		return x
	} else {
		return 0
	}
}

//loops
/*
for i:= 1; i < 10; i++ {

}

for i < 10 {

}
for { //while (true)

}
*/

//array, slice, ranges

/*
var a [10]int

a[3] = 12 //set
i := a[3] //get

//declare and init
var b = [2]int{1, 2}
b := [2]int{1, 2}
a := [...]int{1, 2} //elipsis, 编译器指定长度
*/

//slice is array but length is not unspecified.

//var a []int //declare
func testSlice() {
	//var a = []int {1, 2, 3} //declare and initialize
	a := []int{1, 2, 4}

	a = append(a, 5, 6)

	//loop over an array or slice
	for k, v := range a {
		fmt.Printf("Key %d, val %d\n", k, v)
	}

	chars := []string{"test", "test1"}

	for k, v := range chars {
		fmt.Printf("Key %d, val %s\n", k, v)
	}
	//create via make
	b := make([]byte, 5, 5) //first length, second capacity(optional)
	b[1] = 12
	fmt.Println(b)

	//create slice from array
	x := [3]string{"test1", "test2", "test3"}
	s := x[:]
	fmt.Println(s)
}

//maps

type Vertex struct {
	lat float64
	lng float64
}

func testMap() {
	var m map[string]int
	m = make(map[string]int) //why call make

	m["age"] = 42
	fmt.Println(m["age"])

	delete(m, "age")
	elem, ok := m["age"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("age is deleted")
	}

	//struct map
	var m1 = map[string]Vertex{
		"Bell Labs": {40.121, -74.2121},
		"Google":    {37.121, -122.12121},
	}
	fmt.Println(m1)
}

//struct

//declare method on struct
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.lat*v.lat + v.lng*v.lng)
}

//use point
//the struct value is not copied for the method call
func (v *Vertex) Increase(incr float64) {
	v.lat += incr
}

func testStruct() {
	//1.creating
	var v = Vertex{1.12, 2.33}
	fmt.Println(v)
	var v1 = Vertex{lat: 12.212, lng: 2.3333} //with keys
	fmt.Println(v1)
	var v2 = []Vertex{
		{1.211, 2.112},
		{2.12, 3.311},
	}

	//print
	for _, val := range v2 {
		fmt.Printf("lat: %.2f, lng: %.2f\n", val.lat, val.lng)
	}

	fmt.Println(v.Abs())
	v.Increase(1.0)
	fmt.Println(v)
	//anonymous struct
	//cheaper and safer than using `map[string]interface{}`

	point := struct {
		X, Y int
		V    Vertex
	}{1, 2, Vertex{12.21, 31.11}}
	fmt.Println(point)

	//pointers(指针)

	p := Vertex{21.2, 2.2} //p is a Vertex

	q := &p //q is a pointer to a Vertex
	fmt.Println(q)
	r := &Vertex{2.1, 3.2} //r is also a pointer to a Vertex
	fmt.Println(r)

	var s *Vertex = new(Vertex)
	s.lat = 12.1
	s.lng = 1212.1
	fmt.Println(s)
}

//interface
type Awesomizer interface {
	Awesomize() string
}

type Foo struct {
}

//types implicitly satisfy an interface if they implement all required methods
func (foo Foo) Awesomize() string {
	return "Awesome!"
}

func testInterface() {
	var f = new(Foo)
	fmt.Println(f.Awesomize())
}

//embedding
//there is no subclassing in Go. Instead, there is interface and struct embedding

//ReadWriter implementations must satisfy both Reader and Writer
/*
type ReadWriter interface {
	Writer
	Reader
}
*/

type Server struct {
	Host string
	Port int
	*log.Logger
}

func testEmbedding() {
	server := &Server{
		"localhost",
		8080,
		log.New(os.Stdout, "[info]", log.Ldate|log.Ltime|log.Lshortfile),
	}
	// methods implemented on the embedded struct are passed through
	server.Printf("Host: %s, Port: %d\n", server.Host, server.Port)
}

//Errors
//There is no exception handling, Functions that might produce an error just declare and addtional return value
//of type Error

/*
Error interface:
~~~
type error Error {
	Error() string
}
~~~
 */

//Concurrency

//Goroutines
//lightweight threads(managed by Go, not OS threads)
func doStuff(s string)  {
	fmt.Println(s)
}
func testGoroutines() {
	//using a named function is a goroutine
	go doStuff("light")

	//anonymous inner function
	go func(x int) {

	}(42)
}

//Channels
/*
~~~
ch := make(chan int) //create a channel of type int
ch <- 42 //send a value to the channel ch
v := <-ch //Receive a value from ch
~~~
//Non-buffered channels block. Read blocks when no value is available, write blocks until there is read

//create a buffered channel. Writing to a buffered channels does not block if less than <buffer size> unread values have been written.
ch := make(chan int, 100)

close(ch) //closes the channel (only sender should close)

//read from channel and test if it has been closed
v, ok := <-ch

// if ok is false, channel has been closed

// select blocks on multiple channel operations, if one unblocks, the corresponding case is executed
func doStuff(channelOut, channelIn chan int) {
    select {
    case channelOut <- 42:
        fmt.Println("We could write to channelOut!")
    case x := <- channelIn:
        fmt.Println("We could read from channelIn")
    case <-time.After(time.Second * 1):
        fmt.Println("timeout")
    }
}
*/

//Channel Axioms

/*
1. a send to a nil channel blocks forever

~~~
var c chan string
c <- "Hello"
//fatal error: all goroutines are sleep - deadlock!
~~~

2. a receive from an nil channel blocks server

~~~
var c chan string
fmt.Println(<-c)
//fatal error: all goroutines are sleep - deadlock!
~~~

3. a send to closed channel panics

~~~
var c = make(chan string, 1)
c <- "Hello"
close(c)
c <- "Hello"
// panic: send on closed channel
~~~

4. A receive from a closed channel returns zero value immediately

~~~
var c = make(chan int, 2)
c <- 1
c <- 2
close(c)
for i := 0; i < 3;i++ {
	fmt.Printf("%d", <-c)
}
//1 2 0
~~~

 */
func testChannel() {

}


//make vs new

/**

~~~
sum(1, 2, 3)

nums := []int{1, 2, 3}
sum(nums...)
~~~

function as value:

~~~
add := func (a, b int) int {
	return a + b
}
fmt.Println(add(1, 2))
~~~

 */
func sum(args ...int) (total int) {
	total = 0
	for _, val := range args {
		total += val
	}

	return
}

//return the os flag
// -> os.GOOS
func sysFlag() (flag string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		flag = "darwin"
		break
	case "windows":
		flag = "windows"
		break
	case "linux":
		flag = "linux"
		break
	default:
		flag = "default"
	}

	return
}

func main() {
	fmt.Println("hello world")

	data := []int{10, 12, 13}

	fmt.Println(sum(data...))

	fmt.Println(sum(1, 2, 3, 4))

	fmt.Println(sysFlag())
	fmt.Println("---------slice---------")
	testSlice()
	fmt.Println("---------map---------")
	testMap()
	fmt.Println("---------struct---------")
	testStruct()
	fmt.Println("---------interface---------")
	testInterface()

	testEmbedding()

	testGoroutines()

	testChannel()
}
