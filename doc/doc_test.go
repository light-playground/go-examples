package doc

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println("1+2=", sum)
	//Output:
	//1+2=3
}

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Fail()
	}
}
