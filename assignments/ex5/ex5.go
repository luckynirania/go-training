package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Datastorer interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	Remove(key string) error
}

type InMemDatabase struct {
	PersonHobbyPair map[string][]string
}

func (i *InMemDatabase) Put(key string, value interface{}) error {
	token, ok := value.([]string)
	if !ok {
		return errors.New("invalid format")
	}
	i.PersonHobbyPair[key] = token
	return nil
}

func (i *InMemDatabase) Get(key string) (interface{}, bool, error) {
	token, ok := i.PersonHobbyPair[key]
	if !ok {
		return "", false, errors.New("invalid key")
	}
	return token, true, nil
}

func (i *InMemDatabase) Remove(key string) error {
	_, ok := i.PersonHobbyPair[key]
	if !ok {
		return errors.New("invalid key")
	}
	delete(i.PersonHobbyPair, key)
	return nil
}

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

	DataBase := InMemDatabase{PersonHobbyPair: data}

	// Put Test
	ok := DataBase.Put("Lokesh", []string{"data"})
	fmt.Println("Error Status in Put :- ", ok)

	ok = DataBase.Put("Rahul", "data")
	fmt.Println("Error Status in Put :- ", ok)
	fmt.Println()

	// Get Test
	hobby, present, err := DataBase.Get("Julia")
	if present {
		fmt.Println("Hobbies :- ", hobby)
	} else {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(hobby))

	hobby, present, err = DataBase.Get("Rahul")
	if present {
		fmt.Println("Hobbies :- ", hobby)
	} else {
		fmt.Println(err)
	}
	fmt.Println()

	// Remove Test
	ok = DataBase.Remove("Julia")
	if ok == nil {
		fmt.Println("Successfully removed")
	} else {
		fmt.Println(ok)
	}

	ok = DataBase.Remove("Julia")
	if ok == nil {
		fmt.Println("Successfully removed")
	} else {
		fmt.Println(ok)
	}
}
