package cgoexample

/*
#include <stdio.h>

// Define a simple C function
void sayHello() {
    printf("Hello from C!\n");
}
*/
import "C"

func Example1() {
	// Call the C function from Go
	C.sayHello()
}
