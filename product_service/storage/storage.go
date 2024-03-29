package storage

import (
	"EXAM3_with_mongodb/product_service/pkg/logger"
	m "EXAM3_with_mongodb/product_service/storage/mongo"
	"EXAM3_with_mongodb/product_service/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	Db          *mongo.Database
	productRepo repo.ProductStorageI
}

func NewStoragePg(Db *mongo.Database, log logger.Logger) *storagePg {
	return &storagePg{
		Db:          Db,
		productRepo: m.NewProductRepo(Db, log),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
