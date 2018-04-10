package main

import (
	"fmt"
)

func main() {

	var aaa string
	aaa = "sample string"
	bbb := []byte(aaa)

	fmt.Printf("bbb = %s\n", bbb)
}
