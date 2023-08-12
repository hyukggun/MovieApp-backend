package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"movie-app/apis"
	"movie-app/db"
	"movie-app/ent"
	"movie-app/ent/user"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func client() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=papayetoo dbname=movieapp password=rhkdgus0322 sslmode=disable")
	return client, err
}

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

func loginHandler(res http.ResponseWriter, req *http.Request) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(req *http.Request) {
		fmt.Printf("Login Handler %v\n", req)
		defer wg.Done()
		if req.Method != http.MethodGet {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		queries := req.URL.Query()
		fmt.Println(queries)
		email, password := queries.Get("email"), queries.Get("password")
		client, err := client()
		if err != nil {
			http.Error(res, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer client.Close()

		ctx := context.Background()
		user, err := client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
		if err != nil {
			http.Error(res, "No matching user", http.StatusInternalServerError)
			return
		}

		compareResult := CheckPasswordHash(password, user.Password)
		fmt.Printf("compareResult : %v", compareResult)
		res.Header().Set("Content-Type", "application/json")
		result := map[string]string{
			"result": "true",
		}
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, "Encoding error", http.StatusInternalServerError)
		}
		res.Write(data)
	}(req)
	wg.Wait()
}

func signUpHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method no allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	fmt.Println(string(body))
	defer req.Body.Close()
	if err != nil {
		http.Error(res, "reading body failed", http.StatusInternalServerError)
		return
	}

	jsonMap := map[string]string{}
	if err := json.Unmarshal(body, &jsonMap); err != nil {
		http.Error(res, "parsing body failed", http.StatusInternalServerError)
		return
	}
	fmt.Println(jsonMap)

	go func() {
		name, email, password := jsonMap["name"], jsonMap["email"], jsonMap["password"]

		encPassword, err := GenerateFromPassword(password)
		if err != nil {
			http.Error(res, "Password encryption failed", http.StatusInternalServerError)
			return
		}

		client, err := client()
		defer client.Close()
		if err != nil {
			http.Error(res, "Connect to the database failed", http.StatusInternalServerError)
			return
		}
		ctx := context.Background()
		user, err := client.User.Create().SetName(name).SetEmail(email).SetPassword(string(encPassword)).Save(ctx)
		if err != nil {
			http.Error(res, "Connect to the database failed", http.StatusInternalServerError)
			return
		}
		jsonMap := map[string]bool{
			"result": user != nil,
		}
		data, err := json.Marshal(jsonMap)
		if err != nil {
			http.Error(res, "Connect to the database failed", http.StatusInternalServerError)
			return
		}
		res.Write(data)
	}()
}

func GenerateFromPassword(password string) ([]byte, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return p, err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/signUp", signUpHandler)
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/now_playing", nowPlaying)
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/users/", handleUserByEmail)

	fmt.Println("Server Listening 8000 port")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
