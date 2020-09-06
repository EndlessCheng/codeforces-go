package main

// github.com/EndlessCheng/codeforces-go
func modifyString(s string) (ans string) {
	t := []byte(s)
	for i, b := range t {
		if b == '?' {
			t[i] = 'a'
			if i > 0 && t[i-1] == 'a' || i+1 < len(t) && t[i+1] == 'a' {
				t[i] = 'b'
				if i > 0 && t[i-1] == 'b' || i+1 < len(t) && t[i+1] == 'b' {
					t[i] = 'c'
				}
			}
		}
	}
	return string(t)
}
