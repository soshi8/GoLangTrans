package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	aaa := []byte{115, 97, 109, 112}

	bbb := binary.LittleEndian.Uint32(aaa)
	fmt.Printf("LE1:%x\n", bbb)
}
