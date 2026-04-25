package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload);
	if err != nil {
		log.Printf("Failed at Marshal JSON Response %v", payload)
		w.WriteHeader(500)
		w.Write([]byte("Internal Error occured"))
		return
	}
	w.WriteHeader(code);
	w.Write(data);
}