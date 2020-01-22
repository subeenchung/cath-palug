package models

import (
	"errors"
	"database/sql"
)
type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
	Admin bool `json:"admin"`
}

func (db *DB) GetUser(username string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT username FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("Err: user does not exist")
		}else{
			return false, err
		}
	}
		
	return true, nil
}

func (db *DB) GetUserPassword(username string)([]byte, error){
	var pwd []byte
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return []byte(""), errors.New("Err: user does not exist")
		} else {
			return []byte(""), err
		}
	}
	return pwd, nil
}
