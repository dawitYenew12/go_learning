package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define flags
	filePath := flag.String("f", "", "File or folder to share")
	port := flag.String("p", "9090", "Port to serve on")
	flag.Parse()

	// Validate input
	if *filePath == "" {
		log.Fatal("Please provide a file or folder to share using -f flag")
	}

	// Get absolute path
	absPath, err := filepath.Abs(*filePath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	// Check if file exists
	fileInfo, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		log.Fatalf("File or folder does not exist: %s", absPath)
	}

	// If it's a directory, serve files within it
	if fileInfo.IsDir() {
		fmt.Printf("Serving directory '%s' on http://localhost:%s\n", absPath, *port)
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(absPath))))
	} else {
		// If it's a file, serve only that file
		fmt.Printf("Serving file '%s' on http://localhost:%s\n", absPath, *port)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, absPath)
		})
	}

	addr := fmt.Sprintf("localhost:%s", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
