package db

type Row struct {
	Value string
}

type Database interface {
	Next() Row
}

type MockDB struct {}

func (db* MockDB) Next() Row {
	return Row { "Bruce Wayne" }
}

