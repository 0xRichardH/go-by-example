package purchase

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

var _ Repository = (*MongoRepository)(nil)

func NewMongoRepo(ctx context.Context, connectionStr string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")
	return &MongoRepository{purchases: purchases}, nil
}

func (r *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoPurchase := toMongoPurchase(purchase)
	_, err := r.purchases.InsertOne(ctx, mongoPurchase)
	if err != nil {
		return fmt.Errorf("failed to store purchase: %w", err)
	}

	return nil
}

type mongoPurchase struct {
	ID                uuid.UUID          `bson:"id"`
	Store             store.Store        `bson:"store"`
	ProductsPurchased []coffeeco.Product `bson:"products_purchased"`
	Total             int64              `bson:"total"`
	PaymentMethod     payment.Means      `bson:"payment_method"`
	CardToken         *string            `bson:"card_token"`
	TimeOfPurchase    time.Time          `bson:"time_of_purchase"`
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		ID:                p.id,
		Store:             p.Store,
		ProductsPurchased: p.ProductsToPurchase,
		Total:             p.total.Amount(),
		PaymentMethod:     p.PaymentMethod,
		CardToken:         p.CardToken,
		TimeOfPurchase:    p.timeOfPurchase,
	}
}
