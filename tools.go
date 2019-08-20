package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func askUser(question string) (bool, error) {
	fmt.Printf("%s: %s (y/N): ", appName, question)
	reader := bufio.NewReader(os.Stdin)

	char, _, err := reader.ReadRune()
	if err != nil {
		return false, err
	}

	if unicode.ToLower(char) == 'y' {
		return true, nil
	} else {
		return false, nil
	}
}

func pwd() (string, error) {
	ans, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return ans, nil
}
