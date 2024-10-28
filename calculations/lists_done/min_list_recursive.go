package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[1] { // Vergleicht die ersten beiden Zahlen
		nums[1] = nums[0] // Setzt die kleinere Zahl an die zweite Stelle
	}
	return MinListRecursive(nums[1:]) // Entfernt das erste Element und ruft die Funktion erneut auf
	// Wiederholt die Funktion so lange bis nur noch ein Element in der Liste ist

}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
