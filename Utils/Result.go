package Utils

import "net/http"

type Result struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}

func NewResult() *Result {
	return &Result{}
}

func NewResultSuccess200(data interface{}) *Result {
	return &Result{
		Data: data,
		Msg:  "",
		Code: http.StatusOK,
	}
}

func NewResultError500(msg string) *Result {
	return &Result{
		Data: false,
		Msg:  msg,
		Code: http.StatusInternalServerError,
	}
}

func NewResultError404(msg string) *Result {
	return &Result{
		Data: false,
		Msg:  msg,
		Code: http.StatusNotFound,
	}
}

func NewResultError405(msg string) *Result {
	return &Result{
		Data: false,
		Msg:  msg,
		Code: http.StatusMethodNotAllowed,
	}
}

func NewResultError403(msg string) *Result {
	return &Result{
		Data: false,
		Msg:  msg,
		Code: http.StatusForbidden,
	}
}

func (this *Result) SetData(data interface{}) *Result {
	this.Data = data
	return this
}

func (this *Result) SetMsg(msg string) *Result {
	this.Msg = msg
	return this
}

func (this *Result) SetCode(code int) *Result {
	this.Code = code
	return this
}
