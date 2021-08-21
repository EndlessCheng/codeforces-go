package main

// github.com/EndlessCheng/codeforces-go
func minTimeToType(s string) int {
	cur := 'a'
	ans := len(s)
	for _, b := range s {
		d := int(b - cur)
		if d < 0 {
			d = -d
		}
		ans += min(d, 26-d)
		cur = b
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
