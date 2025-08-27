package main

// github.com/EndlessCheng/codeforces-go
func appealSum(s string) (ans int64) {
	last := [26]int{}
	for i := range last {
		last[i] = -1
	}
	n := len(s)
	for i, c := range s {
		c -= 'a'
		ans += int64(i-last[c]) * int64(n-i)
		last[c] = i
	}
	return
}

func appealSum2(s string) (ans int64) {
	last := [26]int{}
	for i := range last {
		last[i] = -1 // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
	}
	sumG := 0
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c]
		ans += int64(sumG)
		last[c] = i
	}
	return
}
