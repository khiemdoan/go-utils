package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	subsequence := randomutils.Choices(sequence, 2)
	fmt.Printf("random choice: %v\n", subsequence)
}
