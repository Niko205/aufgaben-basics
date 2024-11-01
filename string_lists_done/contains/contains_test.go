package contains

import "fmt"

func ExampleContains() {
	fmt.Println(Contains([]string{"apple", "banana", "cherry"}, "banana"))
	fmt.Println(Contains([]string{"apple", "banana", "cherry"}, "orange"))
	fmt.Println(Contains([]string{}, "banana"))

	// Output:
	// true
	// false
	// false
}
