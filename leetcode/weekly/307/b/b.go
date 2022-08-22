package main

import "strings"

// https://space.bilibili.com/206214
func largestPalindromic(num string) string {
	cnt := ['9' + 1]int{}
	for _, d := range num {
		cnt[d]++
	}
	if cnt['0'] == len(num) { // 特判最特殊的情况：num 全是 0
		return "0"
	}

	s := []byte{}
	for i := '9'; i > '0' || i == '0' && len(s) > 0; i-- { // 如果填了数字，则可以填 0
		s = append(s, strings.Repeat(string(i), cnt[i]/2)...)
	}

	j := len(s) - 1
	for i := byte('9'); i >= '0'; i-- {
		if cnt[i]&1 > 0 { // 还可以填一个变成奇回文串
			s = append(s, i)
			break
		}
	}
	for ; j >= 0; j-- { // 添加镜像部分
		s = append(s, s[j])
	}
	return string(s)
}
