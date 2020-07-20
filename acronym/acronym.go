// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s2 := ""
	s = strings.ReplaceAll(s, " ", " ")
	//here it is advisible to use replace all rather than using replace
	//because if we use replace then wee have to mention an integer "n" which indicates how many times the character should be replaced
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ",", " ")
	words := strings.Fields(s)
	//it is advisible to use strings.Fields than strings.Split
	for _, v := range words {
		s2 += string(v[0])
	}

	return strings.ToUpper(s2)
}
