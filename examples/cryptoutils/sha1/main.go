package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/cryptoutils"
)

func main() {
	sha1 := cryptoutils.Sha1([]byte("abc"))
	fmt.Printf("MD5: %s\n", sha1)
}
