package strings

import "fmt"

func ExampleCheckParentheses() {
	fmt.Println(CheckParentheses("()"))
	fmt.Println(CheckParentheses("(())"))
	fmt.Println(CheckParentheses("()()"))
	fmt.Println(CheckParentheses("(a)b(ac)cb"))
	fmt.Println(CheckParentheses("())"))
	fmt.Println(CheckParentheses("(()"))
	fmt.Println(CheckParentheses("()("))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}
