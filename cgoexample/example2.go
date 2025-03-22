package cgoexample

/*
#include <stdio.h>

// A C function that takes an integer and returns its square
int square(int n) {
    return n * n;
}
*/
import "C"
import "fmt"

func Example2() {
	num := 5
	// Call the C function and pass a Go integer
	result := C.square(C.int(num))
	fmt.Printf("Square of %d is %d\n", num, result)
}
