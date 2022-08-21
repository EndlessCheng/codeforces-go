package main

// https://space.bilibili.com/206214
func largestPalindromic(num string) string {
	cnt := ['9' + 1]int{}
	for _, d := range num {
		cnt[d]++
	}
	if cnt['0'] == len(num) { // 特判
		return "0"
	}

	s := []byte{}
	for i := byte('9'); i > '0'; i-- {
		for j := 0; j < cnt[i]/2; j++ {
			s = append(s, i)
		}
	}
	// 如果填了数字，则可以填 0
	if len(s) > 0 {
		for j := 0; j < cnt['0']/2; j++ {
			s = append(s, '0')
		}
	}

	j := len(s) - 1
	for i := byte('9'); i >= '0'; i-- {
		if cnt[i]&1 > 0 { // 还可以填一个变成奇回文串
			s = append(s, i)
			break
		}
	}
	for ; j >= 0; j-- { // 补充剩余部分
		s = append(s, s[j])
	}
	return string(s)
}
