// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/mickbai/gf.

package gutil_test

import (
	"github.com/mickbai/gf/frame/g"
	"testing"

	"github.com/mickbai/gf/test/gtest"
	"github.com/mickbai/gf/util/gutil"
)

func Test_SliceToMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2", "v2",
		}
		m := gutil.SliceToMap(s)
		t.Assert(len(m), 2)
		t.Assert(m, g.Map{
			"K1": "v1",
			"K2": "v2",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		s := g.Slice{
			"K1", "v1", "K2",
		}
		m := gutil.SliceToMap(s)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
}
