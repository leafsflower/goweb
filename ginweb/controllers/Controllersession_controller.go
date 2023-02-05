package controllers

import (
	"awesomeProject1/models"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	fmt.Println("HOme", session.Get("secret"))
	loginuser := session.Get("secret")
	fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true

	} else {
		return false
	}

}

func HomeGet(c *gin.Context) {
	islogin := GetSession(c)
	fmt.Println(islogin)
	db := utils.InitDb()
	var posts []models.Post
	db.Find(&posts)
	fmt.Println(posts[0].CreatedAt.String()[:11])

	c.HTML(http.StatusOK, "index.html", gin.H{
		"IsLogin": islogin,
		"posts":   posts,
	})

}

func ExitGet(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println("111", session.Get("secret"))
	session.Delete("secret")

	err := session.Save()
	if err != nil {
		return
	}
	fmt.Println("delete session")

	c.Redirect(http.StatusMovedPermanently, "/index")

}

func NewExitGet(c *gin.Context) {

	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("secret")
	session.Save()
	fmt.Println("delete session...", session.Get("secret"))
	c.Redirect(http.StatusFound, "/")
}
