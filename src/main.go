package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cast", castHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func castHandler(res http.ResponseWriter, req *http.Request) {

}
