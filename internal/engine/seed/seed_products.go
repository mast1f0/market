package seed

import (
	"log"
	"market/internal/adapters/storage/orm"
)

func SeedProducts(productRepo *orm.ProductRepository) {
	for _, product := range products {
		if _, err := productRepo.CreateProduct(&product); err != nil {
			log.Println(err)
		}
	}
}
