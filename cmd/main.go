package main

import (
	"barcode/database"
	"barcode/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Barcode backend started")

	db, err := database.InitializeDB()
	if err != nil {
		log.Panicln("Error in Connecting to Database:", err)
	}

	srv := server.NewServer(db)
	if err != nil {
		log.Panicln("Error in creating server:", err)
	}

	mux := server.Router(srv)
	http.Handle("/", mux)
	http.ListenAndServe(fmt.Sprintf(":%s", "8081"), mux)

}
