package store

import (
	"github.com/google/uuid"

	"DDD/coffeeco/internal"
)

type Store struct {
	ID uuid.UUID
	Location string
	ProductsForSale []internal.Product
}