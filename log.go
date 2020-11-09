package log

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
	logger = NewLogger()
)

/**
 * init logger
 */
func init() {
	SetCallDepth(3)
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
	logger.SetLogger(logName, flag, level)
}

/**
 * set call depth
 * @param int callDepth:
 */
func SetCallDepth(callDepth int) {
	logger.SetCallDepth(callDepth)
}

/**
 * @description:			to record Trace level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Trace(format string, v ...interface{}) {
	logger.Trace(format, v...)
}

/**
 * @description:			to record Info level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

/**
 * @description:			to record Warning level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Warning(format string, v ...interface{}) {
	logger.Warning(format, v...)
}

/**
 * @description:			to record Error level log
 * @param string format: 	the format information
 * @param ...interface{} v:	the format params
 */
func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}
