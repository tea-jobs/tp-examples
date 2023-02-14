package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tea-jobs/tp-examples/jwt"
)

func main() {
	http.HandleFunc("/generateToken", generateToken)
	http.HandleFunc("/parseToken", parseToken)

	log.Println("Start http server with port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

func generateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := r.FormValue("data")
		key := r.FormValue("key")
		log.Printf("Generate token with key: %s, data: %s\n", key, data)
		token, err := jwt.GenerateToken(data, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, token)
	}
}

func parseToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		f := struct {
			Token string `json:"token"`
			Key   string `json:"key"`
		}{}

		err := dec.Decode(&f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Parsing \n\ttoken: %s\n\tkey: %s\n", f.Token, f.Key)
		data, err := jwt.ParseToken(f.Token, f.Key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send back data after parsing
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("Data: %s", data))
	}
}
