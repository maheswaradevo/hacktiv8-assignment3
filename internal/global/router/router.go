package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment3/internal/weather"
)

func Init(router *mux.Router, db *sql.DB) {
	webRouter := router.NewRoute().PathPrefix("/api/v1").Subrouter()

	weatherService := weather.ProvideWeatherService(db)
	weatherHandler := weather.ProvideWeatherHandler(webRouter, weatherService)
	weatherHandler.InitHandler()
}
