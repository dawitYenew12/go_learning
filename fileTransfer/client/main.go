package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile uploads a file to the server
func UploadFile(filePath, serverAddress string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a buffer to hold the multipart data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the form file field
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	// Copy file content to the form file
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("error closing writer: %v", err)
	}

	// Make the POST request to upload the file
	url := fmt.Sprintf("http://%s:8080/upload", serverAddress)
	resp, err := http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Print the server's response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	fmt.Println("Server response:", string(respBody))
	return nil
}

// DownloadFile downloads a file from the server
func DownloadFile(fileName, savePath, serverAddress string) error {
	url := fmt.Sprintf("http://%s:8080/download/%s", serverAddress, fileName)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Create the file locally
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("error creating local file: %v", err)
	}
	defer file.Close()

	// Write the file content to the disk
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing file content: %v", err)
	}

	fmt.Println("File downloaded successfully:", savePath)
	return nil
}

func main() {
	// Command-line arguments for upload/download and file paths
	uploadFlag := flag.Bool("upload", false, "Upload a file")
	downloadFlag := flag.Bool("download", false, "Download a file")
	serverAddress := flag.String("server", "localhost", "Server IP address")
	filePath := flag.String("file", "", "Path to the file to upload or download")
	savePath := flag.String("save", "", "Path to save the downloaded file")

	flag.Parse()

	if *uploadFlag {
		if *filePath == "" {
			fmt.Println("Please provide the file path to upload")
			os.Exit(1)
		}
		err := UploadFile(*filePath, *serverAddress)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else if *downloadFlag {
		if *filePath == "" || *savePath == "" {
			fmt.Println("Please provide the file name to download and the save path")
			os.Exit(1)
		}
		err := DownloadFile(*filePath, *savePath, *serverAddress)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Please specify whether to upload or download a file using -upload or -download")
	}
}
