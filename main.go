package main

import (
	"log"
	"net/http"
	sessioninfo "todo/SessionInfo"
	"todo/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/martini-contrib/method"
)

var LoginInfo sessioninfo.SessionInfo

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	store := cookie.NewStore([]byte("select"))
	engine.Use(sessions.Sessions("mysession", store))

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"UserId": "",
		})
	})
	engine.POST("/login", controller.NewLogin().LoginK)

	engine.GET("/singup", func(c *gin.Context) {
		c.HTML(200, "singup.html", gin.H{})
	})
	engine.POST("/singup", controller.NewLogin().SingUp)
	engine.Use(Override)
	menu := engine.Group("/menu")
	menu.Use(sessionCheck())
	{
		//menu.GET("/top", controller.GetMenu)
		menu.GET("/top", controller.NewTodo().List)
		menu.POST("/top", controller.NewTodo().CreateTodo)

		menu.GET("/top/:id", controller.NewTodo().Get)
		menu.POST("/update/:id", controller.NewTodo().Update)
		menu.POST("/top/:id", controller.NewTodo().Delete)

	}

	engine.POST("/logout", controller.PostLogout)

	engine.Run(":8080")
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		LoginInfo.Name = session.Get("name")

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.Name == nil {
			log.Println(session)
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.Set("name", LoginInfo.Name) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}

var overrideHandler = method.Override()

func Override(c *gin.Context) {
	overrideHandler.ServeHTTP(c.Writer, c.Request)
}
