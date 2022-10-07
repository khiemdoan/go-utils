package randomutil

import (
	"bytes"
	"math/rand"
)

var (
	AsciiLowercase        []rune
	AsciiUppercase        []rune
	AsciiLetters          []rune
	Digits                []rune
	AsciiLettersAndDigits []rune
	Punctuation           []rune
	Printable             []rune
)

func init() {
	AsciiLowercase = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	AsciiUppercase = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	AsciiLetters = append(AsciiLetters, AsciiLowercase...)
	AsciiLetters = append(AsciiLetters, AsciiUppercase...)
	Digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	AsciiLettersAndDigits = append(AsciiLettersAndDigits, AsciiLetters...)
	AsciiLettersAndDigits = append(AsciiLettersAndDigits, Digits...)
	Punctuation = []rune{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
	Printable = append(Printable, AsciiLowercase...)
	Printable = append(Printable, AsciiUppercase...)
	Printable = append(Printable, Digits...)
	Printable = append(Printable, Punctuation...)
}

// Random an integer in range [start, stop)
func RandInt(start, stop int) int {
	delta := stop - start
	return start + rand.Intn(delta)
}

// Random a float in range [start, stop)
func RandFloat32(start, stop float32) float32 {
	delta := stop - start
	return start + delta*rand.Float32()
}

// Random a float in range [start, stop)
func RandFloat64(start, stop float64) float64 {
	delta := stop - start
	return start + delta*rand.Float64()
}

// Generate a random string with length n from set of vocaburary
func String(n int, vocab []rune) string {
	var bb bytes.Buffer
	bb.Grow(n)
	l := len(vocab)
	for i := 0; i < n; i++ {
		bb.WriteRune(vocab[RandInt(0, l)])
	}
	return bb.String()
}

// Return a random element from the non-empty sequence
func Choice[T any](sequence []T) T {
	return sequence[RandInt(0, len(sequence))]
}

// Return a k sized list of elements chosen from the population with replacement.
func Choices[T any](population []T, k int) []T {
	result := []T{}
	for i := 0; i < k; i++ {
		result = append(result, Choice(population))
	}
	return result
}

// Shuffle the sequence in place.
func Shuffle[T any](sequence []T) []T {
	length := len(sequence)
	newSeq := make([]T, length)
	result := make([]T, length)

	copy(newSeq, sequence)
	for i := 0; i < length; i++ {
		j := RandInt(0, len(newSeq))
		result[i] = newSeq[j]
		newSeq[j] = newSeq[len(newSeq)-1]
		newSeq = newSeq[:len(newSeq)-1]
	}

	return result
}
