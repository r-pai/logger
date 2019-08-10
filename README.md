## # logger (A logger package for Go)

This Golang logger package can be used to log the application messages to a file.

Its prettly simple to use. 

Create the logger Instance - Use *'CreateLogger'* function  
There are 5 functions available to log, all of these linked to a log level.
- Debug(format string, a ...interface{})
- Info(format string, a ...interface{})
- Warn(format string, a ...interface{})
- Error(format string, a ...interface{})
- Fatal(format string, a ...interface{})

Each of the above functions are assoiated to 5 Loglevels.
- LDebug 
- LInfo
- LWarn
- LError
- LFatal

Any of the log level functions can be called from any goroutine and the logging would not be affected.  

***NOTE**  
    Code is not fully tested in real applications.
 
Please comment on any issues or provide feedback on how to improve.  

## # How to install
```
go get github.com/r-pai/logger
```

## # How to use

### import the package

>import logger "github.com/r-pai/logger"

### Create Logger in main or in init.  

>logger.CreateLogger("./", "Myapp", logger.LDebug)

### Log the messages - 5 Types of LogLevels 

> logger.Debug("Starting Hello LDEBUG")
> logger.Info("Starting Hello LINFO")
> logger.Warn("Starting Hello LWARN")
> logger.Error("Starting Hello LERROR")
> logger.Fatal("Starting Hello LFATAL"

## # Sample Code 1
```golang
package main

import ( 
       logger "github.com/r-pai/logger"
       )

func main() {
        //Create the logger
	logger.CreateLogger("./", "Myapp", logger.LDebug)
	
	logger.Debug("Starting Hello LDEBUG")
        logger.Info("Starting Hello LINFO")
        logger.Warn("Starting Hello LWARN")
        logger.Error("Starting Hello LERROR")
        logger.Fatal("Starting Hello LFATAL"
}
```

## # Sample Code 2
```golang
package main

import (
	logger "github.com/r-pai/logger"
	"time"
)

func number() int {
	num := 15 * 5
	return num
}

func go1() {
	for i := 0; i < 1000; i++ {
		logger.Info"%s %d", "go1 Log message", i)
		time.Sleep(400 * time.Millisecond)
	}
}

func go2() {
	for i := 1000; i < 3000; i++ {
		logger.Info("%s %d", "go2 Log message", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	//Create the logger
	       logger.CreateLogger("./", "Myapp", logger.LDebug)
}

//main entry point
func main() {

	go go1()
	go go2()

	for {
	}

}


```
## # Loggeroutput

Sample output of file : **MyApp_07-09-2019.log**


![Screenshot 2019-07-09 at 10 08 59 AM](https://user-images.githubusercontent.com/33278265/60861025-81bb9a00-a236-11e9-8697-a8e330dfd0f0.png)




