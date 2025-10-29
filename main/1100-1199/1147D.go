package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type uf47 struct {
	fa, dis []int
	cc      int
}

func newUnionFind47(n int) uf47 {
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf47{fa, dis, n}
}

func (u uf47) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] ^= u.dis[u.fa[x]]
		u.fa[x] = root
	}
	return u.fa[x]
}

func (u *uf47) merge(from, to, value int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return u.dis[from]^u.dis[to] == value
	}
	u.dis[x] = value ^ u.dis[to] ^ u.dis[from]
	u.fa[x] = y
	u.cc--
	return true
}

func cf1147D(in io.Reader, out io.Writer) {
	const mod = 998244353
	s := ""
	Fscan(in, &s)
	n := len(s)
	ans := 0
o:
	for st := 1; st < n; st++ {
		u := newUnionFind47(n * 2)
		for i := range n / 2 {
			u.merge(i, n-1-i, 0)
		}
		for i := 1; i < st; i++ {
			if s[i] != '?' && !u.merge(i, 0, int('1'-s[i])) {
				break o
			}
		}
		u.merge(n+st, 0, 0)
		for i, j := st, n-1; i < j; i, j = i+1, j-1 {
			u.merge(n+i, n+j, 0)
		}
		for i := st; i < n; i++ {
			if s[i] != '?' && !u.merge(i, n+i, int(s[i]-'0')) {
				continue o
			}
		}
		m := 1
		for range u.cc - st - 1 {
			m = m * 2 % mod
		}
		ans += m
	}
	Fprint(out, ans%mod)
}

//func main() { cf1147D(os.Stdin, os.Stdout) }
