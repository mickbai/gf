package main

import (
	"fmt"

	"github.com/mickbai/gf/frame/g"
)

func main() {
	fmt.Println(g.Config().Get("none"))
}
