package main

import (
	"os"
)

type order struct {
	prev      string
	unixified string
}

func newOrder(prev, unixified string) *order {
	return &order{prev: prev, unixified: unixified}
}

func (o order) execute() error {
	err := os.Rename(o.prev, o.unixified)
	if err != nil {
		return err
	}
	return nil
}
