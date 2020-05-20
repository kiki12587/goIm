package router

import (
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

		c.HTML(http.StatusOK, "/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", gin.H{
			"title": "登录",
		})
	})

	//登录
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

	//首页
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/register.html", gin.H{
			"title": "注册",
		})
	})

	//提交注册信息
	r.POST("/register", func(c *gin.Context) {
		var reg dao.RegisterModel
		reg.Joinip = c.ClientIP()
		_ = c.ShouldBind(&reg)
		cookie, ok, message := controller.UserRegister(&reg)
		if !ok {
			result = util.RetunMsgFunc(1, message, nil)
		} else {
			result = util.RetunMsgFunc(0, message, cookie)
		}
		c.JSON(http.StatusOK, result)
	})

	//忘记密码
	r.GET("/changePassword", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/changePassword.html", gin.H{
			"title": "修改密码",
		})
	})

	//忘记密码
	r.POST("/changePassword", func(c *gin.Context) {
		var userLoginInfo dao.RegisterModel

		_ = c.ShouldBind(&userLoginInfo)
		cookie, ok, message := controller.UserLoginChange(&userLoginInfo)
		if !ok {
			result = util.RetunMsgFunc(1, message, nil)
		} else {
			result = util.RetunMsgFunc(0, message, cookie)
		}
		c.JSON(http.StatusOK, result)

	})

	//3.运行
	r.Run(":" + config.GetEnv().ServerPort)
	return
}
