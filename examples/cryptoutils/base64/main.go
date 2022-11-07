package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/cryptoutils"
)

func main() {
	base64 := cryptoutils.Base64Encode([]byte("abc"))
	fmt.Printf("Base64: %s\n", base64)
}
