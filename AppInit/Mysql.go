package AppInit

import (
	"dzc.com/Utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func baseDb() {
	var err error
	config := Utils.ShareConfigInstance(false)
	mysqlDsn := config.GetConfigFromKey("mysql_dsn")
	db, err = gorm.Open("mysql",
		mysqlDsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(10)

	fmt.Println("mysql init")
}

func GetDB() *gorm.DB {
	if db == nil {
		baseDb()
	}
	return db
}

