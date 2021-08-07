package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1010C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, k, v int
	Fscan(in, &n, &k)
	g := k
	for ; n > 0 && g > 1; n-- {
		Fscan(in, &v)
		g = gcd(g, v)
	}
	Fprintln(out, k/g)
	for i := 0; i < k; i += g {
		Fprint(out, i, " ")
	}
}

//func main() { CF1010C(os.Stdin, os.Stdout) }
