package main

import (
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/os/gbuild"
)

func main() {
	g.Dump(gbuild.Info())
	g.Dump(gbuild.Map())
}
