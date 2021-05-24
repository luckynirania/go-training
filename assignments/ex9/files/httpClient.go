package httpTraining

import (
	"log"
	"time"
)

type GetThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Client struct {
	Hostname string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
