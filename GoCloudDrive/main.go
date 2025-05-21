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

	err = os.MkdirAll(os.Getenv("ROOT_DIR"), 0755)

	if err != nil {
		slog.Error("Fail to create root dir " + err.Error())
		os.Exit(1)
	}
	
	port := os.Getenv("PORT")

	server :=http.NewServeMux()
 
	server.HandleFunc("GET /hello", handler.Hello)

	// Upload file
	server.HandleFunc("POST /file", handler.UploadFile)
	
	slog.Info("Starting the server on port " + port + " ...")

	err = http.ListenAndServe(":"+port, middleware.RequestLogger(server))

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server closed")
	} else if err != nil {
		slog.Error("Error starting server: %s\n", err.Error())
		os.Exit(1)

	}
}