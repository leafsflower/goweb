package controllers

import (
	"awesomeProject1/models"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Hello",
	})

}

func RegisterPost(c *gin.Context) {
	db := utils.InitDb()
	username := c.PostForm("username")
	password := c.PostForm("password")
	emailcode := c.PostForm("emailcode")
	db.AutoMigrate(&models.Emailregister{})
	var er models.Emailregister
	db.First(&er, "username=?", username)
	if emailcode != er.Emailcode {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return

	}

	db.AutoMigrate(&models.User{})
	var user models.User
	db.First(&user, "username=?", username)
	fmt.Println("ID 是", user.ID)

	if user.ID == 0 {
		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
			return
		}
		db.Create(&models.User{
			Model:    gorm.Model{},
			Username: username,
			Password: string(hasedPassword),
		})
		c.JSON(200, gin.H{"msg": "注册成功"})
		return

	}

	c.JSON(404, gin.H{"msg": "用户已经存在"})

}

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
	})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username", username, "password", password)
	db := utils.InitDb()
	//获取参数

	db.AutoMigrate(&models.User{})
	var user models.User
	db.First(&user, "username=?", username)

	if user.ID == 0 {
		c.JSON(404, gin.H{"msg": "username账号不存在"})
		return
	}
	if user.ID > 0 {

		session := sessions.Default(c)
		session.Set("secret", username)
		err := session.Save()
		fmt.Println("Error!!!", err)

		fmt.Println(session.Get("secret"))

		c.Redirect(http.StatusFound, "/")
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}

}
