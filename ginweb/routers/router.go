package routers

import (
	"awesomeProject1/controllers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	fmt.Println("时间是：", time.Now().String()[:11])
	r.LoadHTMLGlob("views/*")
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	{

		//注册
		r.GET("/register", controllers.RegisterGet)
		r.POST("/register", controllers.RegisterPost)
		//登录
		r.GET("/login", controllers.LoginGet)
		r.POST("/login", controllers.LoginPost)

		//	主页
		r.GET("/", controllers.HomeGet)
		r.GET("/exit", controllers.NewExitGet)

	}
	r.POST("/email", controllers.EmailPost)

	r.GET("/post", controllers.GoAddPost)
	r.GET("/put/:id", controllers.ArticlePut)
	r.POST("/add", controllers.Add)
	r.GET("/article/:id", controllers.Article)
	r.POST("/article/:id", controllers.ArticleUpdate)
	r.GET("/delete/:id", controllers.ArticleDelete)

	//articleRouter.POST("/update", controllers.Update)

	//articleRouter.POST("/delete", controllers.Delete)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	return r
}
