package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
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

	var s, t string
	Fscan(in, &s, &t)
	if len(s) != len(t) {
		Fprint(out, "No")
		return
	}
	s = t + s

	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		if i == len(t) && i+z[i] == n {
			Fprint(out, "Yes\n0")
			return
		}
		if i > len(t) {
			if s[i-1] != s[n-i] {
				Fprint(out, "No") // or break
				return
			}
			if i+z[i] == n {
				Fprint(out, "Yes\n", i-len(t))
				return
			}
		}
	}
	Fprint(out, "No")
}

func main() { run(os.Stdin, os.Stdout) }
