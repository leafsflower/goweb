package controllers

import (
	"awesomeProject1/models"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GoAddPost(c *gin.Context) {
	islogin := GetSession(c)
	fmt.Println(islogin)
	c.HTML(200, "post.html", gin.H{

		"IsLogin": islogin,
	})
}

func Add(c *gin.Context) {

	db := utils.InitDb()
	//获取参数
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")
	ctime := time.Now().String()[:11]

	fmt.Println("创建的博客是：", title)
	err := db.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}

	db.Create(&models.Post{
		Model:   gorm.Model{},
		Title:   title,
		Content: content,
		Tag:     tag,
		Ctime:   ctime,
	})
	c.JSON(200, gin.H{
		"msg":   "注册成功",
		"Title": title,
	})
}

func Article(c *gin.Context) {
	islogin := GetSession(c)
	id := c.Param("id")
	fmt.Println("id is ", id)
	db := utils.InitDb()
	var post models.Post
	db.Where("id=?", id).First(&post)
	fmt.Println(post.CreatedAt.String()[:11])
	createtime := post.CreatedAt.String()[:11]
	c.HTML(200, "article.html", gin.H{
		"id":        id,
		"post":      post,
		"cretetime": createtime,
		"IsLogin":   islogin,
	})

}

func ArticlePut(c *gin.Context) {
	islogin := GetSession(c)

	id := c.Param("id")
	fmt.Println("要修改的id is ", id)
	db := utils.InitDb()
	var post models.Post
	db.Where("id=?", id).First(&post)
	fmt.Println(post.CreatedAt)
	c.HTML(200, "put.html", gin.H{
		"id":      id,
		"post":    post,
		"IsLogin": islogin,
	})

}

func ArticleUpdate(c *gin.Context) {

	id := c.Param("id")
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")
	fmt.Println(id)
	db := utils.InitDb()
	var post models.Post
	db.Model(&post).Where("id = ?", id).Updates(models.Post{
		Model:   gorm.Model{},
		Title:   title,
		Content: content,
		Tag:     tag,
	})
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})

}

func ArticleDelete(c *gin.Context) {
	id := c.Param("id")
	db := utils.InitDb()

	var post models.Post
	db.First(&post, "id = ?", id)
	fmt.Println(post)
	db.Where("id = ?", id).Delete(&post)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})

}
