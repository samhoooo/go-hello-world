package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprint(w, "hello world")
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/api/hello", helloHandler)

    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}