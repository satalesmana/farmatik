package userLoginModel

import (
	"database/sql"
	db "farmatik/app/database"
)

type Handler interface {
	Insert(data *UserLogin) (string, error)
	GetByid(id string) (UserLogin, error)
	Update(id, status string) (string, error)
}

type UserLogin struct {
	ID     string `json:"id"`
	IDUser string `json:"id_user,omitempty"`
	Status string `json:"status,omitempty"`
	Token  string `json:"token,omitempty"`
}

type Db struct {
	database *sql.DB
}

func NewUserLoginHandler() Handler {
	return &Db{database: db.GetCoon()}
}

func (db *Db) Insert(data *UserLogin) (string, error) {
	queryUpdate := `UPDATE user_login SET status='N' WHERE  id_user=?`
	_, errUpdate := db.database.Exec(queryUpdate, &data.IDUser)

	if errUpdate != nil {
		return "", errUpdate
	}

	query := `INSERT INTO user_login(
		id,id_user, status, token  
	) VALUES(?, ?, ?, ?)`

	_, err := db.database.Exec(query,
		&data.ID,
		&data.IDUser,
		&data.Status,
		&data.Token,
	)

	if err != nil {
		return "", err
	}

	return data.ID, nil
}
func (db *Db) GetByid(id string) (UserLogin, error) {

	query := `SELECT id,id_user, status, token  FROM user_login WHERE id=?`
	var output UserLogin

	res, err := db.database.Query(query, id)
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		if err := res.Scan(&output.ID, &output.IDUser, &output.Status, &output.Token); err != nil {
			return output, err
		}
	}

	return output, nil
}

func (db *Db) Update(id, status string) (string, error) {
	query := `UPDATE user_login SET status=? WHERE  id=?`
	_, err := db.database.Exec(query, status, id)
	if err != nil {
		return "", err
	}

	return "Data berhasil perbaharui", nil
}
