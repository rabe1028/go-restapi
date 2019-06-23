package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", rest_func)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type SimpleMessage struct {
	Message string `json:"message"`
}

func rest_func(w http.ResponseWriter, r *http.Request) {
	msg := SimpleMessage{"Hello World!!"}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(msg)
}
