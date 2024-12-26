package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UploadFile(file io.Reader, filename string) error {
	// Sanitize the filename
	safeFilename := filepath.Base(filename)

	// Ensure the 'data' directory exists
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create data directory: %v", err)
	}

	// Define the destination file path
	dstPath := filepath.Join("data", safeFilename)

	// Create the file in the 'data' directory
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer dst.Close()

	// Copy the file contents to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("unable to save file: %v", err)
	}

	return nil
}

func DownloadFile(filename string) (*os.File, error) {
	// Sanitize the filename
	safeFilename := filepath.Base(filename)

	// Define the file path in the 'data' directory
	filePath := filepath.Join("data", safeFilename)

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}

	return file, nil
}
