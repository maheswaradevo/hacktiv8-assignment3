package utils

func ReportWeather(windValue, waterValue int8) string {
	report := ""
	if windValue <= 6 {
		temp := "aman"
		report = temp
	} else if windValue >= 7 && windValue <= 15 {
		temp := "siaga"
		report = temp
	} else if windValue > 15 {
		temp := "bahaya"
		report = temp
	} else if waterValue < 5 {
		temp := "aman"
		report = temp
	} else if waterValue >= 6 && waterValue <= 8 {
		temp := "siaga"
		report = temp
	} else if waterValue > 8 {
		temp := "bahaya"
		report = temp
	}
	return report
}
