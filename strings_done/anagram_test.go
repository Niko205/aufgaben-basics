package strings

import "fmt"

func ExampleIsAnagram() {
	fmt.Println(IsAnagram("ABC", "CBA"))
	fmt.Println(IsAnagram("ABC", "BCA"))
	fmt.Println(IsAnagram("Anna", "annA"))
	fmt.Println(IsAnagram("Anna", "Anna"))
	fmt.Println(IsAnagram("Ampel", "Lampe"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsAnagramIgnoreCase() {
	fmt.Println(IsAnagramIgnoreCase("ABC", "cba"))
	fmt.Println(IsAnagramIgnoreCase("ABC", "bca"))
	fmt.Println(IsAnagramIgnoreCase("Anna", "annA"))
	fmt.Println(IsAnagramIgnoreCase("Anna", "Anna"))
	fmt.Println(IsAnagramIgnoreCase("Ampel", "Lampe"))

	// Output:
	// true
	// true
	// true
	// true
	// true
}
