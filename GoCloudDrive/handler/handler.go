package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	slog.Info("Hello world")		
}

func UploadFile( w http.ResponseWriter, r *http.Request) {
	f, fileHandler, err := r.FormFile("file")

	if err != nil {
		slog.Error("Fail to get file from Form: "+ err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer f.Close()
	slog.Info("file name is: " + fileHandler.Filename)
	slog.Info(fmt.Sprintf("file size is: %d", fileHandler.Size))

	w.WriteHeader((http.StatusOK))
	return
}