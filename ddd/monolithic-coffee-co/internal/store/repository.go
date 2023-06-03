package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (int64, error)
}

type MongoRepository struct {
	storeDiscounts *mongo.Collection
}

var _ Repository = (*MongoRepository)(nil)

func NewMongoRepo(ctx context.Context, connectionStr string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	discounts := client.Database("coffeeco").Collection("store_discounts")
	return &MongoRepository{storeDiscounts: discounts}, nil
}

func (r *MongoRepository) GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (int64, error) {
	var discount int64
	if err := r.storeDiscounts.FindOne(ctx, bson.D{{"store_id", storeID.String()}}).Decode(&discount); err != nil {
		if err == mongo.ErrNoDocuments {
			return discount, ErrNoDiscount
		}

		return discount, fmt.Errorf("failed to get store discount: %w", err)
	}
	return discount, nil
}
