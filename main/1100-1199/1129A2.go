package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1129A2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, x, y, maxR int
	Fscan(in, &n, &m)
	g := make([]struct{ base, extra int }, n+1)
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		if y < x {
			y += n
		}
		if g[x].base == 0 || y < g[x].extra {
			g[x].extra = y
		}
		maxR = max(maxR, g[x].base+g[x].extra)
		g[x].base += n
	}
	for i := 1; i <= n; i++ {
		Fprint(out, maxR-i, " ")
		maxR = max(maxR, g[i].base+g[i].extra)
	}
}

//func main() { CF1129A2(os.Stdin, os.Stdout) }
