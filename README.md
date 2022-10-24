# randomutil
Library containing various random helpers

## Get a random integer

```go
number := randomutil.Rand(0, 20)
```

## Get a random float

```go
number := randomutil.Rand(-10.0, 10.0)
```

## Get a random string

```go
str = randomutil.String(20, randomutil.AsciiLetters)
```

## Get a random choice

```go
sequence := []string{"abc", "def", "ghi", "jkl"}
item := randomutil.Choice(sequence)
```

## Get multi random choices

```go
sequence := []string{"abc", "def", "ghi", "jkl"}
subsequence := randomutil.Choices(sequence, 2)
```

## Shuffle

```go
sequence := []string{"abc", "def", "ghi", "jkl"}
sequence = randomutil.Shuffle(sequence)
```