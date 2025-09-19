package main

import (
	"net/http"

	"l2.18/internal/handler"
	"l2.18/internal/service"
	"l2.18/internal/storage"
)

const portNumber = ":8080"

func main() {
	storage := storage.New()
	srv := service.New(storage)
	h := handler.New(srv)

	s := http.Server{
		Addr:    portNumber,
		Handler: h.Route(),
	}
	s.ListenAndServe()
}
