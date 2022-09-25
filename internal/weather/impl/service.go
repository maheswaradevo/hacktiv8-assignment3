package impl

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/maheswaradevo/hacktiv8-assignment3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment3/internal/utils"
)

type weatherServiceImpl struct {
	repo WeatherRepository
}

func ProvideWeatherService(repo WeatherRepository) *weatherServiceImpl {
	return &weatherServiceImpl{repo: repo}
}

func (w weatherServiceImpl) GetReport(ctx context.Context) (*dto.WeatherResponse, error) {
	res, err := w.repo.GetReport(ctx)
	if err != nil {
		log.Printf("[GetReport] failed to get the weather report")
		return nil, err
	}
	return dto.CreateWeatherResponse(res), nil
}

func (w weatherServiceImpl) UpdateValue(ctx context.Context) error {
	ticker := time.NewTicker(15 * time.Minute)
	quit := make(chan int)

	go func() {
		for {
			select {
			case <-ticker.C:
				windValue := utils.RandWindValue()
				waterValue := utils.RandWaterValue()
				windString := strconv.Itoa(int(windValue)) + " m/s"
				waterString := strconv.Itoa(int(waterValue)) + " m"
				report := utils.ReportWeather(windValue, waterValue)

				err := w.repo.UpdateValue(windString, waterString, report)
				log.Printf("[UpdateValue] value updated!")
				if err != nil {
					log.Printf("[UpdateValue] failed to update the value")
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}
