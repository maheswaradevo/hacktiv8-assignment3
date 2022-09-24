package impl

import (
	"database/sql"
)

type weatherRepositoryImpl struct {
	DB *sql.DB
}

type WeatherRepository interface {
}

func ProvideWeatherRepository(DB *sql.DB) *weatherRepositoryImpl {
	return &weatherRepositoryImpl{DB: DB}
}
