package storage

import (
	"EXAM3/user_service/pkg/logger"
	"EXAM3/user_service/storage/mongodb"
	"EXAM3/user_service/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	Db       *mongo.Database
	userRepo repo.UserStorageI
}

func NewStoragePg(Db *mongo.Database, log logger.Logger) *storagePg {
	return &storagePg{
		Db:       Db,
		userRepo: mongodb.NewUserRepo(Db, log),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
