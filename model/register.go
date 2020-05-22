package model

import (
	"github.com/rs/xid"
	"goIM/dao"
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
	var userStatus dao.RegisterModel
	userStatus = dao.RegisterModel{}
	G_db.Table("goim_users").Select("*").Where("username = ? and password = ?", user.Username, user.Password).Limit(1).Scan(&userStatus)
	if len(userStatus.Openid) == 0 {
		cookie = ""
		ok = false
	} else {
		cookie = userStatus.Openid
		ok = true
	}
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
