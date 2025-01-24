package main

// github.com/EndlessCheng/codeforces-go
const mod = 1_000_000_007

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	g := make([][]int, n)
	fac := 1
	for i := 1; i < n; i++ {
		p := prevRoom[i]
		g[p] = append(g[p], i)
		fac = fac * (i + 1) % mod
	}

	mul := 1
	var dfs func(int) int
	dfs = func(x int) int {
		size := 1
		for _, y := range g[x] {
			size += dfs(y)
		}
		mul = mul * size % mod
		return size
	}
	dfs(0)

	return fac * pow(mul, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
