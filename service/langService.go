package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	str "task/struct"
)

func LangPost(w http.ResponseWriter, r *http.Request) str.ResponseRelation {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	c, errRead := ioutil.ReadAll(r.Body)

	var response str.ResponseRelation

	if errRead != nil {
		fmt.Println("Error")
		response.Status = 500
		response.Desc = "false"
		return response
	}

	var msg str.Lang
	errMarshal := json.Unmarshal(c, &msg)

	if errMarshal != nil {
		response.Status = 500
		response.Desc = "false"
		return response
	}



	response.Status = 200
	response.Desc = msg
	fmt.Println("Language", msg.Language)
	fmt.Println("Appeared", msg.Appeared)
	fmt.Println("Created", msg.Created)
	fmt.Println("Functional", msg.Functional)
	fmt.Println("ObjectOriented", msg.ObjectOriented)
	fmt.Println("InfluencedBy", msg.Relation.InfluencedBy)
	fmt.Println("Influences", msg.Relation.InfluencedBy)
	return response
}
