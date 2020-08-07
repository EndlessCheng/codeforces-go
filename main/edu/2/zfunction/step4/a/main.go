package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
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

	var T int
	var s []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
			if z[i] == n-i {
				Fprintf(out, "%s\n", s[:i])
				continue o
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

func main() { run(os.Stdin, os.Stdout) }
