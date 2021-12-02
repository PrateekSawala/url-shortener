package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"url-shortener/domain/logging"
)

var (
	serverPort = "localhost:3002"
)

func Run() {
	/* Initialize Logging */
	logging.InitializeLogging()

	log := logging.Log("Run")

	handler := setUpServer()
	srv := &http.Server{Addr: serverPort, Handler: handler}
	go func() {
		log.Infof("Starting server at port: %s", serverPort)

		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Info("Server shut down.")
			} else {
				log.Error("failed to start server", err)
			}
		}
	}()
	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM) // interrupt signal sent from terminal, system
	<-sigint

	log.Info("Shutting down server")
}

func setUpServer() http.Handler {
	mux := http.NewServeMux()
	return mux
}
