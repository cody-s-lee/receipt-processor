package main

import (
	"github.com/cody-s-lee/receipt-processor/backend/receipts"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// create a type that satisfies the `receipts.ServerInterface`, which contains an implementation of every operation from the generated code
	server := receipts.NewServer()

	r := chi.NewMux()

	// get an `http.Handler` that we can use
	h := receipts.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:80",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
