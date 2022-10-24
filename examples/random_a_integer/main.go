package main

import (
	"fmt"

	"github.com/khiemdoan/randomutil"
)

func main() {
	number := randomutil.Rand(0, 20)
	fmt.Printf("random integer: %d\n", number)

	number = randomutil.Rand(0, 20)
	fmt.Printf("random integer: %d\n", number)
}
