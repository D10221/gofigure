package inutil

import "testing"

var Value = ""

func Testinit(t *testing.T) {
	if Value == "" {
		t.Error("it doesn't wokr like this")
	}
}

func init() {
	Value = "!"
}

