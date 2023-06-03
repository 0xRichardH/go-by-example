package main

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/purchase"
	"coffeeco/internal/store"
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

const (
	// This is the credentials for mongo if you run docker-compose up in this repo.
	mongoConnectionStr = "mongodb://root:example@localhost:27017"
	// This is the test key from Stripe's documentation. Feel free to use it, no charges will actually be made.
	stripeTestAPIKey = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
)

func main() {
	ctx := context.Background()

	// This is a test token from Stripe's documentation. Feel free to use it, no charges will actually be made.
	cardToken := "tok_visa"

	paymentService, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	purchaseRepo, err := purchase.NewMongoRepo(ctx, mongoConnectionStr)
	if err != nil {
		log.Fatal(err)
	}

	storeRepo, err := store.NewMongoRepo(ctx, mongoConnectionStr)
	if err != nil {
		log.Fatal(err)
	}
	storeService := store.NewService(storeRepo)

	purchaseService := purchase.NewService(paymentService, *storeService, purchaseRepo)

	storeID := uuid.New()
	purchase := purchase.Purchase{
		Store: store.Store{
			ID: storeID,
		},
		ProductsToPurchase: []coffeeco.Product{
			{
				ItemName:  "Late",
				BasePrice: *money.New(3999, "USD"),
			},
		},
		PaymentMethod: payment.MEANS_CARD,
		CardToken:     &cardToken,
	}
	if err := purchaseService.CompletePurchase(ctx, &purchase, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("purchase was successful")
}
