package main

import (
	"fmt"

	"github.com/khiemdoan/go-utils/randomutil"
)

func main() {
	sequence := []string{"abc", "def", "ghi", "jkl"}
	item := randomutil.Choice(sequence)
	fmt.Printf("random choice: %s\n", item)

	item = randomutil.Choice(sequence)
	fmt.Printf("random choice: %s\n", item)
}