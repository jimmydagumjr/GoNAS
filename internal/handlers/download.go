package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jimmydagumjr/GoNAS/internal/services"
)

func FileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Extract filename from URL parameters
	filename := chi.URLParam(r, "filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	// Use the file service to get the file
	file, err := services.DownloadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error accessing file: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	// Get file metadata for the modification time
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Unable to retrieve file information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set headers to indicate a file download
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")

	// Stream the file content to the client
	http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
}
