package jsonxml

import (
	"flag"
	"fmt"
	"log"
)

type Employee struct {
	Name   string `json:"name"	xml:"name"`
	Age    int    `json:"age"	xml:"age"`
	dob    string
	salary int
	Desg   Department `json: "dept"	xml: "dept"`
}

type Department struct {
	Name  string `json: "name"	xml: "dept"`
	level int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Demo() {
	D1 := Department{
		"Technical",
		10,
	}
	D2 := Department{
		"HR",
		12,
	}
	E1 := Employee{
		"Lokesh",
		20,
		"09-Dec-2000",
		200,
		D1,
	}
	E2 := Employee{
		"Raj",
		25,
		"25-Jan-1996",
		350,
		D2,
	}

	// Parsing Flags
	var e1, e2 string
	flag.StringVar(&e1, "e1", "", "Employee 1")
	flag.StringVar(&e2, "e2", "", "Employee 2")
	flag.Parse() // important, otherwise it doesn't work

	if e1 == "" || e2 == "" {
		flag.PrintDefaults()
		log.Fatal("Please provide both names")
	}

	// Writing Employees to JSON
	err := WriteJSON(E1, e1+".json")
	check(err)

	err = WriteJSON(E2, e2+".json")
	check(err)

	// Writing Employees to XML
	err = WriteXML(E1, e1+".xml")
	check(err)

	err = WriteXML(E2, e2+".xml")
	check(err)

	// Reading Employees from JSON
	E1fromJSON, err := ReadJSON(e1 + ".json")
	check(err)
	fmt.Println(E1fromJSON)

	E2fromJSON, err := ReadJSON(e2 + ".json")
	check(err)
	fmt.Println(E2fromJSON)

	// Reading Employees from XML
	E1fromXML, err := ReadXML(e1 + ".xml")
	check(err)
	fmt.Println(E1fromXML)

	E2fromXML, err := ReadXML(e2 + ".xml")
	check(err)
	fmt.Println(E2fromXML)

}
