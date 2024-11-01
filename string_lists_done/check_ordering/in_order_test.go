package check_ordering

import "fmt"

func ExampleCheckOrdering() {
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "a", "f"))
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "f", "a"))
	fmt.Println(CheckOrdering([]string{"a", "b", "c", "d", "e", "f"}, "f", "g"))
	fmt.Println(CheckOrdering([]string{}, "a", "b"))

	// Output:
	// true
	// false
	// false
	// false
}
