package strings

import "strings"

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	b := false
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if s1[i] == s2[j] {
				b = true
			}
		}
		if !b {
			return b
		}
		b = false
	}
	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	b := false
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if s1[i] == s2[j] {
				b = true
			}
		}
		if !b {
			return b
		}
		b = false
	}
	return true
}