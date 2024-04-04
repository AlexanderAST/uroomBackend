package handlers

import "net/http"

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(corsMiddleware)

	s.router.HandleFunc("/hello", s.sayHello()).Methods("OPTIONS", "GET")
	s.router.HandleFunc("/create-news", s.handleCreateNews()).Methods("OPTIONS", "POST")
	s.router.HandleFunc("/delete-news", s.handleNewsDelete()).Methods("OPTIONS", "DELETE")
	s.router.HandleFunc("/get-newsById", s.handleNewsFindByID()).Methods("OPTIONS", "POST")
	s.router.HandleFunc("/get-allNews", s.handleGetAllNews()).Methods("OPTIONS", "GET")
	s.router.HandleFunc("/update-news", s.handleNewsUpdate()).Methods("OPTIONS", "POST")
	s.router.HandleFunc("/upload-photo", s.handlePhotoUpload()).Methods("OPTIONS", "PUT")
	s.router.HandleFunc("/delete-photo", s.handlePhotoDelete()).Methods("OPTIONS", "DELETE")
	s.router.HandleFunc("/get-Photo", s.handleGetPhoto()).Methods("OPTIONS", "GET")
}
