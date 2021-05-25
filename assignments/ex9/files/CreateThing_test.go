package httpTraining

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fixedDate = time.Date(2021, 5, 24, 0, 0, 0, 0, time.UTC)

func returnSuccessResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := GetThingResponse{
		UUID:    "123",
		Name:    "",
		Value:   "",
		Updated: fixedDate,
		Created: fixedDate,
	}

	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("Invalid input")
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		t.Fatalf("Error encoding response")
	}
}

func TestClient_CreateThing(t *testing.T) {
	// fake server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/thing/new", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	type fields struct {
		Hostname string
	}
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetThingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", fields{testServer.URL}, args{"Lokesh", "value"}, &GetThingResponse{"123", "Lokesh", "value", fixedDate, fixedDate}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				Hostname: tt.fields.Hostname,
			}
			got, err := cl.CreateThing(tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateThing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateThing() = %v, want %v", got, tt.want)
			}
		})
	}
}
