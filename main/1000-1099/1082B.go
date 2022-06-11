package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1082B(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var n, c, ans int
	var s string
	Fscan(bufio.NewReader(in), &n, &s)

	s = "S" + s
	a := []int{}
	for i := range s {
		c++
		if i == len(s)-1 || s[i] != s[i+1] {
			if s[i] == 'G' {
				ans = max(ans, c)
			}
			a = append(a, c)
			c = 0
		}
	}

	for i, n := 3, len(a); i < n; i += 2 {
		if a[i-1] == 1 {
			s := a[i-2] + a[i]
			if n >= 6 {
				s++
			}
			ans = max(ans, s)
		} else {
			ans = max(ans, max(a[i], a[i-2])+1)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1082B(os.Stdin, os.Stdout) }
