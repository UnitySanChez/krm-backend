// 存放配置信息
package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"krm-backend/utils/logs"
	"path"
	"runtime"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 //token有效期
	Username   string
	Password   string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数，给结构体赋默认值
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := map[string]interface{}{}
	returnData.Data = data
	return returnData
}
func initLogConfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//文件名和行号
	logrus.SetReportCaller(true)
	//日志格式json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
		// runtime.Frame:帧。可用于获取调用者返回pc值的函数，文件或者是行信息
		//原来是"file":"D:/goland/krm-backend/utils/logs/logs.go:11",
		//加完下面的变成了"file":"logs.go"，因为咱们不需要文件的全路径
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})

}

func init() {
	// 初始化配置
	logs.Debug(nil, "开始加载程序配置")
	// 1. 最先加载环境变量
	viper.AutomaticEnv()

	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8081")
	viper.SetDefault("JWT_SIGN_KEY", "123456")
	viper.SetDefault("JWT_EXPIRE_TIME", 24*60*60) //24小时

	viper.SetDefault("USERNAME", "admin") //配置用户名密码的默认值，生产环境是加密状态
	viper.SetDefault("PASSWORD", "123456")

	//获取端口号配置
	Port = viper.GetString("PORT")
	//获取jwt加密的secret
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	//获取jwt有效期配置
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	//获取用户名和密码配置
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")

	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	initLogConfig(logLevel)

}
