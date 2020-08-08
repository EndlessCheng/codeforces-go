package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
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
	f := func(s1, s2 string) string {
		s := s1 + s2
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
			if i >= len(s1) && i+z[i] == n {
				return s2 + s1[z[i]:]
			}
		}
		return s
	}

	var T int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		if len(s) < len(t) {
			s, t = t, s
		}
		if strings.Contains(s, t) { // 可以不需要，而是在 f 中额外添加一些判定条件
			Fprintln(out, s)
			continue
		}
		ans := f(s, t)
		if s := f(t, s); len(s) < len(ans) {
			ans = s
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
