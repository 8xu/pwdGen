package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	LettersAndNumbers          = 1
	LettersNumbersAndSpecialChars = 2
	AllChars                   = 3
)

func generatePassword(length int, complexity int) (string, error) {
	if length <= 0 || length > 4096 {
		return "", fmt.Errorf("invalid password length: %d", length)
	}
	if complexity < LettersAndNumbers || complexity > AllChars {
		return "", fmt.Errorf("invalid password complexity level: %d", complexity)
	}

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	if complexity >= LettersNumbersAndSpecialChars {
		charset += ")(*&^%$#@!~"
	}
	if complexity == AllChars {
		charset += "`-_=+[{]}\\|;:'\",<.>/?"
	}

	charsetLength := big.NewInt(int64(len(charset)))
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}
	return string(password), nil
}

func main() {
	var length int
	var complexity int
	flag.IntVar(&length, "l", 20, "length of password\n")
	flag.IntVar(&complexity, "c", LettersAndNumbers, "complexity of password\n1 - only letters and numbers\n2 - letters, numbers and special characters\n3 - letters, numbers and all special characters\n")
	flag.Parse()
	password, err := generatePassword(length, complexity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(password)
}
