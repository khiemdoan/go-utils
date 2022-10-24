package main

import (
	"fmt"

	"github.com/khiemdoan/randomutil"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	sequence = randomutil.Shuffle(sequence)
	fmt.Printf("shuffle: %v\n", sequence)
}
