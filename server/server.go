package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	serverPort = "localhost:3002"
)

func Run() {
	handler := setUpServer()
	srv := &http.Server{Addr: serverPort, Handler: handler}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM) // interrupt signal sent from terminal, system
	<-sigint
}

func setUpServer() http.Handler {
	mux := http.NewServeMux()
	return mux
}
