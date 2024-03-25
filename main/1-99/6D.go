package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf6D(in io.Reader, out io.Writer) {
	var n, fi, se int
	Fscan(in, &n, &fi, &se)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]++
	}

	k := (a[0]-1)/se + 1
	type pair struct{ i, c int }
	ans := []pair{{1, k}}
	a[1] -= k * fi
	a[2] -= k * se

	if a[n-1] > 0 {
		k = (a[n-1]-1)/se + 1
		ans = append(ans, pair{n - 2, k})
		a[n-2] -= k * fi
		a[n-3] -= k * se
	}

	dp := make([][17][17]int, n)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(i, preHit, preHp int) int {
		if i == n-1 {
			return 0
		}
		p := &dp[i][preHit][preHp]
		if *p != -1 {
			return *p
		}
		hp := a[i] - preHit*se
		hit := 0
		if preHp > 0 {
			hit = (preHp-1)/se + 1
			hp -= hit * fi
		}
		res := int(1e9)
		for {
			res = min(res, hit+f(i+1, hit, max(hp, 0)))
			if hp <= 0 && hit*se >= a[i+1] {
				*p = res
				return res
			}
			hit++
			hp -= fi
		}
	}

	var pr func(int, int, int)
	pr = func(i, preHit, preHp int) {
		if i == n-1 {
			return
		}
		hp := a[i] - preHit*se
		hit := 0
		if preHp > 0 {
			hit = (preHp-1)/se + 1
			hp -= hit * fi
		}
		res := f(i, preHit, preHp)
		for {
			if hit+f(i+1, hit, max(hp, 0)) == res {
				ans = append(ans, pair{i, hit})
				pr(i+1, hit, max(hp, 0))
				return
			}
			hit++
			hp -= fi
		}
	}
	pr(1, 0, 0)

	s := 0
	for _, p := range ans {
		s += p.c
	}
	Fprintln(out, s)
	for _, p := range ans {
		for c := p.c; c > 0; c-- {
			Fprint(out, p.i+1, " ")
		}
	}
}

//func main() { cf6D(os.Stdin, os.Stdout) }
