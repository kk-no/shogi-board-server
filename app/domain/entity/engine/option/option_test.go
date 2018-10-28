// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"bytes"
	"testing"
)

func TestButton_GetName(t *testing.T) {
	cases := []struct {
		in   Button
		want []byte
	}{
		{Button{[]byte("btn-name")}, []byte("btn-name")},
		{Button{[]byte("")}, []byte("")},
		{Button{[]byte(" ")}, []byte(" ")},
		{Button{[]byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}
func TestButton_Usi(t *testing.T) {
	cases := []struct {
		in   Button
		want []byte
	}{
		{Button{[]byte("btn-name")}, []byte("setoption name btn-name")},
		{Button{[]byte("")}, []byte("setoption name ")},
		{Button{[]byte(" ")}, []byte("setoption name  ")},
		{Button{[]byte("%\n|t\t")}, []byte("setoption name %\n|t\t")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_GetName(t *testing.T) {

}

func TestCheck_Usi(t *testing.T) {

}

func TestFileName_GetName(t *testing.T) {

}

func TestFileName_Usi(t *testing.T) {

}

func TestSelect_GetName(t *testing.T) {

}

func TestSelect_Usi(t *testing.T) {

}

func TestSpin_GetName(t *testing.T) {

}

func TestSpin_Usi(t *testing.T) {

}

func TestString_GetName(t *testing.T) {

}

func TestString_Usi(t *testing.T) {

}

func getNameTestHelper(t *testing.T, i int, o Option, want []byte) {
	t.Helper()
	if !bytes.Equal(o.GetName(), want) {
		t.Errorf(`Option.GetName was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(o.GetName()), string(want))
	}
}

func usiTestHelper(t *testing.T, i int, o Option, want []byte) {
	t.Helper()
	if bytes.Equal(o.Usi(), want) {
		t.Errorf(`Option.Usi was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(o.Usi()), string(want))
	}
}
