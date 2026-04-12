package main

// https://space.bilibili.com/206214
func trafficSignal(timer int) string {
	if timer == 0 {
		return "Green"
	}
	if timer == 30 {
		return "Orange"
	}
	if 30 < timer && timer <= 90 {
		return "Red"
	}
	return "Invalid"
}
