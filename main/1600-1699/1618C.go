package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1618C(_r io.Reader, _w io.Writer) {
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
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		g := [2]int64{}
		for i := range a {
			Fscan(in, &a[i])
			g[i&1] = gcd(g[i&1], a[i])
		}
		ok := [2]bool{true, true}
		for i, v := range a {
			ok[i&1] = ok[i&1] && v%g[i&1^1] > 0
		}
		if ok[0] {
			Fprintln(out, g[1])
		} else if ok[1] {
			Fprintln(out, g[0])
		} else {
			Fprintln(out, 0)
		}
	}
}

//func main() { CF1618C(os.Stdin, os.Stdout) }
