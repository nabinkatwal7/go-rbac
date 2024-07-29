package db

import (
	"os"

	"github.com/nabinkatwal7/go-rbac/model"
)

func LoadDatabase(){
	Connect()
}

// load seed data into the database
func seedData() {
    var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
    var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
    Database.Save(&roles)
	Database.Save(&user)
}