package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf840C(in io.Reader, out io.Writer) {
	const M = 1_000_000_007
	const mx = 301
	C := [mx][mx]int{}
	for i := range C {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % M
		}
	}

	var n, v, s int
	Fscan(in, &n)
	cnt := map[int]int{}
	perm := 1
	for range n {
		Fscan(in, &v)
		for j := 2; j*j <= v; j++ {
			for v%(j*j) == 0 {
				v /= j * j
			}
		}
		cnt[v]++
		perm = perm * cnt[v] % M
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = perm

	i := 0
	for _, c := range cnt {
		// 一共有 s+1 个空隙，其中有 bad 个坏空隙，s+1-bad 个好空隙
		for bad, dp := range f[i][:s+1] {
			if dp == 0 {
				continue
			}
			for b := 1; b <= c; b++ {
				for j := range min(bad, b) + 1 {
					// 把这 c 个数分成 b 块，方案数为 C[c-1][b-1]
					// 选择其中的 j 块插到坏空隙中，方案数为 C[bad][j]
					// 剩下的 b-j 块插到好空隙中，方案数为 C[s+1-bad][b-j]
					// 坏空隙减少了 j，增加了 c-b
					f[i+1][bad-j+c-b] = (f[i+1][bad-j+c-b] + dp*C[c-1][b-1]%M*C[bad][j]%M*C[s+1-bad][b-j]) % M
				}
			}
		}
		i++
		s += c
	}
	Fprint(out, f[i][0])
}

//func main() { cf840C(os.Stdin, os.Stdout) }
