package seeder

import (
	configSeeder "farmatik/app/database/seeder/item_seeder/config_seeder"
	userSeeder "farmatik/app/database/seeder/item_seeder/user_seeder"
)

func Seed() {
	configSeeder.NewSeederHandler().AppConfigSeeder()
	userSeeder.NewSeederHandler().UserSeeder()
}
