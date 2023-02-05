package controllers

import (
	"awesomeProject1/models"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func EmailPost(c *gin.Context) {
	mailTo := []string{
		//可以是多个接收人
		"1756863374@qq.com",
	}
	db := utils.InitDb()
	username := c.PostForm("username")

	if username == "" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return

	}

	err1 := db.AutoMigrate(&models.Emailregister{})
	if err1 != nil {
		return
	}
	var er models.Emailregister
	db.First(&er, "username=?", username)
	fmt.Println("ID 是", er.ID)

	if er.ID > 0 {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return

	}

	subject := "Hello World!"
	var str1 string // 邮件主题
	str1 = strconv.Itoa(int(utils.SyCode()))
	fmt.Println(str1)
	body := "hello" + str1 // 邮件正文

	err := utils.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println("Send fail! - ", err)
		return
	}
	fmt.Println("Send successfully!")

	db.Create(&models.Emailregister{
		Model:     gorm.Model{},
		Username:  username,
		Emailcode: str1,
	})

	c.JSON(200, gin.H{"msg": "邮件发送成功"})

}
