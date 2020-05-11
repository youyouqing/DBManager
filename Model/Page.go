package Model

type PageModel struct {
	Page    int `json:"page" form:"page" binding:"required"`
	PageNum int `json:"pageNum" form:"pageNum" binding:"required"`
}
