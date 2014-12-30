package deferr

type TodoAdapter interface {
	List() []*Todo
	Push(t *Todo) error
	Pop() (*Todo, error)
	Defer() error
}

type Todo struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}
