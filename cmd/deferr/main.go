package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/marconi/deferr"
)

func main() {
	store := &deferr.StoreHandler{}
	repo := deferr.NewTodoRepo(store)
	manager := deferr.NewTodoManager(repo)
	webHandler := deferr.NewWebHandler(manager)

	router := httprouter.New()
	router.GET("/todos", webHandler.List)
	router.POST("/todos", webHandler.Add)
	router.DELETE("/todos", webHandler.Pop)

	host := "localhost:8888"
	log.Println("Listening on " + host)
	log.Fatal(http.ListenAndServe(host, router))
}
