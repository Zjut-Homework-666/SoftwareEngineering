package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name  string `gorm:"column:" json:"name"`
	Sex   string `gorm:"column:" json:"sex"`
	Phone int    `gorm:"column:" json:"phone"`
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

type SeatDetailInfo struct {
	Flight string `gorm:"column:" json:"flight"`
	Seat   string `gorm:"column:" json:"seat"`
	Price  int    `gorm:"column:" json:"price"`
	Status string `gorm:"column:" json:"status"`
}

// 结构体的嵌套加载请使用gorm中的preload
type OrderInfo struct {
	OrderTime   string     `gorm:"column:" json:"orderTime"`
	Price       string     `gorm:"column:" json:"price"`
	OrderId     int        `gorm:"column:" json:"orderId"`
	OrderStatus string     `gorm:"column:" json:"orderStatus"`
	UserInfo    UserInfo   `json:"userInfo"`
	FlightSeat  FlightSeat `json:"flightSeat"`
}

type ResponeInfo struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type ReserveInfo struct {
	UserInfo   UserInfo   `json:"userInfo"`
	FlightSeat FlightSeat `json:"flightSeat"`
}

type ReserverReturn struct {
	ResponeInfo ResponeInfo `json:"responeInfo"`
	PayUrl      string      `json:"payUrl"`
	CancelUrl   string      `json:"cancelUrl"`
}

type PayReturn struct {
	ResponeInfo ResponeInfo `json:"responeInfo"`
	FlightInfo  FlightInfo  `json:"flightInfo"`
	OrderInfo   OrderInfo   `json:"orderInfo"`
}

type CheckInfo struct {
	UserInfo UserInfo `json:"userInfo"`
}

type CheckReturn struct {
	Name        string      `json:"name"`
	FlightInfo  FlightInfo  `json:"flightInfo"`
	FlightSeat  FlightSeat  `json:"flightSeat"`
	ResponeInfo ResponeInfo `json:"responeInfo"`
}

type FlightReturn struct {
	Flights     []FlightDetailInfo `json:"flights"`
	ResponeInfo ResponeInfo        `json:"responeInfo"`
}

type SeatReturn struct {
	Seats       []SeatDetailInfo `json:"seats"`
	ResponeInfo ResponeInfo      `json:"responeInfo"`
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
	username := "root"
	password := "12345678"
	host := "121.5.157.39"
	port := "3306"
	database := "AirTicket"
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

var db = InitDB()

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
		reserveInfo := ReserveInfo{}
		c.BindJSON(&reserveInfo)
		// TODO:王瑞沣,难度⭐⭐
		// TODO:二维码部分,严伟志难度⭐⭐
	})

	router.GET("/pay", func(c *gin.Context) {
		// TODO:王瑞沣,难度⭐⭐⭐
	})

	router.GET("/flights", func(c *gin.Context) {
		flightReturn := FlightReturn{}
		flightReturn.ResponeInfo.Msg = "Success"
		flight := c.Query("flight")
		if flight != "" {
			result := db.Table(flight).Find(&flightReturn.Flights, FlightDetailInfo{Flight: flight})
			if result.Error != nil {
				flightReturn.ResponeInfo.Msg = result.Error.Error()
				flightReturn.ResponeInfo.Code = 1
			}
			c.JSON(http.StatusOK, flightReturn)
			return
		}
		limit, _ := strconv.Atoi(c.Query("limit"))
		pages, _ := strconv.Atoi(c.Query("pages"))
		arrPlace := "%" + c.Query("arrPlace") + "%"
		depPlace := "%" + c.Query("depPlace") + "%"
		price, _ := strconv.Atoi(c.Query("price"))
		status, _ := strconv.Atoi(c.Query("status"))
		sort, _ := strconv.Atoi(c.Query("sort"))
		statuss := []string{"已满", "停飞", "未开售", "售票中"}
		for i := 0; i < 4; i++ {
			if status%2 == 0 {
				statuss[i] = ""
			}
			status /= 2
		}
		order := ""
		orderList := []string{"arrTime", "depTime", "price", "seats"}
		order = orderList[sort/2]
		if sort%2 == 1 {
			order += " desc"
		}
		fmt.Println(statuss, order)
		str := "arrPlace LIKE ? AND depPlace LIKE ? AND price < ? AND status IN (?)"
		result := db.Table("Flights").Where(str, arrPlace, depPlace, price, statuss).Order(order).Limit(limit).Offset(limit * (pages - 1)).Find(&flightReturn.Flights)
		if result.Error != nil {
			flightReturn.ResponeInfo.Msg = result.Error.Error()
			flightReturn.ResponeInfo.Code = 1
		}
		c.JSON(http.StatusOK, flightReturn)
	})

	router.POST("/check", func(c *gin.Context) {
		checkInfo := CheckInfo{}
		c.BindJSON(&checkInfo)
		// TODO:杨子博,难度⭐⭐
	})

	router.GET("/seats", func(c *gin.Context) {
		// TODO:杨子博,难度⭐
	})

	router.Run(":8089")
}
