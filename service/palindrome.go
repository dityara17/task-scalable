package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	str "task/struct"
)

func IsPalindrome(w http.ResponseWriter, r *http.Request) str.ResponseParam {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	c, errRead := ioutil.ReadAll(r.Body)

	var response str.ResponseParam

	if errRead != nil {
		fmt.Println("Error")
		response.Status = 500
		response.Desc = false
		return response
	}

	var msg str.RequestParam
	errMarshal := json.Unmarshal(c, &msg)

	if errMarshal != nil {
		response.Status = 500
		response.Desc = false
		return response
	}

	panjang := len(msg.Param)

	for i := 0; i <= panjang-1; i++ {
		if msg.Param[i] != msg.Param[panjang-1-i] {
			response.Status = 200
			response.Desc = false
			return response
		}
	}

	response.Status = 200
	response.Desc = true
	return response
}
