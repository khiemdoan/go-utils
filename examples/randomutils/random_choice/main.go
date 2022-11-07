package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutils"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	item := randomutils.Choice(sequence)
	fmt.Printf("random choice: %s\n", item)

	item = randomutils.Choice(sequence)
	fmt.Printf("random choice: %s\n", item)
}
