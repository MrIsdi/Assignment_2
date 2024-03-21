package main

import (
	"assignment2/src/config"
	"assignment2/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	routes.Routes()
}
