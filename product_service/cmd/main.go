package main

import (
	"EXAM3_with_mongodb/product_service/config"
	pb "EXAM3_with_mongodb/product_service/genproto/product_service"
	"EXAM3_with_mongodb/product_service/pkg/db"
	"EXAM3_with_mongodb/product_service/pkg/logger"
	"EXAM3_with_mongodb/product_service/service"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log := logger.New(cfg.LogLevel, "product-service")
	defer logger.Cleanup(log)

	log = logger.WithFields(log, logger.String("file", "log.txt"))

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	postgres, err := db.New(*cfg)
	if err != nil {
		log.Fatal("sql connection to postgres error", logger.Error(err))
	}

	authorizationService := service.NewProductService(postgres, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, authorizationService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
