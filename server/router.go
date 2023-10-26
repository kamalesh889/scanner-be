package server

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/login", s.Login).Methods("POST")

	return r
}
