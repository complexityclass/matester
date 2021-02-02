package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"matester/pkg/api"
	"matester/pkg/db"
	"matester/pkg/store"
	"math"
	"net/http"
)

type App struct {
	db         db.Database
	auth       store.AuthValidator
	controller store.UsersController
}

func NewApp(database db.Database) App {
	var auth = store.NewAuthValidator()
	var controller = store.NewUsersController(database)

	fmt.Println("Started matester app")

	return App{db: database, auth: auth, controller: controller}
}

func (app *App) Close() {
	app.db.Close()
	fmt.Println("Finished matester app")
}

func (app *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	auth, err := app.checkAuth(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := app.db.GetUser(auth.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

type SignUpModel struct {
	Pass string   `json:"pass"`
	UserProfile api.UserProfile `json:"profile"`
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

	var userProfile = model.UserProfile
	_, noUser := app.db.GetUserId(userProfile.Login)
	if noUser == nil {
		http.Error(w, "user with login already exists", http.StatusConflict)
		return
	}

	app.SignUpUserInternal(&userProfile, model.Pass)

	w.WriteHeader(http.StatusOK)
}

func (app *App) SignUpUserInternal(profile *api.UserProfile, pass string) {
	var user = &profile.User
	app.auth.AuthoriseUser(user, pass)
	app.db.SaveUser(user)

	var userId, _ = app.db.GetUserId(user.Login)
	for _, hobby := range profile.Hobbies {
		_ = app.db.SaveHobby(hobby, userId)
	}
}

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	_, err := app.checkAuth(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	users, ok1 := r.Form["user"]
	if !ok1 || len(users) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var userLogin = users[0]
	user, err := app.db.GetUserProfile(userLogin)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *App) GetUsersList(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	_, err := app.checkAuth(w, r)
	if err != nil {
		return
	}

	var users = app.controller.List(math.MaxInt32)
	b, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *App) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	user, err := app.checkAuth(w, r)
	if err != nil {
		return
	}
	var userId, getErr = app.db.GetUserId(user.Login)
	if getErr != nil {
		return
	}

	var users = app.controller.Friends(userId, math.MaxInt32)
	b, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *App) LinkFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	user, err := app.checkAuth(w, r)
	if err != nil {
		return
	}

	r.ParseForm()
	friends, ok1 := r.Form["user"]
	if !ok1 || len(friends) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var friendLogin = friends[0]

	userId, userErr := app.db.GetUserId(user.Login)
	friendId, friendErr := app.db.GetUserId(friendLogin)
	if userErr != nil || friendErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	saveErr := app.db.SaveFriend(userId, friendId)
	if saveErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (app *App) UnLinkFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		return
	}
	user, err := app.checkAuth(w, r)
	if err != nil {
		return
	}

	r.ParseForm()
	friends, ok1 := r.Form["user"]
	if !ok1 || len(friends) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var friendLogin = friends[0]

	userId, userErr := app.db.GetUserId(user.Login)
	friendId, friendErr := app.db.GetUserId(friendLogin)
	if userErr != nil || friendErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	saveErr := app.db.DeleteFriend(userId, friendId)
	if saveErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
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
