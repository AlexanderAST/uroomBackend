package handlers

import "net/http"

func (s *server) sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]interface{}{"status": "hello world"})

	}
}
