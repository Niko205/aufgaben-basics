package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	x := s
	for i := 1; i < n; i++ {
		x = x + sep + s
	}
	return x
}
