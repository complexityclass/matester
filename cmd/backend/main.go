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

	// Uncomment to enrich with test data
	//test(app, database)

	var addr = getPortAddr()
	http.HandleFunc("/", app.LoginUser)
	http.HandleFunc("/register", app.SignUpUser)
	http.HandleFunc("/users", app.GetUsersList)
	http.HandleFunc("/friend", app.LinkFriends)
	http.HandleFunc("/friends", app.GetFriendsList)
	fmt.Println("Starting Server at port %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
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
