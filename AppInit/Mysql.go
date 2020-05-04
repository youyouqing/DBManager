package AppInit


import (
	"dzc.com/Utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
func init() {
	var err error
	config := Utils.ShareConfigInstance(false)
	mysqlDsn := config.GetConfigFromKey("mysql_dsn")
	db, err = gorm.Open("mysql",
		mysqlDsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

	fmt.Println("mysql init")
}
func  GetDB() *gorm.DB {
	return db
}
