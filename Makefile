# proto paths
ORDER_PROTO_PATH := ./order-service/pkg/api/v1
PRODUCT_PROTO_PATH := ./product-service/pkg/api/v1

GO_RUN := go run
PROTOC := protoc

# proto
.PHONY: protos
protos:
	@echo "Generating protobuf files for Order Service..."
	@$(PROTOC) --proto_path=$(ORDER_PROTO_PATH) --go_out=$(ORDER_PROTO_PATH) --go-grpc_out=$(ORDER_PROTO_PATH) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(ORDER_PROTO_PATH)/*.proto
	@echo "Generating protobuf files for Product Service..."
	@$(PROTOC) --proto_path=$(PRODUCT_PROTO_PATH) --go_out=$(PRODUCT_PROTO_PATH) --go-grpc_out=$(PRODUCT_PROTO_PATH) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PRODUCT_PROTO_PATH)/*.proto

# services
.PHONY: run-all run-order-service run-product-service run-api-gateway
run-all: run-order-service run-product-service run-api-gateway

run-order-service:
	@echo "Running Order Service..."
	@$(GO_RUN) ./order-service/cmd/main.go &

run-product-service:
	@echo "Running Product Service..."
	@$(GO_RUN) ./product-service/cmd/main.go &

run-api-gateway:
	@echo "Running API Gateway..."
	@$(GO_RUN) ./api-gateway/cmd/main.go &

# clean
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(ORDER_PROTO_PATH)/*.pb.go $(PRODUCT_PROTO_PATH)/*.pb.go
	@echo "Cleanup complete!"
