package initializers

import (
	"os"

	"github.com/3AM-Developer/dae/database"
	"github.com/3AM-Developer/dae/models"
)

func SyncDB() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Role{})
	database.DB.AutoMigrate(&models.OneTimeURL{})
	seedData()
}

func seedData() {
	var roles = []models.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []models.User{{Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}

	database.DB.Save(&roles)
	database.DB.Save(&user)
}
