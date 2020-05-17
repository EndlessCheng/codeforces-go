package main

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	for i, s := range startTime {
		if s <= queryTime && queryTime <= endTime[i] {
			ans++
		}
	}
	return
}
