package algorith

import (
	"testing"
	"fmt"
)

func TestFib2(t *testing.T) {
	var (
		in     = 7
		expect = 13
	)

	actual := Fib(in)
	if actual != expect {
		t.Errorf("Fid(%d) = %d; expected %d", in, actual, expect)
	}

}

func TestFib(t *testing.T) {
	res := Fib(1)

	if res != 1 {
		t.Error("aaaa")
	}

	res = Fib(2)

	if res != 1 {
		t.Failed()
	}

	res = Fib(3)

	if res != 2 {
		t.Failed()
	}

	res = Fib(4)
	if res != 3 {
		t.Failed()
	}
}

//table-driven
func TestFib3(t *testing.T) {
	var fibTests = []struct {
		in       int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)

		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

func ExampleFib() {
	res := Fib(10)
	fmt.Println(res)
	// Output:
	// 55
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++  {
		Fib(1000)
	}
}

func ExampleFoo_Awesomize() {
	s := []string{"a", "b", "c", "d"}

	fmt.Println(s)
	fmt.Println(s[2:])
	fmt.Println(s[2])

	//Output: aa
}