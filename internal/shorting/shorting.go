package shorting

import (
	"math/rand"
	"strings"
)

const (
	Size     = 10
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func GenerateShortLink() string {
	length := len(Alphabet)
	var result strings.Builder
	result.Grow(Size)
	for i := 0; i < Size; i++ {
		result.WriteByte(Alphabet[rand.Intn(length)])
	}
	return result.String()
}
