package users

import (
	"bookstore/bookstore_users-api/domain/users"
	"bookstore/bookstore_users-api/services"
	"bookstore/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
