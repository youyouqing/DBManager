package Model

const TableUsers = "db_users"

type Users struct {
	Id int `json:"id"`
	InsertUsers
}

type InsertUsers struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	Salt     string `json:"sale"`
}

type RegisterBindModel struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
