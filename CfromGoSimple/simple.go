package main

/*
#include <stdio.h>

void foo() {
	printf("You just called foo\n");
}
*/
import "C"

func main() {
	C.foo()
}
