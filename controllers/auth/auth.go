package auth

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登陆的逻辑
func Login(r *gin.Context) {
	returnData := config.NewReturnData() //构造函数，里面的东西用来替换gin.H{"message": err.Error(), "status": "401"}
	//1.获取前端传来的用户名和密码
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		//r.JSON(200, gin.H{"message": err.Error(), "status": "401"})
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	logs.Debug(map[string]interface{}{"username": userInfo.Username, "password": userInfo.Password}, "开始验证登陆信息")

	//2.验证登陆信息
	//数据库或者环境变量
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		ss, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"username": userInfo.Username, "错误信息": err.Error()}, "用户名密码正确但生成token失败")
			returnData.Status = 401
			returnData.Message = "生成token失败"
			r.JSON(200, returnData)
			return
		}
		//token正常生成，返回给前端
		logs.Info(map[string]interface{}{"username": userInfo.Username, "token": ss}, "用户名密码正确，token生成成功")
		date := map[string]interface{}{}
		date["token"] = ss
		returnData.Status = 200
		returnData.Message = "登陆成功"
		returnData.Data = date
		r.JSON(200, returnData)
		return
	} else {
		//用户名密码错误
		returnData.Status = 401
		returnData.Message = "用户名密码错误"
		logs.Info(map[string]interface{}{"username": userInfo.Username, "password": userInfo.Password, "conusername": config.Username, "conpassword": config.Password}, "用户名密码错误")
		r.JSON(401, returnData)
		return
	}

}

// 登出
func Logout(r *gin.Context) {
	//退出
	//实现退出逻辑，如把token存在redis里面
	r.JSON(200, gin.H{"message": "退出成功", "status": "200"})
	logs.Debug(nil, "用户已经退出")

}
