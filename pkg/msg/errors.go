// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msg

import "fmt"

type Error struct {
	Code, Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s", e.Code, e.Msg)
}

func (e *Error) WithMsg(m string) *Error {
	return &Error{e.Code, m}
}

var (
	// http
	NotFound          = &Error{Code: "NotFound"}
	MethodNotAllowed  = &Error{Code: "MethodNotAllowed"}
	FailedToReadBody  = &Error{Code: "FailedToReadBody"}
	FailedToParseBody = &Error{Code: "FailedToParseBody"}

	// engine
	FailedToStart          = &Error{Code: "FailedToStart"}
	FailedToShutdown       = &Error{Code: "FailedToShutdown"}
	EngineIsNotRunning     = &Error{Code: "EngineIsNotRunning"}
	EngineIsAlreadyRunning = &Error{Code: "EngineIsAlreadyRunning"}
	InvalidIdSyntax        = &Error{Code: "InvalidIdSyntax"}
	InvalidOptionSyntax    = &Error{Code: "InvalidOptionSyntax"}
	UnknownOption          = &Error{Code: "UnknownOption"}
	UnknownOptionType      = &Error{Code: "UnknownOptionType"}
)