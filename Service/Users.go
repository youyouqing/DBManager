package Service

import (
	"dzc.com/AppInit"
	"dzc.com/Model"
	"dzc.com/Utils"
	"github.com/jinzhu/gorm"
)

const UsersTableName = "db_users"

type UserService struct {
	dblink *gorm.DB
}

func ShareUserService() *UserService {
	return &UserService{
		AppInit.GetDB().Table(UsersTableName),
	}
}

func (this *UserService) NameExist(name string) bool {
	var count int
	this.dblink.Where("name = ? ", name).Count(&count)
	return count > 0
}

func (this *UserService) LoginUserInfo(name string, password string) Model.Users {
	var UserModel Model.Users
	this.dblink.Where("name = ? AND password = ? ", name, password).Find(&UserModel)
	if UserModel.Id > 0 {
		this.dblink.Model(&UserModel).Updates(&Model.Users{
			InsertUsers: Model.InsertUsers{
				UpdateAt: Utils.NowTimeString(),
			},
		})
		this.UpdateToken(&UserModel)
	}
	return UserModel
}

func (this *UserService) Regist(name string, password string) string {
	md5Psw := Utils.Md5(password)
	salt := Utils.RandIntString(6)
	createAt := Utils.NowTimeString()
	token := Utils.Md5(md5Psw + salt + createAt)
	this.dblink.Create(Model.InsertUsers{
		Name:     name,
		Password: md5Psw,
		Token:    token,
		CreateAt: createAt,
		UpdateAt: createAt,
		Salt:     salt,
	})
	return token
}

func (this *UserService) UpdateToken(userModel *Model.Users) *Model.Users {
	token := Utils.Md5(userModel.Password + userModel.Salt + Utils.NowTimeString())
	this.dblink.Model(&userModel).Updates(&Model.Users{
		InsertUsers: Model.InsertUsers{
			Token: token,
		},
	})
	userModel.Token = token
	return userModel
}

func (this *UserService) UserInfoFromToken(token string) (Model.Users, bool) {
	var UserModel Model.Users
	this.dblink.Where("token = ? ", token).Find(&UserModel)
	if UserModel.Id > 0 {
		return UserModel, true
	}
	return UserModel, false
}
