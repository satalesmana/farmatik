package database

import (
	"database/sql"
	"farmatik/app/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn   *sql.DB
	err      error
	Basepath string
)

func init() {
	cfg := config.GetConfig()
	user := cfg.Database.Username
	password := cfg.Database.Password
	host := cfg.Database.Host
	port := cfg.Database.Port
	dbname := cfg.Database.Name

	log.Printf("Connecting to database " + host + " ...")
	dbConn, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err.Error())
	}
}

func GetConfig() {
	panic("unimplemented")
}

func GetCoon() *sql.DB {
	return dbConn
}
