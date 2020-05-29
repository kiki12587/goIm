package model

import (
	"fmt"
	"github.com/rs/xid"
	"goIM/config"
	"goIM/dao"
	"goIM/util"
	"time"
)

/**
保存用户注册信息
*/
func SaveUserInfo(user *dao.RegisterModel) (cookie string, ok bool, message string) {
	var userStatus dao.RegisterModel
	userStatus = dao.RegisterModel{}
	G_db.Table("goim_users").Select("*").Where("username = ?", user.Username).Limit(1).Scan(&userStatus)
	if len(userStatus.Openid) != 0 {
		cookie = ""
		ok = false
		message = "该邮箱已被注册过"
	} else {
		user.Groupid = 1
		user.Joindate = time.Now().Unix()
		user.Register_type = 1
		user.Remark = time.Unix(1389058332, 0).Format("2006-01-02 15:04:05") + ":成功注册"
		user.Openid = xid.New().String()
		user.Status = 1
		_ = G_db.Table("goim_users").Create(&user)
		cookie = user.Openid
		ok = true
		message = "注册成功"
	}

	return
}

/**
校验用户信息
*/

func CheckUserInfo(user *dao.RegisterModel) (cookie string, ok bool) {
	timeUnix := time.Now().Unix()
	var userStatus dao.RegisterModel
	userStatus = dao.RegisterModel{}
	G_db.Table("goim_users").Select("*").Where("username = ? and password = ?", user.Username, user.Password).Limit(1).Scan(&userStatus)
	if timeUnix > userStatus.Expire {
		_ = updateUserSignAndExpire(userStatus.Openid)

	}

	cookie = userStatus.Openid
	ok = true

	return
}

/**
修改用户密码
*/
func ChangeUserInfo(user *dao.RegisterModel) (cookie string, ok bool, message string) {
	var userStatus dao.RegisterModel
	userStatus = dao.RegisterModel{}

	G_db.Table("goim_users").Select("*").Where("username = ?", user.Username).Limit(1).Scan(&userStatus)
	if len(userStatus.Openid) == 0 {
		cookie = ""
		ok = false
		message = "用户不存在"
	} else {

		err := G_db.Table("goim_users").Where("username = ? ", user.Username).Update("password", user.Password).Error
		if err != nil {
			cookie = ""
			ok = false
			message = "密码修改失败"
		}
		cookie = userStatus.Openid
		ok = true
		message = "密码修改成功,跳转到登录"
	}
	return
}

/**
更新用户IM usersign 签名  expire 过期更新
*/

func updateUserSignAndExpire(openid string) (err error) {
	var userTengXunImInfo dao.TengXunYunIm
	userTengXunImInfo.Sdkappid = 1400362640
	userTengXunImInfo.Key = "34e1bb60f73163964f7a857101e718348faa10de88f3d6ce91e3119a4196b171"
	userTengXunImInfo.Identifier = openid
	userTengXunImInfo.Expire = 2592000
	sign, err := util.GetTengXunImSign(&userTengXunImInfo)
	if err != nil {
		//1.记录日志到mongoDB
		fmt.Println("腾讯签名错误:", err)
	} else { //更新用户信息
		G_db.Table("goim_users").Where("openid = ?", openid).Update(map[string]interface{}{"usersig": sign, "expire": int(time.Now().Unix()) + config.GetEnv().Expire})
		err = nil
	}
	return
}
