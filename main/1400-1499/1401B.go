package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1401B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, a, b, c, x, y, z int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &x, &y, &z)
		ans := min(c, y) * 2
		cnt := c - min(c, y) + a
		if cnt < z {
			ans -= (z - cnt) * 2
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1401B(os.Stdin, os.Stdout) }
