package main

import (
	"fmt"
)

func main() {

	var aaa []byte
	//	aaa = []byte("sample string")
	aaa = []byte{115, 97, 109, 112, 108, 101}
	bbb := string(aaa)

	fmt.Printf("bbb = %s", bbb)
}
