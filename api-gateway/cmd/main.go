package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	orderpb "order-service/pkg/api/v1"
	productpb "product-service/pkg/api/v1"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	r := mux.NewRouter()

	// conn to microservices as a client

	productConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}
	defer productConn.Close()
	productClient := productpb.NewProductServiceClient(productConn)

	orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)


	// handlers gateway

	r.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := productClient.GetProduct(context.Background(), &productpb.GetProductRequest{Id: int32(id)})
		if err != nil {
			http.Error(w, "Failed to get product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}).Methods("GET")

	r.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		var order orderpb.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid order data", http.StatusBadRequest)
			return
		}

		response, err := orderClient.CreateOrder(context.Background(), &order)
		if err != nil {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		var newProduct productpb.AddProductRequest
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		product, err := productClient.AddProduct(context.Background(), &newProduct)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}).Methods("POST")
	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := productClient.ListProducts(context.Background(), &productpb.ListProductsRequest{})
		if err != nil {
			http.Error(w, "Failed to get products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}).Methods("GET")
	r.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var updateProduct productpb.UpdateProductRequest
		if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}
		updateProduct.Id = int32(id)

		product, err := productClient.UpdateProduct(context.Background(), &updateProduct)
		if err != nil {
			http.Error(w, "Failed to update product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}).Methods("PUT")

	log.Println("API Gateway is running on port 8080")
	http.ListenAndServe(":8080", r)
}
