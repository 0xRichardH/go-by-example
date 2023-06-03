package loyalty

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/store"
	"context"

	"github.com/google/uuid"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	FreeDriksAvailable                    int
	RemainingDrinkPurchasesUntilFreeDrink int
	store                                 store.Store
	coffeeLover                           coffeeco.CoffeeLover
}

func (c *CoffeeBux) AddStamp() {
	if c.RemainingDrinkPurchasesUntilFreeDrink == 1 {
		c.RemainingDrinkPurchasesUntilFreeDrink = 10
		c.FreeDriksAvailable += 1
	} else {
		c.RemainingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, products []coffeeco.Product) error {
	productsLen := len(products)
	if productsLen <= 0 {
		return ErrInvalidPurchaseCount
	}

	if c.FreeDriksAvailable < productsLen {
		return ErrInsuffientFreeDriks
	}

	c.FreeDriksAvailable = c.FreeDriksAvailable - productsLen
	return nil
}
