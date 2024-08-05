package main

import (
	"log"
	"net"

	grpcorder "order-service/internal/delivery/grpc"
	"order-service/internal/repository/memory"
	"order-service/internal/usecase"
	pb "order-service/pkg/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := memory.NewOrderMemoryRepository()
	productConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}
	defer productConn.Close()

	usecase := usecase.NewOrderUsecase(repo, productConn)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	handler := grpcorder.NewOrderHandler(usecase)

	pb.RegisterOrderServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("OrderService gRPC server is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
