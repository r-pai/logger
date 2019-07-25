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


## # Sample Code 1
```golang
package main

import "github.com/r-pai/logger"

func main() {
        //Create the logger
	logger.CreateLogger("./MyApp", logger.LDEBUG)
	
	logger.Log(logger.LDEBUG, "Starting Hello LDEBUG")
	logger.Log(logger.LINFO, "Starting Hello LINFO")
	logger.Log(logger.LWARN, "Starting Hello LWARN")
	logger.Log(logger.LERROR, "Starting Hello LERROR")
	logger.Log(logger.LFATAL, "Starting Hello LFATAL")
}
```

## # Sample Code 2
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


![Screenshot 2019-07-09 at 10 08 59 AM](https://user-images.githubusercontent.com/33278265/60861025-81bb9a00-a236-11e9-8697-a8e330dfd0f0.png)




