package users

import (
	"fmt"

	"github.com/HeWiTo/bookstore_users-api/utils/datenow"

	"github.com/HeWiTo/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*Users)
)

func (user *Users) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = ""
	user.UpdatedAt = ""

	return nil
}

func (user *Users) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	user.CreatedAt = datenow.GetNowString()

	usersDB[user.ID] = user
	return nil
}
