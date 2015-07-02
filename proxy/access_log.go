// Save access record to log file.
package proxy

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	LOG_FORMAT = "user{%s} appid{%d} visit scope{%s} from ip{%s}"
)

var (
	LoggerForParse *log.Logger
)

type LogWrapper struct {
	AppId    int64
	Scope    string
	ClientIp string
	UID      string // User ID , such as smartphone mac address.
}

// // put log file to path
func InitAccessLogger(logPath string) {
	if logPath == "" {
		log.Println("Warn! Access log folder not been assigned, all info will print in console")
		// log path is empty.
		initLoggerHandle(os.Stdout)
	} else {
		// prepare a file to write log.
		f, err := CreateOrAppendFile(logPath)
		if err != nil {
			fmt.Errorf("error opening file: %v", err)
		}
		defer f.Close()
		initLoggerHandle(f)
	}
}

// Create a file if not exists , append this file.
func CreateOrAppendFile(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func initLoggerHandle(analysisHandle io.Writer) {
	LoggerForParse = log.New(analysisHandle, "PROXY ", 0)
}

// Write result to access log,
func writeAccessLog(raw string) {
	LoggerForParse.Println(timeFormat(), raw)
}

// return access log time format.
func timeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// write access log in file for analysis.
func WriteAccessLogByApiInfo(a LogWrapper) {
	content := fmt.Sprintf(LOG_FORMAT, a.UID, a.AppId, a.Scope, a.ClientIp)
	writeAccessLog(content)
}
