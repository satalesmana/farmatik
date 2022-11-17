package database

import (
	"farmatik/app/database/migration"
)

func Migrate() {
	db := GetCoon()
	migration.UserMigration(db)
	migration.AppConfig(db)
	migration.UserLoginMigration(db)
	migration.ProductMigration(db)
	migration.ProductHargaJualMigration(db)
	migration.ProductMutation(db)
	migration.PenjualanMigration(db)
	migration.PenjualanDetailMigration(db)
}
