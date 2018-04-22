package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	var aaa uint16 = 4011
	bbb := make([]byte, 4)

	binary.LittleEndian.PutUint16(bbb, aaa)
	fmt.Printf("LE1:%x\n", bbb)
}
