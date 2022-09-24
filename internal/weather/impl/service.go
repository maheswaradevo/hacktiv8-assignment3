package impl

type weatherServiceImpl struct {
	repo WeatherRepository
}

func ProvideWeatherService(repo WeatherRepository) *weatherServiceImpl {
	return &weatherServiceImpl{repo: repo}
}
