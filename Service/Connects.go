package Service

import (
	"dzc.com/AppInit"
	"dzc.com/Model"
	"dzc.com/Utils"
	"github.com/jinzhu/gorm"
)

const ConnectTableName = "db_connects"

type ConnectsService struct {
	dblink *gorm.DB
}

func ShareConnectsService() *ConnectsService {
	return &ConnectsService{
		AppInit.GetDB().Table(ConnectTableName).Debug(),
	}
}

func (this *ConnectsService) Add(insertMod Model.InsertConnects) (bool, string) {
	if this.NameExist(insertMod.Name) {
		return false, insertMod.Name + "已存在"
	}
	insertMod.CreateAt = Utils.NowTimeString()
	insertMod.UpdateAt = insertMod.CreateAt
	this.dblink.Create(insertMod)
	return true, ""
}

func (this *ConnectsService) Detial(uuid string) Model.Connects {
	var mod Model.Connects
	this.dblink.Where("uuid = ?", uuid).First(&mod)
	return mod
}

func (this *ConnectsService) List(uid int, page int, pageNum int) []Model.Connects {
	var connectMod []Model.Connects
	this.dblink.Where("uid = ?", uid).Offset((page - 1) * pageNum).Limit(pageNum).Find(&connectMod)
	return connectMod
}

func (this *ConnectsService) Total(uid int) int {
	var count int
	this.dblink.Where("uid = ?", uid).Count(&count)
	return count
}

func (this *ConnectsService) NameExist(name string) bool {
	var count int
	this.dblink.Where("name = ? ", name).Count(&count)
	return count > 0
}
