## # logger (A logger package for Go)

This Golang logger package can be used to log the application messages to a file.

Its prettly simple to use. 

*'CreateLogger'* function - creates the logger and 'Log' functions log messages.   
There are 5 Loglevels currently which can be set to log 
- LDEBUG 
- LINFO 
- LWARN
- LERROR
- LFATAL

The *'Log'* function can be called from any goroutine and the logging would not be affected.  

***NOTE**  
    Code is not fully tested in real applications.
 
***TODO**  
- Add new functions to each loglevel to keep it simple  
- For each log level define custome types to add more information into logs  

Please comment on any issues or provide feedback on how to improve.  

## # How to install
```
go get github.com/r-pai/logger
```

## # How to use

### import the package

>import "github.com/r-pai/logger"

### Create Logger in main or in init.  
The 1st param is the FullPathfilename and 2nd the LogLevel for the application

>logger.CreateLogger("./MyApp", logger.LDEBUG)

### Log the messages - 5 Types of LogLevels 

>logger.Log(logger.LDEBUG , "Starting Hello LDEBUG")  
>logger.Log(logger.LINFO, "Starting Hello LINFO")  
>logger.Log(logger.LWARN, "Starting Hello LWARN")  
>logger.Log(logger.LERROR, "Starting Hello LERROR")  
>logger.Log(logger.LFATAL, "Starting Hello LFATAL")  


## # Sample Code
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
## # Loggeroutput

Sample output of file : **MyApp_07-09-2019.log**

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



