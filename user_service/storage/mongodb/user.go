package mongodb

import (
	pb "EXAM3/user_service/genproto/user_service"
	"EXAM3/user_service/pkg/logger"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	Db  *mongo.Database
	log logger.Logger
}

func NewUserRepo(Db *mongo.Database, log logger.Logger) *userRepo {
	return &userRepo{
		Db:  Db,
		log: log,
	}
}

func (r *userRepo) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	collection := r.Db.Collection("users")
	iiid := uuid.NewString()
	req.Id = iiid
	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *userRepo) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	collection := r.Db.Collection("users")
	var user pb.User
	filter := bson.M{"id": req.UserId}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUserById(ctx context.Context, req *pb.User) (*pb.User, error) {
	collection := r.Db.Collection("users")
	filter := bson.M{"id": req.Id}
	update := bson.M{
		"$set": bson.M{
			"name":          req.Name,
			"age":           req.Age,
			"username":      req.Username,
			"email":         req.Email,
			"password":      req.Password,
			"refresh_token": req.RefreshToken,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *userRepo) Delete(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	collection := r.Db.Collection("users")
	var user pb.User
	filter := bson.M{"id": req.UserId}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	filter = bson.M{"id": req.UserId}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *userRepo) ListUsers(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	collection := r.Db.Collection("users")

	offset := (req.Page - 1) * req.Limit

	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(req.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*pb.User
	for cursor.Next(ctx) {
		var user pb.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetAllUserResponse{
		Users: users,
	}, nil
}

func (r *userRepo) CheckField(ctx context.Context, req *pb.CheckFieldRequest) (*pb.CheckFieldResponse, error) {
	collection := r.Db.Collection("users")
	filter := bson.M{req.Field: req.Data}
	var status pb.CheckFieldResponse
	err := collection.FindOne(context.TODO(), filter).Decode(&status)
	if err != nil {
		return &pb.CheckFieldResponse{
			Status: false,
		}, nil
	}
	return &pb.CheckFieldResponse{
		Status: true,
	}, nil
}
