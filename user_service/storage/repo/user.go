package repo

import (
	pb "EXAM3/user_service/genproto/user_service"
	"context"
)

type UserStorageI interface {
	Create(ctx context.Context, request *pb.User) (*pb.User, error)
	// GetUserByUsername(ctx context.Context, request *pb.Username) (*pb.User, error)
	GetUserById(ctx context.Context, request *pb.UserId) (*pb.User, error)
	// GetUserByEmail(ctx context.Context, request *pb.Email) (*pb.User, error)
	UpdateUserById(ctx context.Context, request *pb.User) (*pb.User, error)
	Delete(ctx context.Context, request *pb.UserId) (*pb.Empty, error)
	ListUsers(ctx context.Context, request *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error)
	CheckField(ctx context.Context, request *pb.CheckFieldRequest) (*pb.CheckFieldResponse, error)
}
