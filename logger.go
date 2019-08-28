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
	logger *log.Logger

	//Configs
	logFileName   string
	logFilePath   string
	logTimeFormat string
	logLevel      gloglevel
	logAppRoot    string
	logJSONFormat bool

	//LoggerChannels
	//Send Channel is to send log message to logRoutine
	glogSendChannel chan string

	//Recv Channel is to get ack from logroutine for the Log message
	glogRecvChannel chan int
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

type logDetails struct {
	Severity string `json:"severity"`
	Datetime string `json:"datetime"`
	FileName string `json:"filename"`
	LineNo   int    `json:"lineno"`
	Message  string `json:"message"`
}

var (
	//Log Level Map has log level string
	glogLevelMap map[gloglevel]string
)

func initLogger() {
	glogLevelMap = make(map[gloglevel]string)

	glogLevelMap[LDebug] = "DEBUG"
	glogLevelMap[LInfo] = "INFO"
	glogLevelMap[LWarn] = "WARN"
	glogLevelMap[LError] = "ERROR"
	glogLevelMap[LFatal] = "FATAL"
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

	glogger := GLogger{
		log.New(file, "", 0),
		logFileName,
		logFilePath,
		logTimeFormat,
		loglevel,
		"",
		false,
		make(chan string),
		make(chan int)}

	go glogger.logRoutine(glogger.glogSendChannel)

	return &glogger
}

func (glogger *GLogger) isNil() bool {
	if glogger != nil {
		return true
	}
	return false
}

//SetAppRootFolder sets the root folder
func (glogger *GLogger) SetAppRootFolder(rootFolderName string) {
	if glogger.isNil() {
		glogger.logAppRoot = rootFolderName
	}
}

//SetLogTimeFormat sets the root folder
//Time format can be constants from time package
/*
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
*/
func (glogger *GLogger) SetLogTimeFormat(logTimeFormat string) {
	if glogger.isNil() {
		glogger.logTimeFormat = logTimeFormat
	}
}

//SetJSONLog sets the root folder
func (glogger *GLogger) SetJSONLog(set bool) {
	glogger.logJSONFormat = set
}

func (glogger *GLogger) internalLog(loglevel gloglevel, message string) {
	if loglevel >= glogger.logLevel {

		fileName, fLineNo := glogger.getFName()
		logD := logDetails{
			glogLevelMap[loglevel],
			time.Now().Format(glogger.logTimeFormat),
			fileName,
			fLineNo,
			message}
		logMessage := glogger.format(logD)

		glogger.sendToLogRoutine(logMessage)
	}
}

func (glogger *GLogger) sendToLogRoutine(logMessage string) {
	glogger.glogSendChannel <- logMessage
	_ = <-glogger.glogRecvChannel
}

func (glogger *GLogger) logRoutine(logSendChannel chan string) {
	for {
		logMessage := <-logSendChannel
		glogger.logger.Println(logMessage)
		glogger.glogRecvChannel <- 1
	}
}

func (glogger *GLogger) format(logD logDetails) string {
	if glogger.logJSONFormat {
		jformat, _ := json.Marshal(logD)
		return string(jformat)
	}

	format := fmt.Sprintf("%s %-6v %s:%d %s", logD.Datetime, logD.Severity, logD.FileName, logD.LineNo, logD.Message)
	return format
}

func (glogger *GLogger) getFName() (fName string, fLine int) {
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

//Debug logs message in debug log level
func (glogger *GLogger) Debug(format string, a ...interface{}) {
	glogger.internalLog(LDebug, fmt.Sprintf(format, a...))
}

//Info logs message in info log level
func (glogger *GLogger) Info(format string, a ...interface{}) {
	glogger.internalLog(LInfo, fmt.Sprintf(format, a...))
}

//Warn logs message in warning log level
func (glogger *GLogger) Warn(format string, a ...interface{}) {
	glogger.internalLog(LWarn, fmt.Sprintf(format, a...))
}

//Error logs message in error log level
func (glogger *GLogger) Error(format string, a ...interface{}) {
	glogger.internalLog(LError, fmt.Sprintf(format, a...))
}

//Fatal logs message in fatal log level
func (glogger *GLogger) Fatal(format string, a ...interface{}) {
	glogger.internalLog(LFatal, fmt.Sprintf(format, a...))
}
