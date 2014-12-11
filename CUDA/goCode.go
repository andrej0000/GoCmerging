package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lkernel
#include <kernel.h>
*/
import "C"

import (
	"fmt"
)
func main() {
	fmt.Println(int(C.foo(2048)))
}
