package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/go-rbac/db"
	"github.com/nabinkatwal7/go-rbac/utils"
)

func serveApplication(){
	router := gin.Default()

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func main(){
	utils.LoadEnv()
	db.LoadDatabase()
	serveApplication()
}