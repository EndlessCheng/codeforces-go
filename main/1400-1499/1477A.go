package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1477A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n int
	var s, x0, v int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &x0)
		s -= x0
		g := int64(0)
		for n--; n > 0; n-- {
			Fscan(in, &v)
			g = gcd(g, v-x0)
		}
		if g == 0 || s%g != 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1477A(os.Stdin, os.Stdout) }
