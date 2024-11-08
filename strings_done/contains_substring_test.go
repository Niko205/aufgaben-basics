package strings

import "fmt"

func ExampleContainsSubstring() {
	fmt.Println(ContainsSubstring("Hallo", "Ha"))
	fmt.Println(ContainsSubstring("Hallo", "llo"))
	fmt.Println(ContainsSubstring("Hallo", "Ho"))

	// Output:
	// true
	// true
	// false
}
