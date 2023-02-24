package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Users struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var user Users = Users{
	Id:      1,
	Name:    "nariman",
	Age:     35,
	Address: "yasamal",
}

var users []Users

func main() {

	http.HandleFunc("/", loggingMiddleware(myMiddleware(usersHandler)))
	http.HandleFunc("/user", myMiddleware(userHandler))
	http.HandleFunc("/users", myMiddleware(getUsersHandler))
	http.ListenAndServe(":8082", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from Middleware")
		next(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", r.Method, r.URL, r.RemoteAddr)
		next(w, r)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from Handler")
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(users)
	w.Write(resp)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := r.URL.Query().Get("id")
	fmt.Println("id =>", id)
	for _, item := range users {
		if item.Name == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, _ := ioutil.ReadAll(r.Body)
	var newUser Users
	json.Unmarshal(reqBytes, &newUser)
	users = append(users, newUser)
	fmt.Println(users)
}
