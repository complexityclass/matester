package main

import (
	"fmt"
	"matester/pkg/db"
	"matester/pkg/store"
)

func main() {
	var database = db.MockDB{}
	var provider = store.NewUsersProvider(&database)
	users, err := provider.List(10)

	if err == nil {
		for _, user := range users {
			fmt.Printf("hello I'm %s", user.Name)
			fmt.Println("")
		}
	}
}
