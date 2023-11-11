package database

import "github.com/reyhanmichiels/AquaFarmManagement/domain"

func Migrate() {
	DB.Migrator().DropTable(
		&domain.Farm{},
		&domain.Pond{},
		&domain.ApiCall{},
	)

	DB.AutoMigrate(
		&domain.Farm{},
		&domain.Pond{},
		&domain.ApiCall{},
	)
}
