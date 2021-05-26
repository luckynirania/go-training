package serverStruct

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	err := s.db.DeleteThing(uuid)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
