package main

import (
	"github.com/mickbai/gf/container/gmap"
	"github.com/mickbai/gf/frame/g"
)

func main() {
	m1 := gmap.New(true)
	m1.Set("1", "1")

	m2 := m1.Map()
	m2["2"] = "2"

	g.Dump(m1.Clone())
	g.Dump(m2)
	//output:
	//{
	//	"1": "1"
	//}
	//
	//{
	//	"1": "1",
	//	"2": "2"
	//}
}
