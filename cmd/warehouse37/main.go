package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"warehouse37/api"
	"warehouse37/processing"
	"warehouse37/query"
	"warehouse37/registry"
	"warehouse37/store"
)

func main() {
	path := flag.String("db", "warehouse37.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, err := store.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	r := registry.New(s)
	p := processing.New(s)
	q := query.New(s)
	log.Printf("warehouse37 listening on %s", *addr)
	if err = http.ListenAndServe(*addr, api.New(r, p, q).Handler()); err != nil && !os.IsTimeout(err) {
		log.Fatal(err)
	}
}
