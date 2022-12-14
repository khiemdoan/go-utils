package randomutil

import (
	"bytes"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

var (
	AsciiLowercase        []rune
	AsciiUppercase        []rune
	AsciiLetters          []rune
	Digits                []rune
	AsciiLettersAndDigits []rune
	Punctuation           []rune
	Printable             []rune
	random                *rand.Rand
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

	seed := time.Now().UnixMilli()
	random = rand.New(rand.NewSource(seed))
}

// Random a number in range [start, stop)
func Rand[T constraints.Integer | constraints.Float](start, stop T) T {
	delta := stop - start
	return start + T(float64(delta)*(random.Float64()))
}

// Generate a random string with length n from set of vocaburary
func String(n int, vocab []rune) string {
	var bb bytes.Buffer
	bb.Grow(n)
	l := len(vocab)
	for i := 0; i < n; i++ {
		bb.WriteRune(vocab[Rand(0, l)])
	}
	return bb.String()
}

// Return a random element from the non-empty sequence
func Choice[T any](sequence []T) T {
	return sequence[Rand(0, len(sequence))]
}

// Return a k sized list of elements chosen from the population with replacement.
func Choices[T any](population []T, k int) []T {
	result := make([]T, k)
	for i := 0; i < k; i++ {
		result[i] = Choice(population)
	}
	return result
}

// Shuffle the sequence in place.
func Shuffle[T any](sequence []T) []T {
	result := make([]T, len(sequence))
	copy(result, sequence)
	random.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
