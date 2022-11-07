package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	number := randomutils.Rand(-10.0, 10.0)
	fmt.Printf("random integer: %f\n", number)

	number = randomutils.Rand(-10.0, 10.0)
	fmt.Printf("random integer: %f\n", number)
}
