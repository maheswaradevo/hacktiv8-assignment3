package weather

import (
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-assignment3/internal/weather/impl"
)

type WeatherService interface {
}

func ProvideWeatherService(DB *sql.DB) WeatherService {
	repo := impl.ProvideWeatherRepository(DB)
	return impl.ProvideWeatherService(repo)
}
