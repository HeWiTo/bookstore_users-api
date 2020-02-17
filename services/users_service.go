package services

import (
	"github.com/HeWiTo/bookstore_users-api/domain/users"
	"github.com/HeWiTo/bookstore_users-api/utils/errors"
)

func CreateUser(user users.Users) (*users.Users, *errors.RestErr) {
	return &user, nil
}
