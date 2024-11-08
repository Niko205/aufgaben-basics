package strings

import "fmt"

func ExamplePositionOf() {
	fmt.Println(PositionOf("Hallo", 'H'))
	fmt.Println(PositionOf("Hallo", 'l'))
	fmt.Println(PositionOf("Hallo", 'o'))
	fmt.Println(PositionOf("Hallo", 'x'))
	fmt.Println(PositionOf("", 'a'))

	// Output:
	// 0
	// 2
	// 4
	// 5
	// 0
}
