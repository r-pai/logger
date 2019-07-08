package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//Loglevel type as int
type Loglevel int

//Exported Logger Variables
var (
	Logger         *log.Logger
	LOGGERLOGLEVEL = LINFO
)

//Exported LogLevel constant
const (
	LDEBUG Loglevel = iota
	LINFO
	LWARN
	LERROR
	LFATAL
)

var (
	logLevelMap map[Loglevel]string
	logChannel  chan string
)

//initLogLevelMap Inizlizes the log level Map
func initLogLevelMap() {
	logLevelMap = make(map[Loglevel]string)

	logLevelMap[LDEBUG] = "DEBUG"
	logLevelMap[LINFO] = "INFO"
	logLevelMap[LWARN] = "WARN"
	logLevelMap[LERROR] = "ERROR"
	logLevelMap[LFATAL] = "FATAL"
}

//initLogLevelMap Inizlizes the log level Map
func initLogRoutine() {
	logChannel = make(chan string)

	//Infinite loop go routine
	go LogRoutine(logChannel)
}

//init function will Initialize the logger
func init() {

	//Init LogLevel Details
	initLogLevelMap()

	//start Go Routine
	initLogRoutine()

	currentTime := time.Now()
	logFileName := fmt.Sprintf("%s_%s.log", "helloApp", currentTime.Format("01-02-2006"))

	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	Logger = log.New(file, "", log.LstdFlags)
}

//Log is the entry point to log message
func Log(logLevel Loglevel, message string) {

	if true == validateLogLevel(logLevel) {

		//Get the FileName and Line number of the Log function caller
		fileLine := ""
		_, fPath, fLine, ok := runtime.Caller(1)
		if ok {
			fName := filepath.Base(fPath)
			fileLine = fmt.Sprintf("%s:%d", fName, fLine)
		}

		//format the message and send to Log routine
		logMessage := fmt.Sprintf("%-6v %-12v %s", logLevelMap[logLevel], fileLine, message)
		sendToLogRoutine(logMessage)
	}
}

//validateLogLevel formats the message to desired type
func validateLogLevel(logLevel Loglevel) bool {
	if logLevel >= LOGGERLOGLEVEL {
		return true
	}
	return false
}

//sendToLogRoutine will send the log message to the routine
func sendToLogRoutine(logMessage string) {
	logChannel <- logMessage
}

//LogRoutine purpose is to write log to the file
func LogRoutine(logChannel chan string) {
	fmt.Println("Started LogRoutine")
	for {
		logMessage := <-logChannel
		Logger.Println(logMessage)
	}
}
