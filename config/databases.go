package config

/*
	数据库配置文件
*/

const env  = "TEST" //LOCAL 开发环境 TEST 线上测试环境  ONLINE 正式环境

//mysql
type Mysql struct {
	DB_HOST         string //数据库服务器
	DB_DATABASE     string //数据库名称
	DB_USERNAME     string //数据库登录名
	DB_PASSWORD     string //数据库密码
	DB_PORT         string //数据库端口
	CHARSET         string //字符集
	SetMaxIdleConns int    //默认打开数据库的连接数
	SetMaxOpenConns int    //最大打开数据库的连接数

}

//mongodb
type Mongodb struct {
	URL string
}

//RabbitMQ
type RabbitMQ struct {
	URL string
} 

//redis
type Redis struct {
	REDIS_HOST     string
	REDIS_PASSWORD string
	REDIS_PORT     string
}

//mysql配置
func GetMysqlConf() Mysql {
	switch env {
		case "LOCAL":
			return Mysql{
				DB_HOST:         "192.168.31.231",
				DB_DATABASE:     "browser",
				DB_USERNAME:     "root",
				DB_PASSWORD:     "123456",
				DB_PORT:         "3306",
				CHARSET:         "utf8",
				SetMaxIdleConns: 10,
				SetMaxOpenConns: 10,
			}
		case "TEST":
			return Mysql{
				DB_HOST:         "10.200.124.23",
				DB_DATABASE:     "browser",
				DB_USERNAME:     "Mysqladmin",
				DB_PASSWORD:     "Mysql1707!",
				DB_PORT:         "3306",
				CHARSET:         "utf8",
				SetMaxIdleConns: 10,
				SetMaxOpenConns: 10,
			}
		default:
			return Mysql{
				DB_HOST:         "192.168.31.231",
				DB_DATABASE:     "browser",
				DB_USERNAME:     "root",
				DB_PASSWORD:     "123456",
				DB_PORT:         "3306",
				CHARSET:         "utf8",
				SetMaxIdleConns: 10,
				SetMaxOpenConns: 10,
			}
	}

}

//mongodb配置
func GetMongodbConf() Mongodb {
	return Mongodb{
		URL: "mongodb://hhq163:bx123456@192.168.31.231:27017/live_game?connect=direct&maxPoolSize=10",
	}
}

//redis配置
func GetRedisConf() Redis {
	switch env {
	case "LOCAL":
		return Redis{
			REDIS_HOST:     "192.168.31.230",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}
	case "TEST":
		return Redis{
			REDIS_HOST:     "10.200.124.21",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}
	default:
		return Redis{
			REDIS_HOST:     "192.168.31.230",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}


	}


}

//RabbitMQ配置信息
func GetRabbitMQ() RabbitMQ  {
	switch env {
	case "LOCAL":
		return RabbitMQ{
			URL: "amqp://lebo2017:ljf12345@192.168.31.230:5672/test",
		}
	case "TEST":
		return RabbitMQ{
			URL: "amqp://lebo2017:ljf12345@103.196.124.21:5672/test",
		}
	default:
		return RabbitMQ{
			URL: "amqp://lebo2017:ljf12345@192.168.31.230:5672/test",
		}

	}

}
