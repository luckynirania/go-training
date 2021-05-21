package hello

import (
	"fmt"
)

const myConstString = "golang"

func Test() {
	fmt.Printf("my-const-string: %s\n", myConstString)

	var status bool // uninitialized -> default (=false)
	fmt.Printf("status: %v\n", status)

	// := short notation: derives type from right-hand-side
	idx := 256
	fmt.Printf("idx: %d\n", idx)

	longString := `{
        "why": "Useful to embed json in source"
    }`
	fmt.Printf("my-long-string: %s\n", longString)
}
