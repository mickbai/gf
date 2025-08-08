package main

import (
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/os/gfile"
	"github.com/mickbai/gf/os/glog"
)

// 设置日志输出路径
func main() {
	path := "/tmp/glog"
	glog.SetPath(path)
	glog.Println("日志内容")
	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)
}
