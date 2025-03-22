package cgotest

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"
import "fmt"

func Example3() {
	x := 2.0
	// Call the C math library's sqrt function
	result := C.sqrt(C.double(x))
	fmt.Printf("Square root of %.2f is %.2f\n", x, result)
}
