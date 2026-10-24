package main

// https://space.bilibili.com/206214
// 每月天数的前缀和（注：最后一个数可以省略）
var daysPreSum = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}

func calcDays(s string) int {
	month := (s[0]-'0')*10 + (s[1] - '0')
	day := (s[3]-'0')*10 + (s[4] - '0')
	return daysPreSum[month-1] + int(day)
}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	left := calcDays(max(arriveAlice, arriveBob)) // 直接比较字符串再计算
	right := calcDays(min(leaveAlice, leaveBob))
	return max(right-left+1, 0) // 答案不能为负数
}
