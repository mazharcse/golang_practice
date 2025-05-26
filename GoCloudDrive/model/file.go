package model

import (
	"go-cloud-drive/utils"
	"log/slog"
	"time"
)

const (
    FileTable = "file"
    //qInsertFileMeta="INSERT INTO file (name,location,size) VALUES ($1, $2, $3) RETURNING id, uploaded_at" = "INSERT INTO " + FileTable + " (name, location, size) VALUES ($1, $2, $3) RETURING id, uploaded_at"
    qInsertFileMeta = "INSERT INTO " + FileTable + " (name, location, size) VALUES ($1, $2, $3) RETURNING id, uploaded_at"
)

 type FileMeta struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Location string `json:"location"`
    UploadedAt string `json:"uploaded_at"`
    Status int64 `json:"status"`
    Size int64 `json:"size"`
 }

 func InsertFileMeta(f FileMeta) (FileMeta, error) {
    stmnt, err := utils.GetDB().DB.Prepare(qInsertFileMeta)
    if err != nil {
        slog.Error("Fail to prepare statement: "+ err.Error())
        return f, err
    }

    defer stmnt.Close()

    var uploadedAt time.Time
    err = stmnt.QueryRow(f.Name, f.Location, f.Size).Scan(&f.Id, &uploadedAt)
    if err != nil {
        slog.Error("Fail to query row: " + err.Error())
        return f, err
    }
    f.UploadedAt = uploadedAt.Format("2006-01-12 15:04:03")
    return f, nil
 }