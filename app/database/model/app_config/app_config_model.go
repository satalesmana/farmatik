package appconfig

import (
	"database/sql"
	db "farmatik/app/database"
	"fmt"
	"strings"
)

type Handler interface {
	Update(data *AppConfig) (int64, error)
	GetById(id AppConfigRequest) ([]AppConfig, error)
}
type AppConfigRequest struct {
	ID []string `json:"id,omitempty"`
}

type AppConfig struct {
	ID          string `json:"id,omitempty"`
	Keterangan  string `json:"keterangan,omitempty"`
	ConfigValue string `json:"config_value,omitempty"`
}

type Db struct {
	database *sql.DB
}

func NewConfigHandler() Handler {
	return &Db{database: db.GetCoon()}
}

func (db *Db) Update(data *AppConfig) (int64, error) {
	return 0, nil
}

func (db *Db) GetById(request AppConfigRequest) ([]AppConfig, error) {
	var data []AppConfig
	var idConfig string
	for _, v := range request.ID {
		idConfig += fmt.Sprintf("%s','", v)
	}
	idConfig = strings.TrimSuffix(idConfig, "','")

	query := `SELECT id, keterangan, configValue FROM app_config WHERE id IN('` + idConfig + `')`
	res, err := db.database.Query(query)
	if err != nil {
		return data, err
	}
	defer res.Close()

	for res.Next() {
		var config AppConfig
		if err := res.Scan(
			&config.ID,
			&config.Keterangan,
			&config.ConfigValue,
		); err != nil {
			return data, err
		}
		data = append(data, config)
	}

	return data, nil
}
