package main

import (
	"encoding/json"
	"net/http"
)

type MyData struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

func jsonApi(w http.ResponseWriter, r *http.Request) {
	data:= []MyData{
		{Name: "Prajwal", Value: 100},
		{Name: "Arjun", Value: 200},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/users", jsonApi)

	http.ListenAndServe(":8080", nil)
}