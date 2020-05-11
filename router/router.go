package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goIM/config"
	"goIM/controller"
	"goIM/dao"
	"html/template"
	"net/http"
)

//初始化静态资源以及路由
func InitWebHtml() (err error) {
	r := gin.Default()

	//防止字符被转义
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("tpl/*")
	r.Static(config.GetEnv().Static, "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		var reg dao.Register
		err := c.ShouldBind(&reg)
		register, err := controller.UserRegister(&reg)
		fmt.Printf("密码:%v\n", register)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": reg.Username,
				"password": reg.Password,
			})
		}
	})

	//3.运行
	r.Run(":" + config.GetEnv().ServerPort)
	return
}
