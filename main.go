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

var complexities = map[string]int{
	"-easy":   8,
	"-medium": 12,
	"-hard":   20,
}

const charset ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$*_=+?"

func numberGenerationComplexity(complexityKey string) (string, error) {

	limit, exists := complexities[complexityKey]
	if !exists {
		return "", fmt.Errorf("invalid command")
	}
	var builder strings.Builder
	builder.Grow(limit)
	
	maxBig := big.NewInt(int64(len(charset)))

	for range limit{
		r, err := rand.Int(rand.Reader, maxBig)
		if err != nil {
			return "", fmt.Errorf("failed to generate random password: %w", err)
		}
		builder.WriteString(string(charset[r.Int64()]))
	}

	return builder.String(), nil
}
func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <complexity> (options: -easy, -medium, -hard)", filepath.Base(os.Args[0]))
	}
	
	lvl := os.Args[1]

	password,err := numberGenerationComplexity(lvl)


	if err != nil{
		log.Printf("Error generating password for level %s: %v", lvl, err)
		return}


	fmt.Printf("Level %s: %s (complexityKey: %d)\n", lvl, password, len(password))
}