/*
@Time : 2020/5/28 21:41
@Author : HP
@File : im
@Software: GoLand
*/
package dao

type TengXunYunIm struct {
	Sdkappid   int    `json:"sdkappid"`
	Key        string `json:"key"`
	Identifier string `json:"identifier"`
	Expire     int    `json:"expire"`
}
