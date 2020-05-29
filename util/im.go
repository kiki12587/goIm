/*
@Time : 2020/5/29 23:02
@Author : HP
@File : im
@Software: GoLand
*/
package util

import (
	"goIM/dao"
	"goIM/vender/tencentyun"
)

func GetTengXunImSign(userTengXunImInfo *dao.TengXunYunIm) (sig string, err error) {
	sig, err = tencentyun.GenSig(userTengXunImInfo.Sdkappid, userTengXunImInfo.Key, userTengXunImInfo.Identifier, userTengXunImInfo.Expire)
	return
}
