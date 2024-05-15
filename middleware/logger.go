package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func GinLoggerPrintConf() gin.HandlerFunc {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// End timer
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := statusCodeColor(statusCode)
		methodColor := methodColor(method)
		resetColor := resetColor()

		log.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   clientIP,
			"method":      method,
			"path":        path,
		}).Infof("%s %3d %s| %13v | %15s |%s %-7s %s %s",
			statusColor, statusCode, resetColor,
			latency,
			clientIP,
			methodColor, method, resetColor,
			path,
		)
	}
}

func statusCodeColor(code int) string {
	switch {
	case code >= 400 && code <= 499:
		return "\033[31m" // Red
	case code >= 500:
		return "\033[31m" // Red
	default:
		return "\033[32m" // Green
	}
}

func methodColor(method string) string {
	switch method {
	case "POST":
		return "\033[36m" // Cyan
	case "GET":
		return "\033[32m" // Green
	case "PUT":
		return "\033[33m" // Yellow
	case "DELETE":
		return "\033[31m" // Red
	default:
		return "\033[0m" // No color
	}
}

func resetColor() string {
	return "\033[0m"
}
