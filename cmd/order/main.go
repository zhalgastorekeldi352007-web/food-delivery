package main

import (
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "food-delivery/internal/config"
    "food-delivery/internal/infra"
    "food-delivery/internal/logger"
    "food-delivery/internal/order"
    "food-delivery/proto"
    "github.com/nats-io/nats.go"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "go.uber.org/zap"
    "google.golang.org/grpc"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        panic("failed to load config: " + err.Error())
    }

    log, err := logger.InitLogger(cfg.LogLevel)
    if err != nil {
        panic("failed to init logger: " + err.Error())
    }
    defer log.Sync()

    db, err := infra.NewPostgresDB(cfg.DBURL)
    if err != nil {
        log.Fatal("failed to connect to database", zap.Error(err))
    }

    redis, err := infra.NewRedisClient(cfg.RedisURL)
    if err != nil {
        log.Fatal("failed to connect to redis", zap.Error(err))
    }

    nc, err := nats.Connect(cfg.NatsURL)
    if err != nil {
        log.Fatal("failed to connect to NATS", zap.Error(err))
    }
    defer nc.Close()

    repo := order.NewRepository(db, redis, nc)
    service := order.NewService(repo, log)

    // Start metrics server
    go func() {
        http.Handle("/metrics", promhttp.Handler())
        if err := http.ListenAndServe(":"+cfg.HTTPPort, nil); err != nil && err != http.ErrServerClosed {
            log.Error("metrics server failed", zap.Error(err))
        }
    }()

    // Start gRPC server
    lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
    if err != nil {
        log.Fatal("failed to listen", zap.Error(err))
    }

    grpcServer := grpc.NewServer()
    proto.RegisterOrderServiceServer(grpcServer, &orderServer{service: service})

    go func() {
        log.Info("starting gRPC server", zap.String("port", cfg.GRPCPort))
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatal("failed to serve", zap.Error(err))
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Info("shutting down server...")
    grpcServer.GracefulStop()
}
