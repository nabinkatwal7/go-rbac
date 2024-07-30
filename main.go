package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/go-rbac/controller"
	"github.com/nabinkatwal7/go-rbac/db"
	"github.com/nabinkatwal7/go-rbac/model"
	"github.com/nabinkatwal7/go-rbac/utils"
)

func serveApplication() {
	router := gin.Default()

	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:4200"}
	router.Use(cors.New(config))

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(utils.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
	adminRoutes.POST("/room/add", controller.CreateRoom)
	adminRoutes.PUT("/room/:id", controller.UpdateRoom)
	adminRoutes.GET("/room/bookings", controller.GetBookings)

	publicRoutes := router.Group("/api/view")
	publicRoutes.GET("/rooms", controller.GetRooms)
	publicRoutes.GET("/room/:id", controller.GetRoom)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(utils.JWTAuthCustomer())
	protectedRoutes.GET("/rooms/booked", controller.GetUserBookings)
	protectedRoutes.POST("/room/book", controller.CreateBooking)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

func loadDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Role{})
	db.Database.AutoMigrate(&model.Room{})
	db.Database.AutoMigrate(&model.Booking{})

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
