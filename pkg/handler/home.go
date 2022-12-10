package handler

import (
	"log"
	"net/http"
)

var Home = func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("ISA Demo"))
	if err != nil {
		log.Fatalf("writing the Home page: %v", err)
	}
}
