package main

// https://space.bilibili.com/206214
func minExtraChar(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1 // 不选
		for j := 0; j <= i; j++ { // 枚举选哪个
			if has[s[j:i+1]] {
				f[i+1] = min(f[i+1], f[j])
			}
		}
	}
	return f[n]
}

func min(a, b int) int { if b < a { return b }; return a }
