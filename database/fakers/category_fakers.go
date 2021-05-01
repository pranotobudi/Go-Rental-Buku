package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pranotobudi/Go-Rental-Buku/database"
	"gorm.io/gorm"
)

func CategoryFaker(db *gorm.DB) *database.Category {
	return &database.Category{
		Name: faker.Name(),
	}
}
