package customAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func GETAll(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("./data")
	var list []ThingResponse

	var thingResponse ThingResponse
	for _, file := range files {
		f, _ := os.Open("./data/" + file.Name())

		_ = json.NewDecoder(f).Decode(&thingResponse)
		list = append(list, thingResponse)

		f.Close()
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
