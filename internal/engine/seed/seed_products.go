package seed

import (
	"log"
	"market/internal/core/ports"
)

func SeedProducts(productRepo ports.ProductRepository) {
	for _, product := range products {
		if _, err := productRepo.CreateProduct(&product); err != nil {
			log.Println(err)
		}
	}
}
