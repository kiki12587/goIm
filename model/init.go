package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	G_db *gorm.DB
	tag  int
)

//初始化数据库连接
func InitModel() {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open("mysql",
		"root:zxcZXC112233@/goim?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		//如果连接失败 尝试重连 让系统等待3秒钟
		time.Sleep(time.Second * 3)
		tag++
		//超过60次重连则失败
		if tag == 60 {
			fmt.Printf("尝试重连失败,%v次\n", tag)
			fmt.Println(err)

			return
		} else {
			fmt.Printf("尝试重连mysql第%v次\n", tag)
			InitModel()
		}
		return
	} else {
		fmt.Println("mysql connection succedssed")
	}
	G_db = db
	return
}
