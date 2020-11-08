package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/**
 * log type
 */
const (
	None        uint = 0x00 // no log output
	Console     uint = 0x01 // console log
	File        uint = 0x02 // file log
	ConsoleFile uint = 0x03 // both console and file log
)

/**
 * log level
 */
type LogLevel uint8
const (
	LevelTrace   LogLevel = 0
	LevelInfo    LogLevel = 1
	LevelWarning LogLevel = 2
	LevelError   LogLevel = 3
	logPath               = "log"
)

/**
 * log obj
 */
var (
	logTrace   *log.Logger // Trace log
	logInfo    *log.Logger // Info log
	logWarning *log.Logger // Warning log
	logError   *log.Logger // Error log
	logLevel   LogLevel    // log level
)

/**
 * init log obj
 * @description: default init log obj, default will output nothing.
 */
func init() {
	logTrace = log.New(ioutil.Discard, "[T] ", log.Ldate|log.Ltime|log.Lshortfile)
	logInfo = log.New(ioutil.Discard, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
	logWarning = log.New(ioutil.Discard, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
	logError = log.New(ioutil.Discard, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
}

/**
 * set log options
 * @description: 			If the SetLogger function is not called, there is nothing to output
 * @param string logName: 	the log file name
 * @param uint flag: 		the log flag (None, Console, File, ConsoleFile)
 * @param LogLevel level: 	the log level (LevelTrace, LevelInfo, LevelWarning, LevelError)
 * @return
 */
func SetLogger(logName string, flag uint, level LogLevel) {
	var (
		out  io.Writer
		file *os.File
		err  error
	)

	if flag == None {
		log.Println("log flag is None, there is nothing to output.")
	}
	// if log contain file log type, then create log dir and log file.
	if flag&File == File {
		if logName == "" {
			logName = "default.log"
		}
		if !strings.HasSuffix(logName, ".log") {
			logName += ".log"
		}

		if err = os.Mkdir(logPath, 0666); err != nil && !os.IsExist(err) {
			log.Println("log folder is not Exist, create folder log err:", err.Error())
		}

		logPath := logPath + string(os.PathSeparator) + logName
		file, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Failed to open error log file:", err)
		}
	}

	// switch flag to confirm log type
	switch flag {
	case None:
		out = ioutil.Discard
	case Console:
		out = os.Stdout
	case File:
		out = file
	default: // ConsoleFile and others
		out = io.MultiWriter(os.Stdout, file)
	}

	// set log level and log output
	logLevel = level
	logTrace.SetOutput(out)
	logInfo.SetOutput(out)
	logWarning.SetOutput(out)
	logError.SetOutput(out)
}

/**
 * @description:			to record Trace level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Trace(format string, v ...interface{}) {
	if logLevel <= LevelTrace {
		logTrace.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @description:			to record Info level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Info(format string, v ...interface{}) {
	if logLevel <= LevelInfo {
		logInfo.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @description:			to record Warning level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Warning(format string, v ...interface{}) {
	if logLevel <= LevelWarning {
		logWarning.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @description:			to record Error level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Error(format string, v ...interface{}) {
	if logLevel <= LevelError {
		logError.Println(fmt.Sprintf(format, v...))
	}
}
