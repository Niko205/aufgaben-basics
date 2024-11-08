package strings

import "fmt"

func ExampleConcatN() {
	fmt.Println(ConcatN("A", "", 5))
	fmt.Println(ConcatN("A", "B", 5))
	fmt.Println(ConcatN("AB", "C", 5))

	// Output:
	// AAAAA
	// ABABABABA
	// ABCABCABCABCAB
}
