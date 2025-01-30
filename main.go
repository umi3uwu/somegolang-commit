package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = make(map[string]User)
var mu sync.Mutex

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	mu.Lock()
	defer mu.Unlock()
	users[user.ID] = user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	http.Error(w, "Get Only", http.StatusMethodNotAllowed)
	//	return
	//}
	user :=
		User{}
	json.NewDecoder(r.Body).Decode(&user)
	mu.Lock()
	defer mu.Unlock()
	user, ok := users[user.ID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Delete Only", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	delete(users, "1")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/delete", deleteHandler)
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
	}
}
