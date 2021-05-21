package main

import "fmt"

type Student struct {
	Name string
}

type HasNamer interface {
	HasName() bool
}

func (s Student) HasName() bool {
	return s.Name != ""
}

func main() {
	lokesh := Student{"Enjoying"}

	fmt.Println(lokesh)
}
