package httpTraining

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
