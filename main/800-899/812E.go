package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF812E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, o int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		g[v-1] = append(g[v-1], w)
	}
	odd := []int{}
	even := map[int]int{}
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if len(g[v]) == 0 {
			o = d
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d^1)
			}
		}
		if d == o {
			odd = append(odd, a[v])
		} else {
			even[a[v]]++
		}
	}
	f(0, -1, 0)
	xor := 0
	for _, v := range odd {
		xor ^= v
	}
	ans := int64(0)
	if xor > 0 {
		for _, v := range odd {
			ans += int64(even[xor^v])
		}
	} else {
		co := int64(len(odd))
		ce := int64(n) - co
		ans = co*(co-1)/2 + ce*(ce-1)/2
		for _, v := range odd {
			ans += int64(even[v])
		}
	}
	Fprint(out, ans)
}

//func main() { CF812E(os.Stdin, os.Stdout) }
