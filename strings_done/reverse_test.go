package strings

import "fmt"

func ExampleReverse() {
	fmt.Println(Reverse("ABC"))
	fmt.Println(Reverse("Hallo"))
	fmt.Println(Reverse(""))

	// Output:
	// CBA
	// ollaH
	//
}

func ExampleIsReverse() {
	fmt.Println(IsReverse("ABC", "CBA"))
	fmt.Println(IsReverse("Hallo", "ollaH"))
	fmt.Println(IsReverse("", ""))
	fmt.Println(IsReverse("ABC", "DEF"))
	fmt.Println(IsReverse("ABC", "D"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("ABCBA"))
	fmt.Println(IsPalindrome("ANNA"))
	fmt.Println(IsPalindrome("Anna"))
	fmt.Println(IsPalindrome("atoyotasatoyota"))

	// Output:
	// true
	// true
	// false
	// true
}
