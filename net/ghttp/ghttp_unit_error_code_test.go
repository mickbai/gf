// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/mickbai/gf.

// static service testing.

package ghttp_test

import (
	"fmt"
	"github.com/mickbai/gf/errors/gerror"
	"github.com/mickbai/gf/net/ghttp"
	"testing"
	"time"

	"github.com/mickbai/gf/frame/g"
	"github.com/mickbai/gf/test/gtest"
)

func Test_Error_Code(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(func(r *ghttp.Request) {
				r.Middleware.Next()
				r.Response.ClearBuffer()
				r.Response.Write(gerror.Code(r.GetError()))
			})
			group.ALL("/", func(r *ghttp.Request) {
				panic(gerror.NewCode(10000, "test error"))
			})
		})
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.GetContent("/"), "10000")
	})
}
