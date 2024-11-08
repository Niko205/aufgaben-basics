package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	s := ""
	l := len(s1)
	if len(s2) > len(s1) {
		l = len(s2)
	}

	for i := 0; i < l; i++ {
		if i < len(s1) {
			s = s + string(s1[i])
		}
		if i < len(s2) {
			s = s + string(s2[i])
		}
	}
	return s
}
