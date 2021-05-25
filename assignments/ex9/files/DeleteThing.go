package httpTraining

import (
	"fmt"
	"net/http"
)

func (cl *Client) DeleteThing(thingUUID string) error {
	client := http.Client{}
	// todo := GetThingResponse{"", "", "", time.Now(), time.Now()}
	// jsonReq, _ := json.Marshal(todo)
	req, err := http.NewRequest("DELETE", cl.Hostname+"/thing/"+thingUUID, nil)
	check(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	httpResponse, err := client.Do(req)
	check(err)

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return fmt.Errorf("error Deleting thing: http-status %d", httpResponse.StatusCode)
	}
	return nil
}
