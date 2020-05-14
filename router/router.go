package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goIM/config"
	"goIM/controller"
	"goIM/dao"
	"goIM/util"
	"html/template"
	"net/http"
)

var result *util.ReturnMsg

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		username := c.Query("username")
		password := c.Query("password")

		if username == "ft" && password == "123" {
			//c.JSON(http.StatusOK, gin.H{"message": "身份验证成功"})
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

//初始化静态资源以及路由
func InitWebHtml() (err error) {
	r := gin.Default()
	//r.Use(Validate()) //使用validate()中间件身份验证
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

	r.POST("/login", func(c *gin.Context) {
		var userLoginInfo dao.RegisterModel

		_ = c.ShouldBind(&userLoginInfo)
		cookie, ok := controller.UserLoginCheck(&userLoginInfo)
		if !ok {
			result = util.RetunMsgFunc(1, "账号或密码错误", nil)
		} else {
			result = util.RetunMsgFunc(0, "登录成功", cookie)
		}
		c.JSON(http.StatusOK, result)

	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/index.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		var reg dao.RegisterModel
		reg.Joinip = c.ClientIP()
		_ = c.ShouldBind(&reg)
		fmt.Printf("%#v\n", reg)
		ok := controller.UserRegister(&reg)
		if ok {
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
