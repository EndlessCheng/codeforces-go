package main

// github.com/EndlessCheng/codeforces-go
func decode(encoded []int, first int) (ans []int) {
	ans = make([]int, len(encoded)+1)
	ans[0] = first
	for i, e := range encoded {
		ans[i+1] = ans[i] ^ e
	}
	return
}
