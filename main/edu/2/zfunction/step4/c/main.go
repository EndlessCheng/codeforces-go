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
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		z := make([]int, n)
		d := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
			d[0]++
			d[z[i]]--
		}
		c := 1
		for _, v := range d {
			c += v
			Fprint(out, c, " ")
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
