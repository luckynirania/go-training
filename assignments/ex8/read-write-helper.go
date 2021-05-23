package ex8

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

func WriteJSON(e Employee, output string) error {
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

func ReadJSON(input string) (Employee, error) {
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

func WriteXML(e Employee, output string) error {
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

func ReadXML(input string) (Employee, error) {
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
