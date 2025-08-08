package main

import (
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/net/ghttp"
)

// https://github.com/mickbai/gf/issues/437
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("client/layout.html")
	})
	s.SetPort(8199)
	s.Run()
}
