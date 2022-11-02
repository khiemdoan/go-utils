package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutil"
)

func main() {
	number := randomutil.Rand(-10.0, 10.0)
	fmt.Printf("random integer: %f\n", number)

	number = randomutil.Rand(-10.0, 10.0)
	fmt.Printf("random integer: %f\n", number)
}
