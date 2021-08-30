package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 全局变量
var (
	Loger          *log.Logger // 官方日志库
	LogLevel       int         // 日志级别
	IsAllName      bool        // 是否打印全路径名称
	FileNameMaxLen int         // 文件名长度
)

type structP *struct{}

// 日志级别
const (
	LOG_INFO  = 0
	LOG_ERROR = 1
	LOG_DEBUG = 2
)

// SetLogLeve 修改日志级别
func SetLogLeve(logLevel int) {
	LogLevel = logLevel
}

// GetLogLevel 获取日志级别
func GetLogLevel() int {
	return LogLevel
}

// SetLongNameMaxLen 修改文件名长度
func SetLongNameMaxLen(fileNameMaxLen int) {
	FileNameMaxLen = fileNameMaxLen
}

// GetLogLeve 获取文件名长度
func GetLongNameMaxLen() int {
	return FileNameMaxLen
}

// SetIsAllName 修改是否为全路径名称
func SetIsAllName(isAllName bool) {
	IsAllName = isAllName
}

// GetIsAllName 获取全路径名称标识
func GetIsIsAllName() bool {
	return IsAllName
}

// InitLog 日志初始化
func InitLog(fileName string, logLevel int) {
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	LogLevel = logLevel
	FileNameMaxLen = 20
	Loger = log.New(logFile, "", log.Ldate|log.Ltime)
	Loger.Println("initLog SUCC!!")
	Loger.Printf("LogLevel is [%s]\n", getLevel(LogLevel))
}

// TlogPrintln 打印日志对象
func TlogPrintln(logLevel int, v ...interface{}) {

	if logLevel > LogLevel {
		return
	}
	logStr := fmt.Sprintln(v...)
	tLogPrintf(getLevel(logLevel), logStr)
}

// TlogPrintf 格式化打印
func TlogPrintf(logLevel int, format string, v ...interface{}) {

	if logLevel > LogLevel {
		return
	}
	logStr := fmt.Sprintf(format, v...)
	tLogPrintf(getLevel(logLevel), logStr)
}

// getLevel 补全级别
func getLevel(logLevel int) string {
	var logLevelStr string
	switch logLevel {
	case 0:
		logLevelStr = " INFO"
	case 1:
		logLevelStr = "ERROR"
	case 2:
		logLevelStr = "DEBUG"
	}
	return logLevelStr
}

// 最终字符串输出
func tLogPrintf(logLevelStr string, buf string) {
	fileName, Line := getFileName()
	//fmt.Println(fileName)
	Loger.Printf("[ %-*s ]:[%5d][ %5s ]: %s", FileNameMaxLen, fileName, Line, logLevelStr, buf)
}

// 获取日志打印的文件名
func getFileName() (file string, line int) {

	//l.mu.Unlock()
	var ok bool
	_, filetmp, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	paths, file := filepath.Split(filetmp)
	if IsAllName {
		file = paths + file
	}
	//l.mu.Lock()
	return
}
