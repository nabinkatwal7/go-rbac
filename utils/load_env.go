package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}else{
		fmt.Println("Env loaded")
	}
}