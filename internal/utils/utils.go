package utils

import (
	"math/rand/v2"
)

func GenerateURLCode() string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""

	for i := 0; i < 4; i++ {
		randomIndex := rand.IntN(len(charset))
		result += string(charset[randomIndex])
	}

	return result
}
