package hello

import (
	"fmt"
)

func Loops() {
	// for loop

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// while-like for
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// iterate
	values := []string{"a", "b", "c"}
	for idx, value := range values {
		fmt.Printf("%d:%s\n", idx, value)
	}
}
