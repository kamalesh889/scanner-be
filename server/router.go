package server

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/login", s.Login).Methods("POST")
	r.HandleFunc("/create/store", s.CreateStore).Methods("POST")

	r.HandleFunc("/get/product/{barcode}", s.GetProduct).Methods("GET")

	return r
}
