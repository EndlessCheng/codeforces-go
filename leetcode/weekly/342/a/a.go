package main

// https://space.bilibili.com/206214
func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
