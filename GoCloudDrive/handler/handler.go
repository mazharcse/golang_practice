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
	"strconv"
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
// Get files url:/files/search=abc&page=1&limit=20
func GetFiles(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		slog.Error("Fail to convert page to int: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		slog.Error("Fail to convert limit to int: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := model.GetFileMetas(search, page, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	return
}