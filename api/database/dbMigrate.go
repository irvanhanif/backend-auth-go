package database

import "trigonal/backend-auth/api/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}