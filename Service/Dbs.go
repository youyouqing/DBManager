package Service

import (
	"dzc.com/Model"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"sync"
)

type DbsService struct {
	sync.Mutex
}

var dblinks map[string]*gorm.DB

func init() {
	dblinks = map[string]*gorm.DB{}
}

func ShareDbsService() *DbsService {
	return &DbsService{}
}

func (this *DbsService) SetDb(uuid string, connectMod Model.InsertConnects) (*gorm.DB, error) {
	db, err := connectDb(connectMod.Host, connectMod.Port, connectMod.Dbname, connectMod.Username, connectMod.Password)
	if err == nil {
		dblinks[uuid] = db
	}
	fmt.Print(dblinks)
	return db, err
}

func (this *DbsService) GetDb(uuid string) (*gorm.DB, error) {
	this.Lock()
	var db *gorm.DB
	if db, ok := dblinks[uuid]; ok {
		this.Unlock()
		return db, nil
	}
	//  取不到就去数据库取一次  重新设置db
	var mod Model.InsertConnects
	ShareConnectsService().dblink.Where("uuid = ?", uuid).First(&mod)
	if !reflect.DeepEqual(mod, Model.InsertConnects{}) {
		db, err := this.SetDb(uuid, mod)
		if err != nil {
			this.Unlock()
			return db, err
		}
		fmt.Print(db)
		this.Unlock()
		return db, nil
	}
	this.Unlock()
	return db, fmt.Errorf("没有对应的db连接")
}

// root:zhicongdai@tcp(127.0.0.1:3306)/Mysql?charset=utf8mb4&parseTime=True&loc=Local
func connectDb(host string, port string, dbName string, userName string, password string) (db *gorm.DB, dbErr error) {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, host, port, dbName)
	//fmt.Print(mysqlDsn)
	dbCon, dbErr := gorm.Open("mysql",
		mysqlDsn)
	if dbCon != nil {
		//defer dbCon.Close()
		dbCon.SingularTable(true)
		dbCon.DB().SetMaxIdleConns(4)
		dbCon.DB().SetMaxOpenConns(20)
	}
	return dbCon, dbErr
}
