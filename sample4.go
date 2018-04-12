package main

import (
	"fmt"
)

func main() {

	var aaa []byte
	var bbb = make([]byte, 6)
	aaa = []byte{115, 97, 109, 112, 108, 101}
	copy(bbb, aaa[:])

	fmt.Printf("bbb = %s", bbb)
}
