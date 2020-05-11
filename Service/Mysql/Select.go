package Mysql

import (
	"dzc.com/Model"
	"dzc.com/Service"
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SelectService struct {
	ctx *gin.Context
}

func ShareSelectService(ctx *gin.Context) *SelectService {
	return &SelectService{
		ctx: ctx,
	}
}

func (this *SelectService) getDb() (*gorm.DB, error) {
	conMod, _ := Utils.GetConnInfo(this.ctx)
	db, err := Service.ShareDbsService().GetDb(conMod.Uuid)
	return db, err
}

func (this *SelectService) ShowDatabases(dbs Model.SelectDbs) ([]Model.Databases, error) {
	db, err := this.getDb()
	var dbMods []Model.Databases
	if err != nil {
		return nil, err
	}
	db.Raw("SHOW DATABASES").Find(&dbMods)
	return dbMods, err
}

func (this *SelectService) ShowTables(dbName string) ([]Model.TablesInfo, error) {
	db, err := this.getDb()
	var tablesMods []Model.TablesInfo
	if err != nil {
		return nil, err
	}
	db.Raw("select * from information_schema.tables where table_schema=? and table_type='base table'", dbName).Find(&tablesMods)
	return tablesMods, err
}

func (this *SelectService) ShowTableInfo(dbName string, tableName string) ([]Model.TableInfo, error) {
	db, err := this.getDb()
	var tableMods []Model.TableInfo
	if err != nil {
		return nil, err
	}
	db.Raw("select * from information_schema.COLUMNS where TABLE_NAME = ? and TABLE_SCHEMA=? ", tableName, dbName).Find(&tableMods)
	return tableMods, err
}
