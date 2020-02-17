package users

import (
	"net/http"

	"github.com/HeWiTo/bookstore_users-api/utils/errors"

	"github.com/HeWiTo/bookstore_users-api/services"

	"github.com/HeWiTo/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// Handle the error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement me!")
// }
