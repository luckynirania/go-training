package main

import (
	serverStruct "lokesh/assignments/ex11/files"

	"github.com/ldej/go-training/examples/db/inmemory"
)

func main() {
	dbService := inmemory.NewDB()
	server := serverStruct.NewServer(dbService)
	server.ListenAndServe(":8080")
}
