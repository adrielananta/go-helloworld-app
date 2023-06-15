package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	//Return string "Hello World" and status OK 200
	fmt.Fprint(w, "Hello World")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	//Request on path "/" with method GET calls function HelloWorld
	router.HandleFunc("/", helloWorld).Methods("GET")
	//Server listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
