package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Log the timestamp and request details
	log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

	// Your request handling logic here
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a custom logger that writes log entries to a file
	logFile, err := os.Create("server.log")
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	// Set the custom logger as the default logger
	log.SetOutput(logger.Writer())

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
