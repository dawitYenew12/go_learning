package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("request: ", r)
// 	fmt.Println("url path: ", r.URL.Path)
// 	if (r.URL.Path == "/favicon.ico") {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Println("Received request for:", r.URL.Path, "from", r.RemoteAddr)
// 	fmt.Fprintf(w, "Hello, World!")
// }

func staticFileHandler (w http.ResponseWriter, r *http.Request) {
	baseDir := "./static"
	filePath := filepath.Join(baseDir, r.URL.Path)
	fmt.Println("file path: ", filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, filePath)
}

func main() {
	// http.HandleFunc("/", handleRequest)
	http.HandleFunc("/", staticFileHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}