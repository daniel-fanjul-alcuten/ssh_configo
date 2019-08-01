package sshconfigo

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseLine_Empty(t *testing.T) {
	for _, r := range []struct {
		input         string
		expectedLine  Line
		expectedError error
	}{
		// empty
		{``, Line{}, nil},
		// prefix
		{` `, Line{Prefix: String{0, " "}}, nil},
		// comment
		{`#foo`, Line{Comment: String{0, "#foo"}}, nil},
		// prefix and comment
		{` #foo`, Line{Prefix: String{0, " "}, Comment: String{1, "#foo"}}, nil},
		// prefix and comment with space
		{` # foo`, Line{Prefix: String{0, " "}, Comment: String{1, "# foo"}}, nil},
		// unexpected character
		{` @`, Line{Prefix: String{0, " "}},
			errors.New("Unexpected character '@' for remaining \"@\"")},
		// unexpected character with trailing text
		{` @ foo`, Line{Prefix: String{0, " "}},
			errors.New("Unexpected character '@' for remaining \"@ foo\"")},
		// keyword
		{`foo`, Line{Keyword: String{0, "foo"}}, nil},
	} {
		l, err := ParseLine(r.input)
		if v, e := fmt.Sprintf("%#v", l), fmt.Sprintf("%#v", r.expectedLine); v != e {
			t.Errorf(`For "%v": %v`, r.input, v)
		}
		if v, e := fmt.Sprintf("%#v", err), fmt.Sprintf("%#v", r.expectedError); v != e {
			t.Errorf(`For:"%v": %v`, r.input, v)
		}
	}
}
