package log

import (
	"github.com/op/go-logging"
	"os"
	"strings"
	"time"
)

var log *logging.Logger

/*
日志系统初始化
*/
func init() {
	log = logging.MustGetLogger("example")
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
	)
	s := strings.Join(strings.Split(time.Now().UTC().Format(time.UnixDate), " "), "-")
	logFile, err := os.Create("log/wx-" + s + ".log")
	CheckError(err)
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

/*
获取 Logger 指针对象
*/
func GetLogger() *logging.Logger {
	return log
}

/*
错误日志打印, 同时错误日志保存
*/
func CheckError(err error) {
	if err != nil {
		log.Errorf("Error is %s", err)
	}
}
