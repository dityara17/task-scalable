package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"task/service"
)

func handlers() {
	r := mux.NewRouter()
	fmt.Println("Start Processs...")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := service.IsPalindrome(w, r)
		service.SendResponses(w, s)
	}).Methods("POST")

	r.HandleFunc("/lang", func(w http.ResponseWriter, r *http.Request) {
		s := service.LangPost(w, r)
		service.SendResponses(w, s)
	}).Methods("POST")

	http.Handle("/", r)
}

func main() {

	handlers()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
