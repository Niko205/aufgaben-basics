package strings

import "strings"

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	sl := make([]string, len(s))
	z := 0
	for i := len(s) - 1; i >= 0; i-- {
		sl[z] = string(s[i])
		z++
	}
	return strings.Join(sl, "")
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {

	return s1 == Reverse(s2)

}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {

	return IsReverse(s, s)

}
