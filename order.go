package main

import (
	"os"
)

type order struct {
	prev      string
	unixified string
}

func (o order) execute() error {
	err := os.Rename(o.prev, o.unixified)
	if err != nil {
		return err
	}
	return nil
}
