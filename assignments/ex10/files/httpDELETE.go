package customAPI

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func DELETE(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)["uuid"]

	// Check and Remove
	e := os.Remove("./data/" + u + ".json")
	if e != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	thingResponse := ThingResponse{"", "", "", time.Now(), time.Now()}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
