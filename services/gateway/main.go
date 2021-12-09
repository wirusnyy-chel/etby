package main

import (
	in "etby/services/gateway/internal"
	"net/http"
)

type config struct {
	port string
	dev  bool
}

func main() {
	conf := initConfig()
	mux := in.Mux()
	server := http.Server{
		Addr:    conf.port,
		Handler: mux,
	}
	server.ListenAndServe()
}

func initConfig() config {
	cof := config{
		port: ":8081",
		dev:  true,
	}
	return cof
}
