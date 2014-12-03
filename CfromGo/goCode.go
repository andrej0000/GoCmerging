package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lccode
#include <cCode.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	C.cFunction()

	var x, y int
	x = 5
	y = 6

	fmt.Println(int(C.cAdd(C.int(x), C.int(y))))
}
