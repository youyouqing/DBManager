package Service

import (
	"dzc.com/Model"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)


type DbsService struct {
	sync.Mutex
}

var	dblinks map[string] *gorm.DB

func init() {
	dblinks = map[string]*gorm.DB{}
}

func ShareDbsService() *DbsService {
	return &DbsService{
	}
}

func (this *DbsService)CheckDb(connectMod Model.Connects) (interface{}) {
	var errDbMsg interface{}
	defer func() {
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println("a")
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			fmt.Println("b")
		}
		fmt.Println("e")
	}()
	connectDb(connectMod.Host, connectMod.Port, connectMod.Dbname, connectMod.Username, connectMod.Password)
	return errDbMsg
}

func (this *DbsService)SetDb(connectMod Model.InsertConnects) (*gorm.DB, error) {
	this.Lock()
	db,err := connectDb(connectMod.Host,connectMod.Port,connectMod.Dbname,connectMod.Username,connectMod.Password)
	if err == nil {
		dblinks[connectMod.Name] = db
	}
	this.Unlock()
	return db,err
}

func (this *DbsService)GetDb(name string) (*gorm.DB, error) {
	var db *gorm.DB
	if db,ok := dblinks[name]; ok {
		return db,nil
	}
	return db,fmt.Errorf("没有对应的db连接")
}



// root:zhicongdai@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local
func connectDb(host string, port string, dbName string, userName string, password string) (db *gorm.DB, dbErr error) {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, host, port, dbName)
	//fmt.Print(mysqlDsn)
	dbCon, dbErr := gorm.Open("mysql",
		mysqlDsn)
	defer dbCon.Close()
	if dbErr != nil {
		panic(dbErr)
	}
	dbCon.SingularTable(true)
	dbCon.DB().SetMaxIdleConns(4)
	dbCon.DB().SetMaxOpenConns(20)
	return dbCon, dbErr
}