package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vishnusomank/GoXploitDB/utils"
)

var DB = utils.XPLOITDB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "XploitDB.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&XploitDB{})

	DB = database

}
