package model

import (
	"drukdeals/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"` // hashed
}

func (u *User) Create() error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	query := "INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3) RETURNING user_id"
	return db.DB.QueryRow(query, u.Fullname, u.Email, string(hashed)).Scan(&u.UserID)
}

func (u *User) GetByEmail() error {
	query := "SELECT user_id, fullname, email, password FROM users WHERE email = $1"
	return db.DB.QueryRow(query, u.Email).Scan(&u.UserID, &u.Fullname, &u.Email, &u.Password)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
