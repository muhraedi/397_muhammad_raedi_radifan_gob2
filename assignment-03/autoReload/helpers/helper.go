package helpers

func CheckWaterStatus(water int) string {
	var waterStatus string
	if water <= 5 {
		waterStatus = "Aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Siaga"
	} else if water > 8 {
		waterStatus = "Bahaya"
	}
	return waterStatus
}

func CheckWindStatus(wind int) string {
	var windStatus string
	if wind <= 6 {
		windStatus = "Aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Siaga"
	} else if wind > 15 {
		windStatus = "Bahaya"
	}
	return windStatus
}
