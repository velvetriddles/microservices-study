package main

import (
	"log"
	"net"

	grpcproduct "product-service/internal/delivery/grpc"
	"product-service/internal/repository/memory"
	"product-service/internal/usecase"
	pb "product-service/pkg/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := memory.NewProductMemoryRepository()
	usecase := usecase.NewProductUsecase(repo)
	handler := grpcproduct.NewProductHandler(usecase)

	pb.RegisterProductServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("ProductService gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
