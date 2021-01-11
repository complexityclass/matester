package main

import (
	"fmt"
	"matester/pkg/db"
	"matester/pkg/store"
	"net/http"
	"log"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	login, pass, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
        w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		fmt.Println("CMP No Auth")
        return
	}

	var database = db.OpenDB()
	authRow, err := database.Credential(login)
	if err == nil {
		panic("db panic!")
	}

	var validator = store.NewAuthValidator()
	if !validator.IsAuthorised(login, authRow) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
        w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		fmt.Printf("%s: %s is wrong", login, pass)
        return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"meesage": "Welcome to matester!"}`))
	fmt.Println("Everything is ok!")
    return
}

func main() {
	http.HandleFunc("/", greeting)
    fmt.Println("Starting Server at port :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
