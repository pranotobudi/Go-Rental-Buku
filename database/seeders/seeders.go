package seeders

import (
	"github.com/pranotobudi/Go-Rental-Buku/database"
	"github.com/pranotobudi/Go-Rental-Buku/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.CategoryFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	database.InitDBTable()
	database.CategoryDataSeedInit(10)
	database.BookDataSeedInit(20)
	database.UserDataSeedInit(5)
	database.LoanAndFinePaymentDataSeedInit(30)

	// for _, seed := range RegisterSeeders(db) {
	// 	err := db.Debug().Create(seed.Seeder).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
