package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/go-rbac/model"
	"github.com/nabinkatwal7/go-rbac/utils"
	"gorm.io/gorm"
)

func CreateBooking(c *gin.Context) {
	var input model.Booking
	var user_id = utils.CurrentUser(c).ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.UserID != 0 {
		user_id = input.RoomID
	}

	booking := model.Booking{
		Status: "NOT PAID",
		RoomID: input.RoomID,
		UserID: user_id,
	}

	savedBooking, err := booking.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedBooking, "message": "Booking created successfully"})
}

func GetBookings(c *gin.Context) {
	var Booking []model.Booking
	err := model.GetBookings(&Booking)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, Booking)
}

func GetUserBookings(c *gin.Context) {
	var Booking []model.Booking
	var user_id = utils.CurrentUser(c).ID
	err := model.GetUserBookings(&Booking, user_id)
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

	c.JSON(http.StatusOK, Booking)
}
