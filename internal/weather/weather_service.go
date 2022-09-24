package weather

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-assignment3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment3/internal/weather/impl"
)

type WeatherService interface {
	UpdateValue(ctx context.Context) error
	GetReport(ctx context.Context) (*dto.WeatherResponse, error)
}

func ProvideWeatherService(DB *sql.DB) WeatherService {
	repo := impl.ProvideWeatherRepository(DB)
	return impl.ProvideWeatherService(repo)
}
