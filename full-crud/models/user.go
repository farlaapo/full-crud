package models

import (
	"errors"
	"full-crud/db"
	"full-crud/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	raw := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := raw.Scan(&u.ID, &retrievedPassword) // Use & to pass the address

	if err != nil {
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}
