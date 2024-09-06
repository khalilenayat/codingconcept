package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var (
	users = map[string]User{}
)

func main() {
	http.HandleFunc("/createuser", adduser)
	fmt.Println("server started")
	log.Fatalf("server start nahi hua,err :%v\n", http.ListenAndServe(":8000", nil))

}

func adduser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users[user.Name] = user

	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(user)

}
