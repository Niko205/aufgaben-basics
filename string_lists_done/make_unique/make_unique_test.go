package make_unique

import "fmt"

func ExampleMakeUnique() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}

	s1 = MakeUnique(s1)
	fmt.Println(s1)

	// Output: [a b a2 c b2 a3]
}
