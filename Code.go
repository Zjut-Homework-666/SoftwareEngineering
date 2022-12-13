package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var ip = "121.5.157.39"
var host = "127.0.0.1"
var port = ":8089"

type UserInfo struct {
	Name  string `gorm:"column:name" json:"name"`
	Sex   string `gorm:"column:sex" json:"sex"`
	Phone int    `gorm:"column:phone" json:"phone"`
	Mail  string `gorm:"column:mail" json:"mail"`
	Id    string `gorm:"column:id" json:"id"`
}

type FlightSeat struct {
	Flight string `gorm:"column:flight" json:"flight"`
	Seat   int    `gorm:"column:seat" json:"seat"`
}

type FlightInfo struct {
	Flight   string `gorm:"column:flight" json:"flight"`
	ArrPlace string `gorm:"column:arrPlace" json:"arrPlace"`
	DepPlace string `gorm:"column:depPlace" json:"depPlace"`
	ArrTime  string `gorm:"column:arrTime" json:"arrTime"`
	DepTime  string `gorm:"column:depTime" json:"depTime"`
}

type FlightDetailInfo struct {
	Flight      string `gorm:"column:flight" json:"flight"`
	ArrPlace    string `gorm:"column:arrPlace" json:"arrPlace"`
	DepPlace    string `gorm:"column:depPlace" json:"depPlace"`
	ArrTime     string `gorm:"column:arrTime" json:"arrTime"`
	DepTime     string `gorm:"column:depTime" json:"depTime"`
	LowestPrice int    `gorm:"column:lowestPrice" json:"lowestPrice"`
	SeatLeft    int    `gorm:"column:seatLeft" json:"seatLeft"`
	Status      string `gorm:"column:status" json:"status"`
}

type SeatDetailInfo struct {
	Flight string `gorm:"column:flight" json:"flight"`
	Seat   int    `gorm:"column:seat" json:"seat"`
	Price  int    `gorm:"column:price" json:"price"`
	Status string `gorm:"column:status" json:"status"`
}

// 结构体的嵌套加载请使用gorm中的preload
type OrderInfo struct {
	OrderId     int        `gorm:"column:orderId;primary_key" json:"orderId"`
	OrderTime   time.Time  `gorm:"column:orderTime" json:"orderTime"`
	Price       int        `gorm:"column:price" json:"price"`
	OrderStatus string     `gorm:"column:orderstatus" json:"orderStatus"`
	UserInfo    UserInfo   `gorm:"embedded" json:"userInfo"`
	FlightSeat  FlightSeat `gorm:"embedded" json:"flightSeat"`
}

type ResponeInfo struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type ReserveInfo struct {
	UserInfo   UserInfo   `json:"userInfo"`
	FlightSeat FlightSeat `json:"flightSeat"`
}

type ReserveReturn struct {
	ResponeInfo ResponeInfo `json:"responeInfo"`
	PayUrl      string      `json:"payUrl"`
	CancelUrl   string      `json:"cancelUrl"`
	OrderId     int         `json:"orderId"`
}

type ReserveStatusReturn struct {
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
	host := ip
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
func Fuzz(str string) string {
	str = "%" + str + "%"
	return str
}
func Md5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	bytes := c.Sum(nil)
	return hex.EncodeToString(bytes)
}

// 发送邮件给对应的邮箱
func Notice() {
	//TODO:严伟志,难度⭐⭐⭐⭐⭐
}

