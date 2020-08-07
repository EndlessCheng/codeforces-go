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

	var t string
	Fscan(in, &t)
	s := make([]byte, len(t), len(t)*2+1)
	copy(s, t)
	s = append(s, '#')
	for i := len(t) - 1; i >= 0; i-- {
		s = append(s, t[i])
	}

	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		if i > len(t) && i+z[i] == n {
			Fprint(out, z[i])
			break
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
