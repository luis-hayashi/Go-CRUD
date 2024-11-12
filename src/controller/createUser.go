package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luis-hayashi/Go-CRUD/src/configuration/rest_err"
	"github.com/luis-hayashi/Go-CRUD/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrects fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}