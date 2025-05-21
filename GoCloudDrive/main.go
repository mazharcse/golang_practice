package main

import (
	"errors"
	"go-cloud-drive/handler"
	"go-cloud-drive/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
)
  

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file" + err.Error())
		os.Exit(1)
	}
	
	port := os.Getenv("PORT")

	server :=http.NewServeMux()

	// server.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
	// 	slog.Info("Hello world")
	// })
 
	server.HandleFunc("GET /hello", handler.Hello)
	
	slog.Info("Starting the server on port " + port + " ...")

	err = http.ListenAndServe(":"+port, middleware.RequestLogger(server))

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server closed")
	} else if err != nil {
		slog.Error("Error starting server: %s\n", err.Error())
		os.Exit(1)

	}
}