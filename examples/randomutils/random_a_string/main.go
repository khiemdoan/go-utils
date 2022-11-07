package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	str := randomutils.String(20, randomutils.AsciiLetters)
	fmt.Printf("random string: %s\n", str)

	str = randomutils.String(20, randomutils.AsciiLetters)
	fmt.Printf("random string: %s\n", str)
}
