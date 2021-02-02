package main

import (
	"fmt"
	"log"
	"matester/pkg/db"
	"net/http"
	"os"
)

func main() {
	var database = db.OpenDB()
	var app = NewApp(database)
	defer app.Close()
	var mux = http.NewServeMux()

	var addr = getPortAddr()
	mux.Handle("/", commonMiddleware(http.HandlerFunc(app.LoginUser)))
	mux.Handle("/register", commonMiddleware(http.HandlerFunc(app.SignUpUser)))
	mux.Handle("/users", commonMiddleware(http.HandlerFunc(app.GetUsersList)))
	mux.Handle("/user", commonMiddleware(http.HandlerFunc(app.GetUser)))
	mux.Handle("/friend", commonMiddleware(http.HandlerFunc(app.LinkFriends)))
	mux.Handle("/friends", commonMiddleware(http.HandlerFunc(app.GetFriendsList)))
	mux.Handle("/unfriend", commonMiddleware(http.HandlerFunc(app.UnLinkFriends)))
	fmt.Println("Starting Server at port %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getPortAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set")
	}
	fmt.Println("Bind port %s", port)

	return fmt.Sprintf(":%s", port)
}
