package models

import (
	_ "github.com/lib/pq"
	"database/sql"
)

type DataStore interface{
	GetUser(username string) (bool, error)
	GetUserPassword(username string)([]byte, error)
}

type DB struct {
	*sql.DB
}

func NewDB(datasource string) (*DB, error) {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
