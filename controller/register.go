package controller

import (
	"errors"
	"fmt"
	"goIM/dao"
	"goIM/util"
)

func UserRegister(data *dao.Register) (inter interface{}, err error) {
	if data != nil {
		encryptCode := util.AesEncrypt(data.Password)
		//fmt.Println("解密：", decryptCode)

		err = nil
	} else {
		inter = nil
		err = errors.New("无用户注册信息")
	}
	return

}
