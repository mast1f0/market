package seed

import (
	"context"
	"log"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

func SeedCategories(db ports.CategoryRepository) {
	ctx := context.Background()
	categories := []domain.Category{
		{Name: "Электроника"},
		{Name: "Одежда"},
		{Name: "Книги"},
		{Name: "Спорт"},
		{Name: "Игры"},
		{Name: "Дом"},
	}

	for _, c := range categories {
		_, err := db.CreateCategory(ctx, c.Name)
		if err != nil {
			log.Println(err)
		}
	}
}
