package Service

import (
	"dzc.com/AppInit"
	"dzc.com/Model"
	"dzc.com/Utils"
	"github.com/jinzhu/gorm"
	"time"
)

type UserService struct {
}

func ShareUserService() *UserService {
	return &UserService{}
}

func getUserDb() *gorm.DB {
	baseDb := AppInit.GetDB()
	return baseDb.Table(Model.TableUsers)
}

func (this *UserService) NameExist(name string) bool {
	var count int
	getUserDb().Where("name = ? ", name).Count(&count)
	return count > 0
}

func (this *UserService) LoginUserInfo(name string, password string) Model.Users {
	var UserModel Model.Users
	getUserDb().Where("name = ? AND password = ? ", name, password).Find(&UserModel)
	if UserModel.Id > 0 {
		getUserDb().Model(&UserModel).Updates(&Model.Users{
			InsertUsers: Model.InsertUsers{
				UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
			},
		})
		this.UpdateToken(&UserModel)
	}
	return UserModel
}

func (this *UserService) Regist(name string, password string) string {
	md5Psw := Utils.Md5(password)
	salt := Utils.RandIntString(6)
	createAt := time.Now().Format("2006-01-02 15:04:05")
	token := Utils.Md5(md5Psw + salt + createAt)
	getUserDb().Create(Model.InsertUsers{
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
	token := Utils.Md5(userModel.Password + userModel.Salt + time.Now().Format("2006-01-02 15:04:05"))
	getUserDb().Model(&userModel).Updates(&Model.Users{
		InsertUsers: Model.InsertUsers{
			Token: token,
		},
	})
	userModel.Token = token
	return userModel
}

func (this *UserService) UserInfoFromToken(token string) (Model.Users, bool) {
	var UserModel Model.Users
	getUserDb().Where("token = ? ", token).Find(&UserModel)
	if UserModel.Id > 0 {
		return UserModel, true
	}
	return UserModel, false
}
