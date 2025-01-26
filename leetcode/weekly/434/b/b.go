package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func countMentions(numberOfUsers int, events [][]string) []int {
	// 按照时间戳排序，时间戳相同的，离线事件排在前面
	slices.SortFunc(events, func(a, b []string) int {
		ta, _ := strconv.Atoi(a[1])
		tb, _ := strconv.Atoi(b[1])
		return cmp.Or(ta-tb, int(b[0][0])-int(a[0][0]))
	})

	ans := make([]int, numberOfUsers)
	onlineT := make([]int, numberOfUsers)
	for _, e := range events {
		curT, _ := strconv.Atoi(e[1])
		if e[0][0] == 'O' {
			i, _ := strconv.Atoi(e[2])
			onlineT[i] = curT + 60
		} else if e[2][0] == 'A' {
			for i := range ans {
				ans[i]++
			}
		} else if e[2][0] == 'H' {
			for i, t := range onlineT {
				if t <= curT { // 在线
					ans[i]++
				}
			}
		} else {
			for _, s := range strings.Split(e[2], " ") {
				i, _ := strconv.Atoi(s[2:])
				ans[i]++
			}
		}
	}
	return ans
}
