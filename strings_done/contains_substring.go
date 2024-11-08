package strings

import "strings"

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	// TODO
	return strings.Contains(s, t)
}
