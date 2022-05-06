package main

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/controller"
	"jwtDemo/middleware"
	"jwtDemo/opdb"
	"log"
)

func main() {
	//从配置文件中读取数据库的配置信息并连接数据库
	if dbErr := opdb.InitMySqlConn(); dbErr != nil {
		log.Panicln(dbErr)
	}
	defer opdb.DB.Close()
	//初始化表结构
	opdb.InitModel()
	router := gin.Default()
	v1 := router.Group("apis/v1")
	{ //注册
		v1.POST("/register", controller.RegisterUser)
		//登陆，返回token
		v1.POST("/login", controller.Login)
	}
	sv1 := v1.Group("/auth")
	//检验token
	sv1.Use(middleware.JWTAuth)
	{ //测试
		sv1.GET("/sayHello", controller.SayHello)
		//更新token
		sv1.GET("/refresh", controller.Refresh)
	}
	router.Run(":8080")
}
