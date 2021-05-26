package serverStruct

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ldej/go-training/examples/db"
)

type server struct {
	router *mux.Router
	db     db.DB
}

func NewServer(db db.DB) *server {
	s := &server{
		db: db,
	}
	s.Routes()
	return s
}

func (s *server) Routes() {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/thing", s.GetAllThing).Methods("GET")
	s.router.HandleFunc("/thing/{uuid}", s.GetThing).Methods("GET")
	s.router.HandleFunc("/thing/{uuid}", s.PutThing).Methods("PUT")
	s.router.HandleFunc("/thing/{uuid}", s.DeleteThing).Methods("DELETE")
	s.router.HandleFunc("/thing/new", s.PostThing).Methods("POST")
}

func (s *server) ListenAndServe(addr string) {
	hs := &http.Server{Addr: addr, Handler: s.router}
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type ThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
