package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
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

func main() {
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
	err := writeJSON(E1, e1+".json")
	if err != nil {
		log.Fatal(err)
	}

	err = writeJSON(E2, e2+".json")
	if err != nil {
		log.Fatal((err))
	}

	// Writing Employees to XML
	err = writeXML(E1, e1+".xml")
	if err != nil {
		log.Fatal(err)
	}

	err = writeXML(E2, e2+".xml")
	if err != nil {
		log.Fatal((err))
	}

	// Reading Employees from JSON
	E1fromJSON, err := readJSON(e1 + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(E1fromJSON)

	E2fromJSON, err := readJSON(e2 + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(E2fromJSON)

	// Reading Employees from XML
	E1fromXML, err := readXML(e1 + ".xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(E1fromXML)

	E2fromXML, err := readXML(e2 + ".xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(E2fromXML)

}
func writeJSON(e Employee, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	employeeJSON, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	} else {
		if _, err := file.Write(employeeJSON); err != nil {
			return err
		}
		// fmt.Println(string(employeeJSON))
	}
	return nil
}

func readJSON(input string) (Employee, error) {
	f, err := os.Open(input)
	if err != nil {
		return Employee{}, err
	}
	defer f.Close()

	var employee Employee
	err = json.NewDecoder(f).Decode(&employee)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func writeXML(e Employee, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	employeeXML, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	} else {
		if _, err := file.Write(employeeXML); err != nil {
			return err
		}
		// fmt.Println(string(employeeXML))
	}
	return nil
}

func readXML(input string) (Employee, error) {
	f, err := os.Open(input)
	if err != nil {
		return Employee{}, err
	}
	defer f.Close()

	var employee Employee
	err = xml.NewDecoder(f).Decode(&employee)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}
