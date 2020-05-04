package config

const (
	// 服务类

	// 服务端口
	SERVER_PORT = ":8888"


	// 数据库类

	// mysql连接配置
	MYSQL_DSN = "root:zhicongdai@tcp(127.0.0.1:3306)/DBManager?charset=utf8mb4&parseTime=True&loc=Local"
	// redis
	REDIS_HOST = "127.0.0.1"
	REDIS_PORT = 6379
	REDIS_AUTH = "zhicongdai"

	AppName = "DBManager"
	// 日志文件
	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"
	AppGrpcLogName   = "log/" + AppName + "-grpc.log"



)
