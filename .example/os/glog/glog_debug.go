package main

import (
	"time"

	"github.com/mickbai/gf/os/glog"
	"github.com/mickbai/gf/os/gtime"
	"github.com/mickbai/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
