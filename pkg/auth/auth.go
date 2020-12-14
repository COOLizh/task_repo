// Package auth for user's authentication
package auth

import (
	"github.com/COOLizh/task_repo/pkg/db"
	"github.com/COOLizh/task_repo/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(user *models.User) (err error) {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hash)

	_, err = db.AddUser(user)
	return
}

// LoginUser handles user login
func LoginUser(user *models.User) (res models.LoginResponse, err error) {
	userFromDB, err := db.GetUserByUsername(user.Username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(userFromDB.Password),
		[]byte(user.Password),
	); err != nil {
		return
	}

	token, err := CreateToken(userFromDB)
	if err != nil {
		return
	}

	res.Authorization = token
	return
}
