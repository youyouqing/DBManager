package Model

const TableConnects = "db_connects"

type Connects struct {
	Id int `json:"id"`
	InsertConnects
}

type InsertConnects struct {
	Host     string `json:"host" form:"host" binding:"required"`
	Port     string `json:"port" form:"port" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Uid      int    `json:"uid"`
	Dbname   string `json:"dbname" form:"dbname" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"Update_at"`
}
