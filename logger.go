package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//GLogger Struct
type GLogger struct {
	logger        *log.Logger
	logFileName   string
	logFilePath   string
	logTimeFormat string
	logLevel      gloglevel
	logAppRoot    string
	logJSONFormat bool
}

//Exported loglevel constant
const (
	LDebug gloglevel = iota
	LInfo
	LWarn
	LError
	LFatal
)

type gloglevel int

type jsonLog struct {
	Severity string `json:"severity"`
	Datetime string `json:"datetime"`
	FileName string `json:"filename"`
	LineNo   int    `json:"lineno"`
	Message  string `json:"message"`
}

var (
	//Logger pointer
	glogger GLogger

	//Log Level Map has log level string
	glogLevelMap map[gloglevel]string

	//Send Channel is to send log message to logRoutine
	glogSendChannel chan string

	//Recv Channel is to get ack from logroutine for the Log message
	glogRecvChannel chan int
)

func initLogger() {
	glogLevelMap = make(map[gloglevel]string)

	glogLevelMap[LDebug] = "DEBUG"
	glogLevelMap[LInfo] = "INFO"
	glogLevelMap[LWarn] = "WARN"
	glogLevelMap[LError] = "ERROR"
	glogLevelMap[LFatal] = "FATAL"

	glogSendChannel = make(chan string)
	glogRecvChannel = make(chan int)

	go logRoutine(glogSendChannel)
}

//CreateLogger function will create the logger
//TODO: 1. Validate the filePath
//      2. logging error in case of failure
func CreateLogger(logFilePath string, logFileName string, loglevel gloglevel) *GLogger {

	initLogger()

	logPathFileName := logFilePath + "/" + logFileName
	logFile := fmt.Sprintf("%s_%s.log", logPathFileName, time.Now().Format("01-02-2006"))
	logTimeFormat := "2006-01-02 15:04:05.000000"

	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error: Could not open file: %v", err)
	}

	glogger = GLogger{log.New(file, "", 0), logFileName, logFilePath, logTimeFormat, loglevel, "", false}
	return &glogger
}

func (g *GLogger) isNil() bool {
	if g != nil {
		return true
	}
	return false
}

//SetAppRootFolder sets the root folder
func (g *GLogger) SetAppRootFolder(rootFolderName string) {
	if g.isNil() {
		g.logAppRoot = rootFolderName
	}
}

//SetLogTimeFormat sets the root folder
func (g *GLogger) SetLogTimeFormat(logTimeFormat string) {
	if g.isNil() {
		g.logTimeFormat = logTimeFormat
	}
}

//SetJSONLog sets the root folder
func (g *GLogger) SetJSONLog(set bool) {
	g.logJSONFormat = set
}

func internalLog(loglevel gloglevel, message string) {
	if loglevel >= glogger.logLevel {

		fileName, fLineNo := getFName()
		logMessage := jsonLog{glogLevelMap[loglevel],
			time.Now().Format(glogger.logTimeFormat),
			fileName,
			fLineNo,
			message}.format()

		sendToLogRoutine(logMessage)
	}
}

func sendToLogRoutine(logMessage string) {
	glogSendChannel <- logMessage
	_ = <-glogRecvChannel
}

func logRoutine(logSendChannel chan string) {
	for {
		logMessage := <-logSendChannel
		glogger.logger.Println(logMessage)
		glogRecvChannel <- 1
	}
}

func (j jsonLog) format() string {
	if glogger.logJSONFormat {
		jformat, _ := json.Marshal(j)
		return string(jformat)
	}

	format := fmt.Sprintf("%s %-6v %s:%d %s", j.Datetime, j.Severity, j.FileName, j.LineNo, j.Message)
	return format
}

func getFName() (fName string, fLine int) {
	_, fPath, fLine, ok := runtime.Caller(3)
	if ok {
		if glogger.logAppRoot == "" {
			fName = filepath.Base(fPath)
		} else {
			s := strings.Split(fPath, glogger.logAppRoot)
			if len(s) >= 2 {
				fName = s[1][1:]
			} else {
				fName = filepath.Base(fPath)
			}

		}
	}
	return
}

//Debug Function
func Debug(format string, a ...interface{}) {
	internalLog(LDebug, fmt.Sprintf(format, a...))
}

//Info Function
func Info(format string, a ...interface{}) {
	internalLog(LInfo, fmt.Sprintf(format, a...))
}

//Warn Function
func Warn(format string, a ...interface{}) {
	internalLog(LWarn, fmt.Sprintf(format, a...))
}

//Error Function
func Error(format string, a ...interface{}) {
	internalLog(LError, fmt.Sprintf(format, a...))
}

//Fatal Function
func Fatal(format string, a ...interface{}) {
	internalLog(LFatal, fmt.Sprintf(format, a...))
}
