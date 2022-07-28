package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vishnusomank/sbom-poc/utils/constants"
)

var DB = constants.DB
var POLICYDB = constants.POLICYDB
var SBOMPOLICYDB = constants.SBOMPOLICYDB
var BINARYPATHDB = constants.BINARYPATHDB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "SBOM.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&SBOM{})

	DB = database

	policydatabase, err := gorm.Open("sqlite3", "PolicyDB.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	policydatabase.AutoMigrate(&PolicyDB{})

	POLICYDB = policydatabase

	binarypathdb, err := gorm.Open("sqlite3", "BinaryPathDB.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	binarypathdb.AutoMigrate(&BinaryPathDB{})

	BINARYPATHDB = binarypathdb
}
