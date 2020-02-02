package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func testLink(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "derp");
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", testLink)
	log.Fatal(http.ListenAndServe(":8080", router))

}

