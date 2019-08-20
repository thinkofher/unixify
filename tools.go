package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func notify(s string) string {
	return fmt.Sprintf("%s: %s", appName, s)
}

func askUser(question string) (bool, error) {
	fmt.Printf("%s (y/N): ", notify(question))
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

// pwd returns current working directory
func pwd() (string, error) {
	ans, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return ans, nil
}
