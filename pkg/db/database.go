package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"matester/pkg/api"
	"os"
	"time"
)

type Database interface {
	AuthorisedUser(login string) (*api.User, error)
	GetUserId(name string) (int, error)
	GetUser(name string) (*api.User, error)
	SaveUser(user *api.User)
	SaveFriend(userId int, friendId int) error
	QueryUsersList() []api.User
	QueryFriendsList(userId int) []api.User
	Close()
}

type DatabaseImpl struct {
	Status bool
	db     *sql.DB
}

type Row struct {
	Value string
}

type UserRow struct {
	Id   int
	name string
}

func OpenDB() Database {
	dataSourceName := os.Getenv("MATESTER_DB")
	if dataSourceName == "" {
		fmt.Println("No DB config")
	}

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d := DatabaseImpl{Status: false, db: db}
	d.Status = true
	var database Database
	database = &d

	return database
}

func (d *DatabaseImpl) AuthorisedUser(login string) (*api.User, error) {
	userId, err := d.GetUserId(login)
	if err != nil {
		return nil, err
	}

	stmtOut, err := d.db.Prepare("SELECT token FROM auth WHERE user_id = ?")
	if err != nil {
		return nil, errors.New("No user credentials")
	}
	defer stmtOut.Close()

	var token string
	err = stmtOut.QueryRow(userId).Scan(&token)
	if err != nil {
		panic(err.Error())
	}

	return &api.User{
		Login: login,
		Token: token,
	}, nil
}

func (d *DatabaseImpl) SaveUser(user *api.User) {
	userStmt, err := d.db.Prepare("INSERT INTO users(login, first_name, last_name, birth_date, job_title, city) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Errorf("Can't create insert query for new user")
		return
	}
	defer userStmt.Close()
	_, err = userStmt.Exec(user.Login, user.FirstName, user.LastName, user.BirthDate, user.JobTitle, user.City)
	if err != nil {
		fmt.Printf("Can't insert new user")
	}

	userId, err := d.GetUserId(user.Login)
	if err != nil {
		panic(err)
	}

	authStmt, err := d.db.Prepare("INSERT INTO auth(user_id, token) values (?, ?)")
	if err != nil {
		panic(err)
	}
	defer authStmt.Close()
	_, err = authStmt.Exec(userId, user.Token)
	if err != nil {
		panic(err)
		fmt.Printf("Can't insert new user auth")
	}
}

func (d *DatabaseImpl) SaveFriend(userId int, friendId int) error {
	addStmt, err := d.db.Prepare("INSERT INTO friends(fst, snd) values (?, ?)")
	if err != nil {
		return err
	}
	defer addStmt.Close()
	_, err = addStmt.Exec(userId, friendId)
	if err != nil {
		return err
	}

	return nil
}

func (d *DatabaseImpl) QueryUsersList() []api.User {
	rows, err := d.db.Query("SELECT * FROM users")
	if err != nil {
		return make([]api.User, 0)
	}

	return d.rowsToUsers(rows)
}

func (d *DatabaseImpl) QueryFriendsList(userId int) []api.User {
	rows, err := d.db.Query("SELECT u.* FROM users AS u INNER JOIN ((SELECT DISTINCT fst FROM friends WHERE snd = ?) UNION (SELECT DISTINCT snd FROM friends WHERE fst = ?)) AS fr ON u.user_id = fr.fst", userId, userId)
	if err != nil {
		return make([]api.User, 0)
	}

	return d.rowsToUsers(rows)
}

func (d *DatabaseImpl) Close() {
	d.db.Close()
}

func (d *DatabaseImpl) GetUserId(name string) (int, error) {
	stmtId, err := d.db.Prepare("SELECT user_id FROM users WHERE login = ?")
	if err != nil {
		return -1, err
	}
	defer stmtId.Close()

	var id int
	err = stmtId.QueryRow(name).Scan(&id)
	if err != nil {
		return -1, errors.New("no such user")
	}

	return id, nil
}

func (d *DatabaseImpl) GetUser(name string) (*api.User, error) {
	rows, err := d.db.Query("SELECT * FROM users where login = ?", name)
	if err != nil {
		return nil, err
	}
	var users = d.rowsToUsers(rows)

	return &users[0], nil
}

func (d *DatabaseImpl) rowsToUsers(rows *sql.Rows) []api.User {
	var res []api.User
	for rows.Next() {
		var id string
		var user api.User
		err := rows.Scan(&id, &user.Login, &user.FirstName, &user.LastName, &user.BirthDate, &user.JobTitle, &user.City)
		if err != nil {
			continue
		}
		res = append(res, user)
	}

	return res
}
