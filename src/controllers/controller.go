package controllers

import (
	"assignment2/src/config"

	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()
