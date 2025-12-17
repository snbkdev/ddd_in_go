package loyalty

import (
	"DDD/coffeeco/internal"
	"DDD/coffeeco/internal/store"

	"github.com/google/uuid"
)

type CoffeeBux struct {
	ID uuid.UUID
	store store.Store
	coffeeLover internal.CoffeeLover
	FreeDrinksAvailable int
	RemainingDrinkPurchasesUntilFreeDrink int
}