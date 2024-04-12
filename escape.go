package gojsonpath

import (
	"strings"
	"unicode"
)

const acceptableChars = "_*$.,@!~#%^"

// MakeKeyString quotes or escapes the input string if necessary to be
// used in a path expression
func MakeKeyString(input string) string {
	// If input has all digits, letters, and _, no need for escape or
	// quote
	bad := false
	for i, r := range input {
		if i == 0 && unicode.IsDigit(r) {
			bad = true
			break
		}
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && strings.IndexRune(acceptableChars, r) == -1 {
			bad = true
			break
		}
	}
	if !bad {
		return input
	}

	// Does it have double-quote?
	if strings.IndexRune(input, '"') != -1 {
		if strings.IndexRune(input, '\'') != -1 {
			out := make([]rune, 0, len(input)+2)
			out = append(out, '"')
			for _, r := range input {
				if r == '"' {
					out = append(out, '\\', '"')
				} else {
					out = append(out, r)
				}
			}
			out = append(out, '"')
			return string(out)
		}
		return "'" + input + "'"
	}
	return `"` + input + `"`
}

// Joins strings to build a path, escaping them as necessary
func Join(parts ...string) string {
	bldr := strings.Builder{}
	for _, part := range parts {
		bldr.WriteRune('/')
		bldr.WriteString(MakeKeyString(part))
	}
	ret := bldr.String()
	if len(ret) == 0 {
		return "/"
	}
	return ret
}
