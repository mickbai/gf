package main

import (
	"github.com/mickbai/gf/frame/g"
)

func main() {
	g.DB().Model("user").Distinct().CountColumn("uid,name")
}
