package logs

import (
	"fmt"
	"net/http"
)

// Handler is a middleware that logs the request.
func Handler(f func(w http.ResponseWriter, r *http.Request), log func(s string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		log(fmt.Sprintf("request: %v", r.URL.Path))
	}
}
