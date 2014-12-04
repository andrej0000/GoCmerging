package main

/*
#cgo LDFLAGS: -L. -lccode
#include "cCode.h"
*/
import "C"

import (
	"fmt"
)

//export GoFunc
func GoFunc() {
	fmt.Println("You just called GoFunc")
}

func main() {
	C.testFun()
}
