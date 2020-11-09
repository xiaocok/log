/**
 * @author: gitteamer
 * @date: 2020/11/9
 * @note: Logger struct
 */
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
 * new Logger obj and init log
 * @description: default init log obj, default will output nothing.
 */
func NewLogger() *Logger {
	l := new(Logger)
	l.logTrace = log.New(ioutil.Discard, "[T] ", log.Ldate|log.Ltime|log.Lshortfile)
	l.logInfo = log.New(ioutil.Discard, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
	l.logWarning = log.New(ioutil.Discard, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
	l.logError = log.New(ioutil.Discard, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	l.callDepth = 2

	return l
}

/**
 * Logger struct
 */
type Logger struct {
	logTrace   *log.Logger // Trace log
	logInfo    *log.Logger // Info log
	logWarning *log.Logger // Warning log
	logError   *log.Logger // Error log
	logLevel   LogLevel    // log level
	callDepth  int         // log filename call depth
}

/**
 * set log options
 * @description: 			If the SetLogger function is not called, there is nothing to output
 * @param string logName: 	the log file name
 * @param uint flag: 		the log flag (None, Console, File, ConsoleFile)
 * @param LogLevel level: 	the log level (LevelTrace, LevelInfo, LevelWarning, LevelError)
 * @return
 */
func (l *Logger) SetLogger(logName string, flag uint, level LogLevel) {
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

		if err = os.Mkdir(logPath, os.ModePerm); err != nil && !os.IsExist(err) {
			log.Println("log folder is not Exist, create folder log err:", err.Error())
		}

		logPath := logPath + string(os.PathSeparator) + logName
		file, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
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
	l.logLevel = level
	l.logTrace.SetOutput(out)
	l.logInfo.SetOutput(out)
	l.logWarning.SetOutput(out)
	l.logError.SetOutput(out)
}

/**
 * set call depth
 * @param int callDepth:
 */
func (l *Logger) SetCallDepth(callDepth int) {
	l.callDepth = callDepth
}

/**
 * @description:			to record Trace level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func (l *Logger) Trace(format string, v ...interface{}) {
	if l.logLevel <= LevelTrace {
		_ = l.logTrace.Output(l.callDepth, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

/**
 * @description:			to record Info level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func (l *Logger) Info(format string, v ...interface{}) {
	if l.logLevel <= LevelInfo {
		_ = l.logTrace.Output(l.callDepth, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

/**
 * @description:			to record Warning level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func (l *Logger) Warning(format string, v ...interface{}) {
	if l.logLevel <= LevelWarning {
		_ = l.logTrace.Output(l.callDepth, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}

/**
 * @description:			to record Error level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func (l *Logger) Error(format string, v ...interface{}) {
	if l.logLevel <= LevelError {
		_ = l.logTrace.Output(l.callDepth, fmt.Sprintln(fmt.Sprintf(format, v...)))
	}
}
