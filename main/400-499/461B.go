package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF461B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	pow := func(x int64) int64 {
		res := int64(1)
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for v := 1; v < n; v++ {
		Fscan(in, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	isBlack := make([]bool, n)
	for i := range isBlack {
		Fscan(in, &isBlack[i])
	}
	var f func(int, int) (int64, int64)
	f = func(v, fa int) (notBlack, black int64) {
		if isBlack[v] {
			black = 1
			for _, w := range g[v] {
				if w != fa {
					nb, b := f(w, v)
					black = black * (nb + b) % mod
				}
			}
		} else {
			notBlack = 1
			for _, w := range g[v] {
				if w != fa {
					nb, b := f(w, v)
					notBlack = notBlack * (nb + b) % mod
					black = (black + b*pow(nb+b)) % mod
				}
			}
			black = black * notBlack % mod
		}
		return
	}
	_, ans := f(0, -1)
	Fprint(out, ans)
}

//func main() { CF461B(os.Stdin, os.Stdout) }
