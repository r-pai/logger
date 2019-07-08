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
	"learn1/logger"
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

func main() {

  //All the available
	logger.Log(logger.LDEBUG, "Starting Hello LDEBUG")
	logger.Log(logger.LINFO, "Starting Hello LINFO")
	logger.Log(logger.LWARN, "Starting Hello LWARN")
	logger.Log(logger.LERROR, "Starting Hello LERROR")
	logger.Log(logger.LFATAL, "Starting Hello LFATAL")

  //Writing from go routines
	go go1()
	go go2()

	for {
	}

}

```

## TODO
1. Add method to set the log file Name
2. Add method to set the logfile path 
3. Add method to set the loglevel



