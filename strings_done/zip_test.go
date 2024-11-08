package strings

import "fmt"

func ExampleZip() {
	fmt.Println(Zip("AAA", "BBB"))
	fmt.Println(Zip("ABABCD", "CDCDAB"))

	// Output:
	// ABABAB
	// ACBDACBDCADB
}

func ExampleZip_different_length() {
	fmt.Println(Zip("A", "BBB"))
	fmt.Println(Zip("AAA", "B"))

	// Output:
	// ABBB
	// ABAA
}
