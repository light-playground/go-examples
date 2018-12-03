package main

/*
#include <stdio.h>
void hello() {
	printf("Hello, Cgo! -- From C world.\n");
}
*/
import "C"

func main() {
	C.hello()
}
