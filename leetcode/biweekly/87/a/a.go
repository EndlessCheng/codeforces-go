package main

// https://space.bilibili.com/206214
var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func calcDays(s string) (day int) {
	for _, d := range days[:s[0]&15*10+s[1]&15-1] {
		day += d
	}
	return day + int(s[3]&15*10+s[4]&15)
}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	right := calcDays(min(leaveAlice, leaveBob))
	left := calcDays(max(arriveAlice, arriveBob))
	return max(right - left + 1, 0)
}
