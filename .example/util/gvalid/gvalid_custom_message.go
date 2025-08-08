package main

import (
	"fmt"
	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check("", "required", nil)
	fmt.Println(err.String())
}
