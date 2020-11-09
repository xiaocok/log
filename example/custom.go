/**
 * @author: gitteamer
 * @date: 2020/11/9
 * @note:
 */
package main

import "github.com/gitteamer/log"

func main() {
	logger := log.NewLogger()
	logger.SetLogger("custom", log.ConsoleFile, log.LevelTrace)
	logger.SetCallDepth(2)

	logger.Trace("this is trace log.")
	logger.Info("this is info log.")
	logger.Warning("this is warning log.")
	logger.Error("this is error log.")
}
