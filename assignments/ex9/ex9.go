package main

import (
	"fmt"
	"log"
	httpTraining "lokesh/assignments/ex9/files"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client := httpTraining.Client{Hostname: "https://api-ldej-nl.el.r.appspot.com"}

	// GetALl
	resp, err := client.GetAll()
	if err != nil {
		check(err)
	}
	for key, value := range *resp {
		fmt.Println(key, value)
	}
	fmt.Println()
	// ---------------------------------------------------------------------------------------------

	// Delete
	// err = client.DeleteThing("12db965a-5746-40f1-a3b0-5763ed9269d4")
	// if err != nil {
	// 	check(err)
	// }
	// fmt.Println("200 OK")
	// fmt.Println()
	// // ---------------------------------------------------------------------------------------------

	// // GetThing usung UUID
	// respi, err := client.GetThingOnUUID("97fb41b1-0bdc-488a-9e04-b34eb70d919d")
	// if err != nil {
	// 	check(err)
	// }
	// fmt.Printf("%v", respi)
	// fmt.Println()
	// // ---------------------------------------------------------------------------------------------

	// // Post a new thing
	// respi, err = client.CreateThing("Lokesh", "Posting here")
	// if err != nil {
	// 	check(err)
	// }
	// fmt.Printf("%v", respi)
	// fmt.Println()
	// // // ---------------------------------------------------------------------------------------------

	// // // PUT a Thing
	// respi, err = client.PutThing("97fb41b1-0bdc-488a-9e04-b34eb70d919d", "Lokesh", "Putting here")
	// if err != nil {
	// 	check(err)
	// }
	// fmt.Printf("%v", respi)
	// fmt.Println()
	// ---------------------------------------------------------------------------------------------
}
