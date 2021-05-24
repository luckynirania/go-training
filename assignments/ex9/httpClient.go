package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GetThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Client struct {
	Hostname string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (cl *Client) PutThing(thingUUID string, name string, value string) (*GetThingResponse, error) {
	values := map[string]string{"name": name, "value": value}
	json_data, err := json.Marshal(values)
	check(err)

	client := http.Client{}

	req, err := http.NewRequest("PUT", cl.Hostname+"/thing/"+thingUUID, bytes.NewBuffer(json_data))
	check(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	httpResponse, err := client.Do(req)
	check(err)

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error performing PUT thing: http-status %d", httpResponse.StatusCode)
	}

	var resp GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl *Client) CreateThing(name string, value string) (*GetThingResponse, error) {
	values := map[string]string{"name": name, "value": value}
	json_data, err := json.Marshal(values)
	check(err)

	client := http.Client{}

	httpResponse, err := client.Post(cl.Hostname+"/thing/new", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error posting thing: http-status %d", httpResponse.StatusCode)
	}

	var resp GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cl *Client) GetThingOnUUID(thingUUID string) (*GetThingResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/thing/%s", cl.Hostname, thingUUID))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	var resp GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func main() {
	client := Client{Hostname: "https://api-ldej-nl.el.r.appspot.com"}

	// GetThings
	resp, err := client.GetThingOnUUID("97fb41b1-0bdc-488a-9e04-b34eb70d919d")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", resp)
	fmt.Println()

	// Post a new thing
	// resp, err = client.CreateThing("Lokesh", "Posting here")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Printf("%v", resp)
	// fmt.Println()

	// PUT a Thing
	// resp, err = client.PutThing("97fb41b1-0bdc-488a-9e04-b34eb70d919d", "Lokesh", "Putting here")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Printf("%v", resp)
	// fmt.Println()
}
