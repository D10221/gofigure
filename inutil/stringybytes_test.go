package inutil

import (
	"testing"
	"unicode/utf8"

	"fmt"
)

func Test_Strings(t *testing.T) {
	// ⌘:= '⌘' == \u2318
	s := "Hello\u2318"

	t.Logf("byte[x] in []string")
	for i := 0; i < len(s); i++ {
		t.Logf("%x", s[i])
	}

	t.Logf("for ... range []escaped")
	for i, value := range s {
		t.Log("%d: %#U", i, value)
	}

	t.Logf("byte")
	if s[5:] != "⌘" {
		t.Errorf("byte? %q", s[5:])
	}

	t.Logf("for ... range\n")
	const nihongo = "日本2 語 1"
	for index, runeValue := range nihongo {
		t.Logf("%#U starts at byte position %d\n", runeValue, index)
	}
	t.Logf("DecodeRuneInString\n")
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		t.Logf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	a := fmt.Sprintf("%q", "\xd8\xb9")
	b := fmt.Sprintf("%q", "\u0639")
	if b != a {
		t.Error("!=", a, b)
	}

}