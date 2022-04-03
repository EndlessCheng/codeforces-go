package main

// github.com/EndlessCheng/codeforces-go
func parseTime(s string) int {
	return (int(s[0]&15)*10+int(s[1]&15))*60 + int(s[3]&15)*10 + int(s[4]&15)
}

func convertTime(current, correct string) (ans int) {
	diff := parseTime(correct) - parseTime(current)
	if diff < 0 {
		diff += 1440
	}
	for _, inc := range []int{60, 15, 5, 1} { // 从大往小贪心
		ans += diff / inc
		diff %= inc
	}
	return
}
