package controller

import (
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwtDemo/dfst"
	"jwtDemo/middleware"
	"jwtDemo/opdb"
	"log"
	"net/http"
	"time"
)

//注册接口的回调函数
func RegisterUser(c *gin.Context) {
	var userInfo dfst.RegisterInfo
	//将body中的json内容反射到结构体中
	bindErr := c.BindJSON(&userInfo)
	fmt.Println(userInfo)
	//如果错误说明client那边发来的数据不对，直接报错即可
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户注册解析数据失败" + bindErr.Error(),
			"data":   nil,
		})
		return
	}
	//将用户信息插入数据库中
	err := opdb.Register(userInfo)
	//如果出错说明该用户已经注册过了
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "注册失败" + err.Error(),
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "success",
		"data":   nil,
	})
}

//登陆接口的回调函数
func Login(c *gin.Context) {
	var loginReq dfst.LoginReq
	//将body中的json内容反射到结构体中
	bindErr := c.BindJSON(&loginReq)
	fmt.Println(loginReq)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户数据解析失败" + bindErr.Error(),
			"data":   nil,
		})
		return
	}
	//验证账号密码是否正确，即是否在数据库中存在，注册过
	pass, dbErr := opdb.LoginPass(loginReq)
	if !pass {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "账号或密码错误" + dbErr.Error(),
			"data":   nil,
		})
		return
	}
	//创建一个token
	token := generateToken(c, loginReq)
	data := dfst.LoginResult{
		Name:  loginReq.Name,
		Token: token,
	}
	//将对象名和token返回
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
}

//创建一个token
func generateToken(c *gin.Context, loginReq dfst.LoginReq) (token string) {
	// 构造SignKey: 签名和解签名需要使用一个值
	jwt := middleware.NewJWT()
	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		Name: loginReq.Name,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 签名过期时间
			Issuer:    "wxf.top",                // 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err := jwt.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
	log.Println("create token", token)
	return
}

//sayHello接口的回调函数
func SayHello(c *gin.Context) {
	//从上下文中获取claims字段的值，否则就panic
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "hello wxf",
		"data":   claims,
	})
}
func Refresh(c *gin.Context) {
	//从上下文中获取claims字段的值，否则就panic
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	//更新时间
	claims.NotBefore = time.Now().Unix() - 1000
	claims.ExpiresAt = time.Now().Unix() + 3600

	j := middleware.NewJWT()
	//创建出了新的token
	token, err := j.CreateToken(*claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
	log.Println("refresh token", token)
	data := dfst.LoginResult{
		Name:  claims.Name,
		Token: token,
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	return
}
