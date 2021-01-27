package main

import (
	"fmt"
	"log"
	"matester/pkg/api"
	"matester/pkg/db"
	"net/http"
	"os"
)

func main() {
	var database = db.OpenDB()
	var app = NewApp(database)
	defer app.Close()
	var mux = http.NewServeMux()

	// Uncomment to enrich with test data
	//test(app, database)

	var addr = getPortAddr()
	mux.Handle("/", commonMiddleware(http.HandlerFunc(app.LoginUser)))
	mux.Handle("/register", commonMiddleware(http.HandlerFunc(app.SignUpUser)))
	mux.Handle("/users", commonMiddleware(http.HandlerFunc(app.GetUsersList)))
	mux.Handle("/user", commonMiddleware(http.HandlerFunc(app.GetUser)))
	mux.Handle("/friend", commonMiddleware(http.HandlerFunc(app.LinkFriends)))
	mux.Handle("/friends", commonMiddleware(http.HandlerFunc(app.GetFriendsList)))
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

func test(app App, database db.Database) {
	var mocks = testData()
	for _, mock := range mocks {
		app.SignUpUserInternal(&mock.user, mock.pass)
	}
	for _, mock := range mocks {
		var user = mock.user
		id, err := database.GetUserId(user.Login)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s id is %d pass is %s \n", user.Login, id, mock.pass)
	}
}

type Mock struct {
	user api.User
	pass string
}

func testData() []Mock {
	return []Mock{
		Mock{
			user: api.User{
				Login: "justfy",
				Token: "",
			},
			pass: "123",
		},
		Mock{
			user: api.User{
				Login: "complexity",
				Token: "",
			},
			pass: "123456",
		},
		Mock{
			user: api.User{
				Login: "lebron",
				Token: "",
			},
			pass: "12345678",
		},
	}
}
