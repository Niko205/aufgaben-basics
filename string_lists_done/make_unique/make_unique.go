package make_unique

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
func MakeUnique(strings []string) []string {
	a := '1'

	for i := 0; i < len(strings); i++ {
		for j := 0; j < i; j++ {
			if strings[i][0] == strings[j][0] {
				a++
			}
		}
		if a > '1' {
			strings[i] = strings[i] + string(a)
		}
		a = '1'
	}
	return strings
}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
