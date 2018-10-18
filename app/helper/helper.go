package helper

/*
	项目公共函数包
*/

import (
	"browser-manage/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var GlobalMongdbSession *mgo.Session

const (
	UNSERNAME       string = `^[a-zA-Z0-9]{6,50}+$` //匹配用户数字 字母下划线 汉字 大于3位
	PWD             string = `[\S]{6,40}+$`                     //匹配任意字符 6到40位
	maxURLRuneCount int    = 2083                                 //url 最大长度
	minURLRuneCount int    = 3                                    // 最小长度
	URLIP           string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLUsername     string = `(\S+(:\S*)?@)`
	URLSubdomain    string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	IP              string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLPort         string = `(:(\d{1,5}))`
	URL             string = `^` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?$`
)

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
	return net.ParseIP(ip) != nil
}

//验证是否为用户名格式
func CheckUserName(userName string) bool {
	layout, err := regexp.MatchString(UNSERNAME, userName)
	if !layout || err != nil {
		return false
	}
	return true
}

//验证是否为密码格式
func CheckPwd(pwd string) bool {
	layout, err := regexp.MatchString(PWD, pwd)
	if !layout || err != nil {
		return false
	}
	return true
}

//验证是否为域名格式的
func CheckDomain(domain string) bool {
	// 验证域名格式
	if domain == "" || utf8.RuneCountInString(domain) >= maxURLRuneCount || len(domain) <= minURLRuneCount || strings.HasPrefix(domain, ".") {
		return false
	}
	return regexp.MustCompile(URL).MatchString(domain)
}

//验证是否为合法端口格式
func CheckPort(port string) bool {
	if i, err := strconv.Atoi(port); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

//链接RabbitMQ
func GetRabbitMq() (*amqp.Connection, *amqp.Channel) {
	RabbitMqConfig := config.GetRabbitMQ()
	conn, err := amqp.Dial(RabbitMqConfig.URL)
	if err != nil {
		panic("Failed to connect to RabbitMQ")
	}
	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}
	return conn, ch
}
