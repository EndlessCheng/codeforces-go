package main

// github.com/EndlessCheng/codeforces-go
const mod, mx int = 1e9 + 7, 1e5

var F [mx + 1]int

func init() {
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
}

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := prevRoom[w]
		g[v] = append(g[v], w)
	}
	mul := 1
	var f func(int) int
	f = func(v int) int {
		sz := 1
		for _, w := range g[v] {
			sz += f(w)
		}
		mul = mul * sz % mod
		return sz
	}
	f(0)
	return F[n] * pow(mul, mod-2) % mod
}

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
