package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var user = User{}

func main() {

	http.HandleFunc("/home", handler)
	http.ListenAndServe("127.0.0:8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		{
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(user)
		}
	case "GET":
		{
			w.Header().Set("Content-Type", "Application/json")
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
