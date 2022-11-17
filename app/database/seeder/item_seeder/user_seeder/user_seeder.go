package userSeeder

import (
	"database/sql"
	db "farmatik/app/database"
	"log"
	"strconv"
)

type Handler interface {
	UserSeeder()
}

type uscase struct {
	database *sql.DB
}

func NewSeederHandler() Handler {
	return &uscase{
		database: db.GetCoon(),
	}
}

func (uc *uscase) UserSeeder() {
	query := `REPLACE INTO user(id, name, email, password) 
	VALUES('1','Administrator', 'admin@mail.com', '$2a$08$5Bdb1KZUQxhrUvllhcoE8.zGj8d7.8puyi.EJvMqyyzouUFS5nZOO')`
	// default password => admin123#@!

	res, err := uc.database.Exec(query)
	if err != nil {
		log.Println(err)
	}

	intRes, errAft := res.RowsAffected()
	if err != nil {
		log.Println(errAft)
	}

	strIDUser := strconv.Itoa(int(intRes))
	log.Println("[USER SEEDER] " + strIDUser + " Rows Affected ")

}
