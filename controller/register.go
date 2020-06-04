package controller

import (
	"goIM/dao"
	"goIM/model"
	"goIM/util"
)

//用户注册
func UserRegister(data *dao.RegisterModel) (cookie string, ok bool, message string) {

	encryptCode := util.AesEncrypt(data.Password)
	data.Password = encryptCode
	cookie, ok, message = model.SaveUserInfo(data)
	return
}

/**
用户登录检查
*/
func UserLoginCheck(data *dao.RegisterModel) (cookie string, ok bool) {

	encryptCode := util.AesEncrypt(data.Password)
	data.Password = encryptCode
	cookie, ok = model.CheckUserInfo(data)
	return

}

/**
用户密码修改
*/

func UserLoginChange(data *dao.RegisterModel) (cookie string, ok bool, message string) {

	encryptCode := util.AesEncrypt(data.Password)
	data.Password = encryptCode
	cookie, ok, message = model.ChangeUserInfo(data)
	return

}
