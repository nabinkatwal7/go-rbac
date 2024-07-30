package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nabinkatwal7/go-rbac/model"
	"gorm.io/gorm"
)

func Register(context *gin.Context) {
	var input model.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedUser})
}

func Login(context *gin.Context) {
	var input model.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s is required", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model.GetUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"username": input.Username,
		"message":  "Login Successful",
	})
}

func GetUsers(context *gin.Context) {
	var user []model.User

	err := model.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, user)
}

func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	err := model.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var User model.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.BindJSON(&User)
	err = model.UpdateUser(&User)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, User)
}
