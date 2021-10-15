package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1555C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := int(2e9)
		s0, s1 := 0, 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			s0 += a[i]
		}
		for _, v := range a {
			s0 -= v
			ans = min(ans, max(s0, s1))
			Fscan(in, &v)
			s1 += v
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1555C(os.Stdin, os.Stdout) }
