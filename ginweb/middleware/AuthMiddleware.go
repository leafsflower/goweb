package middleware

import (
	"awesomeProject1/models"
	"awesomeProject1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	获取认证信息
		tokenString := ctx.GetHeader("Authorization")

		fmt.Println("helo", tokenString)

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			ctx.Abort()
			return

		}
		tokenString = tokenString[7:]
		token, claim, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			ctx.Abort()
			return

		}

		userId := claim.UserId
		db := utils.InitDb()
		var user models.User
		db.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足3"})
			ctx.Abort()
			return

		}
		ctx.Set("user", user)
		ctx.Next()

	}
}
