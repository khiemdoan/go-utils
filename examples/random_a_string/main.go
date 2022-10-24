package main

import (
	"fmt"

	"github.com/khiemdoan/randomutil"
)

func main() {
	str := randomutil.String(20, randomutil.AsciiLetters)
	fmt.Printf("random string: %s\n", str)

	str = randomutil.String(20, randomutil.AsciiLetters)
	fmt.Printf("random string: %s\n", str)
}
