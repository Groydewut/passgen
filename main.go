package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

const EASY = 8
const MEDIUM = 12
const HARD = 20

func numberGenerationComplexity(length string) string {

	var strSlice []string
	switch length {
	case "-easy":
		for range EASY {
			r, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				log.Println(err)
			}
			numStr := strconv.FormatInt(r.Int64(), 10)
			strSlice = append(strSlice, numStr)
		}
		result := strings.Join(strSlice, "")
		return result

	case "-medium":
		for range MEDIUM {
			r, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				log.Println(err)
			}
			numStr := strconv.FormatInt(r.Int64(), 10)
			strSlice = append(strSlice, numStr)
		}
		result := strings.Join(strSlice, "")
		return result
	case "-hard":
		for range HARD {
			r, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				log.Println(err)
			}
			numStr := strconv.FormatInt(r.Int64(), 10)
			strSlice = append(strSlice, numStr)
		}
		result := strings.Join(strSlice, "")
		return result
	default:
		return "invalid command"

	}
}
func main() {

	fmt.Println(numberGenerationComplexity("-easy"))
	fmt.Println(numberGenerationComplexity("-medium"))
	fmt.Println(numberGenerationComplexity("-hard"))
	fmt.Println(len(numberGenerationComplexity("-easy")))
	fmt.Println(len(numberGenerationComplexity("-medium")))
	fmt.Println(len(numberGenerationComplexity("-hard")))
}
