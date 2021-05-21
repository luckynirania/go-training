package main

import "fmt"

type address struct {
	Street   string
	Locality string
	State    string
	Pin      int
}

type Student struct {
	Name    string
	Roll    string
	Address address
	Class   int
	Section string
	Courses []course
}

type Teacher struct {
	Name    string
	ID      string
	Address address
	Subject string
}

type course struct {
	Code        string
	Name        string
	Books       []string
	Instructors []Teacher
}

func main() {
	fmt.Println("Start")
	a1 := address{
		Street:   "Piprali",
		Locality: "Dev Gen Store",
		State:    "Rajasthan",
		Pin:      335568,
	}
	a2 := address{
		Street:   "Guru",
		Locality: "Saheb",
		State:    "Rajasthan",
		Pin:      335968,
	}
	a3 := address{
		Street:   "Ram",
		Locality: "Light Heavy",
		State:    "Delhi",
		Pin:      112584,
	}
	a4 := address{
		Street:   "Shyam",
		Locality: "Dings",
		State:    "Haryana",
		Pin:      225698,
	}

	t1 := Teacher{
		"Laurence",
		"985EY",
		a1,
		"CS",
	}

	t2 := Teacher{
		"Glory",
		"565EY",
		a4,
		"HR",
	}

	c1 := course{
		"CS101",
		"GoLang",
		[]string{"Book 1", "Book 2"},
		[]Teacher{t1, t2},
	}

	c2 := course{
		"Hr101",
		"Management",
		[]string{"Book 2", "Book 3", "Book 4"},
		[]Teacher{t2},
	}

	s1 := Student{
		"Lokesh",
		"170010009",
		a2,
		12,
		"E",
		[]course{c1, c2},
	}

	s2 := Student{
		"Rahul",
		"170010089",
		a3,
		12,
		"F",
		[]course{c2},
	}

	fmt.Println(s1)
	fmt.Println()
	fmt.Println(s2)
}
