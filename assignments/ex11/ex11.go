package main

import (
	serverStruct "lokesh/assignments/ex11/files"

	"github.com/ldej/go-training/examples/db/sqlite3"
)

func main() {
	dbService, _ := sqlite3.NewDB("Training")
	server := serverStruct.NewServer(dbService)
	server.ListenAndServe(":8080")
}
