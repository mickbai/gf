package main

import (
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		g.Dump(r.GetForm("array"))
		r.Response.WriteTpl("form.html")
	})
	s.SetPort(8199)
	s.Run()
}
