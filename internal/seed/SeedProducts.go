package seed

import (
	"market/internal/adapters/storage/orm"
	"market/internal/core/domain"
	"math/rand"
)

func SeedProducts(db *orm.Storage) {

	categories := db.GetCategories()

	products := []string{
		"Смартфон", "Ноутбук", "Футболка", "Кроссовки",
		"Книга по Go", "Гантели", "PlayStation",
		"Стол", "Наушники", "Рюкзак",
	}

	images := []string{
		"https://picsum.photos/300",
		"https://picsum.photos/301",
		"https://picsum.photos/302",
		"https://picsum.photos/303",
	}

	for i := 0; i < 20; i++ {
		cat := categories[rand.Intn(len(categories))]

		product := domain.Product{
			Name:        products[rand.Intn(len(products))],
			Description: "Отличный товар для повседневного использования",
			Price:       float64(rand.Intn(10000)) / 100,
			CategoryID:  uint(cat.ID),
			ImageURL:    images[rand.Intn(len(images))],
		}

		db.CreateProduct(&product)
	}
}
