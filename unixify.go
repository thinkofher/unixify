package main

import "strings"

func unixify(filename string) string {
	return strings.ToLower(strings.ReplaceAll(filename, " ", "_"))
}
