package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// Handle Ctrl+C or SIGTERM to clean up before exiting
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting down server...")
		// Remove the log file when shutting down
		os.Remove("server.log")
		os.Exit(0)
	}()

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
