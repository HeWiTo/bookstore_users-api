package services

import (
	"github.com/HeWiTo/bookstore_users-api/domain/users"
	"github.com/HeWiTo/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.Users, *errors.RestErr) {
	result := &users.Users{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.Users) (*users.Users, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
