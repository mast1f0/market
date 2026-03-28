package seed

import (
	"log"
	"market/internal/adapters/storage/orm"
	"market/internal/core/domain"
)

func SeedCategories(db *orm.Storage) {
	categories := []domain.Category{
		{Name: "Электроника"},
		{Name: "Одежда"},
		{Name: "Книги"},
		{Name: "Спорт"},
		{Name: "Игры"},
		{Name: "Дом"},
	}

	for _, c := range categories {
		_, err := db.CreateCategory(&c)
		if err != nil {
			log.Println(err)
		}
	}
}