func main() {
	db := InitDB()

	//初始化一个管道和Map
	orderChan := make(chan int, 1)
	orderCancelMap := make(map[int](context.CancelFunc))
	orderCtxMap := make(map[int](context.Context))
	orderUrlMap := make(map[int]string)
	//一直查看管道内是否有数据，有的话将取消Map中对应ID的协程
	go func() {
		for {
			id := <-orderChan
			if value, ok := orderCancelMap[id]; ok {
				value()
				delete(orderCancelMap, id)
			}
		}
	}()

	router := gin.Default()
	router.Use(CorsMiddleWare()) //避免跨域拦截

	//测试端口
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	router.POST("/reserve", func(c *gin.Context) {
		reserveInfo := ReserveInfo{}
		reserveReturn := ReserveReturn{}
		orderInfo := OrderInfo{}
		flightDetailInfo := FlightDetailInfo{}
		c.BindJSON(&reserveInfo)

		flight := reserveInfo.FlightSeat.Flight
		seat := reserveInfo.FlightSeat.Seat
		seatDetailInfo := SeatDetailInfo{Flight: flight, Seat: seat}
		db.Table("seat").Take(&seatDetailInfo, seatDetailInfo)

		if seatDetailInfo.Status == "空" { //座位状态为空则进行预定
			reserveReturn.ResponeInfo.Code = 0
			reserveReturn.ResponeInfo.Msg = "success"

			//更新座位数据库并获取价格
			db.Table("seat").Where(seatDetailInfo).Update("status", "已预定")

			//生成订单信息
			orderInfo.Price = seatDetailInfo.Price
			orderInfo.FlightSeat = reserveInfo.FlightSeat
			orderInfo.UserInfo = reserveInfo.UserInfo
			orderInfo.OrderTime = time.Now()
			orderInfo.OrderStatus = "未付款"

			//数据库创建订单信息
			db.Table("order").Create(&orderInfo)
			//修改机次状态
			db.Table("flight").Take(&flightDetailInfo, FlightInfo{Flight: flight})
			if flightDetailInfo.SeatLeft == 0 {
				db.Table("flight").Where(flightDetailInfo).Update("status", "已满")
			}

			//部分参数初始化
			orderId := orderInfo.OrderId
			fmt.Println(orderId)
			id := strconv.Itoa(orderId)
			checkCode := Md5(id)
			str1 := "http://" + host + port + "/pay?orderId=" + id
			str2 := "&checkCode=" + checkCode + "&payStatus="

			//生成一个十五分钟后自动取消的协程，并将其取消函数绑定至Map
			d := time.Now().Add(time.Minute * 2)
			ctx, cancel := context.WithDeadline(context.Background(), d)
			orderCancelMap[orderId] = cancel
			orderCtxMap[orderId] = ctx
			orderUrlMap[orderId] = str1 + str2 + "1"

			//返回预定信息
			reserveReturn.ResponeInfo.Msg = "Ticket booked successfully, waiting for payment"
			reserveReturn.OrderId = orderId
			reserveReturn.PayUrl = str1 + str2 + "0"
			reserveReturn.CancelUrl = str1 + str2 + "1"
			c.JSON(http.StatusOK, reserveReturn)
		} else { //座位状态不为空则返回Code和Msg
			reserveReturn.ResponeInfo.Code = 1
			reserveReturn.ResponeInfo.Msg = "Failed. The seat has been reserved"
			c.JSON(http.StatusOK, reserveReturn)
		}

	})

	router.GET("/reserveStatus", func(c *gin.Context) {
		orderId, _ := strconv.Atoi(c.Query("orderId"))

		reserveStatusReturn := ReserveStatusReturn{}
		ctx := orderCtxMap[orderId]
		d := time.Now().Add(time.Minute * 1)
		ctx1, cancel := context.WithDeadline(context.Background(), d)
		defer cancel()
		//等待协程取消
		select {
		case <-ctx1.Done():
			reserveStatusReturn.ResponeInfo.Msg = "Payment not completed"
			reserveStatusReturn.ResponeInfo.Code = 1
			c.JSON(http.StatusOK, reserveStatusReturn)
		case <-ctx.Done():
			cancelUrl := orderUrlMap[orderId]
			//因为超时导致的取消
			if ctx.Err() == context.DeadlineExceeded {
				resp, err := http.Get(cancelUrl)
				if err != nil {
					reserveStatusReturn.ResponeInfo.Msg = err.Error()
					reserveStatusReturn.ResponeInfo.Code = 2
				} else {
					reserveStatusReturn.ResponeInfo.Msg = "Order is cancelled after timeout."
					reserveStatusReturn.ResponeInfo.Code = 1
				}
				resp.Body.Close()
			} else { //手动进行的取消，即付款成功或取消付款
				reserveStatusReturn.ResponeInfo.Msg = "The order ends normally."
				reserveStatusReturn.ResponeInfo.Code = 0
			}
			//返回数据
			delete(orderCtxMap, orderId)
			delete(orderUrlMap, orderId)

			reserveStatusReturn.OrderInfo.OrderId = orderId
			db.Table("order").Take(&reserveStatusReturn.OrderInfo)
			reserveStatusReturn.FlightInfo.Flight = reserveStatusReturn.OrderInfo.FlightSeat.Flight
			db.Table("flight").Take(&reserveStatusReturn.FlightInfo)
			c.JSON(http.StatusOK, reserveStatusReturn)
		}

	})

	router.GET("/pay", func(c *gin.Context) {
		// TODO:王瑞沣,难度⭐⭐,更新机次座位数据库与订单数据库
		orderInfo := OrderInfo{}
		flightDetailInfo := FlightDetailInfo{}
		orderId, _ := strconv.Atoi(c.Query("orderId"))
		checkCode := c.Query("checkCode")
		payStatus, _ := strconv.Atoi(c.Query("payStatus"))
		id := strconv.Itoa(orderId)

		if Md5(id) == checkCode {
			db.Table("order").Where("orderId = ?", id).Find(&orderInfo)
			flight := orderInfo.FlightSeat.Flight

			if payStatus == 0 { //付款成功
				db.Table("order").Where("orderId = ?", id).Update("orderstatus", "已付款")
				db.Table("seat").Where(SeatDetailInfo{Flight: flight, Seat: orderInfo.FlightSeat.Seat}).Update("status", "已付款")
			} else { //取消订单
				db.Table("order").Where("orderId = ?", id).Update("orderstatus", "已取消")
				db.Table("seat").Where(SeatDetailInfo{Flight: flight, Seat: orderInfo.FlightSeat.Seat}).Update("status", "空")
				db.Table("flight").Find(&flightDetailInfo, FlightInfo{Flight: flight})
				if flightDetailInfo.SeatLeft != 0 {
					db.Table("flight").Where(FlightInfo{Flight: flight}).Update("status", "售票中")
				}
			}
			// 向管道内放入完成的ID
			orderChan <- orderId
		}
		if payStatus == 0 {
			c.String(http.StatusOK, "付款成功")
		} else {
			c.String(http.StatusOK, "取消成功")
		}
	})

	router.GET("/flights", func(c *gin.Context) {
		flightReturn := FlightReturn{}
		flightReturn.ResponeInfo.Msg = "Success"
		flight := c.Query("flight")
		if flight != "" {
			result := db.Table("flightInfo").Find(&flightReturn.Flights, FlightDetailInfo{Flight: flight})
			if result.Error != nil {
				flightReturn.ResponeInfo.Msg = result.Error.Error()
				flightReturn.ResponeInfo.Code = 1
			}
			c.JSON(http.StatusOK, flightReturn)
			return
		}

		arrPlace := Fuzz(c.Query("arrPlace"))
		depPlace := Fuzz(c.Query("depPlace"))
		date := Fuzz(c.Query("date"))
		minPrice, _ := strconv.Atoi(c.Query("minPrice"))
		maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))
		str := "arrPlace LIKE ? AND depPlace LIKE ? AND arrTime LIKE ? AND lowestPrice > ? AND lowestPrice < ?"
		result := db.Table("flightInfo").Where(str, arrPlace, depPlace, date, minPrice, maxPrice).Find(&flightReturn.Flights)
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

	router.Run(port)
}