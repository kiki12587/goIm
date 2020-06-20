/*
@Time : 2020/6/15 23:17
@Author : HP
@File : users.go
@Software: GoLand
*/
package model

import (
	"goIM/dao"
)

func UserInfoByOpenid(Openid string) (userInfo *dao.RegisterModel, ok bool, message string) {
	userInfo = &dao.RegisterModel{}
	G_db.Table("goim_users").Select("*").Where("openid = ?", Openid).Limit(1).Scan(&userInfo)
	if len(userInfo.Openid) != 0 {
		ok = true
		message = ""
	} else {
		ok = false
		message = "该用户不存在,请确认账号密码"
	}
	return
}
