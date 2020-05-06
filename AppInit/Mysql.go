package AppInit

import (
	"dzc.com/Utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var dbCon *gorm.DB

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

// root:zhicongdai@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local
func ConnectDb(ip string, port string, dbName string, userName string, password string) (db *gorm.DB, dbErr error) {
	defer dbCon.Close()
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, ip, port, dbName)
	dbCon, dbErr = gorm.Open("mysql",
		mysqlDsn)
	if dbErr != nil {
		log.Fatal(dbErr)
		return nil, dbErr
	}
	dbCon.SingularTable(true)
	dbCon.DB().SetMaxIdleConns(4)
	dbCon.DB().SetMaxOpenConns(20)
	return dbCon, dbErr
}

func GetDB() *gorm.DB {
	if db == nil {
		baseDb()
	}
	return db
}

func GetDBCon() *gorm.DB {
	return dbCon
}
