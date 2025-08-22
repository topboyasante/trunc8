package utils

import (
	"strings"
	"testing"
)

func TestGenerateURLCode(t *testing.T) {
	// Test that the function generates a 4-character code
	code := GenerateURLCode()
	if len(code) != 4 {
		t.Errorf("Expected code length to be 4, got %d", len(code))
	}

	// Test that the code only contains valid characters
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range code {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("Generated code contains invalid character: %c", char)
		}
	}

	// Test that multiple calls generate different codes (probabilistic test)
	// With 36^4 possible combinations, collisions should be rare
	codes := make(map[string]bool)
	collisions := 0
	iterations := 1000

	for i := 0; i < iterations; i++ {
		code := GenerateURLCode()
		if codes[code] {
			collisions++
		}
		codes[code] = true
	}

	// Allow for some collisions due to randomness, but not too many
	if collisions > iterations/10 {
		t.Errorf("Too many collisions: %d out of %d iterations", collisions, iterations)
	}
}

func TestGenerateURLCodeConsistency(t *testing.T) {
	// Test that all generated codes follow the expected pattern
	for i := 0; i < 100; i++ {
		code := GenerateURLCode()
		
		// Should be exactly 4 characters
		if len(code) != 4 {
			t.Errorf("Code %s has length %d, expected 4", code, len(code))
		}
		
		// Should only contain uppercase letters and digits
		for _, char := range code {
			if !((char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
				t.Errorf("Code %s contains invalid character: %c", code, char)
			}
		}
	}
}