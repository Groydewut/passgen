package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

type Complexity struct {
	charset  string
	limit    int
	validate func(string) bool
}

var complexities = map[string]Complexity{
	"-easy": {
		charset:  charset_easy,
		limit:    8,
		validate: validateEasyPassword,
	},
	"-medium": {
		charset:  charset_medium,
		limit:    12,
		validate: validateMediumPassword,
	},
	"-hard": {
		charset:  charset_hard,
		limit:    20,
		validate: validateHardPassword,
	},
}

const charset_hard = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$*_=+?"

const charset_medium = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const charset_easy = "abcdefghijklmnopqrstuvwxyz0123456789"

func generatePassword(complexityKey string) (string, error) {
	complexity, exists := complexities[complexityKey]
	if !exists {
		return "", fmt.Errorf("invalid command")
	}

	charset := complexity.charset
	limit := complexity.limit
	maxBig := big.NewInt(int64(len(charset)))

	generationLimit := 1000

	for range generationLimit {
		var builder strings.Builder
		builder.Grow(limit)

		for range limit {
			r, err := rand.Int(rand.Reader, maxBig)
			if err != nil {
				return "", fmt.Errorf("failed to generate random password: %w", err)
			}
			builder.WriteByte(charset[r.Int64()])
		}
		password := builder.String()

		if complexity.validate(password) {
			return password, nil
		}
	}

	return "", fmt.Errorf("failed to generate password within %d attempts", generationLimit)
}

func validateEasyPassword(p string) bool {
	var hasLower, hasDigit bool

	for i := 0; i < len(p); i++ {
		ch := p[i]
		switch {
		case 'a' <= ch && ch <= 'z':
			hasLower = true
		case '0' <= ch && ch <= '9':
			hasDigit = true
		}
		if hasLower && hasDigit {
			return true
		}

	}
	return false
}

func validateMediumPassword(p string) bool {
	var hasLower, hasUpper, hasDigit bool

	for i := 0; i < len(p); i++ {
		ch := p[i]
		switch {
		case 'a' <= ch && ch <= 'z':
			hasLower = true
		case 'A' <= ch && ch <= 'Z':
			hasUpper = true
		case '0' <= ch && ch <= '9':
			hasDigit = true
		}
		if hasLower && hasUpper && hasDigit {
			return true
		}

	}
	return false
}

func validateHardPassword(p string) bool {
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for i := 0; i < len(p); i++ {
		ch := p[i]
		switch {
		case 'a' <= ch && ch <= 'z':
			hasLower = true
		case 'A' <= ch && ch <= 'Z':
			hasUpper = true
		case '0' <= ch && ch <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
		if hasLower && hasUpper && hasDigit && hasSpecial {
			return true
		}

	}
	return false
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <complexity> (options: -easy, -medium, -hard)", filepath.Base(os.Args[0]))
	}

	lvl := os.Args[1]

	password, err := generatePassword(lvl)

	if err != nil {
		log.Printf("Error generating password for level %s: %v", lvl, err)
		return
	}

	fmt.Printf("Level %s: %s (length: %d)\n", lvl, password, len(password))
}
