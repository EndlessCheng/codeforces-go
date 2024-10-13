package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1776L(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, a, b int
	var s string
	Fscan(in, &n, &s, &q)
	c := strings.Count(s, "+")
	c = min(c, n-c)

	d := n - c*2
	if d == 0 {
		Fprint(out, strings.Repeat("YES\n", q))
		return
	}

	for range q {
		Fscan(in, &a, &b)
		if a > b {
			a, b = b, a
		}
		if a < b && a*d%(b-a) == 0 && a*d/(b-a) <= c {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1776L(bufio.NewReader(os.Stdin), os.Stdout) }
