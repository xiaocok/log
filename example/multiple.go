package main

import (
	"github.com/gitteamer/log"
	"time"
)

func main() {
	log.SetLogger("multiple", log.ConsoleFile, log.LevelTrace)

	times := 100
	go func() {
		for i := 0; i < times; i++ {
			log.Trace("this is trace log.")
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			log.Info("this is info log.")
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			log.Warning("this is warning log.")
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			log.Error("this is error log.")
			time.Sleep(time.Microsecond)
		}
	}()

	time.Sleep(time.Second * 10)
}
