package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(s string) (ans int) {
	c := 0
	for _, b := range s {
		if b == '[' {
			c++
		} else if c > 0 {
			c--
		} else {
			c++ // 把最后面的 [ 和 ] 交换
			ans++
		}
	}
	return
}
