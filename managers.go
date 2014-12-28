package deferr

import (
	"strings"

	"github.com/nu7hatch/gouuid"
)

type TodoInteractor interface {
	List() []*Todo
	Add(t *Todo) error
	Pop() (*Todo, error)
	Defer() error
}

type TodoManager struct {
	todoRepo TodoAdapter
}

func NewTodoManager(todoRepo TodoAdapter) *TodoManager {
	return &TodoManager{todoRepo: todoRepo}
}

func (tm *TodoManager) List() []*Todo {
	return tm.todoRepo.List()
}

func (tm *TodoManager) Add(t *Todo) error {
	t.Slug = tm.getSlug()
	return tm.todoRepo.Save(t)
}

// Removes the first item on the list
func (tm *TodoManager) Pop() (*Todo, error) {
	t, err := tm.todoRepo.Pop()
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Defers the first item down to the bottom
func (tm *TodoManager) Defer() error {
	return tm.todoRepo.Defer()
}

func (tm *TodoManager) getSlug() string {
	uuid, _ := uuid.NewV4()
	return strings.Replace(uuid.String(), "-", "", -1)
}
