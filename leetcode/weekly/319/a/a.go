package main

// https://space.bilibili.com/206214
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}
