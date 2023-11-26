package server

import (
	"barcode/utility"

	"github.com/gorilla/mux"
)

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/login", s.Login).Methods("POST")
	r.HandleFunc("user/{id}/create/store", utility.VerifyToken(s.CreateStore)).Methods("POST")
	r.HandleFunc("user/{id}/get/product/{barcode}", utility.VerifyToken(s.GetProduct)).Methods("GET")
	r.HandleFunc("user/{id}/orders", utility.VerifyToken(s.GetOrders)).Methods("GET")

	return r
}

// Store code will be provided as query param to get the schema
