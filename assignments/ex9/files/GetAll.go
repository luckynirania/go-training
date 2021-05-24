package httpTraining

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (cl *Client) GetAll() (*[]GetThingResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/thing", cl.Hostname))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	ResponseData, _ := ioutil.ReadAll(httpResponse.Body)
	var ResponseString map[string]interface{}
	json.Unmarshal(ResponseData, &ResponseString)
	var Things []GetThingResponse

	loki := ResponseString["things"]

	for _, value := range loki.([]interface{}) {
		var resp GetThingResponse
		lok, _ := json.Marshal(value)
		err = json.NewDecoder(bytes.NewReader(lok)).Decode(&resp)
		if err != nil {
			return nil, err
		}
		Things = append(Things, resp)
	}
	return &Things, nil
}
