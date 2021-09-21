package service

import (
	"net/http"
	str "task/struct"
)

func IsPalindrome(w http.ResponseWriter, r *http.Request) str.ResponseParam {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	text := r.URL.Query().Get("text")
	var response str.ResponseParam
	panjang := len(text)

	for i := 0; i <= panjang-1; i++ {
		if text[i] != text[panjang-1-i] {
			response.Status = 200
			response.Desc = false
			return response
		}
	}

	response.Status = 200
	response.Desc = true
	return response
}
