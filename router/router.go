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
		if cookie, err := c.Cookie("user"); err == nil {
			if len(cookie) == 0 { //校验是否有key为auth,value为true的cookie
				c.Next()
				return
			}
		} else {
			c.Abort()
			c.HTML(http.StatusOK, "/login.html", gin.H{
				"title": "登录",
			})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

func LoginValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		if cookie, err := c.Cookie("user"); err == nil {
			if len(cookie) != 0 {
				c.Abort()
				//校验是否有key为auth,value为true的cookie
				c.HTML(http.StatusOK, "index/index.html", gin.H{
					"title": "首页",
				})
			}
		} else {
			c.Next()

		}
	}
}

//初始化静态资源以及路由
func InitWebHtml() (err error) {

	r := gin.Default()

	//防止字符被转义
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("html/*")
	r.Static(config.GetEnv().Static, "./static")

	indexGroup := r.Group("/index")
	indexGroup.Use(Validate()) //使用validate()中间件身份验证
	{

		//首页
		indexGroup.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index/index.html", gin.H{
				"title": "首页",
			})
		})

	}

	r.GET("/", LoginValidate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", gin.H{
			"title": "登录",
		})
	})

	r.GET("/login", LoginValidate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", gin.H{
			"title": "登录",
		})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/register.html", gin.H{
			"title": "注册",
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

	tengxunGroup := r.Group("/tengxun")

	{
		//获取用户签名
		tengxunGroup.GET("/getSig", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index/index.html", gin.H{
				"title": "首页",
			})
		})
	}

	//3.运行
	r.Run(":" + config.GetEnv().ServerPort)
	return
}
