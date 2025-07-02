package main

import (
	. "fmt"
	"index/suffixarray"
	"io"
	"math/bits"
	"unsafe"
)

// https://github.com/EndlessCheng
func cf822E(in io.Reader, out io.Writer) {
	var n, m, x int
	var s, t string
	Fscan(in, &n, &s, &m, &t, &x)

	s += "#" + t
	N := len(s)
	type _tp struct {
		_  []byte
		sa []int32
	}
	sa := (*_tp)(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	rank := make([]int, N)
	for i, p := range sa {
		rank[p] = i
	}
	height := make([]int, N)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < N && j+h < N && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}
	st := make([][18]int, N)
	for i, v := range height {
		st[i][0] = v
	}
	for j := 1; 1<<j <= N; j++ {
		for i := range N - 1<<j + 1 {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
	lcp := func(i, j int) int {
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		return _q(ri+1, rj+1)
	}

	// f[i] 表示如果要用 s[i:] 匹配，对应的 t[j:] 中的 j 最大是多少
	f := make([]int, n)
	for range x {
		nf := make([]int, n+1)
		for i, j := range f {
			l := lcp(i, n+1+j)
			nf[i+l] = max(nf[i+l], j+l)
			nf[i+1] = max(nf[i+1], nf[i])
		}
		if nf[n] == m {
			Fprint(out, "YES")
			return
		}
		f = nf[:n]
	}
	Fprint(out, "NO")
}

//func main() { cf822E(bufio.NewReader(os.Stdin), os.Stdout) }
