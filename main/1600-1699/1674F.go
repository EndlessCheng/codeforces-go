package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1674F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, c, x, y int
	Fscan(in, &n, &m, &q)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		c += bytes.Count(a[i], []byte{'*'})
	}

	ans := c
	for i, r := range a {
		j := c / n
		if i < c%n {
			j++
		}
		ans -= bytes.Count(r[:j], []byte{'*'})
	}

	for ; q > 0; q-- {
		Fscan(in, &x, &y)
		x--
		y--
		p := y*n + x
		if a[x][y] == '*' {
			c--
			if p > c {
				ans--
			}
			if p != c && a[c%n][c/n] == '*' {
				ans++
			}
		} else {
			if p > c {
				ans++
			}
			if p != c && a[c%n][c/n] == '*' {
				ans--
			}
			c++
		}
		Fprintln(out, ans)
		a[x][y] ^= 4
	}
}

//func main() { CF1674F(os.Stdin, os.Stdout) }
