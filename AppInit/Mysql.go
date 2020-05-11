package AppInit

import (
	"dzc.com/Config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func baseDb() {
	var err error
	db, err = gorm.Open("mysql",
		Config.MYSQL_DSN)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(Config.MYSQL_MaxIdleConns)
	db.DB().SetMaxOpenConns(Config.MYSQL_MaxOpenConns)
}

func GetDB() *gorm.DB {
	if db == nil {
		baseDb()
	}
	return db
}
