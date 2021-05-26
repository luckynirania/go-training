package serverStruct

import (
	"encoding/json"
	"net/http"

	"github.com/ldej/go-training/examples/db"
)

func (s *server) GetAllThing(w http.ResponseWriter, r *http.Request) {
	things, _, err := s.db.ListThings(0, s.db.GetSize())
	if err == db.ErrThingNotFound {
		// fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(things)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
