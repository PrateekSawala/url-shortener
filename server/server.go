package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"url-shortener/domain/logging"
	"url-shortener/endpoint"

	"github.com/gorilla/mux"
)

var (
	serverPort = "localhost:3002"
)

func Run() {
	/* Initialize Logging */
	logging.InitializeLogging()
	log := logging.Log("Run")

	go func() {
		log.Infof("Starting server at port: %s", serverPort)
		err := http.ListenAndServe(serverPort, setUpServer())
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

func setUpServer() *mux.Router {
	r := mux.NewRouter()

	// Handler for serving static folder files
	staticFileServer := http.FileServer(http.Dir("./static"))

	r.HandleFunc("/createShortUrl/", endpoint.CreateShortURLHandler)
	r.HandleFunc("/shortUrl/{id}", endpoint.GetOriginalURLHandler)
	r.Handle("/", staticFileServer)
	return r
}
