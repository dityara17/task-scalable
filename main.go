package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"task/service"
)

/**
1. GET /language —> return response body berupa JSON data pada struct awal
dengan proses sebagai berikut: data di-input-kan menggunakan struct (hard coded)
kemudian convert sebagai JSON
2. method lain selain kedua diatas —> return error method not allowed dengan response
body “Method not allowed”
*/

func handlers() {
	r := mux.NewRouter()
	fmt.Println("Start Processs...")

	r.HandleFunc("/palindrome", func(w http.ResponseWriter, r *http.Request) {
		s := service.IsPalindrome(w, r)
		service.SendResponses(w, s)
	}).Methods("GET")

	r.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
		s := service.LangPost(w, r)
		service.SendResponses(w, s)
	}).Methods("POST")

	r.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
		s := service.GetPost(w, r)
		service.SendResponses(w, s)
	}).Methods("GET")



	r.HandleFunc("/language", service.StatusNotAllowed(405, "GET", "POST")).Methods("PUT", "PATCH", "DELETE")
	http.Handle("/", r)
}

func main() {

	handlers()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
