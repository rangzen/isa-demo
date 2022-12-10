package handler

import (
	"log"
	"net/http"
)

type Todo struct {
	rep TodoRepository
}

type TodoRepository interface {
	Query() (string, error)
}

func NewTodo(repository TodoRepository) *Todo {
	return &Todo{
		rep: repository,
	}
}

func (t *Todo) Query() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		get, err := t.rep.Query()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Fatalf("error writing the Todo Query page: %v", err)
			}
			return
		}
		_, err = w.Write([]byte(get))
		if err != nil {
			log.Fatalf("error writing the Todo Query page: %v", err)
		}
	}
}
