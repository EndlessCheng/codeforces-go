package main

// github.com/EndlessCheng/codeforces-go
func minOperations(s string) int {
	diff := 0
	for i, ch := range s {
		// 如果 i 是偶数，把 ch 变成 0，否则变成 1
		if int(ch-'0') != i%2 {
			diff++
		}
	}
	return min(diff, len(s)-diff)
}
