package main

import (
	"time"

	"github.com/mickbai/gf/os/gcron"
	"github.com/mickbai/gf/os/glog"
)

func main() {
	gcron.SetLogLevel(glog.LEVEL_ALL)
	gcron.Add("* * * * * ?", func() {
		glog.Println("test")
	})
	time.Sleep(3 * time.Second)
}
