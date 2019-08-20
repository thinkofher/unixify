package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// flags variables
var dirs bool
var hidden bool
var interactive bool
var verbose bool
var perm bool

var checkers []checker

func main() {
	flag.BoolVar(&perm, "yes", false, "skip asking to rename at startup")
	flag.BoolVar(&dirs, "dirs", false, "include directories")
	flag.BoolVar(&hidden, "hidden", false, "include hidden files")
	flag.BoolVar(&interactive, "interactive", false, "asks before every rename")
	flag.BoolVar(&verbose, "verbose", false, "print current actions")
	flag.Parse()

	// TODO: Print info about program when using --help flag
	// ...

	cwd, err := pwd()
	if err != nil {
		fmt.Println("Cannot get current working directory.")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Println("Cannot get list of files in current working directory.")
		os.Exit(1)
	}

	checkers = append(checkers, isNormal)
	if dirs && hidden {
		checkers = append(checkers, isHiddenOrDir)
	} else {
		if dirs {
			checkers = append(checkers, isDirNonHidden)
		} else if hidden {
			checkers = append(checkers, isHiddenNonDir)
		}
	}

	files = applyCheckers(files, checkers...)

	orders := makeOrders(cwd, files)
	if len(*orders) < 1 {
		fmt.Println("Files are already unixified.")
		os.Exit(0)
	}

	// TODO: Execute orders
	// If user did not provide -yes flag
	if !perm {
		permQuestion := fmt.Sprintf("Do you want to rename %d files?", len(*orders))
		perm, err = askUser(permQuestion)
		if err != nil {
			fmt.Println("Invalid input.")
			os.Exit(1)
		}

		// If user (again) did not approve
		if !perm {
			os.Exit(0)
		}
	}
}
