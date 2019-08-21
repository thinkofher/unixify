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

func init() {
	// replace Usage func in flag package with my
	// own version
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, usage())
	}

	flag.BoolVar(&perm, "yes", false, "skip asking to rename at startup")
	flag.BoolVar(&dirs, "dirs", false, "include directories")
	flag.BoolVar(&hidden, "hidden", false, "include hidden files")
	flag.BoolVar(&interactive, "interactive", false, "asks before every rename")
	flag.BoolVar(&verbose, "verbose", false, "print current actions")
	flag.Parse()
}

func main() {
	// get current working directory
	cwd, err := pwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get current working directory.")
		os.Exit(1)
	}

	// get list of files in current working directory
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get list of files in current working directory.")
		os.Exit(1)
	}

	// apply checkers on basis of used flags
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
		fmt.Printf("Files in '%s' are already unixified.\n", cwd)
		os.Exit(0)
	}

	// If user did not provide -yes flag
	if !perm {
		permQuestion := fmt.Sprintf("Do you want to rename %d files?", len(*orders))
		perm, err = askUser(permQuestion)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input.")
			os.Exit(1)
		}

		// If user (again) did not approve
		if !perm {
			os.Exit(0)
		}
	}

	var orderPerm bool
	for _, o := range *orders {
		if interactive {
			orderPerm, err = askUser(o.interactiveQuestion())
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid input.")
				os.Exit(1)
			}
			if !orderPerm {
				continue
			}
		} else if verbose {
			fmt.Println(notify(o.interactiveMsg()))
		}

		err = o.execute()
		if err != nil {
			fmt.Fprintln(os.Stderr, notify("cannot "+o.interactiveMsg()))
		}
	}
}
