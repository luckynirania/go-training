package customAPI

import (
	"encoding/json"
	"net/http"
	"os"

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

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
