package httpTraining

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
