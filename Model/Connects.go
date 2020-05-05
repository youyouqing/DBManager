package Model

const TableConnects = "db_connects"

type Connects struct {
	Id int `json:"id"`
	InsertConnects
}

type InsertConnects struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
