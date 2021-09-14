package main

// bugfree 的写法
func strWithout3a3b(a, b int) string {
	ans := make([]byte, 0, a+b)
	var ca, cb byte = 'a', 'b'
	if a < b {
		a, b = b, a
		ca, cb = cb, ca
	}
	for a > 0 || b > 0 {
		if a > 0 {
			ans = append(ans, ca)
			a--
		}
		if a > b {
			ans = append(ans, ca)
			a--
		}
		if b > 0 {
			ans = append(ans, cb)
			b--
		}
	}
	return string(ans)
}

// 较麻烦的写法
func strWithout3a3b_(a, b int) (s string) {
	ca, cb := "a", "b"
	if a < b {
		a, b = b, a
		ca, cb = cb, ca
	}
	for b > 0 && a > b {
		s += ca + ca + cb
		a -= 2
		b--
	}
	for a -= b; b > 0; b-- {
		s += "ab"
	}
	for ; a > 0; a-- {
		s += ca
	}
	return
}

