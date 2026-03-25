package auth

import (
	"github.com/gin-gonic/gin"
	"krm-backend/controllers/auth"
)

// 实现登陆接口
func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

// 实现退出接口
func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

func RegisterSubRouter(apiGroup *gin.RouterGroup) {
	//配置登陆功能的路由策略
	authGroup := apiGroup.Group("/auth")
	//登陆功能
	login(authGroup)
	//退出功能
	logout(authGroup)
}
