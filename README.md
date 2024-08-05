# Microservices Demo Project

This project demonstrates a basic microservices architecture with an API Gateway, Order Service, and Product Service.

## Prerequisites

- Protocol Buffers compiler (protoc)

## Getting Started

1. Generate Protocol Buffers code:
    ```
    make protos
    ```

2. Run all services:
    ```
    make run-all
    ```
3. To clean generated protobuf files:
    ```
    make clean
    ```

## Services

### API Gateway

    Accepts requests http and distributes them to services by converting them into data format proto structure

- Port: 8080
- Endpoints:
  - GET /products/{id}
  - POST /orders
  - POST /products
  - GET /products
  - PUT /products/{id}

### Order Service
- gRPC Port: 50052
- Manages order creation and processing

### Product Service
- gRPC Port: 50051
- Manages product information and inventory


