package deferr_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/marconi/deferr"
)

type FakeTodoManager struct {
}

func (ftm *FakeTodoManager) List() []*deferr.Todo {
	return []*deferr.Todo{
		&deferr.Todo{
			Slug: "todo-123",
			Name: "Wash clothes",
		},
	}
}

func (ftm *FakeTodoManager) Push(t *deferr.Todo) error {
	t.Slug = "todo-123"
	return nil
}

func (ftm *FakeTodoManager) Pop() (*deferr.Todo, error) {
	return nil, nil
}

func (ftm *FakeTodoManager) Defer() error {
	return nil
}

func TestWebHandlerSpec(t *testing.T) {
	handler := deferr.NewWebHandler(&FakeTodoManager{})

	Convey("testing web handler", t, func() {
		Convey("should be able to push item", func() {
			body := bytes.NewBufferString("")
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/todos", body)
			handler.Push(w, req, nil)
			So(w.Body.String(), ShouldEqual, `{"message":"Invalid payload."}`)

			body = bytes.NewBufferString(`{"name":"Wash clothes"}`)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", "/todos", body)
			handler.Push(w, req, nil)
			So(w.Body.String(), ShouldEqual, `{"slug":"todo-123","name":"Wash clothes"}`)
		})

		Convey("should be able to list items", func() {
			body := bytes.NewBufferString(`{"name":"Wash clothes"}`)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/todos", body)
			handler.Push(w, req, nil)

			w = httptest.NewRecorder()
			req, _ = http.NewRequest("GET", "/todos", nil)
			handler.List(w, req, nil)
			So(w.Body.String(), ShouldEqual, `[{"slug":"todo-123","name":"Wash clothes"}]`)
		})

		Convey("should be able to pop item", func() {
			// TODO: exercise, implement this yourself
		})

		Convey("should be able to defer item", func() {
			// TODO: exercise, implement this yourself
		})
	})
}
