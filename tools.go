package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// notify returns modified version of given string
// by adding application name at beginning.
func notify(s string) string {
	return fmt.Sprintf("%s: %s", appName, s)
}

// askUser asks the user a given question and
// returns true if user answered positively or
// false if negatively.
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

// pwd returns current working directory.
func pwd() (string, error) {
	ans, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return ans, nil
}

// unixify returns unixy version of given filename.
func unixify(filename string) string {
	return strings.ToLower(strings.ReplaceAll(filename, " ", "_"))
}

// usageVar returns replaceable version
// of given string s, to use in usage
func usageVar(s string) string {
	return fmt.Sprintf("{{%s}}", s)
}

// usage returns filled version of
// usasgeRaw, with key and values from
// config.go
func usage() string {
	var ans = usageRaw
	for key, val := range usageMap {
		ans = strings.ReplaceAll(ans, usageVar(key), val)
	}
	return ans
}
