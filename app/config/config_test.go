// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"sort"
	"testing"

	"github.com/murosan/shogi-board-server/app/lib/stringutil"
	"github.com/murosan/shogi-board-server/app/lib/test_helper"
	"go.uber.org/zap"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		appyml, logyml string
		enginePaths    []string
		engineNames    []string
	}{
		{`
engines:
  com: '/home/user/path/to/engine/bin'
`,
			`
level: 'debug'
encoding: 'console'
encoderConfig:
  messageKey: 'Msg'
  levelKey: 'Level'
  timeKey: 'Time'
  nameKey: 'name'
  callerKey: 'Caller'
  stacktraceKey: 'St'
  levelEncoder: ''
  timeEncoder: 'iso8601'
  durationEncoder: 'string'
  callerEncoder: 'short'
outputPaths:
  - 'stdout'
errorOutputPaths:
  - 'stderr'
`,
			[]string{"/home/user/path/to/engine/bin"},
			[]string{"com"},
		},
	}

	for i, c := range cases {
		conf := NewConfig([]byte(c.appyml), []byte(c.logyml))
		names := conf.GetEngineNames()
		sort.Strings(names)

		// GetEngineNames() と GetEnginePath() のテスト
		if stringutil.SliceEquals(conf.GetEngineNames(), c.engineNames) {
			for j := range names {
				p1 := conf.GetEnginePath(names[j])
				p2 := c.enginePaths[j]
				if p1 != p2 {
					failing(t, "EnginePath", j, p2, p1)
				}
			}
		} else {
			failing(t, "EngineNames", i, c.engineNames, names)
		}
	}
}

// エラーのテスト
func TestNewConfig2(t *testing.T) {
	c := struct {
		appyml, logyml string
		enginePaths    []string
		engineNames    []string
		log            zap.Config
	}{
		`
# invalid syntax
engines
  com: /home/user/path/to/engine/bin
`,
		``,
		[]string{"/home/user/path/to/engine"},
		[]string{"com"},
		zap.Config{},
	}

	errMsg := "Expected panic, but there wasn't.\nInput: " + c.appyml
	testhelper.MustPanic(t, func() {
		NewConfig(
			[]byte(c.appyml),
			[]byte(c.logyml),
		)
	}, errMsg)
}

func failing(t *testing.T, key string, i int, expected, actual interface{}) {
	t.Helper()
	t.Errorf(`%s was not equal to as expected.
i: %d
Expected: %v
Actual:   %v`, key, i, expected, actual)
}
