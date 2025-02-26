package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadPath = "./uploads"

func main() {
	// Ensure upload directory exists
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating upload directory:", err)
		return
	}

	// Handle file upload route
	http.HandleFunc("/upload", uploadFileHandler)

	// Handle file download route
	http.HandleFunc("/download/", downloadFileHandler)

	// Start the server
	fmt.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// uploadFileHandler handles file uploads via POST requests
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form containing the file
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create file on server
	filePath := filepath.Join(uploadPath, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy file contents
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

// downloadFileHandler handles file downloads via GET requests
func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)

	filePath := filepath.Join(uploadPath, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the header to force download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Serve the file content
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error serving file", http.StatusInternalServerError)
	}
}
