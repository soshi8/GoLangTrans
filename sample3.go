package main

import (
	"fmt"
)

func main() {

	var aaa [6]byte
	var bbb []byte
	aaa = [6]byte{115, 97, 109, 112, 108, 101}
	bbb = aaa[:]

	fmt.Printf("bbb = %s", bbb[:])
}
