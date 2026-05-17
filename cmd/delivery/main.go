package main

import (
    "log"
    "net"
    
    "google.golang.org/grpc"
)

type deliveryServer struct{}

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    log.Println("Delivery Service running on port 50053")
    
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
