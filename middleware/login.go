/*
@Time : 2020/5/28 21:31
@Author : HP
@File : login
@Software: GoLand
*/
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goIM/config"
	"net/http"
	"strings"
)

var (
	build strings.Builder
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {

		//这一部分可以替换成从session/cookie中获取，
		if cookie, err := c.Cookie("user"); err == nil {
			if len(cookie) != 0 && cookie != "null" { //校验是否有key为auth,value为true的cookie
				c.Next()
				return
			} else {
				c.Abort()

				c.Redirect(http.StatusMovedPermanently, GetDomainSubString("/login"))
				return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
			}
		} else {
			c.Abort()

			c.Redirect(http.StatusMovedPermanently, GetDomainSubString("/login"))
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

func LoginValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		if cookie, err := c.Cookie("user"); err == nil {
			if len(cookie) != 0 && cookie != "null" {
				c.Abort()
				//校验是否有key为auth,value为true的cookie
				c.Redirect(http.StatusMovedPermanently, GetDomainSubString("/index/index"))
				return
			}
		} else {
			c.Next()

		}
	}
}

func GetDomainSubString(router string) string {

	s1 := config.GetEnv().Domain
	s2 := router

	sList := []string{s1, s2}
	s := strings.Join(sList, "")

	return s
}

func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		c.Next()
	}
}

func SetUserCookieMiddleware(cookie string) gin.HandlerFunc {
	fmt.Println(cookie, 5645645654)
	return func(c *gin.Context) {
		c.SetCookie("user", cookie, 864000, "/goim", "goim.test", false, false)
		c.Next()
	}
}
