package main

import (
	"flag"
	"fmt"
	"os"
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

	// TODO: Print info about program when using --help flag
	// ...

	_, err := pwd()
	if err != nil {
		fmt.Println("Cannot get current working directory.")
		os.Exit(1)
	}
	// TODO: Get []os.Filenfio slice in current working directory
	// ...
	// TODO: Make slice of checkers on the basis of parsed flags
	// ..
	// TODO: Apply checkers to []os.Fileinfo slice
	// ...
	// TODO: Make a []order slice from []os.Fileinfo slice
	// ...
	// TODO: Execute orders
}
