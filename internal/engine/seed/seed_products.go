package seed

import (
	"context"
	"log"
	"market/internal/core/ports"
)

func SeedProducts(productRepo ports.ProductRepository) {
	ctx := context.Background()
	for _, product := range products {
		if _, err := productRepo.CreateProduct(ctx, &product); err != nil {
			log.Println(err)
		}
	}
}
