/*
@Time : 2020/6/15 23:16
@Author : HP
@File : index.go
@Software: GoLand
*/
package controller

import (
	"goIM/dao"
	"goIM/model"
	"goIM/util"
)

func GetUserIdAndSign(cookie string) (userInfo *dao.RegisterModel, ok bool, message string) {
	if len(cookie) == 0 {
		userInfo = nil
		ok = false
		message = "请先登录"
	} else {

		userInfo, ok, message = model.UserInfoByOpenid(util.AesDecrypt(cookie))
	}
	return

}
