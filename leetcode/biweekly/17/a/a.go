package main

// github.com/EndlessCheng/codeforces-go
func decompressRLElist(a []int) (ans []int) {
	for i := 0; i < len(a); i += 2 {
		for j := 0; j < a[i]; j++ {
			ans = append(ans, a[i+1])
		}
	}
	return
}
