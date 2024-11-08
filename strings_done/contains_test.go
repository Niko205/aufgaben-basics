package strings

import "fmt"

func ExampleContains() {
	fmt.Println(Contains("Hallo", 'a'))
	fmt.Println(Contains("Hallo", 'b'))
	fmt.Println(Contains("Hallo", 'H'))
	fmt.Println(Contains("Hallo", 'h'))

	// Output:
	// true
	// false
	// true
	// false
}
