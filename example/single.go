package main

import (
	"github.com/gitteamer/log"
)

func main() {
	log.SetLogger("single", log.ConsoleFile, log.LevelTrace)

	log.Trace("this is trace log.")
	log.Info("this is info log.")
	log.Warning("this is warning log.")
	log.Error("this is error log.")
}
