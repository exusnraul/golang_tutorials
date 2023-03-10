package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("we are looking at mods")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Heythere moders")
}

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GoLang</h1>"))

}
