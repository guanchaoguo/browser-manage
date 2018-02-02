package helper

/*
	项目公共函数包
*/

import (
	"browser-manage/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"regexp"
	"time"
	"github.com/streadway/amqp"
)

var GlobalMongdbSession *mgo.Session

func Self_logger(myerr interface{}) {
	logfile := newLogFile()
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(myerr)

}

//获取redis操作对象
func GetRedis() redis.Conn {
	redisConf := config.GetRedisConf()
	c, err := redis.Dial("tcp", redisConf.REDIS_HOST+":"+redisConf.REDIS_PORT, redis.DialPassword(redisConf.REDIS_PASSWORD))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	return c
}

//获取mongodb操作对象
func GetMongodb() *mgo.Session {
	fmt.Println("1111")
	return GlobalMongdbSession.Clone()
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXZY"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func InitMongodb(session *mgo.Session) {
	GlobalMongdbSession = session
	GlobalMongdbSession.SetPoolLimit(10)

}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}
	if end <= 0 {
		return string(rs[start:])
	}
	return string(rs[start:end])
}

//验证是否为IP格式
func CheckIp(ip string) bool {
	layout, err := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", ip)
	if !layout || err != nil {
		return false
	}
	return true
}

//链接RabbitMQ
func GetRabbitMq() (*amqp.Connection,*amqp.Channel)  {
	RabbitMqConfig := config.GetRabbitMQ()
	conn , err  := amqp.Dial(RabbitMqConfig.URL)
	if err != nil {
		panic( "Failed to connect to RabbitMQ")
	}
	ch,err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}
	return conn,ch
}
