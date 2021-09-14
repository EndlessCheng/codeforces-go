package main

// github.com/EndlessCheng/codeforces-go
const mx, mod int = 1e3, 1e9 + 7

var F, invF [mx + 1]int

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func init() {
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

type uf struct{ fa []int }

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}
func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}
func (u uf) merge(from, to int) { u.fa[u.find(from)] = u.find(to) }

func numOfWays(a []int) int {
	n := len(a)
	vis := make([]bool, n+2)
	u := newUnionFind(n + 1)
	g := make([][2]int, n+1)
	for i := len(a) - 1; i >= 0; i-- {
		fa := a[i]
		// BST 的中序遍历就是 1~n，对于数组 a，必定是先访问 fa-1 和 fa+1 才会访问到 fa
		// 这样就可以倒序遍历 a，用并查集建树，注意 merge 的方向
		if vis[fa-1] {
			son := u.find(fa - 1)
			g[fa][0] = son
			u.merge(son, fa)
		}
		if vis[fa+1] {
			son := u.find(fa + 1)
			g[fa][1] = son
			u.merge(son, fa)
		}
		vis[fa] = true
	}
	ans := 1
	var f func(int) int
	f = func(v int) int {
		if v == 0 {
			return 0
		}
		l, r := f(g[v][0]), f(g[v][1])
		if l > 0 && r > 0 {
			ans = ans * F[l+r] % mod * invF[l] % mod * invF[r] % mod
		}
		return 1 + l + r
	}
	f(a[0])
	return (ans + mod - 1) % mod
}
