package serverStruct

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ldej/go-training/examples/db"
)

func (s *server) GetThing(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	thing, err := s.db.GetThing(uuid)
	if err == db.ErrThingNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	thingResponse := ThingResponse{
		UUID:    thing.UUID,
		Name:    thing.Name,
		Value:   thing.Value,
		Created: thing.Created,
		Updated: thing.Updated,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
