package main

import (
	"sort"
	"strconv"
)

// 排序：正数从小到大，负数从大到小

// github.com/EndlessCheng/codeforces-go
func smallestNumber(num int64) int64 {
	s := []byte(strconv.FormatInt(num, 10))
	if num <= 0 {
		t := s[1:]
		sort.Slice(t, func(i, j int) bool { return t[i] > t[j] }) // 负数，从大到小排序（兼容 num=0）
		ans, _ := strconv.ParseInt(string(s), 10, 64)
		return ans
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 正数，从小到大排序
	i := 0
	for s[i] == '0' { // 如果有前导零就从后面找一个数字
		i++
	}
	s[i], s[0] = s[0], s[i] // 交换，保证没有前导零
	ans, _ := strconv.ParseInt(string(s), 10, 64)
	return ans
}
