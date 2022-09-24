package dto

import "github.com/maheswaradevo/hacktiv8-assignment3/internal/entity"

type WeatherResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Wind  int8 `json:"wind"`
	Water int8 `json:"water"`
}

func CreateWeatherReport(s entity.StatusReport) *WeatherResponse {
	return &WeatherResponse{
		Status: s.Status,
		Data: Data{
			Wind:  s.Weather.Wind,
			Water: s.Weather.Water,
		},
	}
}

func CreateWeatherResponse(s entity.StatusReport) *WeatherResponse {
	res := *CreateWeatherReport(s)
	return &res
}
