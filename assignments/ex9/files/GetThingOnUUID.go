package httpTraining

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
