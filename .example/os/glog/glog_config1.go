package main

import (
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/os/glog"
)

func main() {
	err := glog.SetConfigWithMap(g.Map{
		"prefix": "[TEST]",
	})
	if err != nil {
		panic(err)
	}
	glog.Info(1)
}
