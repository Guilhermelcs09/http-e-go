package main

import (
	"log"
	"net/http"
)

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("local dos usuarios"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
