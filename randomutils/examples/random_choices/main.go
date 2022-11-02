package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutil"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	subsequence := randomutil.Choices(sequence, 2)
	fmt.Printf("random choice: %v\n", subsequence)
}
