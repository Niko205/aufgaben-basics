package strings

import "strings"

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	sl := strings.Split(s, "")
	open := false
	for i := 0; i < len(sl); i++ {
		if sl[i] == "(" {
			open = true
			for j := i + 1; j < len(sl); j++ {
				if sl[j] == ")" {
					open = false
					sl[j] = "a"
					break
				}
			}
			if open {
				return !open
			}
		} else if sl[i] == ")" {
			return false
		}
	}

	return true
}
