package strings

import "strings"

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// TODO
	return strings.Contains(s, string(c))
}
