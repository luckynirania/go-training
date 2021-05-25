package customAPI

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func GET(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)["uuid"]

	// Check if UUID exists
	if _, err := os.Stat("./data/" + u + ".json"); os.IsNotExist(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var thingResponse ThingResponse
	f, _ := os.Open("./data/" + u + ".json")

	_ = json.NewDecoder(f).Decode(&thingResponse)
	f.Close()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
