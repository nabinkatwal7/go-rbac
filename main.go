package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/go-rbac/controller"
	"github.com/nabinkatwal7/go-rbac/db"
	"github.com/nabinkatwal7/go-rbac/model"
	"github.com/nabinkatwal7/go-rbac/utils"
)

func serveApplication() {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func loadDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Role{})

	seedData()
}

func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	db.Database.Create(&roles)
	db.Database.Create(&user)
}

func main() {
	utils.LoadEnv()
	loadDatabase()
	serveApplication()
}

