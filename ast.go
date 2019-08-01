package sshconfigo

import "fmt"

// String is a string inside a line. It does not contain new line endings.
type String struct {
	// Pos is the position of the first character in the line.
	Pos int
	// Text is the contents of the string.
	Text string
}

// StandardKeywords contains all standard keywords.
var StandardKeywords = map[string]struct{}{
	"Host":          struct{}{},
	"AddressFamily": struct{}{},
	// TODO add all standard keywords
}

// IsStandardKeyword returns true if the keyword is contained in
// StandardKeywords.
func IsStandardKeyword(keyword string) bool {
	_, ok := StandardKeywords[keyword]
	return ok
}

// Line is a line inside a ssh_config file. An empty line is represented by a
// Line that does not contain a value for the keyword.
type Line struct {
	// Pos is the line number.
	Pos int
	// Prefix contains the spaces and tabs before the keyword.
	Prefix String
	// Keyword contains the keyword. It's case insensitive.
	Keyword String
	// Assignment is the separator between the keyword and the arguments. It is
	// usually a space, it may contain the '=' character.
	Assignment String
	// Arguments contains the arguments. It's case sensitive.
	Arguments String
	// Postfix contains the spaces and tabs after the arguments.
	Postfix String
	// Comment is the optional string starting the '#' character.
	Comment String
	// NewLine contains the new line endings as they are found in the file.
	NewLine String
}

// IsBlank returns true when the Line does not contain a value for the keyword,
// that is when it is empty or there are only spaces and tabs and / or there is
// only a comment.
func (l Line) IsBlank() bool {
	return l.Keyword.Text == ""
}

// ArgumentsAsString returns the arguments after the escape characters have
// been processed.
func (l Line) ArgumentsAsString() string {
	// TODO Implement Arguments
	return ""
}

// ArgumentsAsSlice returns the arguments after the escape characters have been
// processed.
func (l Line) ArgumentsAsSlice() []string {
	// TODO Implement ArgumentsAsSlice
	return nil
}

// Section represents a slice of Lines starting from the Line with the Host
// specification.
type Section struct {
	// Host is the Line with the Host specification.
	Host Line
	// Lines is the slice of Lines belonging to the Section. It contains all the
	// Lines until the next Section or end of file.
	Lines []Line
}

// ValueAsString returns the value of the given keyword of the given Section.
func (s Section) ValueAsString(keyword string) string {
	// TODO Implement Value
	return ""
}

// ValueAsSlice returns the value of the given keyword of the given Section.
func (s Section) ValueAsSlice(keyword string) []string {
	// TODO Implement Value
	return nil
}

// AST is the data structure that contains the AST of a ssh_config file.
type AST interface {
	// NumLines returns the number of lines in the AST.
	NumLines() int
	// Line returns the Line at the position i.
	Line(int) Line
	// SetLine creates new empty Lines if needed until the given Line is added or
	// replaces at the given position.
	SetLine(Line)
	// DeleteLine deletes the line at the given position.
	DeleteLine(int)
	// Section matches a hostname against the patterns of the Host specifications
	// and returns the first matching Section.
	Section(string) Section
	// Sections returns all Sections in the AST.
	Sections() []Section
}

// ParseLine parses a single line.
func ParseLine(s string) (l Line, err error) {
	i, r := 0, []rune(s)
	accept := func(f func(r rune) bool) (s String) {
		if i < len(r) && f(r[i]) {
			s.Pos = i
			s.Text += string(r[i])
			i++
			for i < len(r) && f(r[i]) {
				s.Text += string(r[i])
				i++
			}
		}
		return
	}
	l.Prefix = accept(func(r rune) bool {
		return r == ' '
	})
	l.Keyword = accept(func(r rune) bool {
		return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
	})
	// TODO Parse l.Assignment
	// TODO Parse l.Arguments
	// TODO Parse l.Postfix
	if i < len(r) && r[i] == '#' {
		l.Comment = accept(func(r rune) bool {
			return true
		})
	}
	// TODO Parse l.NewLine
	if i < len(r) {
		err = fmt.Errorf("Unexpected character '%v' for remaining \"%v\"",
			string(r[i]), string(r[i:]))
		return
	}
	return
}
