package sanitizer

import (
	"fmt"
	"reflect"
	"strings"
)

//
//
// Sanitize a map
//
//
func FixMap(shouldFix bool, m map[string]interface{}, blacklist []string) map[string]interface{} {
	var n = make(map[string]interface{})
	for key, value := range m {
		if (reflect.TypeOf(value).String() == "string") &&
			!HasString(blacklist, key) {
			if shouldFix {
				n[key] = Clean(fmt.Sprintf("%v", value))
			} else {
				n[key] = Dirty(fmt.Sprintf("%v", value))
			}
		} else {
			// Every element must be copied, even if it's blacklisted and/or not a string
			n[key] = value
		}
	}
	return n
}

//
//
// Find out if an array contains an element
//
//
func HasString(ar []string, s string) bool {
	for _, value := range ar {
		if value == s {
			return true
		}
	}
	return false
}

//
//
// Sanitize a string
//
//
func Clean(s string) string {
	if s != "" {
		// Allowed so that they are NOT self-cleaned
		//s = strings.ReplaceAll(s, "&", "&amp;")
		//s = strings.ReplaceAll(s, ";", "&semi;")
		//
		s = strings.ReplaceAll(s, "!", "&excl;")
		s = strings.ReplaceAll(s, "\"", "&quot;")
		s = strings.ReplaceAll(s, "#", "&num;")
		s = strings.ReplaceAll(s, "$", "&dollar;")
		s = strings.ReplaceAll(s, "%", "&percnt;")
		s = strings.ReplaceAll(s, "'", "&apos;")
		s = strings.ReplaceAll(s, "(", "&lpar;")
		s = strings.ReplaceAll(s, ")", "&rpar;")
		s = strings.ReplaceAll(s, "*", "&ast;")
		s = strings.ReplaceAll(s, "+", "&plus;")
		s = strings.ReplaceAll(s, ",", "&comma;")
		s = strings.ReplaceAll(s, "-", "&hyphen;")
		s = strings.ReplaceAll(s, ".", "&period;")
		s = strings.ReplaceAll(s, "/", "&sol;")
		s = strings.ReplaceAll(s, ":", "&colon;")
		s = strings.ReplaceAll(s, "<", "&lt;")
		s = strings.ReplaceAll(s, "=", "&equals;")
		s = strings.ReplaceAll(s, ">", "&gt;")
		s = strings.ReplaceAll(s, "?", "&quest;")
		s = strings.ReplaceAll(s, "@", "&commat;")
		s = strings.ReplaceAll(s, "[", "&lsqb;")
		s = strings.ReplaceAll(s, "\\", "&bsol;")
		s = strings.ReplaceAll(s, "]", "&rsqb;")
		s = strings.ReplaceAll(s, "^", "&hat;")
		s = strings.ReplaceAll(s, "_", "&lowbar;")
		s = strings.ReplaceAll(s, "`", "&grave;")
		s = strings.ReplaceAll(s, "{", "&lcub;")
		s = strings.ReplaceAll(s, "|", "&vert;")
		s = strings.ReplaceAll(s, "}", "&rcub;")
		s = strings.ReplaceAll(s, "~", "&tilde;")
		s = strings.ReplaceAll(s, "--", "")
	}
	return s
}

//
//
// Unsanitize a string
//
//
func Dirty(s string) string {
	if s != "" {
		s = strings.ReplaceAll(s, "&excl;", "!")
		s = strings.ReplaceAll(s, "\"", "&quot;")
		s = strings.ReplaceAll(s, "&num;", "#")
		s = strings.ReplaceAll(s, "&dollar;", "$")
		s = strings.ReplaceAll(s, "&percnt;", "%")
		s = strings.ReplaceAll(s, "&apos;", "'")
		s = strings.ReplaceAll(s, "&lpar;", "(")
		s = strings.ReplaceAll(s, "&rpar;", ")")
		s = strings.ReplaceAll(s, "&ast;", "*")
		s = strings.ReplaceAll(s, "&plus;", "+")
		s = strings.ReplaceAll(s, "&comma;", ",")
		s = strings.ReplaceAll(s, "&hyphen;", "-")
		s = strings.ReplaceAll(s, "&period;", ".")
		s = strings.ReplaceAll(s, "&sol;", "/")
		s = strings.ReplaceAll(s, "&colon;", ":")
		s = strings.ReplaceAll(s, "&lt;", "<")
		s = strings.ReplaceAll(s, "&equals;", "=")
		s = strings.ReplaceAll(s, "&gt;", ">")
		s = strings.ReplaceAll(s, "&quest;", "?")
		s = strings.ReplaceAll(s, "&commat;", "@")
		s = strings.ReplaceAll(s, "&lsqb;", "[")
		s = strings.ReplaceAll(s, "&bsol;", "\\")
		s = strings.ReplaceAll(s, "&rsqb;", "]")
		s = strings.ReplaceAll(s, "&hat;", "^")
		s = strings.ReplaceAll(s, "&lowbar;", "_")
		s = strings.ReplaceAll(s, "&grave;", "`")
		s = strings.ReplaceAll(s, "&lcub;", "{")
		s = strings.ReplaceAll(s, "&vert;", "|")
		s = strings.ReplaceAll(s, "&rcub;", "}")
		s = strings.ReplaceAll(s, "&tilde;", "~")
		//s = strings.ReplaceAll(s, "--", "")
		// Allowed so that they are NOT self-duplicated
		//s = strings.ReplaceAll(s, "&semi;", ";")
		//s = strings.ReplaceAll(s, "&amp;", "&")
	}
	return s
}
