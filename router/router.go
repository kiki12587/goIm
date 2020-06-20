package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goIM/config"
	"goIM/controller"
	"goIM/dao"
	"goIM/middleware"
	"goIM/util"
	"html/template"
	"net/http"
)

var result *util.ReturnMsg

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
	indexGroup.Use(middleware.Validate()) //使用validate()中间件身份验证
	{

		//首页信息
		indexGroup.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index/index.html", gin.H{
				"title": "首页",
			})
		})

		//获取用户ID 和 签名
		indexGroup.POST("/getUserInfo", func(c *gin.Context) {
			var (
				userInfo *dao.RegisterModel
				ok       bool
				message  string
			)

			if cookie, err := c.Cookie("user"); err == nil {
				userInfo, ok, message = controller.GetUserIdAndSign(cookie)
				if !ok {
					result = util.RetunMsgFunc(1, message, nil)
				} else {
					result = util.RetunMsgFunc(0, message, userInfo)
				}
			} else {
				result = util.RetunMsgFunc(1, "签名过期", nil)
			}

			c.JSON(http.StatusOK, result)

		})
	}

	r.GET("/", middleware.LoginValidate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", gin.H{
			"title": "登录",
		})
	})

	r.GET("/login", middleware.LoginValidate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "/login.html", gin.H{
			"title": "登录",
		})
	})

	//登录
	r.POST("/login", func(c *gin.Context) {
		var userLoginInfo dao.RegisterModel

		_ = c.ShouldBind(&userLoginInfo)
		cookie, ok := controller.UserLoginCheck(&userLoginInfo)
		fmt.Println(cookie, ok)
		if !ok {
			result = util.RetunMsgFunc(1, "账号或密码错误", nil)
		} else {
			c.SetCookie("user", cookie, 864000, "/goim", "goim.test", false, false)
			result = util.RetunMsgFunc(0, "登录成功", cookie)
		}
		c.JSON(http.StatusOK, result)

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
			c.SetCookie("user", cookie, 864000, "/goim", "goim.test", false, false)
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
		tengxunGroup.POST("/getSig", func(c *gin.Context) {
			var userTengXunImInfo dao.TengXunYunIm
			_ = c.ShouldBind(&userTengXunImInfo)
			fmt.Println(&userTengXunImInfo)
			sign, err2 := controller.GetTengXunImSign(&userTengXunImInfo)
			if err2 != nil {
				result = util.RetunMsgFunc(1, err2, nil)
			} else {
				result = util.RetunMsgFunc(0, err2, sign)
			}
			c.JSON(http.StatusOK, result)
		})
	}

	//3.运行
	r.Run(":" + config.GetEnv().ServerPort)
	return
}
