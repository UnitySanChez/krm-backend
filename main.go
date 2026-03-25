//项目的总路口

package main

import (
	"github.com/gin-gonic/gin"
	"krm-backend/config"
	"krm-backend/midderwares"
	"krm-backend/routers"
	"krm-backend/utils/logs"
)

func main() {
	//1. 初始化配置
	//2.配置gin
	r := gin.Default()
	r.Use(midderwares.JWTAuth)
	logs.Info(nil, "项目启动成功")
	/*	//测试生产jwt token是否有用
		ss, _ := jwtutil.GenToken("dddd")
		fmt.Println("测试是否可以生成token:", ss)

		//测试解析token是否有用
		//claims, err := jwtutil.ParseToken("hbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9") //测试错误的token观察是否报错
		//claims, err := jwtutil.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRkZGQiLCJpc3MiOiJ0ZXN0Iiwic3ViIjoiaGhoaCIsImV4cCI6MTc3OTM0NjAzMywibmJmIjoxNzc0MTYyMDMzLCJpYXQiOjE3NzQxNjIwMzN9.P_GNOPFBZyDFrul-P9AIN2B1KPFLjBxzKR-AtW3Sz_4")
		claims, err := jwtutil.ParseToken(ss)

		if err != nil {
			//说明解析失败
			fmt.Println("解析token失败:", err.Error())
		} else {
			fmt.Println(claims)
		}*/

	routers.RegisterRouter(r)
	r.Run(config.Port)

}
