package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type xorBasis24 struct {
	b     [60]int
	n, or int
}

func (b *xorBasis24) insert(v int) {
	b.or |= v
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.n++
			return
		}
		v ^= b.b[i]
	}
}

func cf724G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, m, ans int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	var b xorBasis24
	var cnt [60]int
	var cv int
	var dfs func(int, int)
	dfs = func(v, xor int) {
		dis[v] = xor
		cv++
		for i := range cnt {
			cnt[i] += xor >> i & 1
		}
		for _, e := range g[v] {
			w := e.to
			if dis[w] < 0 {
				dfs(w, xor^e.wt)
			} else {
				b.insert(xor ^ e.wt ^ dis[w])
			}
		}
	}
	for i, d := range dis {
		if d >= 0 {
			continue
		}
		b = xorBasis24{}
		cnt = [60]int{}
		cv = 0
		dfs(i, 0)
		for j, c := range cnt {
			if b.or>>j&1 > 0 {
				ans = (ans + cv*(cv-1)/2%mod*pow(2, j+b.n-1)) % mod
			} else {
				ans = (ans + c*(cv-c)%mod*pow(2, j+b.n)) % mod
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf724G(os.Stdin, os.Stdout) }
