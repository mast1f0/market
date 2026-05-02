package orm

import (
	"fmt"
	"market/internal/engine/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB_USER, cfg.DB_PASSWORD,
		cfg.DB_HOST, cfg.DB_PORT,
		cfg.DB_NAME,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < 10; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}
	return db
}
