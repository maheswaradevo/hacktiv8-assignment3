package weather

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment3/internal/constant"
	"github.com/maheswaradevo/hacktiv8-assignment3/internal/utils"
)

type weatherHandler struct {
	r  *mux.Router
	ws WeatherService
}

func (w *weatherHandler) InitHandler() {
	routes := w.r.NewRoute().PathPrefix(constant.WEATHER_API_PATH).Subrouter()

	//Weather
	routes.HandleFunc("", w.updateValue()).Methods(http.MethodPut)
	routes.HandleFunc("/report", w.getReport()).Methods(http.MethodGet)
}

func ProvideWeatherHandler(r *mux.Router, ws WeatherService) *weatherHandler {
	return &weatherHandler{r: r, ws: ws}
}

func (w *weatherHandler) updateValue() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := w.ws.UpdateValue(r.Context())
		if err != nil {
			utils.NewErrorResponse(rw, err)
		}
		utils.NewSuccessResponsWriter(rw, http.StatusCreated, "Updated", nil)
	}
}

func (w *weatherHandler) getReport() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		res, err := w.ws.GetReport(r.Context())
		if err != nil {
			utils.NewErrorResponse(rw, err)
		}

		utils.NewSuccessResponsWriter(rw, http.StatusOK, "SUCCESS", res)
	}
}
