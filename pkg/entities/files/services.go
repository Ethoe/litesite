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

type File struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	FileName   string    `json:"filename"`
	FilePath   string    `json:"file_path"`
	FileSize   uint64    `json:"file_size"`
	UploadDate time.Time `json:"upload_date"`
	URL        string    `json:"url"`
}

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

func GetFilesByUserID(db *sql.DB, user users.User, limit, page int) ([]File, error) {
	query := "SELECT id, filename, file_size, upload_date, url FROM files WHERE user_id = ? LIMIT ? OFFSET ?"
	rows, err := db.Query(query, user.ID, limit, limit*page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fileList []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.FileName, &file.FileSize, &file.UploadDate, &file.URL)
		if err != nil {
			return nil, err
		}
		fileList = append(fileList, file)
	}

	return fileList, nil
}
