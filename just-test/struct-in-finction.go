package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func something() []Student {
	v1 := Student{
		Name: "Lokesh",
		Age:  20,
	}
	v2 := Student{
		Name: "Mohit",
		Age:  22,
	}

	ret := []Student{
		v1, v2,
	}
	return ret
}

func main() {
	lokesh := something()
	fmt.Println(lokesh)
}
