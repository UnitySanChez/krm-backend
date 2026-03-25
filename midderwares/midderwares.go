// 中间件层
package midderwares

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"
)

func JWTAuth(r *gin.Context) {
	//1.除了登陆接口，其他接口都需要验证token
	requestUrl := r.FullPath()
	logs.Debug(map[string]interface{}{"请求路径是：": requestUrl}, "")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Info(map[string]interface{}{"请求路径": requestUrl}, "登陆和退出不需要验证token")
		r.Next()
		return //如果这里没有return,就会继续执行后续的handler函数
	}
	returnData := config.NewReturnData()
	//其他接口需要验证token
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		//说明请求没有携带token,返回401错误
		returnData.Status = 401
		returnData.Message = "请求未携带token，请登陆后尝试"
		r.JSON(200, returnData)
		r.Abort() //立即停止后续所有的handler执行
		return
	}
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "token验证失败"
		r.JSON(200, returnData)
		r.Abort() //立即停止后续所有的handler执行
		return
	}
	//验证成功,就把结果放到context上下文对象中
	r.Set("claims", claims)
	r.Next()
}
