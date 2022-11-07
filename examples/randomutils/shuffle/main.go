package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	sequence = randomutils.Shuffle(sequence)
	fmt.Printf("shuffle: %v\n", sequence)
}
