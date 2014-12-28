package deferr

import (
	"fmt"
)

type Storage interface {
	Query() []interface{}
	Save(i interface{}) error
	Pop() (interface{}, error)
	Defer() error
}

type StoreHandler struct {
	items []interface{}
}

func (sh *StoreHandler) Query() []interface{} {
	return sh.items
}

func (sh *StoreHandler) Save(i interface{}) error {
	sh.items = append(sh.items, i)
	return nil
}

func (sh *StoreHandler) Pop() (interface{}, error) {
	size := len(sh.items)
	if size == 0 {
		return nil, fmt.Errorf("List is empty.")
	}
	item := sh.items[0]
	if size > 1 {
		sh.items = sh.items[1:]
	} else {
		sh.items = nil
	}
	return item, nil
}

func (sh *StoreHandler) Defer() error {
	size := len(sh.items)
	if size <= 1 {
		return nil
	}
	item := sh.items[0]
	sh.items = append(sh.items[1:], item)
	return nil
}
