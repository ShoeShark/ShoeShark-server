package logger

import (
	"fmt"
	"github.com/shoe-shark/shoe-shark-service/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger() {
	cfg := config.GetConfig()
	// 设置日志格式。
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logConf := cfg.Logging
	switch logConf.Level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
	log.SetReportCaller(true) // 打印文件、行号和主调函数。

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "dev" {
		log.SetOutput(os.Stdout)
	} else {
		// 实现日志滚动。
		// Refer to https://www.cnblogs.com/jssyjam/p/11845475.html.
		logger := &lumberjack.Logger{
			Filename:   fmt.Sprintf("%v/%v", logConf.Dir, logConf.Name), // 日志输出文件路径。
			MaxSize:    logConf.MaxSize,                                 // 日志文件最大 size(MB)，缺省 100MB。
			MaxBackups: 10,                                              // 最大过期日志保留的个数。
			MaxAge:     30,                                              // 保留过期文件的最大时间间隔，单位是天。
			LocalTime:  true,                                            // 是否使用本地时间来命名备份的日志。
		}
		log.SetOutput(logger)
	}

}
