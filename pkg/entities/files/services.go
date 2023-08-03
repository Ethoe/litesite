package files

import (
	"cmd/server/main.go/pkg/entities/users"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const dataPath = "/data"
const baseUrl = "https://ethoe.dev/files/%s/%s"

func CreateFile(db *sql.DB, file multipart.File, fileHeader *multipart.FileHeader, user users.User) (string, error) {
	// Upload directory
	uploadPath := filepath.Join(dataPath, user.Username)
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Create file in upload directory
	filePath := filepath.Join(uploadPath, fileHeader.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy data to file
	_, err = io.Copy(dst, file)
	if err != nil {
		os.Remove(filePath)
		return "", err
	}

	url := fmt.Sprintf(baseUrl, user.Username, fileHeader.Filename)

	query := "INSERT INTO files (user_id, filename, file_path, file_size, upload_date, url) VALUES (?, ?, ?, ?, ?, ?)"
	fileSize := fileHeader.Size

	_, err = db.Exec(query, user.ID, fileHeader.Filename, filePath, fileSize, time.Now(), url)
	if err != nil {
		os.Remove(filePath)
		return "", err
	}

	return url, nil
}

func DeleteFile(db *sql.DB, fileId int, user users.User) error {
	// Check if the file exists in the database and belongs to the user
	var filePath string
	err := db.QueryRow("SELECT file_path FROM files WHERE id = ? AND user_id = ?", fileId, user.ID).Scan(&filePath)
	if err != nil {
		return err
	}

	// Delete the file from disk
	if err := os.Remove(filePath); err != nil {
		return err
	}

	// Delete the file metadata from the database
	_, err = db.Exec("DELETE FROM files WHERE id = ?", fileId)
	if err != nil {
		return err
	}

	return nil
}
