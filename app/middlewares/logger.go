package middlewares

import (
	"api/pkg/helpers"
	"api/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取response内容
		w := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ := ioutil.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始的时间
		start := time.Now()
		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		respStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", respStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request body", string(requestBody)))

			// 响应的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if respStatus > 400 && respStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning:"+cast.ToString(respStatus), logFields...)
		} else if respStatus >= 500 && respStatus <= 599 {
			// 除了内部错误，记录 error
			logger.Error("HTTP Error "+cast.ToString(respStatus), logFields...)
		} else {
			logger.Debug("HTTP Access log", logFields...)
		}
	}
}