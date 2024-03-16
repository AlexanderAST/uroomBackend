package handlers

import "net/http"

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.sayHello()).Methods("GET")
	s.router.HandleFunc("/create-news", s.handleCreateNews()).Methods("POST")
	s.router.HandleFunc("/delete-news", s.handleNewsDelete()).Methods("DELETE")
	s.router.HandleFunc("/get-newsById", s.handleNewsFindByID()).Methods("POST")
	s.router.HandleFunc("/get-allNews", s.handleGetAllNews()).Methods("GET")
	s.router.HandleFunc("/update-news", s.handleNewsUpdate()).Methods("POST")
	s.router.HandleFunc("/upload-photo", s.handlePhotoUpload()).Methods("PUT")
	s.router.HandleFunc("/delete-photo", s.handlePhotoDelete()).Methods("DELETE")
}
