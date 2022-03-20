package logs

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/sirupsen/logrus"
)

// InitLogs 初始化日志系统
func InitLogs() {
	// 设置日志输出时加上文件名和方法信息
	logrus.SetReportCaller(true)
	var w io.Writer = &lumberjack.Logger{
		Filename:   "./runtime/logs/runtime.log",
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	f := &logrus.JSONFormatter{}
	logrus.SetFormatter(f)
	logrus.SetOutput(w)
	logrus.Infof("log system init success")
}
