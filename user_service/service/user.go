package service

import (
	pb "EXAM3/user_service/genproto/user_service"
	l "EXAM3/user_service/pkg/logger"
	"EXAM3/user_service/storage"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserService(Db *mongo.Database, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(Db, log),
		logger:  log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req)
	return s.storage.User().Create(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	return s.storage.User().Delete(ctx, req)
}

// func (s *UserService) GetUserByUsername(ctx context.Context, req *pb.Username) (*pb.User, error) {
// 	return s.storage.User().GetUserByUsername(ctx, req)
// }

func (s *UserService) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	return s.storage.User().GetUserById(ctx, req)
}

func (s *UserService) UpdateUserById(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().UpdateUserById(ctx, req)
}

func (s *UserService) ListUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	return s.storage.User().ListUsers(ctx, req)
}
func (s *UserService) CheckField(ctx context.Context, req *pb.CheckFieldRequest) (*pb.CheckFieldResponse, error) {
	return s.storage.User().CheckField(ctx, req)
}

// func (s *UserService) GetUserByEmail(ctx context.Context, req *pb.Email) (*pb.User, error) {
// 	return s.storage.User().GetUserByEmail(ctx, req)
// }
