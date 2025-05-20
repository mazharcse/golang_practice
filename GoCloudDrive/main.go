package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
)



func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	server :=http.NewServeMux()
	
	slog.Info("Starting the server on port 8080 ")

	err := http.ListenAndServe(":8080", server)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server closed")
	} else if err != nil {
		slog.Error("Error starting server: %s\n", err.Error())
		os.Exit(1)

	}
}