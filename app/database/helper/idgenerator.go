package DBHelper

import (
	database "farmatik/app/database"
	"strconv"
)

func GenerateAutoId(id, key, table string) (string, error) {
	var maxID string
	var newId string
	var counter int = 1
	var maxLenId int64 = 10

	db := database.GetCoon()
	query := `SELECT IFNULL(MAX(` + id + `),'')  as kodeTerbesar FROM ` + table
	res, err := db.Query(query)
	if err != nil {
		return newId, err
	}
	defer res.Close()
	for res.Next() {
		if err := res.Scan(
			&maxID,
		); err != nil {
			return newId, err
		}
	}

	str := "1"
	if maxID != "" {
		dataLen := len(maxID)
		lastUrutan := maxID[len(key):dataLen]
		angka, _ := strconv.Atoi(lastUrutan)
		urutan := angka + counter
		str = strconv.Itoa(urutan)
	}

	maxLoop := maxLenId - (int64(len(str)) + int64(len(key)))
	numNuol := ""
	for i := int64(0); i < maxLoop; i++ {
		numNuol = numNuol + "0"
	}

	newId = key + numNuol + str
	return newId, nil
}
