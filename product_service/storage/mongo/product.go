package mongo

import (
	pb "EXAM3_with_mongodb/product_service/genproto/product_service"
	"EXAM3_with_mongodb/product_service/pkg/logger"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	Db  *mongo.Database
	log logger.Logger
}

func NewProductRepo(Db *mongo.Database, log logger.Logger) *productRepo {
	return &productRepo{
		Db:  Db,
		log: log,
	}
}

func (r *productRepo) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	collection := r.Db.Collection("products")
	result, err := collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	var response pb.Product
	filter := bson.M{"_id": result.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *productRepo) GetProductById(ctx context.Context, req *pb.ProductId) (*pb.Product, error) {
	collection := r.Db.Collection("products")

	var response pb.Product
	filter := bson.M{"id": req.ProductId}
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *productRepo) ListProducts(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	collection := r.Db.Collection("products")

	var response pb.GetAllProductResponse

	reqOptions := options.Find()

	reqOptions.SetSkip(int64(req.Page-1) * int64(req.Limit))
	reqOptions.SetLimit(int64(req.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, reqOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product pb.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}

		response.Count++
		response.Products = append(response.Products, &product)
	}

	return &response, nil
}

func (r *productRepo) IncreaseProductAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error) {
	collection := r.Db.Collection("products")

	var response pb.Product
	filter := bson.M{"id": req.ProductId}
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return &pb.ProductAmountResponse{IsEnough: false, Product: nil}, err
	}

	updateReq := bson.M{
		"$set": bson.M{
			"amount":     req.Amount + response.Amount,
			"updated_at": time.Now(),
		},
	}

	err = collection.FindOneAndUpdate(ctx, filter, updateReq).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &pb.ProductAmountResponse{IsEnough: true, Product: &response}, nil
}

func (r *productRepo) DecreaseProductAmount(ctx context.Context, req *pb.ProductAmountRequest) (*pb.ProductAmountResponse, error) {
	collection := r.Db.Collection("products")

	var response pb.Product
	filter := bson.M{"id": req.ProductId}
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return &pb.ProductAmountResponse{IsEnough: false, Product: nil}, err
	}

	if response.Amount == 0 {
		return nil, fmt.Errorf("not enough")
	}

	if response.Amount-req.Amount < 0 {
		return &pb.ProductAmountResponse{IsEnough: false, Product: &response}, err
	}

	updateReq := bson.M{
		"$set": bson.M{
			"amount":     response.Amount - req.Amount,
			"updated_at": time.Now(),
		},
	}

	err = collection.FindOneAndUpdate(ctx, filter, updateReq).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &pb.ProductAmountResponse{IsEnough: true, Product: &response}, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	collection := r.Db.Collection("products")

	var response pb.Product

	filter := bson.M{"id": req.Id}

	updateReq := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
			"price":       req.Price,
			"amount":      req.Amount,
			"updated_at":  time.Now(),
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, updateReq).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *productRepo) DeleteProduct(ctx context.Context, req *pb.ProductId) (*pb.Status, error) {
	collection := r.Db.Collection("products")

	filter := bson.M{"id": req.ProductId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return &pb.Status{Success: false}, err
	}

	return &pb.Status{Success: true}, nil
}

func (r *productRepo) CheckAmount(ctx context.Context, req *pb.ProductId) (*pb.CheckAmountResponse, error) {
	collection := r.Db.Collection("products")

	var checkResult pb.CheckAmountResponse
	var response pb.Product
	filter := bson.M{"id": req.ProductId}
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	checkResult.Amount = response.Amount
	checkResult.ProductId = response.Id

	return &checkResult, nil
}

func (r *productRepo) BuyProduct(ctx context.Context, req *pb.BuyProductRequest) (*pb.Product, error) {
	collection := r.Db.Collection("users_products")

	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	response, err := r.GetProductById(ctx, &pb.ProductId{ProductId: req.ProductId})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *productRepo) GetBoughtProductsByUserId(ctx context.Context, req *pb.UserId) (*pb.GetBoughtProductsResponse, error) {
	collection := r.Db.Collection("users_products")

	var products []*pb.Product

	cursor, err := collection.Find(ctx, bson.M{"user_id": req.UserId})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var order pb.BuyProductRequest
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}

		product, err := r.GetProductById(ctx, &pb.ProductId{ProductId: order.ProductId})
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	response := &pb.GetBoughtProductsResponse{
		Products: products,
	}

	return response, nil
}
