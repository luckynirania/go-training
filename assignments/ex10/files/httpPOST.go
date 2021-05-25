package customAPI

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func POST(w http.ResponseWriter, r *http.Request) {
	u := uuid.New()

	var thingResponse ThingResponse
	thingResponse.UUID = u.String()
	currentTime := time.Now()
	thingResponse.Created = currentTime
	thingResponse.Updated = currentTime

	// ResponseData, _ := ioutil.ReadAll(r.Body)

	// fmt.Println(thingResponse)
	_ = json.NewDecoder(r.Body).Decode(&thingResponse)
	// fmt.Println(thingResponse)

	// Write in Database
	file, _ := os.Create("./data/" + u.String() + ".json")
	defer file.Close()

	dataJSON, _ := json.MarshalIndent(thingResponse, "", "  ")
	file.Write(dataJSON)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
