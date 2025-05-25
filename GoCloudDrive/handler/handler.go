package handler

import (
	"encoding/json"
	"fmt"
	"go-cloud-drive/model"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	// "syscall/js"
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

	newFilePath := filepath.Join(os.Getenv("ROOT_DIR"), fileHandler.Filename)
	newFile, err := os.Create(newFilePath)

	if err != nil {
		slog.Error("Fail to created new file " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, f)

	if err != nil {
		slog.Error("Fail to save file ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newFileMeta := model.FileMeta {
		Name: fileHandler.Filename,
		Location: newFilePath,
		Size: fileHandler.Size,
		Status: 1,
	}

	newFileMeta, err = model.InsertFileMeta(newFileMeta)
	if err != nil {
		slog.Error("Fail to insert file meta: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader((http.StatusOK))
	json.NewEncoder(w).Encode(newFileMeta)
	return
}