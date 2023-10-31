package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := s.LoginService(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *server) CreateStore(w http.ResponseWriter, r *http.Request) {

	var req Storedetails

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := s.Createstore(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (s *server) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	code, found := vars["barcode"]
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	storecode := query.Get("storecode")

	resp, err := s.GetProductService(code, storecode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
