package strings

import "fmt"

func ExampleDuplicateChars() {
	fmt.Println(DuplicateChars("ABC"))
	fmt.Println(DuplicateChars("Hallo"))
	fmt.Println(DuplicateChars(""))

	// Output:
	// AABBCC
	// HHaalllloo
	//
}
