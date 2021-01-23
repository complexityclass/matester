package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"matester/pkg/api"
	"matester/pkg/db"
	"net/http"
	"time"
)

func main() {
	var database = db.OpenDB()
	var app = NewApp(database)
	defer app.Close()
	//test(app, database)
	//testMarshall()

	http.HandleFunc("/", app.LoginUser)
	http.HandleFunc("/register", app.SignUpUser)
	http.HandleFunc("/users", app.GetUsersList)
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func testMarshall() {
	var fname = "Bruce"
	var lname = "Wayne"
	var layoutISO = "2006-01-02"
	t, _ := time.Parse(layoutISO, "1976-12-31")
	var jobTitle = "Senior Crime Investigator"
	user := api.User{
		Login:     "Batman",
		FirstName: &fname,
		LastName:  &lname,
		BirthDate: sql.NullTime{
			Time:  t,
			Valid: false,
		},
		JobTitle: &jobTitle,
		City:     nil,
		Token:    "",
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
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
