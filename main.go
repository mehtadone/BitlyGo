package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type UserResponse struct {
	Username string    `json:"username"`
	APIKey   uuid.UUID `json:"api_key"`
}

type UsersRepo struct {
	Users []User
}

func (u *User) Create() {
	u.CreatedAt = time.Now().UTC().String()
	u.UpdatedAt = time.Now().UTC().String()
}

func (u *User) Response() UserResponse {
	key := uuid.New()
	return UserResponse{
		Username: u.Username,
		APIKey:   key,
	}
}

func main() {
	port := ":8000"
	// Handlers
	http.HandleFunc("/", RootHandler)
	fmt.Printf("Server is running on %v...\n", port)

	err := http.ListenAndServe(port, logRequest(http.DefaultServeMux))
	CheckError(err)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Documentation"))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s \n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
