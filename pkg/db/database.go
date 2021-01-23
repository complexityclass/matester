package db

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Database interface {
	Next() Row
	Credential(login string) (*AuthRow, error)
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

func OpenDB() Database {
	db, err := sql.Open("mysql", "dev:root@tcp(127.0.0.1:3307)/matester_db")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d := DatabaseImpl { Status: false, db: db }
	d.Status = true
	var database Database
	database = &d

	return database
}

func (d *DatabaseImpl) Next() Row {
	return Row { "some" }
}

func (d *DatabaseImpl) ShowTables() string {
	stm, err := d.db.Prepare("SHOW tables")
	if err != nil {
		panic(err)
	}
	defer stm.Close()

	var res string
	err = stm.QueryRow().Scan(&res)
	if err != nil {
		panic(err)
	}

	return res
}

func (d *DatabaseImpl) Credential(login string) (*AuthRow, error) {
	if len(login) < 3 {
		return nil, errors.New("Small login")
	}
	stmtId, err := d.db.Prepare("SELECT user_id FROM users WHERE login = ?")
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
