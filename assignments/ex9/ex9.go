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

	// GetThings
	resp, err := client.GetThingOnUUID("97fb41b1-0bdc-488a-9e04-b34eb70d919d")
	if err != nil {
		check(err)
	}
	fmt.Printf("%v", resp)
	fmt.Println()
	// ---------------------------------------------------------------------------------------------

	// Post a new thing
	resp, err = client.CreateThing("Lokesh", "Posting here")
	if err != nil {
		check(err)
	}
	fmt.Printf("%v", resp)
	fmt.Println()
	// ---------------------------------------------------------------------------------------------

	// PUT a Thing
	resp, err = client.PutThing("97fb41b1-0bdc-488a-9e04-b34eb70d919d", "Lokesh", "Putting here")
	if err != nil {
		check(err)
	}
	fmt.Printf("%v", resp)
	fmt.Println()
	// ---------------------------------------------------------------------------------------------
}
