package cgoexample

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// A C function that takes a string and prints it
void printString(char* str) {
    printf("C says: %s\n", str);
}
*/
import "C"
import "unsafe"

func Example4() {
	str := "Hello from Go!"
	// Convert the Go string to a C string
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr)) // Free the C string to avoid memory leaks

	// Call the C function with the C string
	C.printString(cStr)
}
