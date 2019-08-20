package main

import (
	"os"
	"path/filepath"
	"strings"
)

type checker func(os.FileInfo) bool

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

func isDir(file os.FileInfo) bool { return file.IsDir() }

func isHidden(file os.FileInfo) bool {
	return strings.HasPrefix(file.Name(), ".")
}

func isHiddenNonDir(file os.FileInfo) bool {
	return !isDir(file) && isHidden(file)
}

func isDirNonHidden(file os.FileInfo) bool {
	return !isHidden(file) && isDir(file)
}

func isHiddenOrDir(file os.FileInfo) bool {
	return isDir(file) || isHidden(file)
}

func isNormal(file os.FileInfo) bool {
	return !isDir(file) && !isHidden(file)
}

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
