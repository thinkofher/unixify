package main

import (
	"os"
	"path/filepath"
	"strings"
)

// checker represents function for auditing
// os.FileInfo interfaces.
type checker func(os.FileInfo) bool

// applyCheckers returns slice of os.FileInfo containing
// only os.FileInfo interfaces, fulfilling given checkers.
func applyCheckers(files []os.FileInfo, checkers ...checker) []os.FileInfo {
	var ans []os.FileInfo

	for _, file := range files {
		for _, c := range checkers {
			if c(file) {
				ans = append(ans, file)
				break
			}
		}
	}
	return ans
}

// isDir returns true if given file is directory.
func isDir(file os.FileInfo) bool { return file.IsDir() }

// isHidden returns true if given file is hidden.
func isHidden(file os.FileInfo) bool {
	return strings.HasPrefix(file.Name(), ".")
}

// isHiddenNonDir returns true if given file is hidden
// and is not directory.
func isHiddenNonDir(file os.FileInfo) bool {
	return !isDir(file) && isHidden(file)
}

// isDirNonHidden returns true if given file is directory
// and is not hidden.
func isDirNonHidden(file os.FileInfo) bool {
	return !isHidden(file) && isDir(file)
}

// isHiddenOrDir returns true if given file is directory
// or is hidden.
func isHiddenOrDir(file os.FileInfo) bool {
	return isDir(file) || isHidden(file)
}

// isNormal return true if given file is not directory
// and is not hidden.
func isNormal(file os.FileInfo) bool {
	return !isDir(file) && !isHidden(file)
}

// makeOrders returns pointer to slice of orders created from
// given root filepath and slice of os.FileInfos. Function uses
// filepath.Join function to merge filenames and given root filepath,
// which makes makeOrders portable.
func makeOrders(root string, files []os.FileInfo) *[]order {
	var ans []order
	for _, file := range files {
		prev := filepath.Join(root, file.Name())
		unixified := filepath.Join(root, unixify(file.Name()))

		if prev != unixified {
			ans = append(ans,
				*newOrder(prev, unixified))
		}
	}

	return &ans
}
