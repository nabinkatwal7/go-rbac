package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/go-rbac/model"
	"github.com/nabinkatwal7/go-rbac/utils"
	"gorm.io/gorm"
)

func CreateRoom(c *gin.Context) {
	var input model.Room

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := model.Room{
		Name:     input.Name,
		Location: input.Location,
		UserID:   utils.CurrentUser(c).ID,
	}

	savedRoom, err := room.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedRoom, "message": "Room created successfully"})
}

func GetRooms(c *gin.Context) {
	var Room []model.Room
	err := model.GetRooms(&Room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, Room)
}

func GetRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Room model.Room
	err := model.GetRoom(&Room, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, Room)
}

func UpdateRoom(c *gin.Context) {
	var Room model.Room
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetRoom(&Room, id)
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

	c.BindJSON(&Room)
	err = model.UpdateRoom(&Room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully", "data": Room})
}
