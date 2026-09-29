package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var complexities = map[string]int{
	"-easy":   8,
	"-medium": 12,
	"-hard":   20,
}



func numberGenerationComplexity(length string) (string, error) {

	limit, exists := complexities[length]
	if !exists {
		return "", fmt.Errorf("invalid command")
	}
	var builder strings.Builder
	builder.Grow(limit)
	
	maxBig := big.NewInt(10)

	for range limit{
		r, err := rand.Int(rand.Reader, maxBig)
		if err != nil {
			return "", fmt.Errorf("failed to generate random password: %w", err)
		}
		builder.WriteString(strconv.FormatInt(r.Int64(),10))
		
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


	fmt.Printf("Level %s: %s (length: %d)\n", lvl, password, len(password))
}