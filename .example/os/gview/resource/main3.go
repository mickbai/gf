package main

import (
	"fmt"

	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/os/gres"
	_ "github.com/mickbai/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
