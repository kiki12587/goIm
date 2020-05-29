/*
@Time : 2020/5/28 21:31
@Author : HP
@File : login
@Software: GoLand
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {

		//这一部分可以替换成从session/cookie中获取，
		if cookie, err := c.Cookie("user"); err == nil {
			if len(cookie) != 0 { //校验是否有key为auth,value为true的cookie
				c.Next()
				return
			} else {

				c.Abort()
				c.HTML(http.StatusOK, "/login.html", gin.H{
					"title": "登录",
				})
				return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
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
