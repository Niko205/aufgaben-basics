package strings

import "strings"

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	sl := make([]string, 2*len(s))
	for i := 0; i < len(s); i++ {
		sl[i*2] = string(s[i])
		sl[i*2+1] = string(s[i])

	}

	return strings.Join(sl, "")
}
