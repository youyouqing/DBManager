package Logger

import (
	"bytes"
	"dzc.com/Config"
	"dzc.com/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/json-iterator/go"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"time"
)

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewBodyWriter(resWriter gin.ResponseWriter) *BodyWriter {
	return &BodyWriter{body: bytes.NewBufferString(""), ResponseWriter: resWriter}
}

func (this *BodyWriter) Write(buf []byte) (int, error) {
	this.body.Write(buf)
	return this.ResponseWriter.Write(buf)
}

func (this *BodyWriter) WriteString(s string) (int, error) {
	this.body.WriteString(s)
	return this.ResponseWriter.WriteString(s)
}

var accessLogChannel = make(chan string, 100)

func SetUp() gin.HandlerFunc {

	go handleWriteLog() // 异步写日志

	return func(context *gin.Context) {
		bodyWriter := NewBodyWriter(context.Writer)
		context.Writer = bodyWriter

		startTime := time.Now().UnixNano() // 开始时间
		context.Next()

		responseBody := bodyWriter.body.String() // 响应参数

		endTime := time.Now().UnixNano() // 结束时间

		// 日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = context.Request.Method
		accessLogMap["request_uri"] = context.Request.RequestURI
		accessLogMap["request_proto"] = context.Request.Proto
		accessLogMap["request_ua"] = context.Request.UserAgent()
		accessLogMap["request_referer"] = context.Request.Referer()
		accessLogMap["request_post_data"] = context.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = context.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_string"] = responseBody

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		jsonByte, err := json.Marshal(&accessLogMap)
		if err != nil {
			log.Panic(err)
		} else {
			jsonByteString := string(jsonByte)
			accessLogChannel <- jsonByteString
		}
	}
}

func handleWriteLog() {
	logFileName := Utils.CreateDateDir(Config.AppAccessLogName, os.ModePerm)
	if f, err := os.OpenFile(logFileName+"/"+Config.AccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		log.Println(err)
	} else {
		for accessLog := range accessLogChannel {
			_, _ = f.WriteString(accessLog + "\n")
		}
	}
	return
}
