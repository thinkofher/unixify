package main

import (
	"fmt"
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

func (o order) interactiveMsg() string {
	return fmt.Sprintf("rename %s -> %s", o.prev, o.unixified)
}

func (o order) interactiveQuestion() string {
	return fmt.Sprintf("%s ?", o.interactiveMsg())
}
