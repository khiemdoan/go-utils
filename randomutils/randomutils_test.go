package randomutil

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestRandInt(t *testing.T) {
	start := -100
	stop := 100
	n := Rand(start, stop)
	fmt.Println(n)
	if n < start || n >= stop {
		t.Error("Rand was incorrect")
	}
}

func TestRandInt2(t *testing.T) {
	start := 0
	stop := 1
	n := Rand(start, stop)
	fmt.Println(n)
	if n != 0 {
		t.Error("Rand was incorrect")
	}
}

func TestRandFloat32(t *testing.T) {
	var start float32 = -100.0
	var stop float32 = 100.0
	n := Rand(start, stop)
	fmt.Println(n)
	if n < start || n >= stop {
		t.Error("Rand was incorrect")
	}
}

func TestRandFloat64(t *testing.T) {
	var start float64 = -100.0
	var stop float64 = 100.0
	n := Rand(start, stop)
	fmt.Println(n)
	if n < start || n >= stop {
		t.Error("Rand was incorrect")
	}
}

func TestString(t *testing.T) {
	length := 10
	str := String(length, AsciiLetters)
	fmt.Println(str)
	if len(str) != length {
		t.Error("String was incorrect length")
	}
	for _, c := range str {
		if !slices.Contains(AsciiLetters, c) {
			t.Error("String was incorrect characters")
		}
	}
}

func TestChoice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := Choice(seq)
	fmt.Println(n)
	if !slices.Contains(seq, n) {
		t.Error("Choice was incorrect")
	}
}

func TestChoices(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	length := 2
	choices := Choices(seq, length)
	if len(choices) != length {
		t.Error("Choices was incorrect")
	}
	for _, c := range choices {
		if !slices.Contains(seq, c) {
			t.Error("Choices was incorrect")
		}
	}
}

func TestShuffle(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := Shuffle(seq)
	if len(seq) != len(shuffled) {
		t.Error("Shuffle was incorrect")
	}
	for _, c := range shuffled {
		if !slices.Contains(seq, c) {
			fmt.Println(c)
			t.Error("Shuffle was incorrect")
		}
	}
}
