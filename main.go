package main

import (
	"farmatik/app/config"
	migrationdb "farmatik/app/database"
	seederApp "farmatik/app/database/seeder"
	"farmatik/app/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	argLength := len(os.Args[1:])

	//cek argumen
	if argLength > 0 {
		if os.Args[1] == "migrate" {
			migrationdb.Migrate()
		}

		if os.Args[1] == "seed" {
			seederApp.Seed()
		}

	}

	if argLength == 0 {
		cfg := config.GetConfig()
		r := gin.Default()
		routes := routes.Routes(r)
		routes.Run(cfg.App.Host + ":" + cfg.App.Port)
	}
}
