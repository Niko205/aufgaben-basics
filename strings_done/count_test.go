package strings

import "fmt"

func ExampleCountA() {
	fmt.Println(CountA("ABC"))
	fmt.Println(CountA("ABCABC"))
	fmt.Println(CountA("AAAABBBA"))
	fmt.Println(CountA("DEFGH"))

	// Output:
	// 1
	// 2
	// 5
	// 0
}

func ExampleCountChar() {
	fmt.Println(CountChar("ABC", 'A'))
	fmt.Println(CountChar("ABCABC", 'A'))
	fmt.Println(CountChar("AAAABBBA", 'A'))
	fmt.Println(CountChar("DEFGH", 'A'))

	// Output:
	// 1
	// 2
	// 5
	// 0
}
