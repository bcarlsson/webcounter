package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Counter struct {
	Count int `json:"count"`
}

func main() {
	port := flag.String("port", "8888", "Listening port")
	flag.Parse()
	myCounter := Counter{
		Count: 0,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", myCounter.Counter)
	router.HandleFunc("/reset", myCounter.Reset)

	log.Fatal(http.ListenAndServe(":"+*port, router))
}

func (c *Counter) Counter(w http.ResponseWriter, r *http.Request) {
	c.Count++
	json.NewEncoder(w).Encode(c)
}

func (c *Counter) Reset(w http.ResponseWriter, r *http.Request) {
	c.Count = 0
	fmt.Fprint(w, "Counter has been reset\n")
}
