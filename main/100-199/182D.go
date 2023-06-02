package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF182D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	f := func(s string) (string, int) {
		n := len(s)
		match := make([]int, n)
		for i, c := 1, 0; i < n; i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		if m := match[n-1]; m > 0 && n%(n-m) == 0 {
			return s[:n-m], n / (n - m)
		}
		return s, 1
	}

	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	s, cntS := f(s)
	t, cntT := f(t)
	ans := 0
	if s == t {
		g := gcd(cntS, cntT)
		for d := 1; d*d <= g; d++ {
			if g%d == 0 {
				ans++
				if d*d < g {
					ans++
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF182D(os.Stdin, os.Stdout) }
