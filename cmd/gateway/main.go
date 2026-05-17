package main

import (
    "log"
    "net/http"
    "encoding/json"
)

type HealthResponse struct {
    Status string `json:"status"`
    Service string `json:"service"`
}

func main() {
    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(HealthResponse{
            Status: "ok",
            Service: "api-gateway",
        })
    })
    
    // Home endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Food Delivery API Gateway",
            "version": "1.0.0",
            "services": "order:50051, payment:50052, delivery:50053",
        })
    })
    
    // Order endpoint (proxy example)
    http.HandleFunc("/api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "service": "order",
            "endpoint": "/api/v1/orders",
            "method": r.Method,
        })
    })
    
    log.Println("API Gateway running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
