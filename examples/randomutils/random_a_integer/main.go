package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	number := randomutils.Rand(0, 20)
	fmt.Printf("random integer: %d\n", number)

	number = randomutils.Rand(0, 20)
	fmt.Printf("random integer: %d\n", number)
}
