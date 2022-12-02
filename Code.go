package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name  string `gorm:"column:" json:"name"`
	Sex   string `gorm:"column:" json:"sex"`
	Phone string `gorm:"column:" json:"phone"`
	Mail  string `gorm:"column:" json:"mail"`
	Id    string `gorm:"column:" json:"id"`
}

type FlightSeat struct {
	Flight string `gorm:"column:" json:"flight"`
	Seat   string `gorm:"column:" json:"seat"`
}

type FlightInfo struct {
	Flight   string `gorm:"column:" json:"flight"`
	ArrPlace string `gorm:"column:" json:"arrPlace"`
	DepPlace string `gorm:"column:" json:"depPlace"`
	ArrTime  string `gorm:"column:" json:"arrTime"`
	DepTime  string `gorm:"column:" json:"depTime"`
}

// 结构体的嵌套加载请使用gorm中的preload
type OrderInfo struct {
	OrderTime   string `gorm:"column:" json:"orderTime"`
	Price       string `gorm:"column:" json:"price"`
	OrderId     int    `gorm:"column:" json:"orderId"`
	OrderStatus string `gorm:"column:" json:"orderStatus"`
	UserInfo    UserInfo
	FlightSeat  FlightSeat
}

type FlightDetailInfo struct {
	Flight   string `gorm:"column:" json:"flight"`
	ArrPlace string `gorm:"column:" json:"arrPlace"`
	DepPlace string `gorm:"column:" json:"depPlace"`
	ArrTime  string `gorm:"column:" json:"arrTime"`
	DepTime  string `gorm:"column:" json:"depTime"`
	Price    int    `gorm:"column:" json:"price"`
	Seats    int    `gorm:"column:" json:"seats"`
	Status   string `gorm:"column:" json:"status"`
}

type FlightDetailSeat struct {
	Flight string `gorm:"column:" json:"flight"`
	Seat   string `gorm:"column:" json:"seat"`
	Price  int    `gorm:"column:" json:"price"`
	Status string `gorm:"column:" json:"status"`
}

func CorsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func InitDB() *gorm.DB {
	username := ""
	password := ""
	host := "47.96.8.176"
	port := "3306"
	database := "AAA_MIS08"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: args}))
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	return db

}

// var db = InitDB()

//发送邮件给对应的邮箱
func Notice() {
	//TODO:严伟志,难度⭐⭐⭐⭐⭐
}

// 预定完成生成订单后生成线程调用该函数，使其以参数2调用pay
func Cancel() {
	// TODO:王瑞沣,难度⭐⭐
}

func main() {

	router := gin.Default()
	router.Use(CorsMiddleWare()) //避免跨域拦截

	//测试端口
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	router.POST("/reserve", func(c *gin.Context) {
		// TODO:王瑞沣,难度⭐⭐
	})

	router.GET("/pay", func(c *gin.Context) {
		// TODO:王瑞沣,难度⭐⭐⭐
	})

	router.GET("/flights", func(c *gin.Context) {
		//TODO:严伟志,难度⭐⭐⭐⭐⭐
	})

	router.POST("/check", func(c *gin.Context) {
		// TODO:杨子博,难度⭐⭐
	})

	router.GET("/seats", func(c *gin.Context) {
		// TODO:杨子博,难度⭐
	})

	router.Run(":8089")
}
