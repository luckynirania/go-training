/* try some handlerfunction
create some function try to see for url parameters
httpClient tests
*/
package main

import (
	"log"
	"net/http"

	customAPI "lokesh/assignments/ex10/files"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/thing/", customAPI.GETAll).Methods("GET")
	router.HandleFunc("/thing/{uuid}", customAPI.GET).Methods("GET")
	router.HandleFunc("/thing/{uuid}", customAPI.PUT).Methods("PUT")
	router.HandleFunc("/thing/{uuid}", customAPI.DELETE).Methods("DELETE")
	router.HandleFunc("/thing/new", customAPI.POST).Methods("POST")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
