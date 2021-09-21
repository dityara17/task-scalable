package service

import (
	"encoding/json"
	"net/http"
)

func SendResponses(rw http.ResponseWriter, pl interface{}) {
	js, err := json.Marshal(pl)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(js)
	if err != nil {
		panic(err.Error())
	}
}

// StatusNotAllowed is used to set a specific status code
func StatusNotAllowed(code int, allow ...string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(code)
		if len(allow) > 0 {
			w.Write([]byte(`Method not allowed ` ))
		}
	}
}

