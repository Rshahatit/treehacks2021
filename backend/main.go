package main

import (
	"fmt"
	"log"
	"net/http"

	
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//Handle API Requests
	myRouter.HandleFunc("/question/{questionId}", recommendation.SimpleResponse)
	myRouter.HandleFunc("/career", recommendation.Career)
	// myRouter.HandleFunc("/<ENDPOINT>", <FUNCTION NAME>).Methods("POST")
	// myRouter.HandleFunc("/<ENDPOINT>", <FUNCTION NAME>).Methods("GET")
	// myRouter.HandleFunc("/<ENDPOINT>", <FUNCTION NAME>).Methods("PUT")

	// Serve static files
	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/build/")))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API v1.0")
	handleRequests()
}
