package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2029F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		r, b := false, false
		for i := range n {
			if s[i] == s[(i+1)%n] {
				if s[i] == 'R' {
					r = true
				} else {
					b = true
				}
			}
		}

		if r && b {
			Fprintln(out, "NO")
			continue
		}
		if n <= 2 {
			Fprintln(out, "YES")
			continue
		}
		if !r && !b {
			Fprintln(out, "NO")
			continue
		}

		target := byte('R')
		if r {
			target = 'B'
		}
		p := []int{}
		for i := range n {
			if s[i] == target {
				p = append(p, i)
			}
		}

		m := len(p)
		c := 0
		for i := range m {
			c += (p[(i+1)%m] - p[i] + n) % n % 2
		}

		if c == 1 || c == 0 && m < 2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2029F(bufio.NewReader(os.Stdin), os.Stdout) }
