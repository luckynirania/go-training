package loki

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func A2() {
	fmt.Println("Hello World")
	content, err := ioutil.ReadFile("../README.md")
	if err != nil {
		defer fmt.Println("Error bhai in opening", err)
	}
	longStr := strings.ToUpper(string(content))

	message := []byte(longStr)
	err = ioutil.WriteFile("../READ_CAP.md", message, 0644)
	if err != nil {
		defer fmt.Println("Error bhai in writing", err)
	}
}
