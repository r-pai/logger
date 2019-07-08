## logger
logger package for Golang. The purpose of creating this package was to learn golang more , a use case 
to apply what I have been learning in golang for some days now.
The code is not fully tested and there are some more works to be done like setting the loglevel.
Please comment on any issues or any feedback. I am happy to learn from the mistakes.

The log file currently created in the format 'helloApp_07-08-2019.log' 

## install
```
go get github.com/r-pai/logger
```

## usage
```golang
package main

import (
	"fmt"
	"github.com/r-pai/logger"
	"time"
)

func number() int {
	num := 15 * 5
	return num
}

func go1() {
	for i := 0; i < 1000; i++ {
		logger.Log(logger.LDEBUG, fmt.Sprintf("%s %d", "go1 Log message", i))
		time.Sleep(400 * time.Millisecond)
	}
}

func go2() {
	for i := 1000; i < 3000; i++ {
		logger.Log(logger.LINFO, fmt.Sprintf("%s %d", "go2 Log message", i))
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	//Create the logger
	logger.CreateLogger("./MyApp", logger.LDEBUG)
}

//main entry point
func main() {

	logger.Log(logger.LDEBUG, "Starting Hello LDEBUG")
	logger.Log(logger.LINFO, "Starting Hello LINFO")
	logger.Log(logger.LWARN, "Starting Hello LWARN")
	logger.Log(logger.LERROR, "Starting Hello LERROR")
	logger.Log(logger.LFATAL, "Starting Hello LFATAL")

	go go1()
	go go2()

	for {
	}

}


```
## Loggeroutput

Sample output of file : MyApp_07-09-2019.log

2019/07/09 00:57:24 DEBUG  hello.go:37  Starting Hello LDEBUG  
2019/07/09 00:57:24 INFO   hello.go:38  Starting Hello LINFO  
2019/07/09 00:57:24 WARN   hello.go:39  Starting Hello LWARN  
2019/07/09 00:57:24 ERROR  hello.go:40  Starting Hello LERROR  
2019/07/09 00:57:24 FATAL  hello.go:41  Starting Hello LFATAL  
2019/07/09 00:57:24 DEBUG  hello.go:17  go1 Log message 0  
2019/07/09 00:57:24 INFO   hello.go:24  go2 Log message 1000      
2019/07/09 00:57:24 INFO   hello.go:24  go2 Log message 1001    
2019/07/09 00:57:24 INFO   hello.go:24  go2 Log message 1002  
2019/07/09 00:57:24 INFO   hello.go:24  go2 Log message 1003  
2019/07/09 00:57:24 DEBUG  hello.go:17  go1 Log message 1  
2019/07/09 00:57:24 INFO   hello.go:24  go2 Log message 1004  
2019/07/09 00:57:25 INFO   hello.go:24  go2 Log message 1005  
2019/07/09 00:57:25 INFO   hello.go:24  go2 Log message 1006  
2019/07/09 00:57:25 INFO   hello.go:24  go2 Log message 1007  
2019/07/09 00:57:25 DEBUG  hello.go:17  go1 Log message 2  
2019/07/09 00:57:25 INFO   hello.go:24  go2 Log message 1008  



