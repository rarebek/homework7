package service

import (
	pb "EXAM3_with_mongodb/product_service/genproto/product_service"
	l "EXAM3_with_mongodb/product_service/pkg/logger"
	"EXAM3_with_mongodb/product_service/storage"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedProductServiceServer
}

func NewProductService(Db *mongo.Database, log l.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStoragePg(Db, log),
		logger:  log,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	return s.storage.Product().CreateProduct(ctx, req)
}

func (s *ProductService) GetProductById(ctx context.Context, req *pb.ProductId) (*pb.Product, error) {
	return s.storage.Product().GetProductById(ctx, req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	return s.storage.Product().UpdateProduct(ctx, req)
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.ProductId) (*pb.Status, error) {
	return s.storage.Product().DeleteProduct(ctx, req)
}

func (s *ProductService) CheckAmount(ctx context.Context, req *pb.ProductId) (*pb.CheckAmountResponse, error) {
	return s.storage.Product().CheckAmount(ctx, req)
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	return s.storage.Product().ListProducts(ctx, req)
}

func (s *ProductService) IncreaseAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error) {
	return s.storage.Product().IncreaseProductAmount(ctx, req)
}

func (s *ProductService) DecreaseAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error) {
	return s.storage.Product().DecreaseProductAmount(ctx, req)
}

func (s *ProductService) GetBoughtProductsByUserId(ctx context.Context, req *pb.UserId) (*pb.GetBoughtProductsResponse, error) {
	return s.storage.Product().GetBoughtProductsByUserId(ctx, req)
}

func (s *ProductService) BuyProduct(ctx context.Context, req *pb.BuyProductRequest) (*pb.Product, error) {
	return s.storage.Product().BuyProduct(ctx, req)
}
