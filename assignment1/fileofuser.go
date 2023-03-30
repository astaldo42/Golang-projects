package main

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Username          string
	Password          string `json:"-"`
	EncryptedPassword string
}

var Users []User

func (u *User) GetPassword() string {
	return u.Password
}

func againhashing(pass string) string {
	enc := encryptPassword(pass)
	return enc
}

func (u *User) SetPassword(password string) {
	u.Password = password
	u.EncryptedPassword = encryptPassword(password)
}

func Registration(username, password string) {
	user := User{Username: username}
	user.SetPassword(password)
	Users = append(Users, user)
}

func Authorization(username, password string) bool {
	for _, user := range Users {
		if user.Username == username && user.EncryptedPassword == againhashing(password) {
			return true
		}
	}
	return false
}

func encryptPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
