// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"

	"github.com/murosan/shogi-board-server/app/domain/exception"
	"github.com/murosan/shogi-board-server/app/service/logger"
	"go.uber.org/zap"
)

// OptMap Option の一覧を入れておくもの
// 管理しやすいように、オプションの種類別になっている
type OptMap struct {
	Buttons   map[string]*Button   `json:"buttons"`
	Checks    map[string]*Check    `json:"checks"`
	Spins     map[string]*Spin     `json:"spins"`
	Combos    map[string]*Select   `json:"combos"`
	Strings   map[string]*String   `json:"strings"`
	Filenames map[string]*FileName `json:"filenames"`
}

// NewOptMap 新しい OptMap を返す
func NewOptMap() *OptMap {
	return &OptMap{
		Buttons:   make(map[string]*Button),
		Checks:    make(map[string]*Check),
		Spins:     make(map[string]*Spin),
		Combos:    make(map[string]*Select),
		Strings:   make(map[string]*String),
		Filenames: make(map[string]*FileName),
	}
}

// Append OptMap に新しいオプションを追加する
func (om *OptMap) Append(o Option) {
	switch t := o.(type) {
	case *Button:
		om.Buttons[t.GetName()] = t
	case *Check:
		om.Checks[t.GetName()] = t
	case *Spin:
		om.Spins[t.GetName()] = t
	case *Select:
		om.Combos[t.GetName()] = t
	case *String:
		om.Strings[t.GetName()] = t
	case *FileName:
		om.Filenames[t.GetName()] = t
	default:
		panic(exception.UnknownOption)
	}
}

// Update OptMap にある Option の値を更新する
func (om *OptMap) Update(v UpdateOptionValue) (string, error) {
	var (
		opt Option
		ok  bool
	)
	switch v.Type {
	case "button":
		opt, ok = om.Buttons[v.Name]
	case "check":
		opt, ok = om.Checks[v.Name]
	case "spin":
		opt, ok = om.Spins[v.Name]
	case "select":
		opt, ok = om.Combos[v.Name]
	case "string":
		opt, ok = om.Strings[v.Name]
	case "filename":
		opt, ok = om.Filenames[v.Name]
	}

	log("SpecifiedOption", opt)
	if ok {
		s, e := opt.Update(v.Value)
		log("UpdatedOption", opt)
		return s, e
	}

	msg := fmt.Sprintf("OptionName %s was not found.", v.Name)
	logger.Use().Warn("OptMap_Update. Type was not valid", zap.String("msg", msg))
	return "", exception.UnknownOption.WithMsg(msg)
}

func log(key string, opt Option) {
	logger.Use().Debug("OptMap_Update", zap.Any(key, opt))
}
