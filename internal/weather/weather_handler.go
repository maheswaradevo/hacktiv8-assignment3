package weather

import (
	"github.com/gorilla/mux"
)

type weatherHandler struct {
	r  *mux.Router
	ws WeatherService
}

func (w *weatherHandler) InitHandler() {
	// routes := w.r.NewRoute().PathPrefix(constant.WEATHER_API_PATH).Subrouter()

	//Weather
}

func ProvideWeatherHandler(r *mux.Router, ws WeatherService) *weatherHandler {
	return &weatherHandler{r: r, ws: ws}
}
