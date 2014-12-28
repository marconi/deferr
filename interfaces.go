package deferr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"

	"net/http"
)

////////////////////////////////////////////////////////////////////////////
// Repositories
////////////////////////////////////////////////////////////////////////////

type TodoAdapter interface {
	List() []*Todo
	Save(t *Todo) error
	Pop() (*Todo, error)
	Defer() error
}

type TodoRepo struct {
	store Storage
}

func NewTodoRepo(store Storage) *TodoRepo {
	return &TodoRepo{store: store}
}

func (tr *TodoRepo) List() []*Todo {
	items := tr.store.Query()
	todos := make([]*Todo, len(items))
	for i, item := range items {
		todos[i] = item.(*Todo)
	}
	return todos
}

func (tr *TodoRepo) Save(t *Todo) error {
	return tr.store.Save(t)
}

func (tr *TodoRepo) Pop() (*Todo, error) {
	t, err := tr.store.Pop()
	if err != nil {
		return nil, err
	}
	return t.(*Todo), nil
}

func (tr *TodoRepo) Defer() error {
	return tr.store.Defer()
}

////////////////////////////////////////////////////////////////////////////
// Web handlers
////////////////////////////////////////////////////////////////////////////

type WebHandler struct {
	todoManager TodoInteractor
}

func NewWebHandler(todoManager TodoInteractor) *WebHandler {
	return &WebHandler{todoManager: todoManager}
}

func (wh *WebHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	todos := wh.todoManager.List()
	b, _ := json.Marshal(todos)
	fmt.Fprint(w, string(b))
}

func (wh *WebHandler) Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var t Todo
	var result interface{}
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &t); err != nil {
		result = map[string]string{"message": "Invalid payload."}
	} else {
		wh.todoManager.Add(&t)
		result = t
	}
	b, _ := json.Marshal(result)
	fmt.Fprint(w, string(b))
}

func (wh *WebHandler) Pop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var result interface{}
	result, err := wh.todoManager.Pop()
	if err != nil {
		result = map[string]string{"message": err.Error()}
	}
	b, _ := json.Marshal(result)
	fmt.Fprint(w, string(b))
}

func (wh *WebHandler) Defer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := wh.todoManager.Defer()
	msg := "You have successfully procastinated."
	if err != nil {
		msg = err.Error()
	}
	b, _ := json.Marshal(map[string]string{"message": msg})
	fmt.Fprint(w, string(b))
}
