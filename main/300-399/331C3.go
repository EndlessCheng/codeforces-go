package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf331C3(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ x, y int }
	memo := map[pair]pair{}

	var dfs func(int, int) pair
	dfs = func(d, n int) pair {
		if n < 10 {
			first := 0
			if d != 0 || n != 0 {
				first = 1
			}
			return pair{first, n - max(d, n)}
		}

		key := pair{d, n}
		if p, ok := memo[key]; ok {
			return p
		}

		v := 1
		for v <= n/10 {
			v *= 10
		}

		a := dfs(max(d, n/v), n%v)
		b := dfs(d, n/v*v+a.y)
		memo[key] = pair{a.x + b.x, b.y}
		return memo[key]
	}

	Fprint(out, dfs(0, n).x)
}

//func main() { cf331C3(os.Stdin, os.Stdout) }
