package main

import (
	"flag"
)

var dirs bool
var hidden bool
var interactive bool
var verbose bool

func main() {
	flag.BoolVar(&dirs, "dirs", false, "include directories")
	flag.BoolVar(&hidden, "dirs", false, "include hidden files")
	flag.BoolVar(&interactive, "dirs", false, "asks before every rename")
	flag.BoolVar(&verbose, "dirs", false, "print current actions")
	flag.Parse()
}
