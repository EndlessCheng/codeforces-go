package main

// github.com/EndlessCheng/codeforces-go
func findLexSmallestString(s string, A int, shift int) (ans string) {
	n := len(s)
	ans = s
	add := func(b, inc byte) byte { return '0' + (b&15+inc)%10 }

	a := byte(A)
	if shift&1 > 0 {
		for k := 0; k < n; k++ {
			for i := byte(0); i < 10; i++ {
				for i2 := byte(0); i2 < 10; i2++ {
					t := []byte(s)
					for j, b := range t {
						if j&1 == 0 {
							t[j] = add(b, i*a)
						} else {
							t[j] = add(b, i2*a)
						}
					}
					ans = min(ans, string(t))
				}
			}
			s = s[len(s)-shift:] + s[:len(s)-shift]
		}
	} else {
		for k := 0; k < n; k++ {
			for i := byte(0); i < 10; i++ {
				t := []byte(s)
				for j, v := range t {
					if j&1 == 1 {
						t[j] = add(v, i*a)
					}
				}
				ans = min(ans, string(t))
			}
			s = s[len(s)-shift:] + s[:len(s)-shift]
		}
	}
	return
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
