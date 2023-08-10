package initializers

import (
	"github.com/3AM-Developer/dae/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
