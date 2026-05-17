package main

import (
    "log"
    "net"
    
    "google.golang.org/grpc"
)

type paymentServer struct{}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    log.Println("Payment Service running on port 50052")
    
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
