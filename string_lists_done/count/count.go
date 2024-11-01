package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	a := 0
	for i := 0; i < len(strings); i++ {
		if strings[i] == search {
			a++
		}
	}
	return a
}
