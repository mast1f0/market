package seed

import (
	"log"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

func SeedCategories(db ports.CategoryRepository) {
	categories := []domain.Category{
		{Name: "Электроника"},
		{Name: "Одежда"},
		{Name: "Книги"},
		{Name: "Спорт"},
		{Name: "Игры"},
		{Name: "Дом"},
	}

	for _, c := range categories {
		_, err := db.CreateCategory(c.Name)
		if err != nil {
			log.Println(err)
		}
	}
}
