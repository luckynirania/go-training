package main

import "fmt"

func main() {
	data := map[string][]string{
		"Julia":  {"cricket", "drawing"},
		"Sophie": {"drawing"},
		"Mila":   {"drawing"},
		"Emma":   {"tennis", "kabaddi"},
		"Neha":   {"running"},
		"Abhi":   {"photography", "cricket"},
		"Noor":   {"cricket"},
		"Elin":   {"hockey"},
		"Sara":   {"cricket", "kabaddi"},
		"Yara":   {"tennis"},
	}

	HobbyGroup := map[string][]string{}

	for person, hobbies := range data {
		for _, hobby := range hobbies {
			HobbyGroup[hobby] = append(HobbyGroup[hobby], person)
		}
	}

	fmt.Println(data)
	fmt.Println()
	fmt.Println(HobbyGroup)
}
