package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf1778F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1000
	divisors := [mx + 1][]int{}
	for i := mx; i > 0; i-- {
		for j := i; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	ceilSqrt := [mx + 1]int{}
	calcCeilSqrt := func(x int) int {
		res := 1
		for x > 1 {
			p := lpf[x]
			for p2 := p * p; x%p2 == 0; x /= p2 {
				res *= p
			}
			if x%p == 0 {
				res *= p
				x /= p
			}
		}
		return res
	}
	for i := 1; i <= mx; i++ {
		ceilSqrt[i] = calcCeilSqrt(i)
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if k == 0 {
			Fprintln(out, a[0])
			continue
		}

		subGcd := make([]int, n)
		var dfs0 func(int, int)
		dfs0 = func(v, fa int) {
			subGcd[v] = a[v]
			for _, w := range g[v] {
				if w != fa {
					dfs0(w, v)
					subGcd[v] = gcd(subGcd[v], subGcd[w])
				}
			}
		}
		dfs0(0, -1)

		var dfs func(int, int, int) int
		dfs = func(v, fa, targetGcd int) int {
			if subGcd[v]%targetGcd == 0 {
				return 0
			}
			if subGcd[v]*subGcd[v]%targetGcd == 0 {
				return 1
			}
			if a[v]*a[v]%targetGcd > 0 {
				return 1e9
			}
			cnt := 1
			for _, w := range g[v] {
				if w != fa {
					cnt += dfs(w, v, ceilSqrt[targetGcd])
				}
			}
			return cnt
		}

		for _, d := range divisors[a[0]] {
			cnt := 0
			for _, v := range g[0] {
				cnt += dfs(v, 0, d)
			}
			if cnt < k {
				Fprintln(out, a[0]*d)
				break
			}
		}
	}
}

//func main() { cf1778F(bufio.NewReader(os.Stdin), os.Stdout) }
