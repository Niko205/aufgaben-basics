package triangles

import "fmt"

func ExampleDrawTriangle() {
	DrawTriangle(4, " ", "#")
	fmt.Println()
	DrawTriangle(6, "A", "B")
	// Output:
	// #
	// ##
	// # #
	// ####
	//
	// B
	// BB
	// BAB
	// BAAB
	// BAAAB
	// BBBBBB
}
