// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"fmt"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-proxy-server/app/lib/test_helper"
)

func TestFromUsi_Info(t *testing.T) {
	cases := []struct {
		in   string
		want *result.Info
		mpv  int
		err  error
	}{
		{"info time 1141 depth 3 nodes 135125 score cp -1521 pv 3a3b L*4h 4c4d",
			&result.Info{
				Values: map[string]int{
					result.Time:  1141,
					result.Depth: 3,
					result.Nodes: 135125,
				},
				Score: -1521,
				Moves: []shogi.Move{
					{[]int{2, 0}, []int{2, 1}, 0, false},
					{[]int{-1, -1}, []int{3, 7}, 2, false},
					{[]int{3, 2}, []int{3, 3}, 0, false},
				},
			},
			-1, nil},
	}

	for i, c := range cases {
		infoHelper(t, i, c.in, c.want, c.mpv, c.err)
	}
}

func infoHelper(t *testing.T, i int, in string, want *result.Info, mpv int, err error) {
	t.Helper()

	res, mpvKey, e := NewFromUsi().Info(in)
	fmt.Println(res)

	if (e == nil && err != nil) || (e != nil && err == nil) {
		t.Errorf(`(From Usi: Paese Info) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	// エラーだったが、想定と違った。
	if e != nil && err != nil && !strings.Contains(string(e.Error()), string(err.Error())) {
		t.Errorf(`(From Usi: Paese Info) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	if mpvKey != mpv {
		t.Errorf(`(From Usi: Parse Info) The multipv index value was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, mpv, mpvKey)
	}

	if !test_helper.InfoEquals(res, want) {
		t.Errorf(`(From Usi: Parse Info) The value was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, want, res)
	}
}
