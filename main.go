package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	response := helloWorldResponse{Message: fmt.Sprintf("Hello, %s!", request.Name)}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
