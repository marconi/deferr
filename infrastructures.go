package deferr

import (
	"fmt"
)

type Storage interface {
	Query() []interface{}
	Push(i interface{}) error
	Pop() (interface{}, error)
	Defer() error
}

type StoreHandler struct {
	Items []interface{}
}

func (sh *StoreHandler) Query() []interface{} {
	return sh.Items
}

func (sh *StoreHandler) Push(i interface{}) error {
	sh.Items = append(sh.Items, i)
	return nil
}

func (sh *StoreHandler) Pop() (interface{}, error) {
	size := len(sh.Items)
	if size == 0 {
		return nil, fmt.Errorf("List is empty.")
	}
	item := sh.Items[0]
	if size > 1 {
		sh.Items = sh.Items[1:]
	} else {
		sh.Items = nil
	}
	return item, nil
}

func (sh *StoreHandler) Defer() error {
	size := len(sh.Items)
	if size <= 1 {
		return nil
	}
	item := sh.Items[0]
	sh.Items = append(sh.Items[1:], item)
	return nil
}
