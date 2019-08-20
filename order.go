package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// order contains old (prev) pathname
// and new (unixified) pathane.
type order struct {
	prev      string
	unixified string
}

// newOrder returns pointer to order with given previous
// and unixified version of path.
func newOrder(prev, unixified string) *order {
	return &order{prev: prev, unixified: unixified}
}

// execute changes filename of file under
// prev filepath to its unixified version.
//
// Returns error if renaming failed.
func (o order) execute() error {
	err := os.Rename(o.prev, o.unixified)
	if err != nil {
		return err
	}
	return nil
}

// interactiveMsg returns appealing message for user
// informing about changing name of file specified in order.
func (o order) interactiveMsg() string {
	return fmt.Sprintf("rename '%s' -> '%s'", o.prev, o.unixified)
}

// interactiveQuestion returns appealing question
// for user about changing name of file specified in order.
func (o order) interactiveQuestion() string {
	return fmt.Sprintf("%s ?", o.interactiveMsg())
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
