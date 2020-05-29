/*
@Time : 2020/5/15 0:20
@Author : HP
@File : return.go
@Software: GoLand
*/
package util

type ReturnMsg struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func RetunMsgFunc(code int, msg interface{}, data interface{}) *ReturnMsg {
	rm := new(ReturnMsg)
	rm.Code = code
	rm.Msg = msg
	rm.Data = data
	return rm
}
