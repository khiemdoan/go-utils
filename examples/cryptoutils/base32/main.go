package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/cryptoutils"
)

func main() {
	base32 := cryptoutils.Base32Encode([]byte("abc"))
	fmt.Printf("Base32: %s\n", base32)
}
