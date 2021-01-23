package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"matester/pkg/api"
	"matester/pkg/db"
	"matester/pkg/store"
	"net/http"
)

type App struct {
	db   db.Database
	auth store.AuthValidator
}

func NewApp(database db.Database) App {
	var auth = store.NewAuthValidator()
	fmt.Println("Started matester app")

	return App{db: database, auth: auth}
}

func (app *App) Close() {
	app.db.Close()
	fmt.Println("Finished matester app")
}

func (app *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user, err := app.checkAuth(w, r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	greeting := fmt.Sprintf("%s, logged in to matester! \n", user.Login)
	response := fmt.Sprintf(`"{"meesage": %s"}"`, greeting)
	w.Write([]byte(response))
	fmt.Println("Everything is ok!")

	return
}

type SignUpModel struct {
	Pass string   `json:"pass"`
	User api.User `json:"user"`
}

func (app *App) SignUpUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	var model SignUpModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user = model.User
	app.SignUpUserInternal(&user, model.Pass)

	w.WriteHeader(http.StatusOK)
	greeting := fmt.Sprintf("Hey, %s, welcome to matester! \n", user.Login)
	response := fmt.Sprintf(`"{"meesage": %s"}"`, greeting)
	w.Write([]byte(response))
}

func (app *App) SignUpUserInternal(user *api.User, pass string) {
	app.auth.AuthoriseUser(user, pass)
	app.db.SaveUser(user)
}

func (app *App) checkAuth(w http.ResponseWriter, r *http.Request) (*api.User, error) {
	login, pass, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return nil, errors.New("No credentials")
	}

	user, err := app.db.AuthorisedUser(login)

	if err != nil || !app.auth.IsAuthorised(pass, user) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		fmt.Printf("%s: %s is wrong", login, pass)
		return nil, err
	}

	return user, nil
}
