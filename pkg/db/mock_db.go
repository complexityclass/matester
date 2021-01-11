package db

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Database interface {
	Next() Row
	Credential(login string) AuthRow
	Close()
}

type DatabaseImpl struct {
	Status bool
	db *sql.DB
}

type Row struct {
	Value string
}

type AuthRow struct {
	Pass string
	Salt string
}

type UserRow struct {
	Id int
	name string
}

func OpenDB() DatabaseImpl {
	db, err := sql.Open("mysql", "TODO")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d := DatabaseImpl { Status: false, db: db }
	d.Status = true

	return d
}

func (d *DatabaseImpl) Next() Row {
	return Row { "some" }
}

func (d *DatabaseImpl) Credential(login string) (*AuthRow, error) {
	if len(login) < 3 {
		return nil, errors.New("Small login")
	}
	stmtId, err := d.db.Prepare("SELECT user_id FROM user WHERE login = ?")
	if err != nil {
		return nil, errors.New("No such user")
	}
	defer stmtId.Close()

	stmtOut, err := d.db.Prepare("SELECT pass_hash, pass_salt FROM auth WHERE user_id = ?")
	if err != nil {
		return nil, errors.New("No user credentials")
	}
	defer stmtOut.Close()

	var ar AuthRow
	err = stmtOut.QueryRow(1).Scan(&ar.Pass, &ar.Salt)
	if err != nil {
		panic(err.Error())
	}

	return &ar, nil
}

func (d *DatabaseImpl) Close() {
	d.db.Close();
}
