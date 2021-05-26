package serverStruct

import (
	"encoding/json"
	"net/http"
)

func (s *server) PostThing(w http.ResponseWriter, r *http.Request) {
	var thingNew ThingResponse

	_ = json.NewDecoder(r.Body).Decode(&thingNew)

	thingResponse, err := s.db.CreateThing(thingNew.Name, thingNew.Value)

	temp := ThingResponse{
		UUID:    thingResponse.UUID,
		Name:    thingResponse.Name,
		Value:   thingResponse.Value,
		Created: thingResponse.Created,
		Updated: thingResponse.Updated,
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(temp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
