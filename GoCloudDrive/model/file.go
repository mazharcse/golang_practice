package model

import (
	"database/sql"
	"go-cloud-drive/utils"
	"log/slog"
	"strings"
	"time"
)

const (
    FileTable = "file"    
    qInsertFileMeta = "INSERT INTO " + FileTable + " (name, location, size) VALUES ($1, $2, $3) RETURNING id, uploaded_at"
    qGetFileMeta = "SELECT id, name, location, uploaded_at, status, size from " + FileTable
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

 func GetFileMetas(search string, page int, limit int) ([]FileMeta, error) {
    search = strings.TrimSpace(search)
    queryLimit := limit
    queryOffset := limit * (page -1)
    
    var queryStmt = qGetFileMeta
    var err error
    var rows *sql.Rows

    if search != "" {
        queryStmt += " WHERE name ~* $3"
    }

    queryStmt += " ORDER by uploaded_at DESC LIMIT $1 OFFSET $2"
    
    stmt, err := utils.GetDB().DB.Prepare(queryStmt)
    if err != nil {
        slog.Error("Fail to prepare statement: " + err.Error())
        return nil, err
    }

    defer stmt.Close()

    if search != "" {
        rows, err = stmt.Query(queryLimit, queryOffset, search)
    } else {
        rows, err = stmt.Query(queryLimit, queryOffset)
    }

    if err != nil {
        slog.Error("Fail to query rows: " + err.Error())
        return nil, err
    }

    defer rows.Close()

    fileMetas := make([]FileMeta, 0)

    for rows.Next() {
        var fm FileMeta
        err = rows.Scan(&fm.Id, &fm.Name, &fm.Location, &fm.UploadedAt, &fm.Status, &fm.Size)
        if err != nil {
            slog.Error("Fail to scan row: " + err.Error())
            return nil, err
        }
        fileMetas = append(fileMetas, fm)
    }

    return fileMetas, nil

}