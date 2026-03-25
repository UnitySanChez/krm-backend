package routers

import (
	"github.com/gin-gonic/gin"
	"krm-backend/routers/auth"
)

// 注册路由方法
func RegisterRouter(r *gin.Engine) {
	//登陆的路由配置
	//1.登陆login
	//2.退出logout
	//3./api/auth/login   /api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)

}
