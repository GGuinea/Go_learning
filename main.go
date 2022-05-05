package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	StandardDeviation int   `json:"stddev"`
	NumList           []int `json:"data"`
}

func getRandom(length int) int {
	return length
}

func serverRequest(w http.ResponseWriter, r *http.Request) {
	var simpleList = []response{
		{StandardDeviation: 43, NumList: []int{1, 2, 3, 4}},
		{StandardDeviation: 23, NumList: []int{2, 2, 2, 2}},
	}

	requestParam, ok := r.URL.Query()["requests"]
	if !ok || len(requestParam[0]) < 1 {
		log.Println("Problem with URL param 'rquests'")
		return
	}

	lengthParam, ok := r.URL.Query()["length"]
	if !ok || len(lengthParam[0]) < 1 {
		log.Println("Problem with URL 'length' param")
		return
	}

	requestNumber := requestParam[0]
	length := lengthParam[0]
	log.Println("is: " + string(requestNumber))
	log.Println("is: " + string(length))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(simpleList)

}

func main() {

	http.HandleFunc("/random/mean", serverRequest)
	http.ListenAndServe(":8080", nil)
}
