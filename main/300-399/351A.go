package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF351A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, v, s, c int
	Fscanln(in, &n)
	for i := 0; i < n*2; i++ {
		Fscanf(in, "%d.%d", &v, &v)
		if v > 0 {
			s += v
			c++
		}
	}
	ans := int(1e9)
	for i := max(c-n, 0); i <= min(n, c); i++ {
		ans = min(ans, abs(i*1000-s))
	}
	Fprintf(out, "%d.%03d", ans/1000, ans%1000)
}

//func main() { CF351A(os.Stdin, os.Stdout) }
