package customAPI

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func PUT(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)["uuid"]

	// Check if UUID exists
	if _, err := os.Stat("./data/" + u + ".json"); os.IsNotExist(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Open the UUID file, update the values, write into file
	var thingResponse ThingResponse
	f, _ := os.Open("./data/" + u + ".json")

	_ = json.NewDecoder(f).Decode(&thingResponse)
	f.Close()

	_ = json.NewDecoder(r.Body).Decode(&thingResponse)

	thingResponse.Updated = time.Now()

	file, _ := os.Create("./data/" + u + ".json")
	defer file.Close()

	dataJSON, _ := json.MarshalIndent(thingResponse, "", "  ")
	file.Write(dataJSON)

	// Send response back

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
