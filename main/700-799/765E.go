package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF765E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, rt int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(int, int) int
	f = func(v, fa int) int {
		s := map[int]struct{}{}
		mi, mx := int(1e9), 0
		for _, w := range g[v] {
			if w != fa {
				l := f(w, v)
				if l < 0 {
					return -1
				}
				l++
				s[l] = struct{}{}
				if l < mi {
					mi = l
				}
				if l > mx {
					mx = l
				}
			}
		}
		if len(s) == 0 {
			return 0
		}
		if len(s) == 1 {
			return mi
		}
		if len(s) == 2 && fa == 0 {
			return mi + mx
		}
		rt = v
		return -1
	}
	ans := f(1, 0)
	if ans < 0 {
		ans = f(rt, 0)
	}
	Fprint(out, ans>>bits.TrailingZeros(uint(ans)))
}

//func main() { CF765E(os.Stdin, os.Stdout) }
