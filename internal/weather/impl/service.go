package impl

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/maheswaradevo/hacktiv8-assignment3/internal/dto"
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
				windValue := randWindValue()
				waterValue := randWaterValue()
				report := reportWeather(windValue, waterValue)
				err := w.repo.UpdateValue(windValue, waterValue, report)
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
func reportWeather(windValue, waterValue int8) string {
	report := ""
	if windValue < 6 || waterValue < 6 {
		temp := "aman"
		report = temp
	} else if (windValue < 15 && windValue >= 7) || (waterValue < 8 && waterValue >= 6) {
		temp := "siaga"
		report = temp
	} else if (windValue > 15) || (waterValue > 8) {
		temp := "bahaya"
		report = temp
	}
	return report
}

func randWindValue() int8 {
	rand.Seed(time.Now().UnixMicro())
	res := rand.Intn((100 - 1) + 1)
	return int8(res)
}

func randWaterValue() int8 {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn((100 - 1) + 1)
	return int8(res)
}
