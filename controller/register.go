package controller

import (
	"fmt"
	"goIM/dao"
	"goIM/model"
	"goIM/util"
)

func UserRegister(data *dao.RegisterModel) (ok bool) {

	encryptCode := util.AesEncrypt(data.Password)
	data.Password = encryptCode
	fmt.Printf("%#v\n", data)
	ok = model.SaveUserInfo(data)
	if !ok {
		ok = false
	}
	return

}
