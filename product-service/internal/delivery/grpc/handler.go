package grpcproduct

import (
	"context"

	"product-service/internal/domain/model"
	"product-service/internal/usecase"
	pb "product-service/pkg/api/v1"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	usecase *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: uc}
}

func (h *ProductHandler) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	product := &model.Product{
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}
	err := h.usecase.AddProduct(product)
	if err != nil {
		return &pb.AddProductResponse{Success: false, Error: err.Error()}, nil
	}
	return &pb.AddProductResponse{Success: true}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := h.usecase.GetProduct(req.Id)
	if err != nil {
		return &pb.GetProductResponse{Error: err.Error()}, nil
	}
	return &pb.GetProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := h.usecase.ListProducts()
	if err != nil {
		return nil, err
	}

	response := &pb.ListProductsResponse{}
	for _, product := range products {
		response.Products = append(response.Products, &pb.GetProductResponse{
			Id:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	return response, nil
}
