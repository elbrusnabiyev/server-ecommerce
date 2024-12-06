// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"log"

	pb "C:/Users/Elbrus/go/GoApps/Tsoukalos/gRPC_ServiceProductsOne/productinfo/go/server/ecommerce"

	"github.com/gofrs/uuid/v5"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Сервер используется для реализации ecommerce/productinfo
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct реализует ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	// бизнес логика:
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating ProductID", err)
	}
	in.ID = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.ID] = in
	return &pb.ProductID{Value: in.ID}, status.New(codes.OK, "").Err()
}

// GetProduct реализует ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	// бизнес логика:
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Server(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
