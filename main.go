package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    // Kubernetes will call this to check if the app is running
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Get the hostname of the machine/pod running this app
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Hello from Go WebApp! Host: %s, Time: %s", host, time.Now().Format(time.RFC3339))
}

func main() {
    // Set up the routes
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/health", healthHandler)

    // Read the port from an environment variable, default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on :%s", port)
    // Start the server
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
