package count

import "fmt"

func ExampleCount() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}

	fmt.Println(Count(s1, "a"))
	fmt.Println(Count(s1, "b"))
	fmt.Println(Count(s1, "c"))

	// Output:
	// 3
	// 2
	// 1
}
