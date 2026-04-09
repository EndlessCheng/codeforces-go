package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1612F(in io.Reader, out io.Writer) {
	var n, m, q int
	Fscan(in, &n, &m, &q)
	type pair struct{ x, y int }
	has := map[pair]bool{}
	for range q {
		var x, y int
		Fscan(in, &x, &y)
		if n > m {
			x, y = y, x
		}
		has[pair{x, y}] = true
	}
	if n > m {
		n, m = m, n
	}

	f := make([]int, n+1)
	f[1] = 1
	ans := 0
	for f[n] != m {
		ans++
		for i := n; i > 0; i-- {
			if f[i] == 0 {
				continue
			}
			p := i + f[i]
			if has[pair{i, f[i]}] {
				p++
			}
			f[min(n, p)] = max(f[min(n, p)], f[i])
			f[i] = min(p, m)
		}
	}
	Fprintln(out, ans)
}

//func main() { cf1612F(bufio.NewReader(os.Stdin), os.Stdout) }
