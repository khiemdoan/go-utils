package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/cryptoutils"
)

func main() {
	md5 := cryptoutils.Md5([]byte("abc"))
	fmt.Printf("MD5: %s\n", md5)
}
