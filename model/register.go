package model

import (
	"github.com/rs/xid"
	"goIM/dao"
	"time"
)

/**
保存用户注册信息
*/
func SaveUserInfo(user *dao.RegisterModel) (ok bool) {
	user.Groupid = 1
	user.Joindate = time.Now().Unix()
	user.Register_type = 1
	user.Remark = time.Unix(1389058332, 0).Format("2006-01-02 15:04:05") + ":成功注册"
	user.Openid = xid.New().String()
	user.Status = 1
	_ = G_db.Debug().Table("goim_users").Create(&user)
	return true
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
	}
	cookie = userStatus.Openid
	ok = true
	return

}
