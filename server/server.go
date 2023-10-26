package server

import (
	"barcode/model"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router     *mux.Router
	httpClient *http.Client
	db         *model.Database
}

func NewServer(db *model.Database) *server {

	s := &server{}

	s.router = mux.NewRouter()
	s.db = db
	s.httpClient = &http.Client{}

	return s
}
