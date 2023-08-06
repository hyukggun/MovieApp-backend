package main

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-app/apis"
	"movie-app/db"
	"net/http"
)

// Default url
func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if req.Method == "GET" {
		fmt.Println("Get method called")
	}
	if req.Method == "POST" {
		fmt.Println("Post method called")
		data, _ := io.ReadAll(req.Body)
		fmt.Println("Post body : ", string(data))
	}
}

func handleUsers(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allwoed", http.StatusMethodNotAllowed)
		return
	}
	users, err := db.FindUsers()
	if err != nil {
		http.Error(res, "Cannot find users", http.StatusBadRequest)
	}
	data, err := json.Marshal(users)
	if err != nil {
		http.Error(res, "Encoding failed", http.StatusExpectationFailed)
	}
	res.Write(data)
}

func handleUserByEmail(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allwoed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(req.URL.Path)
	email := req.URL.Path[len("/users/")]
	fmt.Println(string(email))
	user, err := db.FindUserByEmail(string(email))
	if err != nil {
		http.Error(res, "Cannot find users", http.StatusBadRequest)
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(res, "Encoding failed", http.StatusExpectationFailed)
	}
	res.Write(data)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func nowPlaying(res http.ResponseWriter, req *http.Request) {
	data := apis.NowPlayingMovies()
	fmt.Fprintf(res, string(data))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/now_playing", nowPlaying)
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/users/", handleUserByEmail)

	fmt.Println("Server Listening 8000 port")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
