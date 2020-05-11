package main

import (
	"goIM/model"
	"goIM/router"
	"time"
)

func main() {

	//初始化路由 和 静态资源
	go router.InitWebHtml()

	//数据库初始化
	go model.InitModel()

	for {
		time.Sleep(500 * time.Second)
	}

}
