package repo

import (
	pb "EXAM3_with_mongodb/product_service/genproto/product_service"
	"context"
)

type ProductStorageI interface {
	CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error)
	GetProductById(ctx context.Context, req *pb.ProductId) (*pb.Product, error)
	ListProducts(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
	IncreaseProductAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error)
	DecreaseProductAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error)
	UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error)
	DeleteProduct(ctx context.Context, req *pb.ProductId) (*pb.Status, error)
	CheckAmount(ctx context.Context, req *pb.ProductId) (*pb.CheckAmountResponse, error)
	GetBoughtProductsByUserId(ctx context.Context, req *pb.UserId) (*pb.GetBoughtProductsResponse, error)
	BuyProduct(ctx context.Context, req *pb.BuyProductRequest) (*pb.Product, error)
}
