package main

// github.com/EndlessCheng/codeforces-go
func findLexSmallestString(s string, A int, shift int) (ans string) {
	add := func(b, inc byte) byte { return '0' + (b&15+inc)%10 }

	n, a, up := len(s), byte(A), byte(shift&1*9)
	ans = s
	for range s {
		for i := byte(0); i < 10; i++ { // 奇
			for i2 := byte(0); i2 <= up; i2++ { // 偶
				t := []byte(s)
				for j, b := range t {
					if j&1 > 0 {
						t[j] = add(b, i*a)
					} else {
						t[j] = add(b, i2*a)
					}
				}
				ans = min(ans, string(t))
			}
		}
		s = s[n-shift:] + s[:n-shift]
	}
	return
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
