package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register struct {
	Username  string `json:"username" form:"username"`
	Password string `json:"password"  form:"password"`

}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLGlob("tpl/*")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"/index.html",nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"/login.html",nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK,"/register.html",nil)
	})

	r.POST("/register", func(c *gin.Context) {
		var reg Register
		err := c.ShouldBind(&reg)
		if err == nil{
			c.JSON(http.StatusOK,gin.H{
				"username":reg.Username,
				"password":reg.Password,
			})
		}
		//title := c.PostForm("title")
		//password := c.PostForm("password")
		//vef := c.PostForm("vef")
		//fmt.Println(title,password,vef)
		//c.HTML(http.StatusOK,"/register.html",nil)
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8080")
}