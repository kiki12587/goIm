/*
@Time : 2020/5/28 21:36
@Author : HP
@File : im
@Software: GoLand
*/
package controller

import (
	"goIM/dao"
	"goIM/vender/tencentyun"
)

func GetTengXunImSign(userTengXunImInfo *dao.TengXunYunIm) (sig string, err error) {
	sig, err = tencentyun.GenSig(userTengXunImInfo.Sdkappid, userTengXunImInfo.Key, userTengXunImInfo.Identifier, userTengXunImInfo.Expire)
	return
}
